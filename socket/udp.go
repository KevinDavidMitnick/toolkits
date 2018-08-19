package socket

import (
	"log"
	"net"
	"time"
)

// 根据 UDP 实现接口
type UDP struct {
	addr     string       // 要连接地址
	port     string       // 要连接端口
	conn     *net.UDPConn // 当前的连接，如果 nil 表示没有连接
	maxRetry int          // 最大重试次数
}

// 建立一个 TCP 对象
func NewUDP(addr string, port string, maxRetry int) *UDP {
	// 创建UDP对象
	udp := new(UDP)
	// 赋值连接地址
	udp.addr = addr
	// 赋值连接端口
	udp.port = port
	// 赋值最大重试此处
	udp.maxRetry = maxRetry
	// 未连接为空
	udp.conn = nil
	// 返回对象
	return udp
}

// 进行连接
func (udp *UDP) connect() error {
	// 创建地址结构
	addr, err := net.ResolveUDPAddr("udp", udp.addr+":"+udp.port)
	if err != nil {
		// 返回错误
		return err
	}
	// 计数器
	var i int = 0
	// 在有效次数内创建连接
	for {
		// 建立TCP连接
		conn, connErr := net.DialUDP("udp", nil, addr)
		if connErr == nil && conn != nil {
			// 设置缓冲区
			conn.SetReadBuffer(1048576)
			conn.SetWriteBuffer(1048576)
			// 将连接保存到对象
			udp.conn = conn
			// 跳出循环,连接成功
			break
		}
		// 判断计数器次数是否达到峰值
		if i > udp.maxRetry {
			return connErr
		}
		// 计数器计数
		i += 1
	}
	// 返回
	return nil
}

// 使用连接
func (udp *UDP) ReadWrite(rw func(conn *net.UDPConn) error) error {
	// 判断连接是否在使用
	for udp.conn != nil {
		log.Printf("connection [%s-%s] in use", udp.addr, udp.port)
		time.Sleep(1 * time.Second)
	}
	// 连接TCP
	connErr := udp.connect()
	// 连接错误则返回
	if connErr != nil {
		return connErr
	}
	// 保证连接的正常关闭
	defer (func() {
		// 断开连接
		closeErr := udp.close()
		if closeErr != nil {
			log.Printf("close the [%s-%s] connection fail", udp.addr, udp.port)
		}
	})()
	// 调用连接方法，传入TCP对象参数，并返回
	return rw(udp.conn)
}

// 断开连接
func (udp *UDP) close() error {
	// 如果连接已经是空
	if udp.conn == nil {
		return nil
	}
	// 断开连接
	closeErr := udp.conn.Close()
	if closeErr != nil {
		return closeErr
	}
	// 清空连接
	udp.conn = nil
	// 返回
	return nil
}
