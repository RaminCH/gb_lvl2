main:
	echo "Main"

build:
	go build -o bin/main main.go

run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

all: hello build


build_and_run: build run

# clean:
#  go clean
#  rm ${BINARY_NAME}-darwin
#  rm ${BINARY_NAME}-linux
#  rm ${BINARY_NAME}-windows

# test:
#  go test ./...

# test_coverage:
#  go test ./... -coverprofile=coverage.out

# dep:
#  go mod download

# vet:
#  go vet

# lint:
#  golangci-lint run --enable-all