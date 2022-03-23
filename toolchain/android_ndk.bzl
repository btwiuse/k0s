def _android_ndk_impl(ctx):

    ctx.file("README.md", '# Hello, Android!')

    is_host_ndk = "ANDROID_NDK_HOME" in ctx.os.environ

    is_bazel_ndk = "ANDROID_NDK_BAZEL" in ctx.os.environ

    path_ndk = ctx.os.environ["ANDROID_NDK_HOME"] if is_host_ndk else ctx.path(".").realpath

    print('ANDROID_NDK_HOME =', path_ndk)
    print('ANDROID_NDK_BAZEL =', is_bazel_ndk)
    print('OS =', ctx.os.name)

    if not is_host_ndk:
        if is_bazel_ndk:
            print('downloading android ndk for', ctx.os.name)
            if ctx.os.name == "mac os x":
                ctx.download_and_extract(
                    "https://dl.google.com/android/repository/android-ndk-r23b-darwin.zip",
                    stripPrefix = "android-ndk-r23b",
                    sha256 = "e067b7402fdae85bfbe8af1822afd573b8e73dce443a8292fdaeb2e8dc3aeb86"
                )
            if ctx.os.name == "linux":
                ctx.download_and_extract(
                    "https://dl.google.com/android/repository/android-ndk-r23b-linux.zip",
                    stripPrefix = "android-ndk-r23b",
                    sha256 = "c6e97f9c8cfe5b7be0a9e6c15af8e7a179475b7ded23e2d1c1fa0945d6fb4382"
                )

    ctx.file("template.bzl", 'ANDROID_NDK_HOME = "{}"'.format(path_ndk))
    ctx.template(
        "android_ndk.bzl",
        "template.bzl",
        substitutions = {},
        executable = False,
    )
    ctx.template(
        "BUILD.bazel",
        "template.bzl",
        substitutions = {},
        executable = False,
    )

android_ndk = repository_rule(
    implementation = _android_ndk_impl,
    environ = [
        "ANDROID_NDK_HOME",
        "ANDROID_NDK_BAZEL"
    ],
)
