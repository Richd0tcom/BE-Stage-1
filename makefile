
postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Madara123 -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root hngx
dropdb:
	docker exec -it postgres15 dropdb --username=root  hngx
migrateup:
	migrate -path db/migrations -database "postgresql://root:Madara123@localhost:5432/hngx?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:Madara123@localhost:5432/hngx?sslmode=disable" -verbose down



.PHONY: sqlc dropdb createdb migrateup migratedown