# Makefile for github.com/loopcontext/auth-api-go
generate:
	GO111MODULE=on go run github.com/loopcontext/go-graphql-orm

reinit:
	GO111MODULE=on go run github.com/loopcontext/go-graphql-orm init

migrate:
	DATABASE_URL=sqlite3://dev.db PORT=8081 go run *.go migrate

automigrate:
	DATABASE_URL=sqlite3://dev.db PORT=8081 go run *.go automigrate

run:
	DATABASE_URL=sqlite3://dev.db PORT=8081 go run *.go start --cors

debug:
	DEBUG=true DATABASE_URL=sqlite3://dev.db PORT=8081 go run *.go start --cors

voyager:
	docker run --rm -v `pwd`/gen/schema.graphql:/app/schema.graphql -p 8082:80 graphql/voyager

build-lambda-function:
	GO111MODULE=on GOOS=linux go build -o main lambda/main.go && zip lambda.zip main && rm main

test-sqlite:
	GO111MODULE=on go build -o app *.go && DATABASE_URL=sqlite3://test.db ./app migrate && (DATABASE_URL=sqlite3://test.db PORT=7777 ./app start& export app_pid=$$! && make test-godog || test_result=$$? && kill $$app_pid && exit $$test_result)
test:
	GO111MODULE=on go build -o app *.go && ./app migrate && (PORT=7777 ./app start& export app_pid=$$! && make test-godog || test_result=$$? && kill $$app_pid && exit $$test_result)
// TODO: add detection of host ip (eg. host.docker.internal) for other OS
test-godog:
	docker run --rm --network="host" -v "${PWD}/features:/godog/features" -e GRAPHQL_URL=http://$$(if [[ $${OSTYPE} == darwin* ]]; then echo host.docker.internal;else echo localhost;fi):8081/graphql loopcontext/godog-graphql
