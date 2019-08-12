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
	// 连接用户名
	Username string
	// 连接密码
	Password string
	// 认证类型 [1: password 2: ssh key]
	AuthType int
	// ~/.ssh/id_rsa
	KeyPath string
}

// ReloadGocmdConfig 初始化ssh config
func Config(ip, port, username, password, keyPath string, authType ...int) *GocmdConfig {
	config := &GocmdConfig{
		Ip:       ip,
		Username: username,
		Password: password,
		KeyPath:  keyPath,
	}
	// default 22
	if len(port) == 0 {
		config.Port = DefaultSshPort
	} else {
		config.Port = port
	}
	// default password
	if len(authType) == 0 {
		config.AuthType = DefaultAuthType
	} else {
		config.AuthType = authType[0]
	}
	return config
}
