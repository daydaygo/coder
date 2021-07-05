// 高并发服务gogc优化&udpPool优化: https://mp.weixin.qq.com/s/EuJ3Pw0s24Nr1h2edn5Sgg
// 其他优化: 日志精简 上报合并
package main

import (
	"fmt"
	"net"
	"sync"
)

// 服务内存不大, 请求量高, 搞个预留, 大大减少gc: qps 8w->14w
var stub = make([]byte, 1024*1024*1024) // 启动hang1-2s, 影响不大, 启动不会那么快加流量
func main() {
	stub[0]= 1 // 价格应用, 防止优化
	fmt.Println("remain a not used byte ", len(stub))
}

var udpPool = sync.Map{} // mem leak: 几w addr 也就几M内存

type poolResWrapper struct {
	net.PacketConn
	err error
}

func UdpPool(network string, addr *net.UDPAddr, reqData []byte) ([]byte, error) {
	var conn net.PacketConn
	var pool *sync.Pool // 并发? 串包? qps 14w->17w
	poolInterface, _ := udpPool.LoadOrStore(addr.String(), &sync.Pool{New: func() interface{} {
		udpConn, err := net.DialUDP(network, nil, addr)
		return &poolResWrapper{udpConn, err}
	}})
	pool = poolInterface.(*sync.Pool)
	pw := pool.Get().(*poolResWrapper)
	conn, err := pw.PacketConn, pw.err
	if err != nil {
		return nil, err
	}
	udpconn := conn.(*net.UDPConn)
	_, err = udpconn.Write(reqData)
	if err != nil {
		_ = pw.PacketConn.Close()
		return nil, err
	}
	pool.Put(pw)
	return reqData,nil
}