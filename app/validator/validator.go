package validator

import (
	"regexp"
	"strconv"
	"strings"
)

// IsValidDomain checks if a domain has a valid format
func IsValidDomain(domain string) bool {
	pattern := `^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$`
	expression := regexp.MustCompile(pattern)
	return expression.MatchString(domain)
}

// IsValidURL checks if a URL has a valid format
func IsValidURL(url string) bool {
	pattern := `^((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+(:[0-9]+)?|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)$`
	expression := regexp.MustCompile(pattern)
	return expression.MatchString(url)
}

// ParsePorts checks if a list of ports written in plain text is valid
func ParsePorts(rawPorts string) ([]int, bool) {
	ports := make(map[int]struct{}, 0)

	split := strings.Split(rawPorts, ",")
	for _, rawPort := range split {
		rawPort := strings.TrimSpace(rawPort)
		ranged := strings.Split(rawPort, "-")
		if len(ranged) == 2 {
			ok := parsePortRange(ports, ranged[0], ranged[1])
			if !ok {
				return nil, false
			}
		} else if len(ranged) == 1 {
			port, err := strconv.ParseInt(rawPort, 10, 32)
			if err != nil || port < 1 || port > 65535 {
				return nil, false
			}
			ports[int(port)] = struct{}{}
		} else {
			return nil, false
		}
	}

	keys := make([]int, 0)
	for k := range ports {
		keys = append(keys, k)
	}

	return keys, true
}

func parsePortRange(ports map[int]struct{}, rawStart string, rawEnd string) bool {
	start, err := strconv.ParseInt(rawStart, 10, 32)
	if err != nil || start < 1 || start > 65535 {
		return false
	}

	end, err := strconv.ParseInt(rawEnd, 10, 32)
	if err != nil || end < 1 || end > 65535 || end <= start {
		return false
	}

	for port := start; port <= end; port++ {
		ports[int(port)] = struct{}{}
	}

	return true
}
