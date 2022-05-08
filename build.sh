docker run --rm -v "$PWD":/app -w /app -e GOOS=windows -e GOARCH=amd64 -e CGO_ENABLED=1 -e CC=x86_64-w64-mingw32-gcc golang:latest sh -c 'apt update -y && apt install gcc-mingw-w64 -y && go build'
