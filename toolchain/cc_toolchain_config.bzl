load("@bazel_tools//tools/build_defs/cc:action_names.bzl", "ACTION_NAMES")
load(
   "@bazel_tools//tools/cpp:cc_toolchain_config_lib.bzl",
   "feature",
   "flag_group",
   "flag_set",
   "tool_path",
)

def _impl_32(ctx):
    # TODO: make hermetic
    tool_paths = [
        tool_path(
            name = "gcc",
            path = "/usr/bin/i686-w64-mingw32-gcc",
        ),
        tool_path(
            name = "ld",
            path = "/usr/bin/i686-w64-mingw32-ld",
        ),
        tool_path(
            name = "ar",
            path = "/usr/bin/i686-w64-mingw32-ar",
        ),
        tool_path(
            name = "cpp",
            path = "/usr/bin/i686-w64-mingw32-g++",
        ),
        tool_path(
            name = "gcov",
            path = "/bin/false",
        ),
        tool_path(
            name = "nm",
            path = "/bin/false",
        ),
        tool_path(
            name = "objdump",
            path = "/bin/false",
        ),
        tool_path(
            name = "strip",
            path = "/bin/false",
        ),
    ]
    features = [
        feature(
            name = "default_linker_flags",
            enabled = True,
            flag_sets = [
                flag_set(
                    actions = [
                        ACTION_NAMES.cpp_link_executable,
                        ACTION_NAMES.cpp_link_dynamic_library,
                        ACTION_NAMES.cpp_link_nodeps_dynamic_library,
                    ],
                    flag_groups = ([
                        flag_group(
                            flags = [
                                "-lstdc++",
                            ],
                        ),
                    ]),
                ),
            ],
        ),
    ]
    return cc_common.create_cc_toolchain_config_info(
        ctx = ctx,
        cxx_builtin_include_directories = [
            "/usr/i686-w64-mingw32/include",
            "/usr/lib/gcc/i686-w64-mingw32"
        ],
        features = features,
        toolchain_identifier = "local",
        host_system_name = "local",
        target_system_name = "windows",
        target_cpu = "@platforms//cpu:x86_32",
        target_libc = "unknown",
        compiler = "mingw",
        abi_version = "unknown",
        abi_libc_version = "unknown",
        tool_paths = tool_paths,
    )

cc_toolchain_config_32 = rule(
    implementation = _impl_32,
    attrs = {},
    provides = [CcToolchainConfigInfo],
)

def _impl_64(ctx):
    # TODO: make hermetic
    tool_paths = [
        tool_path(
            name = "gcc",
            path = "/usr/bin/x86_64-w64-mingw32-gcc",
        ),
        tool_path(
            name = "ld",
            path = "/usr/bin/x86_64-w64-mingw32-ld",
        ),
        tool_path(
            name = "ar",
            path = "/usr/bin/x86_64-w64-mingw32-ar",
        ),
        tool_path(
            name = "cpp",
            path = "/usr/bin/x86_64-w64-mingw32-g++",
        ),
        tool_path(
            name = "gcov",
            path = "/bin/false",
        ),
        tool_path(
            name = "nm",
            path = "/bin/false",
        ),
        tool_path(
            name = "objdump",
            path = "/bin/false",
        ),
        tool_path(
            name = "strip",
            path = "/bin/false",
        ),
    ]
    features = [
        feature(
            name = "default_linker_flags",
            enabled = True,
            flag_sets = [
                flag_set(
                    actions = [
                        ACTION_NAMES.cpp_link_executable,
                        ACTION_NAMES.cpp_link_dynamic_library,
                        ACTION_NAMES.cpp_link_nodeps_dynamic_library,
                    ],
                    flag_groups = ([
                        flag_group(
                            flags = [
                                "-lstdc++",
                            ],
                        ),
                    ]),
                ),
            ],
        ),
    ]
    return cc_common.create_cc_toolchain_config_info(
        ctx = ctx,
        cxx_builtin_include_directories = [
            "/usr/x86_64-w64-mingw32/include",
            "/usr/lib/gcc/x86_64-w64-mingw32"
        ],
        features = features,
        toolchain_identifier = "local",
        host_system_name = "local",
        target_system_name = "windows",
        target_cpu = "@platforms//cpu:x86_64",
        target_libc = "unknown",
        compiler = "mingw64",
        abi_version = "unknown",
        abi_libc_version = "unknown",
        tool_paths = tool_paths,
    )

cc_toolchain_config_64 = rule(
    implementation = _impl_64,
    attrs = {},
    provides = [CcToolchainConfigInfo],
)
