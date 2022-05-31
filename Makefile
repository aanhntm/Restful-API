all: network postgres createdb migrateup server

network:
	docker network create restapi_network
	
postgres:
	docker run --name postgres14 --network restapi_network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
	sleep 5

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
	
dockerbuild:
	docker build -t restapi:latest

dockerrun:
	docker run --name restapi -p 8080:8080 --network restapi_network -e "DB_SOURCE=postgresql://root:secret@postgrest14:5432/order?sslmode=disable" restapi:latest

server:
	go run main.go

.PHONY: network postgres creaedb dropdb migrateup migratedown server dockerbuild dockerrun
