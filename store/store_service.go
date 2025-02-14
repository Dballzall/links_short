package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// StorageService wraps the Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Global declarations: our store service instance and Redis context
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

// CacheDuration specifies the expiration time for cache entries
const CacheDuration = 6 * time.Hour

// InitializeStore sets up the Redis client and returns the store service pointer.
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // default Redis address
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error initializing Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

/*
SaveUrlMapping persists the mapping between the short URL and the original URL.
The userId is passed here but not used in this simple example.
*/
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key URL | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

/*
RetrieveInitialUrl retrieves the original URL based on the short URL
*/
func RetrieveInitialUrl(shortUrl string) (string, error) {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
