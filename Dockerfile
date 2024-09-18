FROM golang:1.22-alpine

WORKDIR /app

# Install air for live reload (older stable version)
RUN go install github.com/cosmtrek/air@v1.40.4
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Install app dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the app source code
COPY . .

RUN swag init

EXPOSE 8080

CMD ["air"]  # or use CMD ["reflex", "-r", ".*\\.go$", "-s", "--", "go", "run", "main.go"]
