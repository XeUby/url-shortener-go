Mini URL Shortener in Go

Interview-friendly REST API for URL shortening written in pure Go (minimal dependencies).

Features
- POST /shorten — create short URL from a long one
- GET /{code} — 302 redirect to the original URL
- GET /stats/{code} — view original URL and click count
- In-memory storage (thread-safe with sync.RWMutex)
- URL validation (http/https only)
- Collision handling (retry on duplicate short codes)
- Unit tests + HTTP integration tests using httptest
- Clean layered architecture: handler → service → storage

Project Structure
url-shortener/
├── main.go
├── handler/
│   └── handler.go
├── service/
│   └── service.go
├── storage/
│   └── storage.go
├── model/
│   └── model.go
└── README.md

Layer Responsibilities
- handler: HTTP routing, request/response parsing, status codes
- service: Business logic (validation, short code generation, retry logic)
- storage: In-memory persistence (map + RWMutex for concurrency safety)
- model: Shared structs (request/response DTOs)

API Endpoints

1. Create Short URL
POST /shorten
Content-Type: application/json

{
  "url": "https://example.com/very/long/path"
}

Response (201 Created)
{
  "short_url": "http://localhost:8080/x7k9p2m"
}

2. Redirect
GET /{code}    e.g. GET /x7k9p2m

→ 302 Found
Location: https://example.com/very/long/path

3. Get Stats
GET /stats/{code}    e.g. GET /stats/x7k9p2m

Response (200 OK)
{
  "code": "x7k9p2m",
  "original_url": "https://example.com/very/long/path",
  "clicks": 5
}

Quick Start

git clone https://github.com/yourusername/url-shortener.git
cd url-shortener
go run .

Server starts on :8080

Optional: set base URL for generated short links
export BASE_URL="https://your-domain.com"
go run .

Running Tests

# Run all tests
go test ./... -v

# Run with coverage
go test ./... -coverprofile=cover.out
go tool cover -html=cover.out

Example Requests (PowerShell)

# Create short URL
Invoke-RestMethod -Method Post -Uri "http://localhost:8080/shorten" `
  -ContentType "application/json" `
  -Body '{"url":"https://google.com"}'

# Check redirect (headers only, no follow)
$code = "x7k9p2m"
Invoke-WebRequest -UseBasicParsing -Uri "http://localhost:8080/$code" -MaximumRedirection 0 |
  Select-Object StatusCode, Headers

# Get stats
Invoke-RestMethod "http://localhost:8080/stats/$code"

Example Requests (curl)

curl -X POST "http://localhost:8080/shorten" \
  -H "Content-Type: application/json" \
  -d '{"url":"https://google.com"}'

curl -i "http://localhost:8080/x7k9p2m"          # redirect
curl "http://localhost:8080/stats/x7k9p2m"       # stats

Possible Improvements / Roadmap
- Persistent storage (PostgreSQL, Redis, SQLite, BadgerDB)
- Custom short codes / vanity URLs
- Link expiration (TTL)
- Rate limiting & abuse protection
- Analytics dashboard
- Prometheus metrics + structured logging (slog/zap)
- Docker + docker-compose
- HTTPS support
- Graceful shutdown
- OpenAPI/Swagger documentation
- CI pipeline (GitHub Actions: test, lint, build)

