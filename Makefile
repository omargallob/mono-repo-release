# Makefile for mono-repo-release

.PHONY: check-deps build test coverage coverage-html

check-deps:
	@echo "Checking for Go..."
	@command -v go >/dev/null 2>&1 || (echo "Go is not installed. Please install Go." && exit 1)
	@echo "Go found: $$(go version)"
	@echo "Checking for Bazel..."
	@command -v bazel >/dev/null 2>&1 || (echo "Bazel is not installed. Please install Bazelisk or Bazel." && exit 1)
	@echo "Bazel found: $$(bazel --version)"
	@echo "Checking for lcov..."
	@command -v lcov >/dev/null 2>&1 || echo "Warning: lcov not found. HTML coverage reports may not be available."
	@echo "Checking for genhtml..."
	@command -v genhtml >/dev/null 2>&1 || echo "Warning: genhtml not found. HTML coverage reports may not be available."
	@echo "Checking for GoReleaser..."
	@command -v goreleaser >/dev/null 2>&1 || (echo "GoReleaser is not installed. Please install GoReleaser." && exit 1)
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

clean:
	bazel clean --expunge
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel run //:gazelle
