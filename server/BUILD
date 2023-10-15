load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "server_lib",
    srcs = ["server.go"],
    importpath = "github.com/venim/hello-bazel/server",
    visibility = ["//visibility:private"],
    deps = [
        "//proto",
        "//web:static",
        "@com_github_improbable_eng_grpc_web//go/grpcweb",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "server",
    embed = [":server_lib"],
    importpath = "",
    visibility = ["//visibility:public"],
)
