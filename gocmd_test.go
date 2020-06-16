package gocmd

import (
	"testing"
)

func TestGoCmd(t *testing.T) {
	// config
	cfg := Config("127.0.0.1", "22", "xc", "123456", "", AuthByPassword)
	// connect
	gcmd, err := Connect(cfg)
	if err != nil {
		t.Fatalf("Connect err-> %v", err)
	}
	// Run
	if err := gcmd.Run("ls"); err != nil {
		t.Fatalf("Run err-> %v", err)
	}
	// CombinedOutput
	output, err := gcmd.CombinedOutput("ls")
	if err != nil {
		t.Fatalf("CombinedOutput err-> %v", err)
	}
	t.Logf("CombinedOutput ls cmd-> %s", output)

	// request pty
	if err := gcmd.RequestPty("vim const.go"); err != nil {
		t.Logf("RequestPty err-> %v", err)
	}
}
