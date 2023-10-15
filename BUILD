load("@bazel_gazelle//:def.bzl", "gazelle")
load("@npm//:defs.bzl", "npm_link_all_packages")
load("@aspect_rules_ts//ts:defs.bzl", "ts_config")
load("@aspect_rules_js//npm:defs.bzl", "npm_link_package")

# gazelle:prefix github.com/venim/hello-bazel
gazelle(name = "gazelle")

gazelle(
    name = "go-get",
    args = [
        "-to_macro=tools/go_third_party.bzl%go_deps",
    ],
    command = "update-repos",
)

ts_config(
    name = "tsconfig",
    src = "tsconfig.json",
    visibility = ["//visibility:public"],
)

npm_link_package(
    name = "node_modules/@hello-bazel/proto",
    src = "//proto:proto_ts_lib",
    visibility = ["//visibility:public"],
)

npm_link_all_packages(name = "node_modules")
