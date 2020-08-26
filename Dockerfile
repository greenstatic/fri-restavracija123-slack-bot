FROM golang:1.15-alpine3.12 AS builder
RUN go version

RUN apk update && apk add --no-cache git ca-certificates make

COPY . /bot
WORKDIR /bot

RUN make docker-build

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bot/bot_amd64 .

# Default values
ENV DEBUG false

ENTRYPOINT ["./bot_amd64"]
