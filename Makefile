run:
	reflex -r '\.go' -s -- sh -c "go run cmd/goutilites/main.go"

test:
	go test github.com/Bendomey/goutilities/test

build:
	go build