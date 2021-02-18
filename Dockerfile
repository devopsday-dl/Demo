FROM golang:1.12.9-alpine3.9 as builder

MAINTAINER <devops008@sina.com>

WORKDIR /tmp

COPY devops.go /tmp

RUN go build devops.go

FROM alpine:latest

WORKDIR /usr/src/app/

COPY --from=builder /tmp/devops /usr/src/app/

CMD ["./devops"]
