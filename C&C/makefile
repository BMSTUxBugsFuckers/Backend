db-create:
	docker run --name CAndC -e POSTGRES_PASSWORD=passwd -p 5432:5432 -d postgres

migrations-new:
	migrate create -ext sql -dir schema -seq $(M_NAME)

migrations-up:
	migrate -path ./schema -database 'postgres://postgres:passwd@localhost:5432/postgres?sslmode=disable' up

migrations-down:
	migrate -path ./schema -database 'postgres://postgres:passwd@localhost:5432/postgres?sslmode=disable' down

run-server:
	go run cmd/main.go