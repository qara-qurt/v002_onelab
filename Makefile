postgres:
	docker run --name v002_onelab -e POSTGRES_PASSWORD=password  -p 5436:5432 -v pgdata2:/var/lib/postgresql/data --rm -d postgres

migrate-up:
	migrate -database "postgres://postgres:password@localhost:5436/postgres?sslmode=disable" -path ./internal/repository/postgres/migration up

migrate-down:
	migrate -database "postgres://postgres:password@localhost:5436/postgres?sslmode=disable" -path ./internal/repository/postgres/migration down

create-migration:
	migrate create -ext sql -dir ./internal/repository/postgres/migration -seq init

run:
	docker run -dp 8090:8090 --name v002_onelab  --rm app

stop:
	docker stop v002_onelab