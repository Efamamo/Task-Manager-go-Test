build:
	@go build -o bin/task-manager ./Delivery/main.go

test:
	go test -v ./...

run: build
	@./bin/task-manager