"""Non-module dependencies for k0s project."""

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
#load("//toolchain:android_ndk.bzl", "android_ndk")

def go_deps():
    pass

def _non_module_dependencies_impl(mctx):
    """Implementation of non_module_dependencies extension."""
    
    # Android NDK repository
    #android_ndk(name = "android_ndk")
    
    # TODO: Migrate docker rules to rules_oci which supports BZLMOD
    # Docker rules - temporarily commented out due to bazel_skylib compatibility issues
    # git_repository(
    #     name = "io_bazel_rules_docker",
    #     branch = "master", 
    #     remote = "https://github.com/bazelbuild/rules_docker.git",
    # )
    
    # Dhall rules
    git_repository(
        name = "rules_dhall",
        branch = "master",
        remote = "https://github.com/humphrej/rules_dhall.git",
    )
    
    # Clojure rules
    git_repository(
        name = "rules_clojure", 
        branch = "master",
        remote = "https://github.com/simuons/rules_clojure.git",
    )
    
    # Grafana rules
    git_repository(
        name = "io_bazel_rules_grafana",
        branch = "main",
        remote = "https://github.com/etsy/rules_grafana.git",
    )

def _setup_non_module_dependencies():
    """Setup function for non-module dependencies."""
    # This would be called after repositories are defined
    # But module extensions can't call load() statements for setup
    # So we'll need to handle setup in MODULE.bazel or WORKSPACE.bzlmod
    pass

non_module_dependencies = module_extension(
    implementation = _non_module_dependencies_impl,
)
