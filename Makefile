build:
	docker-compose build blog-app

run:
	docker-compose up blog-app

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up