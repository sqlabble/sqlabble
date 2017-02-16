#/bin/sh

set -e

rm -rf cover.out cover.out.tmp
echo 'mode: count' > cover.out
for pkg in $(go list ./...); do
  go test -v -race -cover -covermode=count -coverprofile=cover.out.tmp $pkg
  cat cover.out.tmp | tail -n +2 >> cover.out
done
rm -rf cover.out.tmp

goveralls -coverprofile=cover.out -repotoken=$COVERALLS_TOKEN
