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
    srcs = [
        "app_suite_test.go",
        "main_test.go",
    ],
    embed = [":app_lib"],
    deps = [
        "//pkg/lib1",
        "//pkg/lib2",
        "//test",
        "@com_github_onsi_ginkgo_v2//:ginkgo",
        "@com_github_onsi_gomega//:gomega",
    ],
)
