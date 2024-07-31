package ssh

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

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
