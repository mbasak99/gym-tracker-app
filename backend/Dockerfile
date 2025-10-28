FROM golang:1.25-alpine3.22 AS builder

ARG APP_ENV
ARG API_URL
ARG API_PORT

ENV APP_ENV=${APP_ENV}
ENV API_URL=${API_URL}
ENV API_PORT=${API_PORT}

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env /app/.env
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /app/server ./src

FROM alpine:3.18
RUN apk add --no-cache ca-certificates

# create non-root user
RUN addgroup -S app && adduser -S -G app app

COPY --from=builder --chown=app:app /app/server /usr/local/bin/server

# COPY src/models /app/database # idk if I need these yet

RUN mkdir -p /app && chown -R app:app /app

USER app

EXPOSE ${API_PORT}

ENTRYPOINT ["/usr/local/bin/server"]
