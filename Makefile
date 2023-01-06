postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

removepostgres:
	docker stop postgres15 && docker rm postgres15

createdb:
	docker exec -it postgres15 createdb --owner=root --username=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

resetdb:
	yes | make migratedown && make migrateup

test:
	go test -cover -v ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown resetdb
