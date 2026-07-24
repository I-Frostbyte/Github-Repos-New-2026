.PHONY: test

test:
	@echo "Running tests in all subdirectories..."
	@find . name "*_test.go" -exec sh -c 'cd $$(dirname "{}") && go test -v' \;

bench-test:
	@echo "Running tests in all subdirectories..."
	@find . name "*_test.go" -exec sh -c 'cd $$(dirname "{}") && go test -v --bench .' \;