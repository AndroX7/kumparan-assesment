dep:
	go mod tidy
	go mod vendor

# Use this only for development
dev:
	go build -o bin/gaming-company-test app/api/main.go
	./bin/gaming-company-test

createdb:
	npx sequelize-cli db:create

migrate:
	npx sequelize-cli db:migrate

seed:
	go run ./seeders/.

build:
	set GOOS=linux && set GOARCH=amd64 && go build -o bin/gaming-company-test app/api/main.go

docker-compose-up-clean:
	docker-compose rm && \
	docker-compose pull && \
	docker-compose build --no-cache && \
	docker-compose up --force-recreate gaming-company-test redis-server

refresh:
	npx sequelize db:migrate:undo:all
	npx sequelize db:migrate