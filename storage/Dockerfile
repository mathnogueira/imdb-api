FROM golang:alpine as build
WORKDIR /build

COPY . .

RUN go build -o bin/storage main.go

FROM alpine:3.13
WORKDIR /app

COPY --from=build /build/bin/storage .

CMD ["/app/storage"]