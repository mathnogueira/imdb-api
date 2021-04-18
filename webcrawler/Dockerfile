FROM golang:alpine AS build
WORKDIR /build

COPY . .

RUN go build -o bin/webcrawler main.go

FROM alpine:3.13
WORKDIR /app

COPY --from=build /build/bin/webcrawler .

CMD [ "/app/webcrawler" ]