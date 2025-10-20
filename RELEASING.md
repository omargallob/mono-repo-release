# Release Process

This repository uses [release-please](https://github.com/googleapis/release-please) to manage releases and changelogs.

## Making Changes

1. Create a feature branch from `main`
2. Make your changes in the relevant package(s)
3. Use [Conventional Commits](https://www.conventionalcommits.org/) for your commit messages:
   - `fix:` for patches (e.g., `fix(lib1): correct network timeout`)
   - `feat:` for features (e.g., `feat(lib2): add new API endpoint`)
   - `feat!:` or `fix!:` for breaking changes
   - Add scope to indicate package: `fix(lib1):`, `feat(lib2):`, `feat(app):`

## Release Process

Release-please will:
- Create release PRs when commits are pushed to main
- Update changelogs in each package directory
- Create tags in format `NAME_OF_PACKAGE:vX.Y.Z` (e.g., `lib1:v1.2.3`)
- Maintain the root repo version and changelog

### Tags Created

- Repository: `vX.Y.Z` (e.g., `v1.2.3`)
- Packages:
  - `lib1:vX.Y.Z`
  - `lib2:vX.Y.Z`
  - `app:vX.Y.Z`

## Local Testing

To test release-please locally:

```bash
# Install release-please CLI
npm install -g release-please

# Run in dry-run mode
release-please release-pr \
  --dry-run \
  --manifest-file .release-please-manifest.json \
  --config-file .release-please-config.json \
  --repo-url=github.com/omargallob/mono-repo-release
```

## Managing Releases

1. Review and merge PRs with conventional commit messages
2. Release-please will create release PR(s) when changes are detected
3. Review the release PR (check changelog updates and version bumps)
4. Merge the release PR to:
   - Update all affected CHANGELOGs
   - Create GitHub releases
   - Create git tags
   - Bump versions in manifest