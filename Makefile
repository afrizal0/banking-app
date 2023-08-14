postgres:
	docker run --name bankingapp -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=Heilingspeir1 -d postgres

createdb:
	docker exec -it bankingapp createdb --username=postgres --owner=postgres bankingapp

dropdb:
	docker exec -it bankingapp dropdb --username=postgres  bankingapp

migrateup:
	 migrate -path db/migration -database "postgresql://postgres:Heilingspeir1@localhost:5432/bankingapp?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migration -database "postgresql://postgres:Heilingspeir1@localhost:5432/bankingapp?sslmode=disable" -verbose down 

makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

sqlc:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc 