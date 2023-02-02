FROM golang:1.20-alpine
WORKDIR /opt/hello-gotuna
COPY . .
CMD ["go", "run", "./examples/fullapp/cmd/main.go"]
EXPOSE 8888
