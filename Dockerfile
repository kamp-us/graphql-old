# Building the binary of the App
FROM golang:1.17 AS builder
RUN mkdir -p /app
WORKDIR /app
COPY . .

RUN git config \
    --global \
    url."https://usirin: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@github.com".insteadOf \
    "https://github.com"

RUN CGO_ENABLED=0 \
    GIT_TERMINAL_PROMPT=1 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o /bin/graphql /app/cmd/graphql/main.go

# Moving the binary to the 'final Image' to make it smaller
FROM alpine
WORKDIR /app
RUN apk add --no-cache nano git curl
# COPY --from=builder /feedback-api/internal/configs/dev.env .env
COPY --from=builder bin/graphql graphql

CMD ["./graphql"]

EXPOSE 8000
