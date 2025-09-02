package proxy

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServerStartShutdown(t *testing.T) {
	server, err := NewServer("localhost:0")
	assert.NoError(t, err)

	go func() {
		err := server.Start()
		assert.NoError(t, err)
	}()

	time.Sleep(100 * time.Millisecond)

	// Test connection
	addr := server.Addr().String()
	conn, err := net.Dial("tcp", addr)
	assert.NoError(t, err)

	err = conn.Close()
	if err != nil {
		t.Logf("failed to close connection: %v", err)
	}

	server.Shutdown()

	// After shutdown, listener should be nil or closed
	_, err = net.Dial("tcp", addr)
	assert.Error(t, err)

}
