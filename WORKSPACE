# gazelle:repository_macro go_repos.bzl%go_repositories
# gazelle:repo bazel_gazelle
workspace(name = "com_github_btwiuse_k0s")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

git_repository(
    name = "bazel_features",
    branch = "main",
    remote = "https://github.com/bazel-contrib/bazel_features.git",
)

load("@bazel_features//:deps.bzl", "bazel_features_deps")

bazel_features_deps()

# load("//toolchain:android_ndk.bzl", "android_ndk")

# android_ndk(name = "android_ndk")

# load("@android_ndk//:android_ndk.bzl", "ANDROID_NDK_HOME")

# android_sdk_repository(name = "androidsdk")

# android_ndk_repository(
#     name = "androidndk",
#     api_level = 30,
#     path = ANDROID_NDK_HOME,
# )

# register_toolchains(
#     "//toolchain:cc-toolchain-mingw",
#     "//toolchain:cc-toolchain-mingw64",
#     "//toolchain:cc-toolchain-android_amd64",
#     "//toolchain:cc-toolchain-android_386",
#     "//toolchain:cc-toolchain-android_arm64",
#     "//toolchain:cc-toolchain-android_armv7",
# )

git_repository(
    name = "rules_python",
    branch = "main",
    remote = "https://github.com/bazelbuild/rules_python.git",
)

git_repository(
    name = "rules_java",
    branch = "master",
    remote = "https://github.com/bazelbuild/rules_java.git",
)

#load("@rules_java//java:java_library.bzl", "java_library")

#http_archive(
#     name = "rules_python",
#     sha256 = "d70cd72a7a4880f0000a6346253414825c19cdd40a28289bdf67b8e6480edff8",
#     branch = "master",
#     strip_prefix = "rules_python-0.28.0",
#     url = "https://github.com/bazelbuild/rules_python/releases/download/0.28.0/rules_python-0.28.0.tar.gz",
#)

load("@rules_python//python:repositories.bzl", "py_repositories")

py_repositories()

git_repository(
    name = "io_bazel_rules_go",
    # commit = "b8fd0bb7a7c384eca8c9c179754cbf6644e67feb",
    branch = "master",
    remote = "https://github.com/bazelbuild/rules_go.git",
    # tag = "v0.32.0",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    nogo = "@//:nogo",
    version = "1.24.4",
)  # nogo is in the top-level BUILD file of this workspace

