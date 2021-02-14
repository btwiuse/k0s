def _android_ndk_impl(ctx):

    home = "ANDROID_NDK_HOME"

    if ("ANDROID_NDK_HOME") not in ctx.os.environ:
        print("Please specify the NDK path using ANDROID_NDK_HOME")
    else:
        home = ctx.os.environ["ANDROID_NDK_HOME"]

    print('ANDROID_NDK_HOME = "{}"'.format(home))

    ctx.file("template.bzl", 'ANDROID_NDK_HOME = "{}"'.format(home))
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
    local = True,
)
