postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root quote_maker

dropdb:
	docker exec -it postgres dropdb quote_maker

migrateup:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/quote_maker?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/quote_maker?sslmode=disable" -verbose down	

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown
