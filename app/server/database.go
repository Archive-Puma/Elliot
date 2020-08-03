package server

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"

	"github.com/cosasdepuma/elliot/app/config"
	"github.com/cosasdepuma/elliot/app/utils"
)

type DB struct {
	client *redis.Client
}

type Data struct {
	Target     string
	Subdomains []string
}

// NewDatabase TODO Doc
func NewDatabase() *DB {
	return &DB{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.DBHost, config.DBPort),
			Password: config.DBPass,
			DB:       config.DBIndex,
		}),
	}
}

func (db *DB) GetAll() *Data {
	data := new(Data)

	target, err := db.client.Get("target").Result()
	data.Target = target
	if err != nil {
		return nil
	}

	subdomains, err := db.client.LRange("target:subdomains", 0, -1).Result()
	data.Subdomains = subdomains

	log.Println(data)
	return data
}

func (db *DB) SetTarget(domain string) {
	db.client.FlushDB()

	if !utils.IsValidDomain(domain) {
		return
	}

	db.client.Set("target", domain, 0)
}

func (db *DB) SetSubdomains(subdomains []string) {
	db.client.RPush("target:subdomains", subdomains)
}
