# 2020-site
![Go](https://github.com/UoYMathSoc/2020-site/workflows/Go/badge.svg)

requires: [github.com/kyleconroy/sqlc](https://github.com/kyleconroy/sqlc)

## Edit Config
```Shell
$ cp config.toml.example config.toml
$ nano config.toml
```
Set database host to run to "db" when running inside docker, `make run`, and "localhost" when running in develpoment, `make dev`

## Development
Run `make db` to start the database in docker and `make dev` to start an instance of the website locally
> Remember to set `host = "localhost"` inside config.toml

## Production
Run `make run` to start the website and database in docker on port 8080
> Remember to set `host = "db"` inside config.toml

## Generate Database Package
If you add sql to the repo run the following to generate the corrosponding go code 
```Shell
$ sqlc generate
```