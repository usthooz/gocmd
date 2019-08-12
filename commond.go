package gocmd

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// checkClientNewSession
func (gcmd *Gocmd) checkClientNewSession() (*ssh.Session, error) {
	if gcmd.Client == nil {
		return nil, fmt.Errorf("gocmd-> client is nil...")
	}
	// new session
	return gcmd.Client.NewSession()
}

// Run non output
func (gcmd *Gocmd) Run(cmd string) error {
	session, err := gcmd.checkClientNewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	return session.Run(cmd)
}

// CombinedOutput exec command and has output
func (gcmd *Gocmd) CombinedOutput(cmd string) (string, error) {
	session, err := gcmd.checkClientNewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// exec
	buf, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// RequestPty terminal pty
func (gcmd *Gocmd) RequestPty(cmd string) error {
	session, err := gcmd.checkClientNewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// /dev/stdin flag
	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer terminal.Restore(fd, oldState)

	// stdin and stdout
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	// terminal whidth anf height
	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		return err
	}

	// TerminalModes
	modes := ssh.TerminalModes{
		// enable echo
		ssh.ECHO: 1,
		// input speed = 14.4kbaud
		ssh.TTY_OP_ISPEED: 14400,
		// output speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
		return err
	}
	return session.Run(cmd)
}
