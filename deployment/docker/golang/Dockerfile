FROM golang:1.25.1-alpine AS builder

ARG SERVICE_NAME

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY pb ./pb
COPY cmd/${SERVICE_NAME} ./cmd/${SERVICE_NAME}

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/${SERVICE_NAME}/main.go

FROM alpine AS runner

RUN apk add --no-cache tzdata

RUN addgroup -g 1000 appgroup && \
    adduser -D -u 1000 -G appgroup appuser

WORKDIR /app

COPY --from=builder /app/app .

RUN chown -R appuser:appgroup /app

USER appuser

CMD ["./app"]
