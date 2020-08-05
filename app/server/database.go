package server

// === IMPORTS ===

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/cosasdepuma/elliot/app/config"
	"github.com/cosasdepuma/elliot/app/modules"
	"github.com/cosasdepuma/elliot/app/utils"
)

// === STRUCTURES ===

// DB is a structure associated with the Redis database
type DB struct {
	Data   *Data
	client *redis.Client
}

// Data is a structure associated with data stored in the Redis database
type Data struct {
	Target     string
	Subdomains []string
	Whois      *modules.SWhois
}

// === PUBLIC METHODS ===

// NewDatabase generates a new structure associated with the Redis database
func NewDatabase() *DB {
	return &DB{
		Data: nil,
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.DBHost, config.DBPort),
			Password: config.DBPass,
			DB:       config.DBIndex,
		}),
	}
}

// === STRUCTURES METHODS ===

// Purge deletes all the information hosted in the database
func (db *DB) Purge() {
	db.client.FlushAll()
}

// === Updaters ===

// UpdateDomainOSINT updates all public information associated with a domain
func (db *DB) UpdateDomainOSINT() {
	db.GetDomain()
	db.GetSubdomains()
	db.GetWhois()
}

// === Getters: Domain ===

// GetDomain returns the target domain
func (db *DB) GetDomain() {
	db.Data = new(Data)
	target, _ := db.client.Get("target").Result()
	db.Data.Target = target
}

// GetSubdomains returns the sub-domains associated with a domain
func (db *DB) GetSubdomains() {
	// Get the subdomains
	subdomains, _ := db.client.LRange("target:subdomains", 0, -1).Result()
	db.Data.Subdomains = subdomains
}

// GetWhois gets the Whois information stored about the domain
func (db *DB) GetWhois() {
	// Check for Whois information about the domain
	msg, err := db.client.Get("target:whois:error").Result()
	if err == nil {
		db.Data.Whois = &modules.SWhois{Error: msg}
		// Get the information
	} else {
		domain, _ := db.client.Get("target:whois:domain").Result()
		status, _ := db.client.Get("target:whois:status").Result()
		created, _ := db.client.Get("target:whois:created").Result()
		changed, _ := db.client.Get("target:whois:changed").Result()
		mail, _ := db.client.LRange("target:whois:mail", 0, -1).Result()
		phone, _ := db.client.LRange("target:whois:phone", 0, -1).Result()
		db.Data.Whois = &modules.SWhois{
			Domain:   domain,
			Status:   status,
			IsActive: status == "ACTIVE",
			Created:  created,
			Changed:  changed,
			Mail:     mail,
			Phone:    phone,
		}
	}
}

// === Setters: Domain ===

// SetDomain stores the target domain in the database
func (db *DB) SetDomain(domain string) {
	db.Purge()
	// Check if it is a valid domain
	if !utils.IsValidDomain(domain) {
		return
	}
	db.client.Set("target", domain, 0)
}

// SetSubdomains Stores the subdomains in the database
func (db *DB) SetSubdomains(subdomains []string) {
	db.client.RPush("target:subdomains", subdomains)
}

// SetWhois stores relevant Whois information about the domain in the database
func (db *DB) SetWhois(whois *modules.SWhois) {
	if len(whois.Error) > 0 {
		db.client.Set("target:whois:error", whois.Error, 0)
	} else {
		db.client.Set("target:whois:domain", whois.Domain, 0)
		db.client.Set("target:whois:status", whois.Status, 0)
		db.client.Set("target:whois:created", whois.Created, 0)
		db.client.Set("target:whois:changed", whois.Changed, 0)
		db.client.RPush("target:whois:mail", whois.Mail)
		db.client.RPush("target:whois:phone", whois.Phone)
	}
}
