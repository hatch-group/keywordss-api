FROM golang:1.12.1-alpine3.9

COPY . /go/src/keywordss/api

RUN apk update && \
    apk add --no-cache git && \
    apk add --no-cache make

WORKDIR /go/src/keywordss/api
RUN go get -u github.com/golang/dep/cmd/dep && \
    make dep && \
    apk del git

CMD ["make", "run"]
