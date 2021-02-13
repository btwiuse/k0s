def _android_ndk_impl(ctx):
    if ("ANDROID_NDK_HOME") not in ctx.os.environ:
        fail("Please specify the NDK path using ANDROID_NDK_HOME")

    ctx.file("template.bzl", 'ANDROID_NDK = "{}"'.format(ctx.os.environ["ANDROID_NDK_HOME"]))
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
