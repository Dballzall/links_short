-- Create the database if it doesn't exist
CREATE DATABASE IF NOT EXISTS url_shortener;

USE url_shortener;

-- Create the url_mappings table
CREATE TABLE IF NOT EXISTS url_mappings (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    short_url VARCHAR(10) NOT NULL UNIQUE,
    original_url TEXT NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_short_url (short_url)
); 