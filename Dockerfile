FROM golang:1.14-alpine AS builder
ARG COMMIT
ARG VERSION
ENV LIBRARY_PATH="/app/libtensorflow/lib"
ENV LD_LIBRARY_PATH="/app/libtensorflow/lib"
ENV CGO_CFLAGS="-I/app/libtensorflow/include -L/app/libtensorflow/lib"
COPY . /app
WORKDIR /app
RUN mkdir /app/libtensorflow
RUN curl https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-1.15.0.tar.gz | tar -C libtensorflow -xz
RUN go build -ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT}"

FROM alpine:3
ENV LIBRARY_PATH="/libtensorflow/lib"
ENV LD_LIBRARY_PATH="/libtensorflow/lib"
COPY --from=builder /app/libtensorflow /libtensorflow
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
