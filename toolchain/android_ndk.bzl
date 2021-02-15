def _android_ndk_impl(ctx):

    ctx.file("README.md", '# Hello, Android!')

    host_ndk = "ANDROID_NDK_HOME" in ctx.os.environ

    path_ndk = ctx.os.environ["ANDROID_NDK_HOME"] if host_ndk else ctx.path(".").realpath

    print('ANDROID_NDK_HOME =', path_ndk)

    if not host_ndk:
        ctx.download_and_extract("https://dl.google.com/android/repository/android-ndk-r22-linux-x86_64.zip", stripPrefix = "android-ndk-r22")

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
        "ANDROID_NDK_HOME"
    ],
)
