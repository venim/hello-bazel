load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/venim/hello-bazel
gazelle(name = "gazelle")

gazelle(
    name = "go-install",
    args = [
        "-to_macro=tools/go_third_party.bzl%go_deps",
    ],
    command = "update-repos",
)
