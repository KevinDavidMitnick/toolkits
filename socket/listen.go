package socket

import (
	"net"
)

// 监听结构
type Listen struct {
	addr     string // 要监听的地址
	port     string // 要监听的端口
	maxRetry int    // 最大重试次数
}

// 创建监听对象
func NewListen(addr string, port string, maxRetry int) *Listen {
	// 创建对象
	lt := new(Listen)
	// 赋值地址
	lt.addr = addr
	// 赋值端口
	lt.port = port
	// 赋值最大重试次数
	lt.maxRetry = maxRetry
	// 返回对象
	return lt
}

// TCP监听
func (lt *Listen) ListenTCP() (*net.TCPListener, error) {
	// 创建地址结构
	addr, err := net.ResolveTCPAddr("tcp", lt.addr+":"+lt.port)
	if err != nil {
		// 地址结构创建失败
		return nil, err
	}
	// 记录创建监听时的错误
	var lisErr error
	// 记录创建监听时的监听
	var lis *net.TCPListener
	// 在有效次数内创建监听
	for i := 0; i < lt.maxRetry; i++ {
		// 进行监听
		lis, lisErr = net.ListenTCP("tcp", addr)
		if lisErr == nil {
			// 跳出循环
			break
		}
	}
	// 返回
	return lis, lisErr
}

// UDP监听
func (lt *Listen) ListenUDP() (*net.UDPConn, error) {
	// 创建地址结构
	addr, err := net.ResolveUDPAddr("udp", lt.addr+":"+lt.port)
	if err != nil {
		// 地址结构创建失败
		return nil, err
	}
	// 记录连接时的错误
	var lisErr error
	// 记录创建监听时的监听
	var lis *net.UDPConn
	// 在有效次数内创建监听
	for i := 0; i < lt.maxRetry; i++ {
		// 进行监听
		lis, lisErr = net.ListenUDP("udp", addr)
		if lisErr == nil {
			// 跳出循环
			break
		}
	}
	// 返回
	return lis, lisErr
}
