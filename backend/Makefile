postgres:
		docker run --name postgres16.2 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16.2-alpine

createdb:
		docker exec -it postgres16.2 createdb --username=root --owner=root happy_wallet

dropdb:
		docker exec -it postgres16.2 dropdb happy_wallet

migrateup:
		migrate -path db/migration -database "postgres://root:secret@localhost:5432/happy_wallet?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgres://root:secret@localhost:5432/happy_wallet?sslmode=disable" -verbose down

sqlc:
		sqlc generate

test:
		go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test