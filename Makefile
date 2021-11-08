build:
	go build -o 2020-site

run:
	docker-compose up -d

stop:
	docker-compose down

db:
	docker-compose up -d db

dev:
	go run *.go

db-conn:
        docker exec -it 2020-site_db_1 psql -U postgres
