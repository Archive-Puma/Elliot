package portscanner

import "fmt"

// Port TODO: Doc
type Port struct {
	protocol string
	number   int
	status   string
	service  string
	banner   string
}

func newPort(protocol string, number int) *Port {
	port := Port{protocol: protocol, number: number, status: "close", service: "unknown"}
	port.predictService()
	return &port
}

func (port *Port) setBanner(banner string) {
	port.banner = banner
}

func (port *Port) setOpen() {
	port.status = "open"
}

// IsOpen TODO: Doc
func (port Port) IsOpen() bool {
	return port.status == "open"
}

func (port Port) String() string {
	return fmt.Sprintf("%5d/%s\t%-7s\t%-9s\t%s", port.number, port.protocol, port.status, port.service, port.banner)
}

func (port *Port) predictService() {
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
