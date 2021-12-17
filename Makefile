build:
	go build -o 2020-site

run:
	docker-compose up -d --force-recreate

stop:
	docker-compose down

db:
	docker-compose up -d db

dev:
	go run *.go
