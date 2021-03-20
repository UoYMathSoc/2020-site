FROM golang:1.16.2-alpine3.13
RUN mkdir /site
ADD . /site
WORKDIR /site
RUN go mod download
RUN go build -o 2020-site .
EXPOSE 3000
CMD ["/site/2020-site"]
