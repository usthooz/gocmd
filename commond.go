package gocmd

import (
	"fmt"
)

// Run non output
func (gcmd *Gocmd) Run(cmd string) error {
	if gcmd.Client == nil {
		return fmt.Errorf("gocmd-> client is nil...")
	}
	// new session
	session, err := gcmd.Client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	return session.Run(cmd)
}

// CombinedOutput exec command and has output
func (gcmd *Gocmd) CombinedOutput(cmd string) (string, error) {
	if gcmd.Client == nil {
		return "", fmt.Errorf("gocmd-> client is nil...")
	}
	// new session
	session, err := gcmd.Client.NewSession()
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
