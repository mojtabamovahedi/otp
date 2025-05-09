# OTP Service

This project is an OTP (One-Time Password) service built using Go. It provides APIs for generating and verifying OTPs, leveraging Redis for caching and Gin for HTTP handling. The service is designed to be modular, scalable, and easy to integrate into other systems.

## Features

- **Generate OTP**: Generate a secure OTP for a given phone number.
- **Verify OTP**: Verify the OTP for a phone number.
- **Rate Limiting**: Prevent abuse with request rate limiting.
- **Redis Integration**: Use Redis for caching OTPs with expiration policies.
- **Modular Design**: Easily extendable and maintainable codebase.

## Project Structure

- **`api/handler/http`**: Contains HTTP handlers and middleware for the OTP service.
- **`app`**: Application setup and initialization.
- **`config`**: Configuration management for the service.
- **`internal/repository`**: Repository layer for OTP operations.
- **`internal/service`**: Business logic for OTP generation and verification.
- **`pkg/logger`**: Logging utilities using Zap.
- **`pkg/otp`**: OTP generation logic.
- **`pkg/redis`**: Redis connection and caching utilities.

## Prerequisites

- Go 1.24.2 or later
- Docker (for running Redis)

## Getting Started

1. **Clone the repository**:
   ```bash
   git clone https://github.com/mojtabamovahedi/otp.git
   cd otp
   ```

2. **Start Redis**:
   Use the provided `docker-compose.yml` to start a Redis instance:
   ```bash
   docker-compose up -d
   ```

3. **Run the application**:
   ```bash
   go run cmd/main.go
   ```

4. **API Endpoints**:
   - `POST /api/v1/otp/generate`: Generate an OTP for a phone number.
   - `POST /api/v1/otp/verify`: Verify an OTP for a phone number.

## Configuration

The application uses a `config.yaml` file for configuration. Example:

```yaml
server:
  http_port: 8080

redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0
```

## Testing

Run the test suite using:
```bash
go test ./...
```

## Dependencies

- [Gin](https://github.com/gin-gonic/gin): HTTP web framework.
- [Redis](https://github.com/redis/go-redis): Redis client for Go.
- [Zap](https://github.com/uber-go/zap): Logging library.
- [Testify](https://github.com/stretchr/testify): Testing utilities.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.