load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "app_lib",
    srcs = ["main.go"],
    importpath = "github.com/omargallob/mono-repo-release/cmd/app",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/lib1",
        "//pkg/lib2",
    ],
)

go_library(
    name = "app",
    srcs = ["main.go"],
    importpath = "github.com/omargallob/mono-repo-release/cmd/app",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/lib1",
        "//pkg/lib2",
    ],
)

go_test(
    name = "app_test",
    srcs = ["main_test.go"],
    embed = [":app_lib"],
    deps = [
        "//pkg/lib1",
        "//pkg/lib2",
    ],
)

sh_test(
    name = "golangci-lint_test",
    srcs = ["lint.sh"],
    data = [
        "//tools/golangci-lint:golangci-lint",
        "//:go_sources",  # All Go files
    ],
)
