# Simple http service example in Go

## Build and run

```shell
go mod init books
go build
./books
```

## Run inside docker container with dev target

```shell
docker build --target dev . -t go-http
docker run -it -p 8080:8080 -v ${PWD}:/work go-http sh
go mod init books
go build
./books
```

## Build & run with runtime target

```shell
docker build . -t videos
docker run -it -p 8080:8080 videos
```

Navigate to http://localhost:8080