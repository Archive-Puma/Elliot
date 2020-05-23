package portscanner

import (
	"fmt"
)

type sPort struct {
	protocol string
	number   int
	status   string
	service  string
	banner   string
}

func newPort(protocol string, number int) *sPort {
	port := sPort{protocol: protocol, number: number, status: "close", service: "unknown"}
	port.predictService()
	return &port
}

func (port *sPort) setBanner(banner string) {
	port.banner = banner
}

func (port *sPort) setOpen() {
	port.status = "open"
}

// IsOpen
func (port sPort) isOpen() bool {
	return port.status == "open"
}

func (port sPort) string() string {
	return fmt.Sprintf("%5d/%s\t%-7s\t%-9s\t%s", port.number, port.protocol, port.status, port.service, port.banner)
}

func (port *sPort) format() string {
	return fmt.Sprintf("%d,%s,%s,%s", port.number, port.protocol, port.service, port.banner)
}

func (port *sPort) predictService() {
	if port.protocol == "tcp" {
		if predict, ok := mTCP[port.number]; ok {
			port.service = predict
		}
	} else if port.protocol == "udp" {
		if predict, ok := mUDP[port.number]; ok {
			port.service = predict
		}
	}
}
