FROM golang:wheezy

WORKDIR /go/src/github.com/minodisk/sqlabble
RUN go get -u \
      github.com/go-sql-driver/mysql \
      github.com/mattn/goveralls \
      github.com/sergi/go-diff/diffmatchpatch

COPY . .

CMD go test ./...
