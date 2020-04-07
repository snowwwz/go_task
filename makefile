setup:
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/setup

lint:
	go vet ./...
	golint ./...

fmt:
	go fmt ./...

