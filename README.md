# Backend API

A RESTful API backend built with Go, Gin framework, and MySQL database. This project provides authentication and user management functionality with JWT token-based security.

## Features

- 🔐 JWT Authentication
- 👤 User Management (CRUD operations)
- 🔒 Password hashing with bcrypt
- 🌐 CORS support
- ✅ Request validation
- 🗄️ MySQL database with GORM
- 🔄 Environment-based configuration
- 🚀 Hot reload with Air

## Tech Stack

- **Language**: Go 1.25.4
- **Web Framework**: Gin
- **ORM**: GORM
- **Database**: MySQL
- **Authentication**: JWT (golang-jwt/jwt)
- **Password Hashing**: bcrypt
- **Environment**: godotenv
- **Dev Tool**: Air (hot reload)

## Project Structure

```
.
├── config/          # Configuration and environment setup
├── controller/      # HTTP request handlers
├── database/        # Database connection and initialization
├── helpers/         # Utility functions (JWT, password hashing, converters)
├── middlewares/     # Middleware functions (authentication)
├── models/          # Database models
├── routes/          # API route definitions
├── structs/         # Request/Response structures
├── version/         # Application version management
├── .env.example     # Example environment variables
├── .air.toml        # Air configuration for hot reload
├── go.mod           # Go module dependencies
└── main.go          # Application entry point
```

## Prerequisites

- Go 1.25.4 or higher
- MySQL 5.7 or higher
- Air (for development with hot reload)

## Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd backend-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Setup environment variables**
   ```bash
   cp .env.example .env
   ```
   
   Edit `.env` with your configuration:
   ```env
   APP_PORT=3000
   
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=your_password
   DB_NAME=your_database
   
   JWT_SECRET=your_jwt_secret_key
   ```

4. **Create database**
   ```sql
   CREATE DATABASE your_database;
   ```

5. **Install Air for hot reload (optional)**
   ```bash
   go install github.com/air-verse/air@latest
   ```

## Running the Application

### Development mode (with hot reload)
```bash
air
```

### Production mode
```bash
go run main.go
```

### Build and run
```bash
go build -o backend-api
./backend-api
```

The server will start on `http://localhost:3000` (or the port specified in `.env`)

## API Endpoints

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Health check and version info |
| POST | `/api/register` | Register a new user |
| POST | `/api/login` | Login and get JWT token |

### Protected Endpoints (Require Authentication)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/users` | Get all users |
| POST | `/api/users` | Create a new user |
| GET | `/api/users/:id` | Get user by ID |
| PUT | `/api/users/:id` | Update user by ID |
| DELETE | `/api/users/:id` | Delete user by ID |

## API Usage Examples

### Register
```bash
curl -X POST http://localhost:3000/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "username": "johndoe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Login
```bash
curl -X POST http://localhost:3000/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "password": "password123"
  }'
```

Response:
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "id": 1,
    "name": "John Doe",
    "username": "johndoe",
    "email": "john@example.com",
    "created_at": "2025-11-30 10:00:00",
    "updated_at": "2025-11-30 10:00:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Get All Users (Protected)
```bash
curl -X GET http://localhost:3000/api/users \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Create User (Protected)
```bash
curl -X POST http://localhost:3000/api/users \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "username": "janedoe",
    "email": "jane@example.com",
    "password": "password123"
  }'
```

### Update User (Protected)
```bash
curl -X PUT http://localhost:3000/api/users/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Smith",
    "username": "janesmith",
    "email": "jane.smith@example.com"
  }'
```

### Delete User (Protected)
```bash
curl -X DELETE http://localhost:3000/api/users/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Authentication

This API uses JWT (JSON Web Tokens) for authentication. After successful login, you'll receive a token that must be included in the `Authorization` header for protected endpoints:

```
Authorization: Bearer YOUR_JWT_TOKEN
```

Tokens expire after 360 minutes (6 hours) by default.

## Response Format

### Success Response
```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

### Error Response
```json
{
  "success": false,
  "message": "Error message",
  "errors": { ... }
}
```

## Development

### Database Migrations
The application automatically creates tables on startup using GORM's AutoMigrate feature.

### Adding New Routes
1. Define the handler in `controller/`
2. Add the route in `routes/routes.go`
3. Add middleware if needed (e.g., `middlewares.AuthMiddleware()`)

### Environment Variables
- `APP_PORT`: Server port (default: 3000)
- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `JWT_SECRET`: Secret key for JWT signing

## Security Features

- Password hashing with bcrypt
- JWT token authentication
- CORS configuration
- Input validation
- Protected routes with middleware

## License

This project is private and proprietary.

## Author

ibnufth
