default: all

all: clean build

test:
	go test ./...

build:
	go build -o dist/optprov *.go

linux-build:
	GOOS=linux GOARCH=amd64 go build -o dist/optprov cmd/optprov/*

format:
	gofumpt -w -l .

clean:
	rm -r dist || true

docker:
	docker build . -t dennis-tra/optprov:latest

tools:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1
	go install github.com/volatiletech/sqlboiler/v4@v4.6.0
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.6.0

db-reset: migrate-down migrate-up models

models:
	sqlboiler psql

make open-api:
	oapi-codegen -generate types -package types ./pkg/api/openapi.yaml > ./pkg/api/types/types.gen.go
	openapi-generator generate -g typescript-fetch -i ./pkg/api/openapi.yaml -o ./optprov/src/api --global-property client --additional-properties=supportsES6=true,typescriptThreePlus=true

test-db:
	docker run --rm -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=optprov -e POSTGRES_DB=optprov postgres:13

migrate-up:
	migrate -database 'postgres://optprov:password@localhost:5432/optprov?sslmode=disable' -path migrations up

migrate-down:
	migrate -database 'postgres://optprov:password@localhost:5432/optprov?sslmode=disable' -path migrations down

ts-client:
	openapi-generator generate -g typescript-fetch -i ./pkg/api/openapi.yaml -o ./optprov/src/api --global-property client --additional-properties=supportsES6=true,typescriptThreePlus=true

.PHONY: all clean test format tools models migrate-up migrate-down
