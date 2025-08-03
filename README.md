# mono-repo-release

- [mono-repo-release](#mono-repo-release)
  - [Getting Started](#getting-started)
  - [1. Install Go](#1-install-go)
    - [macOS](#macos)
    - [Linux (Debian/Ubuntu)](#linux-debianubuntu)
  - [2. Install Bazel](#2-install-bazel)
    - [macOS](#macos-1)
    - [Linux](#linux)
  - [3. Build and Test](#3-build-and-test)
  - [4. Code Coverage](#4-code-coverage)
    - [Viewing Coverage as HTML](#viewing-coverage-as-html)
  - [5. Code Coverage Trick: Symbolic Link](#5-code-coverage-trick-symbolic-link)
  - [6. Notes](#6-notes)
  - [7. Troubleshooting](#7-troubleshooting)
  - [8. More Info](#8-more-info)

## Getting Started

This repository uses Go and Bazel for building and testing. Follow the steps below to set up your environment and run the project.

---

## 1. Install Go

### macOS
- Install Homebrew if you don't have it:  
  `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
- Install Go:  
  `brew install go`

### Linux (Debian/Ubuntu)
- Install Go:
  ```sh
  sudo apt update
  sudo apt install golang-go
  ```
- Or download from https://go.dev/dl/ for the latest version.

---

## 2. Install Bazel

### macOS
- Install with Homebrew:  
  `brew install bazelisk`

### Linux
- Download Bazelisk (recommended):
  ```sh
  curl -Lo bazelisk https://github.com/bazelbuild/bazelisk/releases/latest/download/bazelisk-linux-amd64
  chmod +x bazelisk
  sudo mv bazelisk /usr/local/bin/bazel
  ```
- Or follow instructions at https://bazel.build/install

---

## 3. Build and Test

```sh
bazel build //...
bazel test //...
```

---

## 4. Code Coverage

To generate a code coverage report:

```sh
bazel coverage //... --combined_report=lcov
```

Bazel will output a file at `bazel-out/_coverage/_coverage_report.dat`.

### Viewing Coverage as HTML

1. Install `lcov` and `genhtml`:
   - macOS: `brew install lcov`
   - Linux: `sudo apt install lcov`
2. Generate HTML report:
   ```sh
   genhtml bazel-out/_coverage/_coverage_report.dat -o genhtml
   ```
3. Open `genhtml/index.html` in your browser.

---

## 5. Code Coverage Trick: Symbolic Link

Some tools expect the coverage file at a specific location. If you encounter issues, create a symbolic link:

```sh
ln -sf $(pwd)/bazel-out/_coverage/_coverage_report.dat coverage.info
```

This makes `coverage.info` point to the actual coverage data, which can help some CI tools or scripts find the report.

---

## 6. Notes
- Make sure your tests call exported functions from your main code to ensure coverage is tracked.
- For Go, coverage is only reported for non-test files.

---

## 7. Troubleshooting
- If you see `lcov: ERROR: no gcov kernel data found`, it usually means the coverage file is missing or not in the expected format. Ensure you use Bazel's output file directly with `genhtml`.

---

## 8. More Info
- [Go Documentation](https://go.dev/doc/)
- [Bazel Documentation](https://bazel.build/)
- [LCOV/Genhtml](http://ltp.sourceforge.net/coverage/lcov.php)
