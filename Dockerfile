FROM golang:1.14-alpine AS builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY . /app
RUN go build

FROM alpine:3
COPY --from=builder /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
