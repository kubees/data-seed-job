FROM golang:1.19-alpine as build
RUN apk add --no-cache git

WORKDIR /src

COPY . /src/
RUN go mod tidy


RUN go build -o app main.go

FROM alpine:3.12

RUN mkdir -p /app
COPY --from=build /src/app /app/app
CMD ["./app/app"]