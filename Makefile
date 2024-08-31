# Default .env file
ENV_FILE := devconfig.env

# Check for MODE and set the appropriate .env file
ifeq ($(ENV),prod)
	ENV_FILE := config.env
endif

# Include the chosen .env file
ifneq (,$(wildcard $(ENV_FILE)))
	include $(ENV_FILE)
	export
else
	$(error Environment file $(ENV_FILE) not found)
endif

# Build the application
all: build

build:
	@echo "Building..."


	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@echo "Starting server in development mode..."
	@echo $(ENV_FILE)
	@ENV=local go run cmd/api/main.go

run-prod:
	@echo "Starting server in production mode..."
	@echo $(ENV_FILE)
	ENV=prod @go run cmd/api/main.go


up:
	docker compose up -d

down:
	docker compose down -v

migrate-create:
	migrate create -ext sql -dir internal/database/migrations -seq geocomp_db

migrate-up:
	migrate -path ./internal/database/migrations -database "$(DATABASE_URL)" -verbose up

migrate-down:
	migrate -path ./internal/database/migrations -database "$(DATABASE_URL)" -verbose down

# Test the application
test:
	@echo "Testing..."
	#	@go test ./tests -v -cover
	go test -v -cover ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main


sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination internal/database/mock/store.go  github.com/Dennisblay/ordering-app-server/internal/database/sqlc Store

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean run run-prod watch docker-up docker-down migrate-create up down sqlc test-crud mock
