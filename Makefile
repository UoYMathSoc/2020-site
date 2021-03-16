build:
	go build -o 2020-site

run: build
	./2020-site

docker: docker-build docker-run

docker-build:
	docker build -t 2020-site .

docker-run:
	docker run -p 80:3000 -d 2020-site

docker-compose:
	docker-compose up -d