# TODO: add unit tests instead of only testing building
load("@bazel_skylib//rules:build_test.bzl", "build_test")
load("@npm//:vite/package_json.bzl", vite_bin = "bin")
load("@npm//:vue-tsc/package_json.bzl", vue_tsc_bin = "bin")

package(default_visibility = ["//visibility:public"])

SRCS = [
    "env.d.ts",
    "index.html",
    "package.json",
    "vite.config.ts",
    "tsconfig.json",
    "tsconfig.config.json",
    "//frontend/app/src",
]

BUILD_DEPS = ["//:node_modules/" + d for d in [
    "@vitejs/plugin-vue",
    "@vitejs/plugin-vue-jsx",
    "@vue/tsconfig",
    "vite",
    "vue",
    "vue-router",
]]

vite_bin.vite(
    name = "build",
    srcs = SRCS + BUILD_DEPS,
    args = ["build"],
    chdir = package_name(),
    out_dirs = ["dist"],
)

vue_tsc_bin.vue_tsc_test(
    name = "type-check",
    args = ["--noEmit"],
    chdir = package_name(),
    data = SRCS + BUILD_DEPS,
    include_declarations = True,
)

vite_bin.vite_binary(
    name = "vite",
    chdir = package_name(),
    data = SRCS + BUILD_DEPS,
)

build_test(
    name = "build_test",
    targets = [
        ":build",
        ":type-check",
    ],
)
