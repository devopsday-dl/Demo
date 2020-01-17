FROM golang:1.12.9-alpine3.9

MAINTAINER <devops008@sina.com>

COPY devops.go /usr/src/app

CMD ["go","run", "devops.go"]
