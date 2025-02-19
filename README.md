# QuickLink - URL Shortener

QuickLink is a URL shortening service built with Go, MySQL, and Docker. It provides a clean web interface and robust API for creating and managing shortened URLs with click tracking capabilities.

## Features

- **Create shortened URLs** with a clean web interface
- **Track click statistics** with visual graphs
- **Recent URLs dashboard**
- **RESTful API endpoints**
- **Click analytics and statistics**
- **Persistent storage with MySQL**
- **Docker containerization** for easy deployment

## Tech Stack

- **Backend:** Go (Gin Framework)
- **Database:** MySQL 8.0
- **Frontend:** HTML/CSS/JavaScript with Chart.js
- **Containerization:** Docker & Docker Compose
- **Database Management:** Adminer

## Environment Variables

Create a `.env` file in the project root. Here’s a sample of what your `.env.example` might look like:

```dotenv
# .env.example

# Application settings
PORT=9808

# MySQL settings
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=quicklink

# Other settings as needed

Set up environment variables by copying the .env.example file to .env and updating the values:
cp .env.example .env

Start the application using Docker Compose:
docker-compose up --build

Access the application:
Web Interface: http://localhost:9808
Adminer (Database Management): http://localhost
