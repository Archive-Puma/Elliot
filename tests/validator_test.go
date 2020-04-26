package test

import (
	"github.com/cosasdepuma/elliot/app/validator"

	"testing"
)

func TestIsValidDomain(t *testing.T) {
	domain := "fsundays.tech"
	if validator.IsValidDomain(domain) {
		t.Log(domain + " passed")
	} else {
		t.Error(domain + " failed")
	}

	domain = "fsundays.tech:8080"
	if !validator.IsValidDomain(domain) {
		t.Log(domain + " passed")
	} else {
		t.Error(domain + " failed")
	}

	domain = "www.fsundays.tech"
	if validator.IsValidDomain(domain) {
		t.Log(domain + " passed")
	} else {
		t.Error(domain + " failed")
	}

	domain = "http://www.fsundays.tech"
	if !validator.IsValidDomain(domain) {
		t.Log(domain + " passed")
	} else {
		t.Error(domain + " failed")
	}

	domain = "fsundays.tech/"
	if !validator.IsValidDomain(domain) {
		t.Log(domain + " passed")
	} else {
		t.Error(domain + " failed")
	}
}

func TestIsValidURL(t *testing.T) {
	URL := "http://fsundays.tech"
	if validator.IsValidURL(URL) {
		t.Log(URL + " passed")
	} else {
		t.Error(URL + " failed")
	}

	URL = "https://fsundays.tech"
	if validator.IsValidURL(URL) {
		t.Log(URL + " passed")
	} else {
		t.Error(URL + " failed")
	}

	URL = "www.fsundays.tech"
	if validator.IsValidURL(URL) {
		t.Log(URL + " passed")
	} else {
		t.Error(URL + " failed")
	}

	URL = "elliot.fsundays.tech"
	if !validator.IsValidURL(URL) {
		t.Log(URL + " passed")
	} else {
		t.Error(URL + " failed")
	}

	URL = "http://www.fsundays.tech"
	if validator.IsValidURL(URL) {
		t.Log(URL + " passed")
	} else {
		t.Error(URL + " failed")
	}

	URL = "http://www.fsundays.tech/"
	if validator.IsValidURL(URL) {
		t.Log(URL + " passed")
	} else {
		t.Error(URL + " failed")
	}

	URL = "http://www.fsundays.tech:80/"
	if validator.IsValidURL(URL) {
		t.Log(URL + " passed")
	} else {
		t.Error(URL + " failed")
	}

	URL = "fsundays.tech/"
	if !validator.IsValidURL(URL) {
		t.Log(URL + " passed")
	} else {
		t.Error(URL + " failed")
	}
}
