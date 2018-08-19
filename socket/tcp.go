package socket

import (
	"log"
	"net"
	"time"
)

// 根据 TCP 实现接口
type TCP struct {
	addr     string       // 要连接地址
	port     string       // 要连接端口
	conn     *net.TCPConn // 当前的连接，如果 nil 表示没有连接
	maxRetry int          // 最大重试次数
}

// 建立一个 TCP 对象
func NewTCP(addr string, port string, maxRetry int) *TCP {
	// 创建TCP对象
	tcp := new(TCP)
	// 赋值地址
	tcp.addr = addr
	// 赋值端口
	tcp.port = port
	// 赋值最大连接数
	tcp.maxRetry = tcp.maxRetry
	// 未连接状态为空
	tcp.conn = nil
	// 返回对象
	return tcp
}

// 进行连接
func (tcp *TCP) connect() error {
	// 创建地址结构
	addr, err := net.ResolveTCPAddr("tcp", tcp.addr+":"+tcp.port)
	if err != nil {
		// 返回错误
		return err
	}
	// 计数器
	var i int = 0
	// 在有效次数内创建连接
	for {
		// 建立TCP连接
		conn, connErr := net.DialTCP("tcp", nil, addr)
		if connErr == nil && conn != nil {
			// 设置缓冲区
			conn.SetReadBuffer(1048576)
			conn.SetWriteBuffer(1048576)
			// 将连接保存到对象
			tcp.conn = conn
			// 跳出循环,连接成功
			break
		}
		// 判断计数器次数是否达到峰值
		if i > tcp.maxRetry {
			return connErr
		}
		// 计数器计数
		i += 1
	}
	// 连接成功，返回
	return nil
}

// 使用连接
func (tcp *TCP) ReadWrite(rw func(conn *net.TCPConn) error) error {
	// 判断连接是否在使用
	for tcp.conn != nil {
		log.Printf("connection [%s-%s] in use", tcp.addr, tcp.port)
		time.Sleep(1 * time.Second)
	}
	// 连接TCP
	connErr := tcp.connect()
	// 连接错误则返回
	if connErr != nil {
		return connErr
	}
	// 保证连接的正常关闭
	defer (func() {
		// 断开连接
		closeErr := tcp.close()
		if closeErr != nil {
			log.Printf("close the [%s-%s] connection fail", tcp.addr, tcp.port)
		}
	})()
	// 调用连接方法，传入TCP对象参数，并返回
	return rw(tcp.conn)
}

// 断开连接
func (tcp *TCP) close() error {
	// 如果连接已经是空
	if tcp.conn == nil {
		return nil
	}
	// 断开连接
	closeErr := tcp.conn.Close()
	if closeErr != nil {
		return closeErr
	}
	// 清空连接
	tcp.conn = nil
	// 返回
	return nil
}
