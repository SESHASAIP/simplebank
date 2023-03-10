postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simpleBank

dropdb:
	docker exec -it postgres12 dropdb --username=root --owner=root simpleBank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable" -verbose up


migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable" -verbose up 1	

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable" -verbose down 1


sqlc:
	sqlc generate

test:
	go test

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go simplebank/db/store Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc server mock


