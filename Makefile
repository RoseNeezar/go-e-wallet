createdb:
	docker exec -it e-wallet-postgres-1 createdb --username=postgres --owner=postgres e_wallet

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres_password@localhost:5432/e_wallet?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres_password@localhost:5432/e_wallet?sslmode=disable" -verbose down

dropdb:
	docker exec -it e-wallet-postgres-1 dropdb e_wallet

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go
	
.PHONY: createdb migrateup migratedown dropdb sqlc test server