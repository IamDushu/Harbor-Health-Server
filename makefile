postgres: 
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root harbordb

dropdb:
	docker exec -it postgres12 dropdb harbordb

migrateup: 
	migrate -path internal/db/migration -database="postgresql://root:secret@localhost:5432/harbordb?sslmode=disable" -verbose up

migratedown: 
	migrate -path internal/db/migration -database="postgresql://root:secret@localhost:5432/harbordb?sslmode=disable" -verbose down

server:
	go run ./cmd/harbor

.PHONY: postgres createdb dropdb migrateup migratedown