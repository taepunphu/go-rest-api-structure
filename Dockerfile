FROM golang:1.20.6-bullseye
ENV config=docker
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY . /app
RUN swag init
EXPOSE 8001
CMD go run main.go