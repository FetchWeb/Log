package test

import (
	"errors"
	"os"
	"testing"

	log "github.com/FetchWeb/Log"
)

// TestLog writes test logs to Stdout.
func TestLog(t *testing.T) {
	if err := log.Startup(os.Stdout, "TEST", true); err != nil {
		t.Fatalf("Failed to startup logger: %v", err)
	}

	log.Info("Test Info message")
	log.Infof("Test Infof %v", errors.New("message"))

	log.Error("Test Error message")
	log.Errorf("Test Errorf %v", errors.New("message"))

	log.Debug("Test Debug message")
	log.Debugf("Test Debugf %v", errors.New("message"))
}
