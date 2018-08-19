package conn_pool

import (
	"errors"
	"fmt"
	"github.com/toolkits/logger"
	"io"
	"sync"
	"time"
)

const (
	MAX_CONN_ERROR = "Maximum connections reached"
	LOG_TAG        = "[conn_pool]"
)

var (
	slowLogEnabled bool  = false
	slowLogLimit   int   = 1000
	ErrMaxConn     error = errors.New(MAX_CONN_ERROR)
)

func EnableSlowLog(enabled bool, limit int) {
	slowLogEnabled = enabled
	slowLogLimit = limit
}

// ConnPool manages the life cycle of connections
type ConnPool struct {
	sync.Mutex

	// New is used to create a new connection when necessary.
	New func() (io.Closer, error)

	// Ping is use to check the conn fetched from pool
	Ping func(io.Closer) error

	Name     string
	MaxConns int
	MaxIdle  int

	conns int
	free  []io.Closer
}

func NewConnPool(name string, max_conns int, max_idle int) *ConnPool {
	return &ConnPool{
		Name:     name,
		MaxConns: max_conns,
		MaxIdle:  max_idle,
	}
}

func (p *ConnPool) Get() (conn io.Closer, err error) {
	if slowLogEnabled {
		start_t := time.Now()
		defer func() {
			end_t := time.Now()
			diff := float64(end_t.UnixNano()-start_t.UnixNano()) / 1000000
			if diff >= float64(slowLogLimit) {
				logger.Debug("%s get conn from pool cost too much, duration: %f ms, pool: %+v", LOG_TAG, diff, p)
			}
		}()
	}
	p.Lock()
	if p.conns >= p.MaxConns && len(p.free) == 0 {
		p.Unlock()
		return nil, ErrMaxConn
	}

	new_conn := false
	if len(p.free) > 0 {
		// return the first free connection in the pool
		conn = p.free[0]
		p.free = p.free[1:]
	} else {
		conn, err = p.New()
		if err != nil {
			p.Unlock()
			return nil, err
		}
		new_conn = true
	}
	p.Unlock()

	err = p.Ping(conn)
	if err != nil {
		p.Lock()
		logger.Error("%s ping conn fail: %v, pool: %+v", LOG_TAG, err, p)
		if !new_conn && p.conns > 0 {
			p.conns -= 1
		}
		p.Unlock()
		conn.Close()
		return nil, err
	}
	if new_conn {
		p.Lock()
		p.conns += 1
		logger.Trace("%s open new conn: %v, pool: %+v", LOG_TAG, conn, p)
		p.Unlock()
	} else {
		logger.Trace("%s get existent conn: %v, pool: %+v", LOG_TAG, conn, p)
	}

	return conn, nil
}

func (p *ConnPool) Release(conn io.Closer) error {
	p.Lock()

	if len(p.free) >= p.MaxIdle {
		logger.Trace("%s auto close conn: %v, pool: %+v", LOG_TAG, conn, p)
		if conn != nil {
			conn.Close()
		}
		p.conns -= 1
	} else {
		p.free = append(p.free, conn)
	}
	logger.Trace("%s release conn: %v, pool: %+v", LOG_TAG, conn, p)

	p.Unlock()
	return nil
}

func (p *ConnPool) CloseClean(conn io.Closer) error {
	if conn != nil {
		conn.Close()
	}
	p.Lock()
	if p.conns > 0 {
		p.conns -= 1
	}
	logger.Trace("%s close_clean conn: %v, pool: %+v", LOG_TAG, conn, p)
	p.Unlock()

	return nil
}

func (p *ConnPool) Destroy() {
	p.Lock()
	defer p.Unlock()

	for _, conn := range p.free {
		if conn != nil {
			logger.Trace("%s destroy conn: %v, pool: %+v", LOG_TAG, conn, p)
			conn.Close()
		}
	}
	p = nil
}

func (p *ConnPool) String() string {
	return fmt.Sprintf("<TcpConnPool Name:%s conns:%d free:%d MaxConns:%d MaxIdle:%d>",
		p.Name, p.conns, len(p.free), p.MaxConns, p.MaxIdle)
}
