run:
	reflex -r '\.go' -s -- sh -c " go run cmd/goutilities/main.go"


test:
	reflex -r '\.go' -s -- sh -c "go test github.com/Bendomey/goutilities/test -v ."

test-and-exit:
	go test github.com/Bendomey/goutilities/test -v .

build:
	go build