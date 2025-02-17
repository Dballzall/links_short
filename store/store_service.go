package store

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// StorageService wraps the MySQL connection
type StorageService struct {
	db *sql.DB
}

// Global declarations: our store service instance
var (
	storeService = &StorageService{}
)

// InitializeStore sets up the MySQL connection and returns the store service pointer
func InitializeStore() *StorageService {
	// Get MySQL connection details from environment variables
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASSWORD")
	mysqlDB := os.Getenv("MYSQL_DATABASE")

	if mysqlHost == "" {
		mysqlHost = "localhost"
	}
	if mysqlUser == "" {
		mysqlUser = "root"
	}
	if mysqlDB == "" {
		mysqlDB = "url_shortener"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		mysqlUser,
		mysqlPass,
		mysqlHost,
		mysqlDB,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to MySQL: %v", err))
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to MySQL: %v", err))
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	storeService.db = db
	fmt.Printf("Connected to MySQL at %s\n", mysqlHost)
	return storeService
}

/*
SaveUrlMapping persists the mapping between the short URL and the original URL.
The userId is passed here but not used in this simple example.
*/
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	stmt, err := storeService.db.Prepare(`
		INSERT INTO url_mappings (short_url, original_url, user_id) 
		VALUES (?, ?, ?)
	`)
	if err != nil {
		panic(fmt.Sprintf("Failed to prepare statement: %v", err))
	}
	defer stmt.Close()

	_, err = stmt.Exec(shortUrl, originalUrl, userId)
	if err != nil {
		panic(fmt.Sprintf("Failed saving URL mapping: %v", err))
	}
}

/*
RetrieveInitialUrl retrieves the original URL based on the short URL
*/
func RetrieveInitialUrl(shortUrl string) (string, error) {
	var originalUrl string
	err := storeService.db.QueryRow(
		"SELECT original_url FROM url_mappings WHERE short_url = ?",
		shortUrl,
	).Scan(&originalUrl)

	if err == sql.ErrNoRows {
		return "", fmt.Errorf("short URL not found")
	}
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}

func GetRecentUrls(limit int) ([]struct {
	ShortUrl    string    `json:"short_url"`
	OriginalUrl string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}, error) {
	var urls []struct {
		ShortUrl    string    `json:"short_url"`
		OriginalUrl string    `json:"original_url"`
		CreatedAt   time.Time `json:"created_at"`
	}

	rows, err := storeService.db.Query(`
		SELECT short_url, original_url, created_at 
		FROM url_mappings 
		ORDER BY created_at DESC 
		LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var url struct {
			ShortUrl    string    `json:"short_url"`
			OriginalUrl string    `json:"original_url"`
			CreatedAt   time.Time `json:"created_at"`
		}
		if err := rows.Scan(&url.ShortUrl, &url.OriginalUrl, &url.CreatedAt); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}
