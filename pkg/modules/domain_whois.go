package modules

// === IMPORTS ===

import (
	"io/ioutil"
	"net"
	"regexp"
	"time"

	"github.com/cosasdepuma/elliot/pkg/utils"
)

// === CONSTANTS ===

const (
	defaultWhoisHost = "whois.iana.org"
	defaultWhoisPort = "43"
)

// === STRUCTURES ===

type whois struct {
	tld     string
	status  bool
	created string
	changed string
	email   []string
	phone   []string
}

// === MODULE METHOD ===

func moduleWhois(domain string, channel *chan *whois) {
	w := &whois{}
	// Get the information
	response := getWhois(domain)
	if len(response) == 0 {
		*channel <- w
		return
	}
	// Parse the response
	w.parse(response)
	// Return the Whois data
	*channel <- w
}

// === STRUCTURES METHODS ===

func (w *whois) parse(data string) {
	// Find essential data
	r := regexp.MustCompile("(?s)domain:\\s+([A-Z]+)\n.+status:\\s+([A-Z]+).+created:\\s+([0-9-]+).+changed:\\s+([0-9-]+)")
	found := r.FindStringSubmatch(data)
	if len(found) == 5 {
		w.tld = found[1]
		w.status = found[2] == "ACTIVE"
		w.created = found[3]
		w.changed = found[4]
	}
	// Find phone data
	r = regexp.MustCompile("(?s)phone:\\s+(?P<phone>[^\n]+)\n")
	phones := r.FindAllStringSubmatch(data, -1)
	if phones != nil {
		for _, phone := range phones {
			w.phone = append(w.phone, string(phone[1]))
		}
		w.phone = utils.FilterDuplicates(w.phone)
	}
	// Find email data
	r = regexp.MustCompile("(?s)e-mail:\\s+(?P<email>[^\n]+)\n")
	mails := r.FindAllStringSubmatch(data, -1)
	if len(mails) > 1 {
		for _, mail := range mails {
			w.email = append(w.email, string(mail[1]))
		}
		w.email = utils.FilterDuplicates(w.email)
	}
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
