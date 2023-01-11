def _android_ndk_impl(ctx):

    ctx.file("README.md", '# Hello, Android!')

    is_host_ndk = "ANDROID_NDK_HOME" in ctx.os.environ and ctx.os.environ["ANDROID_NDK_HOME"]

    is_bazel_ndk = "ANDROID_NDK_BAZEL" in ctx.os.environ and ctx.os.environ["ANDROID_NDK_BAZEL"]

    path_ndk = ctx.os.environ["ANDROID_NDK_HOME"] if is_host_ndk else ctx.path(".").realpath

    print('ANDROID_NDK_HOME =', path_ndk)
    print('ANDROID_NDK_BAZEL =', is_bazel_ndk)
    print('OS =', ctx.os.name)

    if not is_host_ndk:
        if is_bazel_ndk:
            print('downloading android ndk for', ctx.os.name)
            if ctx.os.name == "mac os x":
                ctx.download_and_extract(
                    "https://dl.google.com/android/repository/android-ndk-r25b-darwin.zip",
                    stripPrefix = "android-ndk-r25b",
                    sha256 = "7e12f1f809878d4f5d5a901809277aa31546d36c10730fade2036d7d95b3607a"
                )
            if ctx.os.name == "linux":
                ctx.download_and_extract(
                    "https://dl.google.com/android/repository/android-ndk-r25b-linux.zip",
                    stripPrefix = "android-ndk-r25b",
                    sha256 = "403ac3e3020dd0db63a848dcaba6ceb2603bf64de90949d5c4361f848e44b005"
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
