package test

import (
	"github.com/mrrobotproject/validator"
	"testing"
)

func TestIsValidDomain(t *testing.T) {
	domain := "fsundays.tech"
	if validator.IsValidDomain(domain) { t.Log(domain + " passed")
	} else { t.Error(domain + " failed")}

	domain = "www.fsundays.tech"
	if validator.IsValidDomain(domain) { t.Log(domain + " passed")
	} else { t.Error(domain + " failed")}

	domain = "http://www.fsundays.tech"
	if !validator.IsValidDomain(domain) { t.Log(domain + " passed")
	} else { t.Error(domain + " failed")}

	domain = "fsundays.tech/"
	if !validator.IsValidDomain(domain) { t.Log(domain + " passed")
	} else { t.Error(domain + " failed")}
}