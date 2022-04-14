.DEFAULT_GOAL:=build
BINARY:=server
PACKAGE:=github.com/cudneys/password-generator

VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
COMMIT_HASH="$(git rev-parse --short HEAD)"
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')


clean:
	rm -rvf dist/

prep:
	mkdir -p dist
	mkdir -p dist/linux/amd64
	mkdir -p dist/linux/arm
	mkdir -p dist/darwin/amd64
	mkdir -p dist/darwin/m1
	mkdir -p windows/amd64

build: clean prep
	swag init
	GOOS=linux GOARGH=amd64 go build -o dist/linux/amd64/${BINARY}
	GOOS=linux GOARGH=arm64 go build -o dist/linux/arm/${BINARY}
	GOOS=darwin GOARGH=amd64 go build -o dist/darwin/and64/${BINARY}
	GOOS=darwin GOARGH=arm64 go build  -o dist/darwin/m1/${BINARY} --ldflags="-X ${PACKAGE}/version.Version=${VERSION} -X ${PACKAGE}/version.BuildTimestamp=${BUILD_TIMESTAMP}"
	GOOS=windows GOARGH=amd64 go build  -o dist/windows/and64/${BINARY}.exe

docker: build
	docker build -t ${PACKAGE}:${VERSION}