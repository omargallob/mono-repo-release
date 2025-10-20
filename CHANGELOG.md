# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0](https://github.com/omargallob/mono-repo-release/compare/v1.0.1...v1.1.0) (2025-10-20)


### Features

* **ci:** monorepo release ([#8](https://github.com/omargallob/mono-repo-release/issues/8)) ([9bbaa85](https://github.com/omargallob/mono-repo-release/commit/9bbaa85e58e46730c6f9c28861bace5a5a956f38))

## [1.0.1](https://github.com/omargallob/mono-repo-release/compare/v1.0.0...v1.0.1) (2025-10-16)


### Bug Fixes

* use correct github ation ([9686a75](https://github.com/omargallob/mono-repo-release/commit/9686a75bbc1e7bbe833fd32db2419d2f4e4f5ced))

## 1.0.0 (2025-10-16)


### ⚠ BREAKING CHANGES

* introduction of release please and goreleaser

### Features

* add glance-lint configuration and deps in makefile ([#5](https://github.com/omargallob/mono-repo-release/issues/5)) ([56aaf4d](https://github.com/omargallob/mono-repo-release/commit/56aaf4d27dba91cbd2ec194af78d0dc1b61edbd7))
* **ci:** add code coverage via bazel ([#1](https://github.com/omargallob/mono-repo-release/issues/1)) ([8eb610d](https://github.com/omargallob/mono-repo-release/commit/8eb610df15465aee3e89605590bd60508dc2ab76))
* introduction of release please and goreleaser ([fb2138f](https://github.com/omargallob/mono-repo-release/commit/fb2138f7105ea3f9dfc90213028080c198025f8e))
* switch to ginkgo ([#2](https://github.com/omargallob/mono-repo-release/issues/2)) ([2ea5471](https://github.com/omargallob/mono-repo-release/commit/2ea54717ad7e0486ad4dc82c68742f32c5fc56c0))
* switch to ginkgo testing ([#3](https://github.com/omargallob/mono-repo-release/issues/3)) ([e7cd505](https://github.com/omargallob/mono-repo-release/commit/e7cd505f149eab21dd5a7255790d93d2dcd8f981))


### Bug Fixes

* **lib2:** resolve memory leak ([e0a1e17](https://github.com/omargallob/mono-repo-release/commit/e0a1e17e8e980fa1a1b6242af89a6c6e7f187ccb))

## [Unreleased]

### Added
- Initial project setup with Bazel build system
- GoReleaser configuration for automated releases
- Release Please integration for automated changelog generation
- Conventional commits setup for semantic versioning
- GitHub Actions workflows for CI/CD
- Multi-binary builds for main app and library CLI tools

### Features
- **lib1**: Core library functionality
- **lib2**: Secondary library with version management
- **app**: Main application binary

## [0.1.0] - 2025-10-16

### Added
- Initial release
- Basic project structure
- Bazel build configuration
- Test suites with coverage reporting
