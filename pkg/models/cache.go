package models

import "time"

// Mapable is an interface that defines the methods for cache operations.
type Mapable interface {
	Get(key string) (string, error)
	Set(key string, value string, ttl *time.Duration) (string, error)
}

// CacheCredentials holds the credentials and configuration for a cache.
type CacheCredentials struct {
	Username  string
	Password  string
	Host      string
	Port      string
	CacheType string
	TTL       *time.Duration
}
