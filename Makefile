build:
	GOOS=linux GOARCH=amd64 go build -o build/main cmd/main.go

serve:
	go run cmd/main.go

config:
	cp .env.example .env

tidy:
	go mod tidy
key:
	go run cmd/artisan.go key:generate

gotest:
	go test ./test/...

migration:
	go run cmd/artisan.go make:migration ${table}

migrate-up:
	go run cmd/artisan.go migrate:up

migrate-down:
	go run cmd/artisan.go migrate:down