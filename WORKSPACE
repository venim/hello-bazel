load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# ---------------------------
# Golang
# ---------------------------

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "91585017debb61982f7054c9688857a2ad1fd823fc3f9cb05048b0025c47d023",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.42.0/rules_go-v0.42.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.42.0/rules_go-v0.42.0.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.21.2")

# ---------------------------
# Gazelle
# ---------------------------

http_archive(
    name = "bazel_gazelle",
    sha256 = "d3fa66a39028e97d76f9e2db8f1b0c11c099e8e01bf363a923074784e451f809",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.33.0/bazel-gazelle-v0.33.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.33.0/bazel-gazelle-v0.33.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("//:tools/go_third_party.bzl", "go_deps")

# gazelle:repository_macro tools/go_third_party.bzl%go_deps
go_deps()

# ---------------------------
# Protobufs
# ---------------------------

# http_archive(
#     name = "com_google_protobuf",
#     sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
#     strip_prefix = "protobuf-3.14.0",
#     urls = [
#         "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
#         "https://github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
#     ],
# )

# load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

# protobuf_deps()

http_archive(
    name = "rules_proto",
    sha256 = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
    strip_prefix = "rules_proto-5.3.0-21.7",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

# ---------------------------
# rules_ts
# ---------------------------

http_archive(
    name = "aspect_rules_ts",
    sha256 = "8aabb2055629a7becae2e77ae828950d3581d7fc3602fe0276e6e039b65092cb",
    strip_prefix = "rules_ts-2.0.0",
    url = "https://github.com/aspect-build/rules_ts/releases/download/v2.0.0/rules_ts-v2.0.0.tar.gz",
)

load("@aspect_rules_ts//ts:repositories.bzl", "LATEST_TYPESCRIPT_VERSION", "rules_ts_dependencies")

rules_ts_dependencies(
    # This keeps the TypeScript version in-sync with the editor, which is typically best.
    # ts_version_from = "//:package.json",

    # Alternatively, you could pick a specific version, or use
    ts_version = LATEST_TYPESCRIPT_VERSION,
)

load("@aspect_rules_js//js:repositories.bzl", "rules_js_dependencies")

rules_js_dependencies()

load("@bazel_features//:deps.bzl", "bazel_features_deps")

bazel_features_deps()

# ---------------------------
# rules_rollup
# ---------------------------

http_archive(
    name = "aspect_rules_rollup",
    sha256 = "a0433a0b0206a45d362749d71bc1e4e0dacf5ca2a572b059328f9753392bca80",
    strip_prefix = "rules_rollup-1.0.0",
    url = "https://github.com/aspect-build/rules_rollup/releases/download/v1.0.0/rules_rollup-v1.0.0.tar.gz",
)

load("@aspect_rules_rollup//rollup:dependencies.bzl", "rules_rollup_dependencies")

rules_rollup_dependencies()

# ---------------------------
# rules_nodejs
# ---------------------------

load("@rules_nodejs//nodejs:repositories.bzl", "DEFAULT_NODE_VERSION", "nodejs_register_toolchains")

nodejs_register_toolchains(
    name = "node",
    node_version = DEFAULT_NODE_VERSION,
)

load("@aspect_rules_js//npm:npm_import.bzl", "npm_translate_lock")

npm_translate_lock(
    name = "npm",
    pnpm_lock = "//:pnpm-lock.yaml",
    verify_node_modules_ignored = "//:.bazelignore",
)

load("@npm//:repositories.bzl", "npm_repositories")

npm_repositories()
