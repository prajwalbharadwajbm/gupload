# Gupload - File Upload Service

Gupload is a Go-based file upload service that allows users to securely upload, manage, and store files. The service provides user authentication, file storage management, and fetching files with its metadata.

For routing endpoints, I'm using julienschmidt/httprouter. PostgreSQL as my database. I've written an auth middleware to validate JWT tokens and pass userID down through the request context. Each request has its context downstreamed for any cancellation from the client.
Using environment-based configs makes it easier to store and isolate secrets across environments and deploy to cloud. Log levels are set up as info, debug and error wrapping around zerolog.

## Architecture
![gupload](https://github.com/user-attachments/assets/e62739a4-4c3d-482c-84e5-0214753d4d8a)

## Features

- **User Authentication**: Secure registration and login system using JWT tokens
- **File Upload**: Upload files with automatic storage management
- **Storage Quota**: Each user has a designated storage quota with tracking of used space
- **File Management**: List uploaded files and manage file metadata

## API Endpoints

- **POST /register**: Register a new user
- **POST /login**: Authenticate a user and receive JWT token
- **POST /upload**: Upload a file (requires authentication)
- **GET /storage/remaining**: Check remaining storage quota (requires authentication)
- **GET /files**: List all uploaded files (requires authentication)
- **GET /healthCheck**: Service health check endpoint

## Getting Started

1. Clone the repository
2. Configure your environment variables in .env file
3. Run the server: `go run ./cmd/server/main.go`

## Architecture

The project follows a clean architecture approach:

| Component | Description |
|-----------|-------------|
| **cmd/server** | Main entry point with init for config, DB initiatialization with http server serving httprouter handler with resource limitation added considering any slow request attacks. |
| **internal/config** | Application configuration reads from .env file |
| **internal/db** | Database connections and repositories with db transaction functions |
| **internal/handlers** | HTTP request handlers |
| **internal/middleware** | Authentication and request processing middleware |
| **internal/service** | Business logic services, auth with JWT Generation and Validation. |
| **internal/utils** | Utility functions |
| **internal/validator** | Input validation |
| **internal/interceptor** | Response handling |
| **internal/logger** | Application logging using zerolog |

