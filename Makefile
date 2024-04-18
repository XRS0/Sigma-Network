default:
	docker start sigma-network
	go run cmd/main.go
migrate:
	migrate -path internal/migrations -database "postgres://postgres:postgres@localhost:5432/sigma-network?sslmode=disable" up