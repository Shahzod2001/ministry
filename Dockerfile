FROM golang:latest as builder
WORKDIR /app
COPY . .

RUN git config --global http.sslVerify false

RUN go mod tidy
RUN go build -o elibrary_linux_amd64 cmd/ministry/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/elibrary_linux_amd64 .
COPY --from=builder /app/config/config.yml config/

RUN chmod +x /app/elibrary_linux_amd64

EXPOSE 8585

CMD ["./elibrary_linux_amd64"]
