package portscanner

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// PortScanner TODO: Doc
type PortScanner struct {
	host    string
	lock    sync.Mutex
	pool    chan bool
	timeout time.Duration
	Results map[int]*Port
}

// NewPortScanner TODO: Doc
func NewPortScanner(host string) *PortScanner {
	return &PortScanner{
		host:    host,
		timeout: time.Second * time.Duration(2),
		lock:    sync.Mutex{}, Results: map[int]*Port{}}
}

// SetTimeout TODO: Doc
func (scanner *PortScanner) SetTimeout(timeout time.Duration) {
	scanner.timeout = timeout
}

func (scanner *PortScanner) withPort(port int) string {
	return fmt.Sprintf("%s:%d", scanner.host, port)
}

// CheckTCPPort TODO: Doc
func (scanner *PortScanner) CheckTCPPort(port int) *Port {
	scannedPort := newPort("tcp", port)

	addr, err := net.ResolveTCPAddr("tcp4", scanner.withPort(port))
	if err == nil {
		conn, err := net.DialTimeout("tcp", addr.String(), scanner.timeout)
		if err == nil {
			scannedPort.setOpen()
			buffer := make([]byte, 30)
			conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(2)))
			bytesLen, err := conn.Read(buffer)
			if err == nil {
				banner := string(buffer[:bytesLen])
				scannedPort.setBanner(strings.SplitN(banner, "\n", 1)[0])
			}

		}

	}

	return scannedPort
}
