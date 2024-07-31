package ssh

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

type Client struct {
	user     string
	password string
	host     string
	port     int
	keyPath  string
}

func NewClient(user string, password string, host string, port int, keyPath string) *Client {
	return &Client{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		keyPath:  keyPath,
	}
}

func (c *Client) connect() (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User:            c.user,
		Auth:            c.createAuthMethods(),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%v", c.host, c.port), config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) createAuthMethods() []ssh.AuthMethod {
	var authMethods = []ssh.AuthMethod{
		ssh.Password(c.password),
	}
	keyAuthMethod, err := loadPublicKey(c.keyPath)
	if err != nil {
		return authMethods
	}
	authMethods = append(authMethods, keyAuthMethod)
	return authMethods
}

func loadPublicKey(path string) (ssh.AuthMethod, error) {
	key, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(signer), nil
}
