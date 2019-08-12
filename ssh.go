package gocmd

import (
	"golang.org/x/crypto/ssh"
)

type Gocmd struct {
	// ssh client
	client *ssh.Client
	// ssh client session
	session *ssh.Session
}

type GocmdConfig struct {
	// 连接的IP
	Ip string
	// 连接端口号
	Port string
	// 连接密码
	Password string
	// 认证类型 [1: password 2: ssh key]
	AuthType int
	// ~/.ssh/id_rsa
	keyPath string
}
