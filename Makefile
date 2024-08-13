build:
	@go build -o bin/task-manager ./Delivery/main.go

test:
	@test_packages=$$(./scripts/filter_test_packages.sh);
	go test -v $$test_packages

run: build
	@./bin/task-manager