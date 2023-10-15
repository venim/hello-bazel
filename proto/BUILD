load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@aspect_rules_ts//ts:proto.bzl", "ts_proto_library")
load("@aspect_rules_js//npm:defs.bzl", "npm_package")

package(default_visibility = ["//visibility:public"])

proto_library(
    name = "hello_proto",
    srcs = ["hello.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "proto_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/venim/hello-bazel/proto",
    proto = ":hello_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "proto",
    embed = [":proto_go_proto"],
    importpath = "github.com/venim/hello-bazel/proto",
    visibility = ["//visibility:public"],
)

ts_proto_library(
    name = "proto_ts_proto",
    node_modules = "//:node_modules",
    proto = ":hello_proto",
)

npm_package(
    name = "proto_ts_lib",
    srcs = [
        ":proto_ts_proto",
    ],
    package = "@hello-bazel/proto",
)
