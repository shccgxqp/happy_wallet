DB_URL= postgres://root:secret@localhost:5432/happy_wallet?sslmode=disable

postgres:
		docker stop postgres
		docker rm postgres
		docker run --name postgres --newwork wallet-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16.2-alpine

createdb:
		docker exec -it postgres createdb --username=root --owner=root happy_wallet

dropdb:
		docker exec -it postgres dropdb happy_wallet

migrateup:
		migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
		migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
		migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
		migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
		sqlc generate

test:
		go test -v -cover -short ./... 

server:
		go run main.go

mock:
		mockgen -package mockdb -destination db/mock/store.go github.com/shccgxqp/happy_wallet/backend/db/sqlc Store
		mockgen -package mockwk -destination worker/mock/distributor.go github.com/shccgxqp/happy_wallet/backend/worker TaskDistributor

proto:
		rm -f pb/*.go
		rm -f doc/swagger/*.swagger.json
		protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=happy_wallet \
    proto/*.proto
		statik -src=doc/swagger -dest=doc

evans:
	evans  --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7.2-alpine

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 db_docs db_schema sqlc test server mock proto evans redis new_migration