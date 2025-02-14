package store

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Initialize a test instance of the storage service.
var testStoreService = &StorageService{}

func TestMain(m *testing.M) {
	// Set test environment variables
	os.Setenv("MYSQL_HOST", "localhost")
	os.Setenv("MYSQL_USER", "root")
	os.Setenv("MYSQL_PASSWORD", "your_password")
	os.Setenv("MYSQL_DATABASE", "url_shortener")

	testStoreService = InitializeStore()
	code := m.Run()

	// Clean up test data
	testStoreService.db.Exec("DELETE FROM url_mappings WHERE user_id = ?",
		"e0dba740-fc4b-4977-872c-d360239e6b1a")

	os.Exit(code)
}

// TestStoreInit verifies that the MySQL connection was properly initialized.
func TestStoreInit(t *testing.T) {
	assert.NotNil(t, testStoreService.db, "Database connection should not be nil")
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
