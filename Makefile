start:
	go run app/api/main.go

migrate:
	go run migrations/main.go

seed:
	go run seeders/main.go