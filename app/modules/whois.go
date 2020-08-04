package modules

// === IMPORTS ===

import (
	"io/ioutil"
	"net"
	"regexp"
	"time"
)

// === CONSTANTS ===

const (
	defaultWhoisHost    = "whois.iana.org"
	defaultWhoisPort    = "43"
	defaultErrorMessage = "No whois information available"
)

// === STRUCTURES ===

// SWhois is a structure to store relevant data about a domain following the standard IANA response
type SWhois struct {
	Domain   string
	Status   string
	IsActive bool
	Created  string
	Changed  string
	Mail     []string
	Phone    []string
	Error    string
}

// === STRUCTURES METHODS ===

func (whois *SWhois) parse(data string) {
	// Find essential data
	r := regexp.MustCompile("(?s)domain:\\s+(?P<domain>[A-Z]+)\n.+status:\\s+(?P<status>[A-Z]+).+created:\\s+(?P<created>[0-9-]+).+changed:\\s+(?P<changed>[0-9-]+)")
	found := r.FindStringSubmatch(data)
	if len(found) == 5 {
		whois.Domain = found[1]
		whois.Status = found[2]
		whois.Created = found[3]
		whois.Changed = found[4]
	}
	// Find phone data
	r = regexp.MustCompile("(?s)phone:\\s+(?P<phone>[^\n]+)\n")
	phones := r.FindAllStringSubmatch(data, -1)
	if phones != nil {
		for _, phone := range phones {
			whois.Phone = append(whois.Phone, string(phone[1]))
		}
	}
	// Find email data
	r = regexp.MustCompile("(?s)e-mail:\\s+(?P<email>[^\n]+)\n")
	mails := r.FindAllStringSubmatch(data, -1)
	if len(mails) > 1 {
		for _, mail := range mails {
			whois.Mail = append(whois.Mail, string(mail[1]))
		}
	}
}

// === PUBLIC METHODS ===

// Whois is a concurrent method for obtaining information about a domain.
func Whois(domain string, channel *chan *SWhois) {
	whois := &SWhois{}
	// Get the information
	response := getWhois(domain)
	if len(response) == 0 {
		whois.Error = defaultErrorMessage
		*channel <- whois
		return
	}
	// Parse the response
	whois.parse(response)
	// Return the Whois data
	*channel <- whois
}

// === PRIVATE METHODS ===

func getWhois(domain string) string {
	// Connect to the Whois service
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(defaultWhoisHost, defaultWhoisPort), time.Second*30)
	if err != nil {
		return ""
	}
	defer conn.Close()
	// Send the domain consult
	_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 30))
	_, err = conn.Write([]byte(domain + "\r\n"))
	if err != nil {
		return ""
	}
	// Read the response
	_ = conn.SetReadDeadline(time.Now().Add(time.Second * 30))
	buffer, err := ioutil.ReadAll(conn)
	if err != nil {
		return ""
	}
	return string(buffer)
}
