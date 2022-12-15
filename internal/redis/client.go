package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Note that in a real world usage, the cache duration shouldn't have
// an expiration time, an LRU policy config should be set where the
// values that are retrieved less often are purged automatically from
// the cache and stored back in RDBMS whenever the cache is full
const CacheDuration = 6 * time.Hour

var redisClient *redis.Client

// Initialize and return a new Redis client instance
func InitializeClient(host string, port int, password string, db int) *redis.Client {

	// Initialize new Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})

	// Get the pong message from Redis
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	return redisClient
}

// Store the mapping of the short URL to the original URL in Redis
func StoreUrlMapping(shortUrl string, originalUrl string) {
	err := redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to save URL mapping | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

// Retrive original URL given the short URL
func RetrieveOriginalUrl(shortUrl string) string {
	result, err := redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveOriginalUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
