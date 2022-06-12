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
                    "https://dl.google.com/android/repository/android-ndk-r24-darwin.zip",
                    stripPrefix = "android-ndk-r24",
                    sha256 = "162a7000515be07489f2ed70d6d3a117d236150f83f3fcb601c163349429ba23"
                )
            if ctx.os.name == "linux":
                ctx.download_and_extract(
                    "https://dl.google.com/android/repository/android-ndk-r24-linux.zip",
                    stripPrefix = "android-ndk-r24",
                    sha256 = "caac638f060347c9aae994e718ba00bb18413498d8e0ad4e12e1482964032997"
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
