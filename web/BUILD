load("@aspect_rules_ts//ts:defs.bzl", "ts_project")
load("@aspect_rules_rollup//rollup:defs.bzl", "rollup")

ts_project(
    name = "web_ts",
    srcs = ["main.ts"],
    tsconfig = "//:tsconfig",
    deps = [
        "//:node_modules/@connectrpc/connect",
        "//:node_modules/@connectrpc/connect-web",
        "//proto:proto_ts_proto",
    ],
)

rollup(
    name = "bundle",
    config_file = ":rollup.config.mjs",
    entry_points = {
        "main.js": "bundle",
    },
    node_modules = "//:node_modules",
    sourcemap = "false",
    visibility = ["//visibility:public"],
    deps = [
        "//:node_modules/@bufbuild/protobuf",
        "//:node_modules/@connectrpc/connect",
        "//:node_modules/@connectrpc/connect-web",
        "//:node_modules/@hello-bazel/proto",
    ],
)
