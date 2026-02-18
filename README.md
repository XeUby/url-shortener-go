# Mini URL Shortener (Go) — REST API + Tests

A small, interview-friendly URL shortener written in Go.
It exposes a REST API to create short links, redirects by code, and provides basic stats (clicks).
Storage is in-memory, architecture is layered (handler/service/storage), and the project includes unit + HTTP tests.

## Features

- **POST /shorten** — create short URL from a long URL
- **GET /{code}** — redirect to original URL (302)
- **GET /stats/{code}** — get original URL + clicks
- In-memory storage (thread-safe)
- URL validation (`http/https`)
- Collision handling (retry on duplicate code)
- Unit tests for storage/service + HTTP tests for handlers

## Project Structure


url-shortener/
├── main.go
├── handler/
├── service/
├── storage/
└── model/


Layer responsibilities:
- **handler** — HTTP layer (routing, JSON, status codes)
- **service** — business logic (validation, code generation, retries)
- **storage** — persistence (in-memory map + RWMutex)
- **model** — DTO/domain structs

## API

### 1) Create short URL

**Request**
`POST /shorten`

{ "url": "https://google.com" }

Response

{ "short_url": "http://localhost:8080/abc123xy" }
2) Redirect

GET /{code} → 302 Found with Location: <original_url>

3) Stats

GET /stats/{code}

Response

{
  "code": "abc123xy",
  "original_url": "https://google.com",
  "clicks": 3
}
Run
go run .

Server starts on :8080.

BASE_URL (optional)

By default responses use http://localhost:8080.
You can override:

PowerShell

$env:BASE_URL="http://localhost:8080"
go run .
Test
go test ./... -v
Example requests
PowerShell (recommended)

Create:

Invoke-RestMethod -Method Post -Uri "http://localhost:8080/shorten" `
  -ContentType "application/json" `
  -Body '{"url":"https://google.com"}'

Redirect (show headers only, no auto-follow):

$code = "abc123xy"
Invoke-WebRequest -UseBasicParsing -Uri "http://localhost:8080/$code" -MaximumRedirection 0 |
  Select-Object StatusCode, Headers

Stats:

Invoke-RestMethod "http://localhost:8080/stats/$code"
curl
curl -X POST "http://localhost:8080/shorten" \
  -H "Content-Type: application/json" \
  --data-raw '{"url":"https://google.com"}'
Possible Improvements

Persistent storage (PostgreSQL/Redis)

Custom alias support (/shorten with custom_code)

Expiration / TTL

Rate limiting

Observability: structured logs + metrics

Dockerfile + CI (GitHub Actions)
