run:
	reflex -r '\.go' -s -- sh -c " go run cmd/goutilities/main.go"

test:
	go test github.com/Bendomey/goutilities/test -v .

build:
	go build