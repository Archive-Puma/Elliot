package utils

// === IMPORTS ===

import (
	"regexp"
)

// === PUBLIC METHODS ===

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
