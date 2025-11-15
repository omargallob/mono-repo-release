# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.3.1](https://github.com/omargallob/mono-repo-release/compare/v2.3.0...v2.3.1) (2025-11-15)


### Bug Fixes

* building of packages works ([640e499](https://github.com/omargallob/mono-repo-release/commit/640e49933319fea05ad8f3865eaf4a8db449e564))
* refactored gh action ([eab2169](https://github.com/omargallob/mono-repo-release/commit/eab21694032d6fe9c0aa071ef2a9ffd4496e2c5a))

## [2.3.0](https://github.com/omargallob/mono-repo-release/compare/v2.2.3...v2.3.0) (2025-11-15)


### Features

* separate goreleaser per package ([2ccebb3](https://github.com/omargallob/mono-repo-release/commit/2ccebb3147053d6a8e5627b3b1d6725fd1963512))

## [2.2.3](https://github.com/omargallob/mono-repo-release/compare/v2.2.2...v2.2.3) (2025-11-15)


### Bug Fixes

* broken pakacge name ([5f34212](https://github.com/omargallob/mono-repo-release/commit/5f34212f6acef7c544f339e0476226c1c764d2f7))
* missing version file ([bd8146a](https://github.com/omargallob/mono-repo-release/commit/bd8146a918f501ba021370be6291c12a14c4c15e))

## [2.2.2](https://github.com/omargallob/mono-repo-release/compare/v2.2.1...v2.2.2) (2025-11-15)


### Bug Fixes

* another tweak ([8cf5d6a](https://github.com/omargallob/mono-repo-release/commit/8cf5d6a542e61b9dec620b58d8e97f878359cefd))

## [2.2.1](https://github.com/omargallob/mono-repo-release/compare/v2.2.0...v2.2.1) (2025-11-15)


### Bug Fixes

* moved gorelease to root ([b29ab6e](https://github.com/omargallob/mono-repo-release/commit/b29ab6e53d841241c4d3c2f5179635b3a138a7e9))
* remove goreleaser from  lib folder ([b54c344](https://github.com/omargallob/mono-repo-release/commit/b54c34462de7e91fcb6cdcdad1addb4595b93693))

## [2.2.0](https://github.com/omargallob/mono-repo-release/compare/v2.1.0...v2.2.0) (2025-11-03)


### Features

* doc change ([6d27b12](https://github.com/omargallob/mono-repo-release/commit/6d27b1299fb3b51c80dc5c32f8c17b070fc565da))

## [2.1.0](https://github.com/omargallob/mono-repo-release/compare/v2.0.0...v2.1.0) (2025-10-20)


### Features

* **pkg/lib1:** added a readme ([70f4507](https://github.com/omargallob/mono-repo-release/commit/70f450715193e72819803cc6c6534bb032313544))


### Bug Fixes

* **ci:** set initial bootstap sha ([da41c49](https://github.com/omargallob/mono-repo-release/commit/da41c49158afc5f1d28c307d7fb703f105e8d3e5))

## [2.0.0](https://github.com/omargallob/mono-repo-release/compare/v1.3.0...v2.0.0) (2025-10-20)


### ⚠ BREAKING CHANGES

* introduction of release please and goreleaser

### Features

* add glance-lint configuration and deps in makefile ([#5](https://github.com/omargallob/mono-repo-release/issues/5)) ([56aaf4d](https://github.com/omargallob/mono-repo-release/commit/56aaf4d27dba91cbd2ec194af78d0dc1b61edbd7))
* **ci:** add code coverage via bazel ([#1](https://github.com/omargallob/mono-repo-release/issues/1)) ([8eb610d](https://github.com/omargallob/mono-repo-release/commit/8eb610df15465aee3e89605590bd60508dc2ab76))
* **ci:** monorepo release ([#8](https://github.com/omargallob/mono-repo-release/issues/8)) ([9bbaa85](https://github.com/omargallob/mono-repo-release/commit/9bbaa85e58e46730c6f9c28861bace5a5a956f38))
* introduction of release please and goreleaser ([fb2138f](https://github.com/omargallob/mono-repo-release/commit/fb2138f7105ea3f9dfc90213028080c198025f8e))
* **lib2:** added ascii art function ([#10](https://github.com/omargallob/mono-repo-release/issues/10)) ([4e9b830](https://github.com/omargallob/mono-repo-release/commit/4e9b830309c68089729fa43fc9f3b5514da65d4b))
* **pkg/lib1:** added a readme ([68bd5cd](https://github.com/omargallob/mono-repo-release/commit/68bd5cd98a6267b65e411bfd2a03ddf75e7d694e))
* switch to ginkgo ([#2](https://github.com/omargallob/mono-repo-release/issues/2)) ([2ea5471](https://github.com/omargallob/mono-repo-release/commit/2ea54717ad7e0486ad4dc82c68742f32c5fc56c0))
* switch to ginkgo testing ([#3](https://github.com/omargallob/mono-repo-release/issues/3)) ([e7cd505](https://github.com/omargallob/mono-repo-release/commit/e7cd505f149eab21dd5a7255790d93d2dcd8f981))


### Bug Fixes

* **ci:** commitlint and rp to use correct scopes ([5897dcb](https://github.com/omargallob/mono-repo-release/commit/5897dcbfd4c3f759baff9bb2c08c5ed3f0254296))
* **ci:** update the values for the action ([83bde52](https://github.com/omargallob/mono-repo-release/commit/83bde526641e36513f7596b84a10ca33dea6c416))
* **lib2:** included packages to test updates ([#12](https://github.com/omargallob/mono-repo-release/issues/12)) ([a338f35](https://github.com/omargallob/mono-repo-release/commit/a338f35ad96eb07b24f009403b2ebfc33c21dbd8))
* **lib2:** resolve memory leak ([e0a1e17](https://github.com/omargallob/mono-repo-release/commit/e0a1e17e8e980fa1a1b6242af89a6c6e7f187ccb))
* **pkg/lib1:** update readme ([1d5df23](https://github.com/omargallob/mono-repo-release/commit/1d5df234fc2f4d3071fe8df793544cacafea01f1))
* **pkg/lib2:** added a readme ([4ac4a56](https://github.com/omargallob/mono-repo-release/commit/4ac4a563c2b5f9491b6a8edf54ee79ff04e526c9))
* **pkg/lib2:** included bs readme ([d97451d](https://github.com/omargallob/mono-repo-release/commit/d97451db7fe503d040e530311101029e9dde2fe6))
* use correct github ation ([9686a75](https://github.com/omargallob/mono-repo-release/commit/9686a75bbc1e7bbe833fd32db2419d2f4e4f5ced))

## [1.3.0](https://github.com/omargallob/mono-repo-release/compare/v1.2.1...v1.3.0) (2025-10-20)


### Features

* **pkg/lib1:** added a readme ([68bd5cd](https://github.com/omargallob/mono-repo-release/commit/68bd5cd98a6267b65e411bfd2a03ddf75e7d694e))


### Bug Fixes

* **ci:** commitlint and rp to use correct scopes ([5897dcb](https://github.com/omargallob/mono-repo-release/commit/5897dcbfd4c3f759baff9bb2c08c5ed3f0254296))
* **pkg/lib1:** update readme ([1d5df23](https://github.com/omargallob/mono-repo-release/commit/1d5df234fc2f4d3071fe8df793544cacafea01f1))
* **pkg/lib2:** added a readme ([4ac4a56](https://github.com/omargallob/mono-repo-release/commit/4ac4a563c2b5f9491b6a8edf54ee79ff04e526c9))
* **pkg/lib2:** included bs readme ([d97451d](https://github.com/omargallob/mono-repo-release/commit/d97451db7fe503d040e530311101029e9dde2fe6))

## [1.2.1](https://github.com/omargallob/mono-repo-release/compare/v1.2.0...v1.2.1) (2025-10-20)


### Bug Fixes

* **lib2:** included packages to test updates ([#12](https://github.com/omargallob/mono-repo-release/issues/12)) ([a338f35](https://github.com/omargallob/mono-repo-release/commit/a338f35ad96eb07b24f009403b2ebfc33c21dbd8))

## [1.2.0](https://github.com/omargallob/mono-repo-release/compare/v1.1.0...v1.2.0) (2025-10-20)


### Features

* **lib2:** added ascii art function ([#10](https://github.com/omargallob/mono-repo-release/issues/10)) ([4e9b830](https://github.com/omargallob/mono-repo-release/commit/4e9b830309c68089729fa43fc9f3b5514da65d4b))

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
