postgres:
	docker run --name meurig -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=ordering-app -d postgres
migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/ordering-app?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/ordering-app?sslmode=disable" -verbose down
sql:
	sqlc generate
run:
	DATABASE_CONNECTION_STRING=postgresql://postgres:postgres@localhost:5432/ordering-app?sslmode=disable \
	DB_HOST=localhost \
	DB_PORT=5432 \
	DB_USER=postgres \
	DB_PASSWORD=postgres \
	DB_TABLE=ordering-app \
	go run .
build:
	GOOS=linux go build -o service-product . 
docker_build:
	docker build --no-cache -t izaakdale/service-product .
