postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root order

dropdb:
	docker exec -it postgres14 dropdb order

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/order?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/order?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres creaedb dropdb migrateup migratedown server