git_repository(
    name = "com_google_protobuf",
    # commit = "09745575a923640154bcf307fba8aedff47f240a",
    branch = "main",
    remote = "https://github.com/protocolbuffers/protobuf.git",
    # shallow_since = "1558721209 -0700",
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

git_repository(
    name = "rules_rust",
    # commit = "f32695dcd02d9a19e42b9eb7f29a24a8ceb2b858",
    branch = "main",
    remote = "https://github.com/bazelbuild/rules_rust.git",
)

load("@rules_rust//rust:repositories.bzl", "rules_rust_dependencies", "rust_register_toolchains")

rules_rust_dependencies()

rust_register_toolchains(
    edition = "2021",
    versions = ["1.73.0"],
)

# https://docs.rs/crate/cargo-raze/0.0.19
load("//cargo:crates.bzl", "raze_fetch_remote_crates")

raze_fetch_remote_crates()

git_repository(
    name = "rules_proto",
    branch = "main",
    remote = "https://github.com/bazelbuild/rules_proto.git",
    # tag = "6.0.0",
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies")

# load("@rules_proto//proto:setup.bzl", "rules_proto_setup")
load("@rules_proto//proto:toolchains.bzl", "rules_proto_toolchains")

# install gazelle
git_repository(
    name = "bazel_gazelle",
    # commit = "8c270274e8b64ed9c38e3b8025e7760ded83ebcf",
    # tag = "v0.22.3",
    branch = "master",
    remote = "https://github.com/bazelbuild/bazel-gazelle.git",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

# go_repos

# local_repository(
#     name = "starlark",
#     path = "starlark",
# )

# load("@starlark//:defs.bzl", "print_seq")

# print_seq()

# # https://github.com/google/bazel_rules_install
# git_repository(
#     name = "com_github_google_rules_install",
#     branch = "main",
#     # commit = "e93a17ed42a8a622a78fbf4737309e583f4b3cb4",
#     remote = "https://github.com/google/bazel_rules_install.git",
# )

# load("@com_github_google_rules_install//:deps.bzl", "install_rules_dependencies")

# install_rules_dependencies()

# load("@com_github_google_rules_install//:setup.bzl", "install_rules_setup")

# install_rules_setup()

# install bbox
# git_repository(
#     name = "bbox",
#     # commit = "16d0642dda469579fecf2d2e1efff544e30a60c1",
#     branch = "master",
#     remote = "https://github.com/btwiuse/bbox.git",
# )
#
# load("@bbox//:package.bzl", "register_repositories")
#
# register_repositories()
#
load("//:go_repos.bzl", "go_repositories")

go_repositories()

# https://stackoverflow.com/questions/64929217/bazel-gazelle-error-no-such-package-org-golang-x-tools-go-analysis-internal
gazelle_dependencies()

rules_proto_dependencies()

# rules_proto_setup()

rules_proto_toolchains()

# https://github.com/humphrej/rules_dhall
git_repository(
    name = "rules_dhall",
    branch = "master",
    remote = "https://github.com/humphrej/rules_dhall.git",
)

load("@rules_dhall//:deps.bzl", "load_dhall_dependencies", "load_dhall_k8s_dependencies")

load_dhall_dependencies()

load_dhall_k8s_dependencies()

# https://github.com/simuons/rules_clojure
git_repository(
    name = "rules_clojure",
    branch = "master",
    remote = "https://github.com/simuons/rules_clojure.git",
)

load("@rules_clojure//:repositories.bzl", "rules_clojure_dependencies", "rules_clojure_toolchains")

rules_clojure_dependencies()

rules_clojure_toolchains()

http_archive(
    name = "rules_oci",
    sha256 = "4a276e9566c03491649eef63f27c2816cc222f41ccdebd97d2c5159e84917c3b",
    strip_prefix = "rules_oci-1.7.4",
    url = "https://github.com/bazel-contrib/rules_oci/releases/download/v1.7.4/rules_oci-v1.7.4.tar.gz",
)

load("@rules_oci//oci:dependencies.bzl", "rules_oci_dependencies")

rules_oci_dependencies()

load("@rules_oci//oci:repositories.bzl", "LATEST_CRANE_VERSION", "oci_register_toolchains")

oci_register_toolchains(
    name = "oci",
    crane_version = LATEST_CRANE_VERSION,
    # Uncommenting the zot toolchain will cause it to be used instead of crane for some tasks.
    # Note that it does not support docker-format images.
    # zot_version = LATEST_ZOT_VERSION,
)

# You can pull your base images using oci_pull like this:
load("@rules_oci//oci:pull.bzl", "oci_pull")

# https://github.com/bazelbuild/rules_docker#go_image
git_repository(
    name = "io_bazel_rules_docker",
    branch = "master",
    remote = "https://github.com/bazelbuild/rules_docker.git",
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()
# end go_image

load("@io_bazel_rules_docker//container:container.bzl", "container_pull")

container_pull(
    name = "base",
    registry = "docker.io",
    repository = "library/alpine",
    # 'tag' is also supported, but digest is encouraged for reproducibility.
    # digest = "sha256:deadbeef",
    tag = "latest",
)

git_repository(
    name = "io_bazel_rules_grafana",
    # commit = "{HEAD}", # replace with a real commit hash
    branch = "main",
    remote = "https://github.com/etsy/rules_grafana.git",
)

load("@io_bazel_rules_grafana//grafana:workspace.bzl", grafana_repositories = "repositories")

grafana_repositories()
