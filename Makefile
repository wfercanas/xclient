DB_URL=postgres://fernando@localhost:5432/xclient?sslmode=disable
MIGRATIONS_DIR=migrations

.PHONY: migrate-up migrate-down migrate-force migrate-new

# Apply all pending migrations
migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Roll back the last migration (⚠️ safe only in dev)
migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

# Force set migration version (use carefully)
migrate-force:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $(VERSION)

# Create a new migration file
migrate-new:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)
