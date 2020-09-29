# 2020-site
![Go](https://github.com/UoYMathSoc/2020-site/workflows/Go/badge.svg)
## Generate Database Package
```Shell
$ sqlc generate
```
requires: [github.com/kyleconroy/sqlc](https://github.com/kyleconroy/sqlc)
## Edit Config
```Shell
$ cp config.toml.example config.toml
$ nano config.toml
```

## Build and Run Site
```Shell
$ go build && ./2020-site
```
