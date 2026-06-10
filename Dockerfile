FROM --platform=linux/amd64 golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notely

FROM --platform=linux/amd64 debian:stable-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /app/notely /usr/bin/notely
CMD ["notely"]
