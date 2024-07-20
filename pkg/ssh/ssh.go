package ssh

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
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

func (c *Client) Shell() error {
	client, err := c.connect()
	if err != nil {
		return err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer term.Restore(fd, oldState)

	termWidth, termHeight, err := term.GetSize(fd)
	if err != nil {
		return err
	}

	if err := session.RequestPty("xterm", termHeight, termWidth, ssh.TerminalModes{}); err != nil {
		return err
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	go func() {
		terminalResize := make(chan os.Signal, 1)
		signal.Notify(terminalResize, syscall.SIGWINCH)
		for {
			select {
			case <-terminalResize:
				width, height, _ := term.GetSize(fd)
				session.WindowChange(height, width)
			}
		}
	}()
	if err := session.Shell(); err != nil {
		return err
	}
	if err := session.Wait(); err != nil {
		return err
	}
	return nil
}

func (c *Client) Run(cmd string) (string, error) {
	client, err := c.connect()
	if err != nil {
		return "", err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	result, err := session.Output(cmd)

	return string(result), err
}

func (c *Client) connect() (*ssh.Client, error) {
	var authMethods = []ssh.AuthMethod{
		ssh.Password(c.password),
	}

	keyAuthMethod, err := publicKeyFile(c.keyPath)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		authMethods = append(authMethods, keyAuthMethod)
	}

	config := &ssh.ClientConfig{
		User:            c.user,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%v", c.host, c.port), config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func publicKeyFile(path string) (ssh.AuthMethod, error) {
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
