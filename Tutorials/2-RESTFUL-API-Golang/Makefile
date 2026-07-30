.PHONY: test

test:
	@echo "Running all subdirectory tests..."
	@find . name "*_test.go" -exec sh -c 'cd $$(dirname "{}") && go test -v' \;

bench-test:
	@echo "Running all subdirectory tests..."
	@find . name "*_test.go" -exec sh -c 'cd $$(dirname "{}") && go test -v --bench .' \;