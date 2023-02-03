FROM golang:1.20-alpine AS builder
WORKDIR /opt/hello-gotuna
COPY . .
RUN go build examples/fullapp/cmd/main.go
EXPOSE 8888

FROM alpine AS runtime
WORKDIR /opt/hello-gotuna
COPY --from=builder /opt/hello-gotuna/main .
CMD ["./main"]
