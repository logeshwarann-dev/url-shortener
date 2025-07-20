# URL Shortener

A RESTful API service for shortening long URLs, built with Go and PostgreSQL. This project is part of the [roadmap.sh URL Shortening Service](https://roadmap.sh/projects/url-shortening-service) project challenge.

## Features

- Create short URLs from long URLs
- Retrieve original URLs using short codes
- Update existing short URLs
- Delete short URLs
- Track access statistics for each short URL
- URL validation and error handling
- Unique short code generation

## Tech Stack

- **Language**: Go 1.23.0
- **Framework**: Gin
- **Database**: PostgreSQL (GORM)
- **Container Management**: Docker

## Prerequisites

- Docker Desktop
- Git

## App Setup Steps

1. Clone the repository:
```bash
git clone https://github.com/logeshwarann-dev/url-shortener.git
cd url-shortener
```

2. Navigate to app directory under deploy directory:
```bash
docker-compose up -d
```
App & DB docker images will be pulled and containers will be started.
The application will start on `http://localhost:8080`

3. Test the API endpoints using Postman:
```
Import the Postman collection present in `postman` directory
```

4. To Stop the App, Run below command in app directory:
```bash
docker-compose down
```

## API Endpoints

### Create Short URL
```http
POST /shorten
Content-Type: application/json

{
  "url": "https://www.example.com/some/long/url"
}
```

**Response (201 Created):**
```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

### Retrieve Short URL
```http
GET /shorten/{shortCode}
```

**Response (200 OK):**
```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

### Update Short URL
```http
PUT /shorten/{shortCode}
Content-Type: application/json

{
  "url": "https://www.example.com/some/updated/url"
}
```

**Response (200 OK):**
```json
{
  "id": "1",
  "url": "https://www.example.com/some/updated/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:30:00Z"
}
```

### Delete Short URL
```http
DELETE /shorten/{shortCode}
```

**Response: 204 No Content**

### Get Statistics
```http
GET /shorten/{shortCode}/stats
```

**Response (200 OK):**
```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z",
  "accessCount": 10
}
```

## Error Responses

- **400 Bad Request**: Invalid request body or URL validation errors
- **404 Not Found**: Short URL not found
- **500 Internal Server Error**: Server errors

## Database Schema

The application uses a single table to store URL mappings:

```sql
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    access_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Testing

Run the tests using:
```bash
go test ./...
```

## Usage Example

1. Create a short URL:
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.google.com"}'
```

2. Access the short URL:
```bash
curl http://localhost:8080/shorten/abc123
```

3. View statistics:
```bash
curl http://localhost:8080/shorten/abc123/stats
```

## Project Structure

```
url-shortener/
│   .dockerignore
│   .gitignore
│   Dockerfile    # Docker file of the app service
│   go.mod        # Go module dependencies
│   go.sum        # Dependency checksums
│   LICENSE
│   README.md     # Project documentation
│   
├───.github
│   └───workflows
│           ci-actions.yml   # GitHub Actions for CI
│
├───cmd
│   └───api
│           .env
│           server.go     # Application entry point
│
├───deploy                # Local Testing of the application
│   ├───app
│   │       docker-compose.yml   # Docker compose file for the app
│   │
│   └───local-db-test            # Local DB testing using docker compose (not required)
│           docker-compose.yml
│           postgres-cmd.txt
│
├───internal           # Core Application Module
│   ├───api
│   │   ├───handlers   # HTTP request handlers
│   │   │       create_url.go
│   │   │       delete_url.go
│   │   │       retrieve_url.go
│   │   │       update_url.go
│   │   │
│   │   ├───router
│   │   │       router.go
│   │   │
│   │   └───tests
│   │           uniqueId_test.go
│   │
│   ├───models          # Data models 
│   │       url_info.go
│   │
│   └───repository       # Database operations
│       └───postgres
│               db_connect.go
│               db_connect_test.go
│               queries.go
│               query_exec.go
│               query_test.go
│
├───pkg
│   └───utils            # Utility functions
│           env_vars.go
│           integer_handler.go
│           string_handler.go
│
└───postman
        URL Shortener.postman_collection.json
         
```

## License

This project is licensed under the MIT License.