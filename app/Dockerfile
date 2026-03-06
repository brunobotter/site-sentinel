FROM golang:1.24-alpine AS builder
WORKDIR /src
RUN apk add --no-cache wget

COPY app/go.mod app/go.sum ./
RUN go mod download

COPY app/. .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/site-sentinel .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /out/site-sentinel ./site-sentinel 
EXPOSE 8080
ENTRYPOINT ["./site-sentinel"]