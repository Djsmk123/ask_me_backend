postgres:
	docker run --name postgress-db -p 5432:5432 -e POSTGRES_PASSWORD=1235789 postgres

createdb:
	docker exec -it postgress-db createdb --username=postgres --owner=postgres ask_me

dropdb:
	docker exec -it postgress-db dropdb ask_me -U postgres

mgup:
	migrate -path db/migration -database "postgres://postgres:1235789@localhost:5432/ask_me?sslmode=disable" up
mgup1:
	migrate -path db/migration -database "postgres://postgres:1235789@localhost:5432/ask_me?sslmode=disable" up 1
mgup2:
	migrate -path db/migration -database "postgres://postgres:1235789@localhost:5432/ask_me?sslmode=disable" up 2
	

mgd:
	migrate -path db/migration -database "postgres://postgres:1235789@localhost:5432/ask_me?sslmode=disable" down
mgd1:
	migrate -path db/migration -database "postgres://postgres:1235789@localhost:5432/ask_me?sslmode=disable" down 1
	
mgd2:
	migrate -path db/migration -database "postgres://postgres:1235789@localhost:5432/ask_me?sslmode=disable" down 2
test: 
	go test -v -cover ./...
openDB:
	docker exec -it postgress-db psql -U postgres -d ask_me


mock:
	mockgen --package mock --destination db/mock/store.go "github.com/djsmk123/askmeapi/db/sqlc" Store

sqlcgen:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate	

run:
	go run main.go

.PHONY: postgres createdb dropdb mgup mgup1 mgup2 mgd mgd1 mgd2  test openDB mock 