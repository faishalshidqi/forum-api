FROM golang:1.24-bookworm AS builder
WORKDIR /app
COPY ./ /app
COPY ./.env /app/.env
RUN GIN_MODE=release CGO_ENABLED=0 GOOS=linux go build -o /forumapi

FROM alpine:latest
WORKDIR /
COPY --from=builder /app/.env /.env
COPY --from=builder /forumapi /forumapi
EXPOSE 5000
ENTRYPOINT ["/forumapi"]
