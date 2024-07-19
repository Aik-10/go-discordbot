FROM golang:1.21.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v ./cmd/discordbot/main.go


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD ["./main"]