postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=rootroot -d postgres

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root simple-bank

dropdb:
	docker exec -it postgres13 dropdb simple-bank

migrateup:
	migrate -path db/migration -database "postgresql://root:rootroot@localhost:5432/simple-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:rootroot@localhost:5432/simple-bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server: 
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/florian-nguyen/golang-training/tech-school/simple-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock