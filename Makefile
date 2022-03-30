default: all

all: clean build

test:
	go test ./...

build: clean
	mkdir -p out
	go build -o out/optprov cmd/optprov/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o out/optprov_linux_amd64 cmd/optprov/*

format:
	gofumpt -w -l .

clean:
	rm -r dist || true

tools:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1
	go install github.com/volatiletech/sqlboiler/v4@v4.6.0
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.6.0

db-reset: migrate-down migrate-up models

models:
	sqlboiler psql

make open-api:
	oapi-codegen -generate types -package types ./pkg/api/openapi.yaml > ./pkg/api/types/types.gen.go
	openapi-generator generate -g typescript-fetch -i ./pkg/api/openapi.yaml -o ./frontend/src/api --global-property models --additional-properties=supportsES6=true,typescriptThreePlus=true

database:
	docker run --rm -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=optprov -e POSTGRES_DB=optprov_dev postgres:13

migrate-up:
	migrate -database 'postgres://optprov:password@localhost:5432/optprov?sslmode=disable' -path migrations up

migrate-down:
	migrate -database 'postgres://optprov:password@localhost:5432/optprov?sslmode=disable' -path migrations down
