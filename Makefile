build:
	go build -o 2020-site

run: build
	./2020-site

docker:
	docker build -t 2020-site .
	docker-compose up -d