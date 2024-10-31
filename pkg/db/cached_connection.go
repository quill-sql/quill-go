package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/quill-sql/quill-go/pkg/models"
	"github.com/quill-sql/quill-go/pkg/utils"
)

// DEFAULT_CACHE_TTL is the default TTL for new cache entries (24 hours)
const DEFAULT_CACHE_TTL time.Duration = 24 * 60 * 60

// CachedConnection represents a cached database connection
type CachedConnection struct {
	DatabaseType string
	Pool         DatabaseConnection
	OrgID        interface{}
	TTL          time.Duration
	Cache        models.RedisMapable
}

// NewCachedConnection creates a new CachedConnection
func NewCachedConnection(databaseType string, config interface{}, cacheConfig models.CacheCredentials) (*CachedConnection, error) {
	pool, err := ConnectToDatabase(databaseType, config)
	if err != nil {
		return nil, err
	}

	ttl := DEFAULT_CACHE_TTL
	if cacheConfig.TTL != nil && *cacheConfig.TTL != time.Duration(0) {
		ttl = *cacheConfig.TTL
	}

	cache, err := getCache(cacheConfig)
	if err != nil {
		return nil, err
	}

	return &CachedConnection{
		DatabaseType: databaseType,
		Pool:         pool,
		TTL:          ttl,
		Cache:        cache,
	}, nil
}

// Query executes a query and caches the result
func (c *CachedConnection) Query(text string) (interface{}, error) {
	if c.Cache == nil {
		return RunQueryByDatabase(c.DatabaseType, c.Pool, text)
	}

	key := fmt.Sprintf("%v:%s", c.OrgID, text)
	cachedResult, err := c.Cache.Get(key)
	if err == nil && cachedResult != "" {
		var result *QueryResults
		err = json.Unmarshal([]byte(cachedResult), &result)
		if err == nil {
			return result, nil
		}
	}

	newResult, err := RunQueryByDatabase(c.DatabaseType, c.Pool, text)
	if err != nil {
		return nil, err
	}

	newResultString, err := json.Marshal(newResult)
	if err != nil {
		return nil, err
	}

	err = c.Cache.Set(key, string(newResultString), time.Duration(c.TTL)*time.Second)
	if err != nil {
		return nil, err
	}

	return newResult, nil
}

// Close closes the database connection
func (c *CachedConnection) Close() error {
	return DisconnectFromDatabase(c.DatabaseType, c.Pool)
}

// getCache configures and returns a cache instance or nil if none could be created
func getCache(config models.CacheCredentials) (models.RedisMapable, error) {
	if config.CacheType == "redis" || config.CacheType == "rediss" {
		redisURL := fmt.Sprintf("%s://%s:%s@%s:%s", config.CacheType, config.Username, config.Password, config.Host, config.Port)
		client := redis.NewClient(&redis.Options{
			Addr: redisURL,
		})
		return &RedisCache{client}, nil
	}
	return nil, nil
}

// RedisCache is a wrapper around *redis.Client that implements the Mapable interface
type RedisCache struct {
	client *redis.Client
}

// Get retrieves a value from the Redis cache
func (c *RedisCache) Get(key string) (string, error) {
	value, err := c.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// Set sets a value in the Redis cache
func (c *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	err := c.client.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return err
	}
	return err
}

func MapQueries(queries []string, targetConnection *CachedConnection) ([]map[string]interface{}, error) {
	var mappedArray []map[string]interface{}
	for _, query := range queries {
		queryResult, err := targetConnection.Query(query)
		if err != nil {
			return nil, err
		}
		row, err := utils.StructToMap(queryResult.(*QueryResults).Rows)
		if err == nil {
			mappedArray = append(mappedArray, row)
		}
	}
	return mappedArray, nil
}
