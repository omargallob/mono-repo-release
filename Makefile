# Makefile for mono-repo-release

.PHONY: check-deps build test coverage coverage-html release-check release-snapshot clean

check-deps:
	@echo "Checking for Go..."
	@command -v go >/dev/null 2>&1 || (echo "Go is not installed. Please install Go." && exit 1)
	@echo "Go found: $$(go version)"
	@echo "Checking for Bazel..."
	@command -v bazel >/dev/null 2>&1 || (echo "Bazel is not installed. Please install Bazelisk or Bazel." && exit 1)
	@echo "Bazel found: $$(bazel --version)"
	@echo "Checking for GoReleaser..."
	@command -v goreleaser >/dev/null 2>&1 || echo "Warning: goreleaser not found. Install with 'brew install goreleaser'."
	@echo "Checking for lcov..."
	@command -v lcov >/dev/null 2>&1 || echo "Warning: lcov not found. HTML coverage reports may not be available."
	@echo "Checking for genhtml..."
	@command -v genhtml >/dev/null 2>&1 || echo "Warning: genhtml not found. HTML coverage reports may not be available."
	@echo "Checking for golangci-lint..."
	@command -v golangci-lint >/dev/null 2>&1 || echo "Warning: golangci-lint not found. Please install golangci-lint."
	@echo "All required tools checked."
	
build:
	bazel build //...

test:
	bazel test //...

coverage:
	bazel coverage //... --combined_report=lcov

coverage-html: coverage
	@command -v genhtml >/dev/null 2>&1 || (echo "genhtml not found. Please install lcov/genhtml." && exit 1)
	genhtml bazel-out/_coverage/_coverage_report.dat -o genhtml
	@echo "Open genhtml/index.html in your browser to view the report."

# Release-related targets
release-check:
	@command -v goreleaser >/dev/null 2>&1 || (echo "goreleaser not found. Install with 'brew install goreleaser'." && exit 1)
	goreleaser check

release-snapshot: release-check
	goreleaser release --snapshot --clean

clean:
	bazel clean
	rm -rf genhtml/
	rm -rf dist/
