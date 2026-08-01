.PHONY: test

test:
	@echo "Running all tests in subdirectories..."
	@find . name "*_test.go" -exec sh -c 'cd $$(dirname "{}") && go test -v' \;

bench-test:
	@echo "Running all tests in subdirectories..."
	@find . name "*_test.go" -exec sh -c 'cd $$(dirname "{}") && go test -v --bench .' \;

run:
	go run main.go
