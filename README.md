# GoReddit

A Reddit-like application built with Go, providing a REST API for managing threads, posts, and comments.

## Features

- 📝 Create and manage discussion threads
- 💬 Create posts within threads
- 🗨️ Add comments to posts
- 👍 Voting system for posts and comments
- 🔄 RESTful API endpoints
- 🗄️ PostgreSQL database integration

## Architecture

This application follows a clean architecture pattern with:

- **Domain Models**: Core entities (Thread, Post, Comment) defined in `goreddit.go`
- **Store Layer**: Database operations with interfaces for easy testing
- **Web Layer**: HTTP handlers using Chi router
- **Database**: PostgreSQL with migrations

## Tech Stack

- **Language**: Go 1.25.2
- **Web Framework**: Chi v5 (HTTP router)
- **Database**: PostgreSQL
- **Database Driver**: lib/pq
- **Query Builder**: sqlx
- **ID Generation**: Google UUID

## Project Structure

```
├── cmd/
│   └── main.go           # Application entry point
├── migrations/
│   ├── 01_create_tables.up.sql    # Database schema
│   └── 01_create_tables.down.sql  # Rollback migrations
├── stores/
│   ├── store.go          # Store interface implementations
│   ├── thread_store.go   # Thread database operations
│   ├── post_store.go     # Post database operations
│   └── comment_store.go  # Comment database operations
├── web/
│   └── handler.go        # HTTP request handlers
├── goreddit.go           # Domain models and interfaces
├── go.mod               # Go module dependencies
└── Makefile             # Build and development commands
```

## Prerequisites

- Go 1.25.2 or later
- PostgreSQL
- Docker (optional, for running PostgreSQL)
- [golang-migrate](https://github.com/golang-migrate/migrate) (for database migrations)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/pnlinh/goreddit.git
cd goreddit
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Start PostgreSQL database

Using Docker:
```bash
make db
```

Or start your own PostgreSQL instance with:
- Host: localhost
- Port: 5432
- Database: postgres
- Username: postgres
- Password: secret

### 4. Run database migrations

```bash
make migrate.up
```

### 5. Start the application

```bash
go run cmd/main.go
```

The server will start on `http://localhost:8081`

## API Endpoints

### Threads

- `GET /threads` - List all threads
- `POST /threads` - Create a new thread
  - Form parameters: `title`, `description`
- `DELETE /threads/{id}` - Delete a thread by ID

### Testing the API

A [Yaak](https://yaak.app) REST client collection is provided in `rest/yaak.goreddit.json` with pre-configured requests for testing all API endpoints.

To use the collection:

1. Install [Yaak REST Client](https://yaak.app)
2. Import the collection file: `File > Import > rest/yaak.goreddit.json`
3. The collection includes sample requests for:
   - Creating threads
   - Listing threads
   - Deleting threads

The collection is configured to work with the default server running on `http://localhost:8081`.

### Database Schema

The application uses three main tables:

- **threads**: Discussion topics
  - `id` (UUID, Primary Key)
  - `title` (Text)
  - `description` (Text)

- **posts**: Posts within threads
  - `id` (UUID, Primary Key)
  - `thread_id` (UUID, Foreign Key)
  - `title` (Text)
  - `content` (Text)
  - `votes` (Integer)

- **comments**: Comments on posts
  - `id` (UUID, Primary Key)
  - `post_id` (UUID, Foreign Key)
  - `content` (Text)
  - `votes` (Integer)

## Development

### Available Make Commands

```bash
make db           # Run PostgreSQL in Docker
make adminer      # Run Adminer database management tool
make migrate.up   # Apply database migrations
make migrate.down # Rollback database migrations
make dev          # Run with live reload (requires air)
```

### Live Reload Development

For development with live reload, install [air](https://github.com/cosmtrek/air):

```bash
go install github.com/cosmtrek/air@latest
```

Then run:
```bash
make dev
```

## Testing

Run the test suite with:

```bash
go test ./...
```

## Database Management

You can use Adminer to manage your database:

```bash
make adminer
```

Then visit `http://localhost:8080` and connect with:
- System: PostgreSQL
- Server: localhost
- Username: postgres
- Password: secret
- Database: postgres

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).

## TODO

- [ ] Add authentication and user management
- [ ] Implement post and comment endpoints
- [ ] Add voting functionality
- [ ] Add pagination for list endpoints
- [ ] Add input validation and error handling
- [ ] Add comprehensive test coverage
- [ ] Add API documentation (OpenAPI/Swagger)
- [ ] Add Docker support for the application
- [ ] Add CI/CD pipeline
