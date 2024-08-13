build:
	@go build -o bin/task-manager ./Delivery/main.go

test:
	@go test $(go list ./Tests/... | grep -v '/Tests/repository_test') -v
