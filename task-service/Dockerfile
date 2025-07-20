FROM golang:1.21 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server .

FROM alpine
RUN apk add --no-cache bash netcat-openbsd
WORKDIR /app
COPY --from=builder /app/server .
COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh
EXPOSE 8080

CMD ["./wait-for-it.sh", "mysql", "3306", "--", "./server"]

