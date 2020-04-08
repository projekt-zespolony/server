FROM golang:1.14-alpine AS builder
ARG COMMIT
ARG VERSION
ENV CGO_ENABLED=0
WORKDIR /app
COPY . /app
RUN go build -ldflags="-X main.version=${VERSION} main.commit=${COMMIT}"

FROM alpine:3
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
