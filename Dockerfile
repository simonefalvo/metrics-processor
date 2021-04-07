FROM golang:1.15-alpine AS builder
WORKDIR /go/src/github.com/smvfal/metrics-processor
COPY . .
RUN go build

FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/github.com/smvfal/metrics-processor/metrics-processor .
RUN mkdir data
ENV NATS_URL="http://nats.openfaas:4222" \
    SUBJECT="nats-test"

CMD ./metrics-processor -s $NATS_URL $SUBJECT
