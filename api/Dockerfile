FROM golang:1.23.6-alpine3.21


RUN apk update && apk add git
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest
RUN apk add --no-cache -X http://dl-cdn.alpinelinux.org/alpine/edge/community watchexec

COPY ./ /usr/src/app

RUN go mod download

CMD ["go", "run", "main.go"]
