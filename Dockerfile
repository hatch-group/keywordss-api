FROM golang:1.12.1-alpine3.9

COPY ./api /go/src/github.com/hatch-group/keywordss-api/api
ENV GO111MODULE=on

RUN apk update && \
    apk add --no-cache git && \
    apk add --no-cache make

WORKDIR /go/src/github.com/hatch-group/keywordss-api/api

RUN go get github.com/pilu/fresh

CMD ["make", "run"]
