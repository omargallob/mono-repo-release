#!/bin/bash
set -euo pipefail

PKG=$1
VERSION=$2

# Update version
./tools/update_versions.sh "$VERSION"

# Generate changelog
git-chglog --config .chglog/config.yml -o "packages/$PKG/CHANGELOG.md" --next-tag "$PKG/v$VERSION"

# Commit changes
git add .
git commit -m "release($PKG): prepare $VERSION"

# Tag release
git tag "$PKG/v$VERSION"

# Run GoReleaser
goreleaser release --config ".goreleaser/$PKG.yaml" --rm-dist --skip-validate

# Push
git push origin main "$PKG/v$VERSION"