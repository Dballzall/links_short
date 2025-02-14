package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Initialize a test instance of the storage service.
var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

// TestStoreInit verifies that the Redis client was properly initialized.
func TestStoreInit(t *testing.T) {
	assert.NotNil(t, testStoreService.redisClient, "Redis client should not be nil")
}

// TestInsertionAndRetrieval tests the full round-trip:
// saving a URL mapping and then retrieving it.
func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	// Persist the URL mapping.
	SaveUrlMapping(shortURL, initialLink, userUUId)

	// Retrieve the original URL using the short URL.
	retrievedUrl, err := RetrieveInitialUrl(shortURL)
	assert.NoError(t, err)
	assert.Equal(t, initialLink, retrievedUrl, "The retrieved URL should match the original URL")
}
