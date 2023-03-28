FROM golang:1.20.0-alpine as builder

WORKDIR /go/bmf-tech-client/app

COPY . .

RUN apk update && \
    apk upgrade && \
    apk add --no-cache libc-dev gcc git openssh openssl bash

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o app

FROM --platform=linux/amd64 alpine

COPY --from=builder /go/bmf-tech-client/app ./

CMD ["./app"]
