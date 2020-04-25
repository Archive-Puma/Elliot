package validator

import (
	"regexp"
)

// IsValidDomain TODO: Doc
func IsValidDomain(domain string) bool {
	pattern := `^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$`
	expression := regexp.MustCompile(pattern)
	return expression.MatchString(domain)
}
