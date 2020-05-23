FROM golang:1.14 AS builder
ARG COMMIT
ARG VERSION
ENV LIBRARY_PATH="/libtensorflow/lib"
ENV LD_LIBRARY_PATH="/libtensorflow/lib"
ENV CGO_CFLAGS="-I/libtensorflow/include -L/libtensorflow/lib"
ADD https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-1.15.0.tar.gz /libtensorflow.tar.gz
RUN mkdir -p /libtensorflow
RUN tar -C /libtensorflow -xzf /libtensorflow.tar.gz
RUN rm /libtensorflow.tar.gz
RUN git clone --depth=1 https://github.com/projekt-zespolony/siec-neuronowa /neural
COPY . /app
WORKDIR /app
RUN go build -ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT}"

FROM ubuntu:18.04
ENV LIBRARY_PATH="/libtensorflow/lib"
ENV LD_LIBRARY_PATH="/libtensorflow/lib"
WORKDIR /
COPY --from=builder /model /model
COPY --from=builder /libtensorflow /libtensorflow
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
