load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "server_lib",
    srcs = ["server.go"],
    importpath = "github.com/venim/hello-bazel/server",
    visibility = ["//visibility:private"],
    deps = [
        "//proto",
        "//web",
        "@com_github_improbable_eng_grpc_web//go/grpcweb",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//health",
        "@org_golang_google_grpc//health/grpc_health_v1",
        "@org_golang_x_net//http2",
        "@org_golang_x_net//http2/h2c",
    ],
)

go_binary(
    name = "server",
    embed = [":server_lib"],
    importpath = "",
    visibility = ["//visibility:public"],
)

go_test(
    name = "server_test",
    srcs = ["server_test.go"],
    embed = [":server_lib"],
    deps = [
        "//proto",
        "@com_github_google_go_cmp//cmp",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//credentials/insecure",
    ],
)
