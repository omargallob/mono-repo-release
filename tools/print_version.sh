#!/bin/sh
# Extracts the Version variable from pkg/lib2/version.go for Bazel stamping or scripting.


# Accept file path as first argument, default to ../pkg/lib2/version.go if not provided
if [ -n "$1" ]; then
  VERSION_FILE="$1"
else
  VERSION_FILE="$(dirname "$0")/../pkg/lib2/version.go"
fi

if [ ! -f "$VERSION_FILE" ]; then
  echo "version.go not found at $VERSION_FILE" >&2
  exit 1
fi

# Output in Bazel workspace status format: STABLE_VERSION <value>
grep 'Version' "$VERSION_FILE" | awk -F'"' '{print "STABLE_VERSION "$2}'
