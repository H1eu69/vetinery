postgres:
	docker run -d --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root vetinery_db

dropdb:
	docker exec -it postgres dropdb vetinery_db

migrateup:
	migrate --path db/migration -database "postgresql://root:secret@localhost:5432/vetinery_db?sslmode=disable" -verbose up

enterdb:
	docker exec -it postgres sh
