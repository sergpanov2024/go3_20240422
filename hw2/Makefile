pg_up:
	docker-compose up -d

pg_down:
	docker-compose down

pg_start:
	docker-compose start

pg_stop:
	docker-compose stop

createdb:
	docker-compose exec db createdb --username=postgres --owner=postgres bankstoredb

dropdb:
	docker-compose exec db dropdb -U postgres bankstoredb

psql:
	docker-compose exec db psql -U postgres -d postgres bankstoredb

migrateup1:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/bankstoredb?sslmode=disable" up 1

migrateup:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/bankstoredb?sslmode=disable" up

migratedown:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/bankstoredb?sslmode=disable" down

migratedown1:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/bankstoredb?sslmode=disable" down 1

sqlc:
	/home/user/go/bin/sqlc generate

test:
	go test -timeout 60s -v -cover ./... -count=1

server:
	go run main.go

.PHONY: pg_start pg_stop createdb dropdb migrateup migratedown sqlc test server pg_up pg_down