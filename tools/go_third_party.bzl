"""Go dependencies managed by gazelle"""

load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_deps():
    """Go dependencies"""
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:BjnpXut1btbtgN/6sp+brB2Kbm2LjNXnidYujAVbSoQ=",
        version = "v1.58.3",
    )
