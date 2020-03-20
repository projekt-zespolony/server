FROM golang:1.14-alpine AS builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY . /app
RUN go build

FROM scratch
COPY --from=builder /app/server /server
EXPOSE 80
ENTRYPOINT ["/server"]
