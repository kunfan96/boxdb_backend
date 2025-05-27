package utils

import (
	"fmt"
	"net"
	"time"
)

// 检测端口是否占用
func IsPortUsed(port string) bool {
	addr := fmt.Sprintf(":%s", port)
	conn, err := net.DialTimeout("tcp", addr, time.Second*3)

	// 连接创建成功
	if err == nil {
		// 关闭连接
		conn.Close()

		return true
	}

	return false
}
