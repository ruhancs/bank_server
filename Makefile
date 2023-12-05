usermock:
	mockgen -source=./internal/user/domain/gateway/user_repository.go -destination=internal/user/application/usecase/mock/user_repo_mock.go

accountmock:
	mockgen -source=./internal/user/domain/gateway/user_repository.go -destination=internal/user/application/usecase/mock/user_repo_mock.go

migratedown:
	migrate -path sql/migrations -database "postgresql://postgres:123456@localhost:5432/bank_server?sslmode=disable" -verbose down

migrateup:
	migrate -path sql/migrations -database "postgresql://postgres:123456@localhost:5432/bank_server?sslmode=disable" -verbose up

migrateupdocker:
	migrate -path sql/migrations -database "postgresql://postgres:123456@db:5432/bank_server?sslmode=disable" -verbose up

sqlc:
	sqlc generate

test:
	go test ./...

doc:
	swag init -g cmd/main.go

start:
	docker-compose up -d

.PHONY: migrateup migratedown sqlc accountmock start doc