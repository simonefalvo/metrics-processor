FROM golang:1.15-alpine AS builder
WORKDIR /go/src/github.com/smvfal/nats-sub
COPY . .
RUN go build

FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/github.com/smvfal/nats-sub/nats-sub .
ENV NATS_URL="http://nats.openfaas:4222" \
    SUBJECT="nats-test"

CMD ./nats-sub -s $NATS_URL $SUBJECT
