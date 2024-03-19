# Simple http service example in Go

## Build and run

```shell
go mod init books
go build
./books
```

## Run inside docker container

```shell
docker build --target dev . -t go
docker run -it -p 8080:8080 -v ${PWD}:/work go sh
go mod init books
go build
./books
```
OLD!