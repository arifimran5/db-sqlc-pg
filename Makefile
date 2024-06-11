create_postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

start_postgres:
	docker start postgres12

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root dbSqlcPg

dropdb:
	docker exec -it postgres12 dropdb --username=root --owner=root dbSqlcPg

migrateup:
	goose -dir db/migrations postgres "postgresql://root:secret@localhost:5432/dbSqlcPg" up

migratedown:
	goose -dir db/migrations postgres "postgresql://root:secret@localhost:5432/dbSqlcPg" down

sqlc: 
	sqlc generate

.PHONY: start_postgres createdb dropdb migrateup migratedown sqlc