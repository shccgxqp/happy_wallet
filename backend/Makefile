DB_URL= postgres://root:secret@localhost:5432/happy_wallet?sslmode=disable

postgres:
		docker stop postgres16.2
		docker rm postgres16.2
		docker run --name postgres16.2 --newwork wallet-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16.2-alpine

createdb:
		docker exec -it postgres16.2 createdb --username=root --owner=root happy_wallet

dropdb:
		docker exec -it postgres16.2 dropdb happy_wallet

migrateup:
		migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
		migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
		migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
		migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
		sqlc generate

test:
		go test -v -cover ./... 

server:
		go run main.go

mock:
		mockgen -package mockdb -destination db/mock/store.go github.com/shccgxqp/happy_wallet/backend/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock