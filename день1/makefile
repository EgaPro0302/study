include .env
export

migrate-create:
	migrate create -ext sql -dir migrations -seq create_users_table

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

run-code:
	go run ./cmd/server
