# Muat semua variabel dari .env ke dalam Makefile
include .env
export $(shell sed 's/=.*//' .env)

# Konfigurasi dasar
MIGRATE = migrate
MIGRATIONS_DIR = migrations
DB_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

# Membuat file migration baru
# Contoh penggunaan: make migrate-create name=create_users_table
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "❌ Error: Harap berikan nama migrasi, contoh: make migrate-create name=create_users_table"; \
		exit 1; \
	fi
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

# Menjalankan semua migration (up)
migrate-up:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Rollback 1 langkah
migrate-down:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

# Rollback semua
migrate-down-all:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

# Menampilkan status migrasi
migrate-version:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

# migrate force (to version before error)
migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "❌ Error: Please include version, i.e: make migrate-force version=2"; \
		exit 1; \
	fi
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $(version)
	@echo "✅ Database version forced to $(version)"

mocks:
	mockgen -source=internal/repositories/user_repo.go -destination=mocks/mock_user_repository.go -package=mocks
	mockgen -source=internal/services/cache_service/cache_service.go -destination=mocks/mock_cache_service.go -package=mocks

seed:
	go run cmd/seed/main.go

run:
# 	go run cmd/main.go
	air

