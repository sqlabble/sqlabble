FROM golang:alpine

WORKDIR /go/src/github.com/minodisk/sqlabble
RUN apk --update add git && \
    go get -u \
      github.com/sergi/go-diff/diffmatchpatch

COPY . .

CMD go test -v ./...
