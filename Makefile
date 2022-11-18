DB_URL=postgresql://postgres:1105@localhost:5432/student_db?sslmode=disable

swag-init:
	swag init -g api/api.go -o api/docs

start:
	go run cmd/main.go

migrateup:
	migrate -path migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migration -database "$(DB_URL)" -verbose down


.PHONY: start migrateup migratedown