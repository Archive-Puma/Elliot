package database

// === STRUCTURES ===

type DBSchema struct {
	Domain DomainSchema
}

// === Domain ===

type DomainSchema struct {
	Value      string
	IPv4       string
	IPv6       string
	Subdomains []string
	Whois      DomainWhoisSchema
	Web        DomainWebSchema
}

/// === Domains:Whois ===

type DomainWhoisSchema struct {
	TLD     string
	Status  bool
	Created string
	Changed string
	Emails  []string
	Phones  []string
}

// === Domain:Web ===

type DomainWebSchema struct {
	Url       string
	Rating    string
	TLS       DomainWebTLSSchema
	Server    string
	Redirects []string
	Links     []string
	Js        []string
	Software  DomainWebSoftwareSchema
}

type DomainWebTLSSchema struct {
}

type DomainWebSoftwareSchema struct {
}
