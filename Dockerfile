FROM golang:1.12.6 AS build
WORKDIR /go/src/github.com/Laica-Lunasys/echoback

ENV GOOS linux
ENV CGO_ENABLED 0

COPY . .
RUN go build -a -installsuffix cgo -v -o echoback main.go

FROM alpine
WORKDIR /app

EXPOSE 8080
COPY --from=build /go/src/github.com/Laica-Lunasys/echoback /app/

RUN apk add --no-cache ca-certificates

ENTRYPOINT ["exec", "/app/echoback"]
