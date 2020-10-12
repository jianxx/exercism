load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

## add the external repositories for the Golang toolchain
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "b725e6497741d7fc2d55fcc29a276627d10e43fa5d0bb692692890ae30d98d00",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.24.3/rules_go-v0.24.3.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.24.3/rules_go-v0.24.3.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "b85f48fa105c4403326e9525ad2b2cc437babaa6e15a3fc0b1dbab0ab064bc7c",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.2/bazel-gazelle-v0.22.2.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.2/bazel-gazelle-v0.22.2.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_rules_dependencies()

go_register_toolchains()

gazelle_dependencies()

## add the external repositories for the Rust toolchain
http_archive(
    name = "io_bazel_rules_rust",
    sha256 = "618cba29165b7a893960de7bc48510b0fb182b21a4286e1d3dbacfef89ace906",
    strip_prefix = "rules_rust-5998baf9016eca24fafbad60e15f4125dd1c5f46",
    urls = [
        # Master branch as of 2020-09-24
        "https://github.com/bazelbuild/rules_rust/archive/5998baf9016eca24fafbad60e15f4125dd1c5f46.tar.gz",
    ],
)

load("@io_bazel_rules_rust//rust:repositories.bzl", "rust_repositories")

rust_repositories(
    edition = "2018",
    version = "1.47.0",
)

load("@io_bazel_rules_rust//:workspace.bzl", "rust_workspace")

rust_workspace()

load("//rust/gigasecond/cargo:crates.bzl", "gigasecond_fetch_remote_crates")

gigasecond_fetch_remote_crates()

load("//rust/dot-dsl/cargo:crates.bzl", "dot_dsl_fetch_remote_crates")

dot_dsl_fetch_remote_crates()

load("//rust/robot-name/cargo:crates.bzl", "robot_name_fetch_remote_crates")

robot_name_fetch_remote_crates()

## add the external repositories for the Python toolchain
http_archive(
    name = "rules_python",
    sha256 = "b5668cde8bb6e3515057ef465a35ad712214962f0b3a314e551204266c7be90c",
    strip_prefix = "rules_python-0.0.2",
    url = "https://github.com/bazelbuild/rules_python/releases/download/0.0.2/rules_python-0.0.2.tar.gz",
)

register_toolchains("//python:py_toolchain")
