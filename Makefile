createDB:
	createdb --username=postgres --owner=postgres go_finance

postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine
	
migrationUp:
	migrate -path db/migration -database "postgresql://postgres:postgres@172.17.0.3:5432/go_finance?sslmode=disable" -verbose up

migrationDrop:
	migrate -path db/migration -database "postgresql://postgres:postgres@172.17.0.3:5432/go_finance?sslmode=disable" -verbose down

.PHONY: createDB postgres