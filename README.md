# ShortURL-Go

A scalable URL shortener service built with Go, designed for high performance and reliability.

## 🚀 Features

- **Fast URL Shortening**: Generate short, unique URLs for any long URL
- **Quick Redirects**: Efficiently redirect users from short URLs to original destinations
- **Rate Limiting**: Protect against abuse with configurable rate limiting
- **Redis Caching**: Improve performance with Redis-based caching
- **PostgreSQL Storage**: Reliable, persistent storage for all URL mappings
- **RESTful API**: Clean API for creating and managing shortened URLs
- **Scalable Architecture**: Designed to handle high traffic loads

## 🛠️ Tech Stack

- **Backend**: Go (Golang)
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **Caching**: Redis
- **Configuration**: Environment variables via godotenv

## 📋 API Endpoints

| Method | Endpoint        | Description                  |
| ------ | --------------- | ---------------------------- |
| GET    | `/:shortCode` | Redirect to the original URL |
| POST   | `/api/urls`   | Create a new short URL       |
| GET    | `/health`     | Health check endpoint        |

## 🔧 Installation & Setup

### Prerequisites

- Go 1.21+
- PostgreSQL
- Redis

### Local Development

1. Clone the repository:

   ```bash
   git clone https://github.com/bharabhi01/shorturl-go.git
   cd shorturl-go
   ```
2. Install dependencies:

   ```bash
   go mod download
   ```
3. Set up environment variables by creating a `.env` file:

   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=shorturlgo
   REDIS_HOST=localhost
   REDIS_PORT=6379
   REDIS_PASSWORD=
   SERVER_PORT=8080
   BASE_URL=http://localhost:8080
   RATE_LIMIT=50
   ```
4. Run the application:

   ```bash
   go run cmd/server/main.go
   ```
5. The service will be available at `http://localhost:8080`

### Docker Deployment

1. Build and run using Docker Compose:

   ```bash
   docker-compose up -d
   ```
2. The service will be available at `http://localhost:8080`

## 🚀 Deployment

This service can be deployed on various platforms:

### Render

1. Create PostgreSQL and Redis services on Render
2. Create a Web Service pointing to your GitHub repository
3. Set the required environment variables
4. Deploy and enjoy!

Detailed deployment instructions are available in the [deployment guide](DEPLOYMENT.md).

## 📝 Usage Examples

### Creating a Short URL

```bash
curl -X POST http://localhost:8080/api/urls \
  -H "Content-Type: application/json" \
  -d '{"long_url":"https://www.example.com/very/long/url/that/needs/shortening"}'
```

Response:

```json
{
  "short_url": "http://localhost:8080/abc123",
  "long_url": "https://www.example.com/very/long/url/that/needs/shortening"
}
```

### Using a Short URL

Simply visit the short URL in your browser or use curl:

```bash
curl -L http://localhost:8080/abc123
```

## 🧪 Testing

Run the tests with:

```bash
go test ./...
```

## 📈 Performance

The service is designed to handle high loads with minimal latency:

- Redis caching for frequently accessed URLs
- Efficient database queries
- Concurrent request handling

## 🔒 Security Considerations

- Rate limiting to prevent abuse
- Input validation for all API endpoints
- No exposure of database IDs in URLs

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
