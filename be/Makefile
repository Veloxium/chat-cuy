postgresinit:
	docker run --name postgres_mydb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=24434 -d postgres

postgres:
	docker exec -it postgres_mydb psql

createdb:
	docker exec -it postgres_mydb createdb --username=root --owner=root go-websocket-sesat

dropdb:
	docker exec -it postgres_mydb dropdb go-websocket-sesat

migrateup:
	migrate -path database/migration -database "postgresql://root:24434@localhost:5432/go-websocket-sesat?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://root:24434@localhost:5432/go-websocket-sesat?sslmode=disable" -verbose down

run:
	go run cmd/main.go

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown run
