package validator

import (
	"regexp"
	"strconv"
	"strings"
)

// IsValidDomain TODO: Doc
func IsValidDomain(domain string) bool {
	pattern := `^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$`
	expression := regexp.MustCompile(pattern)
	return expression.MatchString(domain)
}

// IsValidURL TODO: Doc
func IsValidURL(url string) bool {
	pattern := `^((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+(:[0-9]+)?|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)$`
	expression := regexp.MustCompile(pattern)
	return expression.MatchString(url)
}

// ParsePorts TODO: Doc
func ParsePorts(rawPorts string) ([]int, bool) {
	ports := make(map[int]struct{}, 0)

	split := strings.Split(rawPorts, ",")
	for _, rawPort := range split {
		rawPort := strings.TrimSpace(rawPort)
		ranged := strings.Split(rawPort, "-")
		if len(ranged) == 2 {
			rawStart, rawEnd := strings.TrimSpace(ranged[0]), strings.TrimSpace(ranged[1])

			start, err := strconv.ParseInt(rawStart, 10, 32)
			if err != nil || start < 1 || start > 65535 {
				return nil, false
			}

			end, err := strconv.ParseInt(rawEnd, 10, 32)
			if err != nil || end < 1 || end > 65535 || end <= start {
				return nil, false
			}

			for port := start; port <= end; port++ {
				ports[int(port)] = struct{}{}
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

	keys := make([]int, 0, len(ports))
	for k := range ports {
		keys = append(keys, k)
	}

	return keys, true
}
