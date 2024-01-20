# Create main.go
touch main.go

# Create handlers directory and book_handler.go
mkdir handlers
touch handlers/book_handler.go

# Create models directory and book.go
mkdir models
touch models/book.go

# Create repositories directory and book_repository.go
mkdir repositories
touch repositories/book_repository.go

# Create services directory and book_service.go
mkdir services
touch services/book_service.go

# Create config directory and config.go
mkdir config
touch config/config.go

# Create utils directory and helper.go
mkdir utils
touch utils/helper.go

# Create migrations directory and migration1.go
mkdir migrations
touch migrations/migration1.go

# Create Dockerfile
touch Dockerfile

# Create go.mod and go.sum
go mod init crudapp
