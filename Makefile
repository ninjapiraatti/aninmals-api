include .env
DB_URL=postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

migration_up: 
	migrate -path database/migration/ -database $(DB_URL) -verbose up
migration_down: 
	migrate -path database/migration/ -database $(DB_URL) -verbose down