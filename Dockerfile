FROM golang:1.8

WORKDIR /go/src/github.com/minodisk/sqlabble
RUN go get -u \
      github.com/minodisk/caseconv \
      github.com/go-sql-driver/mysql \
      github.com/mattn/go-zglob \
      github.com/mattn/goveralls \
      github.com/sergi/go-diff/diffmatchpatch

COPY . .
RUN go install ./...

CMD go test -race ./...
