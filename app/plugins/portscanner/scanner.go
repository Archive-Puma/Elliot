package portscanner

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

type sPortScanner struct {
	host    string
	lock    sync.Mutex
	pool    chan bool
	timeout time.Duration
	Results map[int]*sPort
}

func newPortScanner(host string) *sPortScanner {
	return &sPortScanner{
		host:    host,
		timeout: time.Second * time.Duration(2),
		lock:    sync.Mutex{}, Results: map[int]*sPort{}}
}

func (scanner *sPortScanner) setTimeout(timeout time.Duration) {
	scanner.timeout = timeout
}

func (scanner *sPortScanner) withPort(port int) string {
	return fmt.Sprintf("%s:%d", scanner.host, port)
}

func (scanner *sPortScanner) checkTCPPort(port int) *sPort {
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
				scannedPort.setBanner(strings.TrimSpace(strings.SplitN(banner, "\n", 1)[0]))
			}

		}

	}

	return scannedPort
}
