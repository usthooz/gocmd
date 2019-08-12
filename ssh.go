package gocmd

import (
	"fmt"
	"io/ioutil"
	"time"

	homedir "github.com/mitchellh/go-homedir"

	"golang.org/x/crypto/ssh"
)

type Gocmd struct {
	// ssh client
	Client *ssh.Client
}

type GocmdConfig struct {
	// 连接的Host
	Host string
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
func Config(host, port, username, password, keyPath string, authType ...int) *GocmdConfig {
	config := &GocmdConfig{
		Host:     host,
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

// Connect ssh connect
func Connect(cfg *GocmdConfig) (*Gocmd, error) {
	sshCfg := &ssh.ClientConfig{
		User:            cfg.Username,
		Timeout:         time.Duration(SshAuthTimeout) * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// auth
	if cfg.AuthType == AuthByPassword {
		sshCfg.Auth = []ssh.AuthMethod{ssh.Password(cfg.Password)}
	} else {
		publicKeyAuthFunc, err := publicKeyAuthFunc(cfg.KeyPath)
		if err != nil {
			return nil, err
		}
		sshCfg.Auth = []ssh.AuthMethod{publicKeyAuthFunc}
	}

	// ssh addr
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	// ssh connect
	sshClient, err := ssh.Dial("tcp", addr, sshCfg)
	if err != nil {
		return nil, err
	}

	return &Gocmd{
		Client: sshClient,
	}, nil
}

// publicKeyAuthFunc use ssh key auth
func publicKeyAuthFunc(keyPath string) (ssh.AuthMethod, error) {
	// get key path
	keyPath, err := homedir.Expand(keyPath)
	if err != nil {
		return nil, err
	}
	// key ssh key
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(signer), nil
}
