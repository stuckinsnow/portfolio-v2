FROM node:18-alpine AS frontend-build

WORKDIR /app
COPY static/ ./static/
COPY package.json ./

RUN npm install
RUN npm run build

FROM golang:1.21-alpine AS backend-build

# Install golangci-lint
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

WORKDIR /app
COPY backend/ ./
RUN go mod download

# Run linting (optional - remove if you don't want to lint during build)
RUN golangci-lint run

RUN go build -o server .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=backend-build /app/server .
COPY --from=frontend-build /app/static/ ./static/

EXPOSE 8080

CMD ["./server"]