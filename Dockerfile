FROM golang:1.8

WORKDIR /go/src/github.com/minodisk/sqlabble
RUN go get -u \
      github.com/golang/dep/...
COPY . .
RUN dep ensure
RUN go install ./cmd/...

CMD sh test.sh
