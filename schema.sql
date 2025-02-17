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

-- Add this to your schema.sql
CREATE TABLE IF NOT EXISTS url_clicks (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    short_url VARCHAR(10) NOT NULL,
    clicked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (short_url) REFERENCES url_mappings(short_url)
); 