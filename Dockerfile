
FROM golang:1.15-alpine as dev

WORKDIR /work

FROM golang:1.15-alpine as build

WORKDIR /books
COPY ./books/* /books/
RUN go build -o books

FROM alpine as runtime 
COPY --from=build /books/books /
COPY ./books/books.json /
CMD ./books