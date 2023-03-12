postgres-run:
	@docker run --name postgrestest -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=asdqwe123 -e POSTGRES_DB=kelindan -d postgres
redis-run:
	@docker run -d -p 6379:6379 --name redistest redis

postgres:
	@docker start postgrestest
redis:
	@docker start redistest

postgres-stop:
	@docker stop postgrestest
redis-stop:
	@docker stop redistest

setup: postgres-run redis-run postgres-stop redis-stop

stop: postgres-stop redis-stop

container: postgres redis

main: 
	@go run ./cmd/main.go

testing:
	@go test -v -cover ./...

run: postgres redis main

test: postgres redis testing stop