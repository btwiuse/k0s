workspace(name = "k0s")

load("//toolchain:android_ndk.bzl", "android_ndk")

android_ndk(name = "android_ndk")

load("@android_ndk//:android_ndk.bzl", "ANDROID_NDK_HOME")

android_sdk_repository(name = "androidsdk")

android_ndk_repository(
    name = "androidndk",
    api_level = 30,
    path = ANDROID_NDK_HOME,
)

register_toolchains(
    "//toolchain:cc-toolchain-mingw",
    "//toolchain:cc-toolchain-mingw64",
    "//toolchain:cc-toolchain-android_amd64",
    "//toolchain:cc-toolchain-android_386",
    "//toolchain:cc-toolchain-android_arm64",
    "//toolchain:cc-toolchain-android_armv7",
)

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "io_bazel_rules_go",
    # commit = "2a0e3a07e9ed9aa9b7afd1a222638ba52166e52d",
    branch = "master",
    remote = "https://github.com/bazelbuild/rules_go.git",
    # tag = "v0.24.3",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    nogo = "@//:nogo",
    version = "1.18.2",
)  # nogo is in the top-level BUILD file of this workspace

git_repository(
    name = "bazel_gazelle",
    # commit = "8c270274e8b64ed9c38e3b8025e7760ded83ebcf",
    # tag = "v0.22.3",
    branch = "master",
    remote = "https://github.com/bazelbuild/bazel-gazelle.git",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

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

rust_register_toolchains()

load("@rules_rust//rust:repositories.bzl", "rust_repositories")

# https://bazelbuild.github.io/rules_rust/
# rust_repositories(
#     edition = "2018",
#     rustfmt_version = "1.50.0",
#     version = "1.50.0",
# )
rust_repositories(
    edition = "2021",
    iso_date = "2022-05-25",
    version = "nightly",
)

# https://docs.rs/crate/cargo-raze/0.0.19
# load("//cargo:crates.bzl", "raze_fetch_remote_crates")

# raze_fetch_remote_crates()

# https://github.com/google/cargo-raze#using-cargo-raze-through-bazel
git_repository(
    name = "cargo_raze",
    branch = "main",
    remote = "https://github.com/google/cargo-raze",
)

load("@cargo_raze//:repositories.bzl", "cargo_raze_repositories")

cargo_raze_repositories()

load("@cargo_raze//:transitive_deps.bzl", "cargo_raze_transitive_deps")

cargo_raze_transitive_deps()

git_repository(
    name = "rules_proto",
    branch = "master",
    remote = "https://github.com/bazelbuild/rules_proto.git",
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

go_repository(
    name = "com_github_xtaci_kcp_go",
    importpath = "github.com/xtaci/kcp-go",
    sum = "h1:TN1uey3Raw0sTz0Fg8GkfM0uH3YwzhnZWQ1bABv5xAg=",
    version = "v5.4.20+incompatible",
)

go_repository(
    name = "com_github_xtaci_lossyconn",
    importpath = "github.com/xtaci/lossyconn",
    sum = "h1:EWU6Pktpas0n8lLQwDsRyZfmkPeRbdgPtW609es+/9E=",
    version = "v0.0.0-20200209145036-adba10fffc37",
)

go_repository(
    name = "com_github_xtaci_smux",
    importpath = "github.com/xtaci/smux",
    sum = "h1:FBPYOkW8ZTjLKUM4LI4xnnuuDC8CQ/dB04HD519WoEk=",
    version = "v1.5.16",
)

go_repository(
    name = "com_github_felixge_httpsnoop",
    importpath = "github.com/felixge/httpsnoop",
    sum = "h1:+nS9g82KMXccJ/wp0zyRW9ZBHFETmMGtkk+2CTTrW4o=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest",
    importpath = "github.com/Azure/go-autorest/autorest",
    sum = "h1:90Y4srNYrwOtAgVo3ndrQkTYn6kf1Eg/AjTFJ8Is2aM=",
    version = "v0.11.18",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_adal",
    importpath = "github.com/Azure/go-autorest/autorest/adal",
    sum = "h1:Mp5hbtOePIzM8pJVRa3YLrWWmZtoxRXqUEzCfJt3+/Q=",
    version = "v0.9.13",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_date",
    importpath = "github.com/Azure/go-autorest/autorest/date",
    sum = "h1:7gUk1U5M/CQbp9WoqinNzJar+8KY+LPI6wiWrP/myHw=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_mocks",
    importpath = "github.com/Azure/go-autorest/autorest/mocks",
    sum = "h1:K0laFcLE6VLTOwNgSxaGbUcLPuGXlNkbVvq4cW4nIHk=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_azure_go_autorest_logger",
    importpath = "github.com/Azure/go-autorest/logger",
    sum = "h1:IG7i4p/mDa2Ce4TRyAO8IHnVhAVF3RFU+ZtXWSmf4Tg=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_azure_go_autorest_tracing",
    importpath = "github.com/Azure/go-autorest/tracing",
    sum = "h1:TYi4+3m5t6K48TGI9AUdb+IzbnSxvnvUMfuitfgcfuo=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_blang_semver",
    importpath = "github.com/blang/semver",
    sum = "h1:cQNTCjp13qL8KC3Nbxr/y2Bqb63oX6wdnnjpJbkM4JQ=",
    version = "v3.5.1+incompatible",
)

go_repository(
    name = "com_github_brancz_gojsontoyaml",
    importpath = "github.com/brancz/gojsontoyaml",
    sum = "h1:DMb8SuAL9+demT8equqMMzD8C/uxqWmj4cgV7ufrpQo=",
    version = "v0.0.0-20190425155809-e8bd32d46b3d",
)

go_repository(
    name = "com_github_burntsushi_xgb",
    importpath = "github.com/BurntSushi/xgb",
    sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
    version = "v0.0.0-20160522181843-27f122750802",
)

go_repository(
    name = "com_github_campoy_embedmd",
    importpath = "github.com/campoy/embedmd",
    sum = "h1:V4kI2qTJJLf4J29RzI/MAt2c3Bl4dQSYPuflzwFH2hY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_chzyer_readline",
    importpath = "github.com/chzyer/readline",
    sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
    version = "v0.0.0-20180603132655-2972be24d48e",
)

go_repository(
    name = "com_github_dgryski_go_jump",
    importpath = "github.com/dgryski/go-jump",
    sum = "h1:qZNIK8jjHgLFHAW2wzCWPEv0ZIgcBhU7X3oDt/p3Sv0=",
    version = "v0.0.0-20170409065014-e1f439676b57",
)

go_repository(
    name = "com_github_docker_spdystream",
    importpath = "github.com/docker/spdystream",
    sum = "h1:cenwrSVm+Z7QLSV/BsnenAOcDXdX4cMv4wP0B/5QbPg=",
    version = "v0.0.0-20160310174837-449fdfce4d96",
)

go_repository(
    name = "com_github_docopt_docopt_go",
    importpath = "github.com/docopt/docopt-go",
    sum = "h1:bWDMxwH3px2JBh6AyO7hdCn/PkvCZXii8TGj7sbtEbQ=",
    version = "v0.0.0-20180111231733-ee0de3bc6815",
)

go_repository(
    name = "com_github_elazarl_goproxy",
    importpath = "github.com/elazarl/goproxy",
    sum = "h1:pEtiCjIXx3RvGjlUJuCNxNOw0MNblyR9Wi+vJGBFh+8=",
    version = "v0.0.0-20191011121108-aa519ddbe484",
)

go_repository(
    name = "com_github_emicklei_go_restful",
    importpath = "github.com/emicklei/go-restful",
    sum = "h1:8KpYO/Xl/ZudZs5RNOEhWMBY4hmzlZhhRd9cu+jrZP4=",
    version = "v2.15.0+incompatible",
)

go_repository(
    name = "com_github_evanphx_json_patch",
    importpath = "github.com/evanphx/json-patch",
    sum = "h1:4onqiflcdA9EOZ4RxV643DvftH5pOlLGNtQ5lPWQu84=",
    version = "v4.12.0+incompatible",
)

go_repository(
    name = "com_github_go_gl_glfw_v3_3_glfw",
    importpath = "github.com/go-gl/glfw/v3.3/glfw",
    sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
    version = "v0.0.0-20200222043503-6f7a984d4dc4",
)

go_repository(
    name = "com_github_go_logr_logr",
    importpath = "github.com/go-logr/logr",
    sum = "h1:ahHml/yUpnlb96Rp8HCvtYVPY8ZYpxq3g7UYchIYwbs=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_go_openapi_jsonpointer",
    importpath = "github.com/go-openapi/jsonpointer",
    sum = "h1:gZr+CIYByUqjcgeLXnQu2gHYQC9o73G2XUeOFYEICuY=",
    version = "v0.19.5",
)

go_repository(
    name = "com_github_go_openapi_jsonreference",
    importpath = "github.com/go-openapi/jsonreference",
    sum = "h1:1WJP/wi4OjB4iV8KVbH73rQaoialJrqv8gitZLxGLtM=",
    version = "v0.19.5",
)

go_repository(
    name = "com_github_go_openapi_spec",
    importpath = "github.com/go-openapi/spec",
    sum = "h1:uH9RQ6vdyPSs2pSy9fL8QPspDF2AMIMPtmK5coSSjtQ=",
    version = "v0.20.3",
)

go_repository(
    name = "com_github_go_openapi_swag",
    importpath = "github.com/go-openapi/swag",
    sum = "h1:gm3vOOXfiuw5i9p5N9xJvfjvuofpyvLA9Wr6QfK5Fng=",
    version = "v0.19.14",
)

go_repository(
    name = "com_github_google_go_jsonnet",
    importpath = "github.com/google/go-jsonnet",
    sum = "h1:as/sAfmjOHqY/OMBR4mv9I8ZY0/jNuqN3u44AicwxPs=",
    version = "v0.14.0",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    sum = "h1:xmapqc1AyLoB+ddYT6r04bD9lIjlOqGaREovi0SzFaE=",
    version = "v2.1.1-0.20190517191504-25dcb96d9e51+incompatible",
)

go_repository(
    name = "com_github_google_pprof",
    importpath = "github.com/google/pprof",
    sum = "h1:K6RDEckDVWvDI9JAJYCmNdQXq6neHJOYx3V6jnqNEec=",
    version = "v0.0.0-20210720184732-4bb14d4b1be1",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    importpath = "github.com/googleapis/gax-go/v2",
    sum = "h1:dp3bWCh+PPO1zjRRiCSczJav13sBvG4UhNyVTa1KqdU=",
    version = "v2.1.1",
)

go_repository(
    name = "com_github_googleapis_gnostic",
    importpath = "github.com/googleapis/gnostic",
    sum = "h1:9fHAtK0uDfpveeqqo1hkEZJcFvYXAiCN3UutL8F9xHw=",
    version = "v0.5.5",
)

go_repository(
    name = "com_github_gophercloud_gophercloud",
    importpath = "github.com/gophercloud/gophercloud",
    sum = "h1:6sjpKIpVwRIIwmcEGp+WwNovNsem+c+2vm6oxshRpL8=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_gregjones_httpcache",
    importpath = "github.com/gregjones/httpcache",
    sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
    version = "v0.0.0-20180305231024-9cad4c3443a7",
)

go_repository(
    name = "com_github_ianlancetaylor_demangle",
    importpath = "github.com/ianlancetaylor/demangle",
    sum = "h1:mV02weKRL81bEnm8A0HT1/CAelMQDBuQIfLw8n+d6xI=",
    version = "v0.0.0-20200824232613-28f6c0f3b639",
)

go_repository(
    name = "com_github_imdario_mergo",
    importpath = "github.com/imdario/mergo",
    sum = "h1:b6R2BslTbIEToALKP7LxUvijTsNI9TAe80pLWN2g/HU=",
    version = "v0.3.12",
)

go_repository(
    name = "com_github_jsonnet_bundler_jsonnet_bundler",
    importpath = "github.com/jsonnet-bundler/jsonnet-bundler",
    sum = "h1:QFLRnKsv79DTZWbqJRjvMMtP0I6wrtNeMZekqRIRFLE=",
    version = "v0.4.1-0.20200708074244-ada055a225fa",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    importpath = "github.com/jstemmer/go-junit-report",
    sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_mailru_easyjson",
    importpath = "github.com/mailru/easyjson",
    sum = "h1:8yTIVnZgCoiM1TgqoeTl+LfU5Jg6/xL3QhGQnimLYnA=",
    version = "v0.7.6",
)

go_repository(
    name = "com_github_munnerz_goautoneg",
    importpath = "github.com/munnerz/goautoneg",
    sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
    version = "v0.0.0-20191010083416-a7dc8b61c822",
)

go_repository(
    name = "com_github_mxk_go_flowrate",
    importpath = "github.com/mxk/go-flowrate",
    sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
    version = "v0.0.0-20140419014527-cca7078d478f",
)

go_repository(
    name = "com_github_nytimes_gziphandler",
    importpath = "github.com/NYTimes/gziphandler",
    sum = "h1:ZUDjpQae29j0ryrS0u/B8HZfJBtBQHjqw2rQ2cqUQ3I=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_peterbourgon_diskv",
    importpath = "github.com/peterbourgon/diskv",
    sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
    version = "v2.0.1+incompatible",
)

go_repository(
    name = "com_github_prometheus_prometheus",
    importpath = "github.com/prometheus/prometheus",
    sum = "h1:7QPitgO2kOFG8ecuRn9O/4L9+10He72rVRJvMXrE9Hg=",
    version = "v2.5.0+incompatible",
)

go_repository(
    name = "com_github_puerkitobio_purell",
    importpath = "github.com/PuerkitoBio/purell",
    sum = "h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_puerkitobio_urlesc",
    importpath = "github.com/PuerkitoBio/urlesc",
    sum = "h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=",
    version = "v0.0.0-20170810143723-de5bf2ad4578",
)

go_repository(
    name = "com_github_robfig_cron_v3",
    importpath = "github.com/robfig/cron/v3",
    sum = "h1:WdRxkvbJztn8LMz/QEvLN5sBU+xKpSqwwUO1Pjr4qDs=",
    version = "v3.0.1",
)

go_repository(
    name = "com_github_sergi_go_diff",
    importpath = "github.com/sergi/go-diff",
    sum = "h1:XU+rvMAioB0UC3q1MFrIQy4Vo5/4VsRDQQXHsEya6xQ=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    sum = "h1:xoax2sJ2DT8S8xA2paPFjDCScCNeWsg75VG0DLRreiY=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_bigquery",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:PQcPefKFdaIzjQFbiyOgAqyx8q5djaE7x9Sqe712DPA=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_datastore",
    importpath = "cloud.google.com/go/datastore",
    sum = "h1:/May9ojXjRkPBNVrq+oWLqmWCkr4OU5uRY29bu0mRyQ=",
    version = "v1.1.0",
)

go_repository(
    name = "com_google_cloud_go_pubsub",
    importpath = "cloud.google.com/go/pubsub",
    sum = "h1:ukjixP1wl0LpnZ6LWtZJ0mX5tBmjp1f8Sqer8Z2OMUU=",
    version = "v1.3.1",
)

go_repository(
    name = "com_google_cloud_go_storage",
    importpath = "cloud.google.com/go/storage",
    sum = "h1:STgFzyU5/8miMl0//zKh2aQeTyeaUH3WN9bSUiJ09bA=",
    version = "v1.10.0",
)

go_repository(
    name = "com_shuralyov_dmitri_gpu_mtl",
    importpath = "dmitri.shuralyov.com/gpu/mtl",
    sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
    version = "v0.0.0-20190408044501-666a987793e9",
)

go_repository(
    name = "in_gopkg_inf_v0",
    importpath = "gopkg.in/inf.v0",
    sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
    version = "v0.9.1",
)

go_repository(
    name = "io_k8s_api",
    importpath = "k8s.io/api",
    replace = "k8s.io/api",
    sum = "h1:xZjKidCirayzX6tHONRQyTNDVIR55TYVqgATqo6ZULY=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_apimachinery",
    importpath = "k8s.io/apimachinery",
    replace = "k8s.io/apimachinery",
    sum = "h1:vhxQ0PPUUU2Ns1b9r4/UFp13UPs8cw2iOoTjnY9faa0=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_autoscaler_vertical_pod_autoscaler",
    importpath = "k8s.io/autoscaler/vertical-pod-autoscaler",
    sum = "h1:mSO9phIinHH3bRuXQkHMIA2uJ4i1WSayDeYr8J3bGjU=",
    version = "v0.10.0",
)

go_repository(
    name = "io_k8s_client_go",
    importpath = "k8s.io/client-go",
    replace = "k8s.io/client-go",
    sum = "h1:85crgh1IotNkLpKYKZHVNI1JT86nr/iDCvq2iWKsql4=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_code_generator",
    importpath = "k8s.io/code-generator",
    replace = "k8s.io/code-generator",
    sum = "h1:FhilVnvwMFVs65SxIQjXSOznGmzJIZEk3CCk/SULBfk=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_component_base",
    importpath = "k8s.io/component-base",
    replace = "k8s.io/component-base",
    sum = "h1:gdvPs4G11e99meQnW4zN+oYOjH8qkLz1sURrAzvKWqc=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_gengo",
    importpath = "k8s.io/gengo",
    sum = "h1:GohjlNKauSai7gN4wsJkeZ3WAJx4Sh+oT/b5IYn5suA=",
    version = "v0.0.0-20210813121822-485abfe95c7c",
)

go_repository(
    name = "io_k8s_klog",
    importpath = "k8s.io/klog",
    sum = "h1:Pt+yjF5aB1xDSVbau4VsWe+dQNzA0qv1LlXdC2dF6Q8=",
    version = "v1.0.0",
)

go_repository(
    name = "io_k8s_klog_v2",
    importpath = "k8s.io/klog/v2",
    sum = "h1:VW25q3bZx9uE3vvdL6M8ezOX79vA2Aq1nEWLqNQclHc=",
    version = "v2.60.1",
)

go_repository(
    name = "io_k8s_kube_openapi",
    importpath = "k8s.io/kube-openapi",
    sum = "h1:E3J9oCLlaobFUqsjG9DfKbP2BmgwBL2p7pn0A3dG9W4=",
    version = "v0.0.0-20211115234752-e816edb12b65",
)

go_repository(
    name = "io_k8s_kube_state_metrics_v2",
    importpath = "k8s.io/kube-state-metrics/v2",
    replace = "github.com/btwiuse/k16s/v2",
    sum = "h1:gOLR+C34lEfYf3VaQoeX+ekL+gv70qi4O/4Au/YznNE=",
    version = "v2.0.0-beta.0.20201224174453-2114e07844a9",
)

go_repository(
    name = "io_k8s_metrics",
    importpath = "k8s.io/metrics",
    replace = "k8s.io/metrics",
    sum = "h1:SxpF5zcFbUCvF3qzY6WPicp4VVFn9VCMHxnEvrwWJoQ=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_sigs_structured_merge_diff_v3",
    importpath = "sigs.k8s.io/structured-merge-diff/v3",
    sum = "h1:dOmIZBMfhcHS09XZkMyUgkq5trg3/jRyJYFZUiaOp8E=",
    version = "v3.0.0",
)

go_repository(
    name = "io_k8s_sigs_structured_merge_diff_v4",
    importpath = "sigs.k8s.io/structured-merge-diff/v4",
    sum = "h1:bKCqE9GvQ5tiVHn5rfn1r+yao3aLQEaLzkkmAkf+A6Y=",
    version = "v4.2.1",
)

go_repository(
    name = "io_k8s_utils",
    importpath = "k8s.io/utils",
    sum = "h1:wxEMGetGMur3J1xuGLQY7GEQYg9bZxKn3tKo5k/eYcs=",
    version = "v0.0.0-20210930125809-cb0fa318a74b",
)

go_repository(
    name = "io_rsc_binaryregexp",
    importpath = "rsc.io/binaryregexp",
    sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
    version = "v0.2.0",
)

go_repository(
    name = "org_golang_x_image",
    importpath = "golang.org/x/image",
    sum = "h1:hVwzHzIUGRjiF7EcUjqNxk3NCfkPxbDKRdnNE1Rpg0U=",
    version = "v0.0.0-20191009234506-e7c1f5e7dbb8",
)

go_repository(
    name = "org_golang_x_mobile",
    importpath = "golang.org/x/mobile",
    sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
    version = "v0.0.0-20190719004257-d2bd2a29d028",
)

go_repository(
    name = "com_github_btwiuse_bcrypt",
    importpath = "github.com/btwiuse/bcrypt",
    sum = "h1:je6QGuEHitc4zAnAL/2w25WyHcSu+Iz/e5t4TRlvG1s=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_infobloxopen_go_trees",
    importpath = "github.com/infobloxopen/go-trees",
    sum = "h1:w66aaP3c6SIQ0pi3QH1Tb4AMO3aWoEPxd1CNvLphbkA=",
    version = "v0.0.0-20200715205103-96a057b8dfb9",
)

go_repository(
    name = "com_github_milosgajdos_tenus",
    importpath = "github.com/milosgajdos/tenus",
    sum = "h1:jmaJzwaY1DUyYVD0lM4U+uvP2kkEg1VahDqRFxIkVBE=",
    version = "v0.0.3",
)

go_repository(
    name = "com_gitea_lunny_log",
    importpath = "gitea.com/lunny/log",
    sum = "h1:r1en/D7xJmcY24VkHkjkcJFa+7ZWubVWPBrvsHkmHxk=",
    version = "v0.0.0-20190322053110-01b5df579c4e",
)

go_repository(
    name = "com_github_lunny_tango",
    importpath = "github.com/lunny/tango",
    sum = "h1:QeUe+2ksZ3LScC+SKhDbS1wbS/ctuyRnZ3fAsL10J4M=",
    version = "v0.5.6",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    importpath = "github.com/mattn/go-sqlite3",
    sum = "h1:jbhqpg7tQe4SupckyijYiy0mJJ/pRyHvXf7JdWK860o=",
    version = "v1.10.0",
)

go_repository(
    name = "com_github_tango_contrib_basicauth",
    importpath = "github.com/tango-contrib/basicauth",
    sum = "h1:tfB+xuTYq48HTLSXD5V99WAn+W+4nqM4KqZAv2ABmfY=",
    version = "v0.0.0-20170526072948-7fbc19aece86",
)

go_repository(
    name = "com_github_buildkite_agent_v3",
    importpath = "github.com/buildkite/agent/v3",
    replace = "github.com/btwiuse/agent/v3",
    sum = "h1:kqXi8K/7hW2+r6mUS6GpfW6GCZjjweqTJ1hqGRbspdI=",
    version = "v3.27.1-0.20210217080418-ae42a28eefa7",
)

go_repository(
    name = "com_github_buildkite_bintest_v3",
    importpath = "github.com/buildkite/bintest/v3",
    sum = "h1:auJ22sFpyy7t6xR7p5FcqAwpNgkP0AyVhEMSRErFmk0=",
    version = "v3.1.0",
)

go_repository(
    name = "com_github_buildkite_interpolate",
    importpath = "github.com/buildkite/interpolate",
    sum = "h1:k6UDF1uPYOs0iy1HPeotNa155qXRWrzKnqAaGXHLZCE=",
    version = "v0.0.0-20200526001904-07f35b4ae251",
)

go_repository(
    name = "com_github_buildkite_shellwords",
    importpath = "github.com/buildkite/shellwords",
    sum = "h1:hiVSLk7s3yFKFOHF/huoShLqrj13RMguWX2yzfvy7es=",
    version = "v0.0.0-20180315084142-c3f497d1e000",
)

go_repository(
    name = "com_github_buildkite_yaml",
    importpath = "github.com/buildkite/yaml",
    sum = "h1:q+sMKdA6L8LyGVudTkpGoC73h6ak2iWSPFiFo/pFOU8=",
    version = "v0.0.0-20181016232759-0caa5f0796e3",
)

go_repository(
    name = "com_github_datadog_datadog_go",
    importpath = "github.com/DataDog/datadog-go",
    sum = "h1:o4QtYjBU/rG58VPh8Ne6F65YiMY5/v5q4WdY/HvRYMQ=",
    version = "v3.7.2+incompatible",
)

go_repository(
    name = "com_github_fortytw2_leaktest",
    importpath = "github.com/fortytw2/leaktest",
    sum = "h1:u8491cBMTQ8ft8aeV+adlcytMZylmA5nnwwkRZjI8vw=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_go_gl_glfw",
    importpath = "github.com/go-gl/glfw",
    sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
    version = "v0.0.0-20190409004039-e6da0acd62b1",
)

go_repository(
    name = "com_github_google_go_querystring",
    importpath = "github.com/google/go-querystring",
    sum = "h1:AnCroh3fv4ZBgVIf1Iwtovgjaw/GiKJo8M8yD/fhyJ8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_google_martian_v3",
    importpath = "github.com/google/martian/v3",
    sum = "h1:d8MncMlErDFTwQGBK1xhv026j9kqhvw1Qv9IbWT1VLQ=",
    version = "v3.2.1",
)

go_repository(
    name = "com_github_mattn_go_zglob",
    importpath = "github.com/mattn/go-zglob",
    sum = "h1:xsEx/XUoVlI6yXjqBK062zYhRTZltCNmYPx6v+8DNaY=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_nightlyone_lockfile",
    importpath = "github.com/nightlyone/lockfile",
    sum = "h1:+2OJrU8cmOstEoh0uQvYemRGVH1O6xtO2oANUWHFnP0=",
    version = "v0.0.0-20180618180623-0ad87eef1443",
)

go_repository(
    name = "com_github_oleiade_reflections",
    importpath = "github.com/oleiade/reflections",
    sum = "h1:I6mXuorHlvwNDFelz7a+j0HaGYSzX7+Gq60DqLVypfc=",
    version = "v0.0.0-20160817071559-0e86b3c98b2f",
)

go_repository(
    name = "com_github_petermattis_goid",
    importpath = "github.com/petermattis/goid",
    sum = "h1:q2e307iGHPdTGp0hoxKjt1H5pDo6utceo3dQVK3I5XQ=",
    version = "v0.0.0-20180202154549-b0b1615b78e5",
)

go_repository(
    name = "com_github_philhofer_fwd",
    importpath = "github.com/philhofer/fwd",
    sum = "h1:UbZqGr5Y38ApvM/V/jEljVxwocdweyH+vmYvRPBnbqQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_qri_io_jsonpointer",
    importpath = "github.com/qri-io/jsonpointer",
    sum = "h1:C8RRfIlExwwrXw28G8LkrpWiHUVT4uLowfvjUYJ2Iec=",
    version = "v0.0.0-20180309164927-168dd9e45cf2",
)

go_repository(
    name = "com_github_qri_io_jsonschema",
    importpath = "github.com/qri-io/jsonschema",
    sum = "h1:vwTGeGWCew89DI4ZwKCaobGAN7ExvZiBzgn4LZHMVOc=",
    version = "v0.0.0-20180607150648-d0d3b10ec792",
)

go_repository(
    name = "com_github_sasha_s_go_deadlock",
    importpath = "github.com/sasha-s/go-deadlock",
    sum = "h1:T7hUw7pBSINuHQyWwMdfIWZZH5M3ju4yXIbuV/Upp+4=",
    version = "v0.0.0-20180226215254-237a9547c8a5",
)

go_repository(
    name = "com_github_tinylib_msgp",
    importpath = "github.com/tinylib/msgp",
    sum = "h1:gWmO7n0Ys2RBEb7GPYB9Ujq8Mk5p2U08lRnmMcGy6BQ=",
    version = "v1.1.2",
)

go_repository(
    name = "in_gopkg_datadog_dd_trace_go_v1",
    importpath = "gopkg.in/DataDog/dd-trace-go.v1",
    sum = "h1:EmglUJuykRsTwsQDcKaAo3CmOunWU6Dqk7U2lo7Pjss=",
    version = "v1.28.0",
)

go_repository(
    name = "io_rsc_quote_v3",
    importpath = "rsc.io/quote/v3",
    sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
    version = "v3.1.0",
)

go_repository(
    name = "io_rsc_sampler",
    importpath = "rsc.io/sampler",
    sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
    version = "v1.3.0",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:EH1Deb8WZJ0xc0WK//leUHXcX9aLE5SymusoTmMZye8=",
    version = "v0.0.0-20220411215600-e5f449aeb171",
)

go_repository(
    name = "cc_mvdan_interfacer",
    importpath = "mvdan.cc/interfacer",
    sum = "h1:WX1yoOaKQfddO/mLzdV4wptyWgoH/6hwLs7QHTixo0I=",
    version = "v0.0.0-20180901003855-c20040233aed",
)

go_repository(
    name = "cc_mvdan_lint",
    importpath = "mvdan.cc/lint",
    sum = "h1:DxJ5nJdkhDlLok9K6qO+5290kphDJbHOQO1DFFFTeBo=",
    version = "v0.0.0-20170908181259-adc824a0674b",
)

go_repository(
    name = "cc_mvdan_unparam",
    importpath = "mvdan.cc/unparam",
    sum = "h1:kAREL6MPwpsk1/PQPFD3Eg7WAQR5mPTWZJaBiG5LDbY=",
    version = "v0.0.0-20200501210554-b37ab49443f7",
)

go_repository(
    name = "com_github_akamai_akamaiopen_edgegrid_golang",
    importpath = "github.com/akamai/AkamaiOPEN-edgegrid-golang",
    sum = "h1:6rJvj+NXjjauunLeS7uGy891F1cuAwsWKa9iGzTjz1s=",
    version = "v0.9.8",
)

go_repository(
    name = "com_github_akavel_rsrc",
    importpath = "github.com/akavel/rsrc",
    sum = "h1:zjWn7ukO9Kc5Q62DOJCcxGpXC18RawVtYAGdz2aLlfw=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_alangpierce_go_forceexport",
    importpath = "github.com/alangpierce/go-forceexport",
    sum = "h1:3ILjVyslFbc4jl1w5TWuvvslFD/nDfR2H8tVaMVLrEY=",
    version = "v0.0.0-20160317203124-8f1d6941cd75",
)

go_repository(
    name = "com_github_alecthomas_assert",
    importpath = "github.com/alecthomas/assert",
    sum = "h1:smF2tmSOzy2Mm+0dGI2AIUHY+w0BUc+4tn40djz7+6U=",
    version = "v0.0.0-20170929043011-405dbfeb8e38",
)

go_repository(
    name = "com_github_alecthomas_chroma",
    importpath = "github.com/alecthomas/chroma",
    sum = "h1:7XDcGkCQopCNKjZHfYrNLraA+M7e0fMiJ/Mfikbfjek=",
    version = "v0.10.0",
)

go_repository(
    name = "com_github_alecthomas_colour",
    importpath = "github.com/alecthomas/colour",
    sum = "h1:JHZL0hZKJ1VENNfmXvHbgYlbUOvpzYzvy2aZU5gXVeo=",
    version = "v0.0.0-20160524082231-60882d9e2721",
)

go_repository(
    name = "com_github_alecthomas_kong",
    importpath = "github.com/alecthomas/kong",
    sum = "h1:Y0ZBCHAvHhTHw7FFJ2FzCAAG4pkbTgA45nc7BpMhDNk=",
    version = "v0.2.4",
)

go_repository(
    name = "com_github_alecthomas_kong_hcl",
    importpath = "github.com/alecthomas/kong-hcl",
    sum = "h1:atLL+K8Hg0e8863K2X+k7qu+xz3M2a/mWFIACAPf55M=",
    version = "v0.1.8-0.20190615233001-b21fea9723c8",
)

go_repository(
    name = "com_github_alecthomas_repr",
    importpath = "github.com/alecthomas/repr",
    sum = "h1:p9Sln00KOTlrYkxI1zYWl1QLnEqAqEARBEYa8FQnQcY=",
    version = "v0.0.0-20180818092828-117648cd9897",
)

go_repository(
    name = "com_github_aliyun_alibaba_cloud_sdk_go",
    importpath = "github.com/aliyun/alibaba-cloud-sdk-go",
    sum = "h1:E273ePcLllLIBGg5BHr3T0Fp1BJTvUyh5Y57ziSy81w=",
    version = "v1.61.112",
)

go_repository(
    name = "com_github_aliyun_aliyun_oss_go_sdk",
    importpath = "github.com/aliyun/aliyun-oss-go-sdk",
    sum = "h1:nWDRPCyCltiTsANwC/n3QZH7Vww33Npq9MKqlwRzI/c=",
    version = "v0.0.0-20190307165228-86c17b95fcd5",
)

go_repository(
    name = "com_github_andreasbriese_bbloom",
    importpath = "github.com/AndreasBriese/bbloom",
    sum = "h1:cTp8I5+VIoKjsnZuH8vjyaysT/ses3EvZeaV/1UkF2M=",
    version = "v0.0.0-20190825152654-46b345b51c96",
)

go_repository(
    name = "com_github_anmitsu_go_shlex",
    importpath = "github.com/anmitsu/go-shlex",
    sum = "h1:9AeTilPcZAjCFIImctFaOjnTIavg87rW78vTPkQqLI8=",
    version = "v0.0.0-20200514113438-38f4b401e2be",
)

go_repository(
    name = "com_github_antihax_optional",
    importpath = "github.com/antihax/optional",
    sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_antlr_antlr4",
    importpath = "github.com/antlr/antlr4",
    sum = "h1:0cEys61Sr2hUBEXfNV8eyQP01oZuBgoMeHunebPirK8=",
    version = "v0.0.0-20200503195918-621b933c7a7f",
)

go_repository(
    name = "com_github_armon_consul_api",
    importpath = "github.com/armon/consul-api",
    sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
    version = "v0.0.0-20180202201655-eb2c6b5be1b6",
)

go_repository(
    name = "com_github_azure_azure_sdk_for_go",
    importpath = "github.com/Azure/azure-sdk-for-go",
    sum = "h1:Cw16jiP4dI+CK761aq44ol4RV5dUiIIXky1+EKpoiVM=",
    version = "v58.0.0+incompatible",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_azure_auth",
    importpath = "github.com/Azure/go-autorest/autorest/azure/auth",
    sum = "h1:TzPg6B6fTZ0G1zBf3T54aI7p3cAT6u//TOXGPmFMOXg=",
    version = "v0.5.8",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_azure_cli",
    importpath = "github.com/Azure/go-autorest/autorest/azure/cli",
    sum = "h1:dMOmEJfkLKW/7JsokJqkyoYSgmR08hi9KrhjZb+JALY=",
    version = "v0.4.2",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_to",
    importpath = "github.com/Azure/go-autorest/autorest/to",
    sum = "h1:oXVqrxakqqV1UZdSazDOPOLvOIz+XA683u8EctwboHk=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_validation",
    importpath = "github.com/Azure/go-autorest/autorest/validation",
    sum = "h1:AgyqjAd94fwNAoTjl/WQXg4VvFeRFpO+UhNyRXqF1ac=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_baiyubin_aliyun_sts_go_sdk",
    importpath = "github.com/baiyubin/aliyun-sts-go-sdk",
    sum = "h1:ZNv7On9kyUzm7fvRZumSyy/IUiSC7AzL0I1jKKtwooA=",
    version = "v0.0.0-20180326062324-cfa1a18b161f",
)

go_repository(
    name = "com_github_beevik_etree",
    importpath = "github.com/beevik/etree",
    sum = "h1:T0xke/WvNtMoCqgzPhkX2r4rjY3GDZFi+FjpRZY2Jbs=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_bombsimon_wsl_v2",
    importpath = "github.com/bombsimon/wsl/v2",
    sum = "h1:/DdSteYCq4lPX+LqDg7mdoxm14UxzZPoDT0taYc3DTU=",
    version = "v2.2.0",
)

go_repository(
    name = "com_github_boombuler_barcode",
    importpath = "github.com/boombuler/barcode",
    sum = "h1:NDBbPmhS+EqABEs5Kg3n/5ZNjy73Pz7SIV+KCeqyXcs=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_bradfitz_go_smtpd",
    importpath = "github.com/bradfitz/go-smtpd",
    sum = "h1:ckJgFhFWywOx+YLEMIJsTb+NV6NexWICk5+AMSuz3ss=",
    version = "v0.0.0-20170404230938-deb6d6237625",
)

go_repository(
    name = "com_github_buger_jsonparser",
    importpath = "github.com/buger/jsonparser",
    sum = "h1:D21IyuvjDCshj1/qq+pCNd3VZOAEI9jy6Bi131YlXgI=",
    version = "v0.0.0-20181115193947-bf1c66bbce23",
)

go_repository(
    name = "com_github_cenkalti_backoff_v4",
    importpath = "github.com/cenkalti/backoff/v4",
    sum = "h1:6Yo7N8UP2K6LWZnW94DLVSSrbobcWdVzAYOisuDPIFo=",
    version = "v4.1.2",
)

go_repository(
    name = "com_github_cespare_xxhash",
    importpath = "github.com/cespare/xxhash",
    sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_cloudflare_cloudflare_go",
    importpath = "github.com/cloudflare/cloudflare-go",
    sum = "h1:VBodKICVPnwmDxstcW3biKcDSpFIfS/RELUXsZSBYK4=",
    version = "v0.10.2",
)

go_repository(
    name = "com_github_codegangsta_cli",
    importpath = "github.com/codegangsta/cli",
    sum = "h1:iX1FXEgwzd5+XN6wk5cVHOGQj6Q3Dcp20lUeS4lHNTw=",
    version = "v1.20.0",
)

go_repository(
    name = "com_github_coreos_bbolt",
    importpath = "github.com/coreos/bbolt",
    sum = "h1:n6AiVyVRKQFNb6mJlwESEvvLoDyiTzXX7ORAUlkeBdY=",
    version = "v1.3.3",
)

go_repository(
    name = "com_github_coreos_etcd",
    importpath = "github.com/coreos/etcd",
    sum = "h1:Zz1aXgDrFFi1nadh58tA9ktt06cmPTwNNP3dXwIq1lE=",
    version = "v3.3.18+incompatible",
)

go_repository(
    name = "com_github_coreos_go_etcd",
    importpath = "github.com/coreos/go-etcd",
    sum = "h1:bXhRBIXoTm9BYHS3gE0TtQuyNZyeEMux2sDi4oo5YOo=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_corpix_uarand",
    importpath = "github.com/corpix/uarand",
    sum = "h1:RMr1TWc9F4n5jiPDzFHtmaUXLKLNUFK0SgCLo4BhX/U=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_cpu_goacmedns",
    importpath = "github.com/cpu/goacmedns",
    sum = "h1:hYAgjnPu7HogTgb8trqQouR/RrBgXq1TPBgmxbK9eRA=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man",
    importpath = "github.com/cpuguy83/go-md2man",
    sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
    version = "v1.0.10",
)

go_repository(
    name = "com_github_crewjam_httperr",
    importpath = "github.com/crewjam/httperr",
    sum = "h1:WXnT88cFG2davqSFqvaFfzkSMC0lqh/8/rKZ+z7tYvI=",
    version = "v0.0.0-20190612203328-a946449404da",
)

go_repository(
    name = "com_github_crewjam_saml",
    importpath = "github.com/crewjam/saml",
    sum = "h1:H9u+6CZAESUKHxMyxUbVn0IawYvKZn4nt3d4ccV4O/M=",
    version = "v0.4.5",
)

go_repository(
    name = "com_github_daaku_go_zipexe",
    importpath = "github.com/daaku/go.zipexe",
    sum = "h1:wV4zMsDOI2SZ2m7Tdz1Ps96Zrx+TzaK15VbUaGozw0M=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_danwakefield_fnmatch",
    importpath = "github.com/danwakefield/fnmatch",
    sum = "h1:y5HC9v93H5EPKqaS1UYVg1uYah5Xf51mBfIoWehClUQ=",
    version = "v0.0.0-20160403171240-cbb64ac3d964",
)

go_repository(
    name = "com_github_datadog_zstd",
    importpath = "github.com/DataDog/zstd",
    sum = "h1:EndNeuB0l9syBZhut0wns3gV1hL8zX8LIu6ZiVHWLIQ=",
    version = "v1.4.5",
)

go_repository(
    name = "com_github_dchest_uniuri",
    importpath = "github.com/dchest/uniuri",
    sum = "h1:74lLNRzvsdIlkTgfDSMuaPjBr4cf6k7pwQQANm/yLKU=",
    version = "v0.0.0-20160212164326-8902c56451e9",
)

go_repository(
    name = "com_github_dgraph_io_badger",
    importpath = "github.com/dgraph-io/badger",
    sum = "h1:mNw0qs90GVgGGWylh0umH5iag1j6n/PeJtNvL6KY/x8=",
    version = "v1.6.2",
)

go_repository(
    name = "com_github_dgraph_io_badger_v2",
    importpath = "github.com/dgraph-io/badger/v2",
    sum = "h1:TRWBQg8UrlUhaFdco01nO2uXwzKS7zd+HVdwV/GHc4o=",
    version = "v2.2007.4",
)

go_repository(
    name = "com_github_dgraph_io_ristretto",
    importpath = "github.com/dgraph-io/ristretto",
    sum = "h1:KoJOtZf+6wpQaDTuOWGuo61GxcPBIfhwRxRTaTWGCTc=",
    version = "v0.0.4-0.20200906165740-41ebdbffecfd",
)

go_repository(
    name = "com_github_dgryski_go_farm",
    importpath = "github.com/dgryski/go-farm",
    sum = "h1:fAjc9m62+UWV/WAFKLNi6ZS0675eEUC9y3AlwSbQu1Y=",
    version = "v0.0.0-20200201041132-a6ae2369ad13",
)

go_repository(
    name = "com_github_dgryski_go_sip13",
    importpath = "github.com/dgryski/go-sip13",
    sum = "h1:RMLoZVzv4GliuWafOuPuQDKSm1SJph7uCRnnS61JAn4=",
    version = "v0.0.0-20181026042036-e10d5fee7954",
)

go_repository(
    name = "com_github_dimchansky_utfbom",
    importpath = "github.com/dimchansky/utfbom",
    sum = "h1:vV6w1AhK4VMnhBno/TPVCoK9U/LP0PkLCS9tbxHdi/U=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_dlclark_regexp2",
    importpath = "github.com/dlclark/regexp2",
    sum = "h1:F1rxgk7p4uKjwIQxBs9oAXe5CqrXlCduYEJvrF4u93E=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_dnaeon_go_vcr",
    importpath = "github.com/dnaeon/go-vcr",
    sum = "h1:r8L/HqC0Hje5AXMu1ooW8oyQyOFv4GxqpL0nRP7SLLY=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_dnsimple_dnsimple_go",
    importpath = "github.com/dnsimple/dnsimple-go",
    sum = "h1:N+q+ML1CZGf+5r4udu9Opy7WJNtOaFT9aM86Af9gLhk=",
    version = "v0.60.0",
)

go_repository(
    name = "com_github_exoscale_egoscale",
    importpath = "github.com/exoscale/egoscale",
    sum = "h1:1FNZVk8jHUx0AvWhOZxLEDNlacTU0chMXUUNkm9EZaI=",
    version = "v0.18.1",
)

go_repository(
    name = "com_github_fatih_structs",
    importpath = "github.com/fatih/structs",
    sum = "h1:Q7juDM0QtcnhCpeyLGQKyg4TOIghuNXrkL32pHAUMxo=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_flynn_go_shlex",
    importpath = "github.com/flynn/go-shlex",
    sum = "h1:BHsljHzVlRcyQhjrss6TZTdY2VfCqZPbv5k3iBFa2ZQ=",
    version = "v0.0.0-20150515145356-3f9db97f8568",
)

go_repository(
    name = "com_github_francoispqt_gojay",
    importpath = "github.com/francoispqt/gojay",
    sum = "h1:d2m3sFjloqoIUQU3TsHBgj6qg/BVGlTBeHDUmyJnXKk=",
    version = "v1.2.13",
)

go_repository(
    name = "com_github_geertjohan_go_incremental",
    importpath = "github.com/GeertJohan/go.incremental",
    sum = "h1:7AH+pY1XUgQE4Y1HcXYaMqAI0m9yrFqo/jt0CW30vsg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_geertjohan_go_rice",
    importpath = "github.com/GeertJohan/go.rice",
    sum = "h1:KkI6O9uMaQU3VEKaj01ulavtF7o1fWT7+pk/4voiMLQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_gliderlabs_ssh",
    importpath = "github.com/gliderlabs/ssh",
    sum = "h1:6zsha5zo/TWhRhwqCD3+EarCAgZ2yN28ipRnGPnwkI0=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_go_acme_lego_v3",
    importpath = "github.com/go-acme/lego/v3",
    sum = "h1:qC5/8/CbltyAE8fGLE6bGlqucj7pXc/vBxiLwLOsmAQ=",
    version = "v3.7.0",
)

go_repository(
    name = "com_github_go_asn1_ber_asn1_ber",
    importpath = "github.com/go-asn1-ber/asn1-ber",
    sum = "h1:pDbRAunXzIUXfx4CB2QJFv5IuPiuoW+sWvr/Us009o8=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_go_chi_chi",
    importpath = "github.com/go-chi/chi",
    sum = "h1:fGFk2Gmi/YKXk0OmGfBh0WgmN3XB8lVnEyNz34tQRec=",
    version = "v4.1.2+incompatible",
)

go_repository(
    name = "com_github_go_cmd_cmd",
    importpath = "github.com/go-cmd/cmd",
    sum = "h1:IK23uTRWxq6UJnNWp8nKO7mVCwnPfbaxA2lhzEKfNj0=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_go_critic_go_critic",
    importpath = "github.com/go-critic/go-critic",
    sum = "h1:sGEEdiuvLV0OC7/yC6MnK3K6LCPBplspK45B0XVdFAc=",
    version = "v0.4.3",
)

go_repository(
    name = "com_github_go_errors_errors",
    importpath = "github.com/go-errors/errors",
    sum = "h1:ljK/pL5ltg3qoN+OtN6yCv9HWSfMwxSx90GJCZQxYNg=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_go_ini_ini",
    importpath = "github.com/go-ini/ini",
    sum = "h1:8+SRbfpRFlIunpSum4BEf1ClTtVjOgKzgBv9pHFkI6w=",
    version = "v1.44.0",
)

go_repository(
    name = "com_github_go_ldap_ldap",
    importpath = "github.com/go-ldap/ldap",
    sum = "h1:HTeSZO8hWMS1Rgb2Ziku6b8a7qRIZZMHjsvuZyatzwk=",
    version = "v3.0.3+incompatible",
)

go_repository(
    name = "com_github_go_ldap_ldap_v3",
    importpath = "github.com/go-ldap/ldap/v3",
    sum = "h1:fU/0xli6HY02ocbMuozHAYsaHLcnkLjvho2r5a34BUU=",
    version = "v3.4.1",
)

go_repository(
    name = "com_github_go_lintpack_lintpack",
    importpath = "github.com/go-lintpack/lintpack",
    sum = "h1:DI5mA3+eKdWeJ40nU4d6Wc26qmdG8RCi/btYq0TuRN0=",
    version = "v0.5.2",
)

go_repository(
    name = "com_github_go_ole_go_ole",
    importpath = "github.com/go-ole/go-ole",
    sum = "h1:/Fpf6oFPoeFik9ty7siob0G6Ke8QvQEuVcuChpwXzpY=",
    version = "v1.2.6",
)

go_repository(
    name = "com_github_go_piv_piv_go",
    importpath = "github.com/go-piv/piv-go",
    sum = "h1:rfjdFdASfGV5KLJhSjgpGJ5lzVZVtRWn8ovy/H9HQ/U=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_go_toolsmith_astcast",
    importpath = "github.com/go-toolsmith/astcast",
    sum = "h1:JojxlmI6STnFVG9yOImLeGREv8W2ocNUM+iOhR6jE7g=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_astcopy",
    importpath = "github.com/go-toolsmith/astcopy",
    sum = "h1:OMgl1b1MEpjFQ1m5ztEO06rz5CUd3oBv9RF7+DyvdG8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_astequal",
    importpath = "github.com/go-toolsmith/astequal",
    sum = "h1:4zxD8j3JRFNyLN46lodQuqz3xdKSrur7U/sr0SDS/gQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_astfmt",
    importpath = "github.com/go-toolsmith/astfmt",
    sum = "h1:A0vDDXt+vsvLEdbMFJAUBI/uTbRw1ffOPnxsILnFL6k=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_astinfo",
    importpath = "github.com/go-toolsmith/astinfo",
    sum = "h1:wP6mXeB2V/d1P1K7bZ5vDUO3YqEzcvOREOxZPEu3gVI=",
    version = "v0.0.0-20180906194353-9809ff7efb21",
)

go_repository(
    name = "com_github_go_toolsmith_astp",
    importpath = "github.com/go-toolsmith/astp",
    sum = "h1:alXE75TXgcmupDsMK1fRAy0YUzLzqPVvBKoyWV+KPXg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_pkgload",
    importpath = "github.com/go-toolsmith/pkgload",
    sum = "h1:4DFWWMXVfbcN5So1sBNW9+yeiMqLFGl1wFLTL5R0Tgg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_strparse",
    importpath = "github.com/go-toolsmith/strparse",
    sum = "h1:Vcw78DnpCAKlM20kSbAyO4mPfJn/lyYA4BJUDxe2Jb4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_toolsmith_typep",
    importpath = "github.com/go-toolsmith/typep",
    sum = "h1:8xdsa1+FSIH/RhEkgnD1j2CJOy5mNllW1Q9tRiYwvlk=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_gofrs_flock",
    importpath = "github.com/gofrs/flock",
    sum = "h1:ekuhfTjngPhisSjOJ0QWKpPQE8/rbknHaes6WVJj5Hw=",
    version = "v0.0.0-20190320160742-5135e617513b",
)

go_repository(
    name = "com_github_gofrs_uuid",
    importpath = "github.com/gofrs/uuid",
    sum = "h1:yyYWMnhkhrKwwr8gAOcOCYxOOscHgDS9yZgBrnJfGa0=",
    version = "v4.2.0+incompatible",
)

go_repository(
    name = "com_github_goji_httpauth",
    importpath = "github.com/goji/httpauth",
    sum = "h1:lBXNCxVENCipq4D1Is42JVOP4eQjlB8TQ6H69Yx5J9Q=",
    version = "v0.0.0-20160601135302-2da839ab0f4d",
)

go_repository(
    name = "com_github_golang_lint",
    importpath = "github.com/golang/lint",
    sum = "h1:2hRPrmiwPrp3fQX967rNJIhQPtiGXdlQWAxKbKw3VHA=",
    version = "v0.0.0-20180702182130-06c8688daad7",
)

go_repository(
    name = "com_github_golangci_check",
    importpath = "github.com/golangci/check",
    sum = "h1:23T5iq8rbUYlhpt5DB4XJkc6BU31uODLD1o1gKvZmD0=",
    version = "v0.0.0-20180506172741-cfe4005ccda2",
)

go_repository(
    name = "com_github_golangci_dupl",
    importpath = "github.com/golangci/dupl",
    sum = "h1:w8hkcTqaFpzKqonE9uMCefW1WDie15eSP/4MssdenaM=",
    version = "v0.0.0-20180902072040-3e9179ac440a",
)

go_repository(
    name = "com_github_golangci_errcheck",
    importpath = "github.com/golangci/errcheck",
    sum = "h1:YYWNAGTKWhKpcLLt7aSj/odlKrSrelQwlovBpDuf19w=",
    version = "v0.0.0-20181223084120-ef45e06d44b6",
)

go_repository(
    name = "com_github_golangci_go_misc",
    importpath = "github.com/golangci/go-misc",
    sum = "h1:9kfjN3AdxcbsZBf8NjltjWihK2QfBBBZuv91cMFfDHw=",
    version = "v0.0.0-20180628070357-927a3d87b613",
)

go_repository(
    name = "com_github_golangci_go_tools",
    importpath = "github.com/golangci/go-tools",
    sum = "h1:/7detzz5stiXWPzkTlPTzkBEIIE4WGpppBJYjKqBiPI=",
    version = "v0.0.0-20190318055746-e32c54105b7c",
)

go_repository(
    name = "com_github_golangci_goconst",
    importpath = "github.com/golangci/goconst",
    sum = "h1:pe9JHs3cHHDQgOFXJJdYkK6fLz2PWyYtP4hthoCMvs8=",
    version = "v0.0.0-20180610141641-041c5f2b40f3",
)

go_repository(
    name = "com_github_golangci_gocyclo",
    importpath = "github.com/golangci/gocyclo",
    sum = "h1:pXTK/gkVNs7Zyy7WKgLXmpQ5bHTrq5GDsp8R9Qs67g0=",
    version = "v0.0.0-20180528144436-0a533e8fa43d",
)

go_repository(
    name = "com_github_golangci_gofmt",
    importpath = "github.com/golangci/gofmt",
    sum = "h1:iR3fYXUjHCR97qWS8ch1y9zPNsgXThGwjKPrYfqMPks=",
    version = "v0.0.0-20190930125516-244bba706f1a",
)

go_repository(
    name = "com_github_golangci_golangci_lint",
    importpath = "github.com/golangci/golangci-lint",
    sum = "h1:VYLx63qb+XJsHdZ27PMS2w5JZacN0XG8ffUwe7yQomo=",
    version = "v1.27.0",
)

go_repository(
    name = "com_github_golangci_gosec",
    importpath = "github.com/golangci/gosec",
    sum = "h1:fUdgm/BdKvwOHxg5AhNbkNRp2mSy8sxTXyBVs/laQHo=",
    version = "v0.0.0-20190211064107-66fb7fc33547",
)

go_repository(
    name = "com_github_golangci_ineffassign",
    importpath = "github.com/golangci/ineffassign",
    sum = "h1:gLLhTLMk2/SutryVJ6D4VZCU3CUqr8YloG7FPIBWFpI=",
    version = "v0.0.0-20190609212857-42439a7714cc",
)

go_repository(
    name = "com_github_golangci_lint_1",
    importpath = "github.com/golangci/lint-1",
    sum = "h1:MfyDlzVjl1hoaPzPD4Gpb/QgoRfSBR0jdhwGyAWwMSA=",
    version = "v0.0.0-20191013205115-297bf364a8e0",
)

go_repository(
    name = "com_github_golangci_maligned",
    importpath = "github.com/golangci/maligned",
    sum = "h1:kNY3/svz5T29MYHubXix4aDDuE3RWHkPvopM/EDv/MA=",
    version = "v0.0.0-20180506175553-b1d89398deca",
)

go_repository(
    name = "com_github_golangci_misspell",
    importpath = "github.com/golangci/misspell",
    sum = "h1:pLzmVdl3VxTOncgzHcvLOKirdvcx/TydsClUQXTehjo=",
    version = "v0.3.5",
)

go_repository(
    name = "com_github_golangci_prealloc",
    importpath = "github.com/golangci/prealloc",
    sum = "h1:leSNB7iYzLYSSx3J/s5sVf4Drkc68W2wm4Ixh/mr0us=",
    version = "v0.0.0-20180630174525-215b22d4de21",
)

go_repository(
    name = "com_github_golangci_revgrep",
    importpath = "github.com/golangci/revgrep",
    sum = "h1:XQKc8IYQOeRwVs36tDrEmTgDgP88d5iEURwpmtiAlOM=",
    version = "v0.0.0-20180812185044-276a5c0a1039",
)

go_repository(
    name = "com_github_golangci_unconvert",
    importpath = "github.com/golangci/unconvert",
    sum = "h1:zwtduBRr5SSWhqsYNgcuWO2kFlpdOZbP0+yRjmvPGys=",
    version = "v0.0.0-20180507085042-28b1c447d1f4",
)

go_repository(
    name = "com_github_google_cel_go",
    importpath = "github.com/google/cel-go",
    sum = "h1:8v9BSN0avuGwrHFKNCjfiQ/CE6+D6sW+BDyOVoEeP6o=",
    version = "v0.7.3",
)

go_repository(
    name = "com_github_google_cel_spec",
    importpath = "github.com/google/cel-spec",
    sum = "h1:hWEzw+1L1UNxfHAbKXYbirsPGlG8ArXNcTnBKvBqRJ0=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_google_certificate_transparency_go",
    importpath = "github.com/google/certificate-transparency-go",
    sum = "h1:BwvtrCU6nQ3cw0bAbwClSGFQvZpY2S5yt8amttqFSnQ=",
    version = "v1.1.2-0.20210623111010-a50f74f4ce95",
)

go_repository(
    name = "com_github_google_go_github",
    importpath = "github.com/google/go-github",
    sum = "h1:N0LgJ1j65A7kfXrZnUDaYCs/Sf4rEjNlfyDHW9dolSY=",
    version = "v17.0.0+incompatible",
)

go_repository(
    name = "com_github_google_monologue",
    importpath = "github.com/google/monologue",
    sum = "h1:0L/piDwninh6sjZ+vCZI7c6RA0R71ET8v1yinZzC9k8=",
    version = "v0.0.0-20191220140058-35abc9683a6c",
)

go_repository(
    name = "com_github_google_trillian",
    importpath = "github.com/google/trillian",
    sum = "h1:hXrMoOL/sEOsFF2Snw1pmE1fn8FH9qktivGvjf4y/wU=",
    version = "v1.3.14-0.20210622121126-870e0cdde059",
)

go_repository(
    name = "com_github_google_trillian_examples",
    importpath = "github.com/google/trillian-examples",
    sum = "h1:dv2J28D109qglM6VfNzAXZ7VddBojviT5oMSs1yeDUY=",
    version = "v0.0.0-20190603134952-4e75ba15216c",
)

go_repository(
    name = "com_github_googleapis_gax_go",
    importpath = "github.com/googleapis/gax-go",
    sum = "h1:silFMLAnr330+NRuag/VjIGF7TLp/LBrV2CJKFLWEww=",
    version = "v2.0.2+incompatible",
)

go_repository(
    name = "com_github_gorilla_csrf",
    importpath = "github.com/gorilla/csrf",
    sum = "h1:60oN1cFdncCE8tjwQ3QEkFND5k37lQPcRjnlvm7CIJ0=",
    version = "v1.6.0",
)

go_repository(
    name = "com_github_gorilla_securecookie",
    importpath = "github.com/gorilla/securecookie",
    sum = "h1:miw7JPhV+b/lAHSXz4qd/nN9jRiAFV5FwjeKyCS8BvQ=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_gostaticanalysis_analysisutil",
    importpath = "github.com/gostaticanalysis/analysisutil",
    sum = "h1:iwp+5/UAyzQSFgQ4uR2sni99sJ8Eo9DEacKWM5pekIg=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_gravitational_trace",
    importpath = "github.com/gravitational/trace",
    sum = "h1:68WxnfBzJRYktZ30fmIjGQ74RsXYLoeH2/NITPktTMY=",
    version = "v0.0.0-20190726142706-a535a178675f",
)

go_repository(
    name = "com_github_h2non_parth",
    importpath = "github.com/h2non/parth",
    sum = "h1:2VTzZjLZBgl62/EtslCrtky5vbi9dd7HrQPQIx6wqiw=",
    version = "v0.0.0-20190131123155-b4df798d6542",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    importpath = "github.com/hashicorp/hcl",
    sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_huandu_xstrings",
    importpath = "github.com/huandu/xstrings",
    sum = "h1:L18LIDzqlW6xN2rEkpdV8+oL/IXWJ1APd+vsdYy4Wdw=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_iancoleman_strcase",
    importpath = "github.com/iancoleman/strcase",
    sum = "h1:dJBk1m2/qjL1twPLf68JND55vvivMupZ4wIzE8CTdBw=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_icrowley_fake",
    importpath = "github.com/icrowley/fake",
    sum = "h1:Mo9W14pwbO9VfRe+ygqZ8dFbPpoIK1HFrG/zjTuQ+nc=",
    version = "v0.0.0-20180203215853-4178557ae428",
)

go_repository(
    name = "com_github_iij_doapi",
    importpath = "github.com/iij/doapi",
    sum = "h1:MZf03xP9WdakyXhOWuAD5uPK3wHh96wCsqe3hCMKh8E=",
    version = "v0.0.0-20190504054126-0bbf12d6d7df",
)

go_repository(
    name = "com_github_jellevandenhooff_dkim",
    importpath = "github.com/jellevandenhooff/dkim",
    sum = "h1:ujPKutqRlJtcfWk6toYVYagwra7HQHbXOaS171b4Tg8=",
    version = "v0.0.0-20150330215556-f50fe3d243e1",
)

go_repository(
    name = "com_github_jessevdk_go_flags",
    importpath = "github.com/jessevdk/go-flags",
    sum = "h1:4IU2WS7AumrZ/40jfhf4QVDMsQwqA7VEHozFRrGARJA=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_jsternberg_zap_logfmt",
    importpath = "github.com/jsternberg/zap-logfmt",
    sum = "h1:1v+PK4/B48cy8cfQbxL4FmmNZrjnIMr2BsnyEmXqv2o=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_juju_ansiterm",
    importpath = "github.com/juju/ansiterm",
    sum = "h1:FaWFmfWdAUKbSCtOU2QjDaorUexogfaMgbipgYATUMU=",
    version = "v0.0.0-20180109212912-720a0952cc2a",
)

go_repository(
    name = "com_github_juju_ratelimit",
    importpath = "github.com/juju/ratelimit",
    sum = "h1:+7AIFJVQ0EQgq/K9+0Krm7m530Du7tIz0METWzN0RgY=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_kballard_go_shellquote",
    importpath = "github.com/kballard/go-shellquote",
    sum = "h1:Z9n2FFNUXsshfwJMBgNA0RU6/i7WVaAegv3PtuIHPMs=",
    version = "v0.0.0-20180428030007-95032a82bc51",
)

go_repository(
    name = "com_github_klauspost_cpuid_v2",
    importpath = "github.com/klauspost/cpuid/v2",
    sum = "h1:i2lw1Pm7Yi/4O6XCSyJWqEHI2MDw2FzUK6o/D21xn2A=",
    version = "v2.0.11",
)

go_repository(
    name = "com_github_kolo_xmlrpc",
    importpath = "github.com/kolo/xmlrpc",
    sum = "h1:TrxPzApUukas24OMMVDUMlCs1XCExJtnGaDEiIAR4oQ=",
    version = "v0.0.0-20190717152603-07c4ee3fd181",
)

go_repository(
    name = "com_github_kylelemons_godebug",
    importpath = "github.com/kylelemons/godebug",
    sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_labbsr0x_bindman_dns_webhook",
    importpath = "github.com/labbsr0x/bindman-dns-webhook",
    sum = "h1:I7ITbmQPAVwrDdhd6dHKi+MYJTJqPCK0jE6YNBAevnk=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_labbsr0x_goh",
    importpath = "github.com/labbsr0x/goh",
    sum = "h1:97aBJkDjpyBZGPbQuOK5/gHcSFbcr5aRsq3RSRJFpPk=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_letsencrypt_pkcs11key",
    importpath = "github.com/letsencrypt/pkcs11key",
    sum = "h1:GfzE+uq7odDW7nOmp1QWuilLEK7kJf8i84XcIfk3mKA=",
    version = "v2.0.1-0.20170608213348-396559074696+incompatible",
)

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    sum = "h1:AqzbZs4ZoCBp+GtejcpCpcxM3zlSMx29dXbUSeVtJb8=",
    version = "v1.10.2",
)

go_repository(
    name = "com_github_libdns_libdns",
    importpath = "github.com/libdns/libdns",
    sum = "h1:Wu59T7wSHRgtA0cfxC+n1c/e+O3upJGWytknkmFEDis=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_linode_linodego",
    importpath = "github.com/linode/linodego",
    sum = "h1:AMdb82HVgY8o3mjBXJcUv9B+fnJjfDMn2rNRGbX+jvM=",
    version = "v0.10.0",
)

go_repository(
    name = "com_github_liquidweb_liquidweb_go",
    importpath = "github.com/liquidweb/liquidweb-go",
    sum = "h1:vIj1I/Wf97fUnyirD+bi6Y63c0GiXk9nKI1+sFFl3G0=",
    version = "v1.6.0",
)

go_repository(
    name = "com_github_logrusorgru_aurora",
    importpath = "github.com/logrusorgru/aurora",
    sum = "h1:9MlwzLdW7QSDrhDjFlsEYmxpFyIoXmYRon3dt0io31k=",
    version = "v0.0.0-20181002194514-a7b3b318ed4e",
)

go_repository(
    name = "com_github_lunixbochs_vtclean",
    importpath = "github.com/lunixbochs/vtclean",
    sum = "h1:xu2sLAri4lGiovBDQKxl5mrXyESr3gUr5m5SM5+LVb8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_magiconair_properties",
    importpath = "github.com/magiconair/properties",
    sum = "h1:b6kJs+EmPFMYGkow9GiUyCyOvIwYetYJ3fSaWak/Gls=",
    version = "v1.8.5",
)

go_repository(
    name = "com_github_mailgun_minheap",
    importpath = "github.com/mailgun/minheap",
    sum = "h1:aOqSQstfwSx9+tcM/xiKTio3IVjs7ZL2vU8kI9bI6bM=",
    version = "v0.0.0-20170619185613-3dbe6c6bf55f",
)

go_repository(
    name = "com_github_mailgun_multibuf",
    importpath = "github.com/mailgun/multibuf",
    sum = "h1:m2FGM8K2LC9Zyt/7zbQNn5Uvf/YV7vFWKtoMcC7hHU8=",
    version = "v0.0.0-20150714184110-565402cd71fb",
)

go_repository(
    name = "com_github_mailgun_timetools",
    importpath = "github.com/mailgun/timetools",
    sum = "h1:Kg/NPZLLC3aAFr1YToMs98dbCdhootQ1hZIvZU28hAQ=",
    version = "v0.0.0-20141028012446-7e6055773c51",
)

go_repository(
    name = "com_github_mailgun_ttlmap",
    importpath = "github.com/mailgun/ttlmap",
    sum = "h1:ZZYhg16XocqSKPGNQAe0aeweNtFxuedbwwb4fSlg7h4=",
    version = "v0.0.0-20170619185759-c1c17f74874f",
)

go_repository(
    name = "com_github_manifoldco_promptui",
    importpath = "github.com/manifoldco/promptui",
    sum = "h1:3V4HzJk1TtXW1MTZMP7mdlwbBpIinw3HztaIlYthEiA=",
    version = "v0.9.0",
)

go_repository(
    name = "com_github_marten_seemann_qpack",
    importpath = "github.com/marten-seemann/qpack",
    sum = "h1:jvTsT/HpCn2UZJdP+UUB53FfUUgeOyG5K1ns0OJOGVs=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_marten_seemann_qtls",
    importpath = "github.com/marten-seemann/qtls",
    sum = "h1:ECsuYUKalRL240rRD4Ri33ISb7kAQ3qGDlrrl55b2pc=",
    version = "v0.10.0",
)

go_repository(
    name = "com_github_marten_seemann_qtls_go1_15",
    importpath = "github.com/marten-seemann/qtls-go1-15",
    sum = "h1:Ci4EIUN6Rlb+D6GmLdej/bCQ4nPYNtVXQB+xjiXE1nk=",
    version = "v0.1.5",
)

go_repository(
    name = "com_github_masterminds_glide",
    importpath = "github.com/Masterminds/glide",
    sum = "h1:M5MOH04TyRiMBVeWHbifqTpnauxWINIubTCOkhXh+2g=",
    version = "v0.13.2",
)

go_repository(
    name = "com_github_masterminds_goutils",
    importpath = "github.com/Masterminds/goutils",
    sum = "h1:5nUrii3FMTL5diU80unEVvNevw1nH4+ZV4DSLVJLSYI=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_masterminds_semver",
    importpath = "github.com/Masterminds/semver",
    sum = "h1:H65muMkzWKEuNDnfl9d70GUjFniHKHRbFPGBuZ3QEww=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_masterminds_semver_v3",
    importpath = "github.com/Masterminds/semver/v3",
    sum = "h1:hLg3sBzpNErnxhQtUy/mmLR2I9foDujNK030IGemrRc=",
    version = "v3.1.1",
)

go_repository(
    name = "com_github_masterminds_sprig_v3",
    importpath = "github.com/Masterminds/sprig/v3",
    sum = "h1:17jRggJu518dr3QaafizSXOjKYp94wKfABxUmyxvxX8=",
    version = "v3.2.2",
)

go_repository(
    name = "com_github_masterminds_vcs",
    importpath = "github.com/Masterminds/vcs",
    sum = "h1:USF5TvZGYgIpcbNAEMLfFhHqP08tFZVlUVrmTSpqnyA=",
    version = "v1.13.0",
)

go_repository(
    name = "com_github_matoous_godox",
    importpath = "github.com/matoous/godox",
    sum = "h1:RHba4YImhrUVQDHUCe2BNSOz4tVy2yGyXhvYDvxGgeE=",
    version = "v0.0.0-20190911065817-5d6d842e92eb",
)

go_repository(
    name = "com_github_mattermost_xml_roundtrip_validator",
    importpath = "github.com/mattermost/xml-roundtrip-validator",
    sum = "h1:qqXczln0qwkVGcpQ+sQuPOVntt2FytYarXXxYSNJkgw=",
    version = "v0.0.0-20201213122252-bcd7e1b9601e",
)

go_repository(
    name = "com_github_mattn_go_tty",
    importpath = "github.com/mattn/go-tty",
    sum = "h1:8TGB3DFRNl06DB1Q6zBX+I7FDoCUZY2fmMS9WGUIIpw=",
    version = "v0.0.0-20180219170247-931426f7535a",
)

go_repository(
    name = "com_github_mattn_goveralls",
    importpath = "github.com/mattn/goveralls",
    sum = "h1:7eJB6EqsPhRVxvwEXGnqdO2sJI0PTsrWoTMXEk9/OQc=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_mholt_acmez",
    importpath = "github.com/mholt/acmez",
    sum = "h1:C8wsEBIUVi6e0DYoxqCcFuXtwc4AWXL/jgcDjF7mjVo=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_microcosm_cc_bluemonday",
    importpath = "github.com/microcosm-cc/bluemonday",
    sum = "h1:SIYunPjnlXcW+gVfvm0IlSeR5U3WZUOLfVmqg85Go44=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_microsoft_go_winio",
    importpath = "github.com/Microsoft/go-winio",
    sum = "h1:iT12IBVClFevaf8PuVyi3UmZOVh4OqnaLxDTW2O6j3w=",
    version = "v0.4.17",
)

go_repository(
    name = "com_github_miekg_pkcs11",
    importpath = "github.com/miekg/pkcs11",
    sum = "h1:iMwmD7I5225wv84WxIG/bmxz9AXjWvTWIbM/TYHvWtw=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_mitchellh_copystructure",
    importpath = "github.com/mitchellh/copystructure",
    sum = "h1:vpKXTN4ewci03Vljg/q9QvCGUDttBOGBIa15WveJJGw=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_mitchellh_go_ps",
    importpath = "github.com/mitchellh/go-ps",
    sum = "h1:i6ampVEEF4wQFF+bkYfwYgY+F/uYJDktmvLPf7qIgjc=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_go_vnc",
    importpath = "github.com/mitchellh/go-vnc",
    sum = "h1:FI2NIv6fpef6BQl2u3IZX/Cj20tfypRF4yd+uaHOMtI=",
    version = "v0.0.0-20150629162542-723ed9867aed",
)

go_repository(
    name = "com_github_mitchellh_reflectwalk",
    importpath = "github.com/mitchellh/reflectwalk",
    sum = "h1:G2LzWKi524PWgd3mLHV8Y5k7s6XUvT0Gef6zxSIeXaQ=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_mohae_deepcopy",
    importpath = "github.com/mohae/deepcopy",
    sum = "h1:RWengNIwukTxcDr9M+97sNutRR1RKhG96O6jWumTTnw=",
    version = "v0.0.0-20170929034955-c48cc78d4826",
)

go_repository(
    name = "com_github_mozilla_tls_observatory",
    importpath = "github.com/mozilla/tls-observatory",
    sum = "h1:1xJ+Xi9lYWLaaP4yB67ah0+548CD3110mCPWhVVjFkI=",
    version = "v0.0.0-20200317151703-4fa42e1c2dee",
)

go_repository(
    name = "com_github_mreiferson_go_httpclient",
    importpath = "github.com/mreiferson/go-httpclient",
    sum = "h1:oKIteTqeSpenyTrOVj5zkiyCaflLa8B+CD0324otT+o=",
    version = "v0.0.0-20160630210159-31f0106b4474",
)

go_repository(
    name = "com_github_namedotcom_go",
    importpath = "github.com/namedotcom/go",
    sum = "h1:o6uBwrhM5C8Ll3MAAxrQxRHEu7FkapwTuI2WmL1rw4g=",
    version = "v0.0.0-20180403034216-08470befbe04",
)

go_repository(
    name = "com_github_naoina_go_stringutil",
    importpath = "github.com/naoina/go-stringutil",
    sum = "h1:rCUeRUHjBjGTSHl0VC00jUPLz8/F9dDzYI70Hzifhks=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_naoina_toml",
    importpath = "github.com/naoina/toml",
    sum = "h1:PT/lllxVVN0gzzSqSlHEmP8MJB4MY2U7STGxiouV4X8=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_nbio_st",
    importpath = "github.com/nbio/st",
    sum = "h1:W6apQkHrMkS0Muv8G/TipAy/FJl/rCYT0+EuS8+Z0z4=",
    version = "v0.0.0-20140626010706-e9e8d9816f32",
)

go_repository(
    name = "com_github_nbutton23_zxcvbn_go",
    importpath = "github.com/nbutton23/zxcvbn-go",
    sum = "h1:AREM5mwr4u1ORQBMvzfzBgpsctsbQikCVpvC+tX285E=",
    version = "v0.0.0-20180912185939-ae427f1e4c1d",
)

go_repository(
    name = "com_github_neelance_astrewrite",
    importpath = "github.com/neelance/astrewrite",
    sum = "h1:D6paGObi5Wud7xg83MaEFyjxQB1W5bz5d0IFppr+ymk=",
    version = "v0.0.0-20160511093645-99348263ae86",
)

go_repository(
    name = "com_github_neelance_sourcemap",
    importpath = "github.com/neelance/sourcemap",
    sum = "h1:bY6ktFuJkt+ZXkX0RChQch2FtHpWQLVS8Qo1YasiIVk=",
    version = "v0.0.0-20200213170602-2833bce08e4c",
)

go_repository(
    name = "com_github_newrelic_go_agent",
    importpath = "github.com/newrelic/go-agent",
    sum = "h1:IB0Fy+dClpBq9aEoIrLyQXzU34JyI1xVTanPLB/+jvU=",
    version = "v2.15.0+incompatible",
)

go_repository(
    name = "com_github_ngdinhtoan_glide_cleanup",
    importpath = "github.com/ngdinhtoan/glide-cleanup",
    sum = "h1:kN4sV+0tp2F1BvwU+5SfNRMDndRmvIfnI3kZ7B8Yv4Y=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_niemeyer_pretty",
    importpath = "github.com/niemeyer/pretty",
    sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
    version = "v0.0.0-20200227124842-a10e7caefd8e",
)

go_repository(
    name = "com_github_nkovacs_streamquote",
    importpath = "github.com/nkovacs/streamquote",
    sum = "h1:E2B8qYyeSgv5MXpmzZXRNp8IAQ4vjxIjhpAf5hv/tAg=",
    version = "v0.0.0-20170412213628-49af9bddb229",
)

go_repository(
    name = "com_github_nrdcg_auroradns",
    importpath = "github.com/nrdcg/auroradns",
    sum = "h1:m/kBq83Xvy3cU261MOknd8BdnOk12q4lAWM+kOdsC2Y=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_nrdcg_dnspod_go",
    importpath = "github.com/nrdcg/dnspod-go",
    sum = "h1:c/jn1mLZNKF3/osJ6mz3QPxTudvPArXTjpkmYj0uK6U=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_nrdcg_goinwx",
    importpath = "github.com/nrdcg/goinwx",
    sum = "h1:AJnjoWPELyCtofhGcmzzcEMFd9YdF2JB/LgutWsWt/s=",
    version = "v0.6.1",
)

go_repository(
    name = "com_github_nrdcg_namesilo",
    importpath = "github.com/nrdcg/namesilo",
    sum = "h1:kLjCjsufdW/IlC+iSfAqj0iQGgKjlbUUeDJio5Y6eMg=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_nxadm_tail",
    importpath = "github.com/nxadm/tail",
    sum = "h1:nPr65rt6Y5JFSKQO7qToXr7pePgD6Gwiw05lkbyAQTE=",
    version = "v1.4.8",
)

go_repository(
    name = "com_github_oklog_ulid",
    importpath = "github.com/oklog/ulid",
    sum = "h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_oneofone_xxhash",
    importpath = "github.com/OneOfOne/xxhash",
    sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_opendns_vegadns2client",
    importpath = "github.com/OpenDNS/vegadns2client",
    sum = "h1:xPMsUicZ3iosVPSIP7bW5EcGUzjiiMl1OYTe14y/R24=",
    version = "v0.0.0-20180418235048-a3fa4a771d87",
)

go_repository(
    name = "com_github_openpeedeep_depguard",
    importpath = "github.com/OpenPeeDeeP/depguard",
    sum = "h1:VlW4R6jmBIv3/u1JNlawEvJMM4J+dPORPaZasQee8Us=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_oracle_oci_go_sdk",
    importpath = "github.com/oracle/oci-go-sdk",
    sum = "h1:oj5ESjXwwkFRdhZSnPlShvLWYdt/IZ65RQxveYM3maA=",
    version = "v7.0.0+incompatible",
)

go_repository(
    name = "com_github_ovh_go_ovh",
    importpath = "github.com/ovh/go-ovh",
    sum = "h1:37VE5TYj2m/FLA9SNr4z0+A0JefvTmR60Zwf8XSEV7c=",
    version = "v0.0.0-20181109152953-ba5adb4cf014",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    importpath = "github.com/pelletier/go-toml",
    sum = "h1:zeC5b1GviRUyKYd6OJPvBU/mcVDVoL1OhT17FCt5dSQ=",
    version = "v1.9.3",
)

go_repository(
    name = "com_github_pquerna_otp",
    importpath = "github.com/pquerna/otp",
    sum = "h1:TBZrpfnzVbgmpYhiYBK+bJ4Ig0+ye+GGNMe2pTrvxCo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_prometheus_tsdb",
    importpath = "github.com/prometheus/tsdb",
    sum = "h1:YZcsG11NqnK4czYLrWd9mpEuAJIHVQLwdrleYfszMAA=",
    version = "v0.7.1",
)

go_repository(
    name = "com_github_quasilyte_go_consistent",
    importpath = "github.com/quasilyte/go-consistent",
    sum = "h1:JoUA0uz9U0FVFq5p4LjEq4C0VgQ0El320s3Ms0V4eww=",
    version = "v0.0.0-20190521200055-c6f3937de18c",
)

go_repository(
    name = "com_github_rainycape_memcache",
    importpath = "github.com/rainycape/memcache",
    sum = "h1:dq90+d51/hQRaHEqRAsQ1rE/pC1GUS4sc2rCbbFsAIY=",
    version = "v0.0.0-20150622160815-1031fa0ce2f2",
)

go_repository(
    name = "com_github_rs_xid",
    importpath = "github.com/rs/xid",
    sum = "h1:mhH9Nq+C1fY2l1XIpgxIiUOfNpRBYH1kKcr+qfKgjRc=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_russellhaering_goxmldsig",
    importpath = "github.com/russellhaering/goxmldsig",
    sum = "h1:lK/zeJie2sqG52ZAlPNn1oBBqsIsEKypUUBGpYYF6lk=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_russross_blackfriday",
    importpath = "github.com/russross/blackfriday",
    sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_sacloud_libsacloud",
    importpath = "github.com/sacloud/libsacloud",
    sum = "h1:td3Kd7lvpSAxxHEVpnaZ9goHmmhi0D/RfP0Rqqf/kek=",
    version = "v1.26.1",
)

go_repository(
    name = "com_github_samfoo_ansi",
    importpath = "github.com/samfoo/ansi",
    sum = "h1:CmSpbxmewNQbzqztaY0bke1qzHhyNyC29wYgh17Gxfo=",
    version = "v0.0.0-20160124022901-b6bd2ded7189",
)

go_repository(
    name = "com_github_satori_go_uuid",
    importpath = "github.com/satori/go.uuid",
    sum = "h1:0uYX9dsZ2yD7q2RtLRtPSdGDWzjeM3TbMJP9utgA0ww=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_securego_gosec",
    importpath = "github.com/securego/gosec",
    sum = "h1:rq2/kILQnPtq5oL4+IAjgVOjh5e2yj2aaCYi7squEvI=",
    version = "v0.0.0-20200401082031-e946c8c39989",
)

go_repository(
    name = "com_github_shirou_gopsutil",
    importpath = "github.com/shirou/gopsutil",
    sum = "h1:WokF3GuxBeL+n4Lk4Fa8v9mbdjlrl7bHuneF4N1bk2I=",
    version = "v0.0.0-20190901111213-e4ec7b275ada",
)

go_repository(
    name = "com_github_shirou_w32",
    importpath = "github.com/shirou/w32",
    sum = "h1:udFKJ0aHUL60LboW/A+DfgoHVedieIzIXE8uylPue0U=",
    version = "v0.0.0-20160930032740-bb4de0191aa4",
)

go_repository(
    name = "com_github_shurcool_component",
    importpath = "github.com/shurcooL/component",
    sum = "h1:Fth6mevc5rX7glNLpbAMJnqKlfIkcTjZCSHEeqvKbcI=",
    version = "v0.0.0-20170202220835-f88ec8f54cc4",
)

go_repository(
    name = "com_github_shurcool_events",
    importpath = "github.com/shurcooL/events",
    sum = "h1:vabduItPAIz9px5iryD5peyx7O3Ya8TBThapgXim98o=",
    version = "v0.0.0-20181021180414-410e4ca65f48",
)

go_repository(
    name = "com_github_shurcool_github_flavored_markdown",
    importpath = "github.com/shurcooL/github_flavored_markdown",
    sum = "h1:qb9IthCFBmROJ6YBS31BEMeSYjOscSiG+EO+JVNTz64=",
    version = "v0.0.0-20181002035957-2122de532470",
)

go_repository(
    name = "com_github_shurcool_go",
    importpath = "github.com/shurcooL/go",
    sum = "h1:aSISeOcal5irEhJd1M+IrApc0PdcN7e7Aj4yuEnOrfQ=",
    version = "v0.0.0-20200502201357-93f07166e636",
)

go_repository(
    name = "com_github_shurcool_go_goon",
    importpath = "github.com/shurcooL/go-goon",
    sum = "h1:llrF3Fs4018ePo4+G/HV/uQUqEI1HMDjCeOf2V6puPc=",
    version = "v0.0.0-20170922171312-37c2f522c041",
)

go_repository(
    name = "com_github_shurcool_gofontwoff",
    importpath = "github.com/shurcooL/gofontwoff",
    sum = "h1:Yoy/IzG4lULT6qZg62sVC+qyBL8DQkmD2zv6i7OImrc=",
    version = "v0.0.0-20180329035133-29b52fc0a18d",
)

go_repository(
    name = "com_github_shurcool_gopherjslib",
    importpath = "github.com/shurcooL/gopherjslib",
    sum = "h1:UOk+nlt1BJtTcH15CT7iNO7YVWTfTv/DNwEAQHLIaDQ=",
    version = "v0.0.0-20160914041154-feb6d3990c2c",
)

go_repository(
    name = "com_github_shurcool_highlight_diff",
    importpath = "github.com/shurcooL/highlight_diff",
    sum = "h1:vYEG87HxbU6dXj5npkeulCS96Dtz5xg3jcfCgpcvbIw=",
    version = "v0.0.0-20170515013008-09bb4053de1b",
)

go_repository(
    name = "com_github_shurcool_highlight_go",
    importpath = "github.com/shurcooL/highlight_go",
    sum = "h1:7pDq9pAMCQgRohFmd25X8hIH8VxmT3TaDm+r9LHxgBk=",
    version = "v0.0.0-20181028180052-98c3abbbae20",
)

go_repository(
    name = "com_github_shurcool_home",
    importpath = "github.com/shurcooL/home",
    sum = "h1:MPblCbqA5+z6XARjScMfz1TqtJC7TuTRj0U9VqIBs6k=",
    version = "v0.0.0-20181020052607-80b7ffcb30f9",
)

go_repository(
    name = "com_github_shurcool_htmlg",
    importpath = "github.com/shurcooL/htmlg",
    sum = "h1:crYRwvwjdVh1biHzzciFHe8DrZcYrVcZFlJtykhRctg=",
    version = "v0.0.0-20170918183704-d01228ac9e50",
)

go_repository(
    name = "com_github_shurcool_httperror",
    importpath = "github.com/shurcooL/httperror",
    sum = "h1:eHRtZoIi6n9Wo1uR+RU44C247msLWwyA89hVKwRLkMk=",
    version = "v0.0.0-20170206035902-86b7830d14cc",
)

go_repository(
    name = "com_github_shurcool_httpfs",
    importpath = "github.com/shurcooL/httpfs",
    sum = "h1:bUGsEnyNbVPw06Bs80sCeARAlK8lhwqGyi6UT8ymuGk=",
    version = "v0.0.0-20190707220628-8d4bc4ba7749",
)

go_repository(
    name = "com_github_shurcool_httpgzip",
    importpath = "github.com/shurcooL/httpgzip",
    sum = "h1:fxoFD0in0/CBzXoyNhMTjvBZYW6ilSnTw7N7y/8vkmM=",
    version = "v0.0.0-20180522190206-b1c53ac65af9",
)

go_repository(
    name = "com_github_shurcool_issues",
    importpath = "github.com/shurcooL/issues",
    sum = "h1:T4wuULTrzCKMFlg3HmKHgXAF8oStFb/+lOIupLV2v+o=",
    version = "v0.0.0-20181008053335-6292fdc1e191",
)

go_repository(
    name = "com_github_shurcool_issuesapp",
    importpath = "github.com/shurcooL/issuesapp",
    sum = "h1:Y+TeIabU8sJD10Qwd/zMty2/LEaT9GNDaA6nyZf+jgo=",
    version = "v0.0.0-20180602232740-048589ce2241",
)

go_repository(
    name = "com_github_shurcool_notifications",
    importpath = "github.com/shurcooL/notifications",
    sum = "h1:TQVQrsyNaimGwF7bIhzoVC9QkKm4KsWd8cECGzFx8gI=",
    version = "v0.0.0-20181007000457-627ab5aea122",
)

go_repository(
    name = "com_github_shurcool_octicon",
    importpath = "github.com/shurcooL/octicon",
    sum = "h1:bu666BQci+y4S0tVRVjsHUeRon6vUXmsGBwdowgMrg4=",
    version = "v0.0.0-20181028054416-fa4f57f9efb2",
)

go_repository(
    name = "com_github_shurcool_reactions",
    importpath = "github.com/shurcooL/reactions",
    sum = "h1:LneqU9PHDsg/AkPDU3AkqMxnMYL+imaqkpflHu73us8=",
    version = "v0.0.0-20181006231557-f2e0b4ca5b82",
)

go_repository(
    name = "com_github_shurcool_users",
    importpath = "github.com/shurcooL/users",
    sum = "h1:YGaxtkYjb8mnTvtufv2LKLwCQu2/C7qFB7UtrOlTWOY=",
    version = "v0.0.0-20180125191416-49c67e49c537",
)

go_repository(
    name = "com_github_shurcool_webdavfs",
    importpath = "github.com/shurcooL/webdavfs",
    sum = "h1:JtcyT0rk/9PKOdnKQzuDR+FSjh7SGtJwpgVpfZBRKlQ=",
    version = "v0.0.0-20170829043945-18c3829fa133",
)

go_repository(
    name = "com_github_skip2_go_qrcode",
    importpath = "github.com/skip2/go-qrcode",
    sum = "h1:MRM5ITcdelLK2j1vwZ3Je0FKVCfqOLp5zO6trqMLYs0=",
    version = "v0.0.0-20200617195104-da1b6568686e",
)

go_repository(
    name = "com_github_skratchdot_open_golang",
    importpath = "github.com/skratchdot/open-golang",
    sum = "h1:fyKiXKO1/I/B6Y2U8T7WdQGWzwehOuGIrljPtt7YTTI=",
    version = "v0.0.0-20160302144031-75fb7ed4208c",
)

go_repository(
    name = "com_github_smallstep_assert",
    importpath = "github.com/smallstep/assert",
    sum = "h1:unQFBIznI+VYD1/1fApl1A+9VcBk+9dcqGfnePY87LY=",
    version = "v0.0.0-20200723003110-82e2b9b3b262",
)

go_repository(
    name = "com_github_smallstep_certificates",
    importpath = "github.com/smallstep/certificates",
    sum = "h1:wW344Q/QpupjKKFKa4PqzEXfwgeq/54dkU/HNvGnwQQ=",
    version = "v0.19.0",
)

go_repository(
    name = "com_github_smallstep_certinfo",
    importpath = "github.com/smallstep/certinfo",
    sum = "h1:XO3f9krMkKFDRG4CJFdb3vVoLaXMDAhCn3g0padk9EI=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_smallstep_cli",
    importpath = "github.com/smallstep/cli",
    sum = "h1:BslbUHuMfj/LbVHxuZ4Hv1sL+vAHHidqia4JRoCBwXs=",
    version = "v0.18.0",
)

go_repository(
    name = "com_github_smallstep_nosql",
    importpath = "github.com/smallstep/nosql",
    sum = "h1:Go3WYwttUuvwqMtFiiU4g7kBIlY+hR0bIZAqVdakQ3M=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_smallstep_truststore",
    importpath = "github.com/smallstep/truststore",
    sum = "h1:JUTkQ4oHr40jHTS/A2t0usEhteMWG+45CDD2iJA/dIk=",
    version = "v0.11.0",
)

go_repository(
    name = "com_github_smallstep_zcrypto",
    importpath = "github.com/smallstep/zcrypto",
    sum = "h1:q0IDrQpquiWcU1nJmjwEremPwG+pT2AAGsDatPgg3Kw=",
    version = "v0.0.0-20210924233136-66c2600f6e71",
)

go_repository(
    name = "com_github_smallstep_zlint",
    importpath = "github.com/smallstep/zlint",
    sum = "h1:mZxeNDk2+xg4eTEuR4y6z2+i6HBvcWjtdscDiUmBRzc=",
    version = "v0.0.0-20180727184541-d84eaafe274f",
)

go_repository(
    name = "com_github_sourcegraph_annotate",
    importpath = "github.com/sourcegraph/annotate",
    sum = "h1:yKm7XZV6j9Ev6lojP2XaIshpT4ymkqhMeSghO5Ps00E=",
    version = "v0.0.0-20160123013949-f4cad6c6324d",
)

go_repository(
    name = "com_github_sourcegraph_go_diff",
    importpath = "github.com/sourcegraph/go-diff",
    sum = "h1:lhIKJ2nXLZZ+AfbHpYxTn0pXpNTTui0DX7DO3xeb1Zs=",
    version = "v0.5.3",
)

go_repository(
    name = "com_github_sourcegraph_syntaxhighlight",
    importpath = "github.com/sourcegraph/syntaxhighlight",
    sum = "h1:qpG93cPwA5f7s/ZPBJnGOYQNK/vKsaDaseuKT5Asee8=",
    version = "v0.0.0-20170531221838-bd320f5d308e",
)

go_repository(
    name = "com_github_spaolacci_murmur3",
    importpath = "github.com/spaolacci/murmur3",
    sum = "h1:7c1g84S4BPRrfL5Xrdp6fOJ206sU9y293DDHaoy0bLI=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_spf13_cast",
    importpath = "github.com/spf13/cast",
    sum = "h1:s0hze+J0196ZfEMTs80N7UlFt0BDuQ7Q+JDnHiMWKdA=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    importpath = "github.com/spf13/jwalterweatherman",
    sum = "h1:ue6voC5bR5F8YxI5S67j9i582FU4Qvo2bmqnqMYADFk=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_spf13_viper",
    importpath = "github.com/spf13/viper",
    sum = "h1:Kq1fyeebqsBfbjZj4EL7gj2IO0mMaiyjYUWcUsl2O44=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_stackexchange_wmi",
    importpath = "github.com/StackExchange/wmi",
    sum = "h1:VIkavFPXSjcnS+O8yTq7NI32k0R5Aj+v39y29VYDOSA=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_subosito_gotenv",
    importpath = "github.com/subosito/gotenv",
    sum = "h1:Slr1R9HxAlEKefgq5jn9U+DnETlIUa6HfgEzj0g5d7s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_tarm_serial",
    importpath = "github.com/tarm/serial",
    sum = "h1:UyzmZLoiDWMRywV4DUYb9Fbt8uiOSooupjTq10vpvnU=",
    version = "v0.0.0-20180830185346-98f6abe2eb07",
)

go_repository(
    name = "com_github_thomasrooney_gexpect",
    importpath = "github.com/ThomasRooney/gexpect",
    sum = "h1:CjexZrggt4RldpEUXFZf52vSO3cnmFaqW6B4wADj05Q=",
    version = "v0.0.0-20161231170123-5482f0350944",
)

go_repository(
    name = "com_github_timakin_bodyclose",
    importpath = "github.com/timakin/bodyclose",
    sum = "h1:ig99OeTyDwQWhPe2iw9lwfQVF1KB3Q4fpP3X7/2VBG8=",
    version = "v0.0.0-20200424151742-cb6215831a94",
)

go_repository(
    name = "com_github_timewasted_linode",
    importpath = "github.com/timewasted/linode",
    sum = "h1:CpHxIaZzVy26GqJn8ptRyto8fuoYOd1v0fXm9bG3wQ8=",
    version = "v0.0.0-20160829202747-37e84520dcf7",
)

go_repository(
    name = "com_github_tomasen_realip",
    importpath = "github.com/tomasen/realip",
    sum = "h1:fb190+cK2Xz/dvi9Hv8eCYJYvIGUTN2/KLq1pT6CjEc=",
    version = "v0.0.0-20180522021738-f0c99a92ddce",
)

go_repository(
    name = "com_github_tommy_muehle_go_mnd",
    importpath = "github.com/tommy-muehle/go-mnd",
    sum = "h1:RC4maTWLKKwb7p1cnoygsbKIgNlJqSYBeAFON3Ar8As=",
    version = "v1.3.1-0.20200224220436-e6f9a994e8fa",
)

go_repository(
    name = "com_github_transip_gotransip",
    importpath = "github.com/transip/gotransip",
    sum = "h1:clyOmELPZd2LuFEyuo1mP6RXpbAW75PwD+RfDj4kBm0=",
    version = "v0.0.0-20190812104329-6d8d9179b66f",
)

go_repository(
    name = "com_github_transip_gotransip_v6",
    importpath = "github.com/transip/gotransip/v6",
    sum = "h1:rOCMY607PYF+YvMHHtJt7eZRd0mx/uhyz6dsXWPmn+4=",
    version = "v6.0.2",
)

go_repository(
    name = "com_github_uber_go_atomic",
    importpath = "github.com/uber-go/atomic",
    sum = "h1:Azu9lPBWRNKzYXSIwRfgRuDuS0YKsK4NFhiQv98gkxo=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_ultraware_funlen",
    importpath = "github.com/ultraware/funlen",
    sum = "h1:Av96YVBwwNSe4MLR7iI/BIa3VyI7/djnto/pK3Uxbdo=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_ultraware_whitespace",
    importpath = "github.com/ultraware/whitespace",
    sum = "h1:If7Va4cM03mpgrNH9k49/VOicWpGoG70XPBFFODYDsg=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_uudashr_gocognit",
    importpath = "github.com/uudashr/gocognit",
    sum = "h1:MoG2fZ0b/Eo7NXoIwCVFLG5JED3qgQz5/NEE+rOsjPs=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_valyala_bytebufferpool",
    importpath = "github.com/valyala/bytebufferpool",
    sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_valyala_fasthttp",
    importpath = "github.com/valyala/fasthttp",
    sum = "h1:Z7kVhKP9NZz+tCSY7AVhCMPPAk7b+e5fq0l/BfdTlFc=",
    version = "v1.13.1",
)

go_repository(
    name = "com_github_valyala_fasttemplate",
    importpath = "github.com/valyala/fasttemplate",
    sum = "h1:tY9CJiPnMXf1ERmG2EyK7gNUd+c6RKGD0IfU8WdUSz8=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_valyala_quicktemplate",
    importpath = "github.com/valyala/quicktemplate",
    sum = "h1:BaO1nHTkspYzmAjPXj0QiDJxai96tlcZyKcI9dyEGvM=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_valyala_tcplisten",
    importpath = "github.com/valyala/tcplisten",
    sum = "h1:0R4NLDRDZX6JcmhJgXi5E4b8Wg84ihbmUKp/GvSPEzc=",
    version = "v0.0.0-20161114210144-ceec8f93295a",
)

go_repository(
    name = "com_github_viant_assertly",
    importpath = "github.com/viant/assertly",
    sum = "h1:5x1GzBaRteIwTr5RAGFVG14uNeRFxVNbXPWrK2qAgpc=",
    version = "v0.4.8",
)

go_repository(
    name = "com_github_viant_toolbox",
    importpath = "github.com/viant/toolbox",
    sum = "h1:6TteTDQ68CjgcCe8wH3D3ZhUQQOJXMTbj/D9rkk2a1k=",
    version = "v0.24.0",
)

go_repository(
    name = "com_github_vulcand_oxy",
    importpath = "github.com/vulcand/oxy",
    sum = "h1:DbBijGo1+6cFqR9jarkMxasdj0lgWwrrFtue6ijek4Q=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_vulcand_predicate",
    importpath = "github.com/vulcand/predicate",
    sum = "h1:Gq/uWopa4rx/tnZu2opOSBqHK63Yqlou/SzrbwdJiNg=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_vultr_govultr",
    importpath = "github.com/vultr/govultr",
    sum = "h1:UnNMixYFVO0p80itc8PcweoVENyo1PasfvwKhoasR9U=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_weppos_publicsuffix_go",
    importpath = "github.com/weppos/publicsuffix-go",
    sum = "h1:YSnfg3V65LcCFKtIGKGoBhkyKolEd0hlipcXaOjdnQw=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_xeipuuv_gojsonpointer",
    importpath = "github.com/xeipuuv/gojsonpointer",
    sum = "h1:J9EGpcZtP0E/raorCMxlFGSTBrsSlaDGf3jU/qvAE2c=",
    version = "v0.0.0-20180127040702-4e3ac2762d5f",
)

go_repository(
    name = "com_github_xeipuuv_gojsonreference",
    importpath = "github.com/xeipuuv/gojsonreference",
    sum = "h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=",
    version = "v0.0.0-20180127040603-bd5ef7bd5415",
)

go_repository(
    name = "com_github_xeipuuv_gojsonschema",
    importpath = "github.com/xeipuuv/gojsonschema",
    sum = "h1:LhYJRs+L4fBtjZUfuSZIKGeVu0QRy8e5Xi7D17UxZ74=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_xordataexchange_crypt",
    importpath = "github.com/xordataexchange/crypt",
    sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
    version = "v0.0.3-0.20170626215501-b2862e3d0a77",
)

go_repository(
    name = "com_github_yuin_goldmark_highlighting",
    importpath = "github.com/yuin/goldmark-highlighting",
    sum = "h1:yHfZyN55+5dp1wG7wDKv8HQ044moxkyGq12KFFMFDxg=",
    version = "v0.0.0-20220208100518-594be1970594",
)

go_repository(
    name = "com_github_zenazn_goji",
    importpath = "github.com/zenazn/goji",
    sum = "h1:RSQQAbXGArQ0dIDEq+PI6WqN6if+5KHu6x2Cx/GXLTQ=",
    version = "v0.9.0",
)

go_repository(
    name = "com_github_zmap_rc2",
    importpath = "github.com/zmap/rc2",
    sum = "h1:Nzukz5fNOBIHOsnP+6I79kPx3QhLv8nBy2mfFhBRq30=",
    version = "v0.0.0-20190804163417-abaa70531248",
)

go_repository(
    name = "com_github_zmap_zcertificate",
    importpath = "github.com/zmap/zcertificate",
    sum = "h1:hxHelFG6LEcCsUyu6oKo4P7ZkmzLLeQhOZlVtaUymBk=",
    version = "v0.0.0-20190521191901-30e388164f71",
)

go_repository(
    name = "com_github_zmap_zcrypto",
    importpath = "github.com/zmap/zcrypto",
    sum = "h1:EoQIHS1co8tkbljRLMADWiRAWLcKI02M/ZtPrAUxjHc=",
    version = "v0.0.0-20190329181646-dff83107394d",
)

go_repository(
    name = "com_github_zmap_zlint",
    importpath = "github.com/zmap/zlint",
    sum = "h1:174pnZ4WOF6mGuOJy7Qm6V3cmWn61CfhAWMxvPhqwmc=",
    version = "v0.0.0-20190516161541-9047d02cf65a",
)

go_repository(
    name = "com_shuralyov_dmitri_app_changes",
    importpath = "dmitri.shuralyov.com/app/changes",
    sum = "h1:hJiie5Bf3QucGRa4ymsAUOxyhYwGEz1xrsVk0P8erlw=",
    version = "v0.0.0-20180602232624-0a106ad413e3",
)

go_repository(
    name = "com_shuralyov_dmitri_html_belt",
    importpath = "dmitri.shuralyov.com/html/belt",
    sum = "h1:SPOUaucgtVls75mg+X7CXigS71EnsfVUK/2CgVrwqgw=",
    version = "v0.0.0-20180602232347-f7d459c86be0",
)

go_repository(
    name = "com_shuralyov_dmitri_service_change",
    importpath = "dmitri.shuralyov.com/service/change",
    sum = "h1:GvWw74lx5noHocd+f6HBMXK6DuggBB1dhVkuGZbv7qM=",
    version = "v0.0.0-20181023043359-a85b471d5412",
)

go_repository(
    name = "com_shuralyov_dmitri_state",
    importpath = "dmitri.shuralyov.com/state",
    sum = "h1:ivON6cwHK1OH26MZyWDCnbTRZZf0IhNsENoNAKFS1g4=",
    version = "v0.0.0-20180228185332-28bcc343414c",
)

go_repository(
    name = "com_sourcegraph_sourcegraph_go_diff",
    importpath = "sourcegraph.com/sourcegraph/go-diff",
    sum = "h1:eTiIR0CoWjGzJcnQ3OkhIl/b9GJovq4lSAVRt0ZFEG8=",
    version = "v0.5.0",
)

go_repository(
    name = "com_sourcegraph_sqs_pbtypes",
    importpath = "sourcegraph.com/sqs/pbtypes",
    sum = "h1:f7lAwqviDEGvON4kRv0o5V7FT/IQK+tbkF664XMbP3o=",
    version = "v1.0.0",
)

go_repository(
    name = "in_gopkg_airbrake_gobrake_v2",
    importpath = "gopkg.in/airbrake/gobrake.v2",
    sum = "h1:7z2uVWwn7oVeeugY1DtlPAy5H+KYgB1KeKTnqjNatLo=",
    version = "v2.0.9",
)

go_repository(
    name = "in_gopkg_asn1_ber_v1",
    importpath = "gopkg.in/asn1-ber.v1",
    sum = "h1:TxyelI5cVkbREznMhfzycHdkp5cLA7DpE+GKjSslYhM=",
    version = "v1.0.0-20181015200546-f715ec2f112d",
)

go_repository(
    name = "in_gopkg_gemnasium_logrus_airbrake_hook_v2",
    importpath = "gopkg.in/gemnasium/logrus-airbrake-hook.v2",
    sum = "h1:OAj3g0cR6Dx/R07QgQe8wkA9RNjB2u4i700xBkIT4e0=",
    version = "v2.1.2",
)

go_repository(
    name = "in_gopkg_h2non_gock_v1",
    importpath = "gopkg.in/h2non/gock.v1",
    sum = "h1:SzLqcIlb/fDfg7UvukMpNcWsu7sI5tWwL+KCATZqks0=",
    version = "v1.0.15",
)

go_repository(
    name = "in_gopkg_ini_v1",
    importpath = "gopkg.in/ini.v1",
    sum = "h1:duBzk771uxoUuOlyRLkHsygud9+5lrlGjdFBb4mSKDU=",
    version = "v1.62.0",
)

go_repository(
    name = "in_gopkg_natefinch_lumberjack_v2",
    importpath = "gopkg.in/natefinch/lumberjack.v2",
    sum = "h1:1Lc07Kr7qY4U2YPouBjpCLxpiyxIVoxqXgkXLknAOE8=",
    version = "v2.0.0",
)

go_repository(
    name = "in_gopkg_ns1_ns1_go_v2",
    importpath = "gopkg.in/ns1/ns1-go.v2",
    sum = "h1:GAcf+t0o8gdJAdSFYdE9wChu4bIyguMVqz0RHiFL5VY=",
    version = "v2.0.0-20190730140822-b51389932cbc",
)

go_repository(
    name = "in_gopkg_square_go_jose_v2",
    importpath = "gopkg.in/square/go-jose.v2",
    sum = "h1:NGk74WTnPKBNUhNzQX7PYcTLUjoq7mzKk2OKbvwk2iI=",
    version = "v2.6.0",
)

go_repository(
    name = "io_opencensus_go_contrib_exporter_ocagent",
    importpath = "contrib.go.opencensus.io/exporter/ocagent",
    sum = "h1:TKXjQSRS0/cCDrP7KvkgU6SmILtF/yV2TOs/02K/WZQ=",
    version = "v0.5.0",
)

go_repository(
    name = "io_opencensus_go_contrib_exporter_stackdriver",
    importpath = "contrib.go.opencensus.io/exporter/stackdriver",
    sum = "h1:hAIKGfweGma95MG+PZZTN2rD2Vc1DKmUmPeRYlG6j7c=",
    version = "v0.13.7",
)

go_repository(
    name = "io_opencensus_go_contrib_resource",
    importpath = "contrib.go.opencensus.io/resource",
    sum = "h1:4r2CANuYhKGmYWP02+5E94rLRcS/YeD+KlxSrOsMxk0=",
    version = "v0.1.1",
)

go_repository(
    name = "net_howett_plist",
    importpath = "howett.net/plist",
    sum = "h1:7CrbWYbPPO/PyNy38b2EB/+gYbjCe2DXBxgtOOZbSQM=",
    version = "v1.0.0",
)

go_repository(
    name = "net_launchpad_gocheck",
    importpath = "launchpad.net/gocheck",
    sum = "h1:Izowp2XBH6Ya6rv+hqbceQyw/gSGoXfH/UPoTGduL54=",
    version = "v0.0.0-20140225173054-000000000087",
)

go_repository(
    name = "org_apache_git_thrift_git",
    importpath = "git.apache.org/thrift.git",
    sum = "h1:OR8VhtwhcAI3U48/rzBsVOuHi0zDPzYI1xASVcdSgR8=",
    version = "v0.0.0-20180902110319-2566ecd5d999",
)

go_repository(
    name = "org_go4",
    importpath = "go4.org",
    sum = "h1:+hE86LblG4AyDgwMCLTE6FOlM9+qjHSYS+rKqxUVdsM=",
    version = "v0.0.0-20180809161055-417644f6feb5",
)

go_repository(
    name = "org_go4_grpc",
    importpath = "grpc.go4.org",
    sum = "h1:tmXTu+dfa+d9Evp8NpJdgOy6+rt8/x4yG7qPBrtNfLY=",
    version = "v0.0.0-20170609214715-11d0a25b4919",
)

go_repository(
    name = "org_golang_x_build",
    importpath = "golang.org/x/build",
    sum = "h1:E2M5QgjZ/Jg+ObCQAudsXxuTsLj7Nl5RV/lZcQZmKSo=",
    version = "v0.0.0-20190111050920-041ab4dc3f9d",
)

go_repository(
    name = "org_golang_x_perf",
    importpath = "golang.org/x/perf",
    sum = "h1:rWw0Gj4DMl/2otJ8CnfTcwOWkpROAc6qhXXoMrYOCgo=",
    version = "v0.0.0-20210220033136-40a54f11e909",
)

go_repository(
    name = "sm_step_go_crypto",
    importpath = "go.step.sm/crypto",
    sum = "h1:4mnZk21cSxyMGxsEpJwZKKvJvDu1PN09UVrWWFNUBdk=",
    version = "v0.16.1",
)

go_repository(
    name = "net_starlark_go",
    importpath = "go.starlark.net",
    sum = "h1:Uo/x0Ir5vQJ+683GXB9Ug+4fcjsbp7z7Ul8UaZbhsRM=",
    version = "v0.0.0-20220328144851-d1966c6b9fcd",
)

go_repository(
    name = "com_github_xiaq_persistent",
    importpath = "github.com/xiaq/persistent",
    sum = "h1:HxX+USVm4XyGwvWS0eJy+GMttkfSRdFcrZ46WtAs5RQ=",
    version = "v0.0.0-20200820214153-3175cfb92e14",
)

go_repository(
    name = "sh_elv_src",
    importpath = "src.elv.sh",
    sum = "h1:FNmdopFKRXUhuHFpfJRCbDL+iukbLqit2slJBs6wEBA=",
    version = "v0.18.0",
)

go_repository(
    name = "io_robpike_ivy",
    importpath = "robpike.io/ivy",
    sum = "h1:XVYgWSm7THVm1bk1jfRBB9xhso459T22gbpotHbhk7M=",
    version = "v0.2.7",
)

go_repository(
    name = "com_github_traefik_yaegi",
    importpath = "github.com/traefik/yaegi",
    sum = "h1:TuuIc0TC4oaWkVngjVAKkFd4fH35B0B95DmbS76uqs8=",
    version = "v0.11.3",
)

go_repository(
    name = "com_github_rancher_dapper",
    importpath = "github.com/rancher/dapper",
    sum = "h1:DfuPkw6hL6iN1qVuuDtaR9CtJTzEmXXpZfnCkf+QzwE=",
    version = "v0.5.8",
)

go_repository(
    name = "com_github_casbin_caddy_authz_v2",
    importpath = "github.com/casbin/caddy-authz/v2",
    sum = "h1:vUUoBIbTa8k02/zF17Qk0EjJIsu3HSU0ee2lX1nKo54=",
    version = "v2.0.0",
)

go_repository(
    name = "com_github_goproxy_goproxy",
    importpath = "github.com/goproxy/goproxy",
    sum = "h1:gqFFOnckEtcxMX65hjJnxsnbEvQVq1Ia1/1ON9ki5wc=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_minio_md5_simd",
    importpath = "github.com/minio/md5-simd",
    sum = "h1:Gdi1DZK69+ZVMoNHRXJyNcxrMA4dSxoYHZSQbirFg34=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_minio_minio_go_v7",
    importpath = "github.com/minio/minio-go/v7",
    sum = "h1:v+RS2/dpRq+XaarlZItHd3MVjjQcN2noRn4HxmVdmg4=",
    version = "v7.0.9",
)

go_repository(
    name = "com_github_minio_sha256_simd",
    importpath = "github.com/minio/sha256-simd",
    sum = "h1:v1ta+49hkWZyvaKwrQB8elexRqm6Y0aMLjCNsrYxo6g=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_goproxyio_goproxy_v2",
    importpath = "github.com/goproxyio/goproxy/v2",
    sum = "h1:Ndmh6mgEPYpY6jKRvZ9j5bTAw5+iRbc9SAgz/91mINU=",
    version = "v2.0.7",
)

go_repository(
    name = "com_github_goproxyio_windows",
    importpath = "github.com/goproxyio/windows",
    sum = "h1:geXWJzWrEzC5S4MjTi8N+fhzuSYgdf4FwcYb4OI/YGU=",
    version = "v0.0.0-20191126033816-f4a809841617",
)

go_repository(
    name = "com_github_ajg_form",
    importpath = "github.com/ajg/form",
    sum = "h1:t9c7v8JUKu/XxOGBU0yjNpaMloxGEJhUkqFRq0ibGeU=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_andrew_d_go_termutil",
    importpath = "github.com/andrew-d/go-termutil",
    sum = "h1:axBiC50cNZOs7ygH5BgQp4N+aYrZ2DNpWZ1KG3VOSOM=",
    version = "v0.0.0-20150726205930-009166a695a2",
)

go_repository(
    name = "com_github_andybalholm_brotli",
    importpath = "github.com/andybalholm/brotli",
    sum = "h1:KqhlKozYbRtJvsPrrEeXcO+N2l6NYT5A2QAFmSULpEc=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_elithrar_simple_scrypt",
    importpath = "github.com/elithrar/simple-scrypt",
    sum = "h1:KIlOlxdoQf9JWKl5lMAJ28SY2URB0XTRDn2TckyzAZg=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_fasthttp_contrib_websocket",
    importpath = "github.com/fasthttp-contrib/websocket",
    sum = "h1:DddqAaWDpywytcG8w/qoQ5sAN8X12d3Z3koB0C3Rxsc=",
    version = "v0.0.0-20160511215533-1f3b11f56072",
)

go_repository(
    name = "com_github_gavv_httpexpect",
    importpath = "github.com/gavv/httpexpect",
    sum = "h1:1X9kcRshkSKEjNJJxX9Y9mQ5BRfbxU5kORdjhlA1yX8=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_imkira_go_interpol",
    importpath = "github.com/imkira/go-interpol",
    sum = "h1:KIiKr0VSG2CUW1hl1jpiyuzuJeKUUpC8iM1AIE7N1Vk=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_jpillora_ansi",
    importpath = "github.com/jpillora/ansi",
    sum = "h1:+Ei5HCAH0xsrQRCT2PDr4mq9r4Gm4tg+arNdXRkB22s=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_jpillora_cookieauth",
    importpath = "github.com/jpillora/cookieauth",
    sum = "h1:1BYnSG4c+/5kutOsY+7/Ba+rRhUfYv61Jrf605CxfRw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jpillora_eventsource",
    importpath = "github.com/jpillora/eventsource",
    sum = "h1:iMFSHw9uUmQyNWKHylAS9HoK9R9ps2NWqsjzKniCFus=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jpillora_ipfilter",
    importpath = "github.com/jpillora/ipfilter",
    sum = "h1:K5zGPjyjgf2MPB+iTULZ7Hl4zXPWOb4JwgxMdogKq20=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_jpillora_opts",
    importpath = "github.com/jpillora/opts",
    sum = "h1:H8vWooV3P9nsqmCcPgxNZyIa7GPOWA1KQFsfAzIkCtE=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_jpillora_requestlog",
    importpath = "github.com/jpillora/requestlog",
    sum = "h1:bg++eJ74T7DYL3DlIpiwknrtfdUA9oP/M4fL+PpqnyA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jpillora_sizestr",
    importpath = "github.com/jpillora/sizestr",
    sum = "h1:4tr0FLxs1Mtq3TnsLDV+GYUWG7Q26a6s+tV5Zfw2ygw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jpillora_velox",
    importpath = "github.com/jpillora/velox",
    sum = "h1:96BPWo+zw2ywf4T/3C9kEKhzGABsjJ4vWSOkf93el/g=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_jpillora_webproc",
    importpath = "github.com/jpillora/webproc",
    sum = "h1:0BmX+F+HDs0uTYzjy3CSU/7oM8yTTuPkqsJim+b5Ndg=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_k0kubun_colorstring",
    importpath = "github.com/k0kubun/colorstring",
    sum = "h1:uC1QfSlInpQF+M0ao65imhwqKnz3Q2z/d8PWZRMQvDM=",
    version = "v0.0.0-20150214042306-9440f1994b88",
)

go_repository(
    name = "com_github_moul_http2curl",
    importpath = "github.com/moul/http2curl",
    sum = "h1:dRMWoAtb+ePxMlLkrCbAqh4TlPHXvoGUSQ323/9Zahs=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_phuslu_geoip",
    importpath = "github.com/phuslu/geoip",
    sum = "h1:pap5n0dO6f2HUOXKGW0OrG0Y9OlxN0uC+XKMvziUm6g=",
    version = "v1.0.20200217",
)

go_repository(
    name = "com_github_rakyll_statik",
    importpath = "github.com/rakyll/statik",
    sum = "h1:OF3QCZUuyPxuGEP7B4ypUa7sB/iHtqOTDYZXGM8KOdQ=",
    version = "v0.1.7",
)

go_repository(
    name = "com_github_yalp_jsonpath",
    importpath = "github.com/yalp/jsonpath",
    sum = "h1:6fRhSjgLCkTD3JnJxvaJ4Sj+TYblw757bqYgZaOq5ZY=",
    version = "v0.0.0-20180802001716-5cc68e5049a0",
)

go_repository(
    name = "com_github_yudai_gojsondiff",
    importpath = "github.com/yudai/gojsondiff",
    sum = "h1:27cbfqXLVEJ1o8I6v3y9lg8Ydm53EKqHXAOMxEGlCOA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_yudai_golcs",
    importpath = "github.com/yudai/golcs",
    sum = "h1:BHyfKlQyqbsFN5p3IfnEUduWvb9is428/nNb5L3U01M=",
    version = "v0.0.0-20170316035057-ecda9a501e82",
)

go_repository(
    name = "com_github_yudai_pp",
    importpath = "github.com/yudai/pp",
    sum = "h1:Q4//iY4pNF6yPLZIigmvcl7k/bPgrcTPIFIcmawg5bI=",
    version = "v2.0.1+incompatible",
)

go_repository(
    name = "xyz_gomodules_jsonpatch_v3",
    importpath = "gomodules.xyz/jsonpatch/v3",
    sum = "h1:Te7hKxV52TKCbNYq3t84tzKav3xhThdvSsSp/W89IyI=",
    version = "v3.0.1",
)

go_repository(
    name = "xyz_gomodules_orderedmap",
    importpath = "gomodules.xyz/orderedmap",
    sum = "h1:fM/+TGh/O1KkqGR5xjTKg6bU8OKBkg7p0Y+x/J9m8Os=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_aead_chacha20",
    importpath = "github.com/aead/chacha20",
    sum = "h1:KjTM2ks9d14ZYCvmHS9iAKVt9AyzRSqNU1qabPih5BY=",
    version = "v0.0.0-20180709150244-8b13a72661da",
)

go_repository(
    name = "com_github_p4gefau1t_trojan_go",
    importpath = "github.com/p4gefau1t/trojan-go",
    sum = "h1:nCXHM+uV61rj4KKMNXPqBZBkz8GoR6sAjOUH00AmB34=",
    version = "v0.10.6",
)

go_repository(
    name = "com_github_refraction_networking_utls",
    importpath = "github.com/refraction-networking/utls",
    sum = "h1:n9NMHJusHylTmtaJ0Qe0VV9dkTZLiwAxHmrI/l98GeE=",
    version = "v0.0.0-20210713165636-0b2885c8c0d4",
)

go_repository(
    name = "com_github_txthinking_runnergroup",
    importpath = "github.com/txthinking/runnergroup",
    sum = "h1:7PflaKRtU4np/epFxRXlFhlzLXZzKFrH5/I4so5Ove0=",
    version = "v0.0.0-20210608031112-152c7c4432bf",
)

go_repository(
    name = "com_v2ray_core",
    importpath = "v2ray.com/core",
    sum = "h1:JWoYsYlCpFOJX5KcmSkAMHOqOjzux+wx/HtgMBkUvSg=",
    version = "v4.19.1+incompatible",
)

go_repository(
    name = "io_h12_socks",
    importpath = "h12.io/socks",
    sum = "h1:Ka3qaQewws4j4/eDQnOdpr4wXsC//dXtWvftlIcCQUo=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_mholt_caddy_l4",
    importpath = "github.com/mholt/caddy-l4",
    sum = "h1:APtotvfGzmVDwL1YKOKcJX2OSg3/bKOhqAiTX13rolk=",
    version = "v0.0.0-20220420174601-aec6535658b1",
)

go_repository(
    name = "com_github_abiosoft_caddy_json_schema",
    importpath = "github.com/abiosoft/caddy-json-schema",
    sum = "h1:LHvWQrY7nt4wSsB/TueV4rePOIjs+cSLz5/O+pRnYmI=",
    version = "v0.0.0-20210703084946-7fd4aca97120",
)

go_repository(
    name = "com_github_asdine_storm",
    importpath = "github.com/asdine/storm",
    sum = "h1:dczuIkyqwY2LrtXPz8ixMrU/OFgZp71kbKTHGrXYt/Q=",
    version = "v2.1.2+incompatible",
)

go_repository(
    name = "com_github_disintegration_imaging",
    importpath = "github.com/disintegration/imaging",
    sum = "h1:w1LecBlG2Lnp8B3jk5zSuNqd7b4DXhcjwek1ei82L+c=",
    version = "v1.6.2",
)

go_repository(
    name = "com_github_dsnet_compress",
    importpath = "github.com/dsnet/compress",
    sum = "h1:iFaUwBSo5Svw6L7HYpRu/0lE3e0BaElwnNO1qkNQxBY=",
    version = "v0.0.2-0.20210315054119-f66993602bf5",
)

go_repository(
    name = "com_github_dsnet_golib",
    importpath = "github.com/dsnet/golib",
    sum = "h1:tFh1tRc4CA31yP6qDcu+Trax5wW5GuMxvkIba07qVLY=",
    version = "v0.0.0-20171103203638-1ea166775780",
)

go_repository(
    name = "com_github_filebrowser_filebrowser_v2",
    importpath = "github.com/filebrowser/filebrowser/v2",
    sum = "h1:2jpAwRnHRf7ZR1Fu5FefFsfhNGG5T27PwQRkpNhTSHM=",
    version = "v2.21.1",
)

go_repository(
    name = "com_github_go_acme_lego",
    importpath = "github.com/go-acme/lego",
    sum = "h1:5fNN9yRQfv8ymH3DSsxla+4aYeQt2IgfZqHKVnK8f0s=",
    version = "v2.5.0+incompatible",
)

go_repository(
    name = "com_github_lucas_clemente_quic_clients",
    importpath = "github.com/lucas-clemente/quic-clients",
    sum = "h1:/P9n0nICT/GnQJkZovtBqridjxU0ao34m7DpMts79qY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_maruel_natural",
    importpath = "github.com/maruel/natural",
    sum = "h1:PEhRT94KBTY4E0KdCYmhvDGWjSFBxc68j2M6PMRix8U=",
    version = "v0.0.0-20180416170133-dbcb3e2e8cf1",
)

go_repository(
    name = "com_github_marusama_semaphore_v2",
    importpath = "github.com/marusama/semaphore/v2",
    sum = "h1:Y29DhhFMvreVgoqF9EtaSJAF9t2E7Sk7i5VW81sqB8I=",
    version = "v2.4.1",
)

go_repository(
    name = "com_github_mholt_archiver",
    importpath = "github.com/mholt/archiver",
    sum = "h1:1dCVxuqs0dJseYEhi5pl7MYPH9zDa1wBi7mF09cbNkU=",
    version = "v3.1.1+incompatible",
)

go_repository(
    name = "com_github_nwaples_rardecode",
    importpath = "github.com/nwaples/rardecode",
    sum = "h1:vSxaY8vQhOcVr4mm5e8XllHWTiM4JF507A0Katqw7MQ=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_sereal_sereal",
    importpath = "github.com/Sereal/Sereal",
    sum = "h1:3trCIB5GsAOIY8NxlfMztCYIhVsW9V5sZ+brsecjaiI=",
    version = "v0.0.0-20190430203904-6faf9605eb56",
)

go_repository(
    name = "com_github_ulikunitz_xz",
    importpath = "github.com/ulikunitz/xz",
    sum = "h1:RsKRIA2MO8x56wkkcd3LbtcE/uMszhb6DpRf+3uwa3I=",
    version = "v0.5.9",
)

go_repository(
    name = "com_github_vmihailenco_msgpack",
    importpath = "github.com/vmihailenco/msgpack",
    sum = "h1:dSLoQfGFAo3F6OoNhwUmLwVgaUXK79GlxNBwueZn0xI=",
    version = "v4.0.4+incompatible",
)

go_repository(
    name = "com_github_xi2_xz",
    importpath = "github.com/xi2/xz",
    sum = "h1:nIPpBwaJSVYIxUFsDv3M8ofmx9yWTog9BfvIu0q41lo=",
    version = "v0.0.0-20171230120015-48954b6210f8",
)

go_repository(
    name = "com_github_agnivade_levenshtein",
    importpath = "github.com/agnivade/levenshtein",
    sum = "h1:3oJU7J3FGFmyhn8KHjmVaZCN5hxTr7GxgRue+sxIXdQ=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_akhenakh_hunspellgo",
    importpath = "github.com/akhenakh/hunspellgo",
    sum = "h1:bYOD6QJnBJY79MJQR1i9cyQePG5oNDZXDKL2bhN/uvE=",
    version = "v0.0.0-20160221122622-9db38fa26e19",
)

go_repository(
    name = "com_github_alexkohler_nakedret",
    importpath = "github.com/alexkohler/nakedret",
    sum = "h1:S/bzOFhZHYUJp6qPmdXdFHS5nlWGFmLmoc8QOydvotE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_andreyvit_diff",
    importpath = "github.com/andreyvit/diff",
    sum = "h1:bvNMNQO63//z+xNgfBlViaCIJKLlCJ6/fmUseuG0wVQ=",
    version = "v0.0.0-20170406064948-c7f18ee00883",
)

go_repository(
    name = "com_github_azure_go_ansiterm",
    importpath = "github.com/Azure/go-ansiterm",
    sum = "h1:UQHMgLO+TxOElx5B5HZ4hJQsoJ/PvUvKRhJHDQXO8P8=",
    version = "v0.0.0-20210617225240-d185dfc1b5a1",
)

go_repository(
    name = "com_github_azure_go_autorest",
    importpath = "github.com/Azure/go-autorest",
    sum = "h1:V5VMDjClD3GiElqLWO7mz2MxNAK/vTfRHdAubSIPRgs=",
    version = "v14.2.0+incompatible",
)

go_repository(
    name = "com_github_bketelsen_crypt",
    importpath = "github.com/bketelsen/crypt",
    sum = "h1:w/jqZtC9YD4DS/Vp9GhWfWcCpuAL58oTnLoI8vE9YHU=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_btwiuse_etcd_v3",
    importpath = "github.com/btwiuse/etcd/v3",
    sum = "h1:v4A+SAVIB5KvGVS7mDyU8ls3uwlgiNCfazwDL/nK1G4=",
    version = "v3.4.15",
)

go_repository(
    name = "com_github_certifi_gocertifi",
    importpath = "github.com/certifi/gocertifi",
    sum = "h1:uH66TXeswKn5PW5zdZ39xEwfS9an067BirqA+P4QaLI=",
    version = "v0.0.0-20200922220541-2c3bb06c6054",
)

go_repository(
    name = "com_github_chzchzchz_goword",
    importpath = "github.com/chzchzchz/goword",
    sum = "h1:0wUHjDfbCAROEAZ96zAJGwcNMkPIheFaIjtQyv3QqfM=",
    version = "v0.0.0-20170907005317-a9744cb52b03",
)

go_repository(
    name = "com_github_cockroachdb_errors",
    importpath = "github.com/cockroachdb/errors",
    sum = "h1:Lap807SXTH5tri2TivECb/4abUkMZC9zRoLarvcKDqs=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_cockroachdb_logtags",
    importpath = "github.com/cockroachdb/logtags",
    sum = "h1:o/kfcElHqOiXqcou5a3rIlMc7oJbMQkeLk0VQJ7zgqY=",
    version = "v0.0.0-20190617123548-eb05cc24525f",
)

go_repository(
    name = "com_github_coreos_go_oidc",
    importpath = "github.com/coreos/go-oidc",
    sum = "h1:mh48q/BqXqgjVHpy2ZY7WnWAbenxRjsz9N1i1YxjHAk=",
    version = "v2.2.1+incompatible",
)

go_repository(
    name = "com_github_coreos_go_systemd_v22",
    importpath = "github.com/coreos/go-systemd/v22",
    sum = "h1:D9/bQk5vlXQFZ6Kwuu6zaiXJ9oTPe68++AzAJc1DzSI=",
    version = "v22.3.2",
)

go_repository(
    name = "com_github_coreos_license_bill_of_materials",
    importpath = "github.com/coreos/license-bill-of-materials",
    sum = "h1:vHRufSa2k8tfkcDdia1vJFa+oiBvvPxW94mg76PPAoA=",
    version = "v0.0.0-20190913234955-13baff47494e",
)

go_repository(
    name = "com_github_docker_go_units",
    importpath = "github.com/docker/go-units",
    sum = "h1:3uh0PgVws3nIA0Q+MwDC8yjEPf9zjRfZZWXZYDct3Tw=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_etcd_io_gofail",
    importpath = "github.com/etcd-io/gofail",
    sum = "h1:Y2I0lxOttdUKz+hNaIdG3FtjuQrTmwXun1opRV65IZc=",
    version = "v0.0.0-20190801230047-ad7f989257ca",
)

go_repository(
    name = "com_github_fatih_structtag",
    importpath = "github.com/fatih/structtag",
    sum = "h1:/OdNE99OxoI/PqaW/SuSK9uxxT3f/tcSZgon/ssNSx4=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_form3tech_oss_jwt_go",
    importpath = "github.com/form3tech-oss/jwt-go",
    sum = "h1:7ZaBxOI7TMoYBfyA3cQHErNNyAWIKUMIwqxEtgHOs5c=",
    version = "v3.2.3+incompatible",
)

go_repository(
    name = "com_github_getsentry_raven_go",
    importpath = "github.com/getsentry/raven-go",
    sum = "h1:no+xWJRb5ZI7eE8TWgIq1jLulQiIoLG0IfYxv5JYMGs=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_globalsign_mgo",
    importpath = "github.com/globalsign/mgo",
    sum = "h1:DujepqpGd1hyOd7aW59XpK7Qymp8iy83xq74fLr21is=",
    version = "v0.0.0-20181015135952-eeefdecb41b8",
)

go_repository(
    name = "com_github_go_openapi_analysis",
    importpath = "github.com/go-openapi/analysis",
    sum = "h1:Ub9e++M8sDwtHD+S587TYi+6ANBG1NRYGZDihqk0SaY=",
    version = "v0.19.16",
)

go_repository(
    name = "com_github_go_openapi_errors",
    importpath = "github.com/go-openapi/errors",
    sum = "h1:9SnKdGhiPZHF3ttwFMiCBEb8jQ4IDdrK+5+a0oTygA4=",
    version = "v0.19.9",
)

go_repository(
    name = "com_github_go_openapi_loads",
    importpath = "github.com/go-openapi/loads",
    sum = "h1:z5p5Xf5wujMxS1y8aP+vxwW5qYT2zdJBbXKmQUG3lcc=",
    version = "v0.20.2",
)

go_repository(
    name = "com_github_go_openapi_runtime",
    importpath = "github.com/go-openapi/runtime",
    sum = "h1:TqagMVlRAOTwllE/7hNKx6rQ10O6T8ZzeJdMjSTKaD4=",
    version = "v0.19.24",
)

go_repository(
    name = "com_github_go_openapi_strfmt",
    importpath = "github.com/go-openapi/strfmt",
    sum = "h1:l2omNtmNbMc39IGptl9BuXBEKcZfS8zjrTsPKTiJiDM=",
    version = "v0.20.0",
)

go_repository(
    name = "com_github_go_openapi_validate",
    importpath = "github.com/go-openapi/validate",
    sum = "h1:QGQ5CvK74E28t3DkegGweKR+auemUi5IdpMc4x3UW6s=",
    version = "v0.20.1",
)

go_repository(
    name = "com_github_gobuffalo_attrs",
    importpath = "github.com/gobuffalo/attrs",
    sum = "h1:hSkbZ9XSyjyBirMeqSqUrK+9HboWrweVlzRNqoBi2d4=",
    version = "v0.0.0-20190224210810-a9411de4debd",
)

go_repository(
    name = "com_github_gobuffalo_depgen",
    importpath = "github.com/gobuffalo/depgen",
    sum = "h1:31atYa/UW9V5q8vMJ+W6wd64OaaTHUrCUXER358zLM4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_gobuffalo_envy",
    importpath = "github.com/gobuffalo/envy",
    sum = "h1:GlXgaiBkmrYMHco6t4j7SacKO4XUjvh5pwXh0f4uxXU=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_gobuffalo_flect",
    importpath = "github.com/gobuffalo/flect",
    sum = "h1:3GQ53z7E3o00C/yy7Ko8VXqQXoJGLkrTQCLTF1EjoXU=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_gobuffalo_genny",
    importpath = "github.com/gobuffalo/genny",
    sum = "h1:iQ0D6SpNXIxu52WESsD+KoQ7af2e3nCfnSBoSF/hKe0=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_gobuffalo_gitgen",
    importpath = "github.com/gobuffalo/gitgen",
    sum = "h1:mSVZ4vj4khv+oThUfS+SQU3UuFIZ5Zo6UNcvK8E8Mz8=",
    version = "v0.0.0-20190315122116-cc086187d211",
)

go_repository(
    name = "com_github_gobuffalo_gogen",
    importpath = "github.com/gobuffalo/gogen",
    sum = "h1:dLg+zb+uOyd/mKeQUYIbwbNmfRsr9hd/WtYWepmayhI=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_gobuffalo_logger",
    importpath = "github.com/gobuffalo/logger",
    sum = "h1:8thhT+kUJMTMy3HlX4+y9Da+BNJck+p109tqqKp7WDs=",
    version = "v0.0.0-20190315122211-86e12af44bc2",
)

go_repository(
    name = "com_github_gobuffalo_mapi",
    importpath = "github.com/gobuffalo/mapi",
    sum = "h1:fq9WcL1BYrm36SzK6+aAnZ8hcp+SrmnDyAxhNx8dvJk=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_gobuffalo_packd",
    importpath = "github.com/gobuffalo/packd",
    sum = "h1:4sGKOD8yaYJ+dek1FDkwcxCHA40M4kfKgFHx8N2kwbU=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_gobuffalo_packr_v2",
    importpath = "github.com/gobuffalo/packr/v2",
    sum = "h1:Ir9W9XIm9j7bhhkKE9cokvtTl1vBm62A/fene/ZCj6A=",
    version = "v2.2.0",
)

go_repository(
    name = "com_github_gobuffalo_syncx",
    importpath = "github.com/gobuffalo/syncx",
    sum = "h1:tpom+2CJmpzAWj5/VEHync2rJGi+epHNIeRSWjzGA+4=",
    version = "v0.0.0-20190224160051-33c29581e754",
)

go_repository(
    name = "com_github_goccy_go_yaml",
    importpath = "github.com/goccy/go-yaml",
    sum = "h1:JuZRFlqLM5cWF6A+waL8AKVuCcqvKOuhJtUQI+L3ez0=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_godbus_dbus_v5",
    importpath = "github.com/godbus/dbus/v5",
    sum = "h1:9349emZab16e7zQvpmsbtjc18ykshndd8y2PG3sgJbA=",
    version = "v5.0.4",
)

go_repository(
    name = "com_github_google_addlicense",
    importpath = "github.com/google/addlicense",
    sum = "h1:cqvo5suPWlsk6r6o42Fs2K66xYCl2tnhVPUYoP3EnO4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_gordonklaus_ineffassign",
    importpath = "github.com/gordonklaus/ineffassign",
    sum = "h1:U/zHjaVG/sECz5xhnh7kPH+Fv/maPbhZPcaTquo5sPg=",
    version = "v0.0.0-20200809085317-e36bfde3bb78",
)

go_repository(
    name = "com_github_gyuho_gocovmerge",
    importpath = "github.com/gyuho/gocovmerge",
    sum = "h1:BGeD3v3lyKZy+ocGtprXiDXjIiXvZDfuyII7Lym7GbQ=",
    version = "v0.0.0-20171205171859-50c7e6afd535",
)

go_repository(
    name = "com_github_hexfusion_schwag",
    importpath = "github.com/hexfusion/schwag",
    sum = "h1:oDcxzjIf33MTX7b8Eu7eO3a/z8mlTT+blyEoVxBmUUg=",
    version = "v0.0.0-20170606222847-b7d0fc9aadaa",
)

go_repository(
    name = "com_github_jmespath_go_jmespath_internal_testify",
    importpath = "github.com/jmespath/go-jmespath/internal/testify",
    sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_joho_godotenv",
    importpath = "github.com/joho/godotenv",
    sum = "h1:Zjp+RcGpHhGlrMbJzXTrZZPrWj+1vfm90La1wgB6Bhc=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_josharian_intern",
    importpath = "github.com/josharian/intern",
    sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_karrick_godirwalk",
    importpath = "github.com/karrick/godirwalk",
    sum = "h1:DynhcF+bztK8gooS0+NDJFrdNZjJ3gzVzC545UNA9iw=",
    version = "v1.16.1",
)

go_repository(
    name = "com_github_magefile_mage",
    importpath = "github.com/magefile/mage",
    sum = "h1:3HiXzCUY12kh9bIuyXShaVe529fJfyqoVM42o/uom2g=",
    version = "v1.10.0",
)

go_repository(
    name = "com_github_markbates_oncer",
    importpath = "github.com/markbates/oncer",
    sum = "h1:JgVTCPf0uBVcUSWpyXmGpgOc62nK5HWUBKAGc3Qqa5k=",
    version = "v0.0.0-20181203154359-bf2de49a0be2",
)

go_repository(
    name = "com_github_markbates_safe",
    importpath = "github.com/markbates/safe",
    sum = "h1:yjZkbvRM6IzKj9tlu/zMJLS0n/V351OZWRnF3QfaUxI=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_mdempsky_unconvert",
    importpath = "github.com/mdempsky/unconvert",
    sum = "h1:Kc3s6QFyh9DLgInXpWKuG+8I7R7lXbnP7mcoOVIt6KY=",
    version = "v0.0.0-20200228143138-95ecdbfc0b5f",
)

go_repository(
    name = "com_github_mgechev_dots",
    importpath = "github.com/mgechev/dots",
    sum = "h1:QASJXOGm2RZ5Ardbc86qNFvby9AqkLDibfChMtAg5QM=",
    version = "v0.0.0-20190921121421-c36f7dcfbb81",
)

go_repository(
    name = "com_github_mgechev_revive",
    importpath = "github.com/mgechev/revive",
    sum = "h1:z3FL6IFFN3JKzHYHD8O1ExH9g/4lAGJ5x1+9rPZgsFg=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_mikefarah_yq_v3",
    importpath = "github.com/mikefarah/yq/v3",
    sum = "h1:lPmsut5Sk7eK2BmDXuvNEvMbT7MkAJBu64Yxr7iJ6nk=",
    version = "v3.0.0-20201202084205-8846255d1c37",
)

go_repository(
    name = "com_github_moby_spdystream",
    importpath = "github.com/moby/spdystream",
    sum = "h1:cjW1zVyyoiM0T7b6UoySUFqzXMoqRckQtXwGPiBhOM8=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_moby_term",
    importpath = "github.com/moby/term",
    sum = "h1:dcztxKSvZ4Id8iPpHERQBbIJfabdt4wUm5qy3wOL2Zc=",
    version = "v0.0.0-20210619224110-3f7ff695adc6",
)

go_repository(
    name = "com_github_montanaflynn_stats",
    importpath = "github.com/montanaflynn/stats",
    sum = "h1:iruDEfMl2E6fbMZ9s0scYfZQ84/6SPL6zC8ACM2oIL0=",
    version = "v0.0.0-20171201202039-1bf9dbcd8cbe",
)

go_repository(
    name = "com_github_pquerna_cachecontrol",
    importpath = "github.com/pquerna/cachecontrol",
    sum = "h1:jWKYCNlX4J5s8M0nHYkh7Y7c9gRVDEb3mq51j5J0F5M=",
    version = "v0.0.0-20201205024021-ac21108117ac",
)

go_repository(
    name = "com_github_tidwall_pretty",
    importpath = "github.com/tidwall/pretty",
    sum = "h1:HsD+QiTn7sK6flMKIvNmpqz1qrpP3Ps6jOKIKMooyg4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_trustmaster_go_aspell",
    importpath = "github.com/trustmaster/go-aspell",
    sum = "h1:92ZQJRegaqnKjz9HY9an696Sw5EmAqRv0eie/U2IE6k=",
    version = "v0.0.0-20200701131845-c2b1f55bec8f",
)

go_repository(
    name = "com_github_vektah_gqlparser",
    importpath = "github.com/vektah/gqlparser",
    sum = "h1:ZsyLGn7/7jDNI+y4SEhI4yAxRChlv15pUHMjijT+e68=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_xdg_scram",
    importpath = "github.com/xdg/scram",
    sum = "h1:u40Z8hqBAAQyv+vATcGgV0YCnDjqSL7/q/JyPhhJSPk=",
    version = "v0.0.0-20180814205039-7eeb5667e42c",
)

go_repository(
    name = "com_github_xdg_stringprep",
    importpath = "github.com/xdg/stringprep",
    sum = "h1:d9X0esnoa3dFsV0FG35rAT0RIhYFlPq7MiP+DW89La0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_google_cloud_go_firestore",
    importpath = "cloud.google.com/go/firestore",
    sum = "h1:9x7Bx0A9R5/M9jibeJeZWqjeVEIxYW9fZYqB9a70/bY=",
    version = "v1.1.0",
)

go_repository(
    name = "in_gopkg_go_playground_assert_v1",
    importpath = "gopkg.in/go-playground/assert.v1",
    sum = "h1:xoYuJVE7KT85PYWrN730RguIQO0ePzVRfFMXadIrXTM=",
    version = "v1.2.1",
)

go_repository(
    name = "in_gopkg_go_playground_validator_v9",
    importpath = "gopkg.in/go-playground/validator.v9",
    sum = "h1:Wk0Z37oBmKj9/n+tPyBHZmeL19LaCoK3Qq48VwYENss=",
    version = "v9.30.0",
)

go_repository(
    name = "in_gopkg_op_go_logging_v1",
    importpath = "gopkg.in/op/go-logging.v1",
    sum = "h1:6D+BvnJ/j6e222UW8s2qTSe3wGBtvo0MbVQG/c5k8RE=",
    version = "v1.0.0-20160211212156-b2cb9fa56473",
)

go_repository(
    name = "io_etcd_go_protodoc",
    importpath = "go.etcd.io/protodoc",
    sum = "h1:QQiUXlqz+d96jyNG71NE+IGTgOK6Xlhdx+PzvfbLHlQ=",
    version = "v0.0.0-20180829002748-484ab544e116",
)

go_repository(
    name = "io_k8s_apiserver",
    importpath = "k8s.io/apiserver",
    replace = "github.com/btwiuse/apiserver",
    sum = "h1:XZJvyNdzpVIMBJiyzKB91WF66HuCzI+SBbKHb7Jikvo=",
    version = "v0.20.4-btwiuse",
)

go_repository(
    name = "io_k8s_kubelet",
    importpath = "k8s.io/kubelet",
    replace = "k8s.io/kubelet",
    sum = "h1:yTD3mHQsoqOKWUAEYna6egOF8qzvdCeBcpTSS4lL6rw=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_sigs_apiserver_network_proxy_konnectivity_client",
    importpath = "sigs.k8s.io/apiserver-network-proxy/konnectivity-client",
    sum = "h1:KQOkVzXrLNb0EP6W0FD6u3CCPAwgXFYwZitbj7K0P0Y=",
    version = "v0.0.27",
)

go_repository(
    name = "io_k8s_sigs_metrics_server",
    importpath = "sigs.k8s.io/metrics-server",
    sum = "h1:XooqhTZu+HltKmPQYkzIarTvAbNvJOVhNLCqsZrznsI=",
    version = "v0.6.1",
)

go_repository(
    name = "org_golang_google_grpc_examples",
    importpath = "google.golang.org/grpc/examples",
    sum = "h1:EYPmxQ3D+iAYj/5Xjl35AYo3VGRY/JAzFA16vMqLb00=",
    version = "v0.0.0-20210223174733-dabedfb38b74",
)

go_repository(
    name = "org_mongodb_go_mongo_driver",
    importpath = "go.mongodb.org/mongo-driver",
    sum = "h1:rh7GdYmDrb8AQSkF8yteAus8qYOgOASWDOv1BWqBXkU=",
    version = "v1.4.6",
)

go_repository(
    name = "tools_gotest_v3",
    importpath = "gotest.tools/v3",
    sum = "h1:4AuOwCGf4lLR9u3YOe2awrHygurzhO/HeQ6laiA6Sx0=",
    version = "v3.0.3",
)

go_repository(
    name = "com_github_tetratelabs_wazero",
    importpath = "github.com/tetratelabs/wazero",
    sum = "h1:pygQVUHBelsX/QYiIwnWATfqeDobAbg4d6AVDX1vPMw=",
    version = "v0.0.0-20220430041858-abd1c79f3335",
)

go_repository(
    name = "com_github_coredns_caddy",
    importpath = "github.com/coredns/caddy",
    sum = "h1:ezvsPrT/tA/7pYDBZxu0cT0VmWk75AfIaf6GSYCNMf0=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_dnstap_golang_dnstap",
    importpath = "github.com/dnstap/golang-dnstap",
    sum = "h1:KRHBoURygdGtBjDI2w4HifJfMAhhOqDuktAokaSa234=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_farsightsec_golang_framestream",
    importpath = "github.com/farsightsec/golang-framestream",
    sum = "h1:/spFQHucTle/ZIPkYqrfshQqPe2VQEzesH243TjIwqA=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_opentracing",
    importpath = "github.com/grpc-ecosystem/grpc-opentracing",
    sum = "h1:MJG/KsmcqMwFAkh8mTnAwhyKoB+sTAnY4CACC110tbU=",
    version = "v0.0.0-20180507213350-8e809c8a8645",
)

go_repository(
    name = "com_github_coredns_coredns",
    importpath = "github.com/coredns/coredns",
    replace = "github.com/btwiuse/coredns",
    sum = "h1:z6g9H1MYQrM/WNns4RkttlqcS6u8MAOJGb9g12tHixg=",
    version = "v1.8.4",
)

go_repository(
    name = "com_github_armon_go_socks5",
    importpath = "github.com/armon/go-socks5",
    sum = "h1:0CwZNZbxp69SHPdPJAN/hZIm0C4OItdklCFmMRWYpio=",
    version = "v0.0.0-20160902184237-e75332964ef5",
)

go_repository(
    name = "com_github_chai2010_gettext_go",
    importpath = "github.com/chai2010/gettext-go",
    sum = "h1:7aWHqerlJ41y6FOsEUvknqgXnGmJyJSbjhAWq5pO4F8=",
    version = "v0.0.0-20160711120539-c6fed771bfd5",
)

go_repository(
    name = "com_github_containerd_containerd",
    importpath = "github.com/containerd/containerd",
    sum = "h1:rQyoYtj4KddB3bxG6SAqd4+08gePNyJjRqvOIfV3rkM=",
    version = "v1.5.7",
)

go_repository(
    name = "com_github_daviddengcn_go_colortext",
    importpath = "github.com/daviddengcn/go-colortext",
    sum = "h1:uVsMphB1eRx7xB1njzL3fuMdWRN8HtVzoUOItHMwv5c=",
    version = "v0.0.0-20160507010035-511bcaf42ccd",
)

go_repository(
    name = "com_github_docker_distribution",
    importpath = "github.com/docker/distribution",
    sum = "h1:a5mlkVzth6W5A4fOsS3D2EO5BUmsJpcB+cRlLU7cSug=",
    version = "v2.7.1+incompatible",
)

go_repository(
    name = "com_github_docker_go_connections",
    importpath = "github.com/docker/go-connections",
    sum = "h1:El9xVISelRB7BuFusrZozjnkIM5YnzCViNKohAFqRJQ=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_exponent_io_jsonpath",
    importpath = "github.com/exponent-io/jsonpath",
    sum = "h1:105gxyaGwCFad8crR9dcMQWvV9Hvulu6hwUh4tWPJnM=",
    version = "v0.0.0-20151013193312-d6023ce2651d",
)

go_repository(
    name = "com_github_fatih_camelcase",
    importpath = "github.com/fatih/camelcase",
    sum = "h1:hxNvNX/xYBp0ovncs8WyWZrOrpBNub/JfaMvbURyft8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_fvbommel_sortorder",
    importpath = "github.com/fvbommel/sortorder",
    sum = "h1:dSnXLt4mJYH25uDDGa3biZNQsozaUWDSWeKJ0qqFfzE=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_golangplus_bytes",
    importpath = "github.com/golangplus/bytes",
    sum = "h1:7xqw01UYS+KCI25bMrPxwNYkSns2Db1ziQPpVq99FpE=",
    version = "v0.0.0-20160111154220-45c989fe5450",
)

go_repository(
    name = "com_github_golangplus_fmt",
    importpath = "github.com/golangplus/fmt",
    sum = "h1:f5gsjBiF9tRRVomCvrkGMMWI8W1f2OBFar2c5oakAP0=",
    version = "v0.0.0-20150411045040-2a5d6d7d2995",
)

go_repository(
    name = "com_github_golangplus_testing",
    importpath = "github.com/golangplus/testing",
    sum = "h1:KhcknUwkWHKZPbFy2P7jH5LKJ3La+0ZeknkkmrSgqb0=",
    version = "v0.0.0-20180327235837-af21d9c3145e",
)

go_repository(
    name = "com_github_jaypipes_ghw",
    importpath = "github.com/jaypipes/ghw",
    sum = "h1:L4Bg16C1brIyEOo025bILWbtJmJMj7Nvv36nbnaTK3M=",
    version = "v0.8.1-0.20220131141055-fb0598ce62c8",
)

go_repository(
    name = "com_github_jaypipes_pcidb",
    importpath = "github.com/jaypipes/pcidb",
    sum = "h1:VIM7GKVaW4qba30cvB67xSCgJPTzkG8Kzw/cbs5PHWU=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_jpillora_chisel",
    importpath = "github.com/jpillora/chisel",
    sum = "h1:eLbzoX+ekDhVmF5CpSJD01NtH/w7QMYeaFCIFbzn9ns=",
    version = "v1.7.7",
)

go_repository(
    name = "com_github_koding_websocketproxy",
    importpath = "github.com/koding/websocketproxy",
    sum = "h1:N7A4JCA2G+j5fuFxCsJqjFU/sZe0mj8H0sSoSwbaikw=",
    version = "v0.0.0-20181220232114-7ed82d81a28c",
)

go_repository(
    name = "com_github_liggitt_tabwriter",
    importpath = "github.com/liggitt/tabwriter",
    sum = "h1:9TO3cAIGXtEhnIaL+V+BEER86oLrvS+kWobKpbJuye0=",
    version = "v0.0.0-20181228230101-89fcab3d43de",
)

go_repository(
    name = "com_github_lithammer_dedent",
    importpath = "github.com/lithammer/dedent",
    sum = "h1:VNzHMVCBNG1j0fh3OrsFRkVUwStdDArbgBWoPAffktY=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_makenowjust_heredoc",
    importpath = "github.com/MakeNowJust/heredoc",
    sum = "h1:sjQovDkwrZp8u+gxLtPgKGjk5hCxuy2hrRejBTA9xFU=",
    version = "v0.0.0-20170808103936-bb23615498cd",
)

go_repository(
    name = "com_github_mitchellh_go_wordwrap",
    importpath = "github.com/mitchellh/go-wordwrap",
    sum = "h1:6GlHJ/LTGMrIJbwgdqdl2eEH8o+Exx/0m8ir9Gns0u4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_morikuni_aec",
    importpath = "github.com/morikuni/aec",
    sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_opencontainers_go_digest",
    importpath = "github.com/opencontainers/go-digest",
    sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_opencontainers_image_spec",
    importpath = "github.com/opencontainers/image-spec",
    sum = "h1:JMemWkRwHx4Zj+fVxWoMCFm/8sYGGrUVojFA6h/TRcI=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_portainer_agent",
    importpath = "github.com/portainer/agent",
    sum = "h1:Obxde+U8uOrNDS5HMHk9EEESviQ69HnnW7GSMMrpa6c=",
    version = "v0.0.0-20220428235705-f483e4d7ecc4",
)

go_repository(
    name = "com_github_portainer_libcrypto",
    importpath = "github.com/portainer/libcrypto",
    sum = "h1:qY8TbocN75n5PDl16o0uVr5MevtM5IhdwSelXEd4nFM=",
    version = "v0.0.0-20210422035235-c652195c5c3a",
)

go_repository(
    name = "com_github_portainer_libhttp",
    importpath = "github.com/portainer/libhttp",
    sum = "h1:GMIjRVV2LADpJprPG2+8MdRH6XvrFgC7wHm7dFUdOpc=",
    version = "v0.0.0-20211208103139-07a5f798eb3f",
)

go_repository(
    name = "io_k8s_cli_runtime",
    importpath = "k8s.io/cli-runtime",
    replace = "k8s.io/cli-runtime",
    sum = "h1:jVU13lBeebHLtarHeHkoIi3uRONFzccmP7hHLzEoQ4w=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_component_helpers",
    importpath = "k8s.io/component-helpers",
    replace = "k8s.io/component-helpers",
    sum = "h1:3XJi6w+AcLd5f3ZwSRfgWuHFnUCmMAaRsUt2+NGDyQ0=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_kubectl",
    importpath = "k8s.io/kubectl",
    replace = "k8s.io/kubectl",
    sum = "h1:Y1gUiigiZM+ulcrnWeqSHlTd0/7xWcQIXjuMnjtHyoo=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_sigs_kustomize",
    importpath = "sigs.k8s.io/kustomize",
    sum = "h1:JUufWFNlI44MdtnjUqVnvh29rR37PQFzPbLXqhyOyX0=",
    version = "v2.0.3+incompatible",
)

go_repository(
    name = "com_github_btwiuse_cadvisor",
    importpath = "github.com/btwiuse/cadvisor",
    sum = "h1:TN7StIvTOGlPb0jSnp3UIUCXkXzD53UUGZL0ZhyGA+E=",
    version = "v0.0.0-20210312172035-34fddda41018",
)

go_repository(
    name = "com_github_checkpoint_restore_go_criu_v4",
    importpath = "github.com/checkpoint-restore/go-criu/v4",
    sum = "h1:WW2B2uxx9KWF6bGlHqhm8Okiafwwx7Y2kcpn8lCpjgo=",
    version = "v4.1.0",
)

go_repository(
    name = "com_github_cilium_ebpf",
    importpath = "github.com/cilium/ebpf",
    sum = "h1:1k/q3ATgxSXRdrmPfH8d7YK0GfqVsEKZAX9dQZvs56k=",
    version = "v0.7.0",
)

go_repository(
    name = "com_github_containerd_console",
    importpath = "github.com/containerd/console",
    sum = "h1:Pi6D+aZXM+oUw1czuKgH5IJ+y0jhYcwBJfx5/Ghn9dE=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_containerd_ttrpc",
    importpath = "github.com/containerd/ttrpc",
    sum = "h1:2/O3oTZN36q2xRolk0a2WWGgh7/Vf/liElg5hFYLX9U=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_containerd_typeurl",
    importpath = "github.com/containerd/typeurl",
    sum = "h1:Chlt8zIieDbzQFzXzAeBEF92KhExuE4p9p92/QmY7aY=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_cyphar_filepath_securejoin",
    importpath = "github.com/cyphar/filepath-securejoin",
    sum = "h1:jCwT2GTP+PY5nBz3c/YL5PAIbusElVrPujOBSCj8xRg=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_euank_go_kmsg_parser",
    importpath = "github.com/euank/go-kmsg-parser",
    sum = "h1:cHD53+PLQuuQyLZeriD1V/esuG4MuU0Pjs5y6iknohY=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_frankban_quicktest",
    importpath = "github.com/frankban/quicktest",
    sum = "h1:yNZif1OkDfNoDfb9zZa9aXIpejNR4F23Wely0c+Qdqk=",
    version = "v1.13.0",
)

go_repository(
    name = "com_github_garyburd_redigo",
    importpath = "github.com/garyburd/redigo",
    sum = "h1:yE/pwKCrbLpLpQICzYTeZ7JsTA/C53wFTJHaEtRqniM=",
    version = "v1.6.2",
)

go_repository(
    name = "com_github_google_cadvisor",
    importpath = "github.com/google/cadvisor",
    sum = "h1:jai6dmBP9QAYluNGqU18yVUTw6uuyAW0AqtZIjvl8Qg=",
    version = "v0.39.0",
)

go_repository(
    name = "com_github_gorilla_sessions",
    importpath = "github.com/gorilla/sessions",
    sum = "h1:DHd3rPN5lE3Ts3D8rKkQ8x/0kqfeNmBAaiSi+o7FsgI=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_influxdb_influxdb",
    importpath = "github.com/influxdb/influxdb",
    sum = "h1:+8XPwrWoNps4WbLfhNhH0ct8LUJAP1q+faViiEpSHYc=",
    version = "v0.9.6-0.20151125225445-9eab56311373",
)

go_repository(
    name = "com_github_jcmturner_aescts_v2",
    importpath = "github.com/jcmturner/aescts/v2",
    sum = "h1:9YKLH6ey7H4eDBXW8khjYslgyqG2xZikXP0EQFKrle8=",
    version = "v2.0.0",
)

go_repository(
    name = "com_github_jcmturner_dnsutils_v2",
    importpath = "github.com/jcmturner/dnsutils/v2",
    sum = "h1:lltnkeZGL0wILNvrNiVCR6Ro5PGU/SeBvVO/8c/iPbo=",
    version = "v2.0.0",
)

go_repository(
    name = "com_github_jcmturner_gofork",
    importpath = "github.com/jcmturner/gofork",
    sum = "h1:J7uCkflzTEhUZ64xqKnkDxq3kzc96ajM1Gli5ktUem8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jcmturner_goidentity_v6",
    importpath = "github.com/jcmturner/goidentity/v6",
    sum = "h1:VKnZd2oEIMorCTsFBnJWbExfNN7yZr3EhJAxwOkZg6o=",
    version = "v6.0.1",
)

go_repository(
    name = "com_github_jcmturner_gokrb5_v8",
    importpath = "github.com/jcmturner/gokrb5/v8",
    sum = "h1:6ZIM6b/JJN0X8UM43ZOM6Z4SJzla+a/u7scXFJzodkA=",
    version = "v8.4.2",
)

go_repository(
    name = "com_github_jcmturner_rpc_v2",
    importpath = "github.com/jcmturner/rpc/v2",
    sum = "h1:7FXXj8Ti1IaVFpSAziCZWNzbNuZmnvw/i6CqLNdWfZY=",
    version = "v2.0.3",
)

go_repository(
    name = "com_github_mesos_mesos_go",
    importpath = "github.com/mesos/mesos-go",
    sum = "h1:jMp9+W3zLu46g8EuP2su2Sjj7ipBh4N/g65c0kzGl/8=",
    version = "v0.0.11",
)

go_repository(
    name = "com_github_mindprince_gonvml",
    importpath = "github.com/mindprince/gonvml",
    sum = "h1:PS1dLCGtD8bb9RPKJrc8bS7qHL6JnW1CZvwzH9dPoUs=",
    version = "v0.0.0-20190828220739-9ebdce4bb989",
)

go_repository(
    name = "com_github_mistifyio_go_zfs",
    importpath = "github.com/mistifyio/go-zfs",
    sum = "h1:aKW/4cBs+yK6gpqU3K/oIwk9Q/XICqd3zOX/UFuvqmk=",
    version = "v2.1.2-0.20190413222219-f784269be439+incompatible",
)

go_repository(
    name = "com_github_moby_sys_mountinfo",
    importpath = "github.com/moby/sys/mountinfo",
    sum = "h1:1O+1cHA1aujwEwwVMa2Xm2l+gIpUHyd3+D+d7LZh1kM=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_mrunalp_fileutils",
    importpath = "github.com/mrunalp/fileutils",
    sum = "h1:NKzVxiH7eSk+OQ4M+ZYW1K6h27RUV3MI6NUTsHhU6Z4=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_opencontainers_runc",
    importpath = "github.com/opencontainers/runc",
    sum = "h1:opHZMaswlyxz1OuGpBE53Dwe4/xF7EZTY0A2L/FpCOg=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_opencontainers_runtime_spec",
    importpath = "github.com/opencontainers/runtime-spec",
    sum = "h1:3snG66yBm59tKhhSPQrQ/0bCrv1LQbKt40LnUPiUxdc=",
    version = "v1.0.3-0.20210326190908-1c3f411f0417",
)

go_repository(
    name = "com_github_opencontainers_selinux",
    importpath = "github.com/opencontainers/selinux",
    sum = "h1:c4ca10UMgRcvZ6h0K4HtS15UaVSBEaE+iln2LVpAuGc=",
    version = "v1.8.2",
)

go_repository(
    name = "com_github_pquerna_ffjson",
    importpath = "github.com/pquerna/ffjson",
    sum = "h1:xoIK0ctDddBMnc74udxJYBqlo9Ylnsp1waqjLsnef20=",
    version = "v0.0.0-20190930134022-aa0246cd15f7",
)

go_repository(
    name = "com_github_rican7_retry",
    importpath = "github.com/Rican7/retry",
    sum = "h1:FqK94z34ly8Baa6K+G8Mmza9rYWTKOJk+yckIBB5qVk=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_seandolphin_bqschema",
    importpath = "github.com/SeanDolphin/bqschema",
    sum = "h1:iCYFd5Qsw6caM2k5/SsITSL9+3kQCr+oz6pnNjWTq90=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_seccomp_libseccomp_golang",
    importpath = "github.com/seccomp/libseccomp-golang",
    sum = "h1:NJjM5DNFOs0s3kYE1WUOr6G8V97sdt46rlXTMfXGWBo=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_syndtr_gocapability",
    importpath = "github.com/syndtr/gocapability",
    sum = "h1:kdXcSzyDtseVEc4yCz2qF8ZrQvIDBJLl4S1c3GCXmoI=",
    version = "v0.0.0-20200815063812-42c35b437635",
)

go_repository(
    name = "com_github_vishvananda_netlink",
    importpath = "github.com/vishvananda/netlink",
    sum = "h1:cPXZWzzG0NllBLdjWoD1nDfaqu98YMv+OneaKc8sPOA=",
    version = "v1.1.1-0.20201029203352-d40f9887b852",
)

go_repository(
    name = "com_github_vishvananda_netns",
    importpath = "github.com/vishvananda/netns",
    sum = "h1:gga7acRE695APm9hlsSMoOoE65U4/TcqNj90mc69Rlg=",
    version = "v0.0.0-20211101163701-50045581ed74",
)

go_repository(
    name = "com_github_willf_bitset",
    importpath = "github.com/willf/bitset",
    sum = "h1:N7Z7E9UvjW+sGsEl7k/SJrvY2reP1A07MrGuCjIOjRE=",
    version = "v1.1.11",
)

go_repository(
    name = "in_gopkg_olivere_elastic_v2",
    importpath = "gopkg.in/olivere/elastic.v2",
    sum = "h1:7cpl3MW8ysa4GYFBXklpo5mspe4NK0rpZTdyZ+QcD4U=",
    version = "v2.0.61",
)

go_repository(
    name = "com_github_ajstarks_svgo",
    importpath = "github.com/ajstarks/svgo",
    sum = "h1:wVe6/Ea46ZMeNkQjjBW6xcqyQA/j5e0D6GytH95g0gQ=",
    version = "v0.0.0-20180226025133-644b8db467af",
)

go_repository(
    name = "com_github_auth0_go_jwt_middleware",
    importpath = "github.com/auth0/go-jwt-middleware",
    sum = "h1:irR1cO6eek3n5uquIVaRAsQmZnlsfPuHNz31cXo4eyk=",
    version = "v0.0.0-20170425171159-5493cabe49f7",
)

go_repository(
    name = "com_github_boltdb_bolt",
    importpath = "github.com/boltdb/bolt",
    sum = "h1:JQmyP4ZBrce+ZQu0dY660FMfatumYDLun9hBCUVIkF4=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_clusterhq_flocker_go",
    importpath = "github.com/clusterhq/flocker-go",
    sum = "h1:eIHD9GNM3Hp7kcRW5mvcz7WTR3ETeoYYKwpgA04kaXE=",
    version = "v0.0.0-20160920122132-2b8b7259d313",
)

go_repository(
    name = "com_github_codegangsta_negroni",
    importpath = "github.com/codegangsta/negroni",
    sum = "h1:+aYywywx4bnKXWvoWtRfJ91vC59NbEhEY03sZjQhbVY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_container_storage_interface_spec",
    importpath = "github.com/container-storage-interface/spec",
    sum = "h1:bD9KIVgaVKKkQ/UbVUY9kCaH/CJbhNxe0eeB4JeJV2s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_containerd_cgroups",
    importpath = "github.com/containerd/cgroups",
    sum = "h1:iJnMvco9XGvKUvNQkv88bE4uJXxRQH18efbKo9w5vHQ=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_containerd_continuity",
    importpath = "github.com/containerd/continuity",
    sum = "h1:UFRRY5JemiAhPZrr/uE0n8fMTLcZsUvySPr1+D7pgr8=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_containerd_fifo",
    importpath = "github.com/containerd/fifo",
    sum = "h1:6PirWBr9/L7GDamKr+XM0IeUFXu5mf3M/BPpH9gaLBU=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_containerd_go_runc",
    importpath = "github.com/containerd/go-runc",
    sum = "h1:oU+lLv1ULm5taqgV/CJivypVODI4SUz1znWjv3nNYS0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_containernetworking_cni",
    importpath = "github.com/containernetworking/cni",
    sum = "h1:7zpDnQ3T3s4ucOuJ/ZCLrYBxzkg0AELFfII3Epo9TmI=",
    version = "v0.8.1",
)

go_repository(
    name = "com_github_coredns_corefile_migration",
    importpath = "github.com/coredns/corefile-migration",
    sum = "h1:7HI4r5S5Fne749a+JDxUZppqBpYoZK8Q53ZVK9cn3aM=",
    version = "v1.0.10",
)

go_repository(
    name = "com_github_fogleman_gg",
    importpath = "github.com/fogleman/gg",
    sum = "h1:WXb3TSNmHp2vHoCroCIB1foO/yQ36swABL8aOVeDpgg=",
    version = "v1.2.1-0.20190220221249-0403632d5b90",
)

go_repository(
    name = "com_github_go_bindata_go_bindata",
    importpath = "github.com/go-bindata/go-bindata",
    sum = "h1:tR4f0e4VTO7LK6B2YWyAoVEzG9ByG1wrXB4TL9+jiYg=",
    version = "v3.1.1+incompatible",
)

go_repository(
    name = "com_github_go_ozzo_ozzo_validation",
    importpath = "github.com/go-ozzo/ozzo-validation",
    sum = "h1:sUy/in/P6askYr16XJgTKq/0SZhiWsdg4WZGaLsGQkM=",
    version = "v3.5.0+incompatible",
)

go_repository(
    name = "com_github_golang_freetype",
    importpath = "github.com/golang/freetype",
    sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
    version = "v0.0.0-20170609003504-e2365dfdc4a0",
)

go_repository(
    name = "com_github_googlecloudplatform_k8s_cloud_provider",
    importpath = "github.com/GoogleCloudPlatform/k8s-cloud-provider",
    sum = "h1:JhyuWIqYrstW7KHMjk/fTqU0xtMpBOHuiTA2FVc7L4E=",
    version = "v0.0.0-20200415212048-7901bc822317",
)

go_repository(
    name = "com_github_heketi_heketi",
    importpath = "github.com/heketi/heketi",
    sum = "h1:ysqc8k973k1lLJ4BOOHAkx14K2nt4cLjsIm+hwWDZDE=",
    version = "v9.0.1-0.20190917153846-c2e2a4ab7ab9+incompatible",
)

go_repository(
    name = "com_github_heketi_tests",
    importpath = "github.com/heketi/tests",
    sum = "h1:oJ/NLadJn5HoxvonA6VxG31lg0d6XOURNA09BTtM4fY=",
    version = "v0.0.0-20151005000721-f3775cbcefd6",
)

go_repository(
    name = "com_github_ishidawataru_sctp",
    importpath = "github.com/ishidawataru/sctp",
    sum = "h1:qPmlgoeRS18y2dT+iAH5vEKZgIqgiPi2Y8UCu/b7Aq8=",
    version = "v0.0.0-20190723014705-7c296d48a2b5",
)

go_repository(
    name = "com_github_jeffashton_win_pdh",
    importpath = "github.com/JeffAshton/win_pdh",
    sum = "h1:UKkYhof1njT1/xq4SEg5z+VpTgjmNeHwPGRQl7takDI=",
    version = "v0.0.0-20161109143554-76bb4ee9f0ab",
)

go_repository(
    name = "com_github_jung_kurt_gofpdf",
    importpath = "github.com/jung-kurt/gofpdf",
    sum = "h1:PJr+ZMXIecYc1Ey2zucXdR73SMBtgjPgwa31099IMv0=",
    version = "v1.0.3-0.20190309125859-24315acbbda5",
)

go_repository(
    name = "com_github_libopenstorage_openstorage",
    importpath = "github.com/libopenstorage/openstorage",
    sum = "h1:GLPam7/0mpdP8ZZtKjbfcXJBTIA/T1O6CBErVEFEyIM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_lpabon_godbc",
    importpath = "github.com/lpabon/godbc",
    sum = "h1:ilqjArN1UOENJJdM34I2YHKmF/B0gGq4VLoSGy9iAao=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_microsoft_hcsshim",
    importpath = "github.com/Microsoft/hcsshim",
    sum = "h1:btRfUDThBE5IKcvI8O8jOiIkujUsAMBSRsYDYmEi6oM=",
    version = "v0.8.21",
)

go_repository(
    name = "com_github_moby_ipvs",
    importpath = "github.com/moby/ipvs",
    sum = "h1:aoZ7fhLTXgDbzVrAnvV+XbKOU8kOET7B3+xULDF/1o0=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_mvdan_xurls",
    importpath = "github.com/mvdan/xurls",
    sum = "h1:OpuDelGQ1R1ueQ6sSryzi6P+1RtBpfQHM8fJwlE45ww=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_quobyte_api",
    importpath = "github.com/quobyte/api",
    sum = "h1:+sOX1gIlC/OaLipqVZWrHgly9Kh9Qo8OygeS0mWAg30=",
    version = "v0.1.8",
)

go_repository(
    name = "com_github_remyoudompheng_bigfft",
    importpath = "github.com/remyoudompheng/bigfft",
    sum = "h1:/NRJ5vAYoqz+7sG51ubIDHXeWO8DlTSrToPu6q11ziA=",
    version = "v0.0.0-20170806203942-52369c62f446",
)

go_repository(
    name = "com_github_robfig_cron",
    importpath = "github.com/robfig/cron",
    sum = "h1:jk4/Hud3TTdcrJgUOBgsqrZBarcxl6ADIjSC2iniwLY=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_rubiojr_go_vhd",
    importpath = "github.com/rubiojr/go-vhd",
    sum = "h1:if3/24+h9Sq6eDx8UUz1SO9cT9tizyIsATfB7b4D3tc=",
    version = "v0.0.0-20200706105327-02e210299021",
)

go_repository(
    name = "com_github_storageos_go_api",
    importpath = "github.com/storageos/go-api",
    sum = "h1:U0SablXoZIg06gvSlg8BCdzq1C/SkHVygOVX95Z2MU0=",
    version = "v2.2.0+incompatible",
)

go_repository(
    name = "com_github_thecodeteam_goscaleio",
    importpath = "github.com/thecodeteam/goscaleio",
    sum = "h1:SB5tO98lawC+UK8ds/U2jyfOCH7GTcFztcF5x9gbut4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_vmware_govmomi",
    importpath = "github.com/vmware/govmomi",
    sum = "h1:gpw/0Ku+6RgF3jsi7fnCLmlcikBHfKBCUcu1qgc16OU=",
    version = "v0.20.3",
)

go_repository(
    name = "io_k8s_apiextensions_apiserver",
    importpath = "k8s.io/apiextensions-apiserver",
    replace = "k8s.io/apiextensions-apiserver",
    sum = "h1:VO/Y5PwBdznMIctX/vvgSNhxffikEmcLC/V1bpbhHhU=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_cloud_provider",
    importpath = "k8s.io/cloud-provider",
    replace = "k8s.io/cloud-provider",
    sum = "h1:wy5PE9CM3V7tTJ1rqAP8c3BIy/uZZoMdGijCOxvPOeo=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_cluster_bootstrap",
    importpath = "k8s.io/cluster-bootstrap",
    replace = "k8s.io/cluster-bootstrap",
    sum = "h1:F9NSfGzWyl8Nhh8xU6p5No2S0RThysxYsFr2EuJZBxo=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_controller_manager",
    importpath = "k8s.io/controller-manager",
    replace = "k8s.io/controller-manager",
    sum = "h1:MwaI7deb9NhqyxaWTcOeocDGWXy0U65hZH+qGA8esSA=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_cri_api",
    importpath = "k8s.io/cri-api",
    replace = "k8s.io/cri-api",
    sum = "h1:AwwzhJMfaxiw8NnEJAUQI+FWlX1mAp9tHODTVxnkEQg=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_csi_translation_lib",
    importpath = "k8s.io/csi-translation-lib",
    replace = "k8s.io/csi-translation-lib",
    sum = "h1:I5qDIDPlVxcxfoQuZ1EoRkR0YtOTwtmLVtDFuafW1Go=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_heapster",
    importpath = "k8s.io/heapster",
    sum = "h1:lUsE/AHOMHpi3MLlBEkaU8Esxm5QhdyCrv1o7ot0s84=",
    version = "v1.2.0-beta.1",
)

go_repository(
    name = "io_k8s_kube_aggregator",
    importpath = "k8s.io/kube-aggregator",
    replace = "k8s.io/kube-aggregator",
    sum = "h1:j/SUwPy1eO+ud3XOUGmH18gISPyerqhXOoNRZDbv3fs=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_kube_controller_manager",
    importpath = "k8s.io/kube-controller-manager",
    replace = "k8s.io/kube-controller-manager",
    sum = "h1:vyZOuF0zU9NRku3snWleE3KzL5Hzy2kD7zJ2tyC8IWo=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_kube_proxy",
    importpath = "k8s.io/kube-proxy",
    replace = "k8s.io/kube-proxy",
    sum = "h1:bKRlKwhOQlMtGCnBJhgWLdNPGvYW6+1yBmG/sj6OkgI=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_kube_scheduler",
    importpath = "k8s.io/kube-scheduler",
    replace = "k8s.io/kube-scheduler",
    sum = "h1:aU/B/mP7cN75dOFUrvZSEnDa3Rp0wvNd8Fy+WlEihlk=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_kubernetes",
    importpath = "k8s.io/kubernetes",
    sum = "h1:qTfB+u5M92k2fCCCVP2iuhgwwSOv1EkAkvQY1tQODD8=",
    version = "v1.13.0",
)

go_repository(
    name = "io_k8s_legacy_cloud_providers",
    importpath = "k8s.io/legacy-cloud-providers",
    replace = "k8s.io/legacy-cloud-providers",
    sum = "h1:E2Twm/y3dYXlUqTFad38IJGpoimH4fExSQtoaLmC5Dw=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_mount_utils",
    importpath = "k8s.io/mount-utils",
    replace = "k8s.io/mount-utils",
    sum = "h1:jfrZDJ9u5Jq7n3lJ9UKpp/RsV6M4sCLi0rmM6jiKoQQ=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_sample_apiserver",
    importpath = "k8s.io/sample-apiserver",
    replace = "k8s.io/sample-apiserver",
    sum = "h1:rAfusAeMxAoDwkvCob7b95inhRSUoHOkXz+56H8Nhzs=",
    version = "v0.20.4",
)

go_repository(
    name = "io_k8s_system_validators",
    importpath = "k8s.io/system-validators",
    sum = "h1:D8clGm3PIsgX0bHoQWQXC8W/rTi6jRKejQGk7z/yuDY=",
    version = "v1.2.0",
)

go_repository(
    name = "org_bitbucket_bertimus9_systemstat",
    importpath = "bitbucket.org/bertimus9/systemstat",
    sum = "h1:N9r8OBSXAgEUfho3SQtZLY8zo6E1OdOMvelvP22aVFc=",
    version = "v0.0.0-20180207000608-0eeff89b0690",
)

go_repository(
    name = "org_gonum_v1_gonum",
    importpath = "gonum.org/v1/gonum",
    sum = "h1:4r+yNT0+8SWcOkXP+63H2zQbN+USnC73cjGUxnDF94Q=",
    version = "v0.6.2",
)

go_repository(
    name = "org_gonum_v1_netlib",
    importpath = "gonum.org/v1/netlib",
    sum = "h1:jRyg0XfpwWlhEV8mDfdNGBeSJM2fuyh9Yjrnd8kF2Ts=",
    version = "v0.0.0-20190331212654-76723241ea4e",
)

go_repository(
    name = "org_gonum_v1_plot",
    importpath = "gonum.org/v1/plot",
    sum = "h1:Qh4dB5D/WpoUUp3lSod7qgoyEHbDGPUWjIbnqdqqe1k=",
    version = "v0.0.0-20190515093506-e2840ee46a6b",
)

go_repository(
    name = "org_modernc_cc",
    importpath = "modernc.org/cc",
    sum = "h1:nPibNuDEx6tvYrUAtvDTTw98rx5juGsa5zuDnKwEEQQ=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_golex",
    importpath = "modernc.org/golex",
    sum = "h1:wWpDlbK8ejRfSyi0frMyhilD3JBvtcx2AdGDnU+JtsE=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_mathutil",
    importpath = "modernc.org/mathutil",
    sum = "h1:93vKjrJopTPrtTNpZ8XIovER7iCIH1QU7wNbOQXC60I=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_strutil",
    importpath = "modernc.org/strutil",
    sum = "h1:XVFtQwFVwc02Wk+0L/Z/zDDXO81r5Lhe6iMKmGX3KhE=",
    version = "v1.0.0",
)

go_repository(
    name = "org_modernc_xc",
    importpath = "modernc.org/xc",
    sum = "h1:7ccXrupWZIS3twbUGrtKmHS2DXY6xegFua+6O3xgAFU=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_dgryski_go_metro",
    importpath = "github.com/dgryski/go-metro",
    sum = "h1:BS21ZUJ/B5X2UVUbczfmdWH7GapPWAhxcMsDnjJTU1E=",
    version = "v0.0.0-20200812162917-85c65e2d0165",
)

go_repository(
    name = "com_github_ebfe_bcrypt_pbkdf",
    importpath = "github.com/ebfe/bcrypt_pbkdf",
    sum = "h1:YtdtTUN1iH97s+6PUjLnaiKSQj4oG1/EZ3N9bx6g4kU=",
    version = "v0.0.0-20140212075826-3c8d2dcb253a",
)

go_repository(
    name = "com_github_h12w_go_socks5",
    importpath = "github.com/h12w/go-socks5",
    sum = "h1:5XxdakFhqd9dnXoAZy1Mb2R/DZ6D1e+0bGC/JhucGYI=",
    version = "v0.0.0-20200522160539-76189e178364",
)

go_repository(
    name = "com_github_phayes_freeport",
    importpath = "github.com/phayes/freeport",
    sum = "h1:JhzVVoYvbOACxoUmOs6V/G4D5nPVUW73rKvXxP4XUJc=",
    version = "v0.0.0-20180830031419-95f893ade6f2",
)

go_repository(
    name = "com_github_pires_go_proxyproto",
    importpath = "github.com/pires/go-proxyproto",
    sum = "h1:EBupykFmo22SDjv4fQVQd2J9NOoLPmyZA/15ldOGkPw=",
    version = "v0.6.1",
)

go_repository(
    name = "com_github_seiflotfy_cuckoofilter",
    importpath = "github.com/seiflotfy/cuckoofilter",
    sum = "h1:pqy40B3MQWYrza7YZXOXgl0Nf0QGFqrOC0BKae1UNAA=",
    version = "v0.0.0-20201222105146-bc6005554a0c",
)

go_repository(
    name = "com_github_v2fly_v2ray_core_v4",
    importpath = "github.com/v2fly/v2ray-core/v4",
    sum = "h1:RsyjltU81bR+LRIzsRAGZV+QslPYPXOl5YPVCU+YGNc=",
    version = "v4.45.0",
)

go_repository(
    name = "com_github_v2fly_vsign",
    importpath = "github.com/v2fly/VSign",
    sum = "h1:p1UzXK6VAutXFFQMnre66h7g1BjRKUnLv0HfmmRoz7w=",
    version = "v0.0.0-20201108000810-e2adc24bf848",
)

go_repository(
    name = "ag_pack_amqp",
    importpath = "pack.ag/amqp",
    sum = "h1:cuNDWLUTbKRtEZwhB0WQBXf9pGbm87pUBXQhvcFxBWg=",
    version = "v0.11.2",
)

go_repository(
    name = "com_github_alecthomas_kingpin",
    importpath = "github.com/alecthomas/kingpin",
    sum = "h1:5svnBTFgJjZvGKyYBtMB0+m5wvrbUHiqye8wRJMlnYI=",
    version = "v2.2.6+incompatible",
)

go_repository(
    name = "com_github_apex_log",
    importpath = "github.com/apex/log",
    sum = "h1:1fyfbPvUwD10nMoh3hY6MXzvZShJQn9/ck7ATgAt5pA=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_apex_logs",
    importpath = "github.com/apex/logs",
    sum = "h1:KmEBVwfDUOTFcBO8cfkJYwdQ5487UZSN+GteOGPmiro=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_aphistic_golf",
    importpath = "github.com/aphistic/golf",
    sum = "h1:2KLQMJ8msqoPHIPDufkxVcoTtcmE5+1sL9950m4R9Pk=",
    version = "v0.0.0-20180712155816-02c07f170c5a",
)

go_repository(
    name = "com_github_aphistic_sweet",
    importpath = "github.com/aphistic/sweet",
    sum = "h1:I4z+fAUqvKfvZV/CHi5dV0QuwbmIvYYFDjG0Ss5QpAs=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_aybabtme_rgbterm",
    importpath = "github.com/aybabtme/rgbterm",
    sum = "h1:WWB576BN5zNSZc/M9d/10pqEx5VHNhaQ/yOVAkmj5Yo=",
    version = "v0.0.0-20170906152045-cc83f3b3ce59",
)

go_repository(
    name = "com_github_azure_azure_amqp_common_go_v2",
    importpath = "github.com/Azure/azure-amqp-common-go/v2",
    sum = "h1:+QbFgmWCnPzdaRMfsI0Yb6GrRdBj5jVL8N3EXuEUcBQ=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_azure_azure_pipeline_go",
    importpath = "github.com/Azure/azure-pipeline-go",
    sum = "h1:6oiIS9yaG6XCCzhgAgKFfIWyo4LLCiDhZot6ltoThhY=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_azure_azure_service_bus_go",
    importpath = "github.com/Azure/azure-service-bus-go",
    sum = "h1:G1qBLQvHCFDv9pcpgwgFkspzvnGknJRR0PYJ9ytY/JA=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_azure_azure_storage_blob_go",
    importpath = "github.com/Azure/azure-storage-blob-go",
    sum = "h1:53qhf0Oxa0nOjgbDeeYPUeyiNmafAFEY95rZLK0Tj6o=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_blakesmith_ar",
    importpath = "github.com/blakesmith/ar",
    sum = "h1:m935MPodAbYS46DG4pJSv7WO+VECIWUQ7OJYSoTrMh4=",
    version = "v0.0.0-20190502131153-809d4375e1fb",
)

go_repository(
    name = "com_github_bombsimon_wsl_v3",
    importpath = "github.com/bombsimon/wsl/v3",
    sum = "h1:E5SRssoBgtVFPcYWUOFJEcgaySgdtTNYzsSKDOY7ss8=",
    version = "v3.1.0",
)

go_repository(
    name = "com_github_btwiuse_dkg",
    importpath = "github.com/btwiuse/dkg",
    sum = "h1:1tIJhP64piTclCgtmtZNwcX2WKCgQsPXV6jENiu/h0A=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_caarlos0_ctrlc",
    importpath = "github.com/caarlos0/ctrlc",
    sum = "h1:2DtF8GSIcajgffDFJzyG15vO+1PuBWOMUdFut7NnXhw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_campoy_unique",
    importpath = "github.com/campoy/unique",
    sum = "h1:V9a67dfYqPLAvzk5hMQOXYJlZ4SLIXgyKIE+ZiHzgGQ=",
    version = "v0.0.0-20180121183637-88950e537e7e",
)

go_repository(
    name = "com_github_cavaliercoder_go_cpio",
    importpath = "github.com/cavaliercoder/go-cpio",
    sum = "h1:hHg27A0RSSp2Om9lubZpiMgVbvn39bsUmW9U5h0twqc=",
    version = "v0.0.0-20180626203310-925f9528c45e",
)

go_repository(
    name = "com_github_cppforlife_cobrautil",
    importpath = "github.com/cppforlife/cobrautil",
    sum = "h1:rPWcUBgMb1ox2eCohCuZ8gsZVe0aB5qBbYaBpdoxfCE=",
    version = "v0.0.0-20200514214827-bb86e6965d72",
)

go_repository(
    name = "com_github_cppforlife_color",
    importpath = "github.com/cppforlife/color",
    sum = "h1:mYQweUIBD+TBRjIeQnJmXr0GSVMpI6O0takyb/aaOgo=",
    version = "v1.9.1-0.20200716202919-6706ac40b835",
)

go_repository(
    name = "com_github_cppforlife_go_cli_ui",
    importpath = "github.com/cppforlife/go-cli-ui",
    sum = "h1:iJlI52aykDztyxFnESp3BVwHJi9Pu8Z9s1Ftb5mb4j0=",
    version = "v0.0.0-20220428182907-73db60c7611a",
)

go_repository(
    name = "com_github_devigned_tab",
    importpath = "github.com/devigned/tab",
    sum = "h1:3mD6Kb1mUOYeLpJvTVSDwSg5ZsfSxfvxGRTxRsJsITA=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_djarvur_go_err113",
    importpath = "github.com/Djarvur/go-err113",
    sum = "h1:uCRZZOdMQ0TZPHYTdYpoC0bLYJKPEHPUJ8MeAa51lNU=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_docker_cli",
    importpath = "github.com/docker/cli",
    sum = "h1:OJ7YkwQA+k2Oi51lmCojpjiygKpi76P7bg91b2eJxYU=",
    version = "v20.10.9+incompatible",
)

go_repository(
    name = "com_github_docker_docker_credential_helpers",
    importpath = "github.com/docker/docker-credential-helpers",
    sum = "h1:zI2p9+1NQYdnG6sMU26EX4aVGlqbInSQxQXLvzJ4RPQ=",
    version = "v0.6.3",
)

go_repository(
    name = "com_github_go_xmlfmt_xmlfmt",
    importpath = "github.com/go-xmlfmt/xmlfmt",
    sum = "h1:khEcpUM4yFcxg4/FHQWkvVRmgijNXRfzkIDHh23ggEo=",
    version = "v0.0.0-20191208150333-d5b6f63a941b",
)

go_repository(
    name = "com_github_google_go_containerregistry",
    importpath = "github.com/google/go-containerregistry",
    sum = "h1:AG8FSAfXglim2l5qSrqp5VK2Xl03PiBf25NiTGGamws=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_google_go_github_v28",
    importpath = "github.com/google/go-github/v28",
    sum = "h1:kORf5ekX5qwXO2mGzXXOjMe/g6ap8ahVe0sBEulhSxo=",
    version = "v28.1.1",
)

go_repository(
    name = "com_github_google_go_replayers_grpcreplay",
    importpath = "github.com/google/go-replayers/grpcreplay",
    sum = "h1:eNb1y9rZFmY4ax45uEEECSa8fsxGRU+8Bil52ASAwic=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_google_go_replayers_httpreplay",
    importpath = "github.com/google/go-replayers/httpreplay",
    sum = "h1:AX7FUb4BjrrzNvblr/OlgwrmFiep6soj5K2QSDW7BGk=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_google_rpmpack",
    importpath = "github.com/google/rpmpack",
    sum = "h1:BW6OvS3kpT5UEPbCZ+KyX/OB4Ks9/MNMhWjqPPkZxsE=",
    version = "v0.0.0-20191226140753-aa36bfddb3a0",
)

go_repository(
    name = "com_github_google_subcommands",
    importpath = "github.com/google/subcommands",
    sum = "h1:/eqq+otEXm5vhfBrbREPCSVQbvofip6kIz+mX5TUH7k=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_google_wire",
    importpath = "github.com/google/wire",
    sum = "h1:kXcsA/rIGzJImVqPdhfnr6q0xsS9gU0515q1EPpJ9fE=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_googlecloudplatform_cloudsql_proxy",
    importpath = "github.com/GoogleCloudPlatform/cloudsql-proxy",
    sum = "h1:sTOp2Ajiew5XIH92YSdwhYc+bgpUX5j5TKK/Ac8Saw8=",
    version = "v0.0.0-20191009163259-e802c2cb94ae",
)

go_repository(
    name = "com_github_gookit_color",
    importpath = "github.com/gookit/color",
    sum = "h1:xOYBan3Fwlrqj1M1UN2TlHOCRiek3bGzWf/vPnJ1roE=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_goreleaser_goreleaser",
    importpath = "github.com/goreleaser/goreleaser",
    sum = "h1:Z+7XPrfGK11s/Sp+a06sx2FzGuCjTBdxN2ubpGvQbjY=",
    version = "v0.136.0",
)

go_repository(
    name = "com_github_goreleaser_nfpm",
    importpath = "github.com/goreleaser/nfpm",
    sum = "h1:BPwIomC+e+yuDX9poJowzV7JFVcYA0+LwGSkbAPs2Hw=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_hashicorp_go_hclog",
    importpath = "github.com/hashicorp/go-hclog",
    sum = "h1:K4ev2ib4LdQETX5cSZBG0DVLk1jwGqSPXBjdah3veNs=",
    version = "v0.16.2",
)

go_repository(
    name = "com_github_hashicorp_go_retryablehttp",
    importpath = "github.com/hashicorp/go-retryablehttp",
    sum = "h1:eu1EI/mbirUgP5C8hVsTNaGZreBDlYiwC1FZWkvQPQ4=",
    version = "v0.7.0",
)

go_repository(
    name = "com_github_jarcoal_httpmock",
    importpath = "github.com/jarcoal/httpmock",
    sum = "h1:cHtVEcTxRSX4J0je7mWPfc9BpDpqzXSJ5HbymZmyHck=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_jingyugao_rowserrcheck",
    importpath = "github.com/jingyugao/rowserrcheck",
    sum = "h1:GmsqmapfzSJkm28dhRoHz2tLRbJmqhU86IPgBtN3mmk=",
    version = "v0.0.0-20191204022205-72ab7603b68a",
)

go_repository(
    name = "com_github_jirfag_go_printf_func_name",
    importpath = "github.com/jirfag/go-printf-func-name",
    sum = "h1:KA9BjwUk7KlCh6S9EAGWBt1oExIUv9WyNCiRz5amv48=",
    version = "v0.0.0-20200119135958-7558a9eaa5af",
)

go_repository(
    name = "com_github_jmoiron_sqlx",
    importpath = "github.com/jmoiron/sqlx",
    sum = "h1:lrdPtrORjGv1HbbEvKWDUAy97mPpFm4B8hp77tcCUJY=",
    version = "v1.2.1-0.20190826204134-d7d95172beb5",
)

go_repository(
    name = "com_github_joefitzgerald_rainbow_reporter",
    importpath = "github.com/joefitzgerald/rainbow-reporter",
    sum = "h1:AuMG652zjdzI0YCCnXAqATtRBpGXMcAnrajcaTrSeuo=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_maratori_testpackage",
    importpath = "github.com/maratori/testpackage",
    sum = "h1:QtJ5ZjqapShm0w5DosRjg0PRlSdAdlx+W6cCKoALdbQ=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_mattn_go_ieproxy",
    importpath = "github.com/mattn/go-ieproxy",
    sum = "h1:qiyop7gCflfhwCzGyeT0gro3sF9AIg9HU98JORTkqfI=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_maxbrunsfeld_counterfeiter_v6",
    importpath = "github.com/maxbrunsfeld/counterfeiter/v6",
    sum = "h1:g+4J5sZg6osfvEfkRZxJ1em0VT95/UOZgi/l7zi1/oE=",
    version = "v6.2.2",
)

go_repository(
    name = "com_github_mgutz_ansi",
    importpath = "github.com/mgutz/ansi",
    sum = "h1:5PJl274Y63IEHC+7izoQE9x6ikvDFZS2mDVS3drnohI=",
    version = "v0.0.0-20200706080929-d51e80ef957d",
)

go_repository(
    name = "com_github_nakabonne_nestif",
    importpath = "github.com/nakabonne/nestif",
    sum = "h1:+yOViDGhg8ygGrmII72nV9B/zGxY188TYpfolntsaPw=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_phayes_checkstyle",
    importpath = "github.com/phayes/checkstyle",
    sum = "h1:CdDQnGF8Nq9ocOS/xlSptM1N3BbrA6/kmaep5ggwaIA=",
    version = "v0.0.0-20170904204023-bfd46e6a821d",
)

go_repository(
    name = "com_github_quasilyte_go_ruleguard",
    importpath = "github.com/quasilyte/go-ruleguard",
    sum = "h1:DvnesvLtRPQOvaUbfXfh0tpMHg29by0H7F2U+QIkSu8=",
    version = "v0.1.2-0.20200318202121-b00d7a75d3d8",
)

go_repository(
    name = "com_github_ryancurrah_gomodguard",
    importpath = "github.com/ryancurrah/gomodguard",
    sum = "h1:DWbye9KyMgytn8uYpuHkwf0RHqAYO6Ay/D0TbCpPtVU=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_sassoftware_go_rpmutils",
    importpath = "github.com/sassoftware/go-rpmutils",
    sum = "h1:+gCnWOZV8Z/8jehJ2CdqB47Z3S+SREmQcuXkRFLNsiI=",
    version = "v0.0.0-20190420191620-a8f1baeba37b",
)

go_repository(
    name = "com_github_sclevine_spec",
    importpath = "github.com/sclevine/spec",
    sum = "h1:1Jwdf9jSfDl9NVmt8ndHqbTZ7XCCPbh1jI3hkDBHVYA=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_securego_gosec_v2",
    importpath = "github.com/securego/gosec/v2",
    sum = "h1:y/9mCF2WPDbSDpL3QDWZD3HHGrSYw0QSHnCqTfs4JPE=",
    version = "v2.3.0",
)

go_repository(
    name = "com_github_smartystreets_go_aws_auth",
    importpath = "github.com/smartystreets/go-aws-auth",
    sum = "h1:hp2CYQUINdZMHdvTdXtPOY2ainKl4IoMcpAXEf2xj3Q=",
    version = "v0.0.0-20180515143844-0c1422d1fdb9",
)

go_repository(
    name = "com_github_smartystreets_gunit",
    importpath = "github.com/smartystreets/gunit",
    sum = "h1:RyPDUFcJbvtXlhJPk7v+wnxZRY2EUokhEYl2EJOPToI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_tdakkota_asciicheck",
    importpath = "github.com/tdakkota/asciicheck",
    sum = "h1:HxLVTlqcHhFAz3nWUcuvpH7WuOMv8LQoCWmruLfFH2U=",
    version = "v0.0.0-20200416200610-e657995f937b",
)

go_repository(
    name = "com_github_tetafro_godot",
    importpath = "github.com/tetafro/godot",
    sum = "h1:Dib7un+rYJFUi8vN0Bk6EHheKy6fv6ZzFURHw75g6m8=",
    version = "v0.4.2",
)

go_repository(
    name = "com_github_tj_assert",
    importpath = "github.com/tj/assert",
    sum = "h1:Rw8kxzWo1mr6FSaYXjQELRe88y2KdfynXdnK72rdjtA=",
    version = "v0.0.0-20171129193455-018094318fb0",
)

go_repository(
    name = "com_github_tj_go_elastic",
    importpath = "github.com/tj/go-elastic",
    sum = "h1:eGaGNxrtoZf/mBURsnNQKDR7u50Klgcf2eFDQEnc8Bc=",
    version = "v0.0.0-20171221160941-36157cbbebc2",
)

go_repository(
    name = "com_github_tj_go_kinesis",
    importpath = "github.com/tj/go-kinesis",
    sum = "h1:m74UWYy+HBs+jMFR9mdZU6shPewugMyH5+GV6LNgW8w=",
    version = "v0.0.0-20171128231115-08b17f58cb1b",
)

go_repository(
    name = "com_github_tj_go_spin",
    importpath = "github.com/tj/go-spin",
    sum = "h1:lhdWZsvImxvZ3q1C5OIB7d72DuOwP4O2NdBg9PyzNds=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_vdemeester_k8s_pkg_credentialprovider",
    importpath = "github.com/vdemeester/k8s-pkg-credentialprovider",
    sum = "h1:czKEIG2Q3YRTgs6x/8xhjVMJD5byPo6cZuostkbTM74=",
    version = "v1.17.4",
)

go_repository(
    name = "com_github_vito_go_interact",
    importpath = "github.com/vito/go-interact",
    sum = "h1:O8xi8c93bRUv2Tb/v6HdiuGc+WnWt+AQzF74MOOdlBs=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_xanzy_go_gitlab",
    importpath = "github.com/xanzy/go-gitlab",
    sum = "h1:tBm+OXv1t+KBsqlXkSDFz+YUjRM0GFsjpOWYOod3Ebs=",
    version = "v0.32.0",
)

go_repository(
    name = "dev_gocloud",
    importpath = "gocloud.dev",
    sum = "h1:EDRyaRAnMGSq/QBto486gWFxMLczAfIYUmusV7XLNBM=",
    version = "v0.19.0",
)

go_repository(
    name = "in_gopkg_fsnotify_fsnotify_v1",
    importpath = "gopkg.in/fsnotify/fsnotify.v1",
    sum = "h1:XNNYLJHt73EyYiCZi6+xjupS9CpvmiDgjPTAjrBlQbo=",
    version = "v1.4.7",
)

go_repository(
    name = "io_gitea_code_sdk_gitea",
    importpath = "code.gitea.io/sdk/gitea",
    sum = "h1:hvDCz4wtFvo7rf5Ebj8tGd4aJ4wLPKX3BKFX9Dk1Pgs=",
    version = "v0.12.0",
)

go_repository(
    name = "io_opencensus_go_contrib_exporter_aws",
    importpath = "contrib.go.opencensus.io/exporter/aws",
    sum = "h1:YsbWYxDZkC7x2OxlsDEYvvEXZ3cBI3qBgUK5BqkZvRw=",
    version = "v0.0.0-20181029163544-2befc13012d0",
)

go_repository(
    name = "io_opencensus_go_contrib_integrations_ocsql",
    importpath = "contrib.go.opencensus.io/integrations/ocsql",
    sum = "h1:kfg5Yyy1nYUrqzyfW5XX+dzMASky8IJXhtHe0KTYNS4=",
    version = "v0.1.4",
)

go_repository(
    name = "org_bazil_fuse",
    importpath = "bazil.org/fuse",
    sum = "h1:FNCRpXiquG1aoyqcIWVFmpTSKVcx2bQD38uZZeGtdlw=",
    version = "v0.0.0-20180421153158-65cc252bf669",
)

go_repository(
    name = "com_github_caddy_dns_alidns",
    importpath = "github.com/caddy-dns/alidns",
    sum = "h1:B2JSeQrx2S3xgO5bte7SANmBhyoFfSd9v/Blo631gRc=",
    version = "v1.0.23",
)

go_repository(
    name = "com_github_caddy_dns_azure",
    importpath = "github.com/caddy-dns/azure",
    sum = "h1:j38dM8EAtK0u3CzKUML736MomOsbvyylMQ2hhZ3gUS0=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_caddy_dns_cloudflare",
    importpath = "github.com/caddy-dns/cloudflare",
    sum = "h1:usrheD4ZMCLeI7CI+vHff7+I+D8wIL8Gr3dTlkSmkBA=",
    version = "v0.0.0-20210607183747-91cf700356a1",
)

go_repository(
    name = "com_github_caddy_dns_digitalocean",
    importpath = "github.com/caddy-dns/digitalocean",
    sum = "h1:5JXI9U73KPkSnEveF8DTiSpZiZkNOgOr0mMJZeDx4s8=",
    version = "v0.0.0-20210809220558-ac6e4fd9e135",
)

go_repository(
    name = "com_github_caddy_dns_dnspod",
    importpath = "github.com/caddy-dns/dnspod",
    sum = "h1:txebUF94u/Y8Wt6+hZCmBlD7UQFaNbUUMcAEq8PjgRk=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_caddy_dns_duckdns",
    importpath = "github.com/caddy-dns/duckdns",
    sum = "h1:SSqTt/kbtV9RmJ4wqys5j2ltDwfF35Tv+UU4qaOmTpY=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_caddy_dns_route53",
    importpath = "github.com/caddy-dns/route53",
    sum = "h1:msy8aKay1JwTgN2u4kFlqREjlavpUX8mjHVFS+ZDqEY=",
    version = "v1.1.3",
)

go_repository(
    name = "com_github_caddy_dns_vultr",
    importpath = "github.com/caddy-dns/vultr",
    sum = "h1:L86RWlbq2cObphvPPnUbL4Ap+gcF72NBUga9i4iPV34=",
    version = "v0.0.0-20211122185502-733392841379",
)

go_repository(
    name = "com_github_digitalocean_godo",
    importpath = "github.com/digitalocean/godo",
    sum = "h1:WYy7MIVVhTMZUNB+UA3irl2V9FyDJeDttsifYyn7jYA=",
    version = "v1.41.0",
)

go_repository(
    name = "com_github_libdns_alidns",
    importpath = "github.com/libdns/alidns",
    sum = "h1:KwuC1AmihOJjoFWXFRFaQx1PcD/jRpY04jo7yNK5zfk=",
    version = "v1.0.2-x2",
)

go_repository(
    name = "com_github_libdns_azure",
    importpath = "github.com/libdns/azure",
    sum = "h1:SVYG+iMKtSpSJZBZ0hjETAMNscPoWRMJI7nnlLonwD4=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_libdns_cloudflare",
    importpath = "github.com/libdns/cloudflare",
    sum = "h1:93WkJaGaiXCe353LHEP36kAWCUw0YjFqwhkBkU2/iic=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libdns_digitalocean",
    importpath = "github.com/libdns/digitalocean",
    sum = "h1:JYi/h0UEECrxY2JCi5FIfZEDFuUJvwihUWdm3bnDu2A=",
    version = "v0.0.0-20210310230526-186c4ebd2215",
)

go_repository(
    name = "com_github_libdns_dnspod",
    importpath = "github.com/libdns/dnspod",
    sum = "h1:xJHDIujgLjvZnpB8/rMoCHUqA/KxSGBqRUXxSIzNzAA=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_libdns_duckdns",
    importpath = "github.com/libdns/duckdns",
    sum = "h1:wkgu98DkwpjduH2fxC2YkCiNkNnfQiHaYskMkEyRZ28=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_libdns_route53",
    importpath = "github.com/libdns/route53",
    sum = "h1:etUVkopzG9xGEt34xfmYbpz6rTgAnv+n0vcV/1Xdc7c=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_libdns_vultr",
    importpath = "github.com/libdns/vultr",
    sum = "h1:ds9Nu9RwQWIHXM/e7264RiUfdyAgNdoZkiWXt0xSIzY=",
    version = "v0.0.0-20211122184636-cd4cb5c12e51",
)

go_repository(
    name = "com_github_vultr_govultr_v2",
    importpath = "github.com/vultr/govultr/v2",
    sum = "h1:cvxH1mC/VTs2oRx3njhIhpypg2EBE70rlxAKKtRU/Zg=",
    version = "v2.11.0",
)

go_repository(
    name = "com_github_adrg_xdg",
    importpath = "github.com/adrg/xdg",
    sum = "h1:s/tV7MdqQnzB1nKY8aqHvAMD+uCiuEDzVB5HLRY849U=",
    version = "v0.3.3",
)

go_repository(
    name = "com_github_aead_siphash",
    importpath = "github.com/aead/siphash",
    sum = "h1:FwHfE/T45KPKYuuSAKyyvE+oPWcaQ+CUmFW0bPlM+kg=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_btcsuite_btcd",
    importpath = "github.com/btcsuite/btcd",
    sum = "h1:Ik4hyJqN8Jfyv3S4AGBOmyouMsYE3EdYODkMbQjwPGw=",
    version = "v0.20.1-beta",
)

go_repository(
    name = "com_github_btcsuite_btclog",
    importpath = "github.com/btcsuite/btclog",
    sum = "h1:bAs4lUbRJpnnkd9VhRV3jjAVU7DJVjMaK+IsvSeZvFo=",
    version = "v0.0.0-20170628155309-84c8d2346e9f",
)

go_repository(
    name = "com_github_btcsuite_btcutil",
    importpath = "github.com/btcsuite/btcutil",
    sum = "h1:yJzD/yFppdVCf6ApMkVy8cUxV0XrxdP9rVf6D87/Mng=",
    version = "v0.0.0-20190425235716-9e5f4b9a998d",
)

go_repository(
    name = "com_github_btcsuite_go_socks",
    importpath = "github.com/btcsuite/go-socks",
    sum = "h1:R/opQEbFEy9JGkIguV40SvRY1uliPX8ifOvi6ICsFCw=",
    version = "v0.0.0-20170105172521-4720035b7bfd",
)

go_repository(
    name = "com_github_btcsuite_goleveldb",
    importpath = "github.com/btcsuite/goleveldb",
    sum = "h1:qdGvebPBDuYDPGi1WCPjy1tGyMpmDK8IEapSsszn7HE=",
    version = "v0.0.0-20160330041536-7834afc9e8cd",
)

go_repository(
    name = "com_github_btcsuite_snappy_go",
    importpath = "github.com/btcsuite/snappy-go",
    sum = "h1:ZA/jbKoGcVAnER6pCHPEkGdZOV7U1oLUedErBHCUMs0=",
    version = "v0.0.0-20151229074030-0bdef8d06723",
)

go_repository(
    name = "com_github_btcsuite_websocket",
    importpath = "github.com/btcsuite/websocket",
    sum = "h1:R8vQdOQdZ9Y3SkEwmHoWBmX1DNXhXZqlTpq6s4tyJGc=",
    version = "v0.0.0-20150119174127-31079b680792",
)

go_repository(
    name = "com_github_btcsuite_winsvc",
    importpath = "github.com/btcsuite/winsvc",
    sum = "h1:J9B4L7e3oqhXOcm+2IuNApwzQec85lE+QaikUcCs+dk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_davidlazar_go_crypto",
    importpath = "github.com/davidlazar/go-crypto",
    sum = "h1:6xT9KW8zLC5IlbaIF5Q7JNieBoACT7iW0YTxQHR0in0=",
    version = "v0.0.0-20170701192655-dcfb0a7ac018",
)

go_repository(
    name = "com_github_flynn_noise",
    importpath = "github.com/flynn/noise",
    sum = "h1:DlTHqmzmvcEiKj+4RYo/imoswx/4r6iBlCMfVtrMXpQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_check_check",
    importpath = "github.com/go-check/check",
    sum = "h1:0gkP6mzaMqkmpcJYCFOLkIBwI7xFExG03bbkOkCvUPI=",
    version = "v0.0.0-20180628173108-788fd7840127",
)

go_repository(
    name = "com_github_gxed_hashland_keccakpg",
    importpath = "github.com/gxed/hashland/keccakpg",
    sum = "h1:wrk3uMNaMxbXiHibbPO4S0ymqJMm41WiudyFSs7UnsU=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_gxed_hashland_murmur3",
    importpath = "github.com/gxed/hashland/murmur3",
    sum = "h1:SheiaIt0sda5K+8FLz952/1iWS9zrnKsEJaOJu4ZbSc=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_huin_goupnp",
    importpath = "github.com/huin/goupnp",
    sum = "h1:wg75sLpL6DZqwHQN6E1Cfk6mtfzS45z8OV+ic+DtHRo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_huin_goutil",
    importpath = "github.com/huin/goutil",
    sum = "h1:vlNjIqmUZ9CMAWsbURYl3a6wZbw7q5RHVvlXTNS/Bs8=",
    version = "v0.0.0-20170803182201-1ca381bf3150",
)

go_repository(
    name = "com_github_ipfs_go_cid",
    importpath = "github.com/ipfs/go-cid",
    sum = "h1:ysQJVJA3fNDF1qigJbsSQOdjhVLsOEoPdh0+R97k3jY=",
    version = "v0.0.7",
)

go_repository(
    name = "com_github_ipfs_go_datastore",
    importpath = "github.com/ipfs/go-datastore",
    sum = "h1:cwOUcGMLdLPWgu3SlrCckCMznaGADbPqE0r8h768/Dg=",
    version = "v0.4.5",
)

go_repository(
    name = "com_github_ipfs_go_detect_race",
    importpath = "github.com/ipfs/go-detect-race",
    sum = "h1:qX/xay2W3E4Q1U7d9lNs1sU9nvguX0a7319XbyQ6cOk=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ipfs_go_ds_badger",
    importpath = "github.com/ipfs/go-ds-badger",
    sum = "h1:J27YvAcpuA5IvZUbeBxOcQgqnYHUPxoygc6QxxkodZ4=",
    version = "v0.2.3",
)

go_repository(
    name = "com_github_ipfs_go_ds_leveldb",
    importpath = "github.com/ipfs/go-ds-leveldb",
    sum = "h1:QmQoAJ9WkPMUfBLnu1sBVy0xWWlJPg0m4kRAiJL9iaw=",
    version = "v0.4.2",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_delay",
    importpath = "github.com/ipfs/go-ipfs-delay",
    sum = "h1:NAviDvJ0WXgD+yiL2Rj35AmnfgI11+pHXbdciD917U0=",
    version = "v0.0.0-20181109222059-70721b86a9a8",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_util",
    importpath = "github.com/ipfs/go-ipfs-util",
    sum = "h1:59Sswnk1MFaiq+VcaknX7aYEyGyGDAA73ilhEK2POp8=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_ipfs_go_ipns",
    importpath = "github.com/ipfs/go-ipns",
    sum = "h1:oq4ErrV4hNQ2Eim257RTYRgfOSV/s8BDaf9iIl4NwFs=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_ipfs_go_log",
    importpath = "github.com/ipfs/go-log",
    sum = "h1:6nLQdX4W8P9yZZFH7mO+X/PzjN8Laozm/lMJ6esdgzY=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_ipfs_go_log_v2",
    importpath = "github.com/ipfs/go-log/v2",
    sum = "h1:G4TtqN+V9y9HY9TA6BwbCVyyBZ2B9MbCjR2MtGx8FR0=",
    version = "v2.1.1",
)

go_repository(
    name = "com_github_jackpal_gateway",
    importpath = "github.com/jackpal/gateway",
    sum = "h1:qzXWUJfuMdlLMtt0a3Dgt+xkWQiA5itDEITVJtuSwMc=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_jackpal_go_nat_pmp",
    importpath = "github.com/jackpal/go-nat-pmp",
    sum = "h1:KzKSgb7qkJvOUTqYl9/Hg/me3pWgBmERKrTGD7BdWus=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_jbenet_go_cienv",
    importpath = "github.com/jbenet/go-cienv",
    sum = "h1:Vc/s0QbQtoxX8MwwSLWWh+xNNZvM3Lw7NsTcHrvvhMc=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_jbenet_go_temp_err_catcher",
    importpath = "github.com/jbenet/go-temp-err-catcher",
    sum = "h1:zpb3ZH6wIE8Shj2sKS+khgRvf7T7RABoLk/+KKHggpk=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_jbenet_goprocess",
    importpath = "github.com/jbenet/goprocess",
    sum = "h1:DRGOFReOMqqDNXwW70QkacFW0YN9QnwLV0Vqk+3oU0o=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_jrick_logrotate",
    importpath = "github.com/jrick/logrotate",
    sum = "h1:lQ1bL/n9mBNeIXoTUoYRlK4dHuNJVofX9oWqBtPnSzI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_k0kubun_go_ansi",
    importpath = "github.com/k0kubun/go-ansi",
    sum = "h1:qGQQKEcAR99REcMpsXCp3lJ03zYT1PkRd3kQGPn9GVg=",
    version = "v0.0.0-20180517002512-3bf9e2903213",
)

go_repository(
    name = "com_github_kami_zh_go_capturer",
    importpath = "github.com/kami-zh/go-capturer",
    sum = "h1:cVtBfNW5XTHiKQe7jDaDBSh/EVM4XLPutLAGboIXuM0=",
    version = "v0.0.0-20171211120116-e492ea43421d",
)

go_repository(
    name = "com_github_kkdai_bstream",
    importpath = "github.com/kkdai/bstream",
    sum = "h1:FOOIBWrEkLgmlgGfMuZT83xIwfPDxEI2OHu6xUmJMFE=",
    version = "v0.0.0-20161212061736-f391b8402d23",
)

go_repository(
    name = "com_github_koron_go_ssdp",
    importpath = "github.com/koron/go-ssdp",
    sum = "h1:68u9r4wEvL3gYg2jvAOgROwZ3H+Y3hIDk4tbbmIjcYQ=",
    version = "v0.0.0-20191105050749-2e1c40ed0b5d",
)

go_repository(
    name = "com_github_kubuxu_go_os_helper",
    importpath = "github.com/Kubuxu/go-os-helper",
    sum = "h1:EJiD2VUQyh5A9hWJLmc6iWg6yIcJ7jpBcwC8GMGXfDk=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_libp2p_go_addr_util",
    importpath = "github.com/libp2p/go-addr-util",
    sum = "h1:7cWK5cdA5x72jX0g8iLrQWm5TRJZ6CzGdPEhWj7plWU=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_libp2p_go_buffer_pool",
    importpath = "github.com/libp2p/go-buffer-pool",
    sum = "h1:QNK2iAFa8gjAe1SPz6mHSMuCcjs+X1wlHzeOSqcmlfs=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_libp2p_go_cidranger",
    importpath = "github.com/libp2p/go-cidranger",
    sum = "h1:ewPN8EZ0dd1LSnrtuwd4709PXVcITVeuwbag38yPW7c=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_libp2p_go_conn_security_multistream",
    importpath = "github.com/libp2p/go-conn-security-multistream",
    sum = "h1:uNiDjS58vrvJTg9jO6bySd1rMKejieG7v45ekqHbZ1M=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_libp2p_go_eventbus",
    importpath = "github.com/libp2p/go-eventbus",
    sum = "h1:VanAdErQnpTioN2TowqNcOijf6YwhuODe4pPKSDpxGc=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_libp2p_go_flow_metrics",
    importpath = "github.com/libp2p/go-flow-metrics",
    sum = "h1:8tAs/hSdNvUiLgtlSy3mxwxWP4I9y/jlkPFT7epKdeM=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_libp2p_go_libp2p",
    importpath = "github.com/libp2p/go-libp2p",
    sum = "h1:tDdrXARSghmusdm0nf1U/4M8aj8Rr0V2IzQOXmbzQ3s=",
    version = "v0.13.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_asn_util",
    importpath = "github.com/libp2p/go-libp2p-asn-util",
    sum = "h1:BM7aaOF7RpmNn9+9g6uTjGJ0cTzWr5j9i9IKeun2M8U=",
    version = "v0.0.0-20200825225859-85005c6cf052",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_autonat",
    importpath = "github.com/libp2p/go-libp2p-autonat",
    sum = "h1:3y8XQbpr+ssX8QfZUHekjHCYK64sj6/4hnf/awD4+Ug=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_blankhost",
    importpath = "github.com/libp2p/go-libp2p-blankhost",
    sum = "h1:3EsGAi0CBGcZ33GwRuXEYJLLPoVWyXJ1bcJzAJjINkk=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_circuit",
    importpath = "github.com/libp2p/go-libp2p-circuit",
    sum = "h1:eqQ3sEYkGTtybWgr6JLqJY6QLtPWRErvFjFDfAOO1wc=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_core",
    importpath = "github.com/libp2p/go-libp2p-core",
    sum = "h1:5K3mT+64qDTKbV3yTdbMCzJ7O6wbNsavAEb8iqBvBcI=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_crypto",
    importpath = "github.com/libp2p/go-libp2p-crypto",
    sum = "h1:k9MFy+o2zGDNGsaoZl0MA3iZ75qXxr9OOoAZF+sD5OQ=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_discovery",
    importpath = "github.com/libp2p/go-libp2p-discovery",
    sum = "h1:Qfl+e5+lfDgwdrXdu4YNCWyEo3fWuP+WgN9mN0iWviQ=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_kad_dht",
    importpath = "github.com/libp2p/go-libp2p-kad-dht",
    sum = "h1:FsriVQhOUZpCotWIjyFSjEDNJmUzuMma/RyyTDZanwc=",
    version = "v0.11.1",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_kbucket",
    importpath = "github.com/libp2p/go-libp2p-kbucket",
    sum = "h1:spZAcgxifvFZHBD8tErvppbnNiKA5uokDu3CV7axu70=",
    version = "v0.4.7",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_loggables",
    importpath = "github.com/libp2p/go-libp2p-loggables",
    sum = "h1:h3w8QFfCt2UJl/0/NW4K829HX/0S4KD31PQ7m8UXXO8=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_mplex",
    importpath = "github.com/libp2p/go-libp2p-mplex",
    sum = "h1:/pyhkP1nLwjG3OM+VuaNJkQT/Pqq73WzB3aDN3Fx1sc=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_nat",
    importpath = "github.com/libp2p/go-libp2p-nat",
    sum = "h1:wMWis3kYynCbHoyKLPBEMu4YRLltbm8Mk08HGSfvTkU=",
    version = "v0.0.6",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_netutil",
    importpath = "github.com/libp2p/go-libp2p-netutil",
    sum = "h1:zscYDNVEcGxyUpMd0JReUZTrpMfia8PmLKcKF72EAMQ=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_noise",
    importpath = "github.com/libp2p/go-libp2p-noise",
    sum = "h1:vqYQWvnIcHpIoWJKC7Al4D6Hgj0H012TuXRhPwSMGpQ=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peer",
    importpath = "github.com/libp2p/go-libp2p-peer",
    sum = "h1:EQ8kMjaCUwt/Y5uLgjT8iY2qg0mGUT0N1zUjer50DsY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peerstore",
    importpath = "github.com/libp2p/go-libp2p-peerstore",
    sum = "h1:2ACefBX23iMdJU9Ke+dcXt3w86MIryes9v7In4+Qq3U=",
    version = "v0.2.6",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_pnet",
    importpath = "github.com/libp2p/go-libp2p-pnet",
    sum = "h1:J6htxttBipJujEjz1y0a5+eYoiPcFHhSYHH6na5f0/k=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_record",
    importpath = "github.com/libp2p/go-libp2p-record",
    sum = "h1:R27hoScIhQf/A8XJZ8lYpnqh9LatJ5YbHs28kCIfql0=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_routing_helpers",
    importpath = "github.com/libp2p/go-libp2p-routing-helpers",
    sum = "h1:xY61alxJ6PurSi+MXbywZpelvuU4U4p/gPTxjqCqTzY=",
    version = "v0.2.3",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_secio",
    importpath = "github.com/libp2p/go-libp2p-secio",
    sum = "h1:rLLPvShPQAcY6eNurKNZq3eZjPWfU9kXF2eI9jIYdrg=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_swarm",
    importpath = "github.com/libp2p/go-libp2p-swarm",
    sum = "h1:hahq/ijRoeH6dgROOM8x7SeaKK5VgjjIr96vdrT+NUA=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_testing",
    importpath = "github.com/libp2p/go-libp2p-testing",
    sum = "h1:PrwHRi0IGqOwVQWR3xzgigSlhlLfxgfXgkHxr77EghQ=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_tls",
    importpath = "github.com/libp2p/go-libp2p-tls",
    sum = "h1:twKMhMu44jQO+HgQK9X8NHO5HkeJu2QbhLzLJpa8oNM=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport_upgrader",
    importpath = "github.com/libp2p/go-libp2p-transport-upgrader",
    sum = "h1:xwj4h3hJdBrxqMOyMUjwscjoVst0AASTsKtZiTChoHI=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_yamux",
    importpath = "github.com/libp2p/go-libp2p-yamux",
    sum = "h1:sX4WQPHMhRxJE5UZTfjEuBvlQWXB5Bo3A2JK9ZJ9EM0=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_libp2p_go_maddr_filter",
    importpath = "github.com/libp2p/go-maddr-filter",
    sum = "h1:4ACqZKw8AqiuJfwFGq1CYDFugfXTOos+qQ3DETkhtCE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_libp2p_go_mplex",
    importpath = "github.com/libp2p/go-mplex",
    sum = "h1:U1T+vmCYJaEoDJPV1aq31N56hS+lJgb397GsylNSgrU=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_libp2p_go_msgio",
    importpath = "github.com/libp2p/go-msgio",
    sum = "h1:lQ7Uc0kS1wb1EfRxO2Eir/RJoHkHn7t6o+EiwsYIKJA=",
    version = "v0.0.6",
)

go_repository(
    name = "com_github_libp2p_go_nat",
    importpath = "github.com/libp2p/go-nat",
    sum = "h1:qxnwkco8RLKqVh1NmjQ+tJ8p8khNLFxuElYG/TwqW4Q=",
    version = "v0.0.5",
)

go_repository(
    name = "com_github_libp2p_go_netroute",
    importpath = "github.com/libp2p/go-netroute",
    sum = "h1:1ngWRx61us/EpaKkdqkMjKk/ufr/JlIFYQAxV2XX8Ig=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_libp2p_go_openssl",
    importpath = "github.com/libp2p/go-openssl",
    sum = "h1:eCAzdLejcNVBzP/iZM9vqHnQm+XyCEbSSIheIPRGNsw=",
    version = "v0.0.7",
)

go_repository(
    name = "com_github_libp2p_go_reuseport",
    importpath = "github.com/libp2p/go-reuseport",
    sum = "h1:XSG94b1FJfGA01BUrT82imejHQyTxO4jEWqheyCXYvU=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_libp2p_go_reuseport_transport",
    importpath = "github.com/libp2p/go-reuseport-transport",
    sum = "h1:OZGz0RB620QDGpv300n1zaOcKGGAoGVf8h9txtt/1uM=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_libp2p_go_sockaddr",
    importpath = "github.com/libp2p/go-sockaddr",
    sum = "h1:tCuXfpA9rq7llM/v834RKc/Xvovy/AqM9kHvTV/jY/Q=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_libp2p_go_stream_muxer",
    importpath = "github.com/libp2p/go-stream-muxer",
    sum = "h1:Ce6e2Pyu+b5MC1k3eeFtAax0pW4gc6MosYSLV05UeLw=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_libp2p_go_stream_muxer_multistream",
    importpath = "github.com/libp2p/go-stream-muxer-multistream",
    sum = "h1:TqnSHPJEIqDEO7h1wZZ0p3DXdvDSiLHQidKKUGZtiOY=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_libp2p_go_tcp_transport",
    importpath = "github.com/libp2p/go-tcp-transport",
    sum = "h1:ExZiVQV+h+qL16fzCWtd1HSzPsqWottJ8KXwWaVi8Ns=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_libp2p_go_ws_transport",
    importpath = "github.com/libp2p/go-ws-transport",
    sum = "h1:9tvtQ9xbws6cA5LvqdE6Ne3vcmGB4f1z9SByggk4s0k=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_libp2p_go_yamux",
    importpath = "github.com/libp2p/go-yamux",
    sum = "h1:P1Fe9vF4th5JOxxgQvfbOHkrGqIZniTLf+ddhZp8YTI=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_libp2p_go_yamux_v2",
    importpath = "github.com/libp2p/go-yamux/v2",
    sum = "h1:vSGhAy5u6iHBq11ZDcyHH4Blcf9xlBhT4WQDoOE90LU=",
    version = "v2.0.0",
)

go_repository(
    name = "com_github_minio_blake2b_simd",
    importpath = "github.com/minio/blake2b-simd",
    sum = "h1:lYpkrQH5ajf0OXOcUbGjvZxxijuBwbbmlSxLiuofa+g=",
    version = "v0.0.0-20160723061019-3f5f724cb5b1",
)

go_repository(
    name = "com_github_mitchellh_colorstring",
    importpath = "github.com/mitchellh/colorstring",
    sum = "h1:62I3jR2EmQ4l5rM/4FEfDWcRD+abF5XlKShorW5LRoQ=",
    version = "v0.0.0-20190213212951-d06e56a500db",
)

go_repository(
    name = "com_github_mr_tron_base58",
    importpath = "github.com/mr-tron/base58",
    sum = "h1:T/HDJBh4ZCPbU39/+c3rRvE0uKBQlU27+QI8LJ4t64o=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_multiformats_go_base32",
    importpath = "github.com/multiformats/go-base32",
    sum = "h1:tw5+NhuwaOjJCC5Pp82QuXbrmLzWg7uxlMFp8Nq/kkI=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_multiformats_go_base36",
    importpath = "github.com/multiformats/go-base36",
    sum = "h1:JR6TyF7JjGd3m6FbLU2cOxhC0Li8z8dLNGQ89tUg4F4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr",
    importpath = "github.com/multiformats/go-multiaddr",
    sum = "h1:1bxa+W7j9wZKTZREySx1vPMs2TqrYWjVZ7zE6/XLG1I=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_dns",
    importpath = "github.com/multiformats/go-multiaddr-dns",
    sum = "h1:YWJoIDwLePniH7OU5hBnDZV6SWuvJqJ0YtN6pLeH9zA=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_fmt",
    importpath = "github.com/multiformats/go-multiaddr-fmt",
    sum = "h1:WLEFClPycPkp4fnIzoFoV9FVd49/eQsuaL3/CWe167E=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_net",
    importpath = "github.com/multiformats/go-multiaddr-net",
    sum = "h1:MSXRGN0mFymt6B1yo/6BPnIRpLPEnKgQNvVfCX5VDJk=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_multiformats_go_multibase",
    importpath = "github.com/multiformats/go-multibase",
    sum = "h1:l/B6bJDQjvQ5G52jw4QGSYeOTZoAwIO77RblWplfIqk=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_multiformats_go_multihash",
    importpath = "github.com/multiformats/go-multihash",
    sum = "h1:hWOPdrNqDjwHDx82vsYGSDZNyktOJJ2dzZJzFkOV1jM=",
    version = "v0.0.15",
)

go_repository(
    name = "com_github_multiformats_go_multistream",
    importpath = "github.com/multiformats/go-multistream",
    sum = "h1:6AuNmQVKUkRnddw2YiDjt5Elit40SFxMJkVnhmETXtU=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_multiformats_go_varint",
    importpath = "github.com/multiformats/go-varint",
    sum = "h1:gk85QWKxh3TazbLxED/NlDVv8+q+ReFJk7Y2W/KhfNY=",
    version = "v0.0.6",
)

go_repository(
    name = "com_github_rivo_uniseg",
    importpath = "github.com/rivo/uniseg",
    sum = "h1:S1pD9weZBuJdFmowNwbpi7BJ8TNftyUImj/0WQi72jY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_schollz_pake_v2",
    importpath = "github.com/schollz/pake/v2",
    sum = "h1:pMd7VRwEytDLFXP3vf5UksO6/zxcEoEhb4jaqb6NRR4=",
    version = "v2.0.7",
)

go_repository(
    name = "com_github_schollz_progressbar_v3",
    importpath = "github.com/schollz/progressbar/v3",
    sum = "h1:BKyefEMgFBDbo+JaeqHcm/9QdSj8qG8sUY+6UppGpnw=",
    version = "v3.8.0",
)

go_repository(
    name = "com_github_smola_gocompat",
    importpath = "github.com/smola/gocompat",
    sum = "h1:6b1oIMlUXIpz//VKEDzPVBK8KG7beVwmHIUEBIs/Pns=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_spacemonkeygo_openssl",
    importpath = "github.com/spacemonkeygo/openssl",
    sum = "h1:/eS3yfGjQKG+9kayBkj0ip1BGhq6zJ3eaVksphxAaek=",
    version = "v0.0.0-20181017203307-c2dcc5cca94a",
)

go_repository(
    name = "com_github_spacemonkeygo_spacelog",
    importpath = "github.com/spacemonkeygo/spacelog",
    sum = "h1:RC6RW7j+1+HkWaX/Yh71Ee5ZHaHYt7ZP4sQgUrm6cDU=",
    version = "v0.0.0-20180420211403-2296661a0572",
)

go_repository(
    name = "com_github_src_d_envconfig",
    importpath = "github.com/src-d/envconfig",
    sum = "h1:/AJi6DtjFhZKNx3OB2qMsq7y4yT5//AeSZIe7rk+PX8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_syndtr_goleveldb",
    importpath = "github.com/syndtr/goleveldb",
    sum = "h1:fBdIW9lB4Iz0n9khmH8w27SJ3QEJ7+IgjPEwGSZiFdE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_tscholl2_siec",
    importpath = "github.com/tscholl2/siec",
    sum = "h1:tZWtuLE+LbUwT4OP1oWBSB9zXA8qmQ5qEm4kV9R72oo=",
    version = "v0.0.0-20191122224205-8da93652b094",
)

go_repository(
    name = "com_github_tyler_smith_go_bip39",
    importpath = "github.com/tyler-smith/go-bip39",
    sum = "h1:5eUemwrMargf3BSLRRCalXT93Ns6pQJIjYQN2nyfOP8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_urfave_cli_v2",
    importpath = "github.com/urfave/cli/v2",
    sum = "h1:qph92Y649prgesehzOrQjdWyxFOp/QVM+6imKHad91M=",
    version = "v2.3.0",
)

go_repository(
    name = "com_github_whyrusleeping_go_keyspace",
    importpath = "github.com/whyrusleeping/go-keyspace",
    sum = "h1:EKhdznlJHPMoKr0XTrX+IlJs1LH3lyx2nfr1dOlZ79k=",
    version = "v0.0.0-20160322163242-5b898ac5add1",
)

go_repository(
    name = "com_github_whyrusleeping_go_logging",
    importpath = "github.com/whyrusleeping/go-logging",
    sum = "h1:fwpzlmT0kRC/Fmd0MdmGgJG/CXIZ6gFq46FQZjprUcc=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_whyrusleeping_mafmt",
    importpath = "github.com/whyrusleeping/mafmt",
    sum = "h1:TCghSl5kkwEE0j+sU/gudyhVMRlpBin8fMBBHg59EbA=",
    version = "v1.2.8",
)

go_repository(
    name = "com_github_whyrusleeping_mdns",
    importpath = "github.com/whyrusleeping/mdns",
    sum = "h1:Y1/FEOpaCpD21WxrmfeIYCFPuVPRCY2XZTWzTNHGw30=",
    version = "v0.0.0-20190826153040-b9b60ed33aa9",
)

go_repository(
    name = "com_github_whyrusleeping_multiaddr_filter",
    importpath = "github.com/whyrusleeping/multiaddr-filter",
    sum = "h1:E9S12nwJwEOXe2d6gT6qxdvqMnNq+VnSsKPgm2ZZNds=",
    version = "v0.0.0-20160516205228-e903e4adabd7",
)

go_repository(
    name = "com_github_x_cray_logrus_prefixed_formatter",
    importpath = "github.com/x-cray/logrus-prefixed-formatter",
    sum = "h1:00txxvfBM9muc0jiLIEAkAcIMJzfthRT6usrui8uGmg=",
    version = "v0.5.2",
)

go_repository(
    name = "in_gopkg_src_d_go_cli_v0",
    importpath = "gopkg.in/src-d/go-cli.v0",
    sum = "h1:mXa4inJUuWOoA4uEROxtJ3VMELMlVkIxIfcR0HBekAM=",
    version = "v0.0.0-20181105080154-d492247bbc0d",
)

go_repository(
    name = "in_gopkg_src_d_go_log_v1",
    importpath = "gopkg.in/src-d/go-log.v1",
    sum = "h1:heWvX7J6qbGWbeFS/aRmiy1eYaT+QMV6wNvHDyMjQV4=",
    version = "v1.0.1",
)

go_repository(
    name = "org_uber_go_goleak",
    importpath = "go.uber.org/goleak",
    sum = "h1:gZAh5/EyT/HQwlpkCy6wTpqfH9H8Lz8zbm3dZh+OyzA=",
    version = "v1.1.12",
)

go_repository(
    name = "com_github_fd_go_nat",
    importpath = "github.com/fd/go-nat",
    sum = "h1:DPyQ97sxA9ThrWYRPcWUz/z9TnpTIGRYODIQc/dy64M=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_ipfs_go_todocounter",
    importpath = "github.com/ipfs/go-todocounter",
    sum = "h1:CMTP+27nBrAvFJtH5D/jXc9/c2ZDVaoBAyLYgY4aw8U=",
    version = "v0.0.0-20161011214330-1e832b829506",
)

go_repository(
    name = "com_github_jbenet_go_context",
    importpath = "github.com/jbenet/go-context",
    sum = "h1:BQSFePA1RWJOlocH6Fxy8MmwDt+yVQYULKfN0RoTN8A=",
    version = "v0.0.0-20150711004518-d14ea06fba99",
)

go_repository(
    name = "com_github_libp2p_go_conn_security",
    importpath = "github.com/libp2p/go-conn-security",
    sum = "h1:9mAB66HuUt8oS9GWLTAmV2o3Qvd9lZJMPUCsieuNPp0=",
    version = "v0.0.0-20190226201940-b2fb4ac68c41",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_host",
    importpath = "github.com/libp2p/go-libp2p-host",
    sum = "h1:arERhb/u8bUN55fZ7YrG7EzJoQbfSRccbZMATNOXF2E=",
    version = "v0.0.0-20190226202054-2aa838333967",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_interface_connmgr",
    importpath = "github.com/libp2p/go-libp2p-interface-connmgr",
    sum = "h1:Z2k4gQ2S+iVqLNKHIsVrY1v7I2b16YwoH5zCOFcVKvM=",
    version = "v0.0.0-20190226202048-6a96ea29ac65",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_interface_pnet",
    importpath = "github.com/libp2p/go-libp2p-interface-pnet",
    sum = "h1:TiljmHO1c0NEBUZ/gEjxRKBhMl2dLy9+WL8iKdbax00=",
    version = "v0.0.0-20180919000501-d240acf619f6",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_metrics",
    importpath = "github.com/libp2p/go-libp2p-metrics",
    sum = "h1:d+QubbWRbDzYvAf5B/xxXkQMJY8nojml2oQlVFkLPf4=",
    version = "v0.0.0-20190226174147-1f0f4db04727",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_net",
    importpath = "github.com/libp2p/go-libp2p-net",
    sum = "h1:Rx1GOAyaeba8cISnhBpvU5Zlz9mzgFuqJdVXTTsQoA4=",
    version = "v0.0.0-20190226201932-e71fff5ba6e9",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_protocol",
    importpath = "github.com/libp2p/go-libp2p-protocol",
    sum = "h1:ftsFA3rhuSWX7hNR1CalvmnCKhmncZG7Nyd79Ll6lh0=",
    version = "v0.0.0-20171212212132-b29f3d97e3a2",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_routing",
    importpath = "github.com/libp2p/go-libp2p-routing",
    sum = "h1:J/JUNkEXrVPAyJlEmMcIJcLEC2BX7Q8Uwm1vGSwnQw8=",
    version = "v0.0.0-20190226202213-efb912a4c564",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport",
    importpath = "github.com/libp2p/go-libp2p-transport",
    sum = "h1:ErBx9a936Bg6OXLfYn/XA4BzA0SnSZ/SKsFq9QPsj2M=",
    version = "v0.0.0-20190226201958-e8580c8a519d",
)

go_repository(
    name = "com_github_libp2p_go_testutil",
    importpath = "github.com/libp2p/go-testutil",
    sum = "h1:IxUJRyKDh3qEm3qwBQ8IA30EG4ChyR2Sgek2ZcaIqMs=",
    version = "v0.0.0-20190226202041-873eaa1a26ba",
)

go_repository(
    name = "com_github_whyrusleeping_base32",
    importpath = "github.com/whyrusleeping/base32",
    sum = "h1:BCPnHtcboadS0DvysUuJXZ4lWVv5Bh5i7+tbIyi+ck4=",
    version = "v0.0.0-20170828182744-c30ac30633cc",
)

go_repository(
    name = "com_github_whyrusleeping_go_notifier",
    importpath = "github.com/whyrusleeping/go-notifier",
    sum = "h1:M/lL30eFZTKnomXY6huvM6G0+gVquFNf6mxghaWlFUg=",
    version = "v0.0.0-20170827234753-097c5d47330f",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_multiplex",
    importpath = "github.com/whyrusleeping/go-smux-multiplex",
    sum = "h1:iqksILj8STw03EJQe7Laj4ubnw+ojOyik18cd5vPL1o=",
    version = "v3.0.16+incompatible",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_multistream",
    importpath = "github.com/whyrusleeping/go-smux-multistream",
    sum = "h1:BdYHctE9HJZLquG9tpTdwWcbG4FaX6tVKPGjCGgiVxo=",
    version = "v2.0.2+incompatible",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_yamux",
    importpath = "github.com/whyrusleeping/go-smux-yamux",
    sum = "h1:IGm/UP/JpEFS6D787sZnZg7RA6fZIR9c/Ms9DeAVNuk=",
    version = "v2.0.8+incompatible",
)

go_repository(
    name = "com_github_whyrusleeping_yamux",
    importpath = "github.com/whyrusleeping/yamux",
    sum = "h1:4CK3aUUJQu0qpKZv5gEWJjNOQtdbdDhVVS6PJ+HimdE=",
    version = "v1.1.5",
)

go_repository(
    name = "com_github_go_kit_log",
    importpath = "github.com/go-kit/log",
    sum = "h1:7i2K3eKTos3Vc0enKCfnVcgHh2olr/MyfboYq7cAcFw=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_go_task_slim_sprig",
    importpath = "github.com/go-task/slim-sprig",
    sum = "h1:p104kn46Q8WdvHunIJ9dAyjPVtrBPhSr3KT2yUst43I=",
    version = "v0.0.0-20210107165309-348f09dbbbc0",
)

go_repository(
    name = "com_github_marten_seemann_qtls_go1_16",
    importpath = "github.com/marten-seemann/qtls-go1-16",
    sum = "h1:o9JrYPPco/Nukd/HpOHMHZoBDXQqoNtUCmny98/1uqQ=",
    version = "v0.1.5",
)

go_repository(
    name = "com_github_marten_seemann_qtls_go1_17",
    importpath = "github.com/marten-seemann/qtls-go1-17",
    sum = "h1:DQjHPq+aOzUeh9/lixAGunn6rIOQyWChPSI4+hgW7jc=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_shopspring_decimal",
    importpath = "github.com/shopspring/decimal",
    sum = "h1:abSATXmQEYyShuxI4/vyW3tV1MrKAJzCZ/0zLUXYbsQ=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_stoewer_go_strcase",
    importpath = "github.com/stoewer/go-strcase",
    sum = "h1:Z2iHWqGXH00XYgqDmNgQbIBxf3wrNq0F3feEy0ainaU=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_thales_e_security_pool",
    importpath = "github.com/thales-e-security/pool",
    sum = "h1:RAPs4q2EbWsTit6tpzuvTFlgFRJ3S8Evf5gtvVDbmPg=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_thalesignite_crypto11",
    importpath = "github.com/ThalesIgnite/crypto11",
    sum = "h1:3MebRK/U0mA2SmSthXAIZAdUA9w8+ZuKem2O6HuR1f8=",
    version = "v1.2.4",
)

go_repository(
    name = "com_sslmate_software_src_go_pkcs12",
    importpath = "software.sslmate.com/src/go-pkcs12",
    sum = "h1:SqYE5+A2qvRhErbsXFfUEUmpWEKxxRSMgGLkvRAFOV4=",
    version = "v0.0.0-20210415151418-c5206de65a78",
)

go_repository(
    name = "org_mozilla_go_pkcs7",
    importpath = "go.mozilla.org/pkcs7",
    sum = "h1:CCriYyAfq1Br1aIYettdHZTy8mBTIPo7We18TuO/bak=",
    version = "v0.0.0-20210826202110-33d05740a352",
)

go_repository(
    name = "sm_step_go_cli_utils",
    importpath = "go.step.sm/cli-utils",
    sum = "h1:2GvY5Muid1yzp7YQbfCCS+gK3q7zlHjjLL5Z0DXz8ds=",
    version = "v0.7.0",
)

go_repository(
    name = "com_gitlab_mjwhitta_cli",
    importpath = "gitlab.com/mjwhitta/cli",
    sum = "h1:GbfhWvEDASibMOAvp2QuSSPfURej85uqx7qECFbN3Z0=",
    version = "v1.9.0",
)

go_repository(
    name = "com_gitlab_mjwhitta_hilighter",
    importpath = "gitlab.com/mjwhitta/hilighter",
    sum = "h1:TmE55gpilZjXrGQWYvT8RADD36kf4oBdlAPIJQDMUnU=",
    version = "v1.10.1",
)

go_repository(
    name = "com_gitlab_mjwhitta_jq",
    importpath = "gitlab.com/mjwhitta/jq",
    sum = "h1:XEjsAfYAP+4Fc3FWgd5O58Gsak3WuOsVE5lH6EsxpII=",
    version = "v1.5.2",
)

go_repository(
    name = "com_gitlab_mjwhitta_jsoncfg",
    importpath = "gitlab.com/mjwhitta/jsoncfg",
    sum = "h1:X6tWIpgSlLoiHOvhK5JMziXTaoXPUec5gkm28jLJTyM=",
    version = "v1.6.6",
)

go_repository(
    name = "com_gitlab_mjwhitta_log",
    importpath = "gitlab.com/mjwhitta/log",
    sum = "h1:g+Pr678Sh/404fnYNnIokKe9cQh/eC21ZdJ7sb+mRQM=",
    version = "v1.6.0",
)

go_repository(
    name = "com_gitlab_mjwhitta_pathname",
    importpath = "gitlab.com/mjwhitta/pathname",
    sum = "h1:GWLgFYO4RjYutbpV4s0096xLAGEIZbLnsoz2r6rcBf0=",
    version = "v1.2.0",
)

go_repository(
    name = "com_gitlab_mjwhitta_safety",
    importpath = "gitlab.com/mjwhitta/safety",
    sum = "h1:Q3N8HoD5Wx60UT4I0n/ivFwM+737u9PILq3G0SWgkbI=",
    version = "v1.11.0",
)

go_repository(
    name = "com_gitlab_mjwhitta_sysinfo",
    importpath = "gitlab.com/mjwhitta/sysinfo",
    sum = "h1:FD/J3aGxyfqE799jKm8v3qPyvQhs2nx90dHFowkJR90=",
    version = "v1.4.6",
)

go_repository(
    name = "com_gitlab_mjwhitta_where",
    importpath = "gitlab.com/mjwhitta/where",
    sum = "h1:UfJSi6vHklt5ufnZ+lnIF+vHmrBFUSbHZt/hF7Nox+k=",
    version = "v1.2.3",
)

go_repository(
    name = "com_github_twitchyliquid64_golang_asm",
    importpath = "github.com/twitchyliquid64/golang-asm",
    sum = "h1:SU5vSMR7hnwNxj24w34ZyCi/FmDZTkS4MhqMhdFk5YI=",
    version = "v0.15.1",
)

go_repository(
    name = "com_github_go_gost_gosocks4",
    importpath = "github.com/go-gost/gosocks4",
    sum = "h1:+k1sec8HlELuQV7rWftIkmy8UijzUt2I6t+iMPlGB2s=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_go_gost_gosocks5",
    importpath = "github.com/go-gost/gosocks5",
    sum = "h1:Hkmp9YDRBSCJd7xywW6dBPT6B9aQTkuWd+3WCheJiJA=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_go_gost_tls_dissector",
    importpath = "github.com/go-gost/tls-dissector",
    sum = "h1:73NGqAs22ey3wJkIYVD/ACEoovuIuOlEzQTEoqrO5+U=",
    version = "v0.0.2-0.20211125135007-2b5d5bd9c07e",
)

go_repository(
    name = "com_github_marten_seemann_qtls_go1_18",
    importpath = "github.com/marten-seemann/qtls-go1-18",
    sum = "h1:qp7p7XXUFL7fpBvSS1sWD+uSqPvzNQK43DH+/qEkj0Y=",
    version = "v0.1.1",
)

go_repository(
    name = "com_gitlab_yawning_bsaes_git",
    importpath = "gitlab.com/yawning/bsaes.git",
    sum = "h1:FpfFs4EhNehiVfzQttTuxanPIT43FtkkCFypIod8LHo=",
    version = "v0.0.0-20190805113838-0a714cd429ec",
)

go_repository(
    name = "com_gitlab_yawning_obfs4_git",
    importpath = "gitlab.com/yawning/obfs4.git",
    sum = "h1:w/f20IHUkUYEp+xYgpKz4Bs78zms0DbjPZCep5lc0xA=",
    version = "v0.0.0-20210511220700-e330d1b7024b",
)

go_repository(
    name = "com_gitlab_yawning_utls_git",
    importpath = "gitlab.com/yawning/utls.git",
    sum = "h1:RL6O0MP2YI0KghuEU/uGN6+8b4183eqNWoYgx7CXD0U=",
    version = "v0.0.12-1",
)

go_repository(
    name = "com_github_alcortesm_tgz",
    importpath = "github.com/alcortesm/tgz",
    sum = "h1:uSoVVbwJiQipAclBbw+8quDsfcvFjOpI5iCf4p/cqCs=",
    version = "v0.0.0-20161220082320-9c5fe88206d7",
)

go_repository(
    name = "com_github_aokoli_goutils",
    importpath = "github.com/aokoli/goutils",
    sum = "h1:7fpzNGoJ3VA8qcrm++XEE1QUe0mIwNeLa02Nwq7RDkg=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_apache_beam",
    importpath = "github.com/apache/beam",
    sum = "h1:iYCqXSPdvj8YW6LS3MleV080RZKf5ULS+L490PjJY7Y=",
    version = "v2.30.0+incompatible",
)

go_repository(
    name = "com_github_benbjohnson_clock",
    importpath = "github.com/benbjohnson/clock",
    sum = "h1:Q92kusRqC1XV2MjkWETPvjJVqKetz1OzxZB7mHJLju8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_cncf_xds_go",
    importpath = "github.com/cncf/xds/go",
    sum = "h1:zH8ljVhhq7yC0MIeUL/IviMtY8hx2mK8cN9wEYb8ggw=",
    version = "v0.0.0-20211011173535-cb28da3451f1",
)

go_repository(
    name = "com_github_fullstorydev_grpcurl",
    importpath = "github.com/fullstorydev/grpcurl",
    sum = "h1:Pp648wlTTg3OKySeqxM5pzh8XF6vLqrm8wRq66+5Xo0=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_go_redis_redis",
    importpath = "github.com/go-redis/redis",
    sum = "h1:K0pv1D7EQUjfyoMql+r/jZqCLizCGKFlFgcHWWmHQjg=",
    version = "v6.15.9+incompatible",
)

go_repository(
    name = "com_github_google_go_licenses",
    importpath = "github.com/google/go-licenses",
    sum = "h1:JtmsUf+m+KdwCOgLG578T0Mvd0+l+dezPrJh5KYnXZg=",
    version = "v0.0.0-20210329231322-ce1d9163b77d",
)

go_repository(
    name = "com_github_google_licenseclassifier",
    importpath = "github.com/google/licenseclassifier",
    sum = "h1:EfzlPF5MRmoWsCGvSkPZ1Nh9uVzHf4FfGnDQ6CXd2NA=",
    version = "v0.0.0-20210325184830-bb04aff29e72",
)

go_repository(
    name = "com_github_groob_finalizer",
    importpath = "github.com/groob/finalizer",
    sum = "h1:5ikpG9mYCMFiZX0nkxoV6aU2IpCHPdws3gCNgdZeEV0=",
    version = "v0.0.0-20170707115354-4c2ed49aabda",
)

go_repository(
    name = "com_github_jhump_protoreflect",
    importpath = "github.com/jhump/protoreflect",
    sum = "h1:npqHz788dryJiR/l6K/RUQAyh2SwV91+d1dnh4RjO9w=",
    version = "v1.9.0",
)

go_repository(
    name = "com_github_kevinburke_ssh_config",
    importpath = "github.com/kevinburke/ssh_config",
    sum = "h1:DowS9hvgyYSX4TO5NpyC606/Z4SxnNYbT+WX27or6Ck=",
    version = "v0.0.0-20201106050909-4977a11b4351",
)

go_repository(
    name = "com_github_letsencrypt_pkcs11key_v4",
    importpath = "github.com/letsencrypt/pkcs11key/v4",
    sum = "h1:qLc/OznH7xMr5ARJgkZCCWk+EomQkiNTOoOF5LAgagc=",
    version = "v4.0.0",
)

go_repository(
    name = "com_github_masterminds_sprig",
    importpath = "github.com/Masterminds/sprig",
    sum = "h1:z4yfnGrZ7netVz+0EDJ0Wi+5VZCSYp4Z0m2dk6cEM60=",
    version = "v2.22.0+incompatible",
)

go_repository(
    name = "com_github_micromdm_scep_v2",
    importpath = "github.com/micromdm/scep/v2",
    sum = "h1:2fS9Rla7qRR266hvUoEauBJ7J6FhgssEiq2OkSKXmaU=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_mwitkow_go_proto_validators",
    importpath = "github.com/mwitkow/go-proto-validators",
    sum = "h1:F6LFfmgVnfULfaRsQWBbe7F7ocuHCr9+7m+GAeDzNbQ=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_nishanths_predeclared",
    importpath = "github.com/nishanths/predeclared",
    sum = "h1:3f0nxAmdj/VoCGN/ijdMy7bj6SBagaqYg1B0hu8clMA=",
    version = "v0.0.0-20200524104333-86fad755b4d3",
)

go_repository(
    name = "com_github_otiai10_copy",
    importpath = "github.com/otiai10/copy",
    sum = "h1:HvG945u96iNadPoG2/Ja2+AUJeW5YuFQMixq9yirC+k=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_otiai10_curr",
    importpath = "github.com/otiai10/curr",
    sum = "h1:TJIWdbX0B+kpNagQrjgq8bCMrbhiuX73M2XwgtDMoOI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_otiai10_mint",
    importpath = "github.com/otiai10/mint",
    sum = "h1:BCmzIS3n71sGfHB5NMNDB3lHYPz8fWSkCAErHed//qc=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_pelletier_go_buffruneio",
    importpath = "github.com/pelletier/go-buffruneio",
    sum = "h1:U4t4R6YkofJ5xHm3dJzuRpPZ0mr5MMCoAWooScCR7aA=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_pseudomuto_protoc_gen_doc",
    importpath = "github.com/pseudomuto/protoc-gen-doc",
    sum = "h1:aNTZq0dy0Pq2ag2v7bhNKFNgBBA8wMCoJSChhd7RciE=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_pseudomuto_protokit",
    importpath = "github.com/pseudomuto/protokit",
    sum = "h1:hlnBDcy3YEDXH7kc9gV+NLaN0cDzhDvD1s7Y6FZ8RpM=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_src_d_gcfg",
    importpath = "github.com/src-d/gcfg",
    sum = "h1:xXbNR5AlLSA315x2UO+fTSSAXCDf+Ar38/6oyGbDKQ4=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_xanzy_ssh_agent",
    importpath = "github.com/xanzy/ssh-agent",
    sum = "h1:wUMzuKtKilRgBAD1sUb8gOwwRr2FGoBVumcjoOACClI=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_spanner",
    importpath = "cloud.google.com/go/spanner",
    sum = "h1:XvqWpUANh4XZbtLj7gwvTYkYAx0D0x0nTOVKrdEP+Uk=",
    version = "v1.20.0",
)

go_repository(
    name = "in_gopkg_src_d_go_billy_v4",
    importpath = "gopkg.in/src-d/go-billy.v4",
    sum = "h1:0SQA1pRztfTFx2miS8sA97XvooFeNOmvUenF4o0EcVg=",
    version = "v4.3.2",
)

go_repository(
    name = "in_gopkg_src_d_go_git_fixtures_v3",
    importpath = "gopkg.in/src-d/go-git-fixtures.v3",
    sum = "h1:ivZFOIltbce2Mo8IjzUHAFoq/IylO9WHhNOAJK+LsJg=",
    version = "v3.5.0",
)

go_repository(
    name = "in_gopkg_src_d_go_git_v4",
    importpath = "gopkg.in/src-d/go-git.v4",
    sum = "h1:SRtFyV8Kxc0UP7aCHcijOMQGPxHSmMOPrzulQWolkYE=",
    version = "v4.13.1",
)

go_repository(
    name = "io_etcd_go_etcd_api_v3",
    importpath = "go.etcd.io/etcd/api/v3",
    sum = "h1:GsV3S+OfZEOCNXdtNkBSR7kgLobAa/SO6tCxRa0GAYw=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_client_pkg_v3",
    importpath = "go.etcd.io/etcd/client/pkg/v3",
    sum = "h1:2aQv6F436YnN7I4VbI8PPYrBhu+SmrTaADcf8Mi/6PU=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_client_v2",
    importpath = "go.etcd.io/etcd/client/v2",
    sum = "h1:ftQ0nOOHMcbMS3KIaDQ0g5Qcd6bhaBrQT6b89DfwLTs=",
    version = "v2.305.0",
)

go_repository(
    name = "io_etcd_go_etcd_client_v3",
    importpath = "go.etcd.io/etcd/client/v3",
    sum = "h1:62Eh0XOro+rDwkrypAGDfgmNh5Joq+z+W9HZdlXMzek=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_etcdctl_v3",
    importpath = "go.etcd.io/etcd/etcdctl/v3",
    sum = "h1:i8DGjR9gBRoS6NEHF3XBxxh7QwL1DyilXMCkHpyy6zM=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_etcdutl_v3",
    importpath = "go.etcd.io/etcd/etcdutl/v3",
    sum = "h1:orNfs85GWmiOl0p23Yi9YRfHNb3Qfdlt0wVFkPTRVxQ=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_pkg_v3",
    importpath = "go.etcd.io/etcd/pkg/v3",
    sum = "h1:ntrg6vvKRW26JRmHTE0iNlDgYK6JX3hg/4cD62X0ixk=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_raft_v3",
    importpath = "go.etcd.io/etcd/raft/v3",
    sum = "h1:kw2TmO3yFTgE+F0mdKkG7xMxkit2duBDa2Hu6D/HMlw=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_server_v3",
    importpath = "go.etcd.io/etcd/server/v3",
    sum = "h1:jk8D/lwGEDlQU9kZXUFMSANkE22Sg5+mW27ip8xcF9E=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_tests_v3",
    importpath = "go.etcd.io/etcd/tests/v3",
    sum = "h1:+uMuHYKKlLUzbW322XrQXbaGM9qiV7vUL+AEPT/KYY4=",
    version = "v3.5.0",
)

go_repository(
    name = "io_etcd_go_etcd_v3",
    importpath = "go.etcd.io/etcd/v3",
    sum = "h1:fs7tB+L/xRDi/+p9qKuaPGCtMX6vkovLRXTqvEE98Ek=",
    version = "v3.5.0",
)

go_repository(
    name = "io_opentelemetry_go_contrib",
    importpath = "go.opentelemetry.io/contrib",
    sum = "h1:ubFQUn0VCZ0gPwIoJfBJVpeBlyRMxu8Mm/huKWYd9p0=",
    version = "v0.20.0",
)

go_repository(
    name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
    importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
    sum = "h1:sO4WKdPAudZGKPcpZT4MJn6JaDmpyLrMPDGGyA1SttE=",
    version = "v0.20.0",
)

go_repository(
    name = "io_opentelemetry_go_otel",
    importpath = "go.opentelemetry.io/otel",
    sum = "h1:7ESuKPq6zpjRaY5nvVDGiuwK7VAJ8MwkKnmNJ9whNZ4=",
    version = "v1.4.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp",
    importpath = "go.opentelemetry.io/otel/exporters/otlp",
    sum = "h1:PTNgq9MRmQqqJY0REVbZFvwkYOA85vbdQU/nVfxDyqg=",
    version = "v0.20.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_metric",
    importpath = "go.opentelemetry.io/otel/metric",
    sum = "h1:HhJPsGhJoKRSegPQILFbODU56NS/L1UE4fS1sC5kIwQ=",
    version = "v0.27.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_oteltest",
    importpath = "go.opentelemetry.io/otel/oteltest",
    sum = "h1:HiITxCawalo5vQzdHfKeZurV8x7ljcqAgiWzF6Vaeaw=",
    version = "v0.20.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk",
    importpath = "go.opentelemetry.io/otel/sdk",
    sum = "h1:LJE4SW3jd4lQTESnlpQZcBhQ3oci0U2MLR5uhicfTHQ=",
    version = "v1.4.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk_export_metric",
    importpath = "go.opentelemetry.io/otel/sdk/export/metric",
    sum = "h1:c5VRjxCXdQlx1HjzwGdQHzZaVI82b5EbBgOu2ljD92g=",
    version = "v0.20.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk_metric",
    importpath = "go.opentelemetry.io/otel/sdk/metric",
    sum = "h1:7ao1wpzHRVKf0OQ7GIxiQJA6X7DLX9o14gmVon7mMK8=",
    version = "v0.20.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_trace",
    importpath = "go.opentelemetry.io/otel/trace",
    sum = "h1:4OOUrPZdVFQkbzl/JSdvGCWIdw5ONXXxzHlaLlWppmo=",
    version = "v1.4.0",
)

go_repository(
    name = "io_opentelemetry_go_proto_otlp",
    importpath = "go.opentelemetry.io/proto/otlp",
    sum = "h1:CMJ/3Wp7iOWES+CYLfnBv+DVmPbB+kmy9PJ92XvlR6c=",
    version = "v0.12.0",
)

go_repository(
    name = "org_bitbucket_creachadair_shell",
    importpath = "bitbucket.org/creachadair/shell",
    sum = "h1:reJflDbKqnlnqb4Oo2pQ1/BqmY/eCWcNGHrIUO8qIzc=",
    version = "v0.0.6",
)

go_repository(
    name = "org_golang_google_grpc_cmd_protoc_gen_go_grpc",
    importpath = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
    sum = "h1:M1YKkFIboKNieVO5DLUEVzQfGwJD30Nv2jfUgzb5UcE=",
    version = "v1.1.0",
)

go_repository(
    name = "sm_step_go_linkedca",
    importpath = "go.step.sm/linkedca",
    sum = "h1:lEkGRDY+u7FudGKt8yEo7nBy5OzceO9s3rl+/sZVL5M=",
    version = "v0.15.0",
)

go_repository(
    name = "com_github_dreamacro_clash",
    importpath = "github.com/Dreamacro/clash",
    sum = "h1:apgeijehqa5XlHbIMFJyq57yMvQWNqDix4/GoY8mttQ=",
    version = "v1.10.0",
)

go_repository(
    name = "com_github_dreamacro_go_shadowsocks2",
    importpath = "github.com/Dreamacro/go-shadowsocks2",
    sum = "h1:8CtbE1HoPPMfrQZGXmlluq6dO2lL31W6WRRE8fabc4Q=",
    version = "v0.1.7",
)

go_repository(
    name = "com_github_fanliao_go_promise",
    importpath = "github.com/fanliao/go-promise",
    sum = "h1:0eU/faU2oDIB2BkQVM02hgRLJjGzzUuRf19HUhp0394=",
    version = "v0.0.0-20141029170127-1890db352a72",
)

go_repository(
    name = "com_github_go_chi_chi_v5",
    importpath = "github.com/go-chi/chi/v5",
    sum = "h1:rDTPXLDHGATaeHvVlLcR4Qe0zftYethFucbjVQ1PxU8=",
    version = "v5.0.7",
)

go_repository(
    name = "com_github_go_chi_cors",
    importpath = "github.com/go-chi/cors",
    sum = "h1:tV1g1XENQ8ku4Bq3K9ub2AtgG+p16SmzeMSGTwrOKdE=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_go_chi_render",
    importpath = "github.com/go-chi/render",
    sum = "h1:4/5tis2cKaNdnv9zFLfXzcquC9HbeZgCnxGnKrltBS8=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_hugelgupf_socketpair",
    importpath = "github.com/hugelgupf/socketpair",
    sum = "h1:/jC7qQFrv8CrSJVmaolDVOxTfS9kc36uB6H40kdbQq8=",
    version = "v0.0.0-20190730060125-05d35a94e714",
)

go_repository(
    name = "com_github_insomniacslk_dhcp",
    importpath = "github.com/insomniacslk/dhcp",
    sum = "h1:efcJu2Vzz6DoSq245deWNzTz6l/gsqdphm3FjmI88/g=",
    version = "v0.0.0-20220119180841-3c283ff8b7dd",
)

go_repository(
    name = "com_github_mdlayher_ethernet",
    importpath = "github.com/mdlayher/ethernet",
    sum = "h1:lez6TS6aAau+8wXUP3G9I3TGlmPFEq2CTxBaRqY6AGE=",
    version = "v0.0.0-20190606142754-0394541c37b7",
)

go_repository(
    name = "com_github_mdlayher_raw",
    importpath = "github.com/mdlayher/raw",
    sum = "h1:aFkJ6lx4FPip+S+Uw4aTegFMct9shDvP+79PsSxpm3w=",
    version = "v0.0.0-20191009151244-50f2db8cc065",
)

go_repository(
    name = "com_github_oschwald_geoip2_golang",
    importpath = "github.com/oschwald/geoip2-golang",
    sum = "h1:GKxT3yaWWNXSb7vj6D7eoJBns+lGYgx08QO0UcNm0YY=",
    version = "v1.6.1",
)

go_repository(
    name = "com_github_oschwald_maxminddb_golang",
    importpath = "github.com/oschwald/maxminddb-golang",
    sum = "h1:Uh/DSnGoxsyp/KYbY1AuP0tYEwfs0sCph9p/UMXK/Hk=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_prashantv_gostub",
    importpath = "github.com/prashantv/gostub",
    sum = "h1:BTyx3RfQjRHnUWaGF9oQos79AlQ5k8WNktv7VGvVH4g=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_u_root_uio",
    importpath = "github.com/u-root/uio",
    sum = "h1:BFvcl34IGnw8yvJi8hlqLFo9EshRInwWBs2M5fGWzQA=",
    version = "v0.0.0-20210528114334-82958018845c",
)

go_repository(
    name = "org_uber_go_automaxprocs",
    importpath = "go.uber.org/automaxprocs",
    sum = "h1:e1YG66Lrk73dn4qhg8WFSvhF0JuFQF0ERIp4rpuV8Qk=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_alexflint_go_filemutex",
    importpath = "github.com/alexflint/go-filemutex",
    sum = "h1:AMzIhMUqU3jMrZiTuW0zkYeKlKDAFD+DG20IoO421/Y=",
    version = "v0.0.0-20171022225611-72bdc8eae2ae",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_credentials",
    importpath = "github.com/aws/aws-sdk-go-v2/credentials",
    sum = "h1:2faRNX8JgZVy7dDxERkaGBqb/xo5Rgmc8JMPL5j1o58=",
    version = "v1.6.2",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_internal_configsources",
    importpath = "github.com/aws/aws-sdk-go-v2/internal/configsources",
    sum = "h1:LZwqhOyqQ2w64PZk04V0Om9AEExtW8WMkCRoE1h9/94=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_internal_endpoints_v2",
    importpath = "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2",
    sum = "h1:ObMfGNk0xjOWduPxsrRWVwZZia3e9fOcO6zlKCkt38s=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_ecr",
    importpath = "github.com/aws/aws-sdk-go-v2/service/ecr",
    sum = "h1:onTF83DG9dsRv6UzuhYb7phiktjwQ++s/n+ZtNlTQnM=",
    version = "v1.10.1",
)

go_repository(
    name = "com_github_aws_smithy_go",
    importpath = "github.com/aws/smithy-go",
    sum = "h1:c7FUdEqrQA1/UVKKCNDFQPNKGp4FQg3YW4Ck5SLTG58=",
    version = "v1.9.0",
)

go_repository(
    name = "com_github_bitly_go_simplejson",
    importpath = "github.com/bitly/go-simplejson",
    sum = "h1:6IH+V8/tVMab511d5bn4M7EwGXZf9Hj6i2xSwkNEM+Y=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_bits_and_blooms_bitset",
    importpath = "github.com/bits-and-blooms/bitset",
    sum = "h1:Kn4yilvwNtMACtf1eYDlG8H77R07mZSPbMjLyS07ChA=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_bmizerany_assert",
    importpath = "github.com/bmizerany/assert",
    sum = "h1:DDGfHa7BWjL4YnC6+E63dPcxHo2sUxDIu8g3QgEJdRY=",
    version = "v0.0.0-20160611221934-b7ed37b82869",
)

go_repository(
    name = "com_github_bshuster_repo_logrus_logstash_hook",
    importpath = "github.com/bshuster-repo/logrus-logstash-hook",
    sum = "h1:pgAtgj+A31JBVtEHu2uHuEx0n+2ukqUJnS2vVe5pQNA=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_bugsnag_bugsnag_go",
    importpath = "github.com/bugsnag/bugsnag-go",
    sum = "h1:rFt+Y/IK1aEZkEHchZRSq9OQbsSzIT/OrI8YFFmRIng=",
    version = "v0.0.0-20141110184014-b1d153021fcd",
)

go_repository(
    name = "com_github_bugsnag_osext",
    importpath = "github.com/bugsnag/osext",
    sum = "h1:otBG+dV+YK+Soembjv71DPz3uX/V/6MMlSyD9JBQ6kQ=",
    version = "v0.0.0-20130617224835-0dd3f918b21b",
)

go_repository(
    name = "com_github_bugsnag_panicwrap",
    importpath = "github.com/bugsnag/panicwrap",
    sum = "h1:nvj0OLI3YqYXer/kZD8Ri1aaunCxIEsOst1BVJswV0o=",
    version = "v0.0.0-20151223152923-e2c28503fcd0",
)

go_repository(
    name = "com_github_checkpoint_restore_go_criu_v5",
    importpath = "github.com/checkpoint-restore/go-criu/v5",
    sum = "h1:TW8f/UvntYoVDMN1K2HlT82qH1rb0sOjpGw3m6Ym+i4=",
    version = "v5.0.0",
)

go_repository(
    name = "com_github_circonus_labs_circonus_gometrics",
    importpath = "github.com/circonus-labs/circonus-gometrics",
    sum = "h1:C29Ae4G5GtYyYMm1aztcyj/J5ckgJm2zwdDajFbx1NY=",
    version = "v2.3.1+incompatible",
)

go_repository(
    name = "com_github_circonus_labs_circonusllhist",
    importpath = "github.com/circonus-labs/circonusllhist",
    sum = "h1:TJH+oke8D16535+jHExHj4nQvzlZrj7ug5D7I/orNUA=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_containerd_aufs",
    importpath = "github.com/containerd/aufs",
    sum = "h1:2oeJiwX5HstO7shSrPZjrohJZLzK36wvpdmzDRkL/LY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_containerd_btrfs",
    importpath = "github.com/containerd/btrfs",
    sum = "h1:osn1exbzdub9L5SouXO5swW4ea/xVdJZ3wokxN5GrnA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_containerd_go_cni",
    importpath = "github.com/containerd/go-cni",
    sum = "h1:YbJAhpTevL2v6u8JC1NhCYRwf+3Vzxcc5vGnYoJ7VeE=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_containerd_imgcrypt",
    importpath = "github.com/containerd/imgcrypt",
    sum = "h1:LBwiTfoUsdiEGAR1TpvxE+Gzt7469oVu87iR3mv3Byc=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_containerd_nri",
    importpath = "github.com/containerd/nri",
    sum = "h1:6QioHRlThlKh2RkRTR4kIT3PKAcrLo3gIWnjkM4dQmQ=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_containerd_zfs",
    importpath = "github.com/containerd/zfs",
    sum = "h1:cXLJbx+4Jj7rNsTiqVfm6i+RNLx6FFA2fMmDlEf+Wm8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_containernetworking_plugins",
    importpath = "github.com/containernetworking/plugins",
    sum = "h1:FD1tADPls2EEi3flPc2OegIY1M9pUa9r2Quag7HMLV8=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_containers_ocicrypt",
    importpath = "github.com/containers/ocicrypt",
    sum = "h1:prL8l9w3ntVqXvNH1CiNn5ENjcCnr38JqpSyvKKB4GI=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_d2g_dhcp4",
    importpath = "github.com/d2g/dhcp4",
    sum = "h1:Xo2rK1pzOm0jO6abTPIQwbAmqBIOj132otexc1mmzFc=",
    version = "v0.0.0-20170904100407-a1d1b6c41b1c",
)

go_repository(
    name = "com_github_d2g_dhcp4client",
    importpath = "github.com/d2g/dhcp4client",
    sum = "h1:suYBsYZIkSlUMEz4TAYCczKf62IA2UWC+O8+KtdOhCo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_d2g_dhcp4server",
    importpath = "github.com/d2g/dhcp4server",
    sum = "h1:+CpLbZIeUn94m02LdEKPcgErLJ347NUwxPKs5u8ieiY=",
    version = "v0.0.0-20181031114812-7d4a0a7f59a5",
)

go_repository(
    name = "com_github_d2g_hardwareaddr",
    importpath = "github.com/d2g/hardwareaddr",
    sum = "h1:itqmmf1PFpC4n5JW+j4BU7X4MTfVurhYRTjODoPb2Y8=",
    version = "v0.0.0-20190221164911-e7d9fbe030e4",
)

go_repository(
    name = "com_github_denverdino_aliyungo",
    importpath = "github.com/denverdino/aliyungo",
    sum = "h1:p6poVbjHDkKa+wtC8frBMwQtT3BmqGYBjzMwJ63tuR4=",
    version = "v0.0.0-20190125010748-a747050bb1ba",
)

go_repository(
    name = "com_github_docker_go_events",
    importpath = "github.com/docker/go-events",
    sum = "h1:+pKlWGMw7gf6bQ+oDZB4KHQFypsfjYlq/C4rfL7D3g8=",
    version = "v0.0.0-20190806004212-e31b211e4f1c",
)

go_repository(
    name = "com_github_docker_go_metrics",
    importpath = "github.com/docker/go-metrics",
    sum = "h1:AgB/0SvBxihN0X8OR4SjsblXkbMvalQ8cjmtKQ2rQV8=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_docker_libtrust",
    importpath = "github.com/docker/libtrust",
    sum = "h1:ZClxb8laGDf5arXfYcAtECDFgAgHklGI8CxgjHnXKJ4=",
    version = "v0.0.0-20150114040149-fa567046d9b1",
)

go_repository(
    name = "com_github_fullsailor_pkcs7",
    importpath = "github.com/fullsailor/pkcs7",
    sum = "h1:RDBNVkRviHZtvDvId8XSGPu3rmpmSe+wKRcEWNgsfWU=",
    version = "v0.0.0-20190404230743-d7302db945fa",
)

go_repository(
    name = "com_github_fxamacker_cbor_v2",
    importpath = "github.com/fxamacker/cbor/v2",
    sum = "h1:aM45YGMctNakddNNAezPxDUpv38j44Abh+hifNuqXik=",
    version = "v2.3.0",
)

go_repository(
    name = "com_github_g07cha_defender",
    importpath = "github.com/g07cha/defender",
    sum = "h1:gWvniJ4GbFfkf700kykAImbLiEMU0Q3QN9hQ26Js1pU=",
    version = "v0.0.0-20180505193036-5665c627c814",
)

go_repository(
    name = "com_github_go_git_gcfg",
    importpath = "github.com/go-git/gcfg",
    sum = "h1:Q5ViNfGF8zFgyJWPqYwA7qGFoMTEiBmdlkcfRmpIMa4=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_go_git_go_billy_v5",
    importpath = "github.com/go-git/go-billy/v5",
    sum = "h1:4pl5BV4o7ZG/lterP4S6WzJ6xr49Ba5ET9ygheTYahk=",
    version = "v5.1.0",
)

go_repository(
    name = "com_github_go_git_go_git_v5",
    importpath = "github.com/go-git/go-git/v5",
    sum = "h1:8WKMtJR2j8RntEXR/uvTKagfEt4GYlwQ7mntE4+0GWc=",
    version = "v5.3.0",
)

go_repository(
    name = "com_github_google_shlex",
    importpath = "github.com/google/shlex",
    sum = "h1:El6M4kTTCOh6aBiKaUGG7oYTSPP8MxqL4YI3kZKwcP4=",
    version = "v0.0.0-20191202100458-e7afc7fbc510",
)

go_repository(
    name = "com_github_hashicorp_cronexpr",
    importpath = "github.com/hashicorp/cronexpr",
    sum = "h1:NJZDd87hGXjoZBdvyCF9mX4DCq5Wy7+A/w+A7q0wn6c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_hashicorp_nomad_api",
    importpath = "github.com/hashicorp/nomad/api",
    sum = "h1:Y/0SMGEylkIwKtWZfKA+SVSIbOAKLMQPz/BjJG4vwHo=",
    version = "v0.0.0-20220211135303-4afc67b7002e",
)

go_repository(
    name = "com_github_j_keck_arping",
    importpath = "github.com/j-keck/arping",
    sum = "h1:742eGXur0715JMq73aD95/FU0XpVKXqNuTnEfXsLOYQ=",
    version = "v0.0.0-20160618110441-2cf9dc699c56",
)

go_repository(
    name = "com_github_marstr_guid",
    importpath = "github.com/marstr/guid",
    sum = "h1:/M4H/1G4avsieL6BbUwCOBzulmoeKVP5ux/3mQNnbyI=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_microsoft_hcsshim_test",
    importpath = "github.com/Microsoft/hcsshim/test",
    sum = "h1:4FA+QBaydEHlwxg0lMN3rhwoDaQy6LKhVWR4qvq4BuA=",
    version = "v0.0.0-20210227013316-43a75bb4edd3",
)

go_repository(
    name = "com_github_mitchellh_osext",
    importpath = "github.com/mitchellh/osext",
    sum = "h1:2+myh5ml7lgEU/51gbeLHfKGNfgEQQIWrlbdaOsidbQ=",
    version = "v0.0.0-20151018003038-5e2d6d41470f",
)

go_repository(
    name = "com_github_moby_locker",
    importpath = "github.com/moby/locker",
    sum = "h1:fOXqR41zeveg4fFODix+1Ch4mj/gT0NE1XJbp/epuBg=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_moby_sys_symlink",
    importpath = "github.com/moby/sys/symlink",
    sum = "h1:MTFZ74KtNI6qQQpuBxU+uKCim4WtOMokr03hCfJcazE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_ncw_swift",
    importpath = "github.com/ncw/swift",
    sum = "h1:4DQRPj35Y41WogBxyhOXlrI37nzGlyEcsforeudyYPQ=",
    version = "v1.0.47",
)

go_repository(
    name = "com_github_opencontainers_runtime_tools",
    importpath = "github.com/opencontainers/runtime-tools",
    sum = "h1:H7DMc6FAjgwZZi8BRqjrAAHWoqEr5e5L6pS4V0ezet4=",
    version = "v0.0.0-20181011054405-1d69bd0f9c39",
)

go_repository(
    name = "com_github_orcaman_concurrent_map",
    importpath = "github.com/orcaman/concurrent-map",
    sum = "h1:lNCW6THrCKBiJBpz8kbVGjC7MgdCGKwuvBgc7LoD6sw=",
    version = "v0.0.0-20190826125027-8c72a8bb44f6",
)

go_repository(
    name = "com_github_portainer_docker_compose_wrapper",
    importpath = "github.com/portainer/docker-compose-wrapper",
    sum = "h1:dg/uvltrR++AVDjjVkXKrinZ/T8YlaKeUAOAmQ1i+dk=",
    version = "v0.0.0-20220407011010-3c7408969ad3",
)

go_repository(
    name = "com_github_portainer_libhelm",
    importpath = "github.com/portainer/libhelm",
    sum = "h1:5e8KAnDa2G3cEHK7aV/ue8lOaoQwBZUzoALslwWkR04=",
    version = "v0.0.0-20210929000907-825e93d62108",
)

go_repository(
    name = "com_github_portainer_portainer_api",
    importpath = "github.com/portainer/portainer/api",
    sum = "h1:gUe/VH30JLUrr9cc6wM6ByWB6rZJVBXql/E+A9e25EM=",
    version = "v0.0.0-20220303203420-547d9c2fde15",
)

go_repository(
    name = "com_github_rkl_digest",
    importpath = "github.com/rkl-/digest",
    sum = "h1:rDj3WeO+TiWyxfcydUnKegWAZoR5kQsnW0wzhggdOrw=",
    version = "v0.0.0-20180419075440-8316caa4a777",
)

go_repository(
    name = "com_github_safchain_ethtool",
    importpath = "github.com/safchain/ethtool",
    sum = "h1:SsRnt87qssm3RltLJze6kM+4fs32twq6mZEcBxbDMVg=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_shopify_logrus_bugsnag",
    importpath = "github.com/Shopify/logrus-bugsnag",
    sum = "h1:UrqY+r/OJnIp5u0s1SbQ8dVfLCZJsnvazdBP5hS4iRs=",
    version = "v0.0.0-20171204204709-577dee27f20d",
)

go_repository(
    name = "com_github_stefanberger_go_pkcs11uri",
    importpath = "github.com/stefanberger/go-pkcs11uri",
    sum = "h1:lIOOHPEbXzO3vnmx2gok1Tfs31Q8GQqKLc8vVqyQq/I=",
    version = "v0.0.0-20201008174630-78d3cae3a980",
)

go_repository(
    name = "com_github_tchap_go_patricia",
    importpath = "github.com/tchap/go-patricia",
    sum = "h1:JvoDL7JSoIP2HDE8AbDH3zC8QBPxmzYe32HHy5yQ+Ck=",
    version = "v2.2.6+incompatible",
)

go_repository(
    name = "com_github_tv42_httpunix",
    importpath = "github.com/tv42/httpunix",
    sum = "h1:G3dpKMzFDjgEh2q1Z7zUUtKa8ViPtH+ocF0bE0g00O8=",
    version = "v0.0.0-20150427012821-b75d8614f926",
)

go_repository(
    name = "com_github_x448_float16",
    importpath = "github.com/x448/float16",
    sum = "h1:qLwI1I70+NjRFUR3zs1JPUCgaCXSh3SW62uAKT1mSBM=",
    version = "v0.8.4",
)

go_repository(
    name = "com_github_yvasiyarov_go_metrics",
    importpath = "github.com/yvasiyarov/go-metrics",
    sum = "h1:+lm10QQTNSBd8DVTNGHx7o/IKu9HYDvLMffDhbyLccI=",
    version = "v0.0.0-20140926110328-57bccd1ccd43",
)

go_repository(
    name = "com_github_yvasiyarov_gorelic",
    importpath = "github.com/yvasiyarov/gorelic",
    sum = "h1:hlE8//ciYMztlGpl/VA+Zm1AcTPHYkHJPbHqE6WJUXE=",
    version = "v0.0.0-20141212073537-a9bba5b9ab50",
)

go_repository(
    name = "com_github_yvasiyarov_newrelic_platform_go",
    importpath = "github.com/yvasiyarov/newrelic_platform_go",
    sum = "h1:ERexzlUfuTvpE74urLSbIQW0Z/6hF9t8U4NsJLaioAY=",
    version = "v0.0.0-20140908184405-b21fdbd4370f",
)

go_repository(
    name = "in_gopkg_yaml_v1",
    importpath = "gopkg.in/yaml.v1",
    sum = "h1:POO/ycCATvegFmVuPpQzZFJ+pGZeX22Ufu6fibxDVjU=",
    version = "v1.0.0-20140924161607-9f9df34309c0",
)

go_repository(
    name = "org_golang_google_cloud",
    importpath = "google.golang.org/cloud",
    sum = "h1:Cpp2P6TPjujNoC5M2KHY6g7wfyLYfIWRZaSdIKfDasA=",
    version = "v0.0.0-20151119220103-975617b05ea8",
)

go_repository(
    name = "com_github_azure_go_ntlmssp",
    importpath = "github.com/Azure/go-ntlmssp",
    sum = "h1:/IBSNwUN8+eKzUzbJPqhK839ygXJ82sde8x3ogr6R28=",
    version = "v0.0.0-20200615164410-66371956d46c",
)

go_repository(
    name = "com_github_caddyserver_forwardproxy",
    importpath = "github.com/caddyserver/forwardproxy",
    sum = "h1:wC4VlHOHWTjslaf/SBQ4nFaU1aVUZFRqyzZRnCJNV7k=",
    version = "v0.0.0-20211013034647-8c6ef2bd4a8f",
)

go_repository(
    name = "com_github_cockroachdb_apd",
    importpath = "github.com/cockroachdb/apd",
    sum = "h1:3LFP3629v+1aKXU5Q37mxmRxX/pIu1nijXydLShEq5I=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_cyberdelia_go_metrics_graphite",
    importpath = "github.com/cyberdelia/go-metrics-graphite",
    sum = "h1:M5QgkYacWj0Xs8MhpIK/5uwU02icXpEoSo9sM2aRCps=",
    version = "v0.0.0-20161219230853-39f87cc3b432",
)

go_repository(
    name = "com_github_elazarl_goproxy_ext",
    importpath = "github.com/elazarl/goproxy/ext",
    sum = "h1:dWB6v3RcOy03t/bUadywsbyrQwCqZeNIEX6M1OtSZOM=",
    version = "v0.0.0-20190711103511-473e67f1d7d2",
)

go_repository(
    name = "com_github_evanphx_json_patch_v5",
    importpath = "github.com/evanphx/json-patch/v5",
    sum = "h1:bAmFiUJ+o0o2B4OiTFeE3MqCOtyo+jjPP9iZ0VRxYUc=",
    version = "v5.5.0",
)

go_repository(
    name = "com_github_freman_caddy2_reauth",
    importpath = "github.com/freman/caddy2-reauth",
    sum = "h1:JBn9TEuc5i0FPckHVOL8NI8Q+R0kNBOBLvBLadgkZ3o=",
    version = "v0.0.0-20200518130136-6064aa96b1a8",
)

go_repository(
    name = "com_github_go_logr_stdr",
    importpath = "github.com/go-logr/stdr",
    sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_go_test_deep",
    importpath = "github.com/go-test/deep",
    sum = "h1:onZX1rnHT3Wv6cqNgYyFOOlgVKJrksuCMCRvJStbMYw=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_golang_jwt_jwt_v4",
    importpath = "github.com/golang-jwt/jwt/v4",
    sum = "h1:XUgk2Ex5veyVFVeLm0xhusUTQybEbexJXrvPNOKkSY0=",
    version = "v4.1.0",
)

go_repository(
    name = "com_github_greenpau_caddy_authorize",
    importpath = "github.com/greenpau/caddy-authorize",
    sum = "h1:mbRmYvwKlMFjr4Fubq8UFEyRUgSp2oLAwqawEJrQvA8=",
    version = "v1.3.24",
)

go_repository(
    name = "com_github_hashicorp_go_kms_wrapping_entropy",
    importpath = "github.com/hashicorp/go-kms-wrapping/entropy",
    sum = "h1:xuTi5ZwjimfpvpL09jDE71smCBRpnF5xfo871BSX4gs=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_hashicorp_go_plugin",
    importpath = "github.com/hashicorp/go-plugin",
    sum = "h1:DXmvivbWD5qdiBts9TpBC7BYL1Aia5sxbRgQB+v6UZM=",
    version = "v1.4.3",
)

go_repository(
    name = "com_github_hashicorp_go_secure_stdlib_base62",
    importpath = "github.com/hashicorp/go-secure-stdlib/base62",
    sum = "h1:6KMBnfEv0/kLAz0O76sliN5mXbCDcLfs2kP7ssP7+DQ=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_hashicorp_go_secure_stdlib_mlock",
    importpath = "github.com/hashicorp/go-secure-stdlib/mlock",
    sum = "h1:cCRo8gK7oq6A2L6LICkUZ+/a5rLiRXFMf1Qd4xSwxTc=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_hashicorp_go_secure_stdlib_parseutil",
    importpath = "github.com/hashicorp/go-secure-stdlib/parseutil",
    sum = "h1:78ki3QBevHwYrVxnyVeaEz+7WtifHhauYF23es/0KlI=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_hashicorp_go_secure_stdlib_password",
    importpath = "github.com/hashicorp/go-secure-stdlib/password",
    sum = "h1:6JzmBqXprakgFEHwBgdchsjaA9x3GyjdI568bXKxa60=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_hashicorp_go_secure_stdlib_strutil",
    importpath = "github.com/hashicorp/go-secure-stdlib/strutil",
    sum = "h1:nd0HIW15E6FG1MsnArYaHfuw9C2zgzM8LxkG5Ty/788=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_hashicorp_go_secure_stdlib_tlsutil",
    importpath = "github.com/hashicorp/go-secure-stdlib/tlsutil",
    sum = "h1:Yc026VyMyIpq1UWRnakHRG01U8fJm+nEfEmjoAb00n8=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_hashicorp_vault_api",
    importpath = "github.com/hashicorp/vault/api",
    sum = "h1:pkDkcgTh47PRjY1NEFeofqR4W/HkNUi9qIakESO2aRM=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_hashicorp_vault_api_auth_approle",
    importpath = "github.com/hashicorp/vault/api/auth/approle",
    sum = "h1:R5yA+xcNvw1ix6bDuWOaLOq2L4L77zDCVsethNw97xQ=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_hashicorp_vault_sdk",
    importpath = "github.com/hashicorp/vault/sdk",
    sum = "h1:kR3dpxNkhh/wr6ycaJYqp6AFT/i2xaftbfnwZduTKEY=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_hashicorp_yamux",
    importpath = "github.com/hashicorp/yamux",
    sum = "h1:b5rjCoWHc7eqmAS4/qyk21ZsHyb6Mxv/jykxvNTkU4M=",
    version = "v0.0.0-20180604194846-3520598351bb",
)

go_repository(
    name = "com_github_jackc_chunkreader",
    importpath = "github.com/jackc/chunkreader",
    sum = "h1:4s39bBR8ByfqH+DKm8rQA3E1LHZWB9XWcrz8fqaZbe0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jackc_chunkreader_v2",
    importpath = "github.com/jackc/chunkreader/v2",
    sum = "h1:i+RDz65UE+mmpjTfyz0MoVTnzeYxroil2G82ki7MGG8=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_jackc_pgconn",
    importpath = "github.com/jackc/pgconn",
    sum = "h1:DzdIHIjG1AxGwoEEqS+mGsURyjt4enSmqzACXvVzOT8=",
    version = "v1.10.1",
)

go_repository(
    name = "com_github_jackc_pgio",
    importpath = "github.com/jackc/pgio",
    sum = "h1:g12B9UwVnzGhueNavwioyEEpAmqMe1E/BN9ES+8ovkE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jackc_pgmock",
    importpath = "github.com/jackc/pgmock",
    sum = "h1:DadwsjnMwFjfWc9y5Wi/+Zz7xoE5ALHsRQlOctkOiHc=",
    version = "v0.0.0-20210724152146-4ad1a8207f65",
)

go_repository(
    name = "com_github_jackc_pgpassfile",
    importpath = "github.com/jackc/pgpassfile",
    sum = "h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jackc_pgproto3",
    importpath = "github.com/jackc/pgproto3",
    sum = "h1:FYYE4yRw+AgI8wXIinMlNjBbp/UitDJwfj5LqqewP1A=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_jackc_pgproto3_v2",
    importpath = "github.com/jackc/pgproto3/v2",
    sum = "h1:r7JypeP2D3onoQTCxWdTpCtJ4D+qpKr0TxvoyMhZ5ns=",
    version = "v2.2.0",
)

go_repository(
    name = "com_github_jackc_pgservicefile",
    importpath = "github.com/jackc/pgservicefile",
    sum = "h1:C8S2+VttkHFdOOCXJe+YGfa4vHYwlt4Zx+IVXQ97jYg=",
    version = "v0.0.0-20200714003250-2b9c44734f2b",
)

go_repository(
    name = "com_github_jackc_pgtype",
    importpath = "github.com/jackc/pgtype",
    sum = "h1:/SH1RxEtltvJgsDqp3TbiTFApD3mey3iygpuEGeuBXk=",
    version = "v1.9.0",
)

go_repository(
    name = "com_github_jackc_pgx_v4",
    importpath = "github.com/jackc/pgx/v4",
    sum = "h1:TgdrmgnM7VY72EuSQzBbBd4JA1RLqJolrw9nQVZABVc=",
    version = "v4.14.0",
)

go_repository(
    name = "com_github_jackc_puddle",
    importpath = "github.com/jackc/puddle",
    sum = "h1:DNDKdn/pDrWvDWyT2FYvpZVE81OAhWrjCv19I9n108Q=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_kardianos_service",
    importpath = "github.com/kardianos/service",
    sum = "h1:bGuZ/epo3vrt8IPC7mnKQolqFeYJb7Cs8Rk4PSOBB/g=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_lxn_walk",
    importpath = "github.com/lxn/walk",
    sum = "h1:NVRJ0Uy0SOFcXSKLsS65OmI1sgCCfiDUPj+cwnH7GZw=",
    version = "v0.0.0-20210112085537-c389da54e794",
)

go_repository(
    name = "com_github_lxn_win",
    importpath = "github.com/lxn/win",
    sum = "h1:H+t6A/QJMbhCSEH5rAuRxh+CtW96g0Or0Fxa9IKr4uc=",
    version = "v0.0.0-20210218163916-a377121e959e",
)

go_repository(
    name = "com_github_mastercactapus_proxyprotocol",
    importpath = "github.com/mastercactapus/proxyprotocol",
    sum = "h1:WpDMFKCYdF8NsoA6OrXyNKyZrzMURqqOP1PE7297RCE=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_mathetake_gasm",
    importpath = "github.com/mathetake/gasm",
    sum = "h1:GOar/82ZJdJ9COsZmNfSKgI/h5eYlBj6WvmUkuFEuDE=",
    version = "v0.0.0-20200928142744-80e74517647c",
)

go_repository(
    name = "com_github_mholt_caddy",
    importpath = "github.com/mholt/caddy",
    sum = "h1:KI6RPGih2GFzWRPG8s9clKK28Ns4ZlVMKR/v7mxq6+c=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_nbrownus_go_metrics_prometheus",
    importpath = "github.com/nbrownus/go-metrics-prometheus",
    sum = "h1:8dM0ilqKL0Uzl42GABzzC4Oqlc3kGRILz0vgoff7nwg=",
    version = "v0.0.0-20210712211119-974a6260965f",
)

go_repository(
    name = "com_github_rogpeppe_go_charset",
    importpath = "github.com/rogpeppe/go-charset",
    sum = "h1:BN/Nyn2nWMoqGRA7G7paDNDqTXE30mXGqzzybrfo05w=",
    version = "v0.0.0-20180617210344-2471d30d28b4",
)

go_repository(
    name = "com_github_rs_zerolog",
    importpath = "github.com/rs/zerolog",
    sum = "h1:uPRuwkWF4J6fGsJ2R0Gn2jB1EQiav9k3S6CSdygQJXY=",
    version = "v1.15.0",
)

go_repository(
    name = "com_github_slackhq_nebula",
    importpath = "github.com/slackhq/nebula",
    sum = "h1:wuIOHsOnrNw3rQx8yPxXiGu8wAtAxxtUI/K8W7Vj7EI=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_tailscale_tscert",
    importpath = "github.com/tailscale/tscert",
    sum = "h1:uFx5aih29p2IaRUF0lJwtVViCXStlvnPPE3NEmM4Ivs=",
    version = "v0.0.0-20220125204807-4509a5fbaf74",
)

go_repository(
    name = "com_google_cloud_go_compute",
    importpath = "cloud.google.com/go/compute",
    sum = "h1:mPL/MzDDYHsh5tHRS9mhmhWlcgClCrCa6ApQCU6wnHI=",
    version = "v1.3.0",
)

go_repository(
    name = "com_google_cloud_go_iam",
    importpath = "cloud.google.com/go/iam",
    sum = "h1:W2vbGCrE3Z7J/x3WXLxxGl9LMSB2uhsAA7Ss/6u/qRY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_google_cloud_go_kms",
    importpath = "cloud.google.com/go/kms",
    sum = "h1:iElbfoE61VeLhnZcGOltqL8HIly8Nhbe5t6JlH9GXjo=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_security",
    importpath = "cloud.google.com/go/security",
    sum = "h1:BhCl33x+KQI4qiZnFrfr2gAGhb2aZ0ZvKB3Y4QlEfgo=",
    version = "v1.3.0",
)

go_repository(
    name = "com_zx2c4_golang_wintun",
    importpath = "golang.zx2c4.com/wintun",
    sum = "h1:Ug9qvr1myri/zFN6xL17LSCBGFDnphBBhzmILHsM5TY=",
    version = "v0.0.0-20211104114900-415007cec224",
)

go_repository(
    name = "com_zx2c4_golang_wireguard_windows",
    importpath = "golang.zx2c4.com/wireguard/windows",
    sum = "h1:OnYw96PF+CsIMrqWo5QP3Q59q5hY1rFErk/yN3cS+JQ=",
    version = "v0.5.1",
)

go_repository(
    name = "in_gopkg_inconshreveable_log15_v2",
    importpath = "gopkg.in/inconshreveable/log15.v2",
    sum = "h1:RlWgLqCMMIYYEVcAR5MDsuHlVkaIPDAF+5Dehzg8L5A=",
    version = "v2.0.0-20180818164646-67afb5ed74ec",
)

go_repository(
    name = "io_filippo_edwards25519",
    importpath = "filippo.io/edwards25519",
    sum = "h1:m0VOOB23frXZvAOK44usCgLWvtsxIoMCTBGJZlpmGfU=",
    version = "v1.0.0-rc.1",
)

go_repository(
    name = "io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp",
    importpath = "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp",
    sum = "h1:SLme4Porm+UwX0DdHMxlwRt7FzPSE0sys81bet2o0pU=",
    version = "v0.29.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_internal_retry",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/internal/retry",
    sum = "h1:j7AwzDdAQBJjcqayAaYbvpYeZzII7cEe5qJTu+De6UY=",
    version = "v1.4.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace",
    sum = "h1:lRpP10E8oTGVmY1nVXcwelCT1Z8ca41/l5ce7AqLAss=",
    version = "v1.4.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracegrpc",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc",
    sum = "h1:buSx4AMC/0Z232slPhicN/fU5KIlj0bMngct5pcZhkI=",
    version = "v1.4.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_internal_metric",
    importpath = "go.opentelemetry.io/otel/internal/metric",
    sum = "h1:9dAVGAfFiiEq5NVB9FUJ5et+btbDQAUIJehJ+ikyryk=",
    version = "v0.27.0",
)

go_repository(
    name = "af_inet_netaddr",
    importpath = "inet.af/netaddr",
    sum = "h1:tvgqez5ZQoBBiBAGNU/fmJy247yB/7++kcLOEoMYup0=",
    version = "v0.0.0-20210903134321-85fa6c94624e",
)

go_repository(
    name = "com_github_bmatcuk_doublestar_v4",
    importpath = "github.com/bmatcuk/doublestar/v4",
    sum = "h1:X0krlUVAVmtr2cRoTqR8aDMrDqnB36ht8wpWTiQ3jsA=",
    version = "v4.0.2",
)

go_repository(
    name = "com_github_dsoprea_go_exif_v2",
    importpath = "github.com/dsoprea/go-exif/v2",
    sum = "h1:bVaiYo8amn7Lu93sz6mTlYB3EtLG9aRcMnM1Eps8fmM=",
    version = "v2.0.0-20200321225314-640175a69fe4",
)

go_repository(
    name = "com_github_dsoprea_go_exif_v3",
    importpath = "github.com/dsoprea/go-exif/v3",
    sum = "h1:rz9dPf+Unge2D5RNIRNFvCc2OrGfhAPuxx4L6412jdI=",
    version = "v3.0.0-20201216222538-db167117f483",
)

go_repository(
    name = "com_github_dsoprea_go_logging",
    importpath = "github.com/dsoprea/go-logging",
    sum = "h1:F/7L5wr/fP/SKeO5HuMlNEX9Ipyx2MbH2rV9G4zJRpk=",
    version = "v0.0.0-20200517223158-a10564966e9d",
)

go_repository(
    name = "com_github_dsoprea_go_utility",
    importpath = "github.com/dsoprea/go-utility",
    sum = "h1:/w4QxepU4AHh3AuO6/g8y/YIIHH5+aKP3Bj8sg5cqhU=",
    version = "v0.0.0-20200711062821-fab8125e9bdf",
)

go_repository(
    name = "com_github_dsoprea_go_utility_v2",
    importpath = "github.com/dsoprea/go-utility/v2",
    sum = "h1:IxIbA7VbCNrwumIYjDoMOdf4KOSkMC6NJE4s8oRbE7E=",
    version = "v2.0.0-20200717064901-2fccff4aa15e",
)

go_repository(
    name = "com_github_dvyukov_go_fuzz",
    importpath = "github.com/dvyukov/go-fuzz",
    sum = "h1:q1oJaUPdmpDm/VyXosjgPgr6wS7c5iV2p0PwJD73bUI=",
    version = "v0.0.0-20210103155950-6a8e9d1f2415",
)

go_repository(
    name = "com_github_getkin_kin_openapi",
    importpath = "github.com/getkin/kin-openapi",
    sum = "h1:j77zg3Ec+k+r+GA3d8hBoXpAc6KX9TbBPrwQGBIy2sY=",
    version = "v0.76.0",
)

go_repository(
    name = "com_github_golang_geo",
    importpath = "github.com/golang/geo",
    sum = "h1:C/hKUcHT483btRbeGkrRjJz+Zbcj8audldIi9tRJDCc=",
    version = "v0.0.0-20200319012246-673a6f80352d",
)

go_repository(
    name = "com_github_golang_jwt_jwt",
    importpath = "github.com/golang-jwt/jwt",
    sum = "h1:IfV12K8xAKAnZqdXVzCZ+TOjboZ2keLg81eXfW3O+oY=",
    version = "v3.2.2+incompatible",
)

go_repository(
    name = "com_github_gomarkdown_markdown",
    importpath = "github.com/gomarkdown/markdown",
    sum = "h1:LP/6EfrZ/LyCc+SXvANDrIJ4sP9u2NAtqyv6QknetNQ=",
    version = "v0.0.0-20200824053859-8c8b3816f167",
)

go_repository(
    name = "com_github_gopherjs_websocket",
    importpath = "github.com/gopherjs/websocket",
    sum = "h1:AG0ZwOFfTjzk5n1PXhWttuw3i7i0TeqypKSqCXhMlyM=",
    version = "v0.0.0-20191103002815-9a42957e2b3a",
)

go_repository(
    name = "com_github_hashicorp_go_envparse",
    importpath = "github.com/hashicorp/go-envparse",
    sum = "h1:v1d9+AJMP6i4p8BSKNU0InuvmIAdZjQLNN19V86AG4Q=",
    version = "v0.0.0-20200406174449-d9cfd743a15e",
)

go_repository(
    name = "com_github_illumos_go_kstat",
    importpath = "github.com/illumos/go-kstat",
    sum = "h1:hk4LPqXIY/c9XzRbe7dA6qQxaT6Axcbny0L/G5a4owQ=",
    version = "v0.0.0-20210513183136-173c9b0a9973",
)

go_repository(
    name = "com_github_josharian_native",
    importpath = "github.com/josharian/native",
    sum = "h1:uhL5Gw7BINiiPAo24A2sxkcDI0Jt/sqp1v5xQCniEFA=",
    version = "v0.0.0-20200817173448-b6b71def0850",
)

go_repository(
    name = "com_github_klauspost_pgzip",
    importpath = "github.com/klauspost/pgzip",
    sum = "h1:qnWYvvKqedOF2ulHpMG72XQol4ILEJ8k2wwRl/Km8oE=",
    version = "v1.2.5",
)

go_repository(
    name = "com_github_kr_fs",
    importpath = "github.com/kr/fs",
    sum = "h1:Jskdu9ieNAYnjxsi0LbQp1ulIKZV1LAFgK1tWhpZgl8=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_lunixbochs_struc",
    importpath = "github.com/lunixbochs/struc",
    sum = "h1:EnfXoSqDfSNJv0VBNqY/88RNnhSGYkrHaO0mmFGbVsc=",
    version = "v0.0.0-20200707160740-784aaebc1d40",
)

go_repository(
    name = "com_github_mdlayher_ethtool",
    importpath = "github.com/mdlayher/ethtool",
    sum = "h1:WgyLFv10Ov49JAQI/ZLUkCZ7VJS3r74hwFIGXJsgZlY=",
    version = "v0.0.0-20210210192532-2b88debcdd43",
)

go_repository(
    name = "com_github_mdlayher_socket",
    importpath = "github.com/mdlayher/socket",
    sum = "h1:qEtkL8n1DAHpi5/AOgAckwGQUlMe4+jhL/GMt+GKIks=",
    version = "v0.0.0-20210307095302-262dc9984e00",
)

go_repository(
    name = "com_github_mholt_archiver_v3",
    importpath = "github.com/mholt/archiver/v3",
    sum = "h1:rDjOBX9JSF5BvoJGvjqK479aL70qh9DIpZCl+k7Clwo=",
    version = "v3.5.1",
)

go_repository(
    name = "com_github_mmarkdown_mmark",
    importpath = "github.com/mmarkdown/mmark",
    sum = "h1:vMeUeDzBK3H+/mU0oMVfMuhSXJlIA+DE/DMPQNAj5C4=",
    version = "v2.0.40+incompatible",
)

go_repository(
    name = "com_github_pierrec_lz4_v4",
    importpath = "github.com/pierrec/lz4/v4",
    sum = "h1:qvY3YFXRQE/XB8MlLzJH7mSzBs74eA2gg52YTk6jUPM=",
    version = "v4.1.2",
)

go_repository(
    name = "com_github_pkg_sftp",
    importpath = "github.com/pkg/sftp",
    sum = "h1:VasscCm72135zRysgrJDKsntdmPN+OuU3+nnHYA9wyc=",
    version = "v1.10.1",
)

go_repository(
    name = "com_github_prometheus_exporter_toolkit",
    importpath = "github.com/prometheus/exporter-toolkit",
    sum = "h1:XtYeVeeC5daG4txbc9+mieKq+/AK4gtIBLl9Mulrjnk=",
    version = "v0.7.0",
)

go_repository(
    name = "com_github_sclevine_agouti",
    importpath = "github.com/sclevine/agouti",
    sum = "h1:8IBJS6PWz3uTlMP3YBIR5f+KAldcGuOeFkFbUWfBgK4=",
    version = "v3.0.0+incompatible",
)

go_repository(
    name = "com_github_sourcegraph_go_lsp",
    importpath = "github.com/sourcegraph/go-lsp",
    sum = "h1:afLbh+ltiygTOB37ymZVwKlJwWZn+86syPTbrrOAydY=",
    version = "v0.0.0-20200429204803-219e11d77f5d",
)

go_repository(
    name = "com_github_sourcegraph_jsonrpc2",
    importpath = "github.com/sourcegraph/jsonrpc2",
    sum = "h1:ohJHjZ+PcaLxDUjqk2NC3tIGsVa5bXThe1ZheSXOjuk=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_v2fly_browserbridge",
    importpath = "github.com/v2fly/BrowserBridge",
    sum = "h1:4Yh46CVE3k/lPq6hUbEdbB1u1anRBXLewm3k+L0iOMc=",
    version = "v0.0.0-20210430233438-0570fc1d7d08",
)

go_repository(
    name = "com_github_v2fly_ss_bloomring",
    importpath = "github.com/v2fly/ss-bloomring",
    sum = "h1:5QefA066A1tF8gHIiADmOVOV5LS43gt3ONnlEl3xkwI=",
    version = "v0.0.0-20210312155135-28617310f63e",
)

go_repository(
    name = "com_gitlab_mjwhitta_errors",
    importpath = "gitlab.com/mjwhitta/errors",
    sum = "h1:rivvgv4HCdIJ6Y6kFfPcirzKnBRewkjgsOktR0Uaj8g=",
    version = "v1.0.0",
)

go_repository(
    name = "io_k8s_klog_hack_tools",
    importpath = "k8s.io/klog/hack/tools",
    sum = "h1:u27Xm1of9MTDM1CZW3hg0Vv04ohywEG152G8mpr9n8Y=",
    version = "v0.0.0-20210917071902-331d2323a192",
)

go_repository(
    name = "io_k8s_sigs_json",
    importpath = "sigs.k8s.io/json",
    sum = "h1:fD1pz4yfdADVNfFmcP2aBEtudwUQ1AlLnRBALr33v3s=",
    version = "v0.0.0-20211020170558-c049b76a60c6",
)

go_repository(
    name = "io_k8s_sigs_mdtoc",
    importpath = "sigs.k8s.io/mdtoc",
    sum = "h1:6ECKhQnbetwZBR6R2IeT2LH+1w+2Zsip0iXjikgaXIk=",
    version = "v1.0.1",
)

go_repository(
    name = "org_go4_intern",
    importpath = "go4.org/intern",
    sum = "h1:pq5gAii+wMY+DsJ5r9I6T7CHjHxHlb4d45gChzX2SsI=",
    version = "v0.0.0-20220301175310-a089fc204883",
)

go_repository(
    name = "org_go4_unsafe_assume_no_moving_gc",
    importpath = "go4.org/unsafe/assume-no-moving-gc",
    sum = "h1:Tx9kY6yUkLge/pFG7IEMwDZy6CS2ajFc9TvQdPCW0uA=",
    version = "v0.0.0-20211027215541-db492cf91b37",
)

go_repository(
    name = "com_github_templexxx_cpu",
    importpath = "github.com/templexxx/cpu",
    sum = "h1:pUEZn8JBy/w5yzdYWgx+0m0xL9uk6j4K91C5kOViAzo=",
    version = "v0.0.7",
)

go_repository(
    name = "com_github_templexxx_xorsimd",
    importpath = "github.com/templexxx/xorsimd",
    sum = "h1:iUZcywbOYDRAZUasAs2eSCUW8eobuZDy0I9FJiORkVg=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_xtaci_kcp_go_v5",
    importpath = "github.com/xtaci/kcp-go/v5",
    sum = "h1:Pwn0aoeNSPF9dTS7IgiPXn0HEtaIlVb6y5UKWPsx8bI=",
    version = "v5.6.1",
)

go_repository(
    name = "com_github_caddyserver_caddy",
    importpath = "github.com/caddyserver/caddy",
    sum = "h1:5B1Hs0UF2x2tggr2X9jL2qOZtDXbIWQb9YLbmlxHSuM=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_caddyserver_caddy_v2",
    importpath = "github.com/caddyserver/caddy/v2",
    sum = "h1:eRHzZ4l3X6Ag3kUt8nj5IxATprhqKq/wToP7OHlXWA0=",
    version = "v2.5.0",
)

go_repository(
    name = "com_github_caddyserver_certmagic",
    importpath = "github.com/caddyserver/certmagic",
    sum = "h1:rdSnjcUVJojmL4M0efJ+yHXErrrijS4YYg3FuwRdJkI=",
    version = "v0.16.1",
)

go_repository(
    name = "com_github_cenkalti_backoff_v3",
    importpath = "github.com/cenkalti/backoff/v3",
    sum = "h1:ske+9nBpD9qZsTBoF41nW5L+AIuFBKMeze18XQ3eG1c=",
    version = "v3.0.0",
)

go_repository(
    name = "com_github_decker502_dnspod_go",
    importpath = "github.com/decker502/dnspod-go",
    sum = "h1:6dwhUFCYbC5bgpebLKn7PrI43e/5mn9tpUL9YcYCdTU=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_greenpau_caddy_auth_jwt",
    importpath = "github.com/greenpau/caddy-auth-jwt",
    sum = "h1:CqM0yyAmLbD5vQvl5XaQkyMyq/AZnjk9ZneakxE8Xzo=",
    version = "v1.2.6",
)

go_repository(
    name = "com_github_greenpau_caddy_auth_portal",
    importpath = "github.com/greenpau/caddy-auth-portal",
    replace = "github.com/btwiuse/caddy-auth-portal",
    sum = "h1:P2dbBIM00iljscsSxxKy1gbUCudvFlnbxymj1sdOswk=",
    version = "v1.3.12-0.20210204101408-068c2618b417",
)

go_repository(
    name = "com_github_greenpau_caddy_trace",
    importpath = "github.com/greenpau/caddy-trace",
    sum = "h1:Hi5vBKCzUS9XKjjy5TGzjr+tQef0NvPE+8lzVs3Uvsw=",
    version = "v1.1.8",
)

go_repository(
    name = "com_github_greenpau_go_identity",
    importpath = "github.com/greenpau/go-identity",
    sum = "h1:SnD/LrSIoH35VZk4IskHD1Zm9+cms4Iu+QfjxNNgnU4=",
    version = "v1.0.19",
)

go_repository(
    name = "com_github_greenpau_versioned",
    importpath = "github.com/greenpau/versioned",
    sum = "h1:ICqCoTG8Xv92BV+bKs52d86pDF/e0zhk3LLELsYMpl4=",
    version = "v1.0.23",
)

go_repository(
    name = "com_github_hairyhenderson_caddy_teapot_module",
    importpath = "github.com/hairyhenderson/caddy-teapot-module",
    sum = "h1:POfr7XzYFIUEzAR5/f1+LZ9sN33GqG7XyscWljyFu5Y=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_iamd3vil_caddy_yaml_adapter",
    importpath = "github.com/iamd3vil/caddy_yaml_adapter",
    sum = "h1:5eTxtJy0pyxzY5a1N3bOap7JonTWkuRjrIEs9sK7ciE=",
    version = "v0.0.0-20200503183711-d479c29b475a",
)

go_repository(
    name = "com_github_jimstudt_http_authentication",
    importpath = "github.com/jimstudt/http-authentication",
    sum = "h1:BcF8coBl0QFVhe8vAMMlD+CV8EISiu9MGKLoj6ZEyJA=",
    version = "v0.0.0-20140401203705-3eca13d6893a",
)

go_repository(
    name = "com_github_marten_seemann_chacha20",
    importpath = "github.com/marten-seemann/chacha20",
    sum = "h1:f40vqzzx+3GdOmzQoItkLX5WLvHgPgyYqFFIO5Gh4hQ=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_mholt_caddy_webdav",
    importpath = "github.com/mholt/caddy-webdav",
    sum = "h1:ccjARF3ytF/kfU4CZK/GyFvjFpiOzEmMI5I00eR4zKQ=",
    version = "v0.0.0-20210914165325-f7b67f8ca1e6",
)

go_repository(
    name = "com_github_mholt_certmagic",
    importpath = "github.com/mholt/certmagic",
    sum = "h1:JOUiX9IAZbbgyjNP2GY6v/6lorH+9GkZsc7ktMpGCSo=",
    version = "v0.8.3",
)

go_repository(
    name = "in_gopkg_mcuadros_go_syslog_v2",
    importpath = "gopkg.in/mcuadros/go-syslog.v2",
    sum = "h1:60g8zx1BijSVSgLTzLCW9UC4/+i1Ih9jJ1DR5Tgp9vE=",
    version = "v2.2.1",
)

rules_proto_dependencies()

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

# https://github.com/bazelbuild/rules_docker#go_image
git_repository(
    name = "io_bazel_rules_docker",
    branch = "master",
    remote = "https://github.com/bazelbuild/rules_docker.git",
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
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

go_repository(
    name = "co_honnef_go_tools",
    importpath = "honnef.co/go/tools",
    sum = "h1:UoveltGrhghAA7ePc+e+QYDHXrBps2PqFZiHkGR/xK8=",
    version = "v0.0.1-2020.1.4",
)

go_repository(
    name = "com_github_abbot_go_http_auth",
    importpath = "github.com/abbot/go-http-auth",
    sum = "h1:QjmvZ5gSC7jm3Zg54DqWE/T5m1t2AfDu6QlXJT0EVT0=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_alecthomas_template",
    importpath = "github.com/alecthomas/template",
    sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
    version = "v0.0.0-20190718012654-fb15b899a751",
)

go_repository(
    name = "com_github_alecthomas_units",
    importpath = "github.com/alecthomas/units",
    sum = "h1:AUNCr9CiJuwrRYS3XieqF+Z9B9gNxo/eANAJCF2eiN4=",
    version = "v0.0.0-20210208195552-ff826a37aa15",
)

go_repository(
    name = "com_github_alexpantyukhin_go_pattern_match",
    importpath = "github.com/alexpantyukhin/go-pattern-match",
    sum = "h1:Bq65TLHVwjKAnS+kaPOi3lyOpw0rnNmKcw7EsY4WzuM=",
    version = "v0.0.0-20200628201436-c57d5ad3f2c5",
)

go_repository(
    name = "com_github_beevik_ntp",
    importpath = "github.com/beevik/ntp",
    sum = "h1:xzVrPrE4ziasFXgBVBZJDP0Wg/KpMwk2KHJ4Ba8GrDw=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_beorn7_perks",
    importpath = "github.com/beorn7/perks",
    sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_btwiuse_asciitransport",
    importpath = "github.com/btwiuse/asciitransport",
    sum = "h1:LmuQIwBS5Lkqaw4uxHR57caequKBx/+LfDhu1JyzN2U=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_btwiuse_gods",
    importpath = "github.com/btwiuse/gods",
    sum = "h1:RamJoBgLnq2tsdeXwKc/+ytJ30rftG+MBDreHGFWpSw=",
    version = "v0.0.0-20190414062120-7e7cf0aebbb0",
)

go_repository(
    name = "com_github_btwiuse_pretty",
    importpath = "github.com/btwiuse/pretty",
    sum = "h1:9gByUaL1XTCtnT0m0+9UrOie1QOQiWJ3FAHeg3C/+Es=",
    version = "v0.0.0-20201017182805-a72bdf3093a3",
)

go_repository(
    name = "com_github_btwiuse_wetty",
    importpath = "github.com/btwiuse/wetty",
    sum = "h1:MSMG+C1wR+bgG5xB2TRZK+5yxYiOmGwe44c/3gpaGT0=",
    version = "v0.0.36",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:ksErzDEI1khOiGPgpwuI7x2ebx/uXQNw7xJpn9Eq1+I=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    importpath = "github.com/census-instrumentation/opencensus-proto",
    sum = "h1:t/LhUZLVitR1Ow2YOnduCsavhwFUklBMoGVYUCqmCqk=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_cespare_xxhash_v2",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:YRXhKfTDauu4ajMg1TPgFO5jnlC2HCbmLXMcTG5cbYE=",
    version = "v2.1.2",
)

go_repository(
    name = "com_github_cirocosta_asciinema_edit",
    importpath = "github.com/cirocosta/asciinema-edit",
    sum = "h1:Bc9sl3YDXJNoo8IM9YFavXKXyrrLm2lbOTWRSQ7fUoo=",
    version = "v0.0.0-20190130154215-1c0971ae232a",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_coreos_go_systemd",
    importpath = "github.com/coreos/go-systemd",
    sum = "h1:iW4rZ826su+pqaw19uhpSCzhj44qo35pNgKFGqzDKkU=",
    version = "v0.0.0-20191104093116-d3cd4ed1dbcf",
)

go_repository(
    name = "com_github_creack_pty",
    importpath = "github.com/creack/pty",
    sum = "h1:n56/Zwd5o6whRC5PMGretI4IdRLlmBXYNjScPaBgsbY=",
    version = "v1.1.18",
)

go_repository(
    name = "com_github_data_dog_go_sqlmock",
    importpath = "github.com/DATA-DOG/go-sqlmock",
    sum = "h1:CWUqKXe0s8A2z6qCgkP4Kru7wC11YoAnoupUKFDnH08=",
    version = "v1.3.3",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_denisbrodbeck_machineid",
    importpath = "github.com/denisbrodbeck/machineid",
    sum = "h1:geKr9qtkB876mXguW2X6TU4ZynleN6ezuMSRhl4D7AQ=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_docker_docker",
    importpath = "github.com/docker/docker",
    sum = "h1:+T9/PRYWNDo5SZl5qS1r9Mo/0Q8AwxKKPtu9S1yxM0w=",
    version = "v20.10.14+incompatible",
)

go_repository(
    name = "com_github_ema_qdisc",
    importpath = "github.com/ema/qdisc",
    sum = "h1:0GHzegkDz/zSrt+Zph1OueNImPdUxoToypnkhhRYTjI=",
    version = "v0.0.0-20200603082823-62d0308e3e00",
)

go_repository(
    name = "com_github_emirpasic_gods",
    importpath = "github.com/emirpasic/gods",
    sum = "h1:QAUIPSaCu4G+POclxeqb3F+WPpdKqFGlw36+yOzGlrg=",
    version = "v1.12.0",
)

go_repository(
    name = "com_github_envoyproxy_go_control_plane",
    importpath = "github.com/envoyproxy/go-control-plane",
    sum = "h1:xvqufLtNVwAhN8NMyWklVgxnWohi+wtMGQMhtxexlm0=",
    version = "v0.10.2-0.20220325020618-49ff273808a1",
)

go_repository(
    name = "com_github_envoyproxy_protoc_gen_validate",
    importpath = "github.com/envoyproxy/protoc-gen-validate",
    sum = "h1:bV5JGEB1ouEzZa0hgVDFFiClrUEuGWRaAc/3mxR2QK0=",
    version = "v0.3.0-java",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:mZcQUHVQUQWoPXXtuf9yuEXKudkV2sx1E06UadKWpgI=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_gdamore_encoding",
    importpath = "github.com/gdamore/encoding",
    sum = "h1:+7OoQ1Bc6eTm5niUzBa0Ctsh6JbMW6Ra+YNuAtDBdko=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_gdamore_tcell",
    importpath = "github.com/gdamore/tcell",
    sum = "h1:vUnHwJRvcPQa3tzi+0QI4U9JINXYJlOz9yiaiPQ2wMU=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:Mn26/9ZMNWSw9C9ERFA1PUxfmGpolnw2v0bKOREu5ew=",
    version = "v1.0.1-0.20190212211648-25d852aebe32",
)

go_repository(
    name = "com_github_go_kit_kit",
    importpath = "github.com/go-kit/kit",
    sum = "h1:dXFJfIHVvUcpSgDOV+Ne6t7jXri8Tfv2uOLHUZ2XNuo=",
    version = "v0.10.0",
)

go_repository(
    name = "com_github_go_logfmt_logfmt",
    importpath = "github.com/go-logfmt/logfmt",
    sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_go_stack_stack",
    importpath = "github.com/go-stack/stack",
    sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_gobwas_httphead",
    importpath = "github.com/gobwas/httphead",
    sum = "h1:s+21KNqlpePfkah2I+gwHF8xmJWRjooY+5248k6m4A0=",
    version = "v0.0.0-20180130184737-2c6c146eadee",
)

go_repository(
    name = "com_github_gobwas_pool",
    importpath = "github.com/gobwas/pool",
    sum = "h1:QEmUOlnSjWtnpRGHF3SauEiOsy82Cup83Vf2LcMlnc8=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_gobwas_ws",
    importpath = "github.com/gobwas/ws",
    sum = "h1:CoAavW/wd/kulfZmSIBt6p24n4j7tHgNVCjsfHVNUbo=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_godbus_dbus",
    importpath = "github.com/godbus/dbus",
    sum = "h1:BWhy2j3IXJhjCbC68FptL43tDKIq8FladmaTs3Xs7Z8=",
    version = "v0.0.0-20190422162347-ade71ed3457e",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    sum = "h1:ErTB+efbowRARo13NNdxyJji2egdxLGQhRaY+DUumQc=",
    version = "v1.6.0",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:ROPKBNFfQgOUMifHyP+KYbvpjbdoFNs+aK7DXlji0Tw=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:81/ik6ipDQS2aGcBfIN5dHDB36BwrStyeAQquSYCV4o=",
    version = "v0.5.7",
)

go_repository(
    name = "com_github_google_gofuzz",
    importpath = "github.com/google/gofuzz",
    sum = "h1:Hsa8mG0dQ46ij8Sl2AYJDUv1oA9/d6Vk+3LG99Oe02g=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_google_renameio",
    importpath = "github.com/google/renameio",
    sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_gorilla_handlers",
    importpath = "github.com/gorilla/handlers",
    sum = "h1:9lRY6j8DEeeBT10CvO9hGW0gmky0BprnvDI5vfhUHH4=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_gorilla_mux",
    importpath = "github.com/gorilla/mux",
    sum = "h1:i40aqfkR1h2SlN9hojwV5ZA91wcXFOvkdNIeFDP5koI=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    sum = "h1:PPwGk2jz7EePpoHN/+ClbZu8SPxiqlu12wZP/3sWmnc=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_hodgesds_perf_utils",
    importpath = "github.com/hodgesds/perf-utils",
    sum = "h1:onWrAGy6RYr7938qNXtSsTr54K4BLx8Hh3EXAr+xy+U=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    sum = "h1:Ysi1UhrSyBltF8f+3RAt4UaqHc+53JJ0jyl0pY0sfck=",
    version = "v1.0.1-0.20180514194441-a1dbeea552b7",
)

go_repository(
    name = "com_github_jsimonetti_rtnetlink",
    importpath = "github.com/jsimonetti/rtnetlink",
    sum = "h1:N527AHMa793TP5z5GNAn/VLPzlc0ewzWdeP/25gDfgQ=",
    version = "v0.0.0-20211022192332-93da33804786",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
    version = "v1.1.12",
)

go_repository(
    name = "com_github_julienschmidt_httprouter",
    importpath = "github.com/julienschmidt/httprouter",
    sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_kr_logfmt",
    importpath = "github.com/kr/logfmt",
    sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
    version = "v0.0.0-20140226030751-b84e30acd515",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:WgNl7dwNpEZ6jJ9k1snq4pZsg7DOEN8hP9Xw0Tsjwk0=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:AkaSdXYQOWeaO3neb8EM634ahkXXe3jYbVh/F9lq+GI=",
    version = "v1.1.8",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_liamg_aminal",
    importpath = "github.com/liamg/aminal",
    sum = "h1:0mLmvBYQUIX3fO9zoOY/PBEb+HW5Jcu3aexD+ioEtc0=",
    version = "v0.9.0",
)

go_repository(
    name = "com_github_lucasb_eyer_go_colorful",
    importpath = "github.com/lucasb-eyer/go-colorful",
    sum = "h1:QIbQXiugsb+q10B+MI+7DI1oQLdmnep86tWFlaaUAac=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_lufia_iostat",
    importpath = "github.com/lufia/iostat",
    sum = "h1:Botv3++V0FnQyhRlSt82DHUBv7XlxFtaNInpLq1jrAU=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_lukesampson_figlet",
    importpath = "github.com/lukesampson/figlet",
    sum = "h1:UtyD+eBVdLYSj5/pjfSR6mtnzMgIiOVcFT024G2l4CY=",
    version = "v0.0.0-20190211215653-8a3ef4a6ac42",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:yVuAays6BHfxijgZPzw+3Zlu5yQgKGP2/hcQbHb7S9Y=",
    version = "v0.0.14",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    sum = "h1:lTGmDsbAYt5DmK6OnoV7EuIF1wEIFAcxld6ypU4OSgU=",
    version = "v0.0.13",
)

go_repository(
    name = "com_github_mattn_go_shellwords",
    importpath = "github.com/mattn/go-shellwords",
    sum = "h1:M2zGm7EW6UQJvDeQxo4T51eKPurbeFbe8WtebGE2xrk=",
    version = "v1.0.12",
)

go_repository(
    name = "com_github_mattn_go_xmlrpc",
    importpath = "github.com/mattn/go-xmlrpc",
    sum = "h1:Y6WEMLEsqs3RviBrAa1/7qmbGB7DVD3brZIbqMbQdGY=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    sum = "h1:I0XW9+e1XWDxdcEniV4rQAIOPUGDq67JSCiRCgGCZLI=",
    version = "v1.0.2-0.20181231171920-c182affec369",
)

go_repository(
    name = "com_github_maxris_w32",
    importpath = "github.com/MaxRis/w32",
    sum = "h1:5GK1kUZ3zfsAKqry3g+4mLoSjdDRLBeZ/93sjJKK8kA=",
    version = "v0.0.0-20180517000239-4f5cfb03fabf",
)

go_repository(
    name = "com_github_mbndr_figlet4go",
    importpath = "github.com/mbndr/figlet4go",
    sum = "h1:mQncVDBpKkAecPcH2IMGpKUQYhwowlafQbfkz2QFqkc=",
    version = "v0.0.0-20190224160619-d6cef5b186ea",
)

go_repository(
    name = "com_github_mdlayher_genetlink",
    importpath = "github.com/mdlayher/genetlink",
    sum = "h1:OoHN1OdyEIkScEmRgxLEe2M9U8ClMytqA5niynLtfj0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mdlayher_netlink",
    importpath = "github.com/mdlayher/netlink",
    sum = "h1:I154BCU+mKlIf7BgcAJB2r7QjveNPty6uNY1g9ChVfI=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_mdlayher_wifi",
    importpath = "github.com/mdlayher/wifi",
    sum = "h1:50p1vPNK43pzCVX10+5MmiOerbBzC1vR6+sLB3FZewE=",
    version = "v0.0.0-20200527114002-84f0b9457fdd",
)

go_repository(
    name = "com_github_mdp_qrterminal",
    importpath = "github.com/mdp/qrterminal",
    sum = "h1:07+fzVDlPuBlXS8tB0ktTAyf+Lp1j2+2zK3fBOL5b7c=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_miekg_dns",
    importpath = "github.com/miekg/dns",
    sum = "h1:Ucfr7IIVyMBz4lRE8qmGUuZ4Wt3/ZGu9hmcMT3Uu4tQ=",
    version = "v1.1.48",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    importpath = "github.com/modern-go/concurrent",
    sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    importpath = "github.com/mwitkow/go-conntrack",
    sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
    version = "v0.0.0-20190716064945-2f068394615f",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    sum = "h1:8xi0RTUf59SOSfEtZMvwTvXYMzG4gV23XVHOZiXNtnE=",
    version = "v1.16.5",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    sum = "h1:7lLHu94wT9Ij0o6EWWclhu0aOh32VxhkwEJvzuWPeak=",
    version = "v1.13.0",
)

go_repository(
    name = "com_github_patrickmn_go_cache",
    importpath = "github.com/patrickmn/go-cache",
    sum = "h1:HRMgzkcYKYpi3C8ajMPV8OFXaaRUnok+kx1WdO15EQc=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    sum = "h1:ZiaPsmm9uiBeaSMRznKsCDNtPCS0T3JVDGF+06gjBzk=",
    version = "v1.12.1",
)

go_repository(
    name = "com_github_prometheus_client_model",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:uq5h0d+GuxiXLJLNABMgp2qUWDPiLvgCzz2dUR+/W/M=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_prometheus_common",
    importpath = "github.com/prometheus/common",
    sum = "h1:RBmGO9d/FVjqHT0yUGQwBJhkwKV+wPCn7KGpvfab0uE=",
    version = "v0.34.0",
)

go_repository(
    name = "com_github_prometheus_node_exporter",
    importpath = "github.com/prometheus/node_exporter",
    sum = "h1:pdJ5824noBu1O2a3a96uiWDH9Ns2ZTE9OxBlPy9Deik=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_prometheus_procfs",
    importpath = "github.com/prometheus/procfs",
    sum = "h1:ncXqc93eJV1Ncr3f6GA3MrIDNkNHvcPonRC2QgZaVkQ=",
    version = "v0.7.4-0.20211011103944-1a7a2bd3279f",
)

go_repository(
    name = "com_github_pupapaik_sysinfo",
    importpath = "github.com/pupapaik/sysinfo",
    sum = "h1:sfRZ+ctR2rYU/CfsaIdAqvxOxiT4dCKlCZRFQwTap3g=",
    version = "v0.0.0-20200106202926-c17dea004cd4",
)

go_repository(
    name = "com_github_riobard_go_shadowsocks2",
    importpath = "github.com/riobard/go-shadowsocks2",
    sum = "h1:nplpXf7LiL7JCovIXiOExF84ZtVZbZuRyc14UXUZblw=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    importpath = "github.com/rogpeppe/go-internal",
    sum = "h1:/FiVV8dS/e+YqF2JvO3yXRFbBLTIuSDkuC7aBOAvL+k=",
    version = "v1.6.1",
)

go_repository(
    name = "com_github_rs_cors",
    importpath = "github.com/rs/cors",
    sum = "h1:KCooALfAYGs415Cwu5ABvv9n9509fSiG5SQJn/AQo4U=",
    version = "v1.8.2",
)

go_repository(
    name = "com_github_shurcool_sanitized_anchor_name",
    importpath = "github.com/shurcooL/sanitized_anchor_name",
    sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_siebenmann_go_kstat",
    importpath = "github.com/siebenmann/go-kstat",
    sum = "h1:GfSdC6wKfTGcgCS7BtzF5694Amne1pGCSTY252WhlEY=",
    version = "v0.0.0-20210513183136-173c9b0a9973",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:dJKuHgqk1NNQlqoA6BTlM1Wf9DOH3NBjQyu0h9+AZZE=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_soundcloud_go_runit",
    importpath = "github.com/soundcloud/go-runit",
    sum = "h1:os5OBNhwOwybXZMNLqT96XqtjdTtwRFw2w08uluvNeI=",
    version = "v0.0.0-20150630195641-06ad41a06c4a",
)

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:Hbg2NidpLE8veEBkEZTL3CvlkUIVzuU9jDplZO54c48=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:5TQK59W5E3v0r2duFAb7P95B6hEeOyEnHRa8MjYSMTY=",
    version = "v1.7.1",
)

go_repository(
    name = "com_github_txthinking_gotun2socks",
    importpath = "github.com/txthinking/gotun2socks",
    sum = "h1:0S59/gOPf/wFIGztKRHfImRwlCi9vpsahVnywWv82LA=",
    version = "v0.0.0-20180829122610-35016fdae05e",
)

go_repository(
    name = "com_github_txthinking_socks5",
    importpath = "github.com/txthinking/socks5",
    sum = "h1:7x8pJcBTdKTBpQbRjZZc9o6CDquXBbvm9UIrR6ZSRJ4=",
    version = "v0.0.0-20210716140126-fa1f52a8f2da",
)

go_repository(
    name = "com_github_txthinking_x",
    importpath = "github.com/txthinking/x",
    sum = "h1:gMWxZxBFRAXqoGkwkYlPX2zvyyKNWJpxOxCrjqJkm5A=",
    version = "v0.0.0-20210326105829-476fab902fbe",
)

go_repository(
    name = "com_github_urfave_negroni",
    importpath = "github.com/urfave/negroni",
    sum = "h1:kIimOitoypq34K7TG7DUaJ9kq/N4Ofuwi1sjz0KipXc=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_vividcortex_gohistogram",
    importpath = "github.com/VividCortex/gohistogram",
    sum = "h1:6+hBz+qvs0JOrrNhhmR7lFxo5sINxBCGXrdtl/UvroE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_vojtechvitek_yaml_cli",
    importpath = "github.com/VojtechVitek/yaml-cli",
    sum = "h1:jsovPZX2Yau3+g0/xpwFXkCIfGRehtyecM0vB+Y6pFo=",
    version = "v0.0.5",
)

go_repository(
    name = "com_github_yrpc_rpheap",
    importpath = "github.com/yrpc/rpheap",
    sum = "h1:dyIrDf2ZEI0SUGZzcK3GrMp+fBcdaZk7tYvDz8sVbrM=",
    version = "v0.0.0-20191222053847-9002d7e5a1a1",
)

go_repository(
    name = "com_github_yrpc_util",
    importpath = "github.com/yrpc/util",
    sum = "h1:w3AUE/DLL+g6BKLgJbOfjVrxqabaA445sYSBatEAoBo=",
    version = "v0.0.0-20191229105456-04e44c1bb858",
)

go_repository(
    name = "com_github_yrpc_yrpc",
    importpath = "github.com/yrpc/yrpc",
    sum = "h1:C/zMo1vpa3Kr0F68ZCfxP9ddwwqbSaLj1cGYcZb6fp0=",
    version = "v0.0.0-20191231114812-451503bf48c2",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    sum = "h1:t9Iw5QH5v4XtlEQaCtUY7x6sCABps8sW0acw7e2WQ6Y=",
    version = "v0.100.2",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
    sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
    version = "v2.2.6",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
    version = "v1.0.0-20201130134442-10cb98267c6c",
)

go_repository(
    name = "in_gopkg_errgo_v2",
    importpath = "gopkg.in/errgo.v2",
    sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
    version = "v2.1.0",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
    version = "v1.4.7",
)

go_repository(
    name = "in_gopkg_russross_blackfriday_v2",
    importpath = "gopkg.in/russross/blackfriday.v2",
    sum = "h1:+FlnIV8DSQnT7NZ43hcVKcdJdzZoeCmJj4Ql8gq5keA=",
    version = "v2.0.0",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    importpath = "gopkg.in/tomb.v1",
    sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
    version = "v1.0.0-20141024135613-dd632973f1e7",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
    version = "v2.4.0",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:h8qDotaEPuJATrMmW04NCwg7v22aHH28wwpauUhK9Oo=",
    version = "v3.0.0-20210107192922-496545a6307b",
)

go_repository(
    name = "io_nhooyr_websocket",
    importpath = "nhooyr.io/websocket",
    sum = "h1:usjR2uOr/zjjkVMy0lW+PPohFok7PCow5sDjLgX4P4g=",
    version = "v1.8.7",
)

go_repository(
    name = "io_rsc_qr",
    importpath = "rsc.io/qr",
    sum = "h1:6vBLea5/NRMVTz8V66gipeLycZMl/+UlFmk8DvqQ6WY=",
    version = "v0.2.0",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:FZR1q0exgwxzPzp/aF+VccGrSfxfPpkBqjIIEq3ru6c=",
    version = "v1.6.7",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:SVYXkUz2yZS9FWb2Gm8ivSlbNQzL2Z/NpPKE3RG2jWk=",
    version = "v0.0.0-20220222213610-43724f9ea8cf",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:oCjezcn6g6A75TGoKYBPgKmVBLexhYLM6MebdrPApP8=",
    version = "v1.46.0",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:OeJjE6G4dgCY4PIXvIRQbE8+RX+uXZyGhUy/ksMGJoc=",
    version = "v0.0.0-20220427172511-eb4f295cb31f",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:GrkO5AtFUU9U/1f5ctbIBXtBGeSJbWwIYfIsTcFMaX4=",
    version = "v0.0.0-20210220032938-85be41e4509f",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:VLliZ0d+/avPrXXH+OakdXhpJuEoBZuwh1m2j7U6Iug=",
    version = "v0.0.0-20210508222113-6edffad5e616",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:kQgndtyPBW/JIYERgdxfwMYh3AVStj88WQTlNDi2a+o=",
    version = "v0.6.0-dev.0.20220106191415-9b9b3d81d5e3",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:HVyaeDAYux4pnY+D/SiwmLOR36ewZ4iGQIIrtnuCjFA=",
    version = "v0.0.0-20220425223048-2871e0cb64e4",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:clP8eMhB30EHdc0bd2Twtq6kgU7yl5ub2cQLSdrv1Dg=",
    version = "v0.0.0-20220223155221-ee480838109b",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:5KslGYwFpkhGh+Q16bwMP3cOontH8FOep7tGV86Y7SQ=",
    version = "v0.0.0-20210220032951-036812b2e83c",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:Js08h5hqB5xyWR789+QqueR6sDE8mk+YvpETZ+F6X9Y=",
    version = "v0.0.0-20220429233432-b5fbb4746d32",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:NXqSWXSRUSCaFuvitrWtU169I3876zRTalMRbfd6LL0=",
    version = "v0.3.8-0.20211004125949-5bd84dd9b33b",
)

go_repository(
    name = "org_golang_x_time",
    importpath = "golang.org/x/time",
    sum = "h1:7zkz7BUtwNFFqcowJ+RIgu2MaV/MapERkDIy+mwPyjs=",
    version = "v0.0.0-20210723032227-1f47c861a9ac",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:QjFRCZxdOhBJ/UNgnBZLbNV13DlbnK0quyivTnXJM20=",
    version = "v0.1.10",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_repository(
    name = "org_modernc_httpfs",
    importpath = "modernc.org/httpfs",
    sum = "h1:AAgIpFZRXuYnkjftxTAZwMIiwEqAfk8aVB2/oA6nAeM=",
    version = "v1.0.6",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    sum = "h1:ECmE8Bn/WFTYwEW/bpKD3M8VtR/zQVbavAoalC1PYyE=",
    version = "v1.9.0",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    sum = "h1:zaiO/rmgFjbmCXdSYJWQcdvOCsthmdaHfr3Gm2Kx4Ec=",
    version = "v1.7.0",
)

go_repository(
    name = "org_uber_go_ratelimit",
    importpath = "go.uber.org/ratelimit",
    sum = "h1:d9qaMM+ODpCq+9We41//fu/sHsTnXcrqd1en3x+GKy4=",
    version = "v0.0.0-20180316092928-c15da0234277",
)

go_repository(
    name = "org_uber_go_tools",
    importpath = "go.uber.org/tools",
    sum = "h1:0mgffUl7nfd+FpvXMVz4IDEaUSmT1ysygQC7qYo7sG4=",
    version = "v0.0.0-20190618225709-2cfd321de3ee",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    sum = "h1:WefMeulhovoZ2sYXz7st6K0sLj7bBhpiFaud4r4zST8=",
    version = "v1.21.0",
)

go_repository(
    name = "tools_gotest",
    importpath = "gotest.tools",
    sum = "h1:VsBPFP1AI068pPrMxtb/S8Zkgf9xEmTLJjfM+P5UIEo=",
    version = "v2.2.0+incompatible",
)

go_repository(
    name = "com_github_afex_hystrix_go",
    importpath = "github.com/afex/hystrix-go",
    sum = "h1:rFw4nCn9iMW+Vajsk51NtYIcwSTkXr+JGrMd36kTDJw=",
    version = "v0.0.0-20180502004556-fa1af6a1f4f5",
)

go_repository(
    name = "com_github_apache_thrift",
    importpath = "github.com/apache/thrift",
    sum = "h1:5hryIiq9gtn+MiLVn0wP37kb/uTeRZgN08WoCsAhIhI=",
    version = "v0.13.0",
)

go_repository(
    name = "com_github_armon_circbuf",
    importpath = "github.com/armon/circbuf",
    sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
    version = "v0.0.0-20150827004946-bbbad097214e",
)

go_repository(
    name = "com_github_armon_go_metrics",
    importpath = "github.com/armon/go-metrics",
    sum = "h1:O2sNqxBdvq8Eq5xmzljcYzAORli6RWCvEym4cJf9m18=",
    version = "v0.3.9",
)

go_repository(
    name = "com_github_armon_go_radix",
    importpath = "github.com/armon/go-radix",
    sum = "h1:F4z6KzEeeQIMeLFa97iZU6vupzoecKdU5TX24SNppXI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_aryann_difflib",
    importpath = "github.com/aryann/difflib",
    sum = "h1:uUXgbcPDK3KpW29o4iy7GtuappbWT0l5NaMo9H9pJDw=",
    version = "v0.0.0-20210328193216-ff5ff6dc229b",
)

go_repository(
    name = "com_github_aws_aws_lambda_go",
    importpath = "github.com/aws/aws-lambda-go",
    sum = "h1:SuCy7H3NLyp+1Mrfp+m80jcbi9KYWAs9/BXwppwRDzY=",
    version = "v1.13.3",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    sum = "h1:zJnJ8Y964DjyRE55UVoMKgOG4w5i88LpN6xSpBX7z84=",
    version = "v1.41.14",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2",
    importpath = "github.com/aws/aws-sdk-go-v2",
    sum = "h1:GzvOVAdTbWxhEMRK4FfiblkGverOkAT0UodDxC1jHQM=",
    version = "v1.11.1",
)

go_repository(
    name = "com_github_bgentry_speakeasy",
    importpath = "github.com/bgentry/speakeasy",
    sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_casbin_casbin_v2",
    importpath = "github.com/casbin/casbin/v2",
    sum = "h1:JApbUmymvG33xIIYJ4G4ijj20tlZ9b8LY4ByTe+Oz+M=",
    version = "v2.8.6",
)

go_repository(
    name = "com_github_cenkalti_backoff",
    importpath = "github.com/cenkalti/backoff",
    sum = "h1:tNowT99t7UNflLxfYYSlKYsBpXdEet03Pg2g16Swow4=",
    version = "v2.2.1+incompatible",
)

go_repository(
    name = "com_github_clbanning_x2j",
    importpath = "github.com/clbanning/x2j",
    sum = "h1:EdRZT3IeKQmfCSrgo8SZ8V3MEnskuJP0wCYNpe+aiXo=",
    version = "v0.0.0-20191024224557-825249438eec",
)

go_repository(
    name = "com_github_cockroachdb_datadriven",
    importpath = "github.com/cockroachdb/datadriven",
    sum = "h1:xD/lrqdvwsc+O2bjSSi3YqY73Ke3LAiSCx49aCesA0E=",
    version = "v0.0.0-20200714090401-bf6692d28da5",
)

go_repository(
    name = "com_github_codahale_hdrhistogram",
    importpath = "github.com/codahale/hdrhistogram",
    sum = "h1:qMd81Ts1T2OTKmB4acZcyKaMtRnY5Y44NuXGX2GFJ1w=",
    version = "v0.0.0-20161010025455-3a0bb77429bd",
)

go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_coreos_pkg",
    importpath = "github.com/coreos/pkg",
    sum = "h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
    version = "v0.0.0-20180928190104-399ea9e2e55f",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man_v2",
    importpath = "github.com/cpuguy83/go-md2man/v2",
    sum = "h1:EoUDS0afbrsXAZ9YQ9jdu/mZ2sXgT1/2yyNng4PGlyM=",
    version = "v2.0.0",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    importpath = "github.com/dgrijalva/jwt-go",
    sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
    version = "v3.2.0+incompatible",
)

go_repository(
    name = "com_github_dustin_go_humanize",
    importpath = "github.com/dustin/go-humanize",
    sum = "h1:opbrjaN/L8gg6Xh5D04Tem+8xVcz6ajZlGCs49mQgyg=",
    version = "v1.0.1-0.20200219035652-afde56e7acac",
)

go_repository(
    name = "com_github_eapache_go_resiliency",
    importpath = "github.com/eapache/go-resiliency",
    sum = "h1:v7g92e/KSN71Rq7vSThKaWIq68fL4YHvWyiUKorFR1Q=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_eapache_go_xerial_snappy",
    importpath = "github.com/eapache/go-xerial-snappy",
    sum = "h1:YEetp8/yCZMuEPMUDHG0CW/brkkEp8mzqk2+ODEitlw=",
    version = "v0.0.0-20180814174437-776d5712da21",
)

go_repository(
    name = "com_github_eapache_queue",
    importpath = "github.com/eapache/queue",
    sum = "h1:YOEu7KNc61ntiQlcEeUIoDTJ2o8mQznoNvUhiigpIqc=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_edsrzf_mmap_go",
    importpath = "github.com/edsrzf/mmap-go",
    sum = "h1:CEBF7HpRnUCSJgGUb5h1Gm7e3VkmVDrR8lvWVLtrOFw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_fatih_color",
    importpath = "github.com/fatih/color",
    sum = "h1:s36xzo75JdqLaaWoiEHk767eHiwo0598uUxyfiPkDsg=",
    version = "v1.10.0",
)

go_repository(
    name = "com_github_franela_goblin",
    importpath = "github.com/franela/goblin",
    sum = "h1:gb2Z18BhTPJPpLQWj4T+rfKHYCHxRHCtRxhKKjRidVw=",
    version = "v0.0.0-20200105215937-c9ffbefa60db",
)

go_repository(
    name = "com_github_franela_goreq",
    importpath = "github.com/franela/goreq",
    sum = "h1:a9ENSRDFBUPkJ5lCgVZh26+ZbGyoVJG7yb5SSzF5H54=",
    version = "v0.0.0-20171204163338-bcd34c9993f8",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    sum = "h1:BCTh4TKNUYmOmMUcQ3IipzF5prigylS7XXjEkfCHuOE=",
    version = "v1.6.0",
)

go_repository(
    name = "com_github_gogo_googleapis",
    importpath = "github.com/gogo/googleapis",
    sum = "h1:zgVt4UpGxcqVOw97aRGxT4svlcmdK35fynLNctY32zI=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_golang_groupcache",
    importpath = "github.com/golang/groupcache",
    sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
    version = "v0.0.0-20210331224755-41bb18bfe9da",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_google_btree",
    importpath = "github.com/google/btree",
    sum = "h1:gK4Kx5IaGY9CD5sPJ36FHiBJ6ZXl0kilRiiCj+jdYp4=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:ATVz3rDvK4xX0nHx57zYSHRVIK/+lFwln9KJr8wvuk0=",
    version = "v0.0.0-20210420193930-a4630ec28c79",
)

go_repository(
    name = "com_github_gorilla_context",
    importpath = "github.com/gorilla/context",
    sum = "h1:AWwleXJkX/nhcU9bZSnZoi3h/qGYqQAGhq6zZe/aQW8=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_middleware",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
    sum = "h1:+9834+KizmvFV7pXQGSXQTsaWhq2GjuNUt0aUU0YBYw=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_prometheus",
    importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
    sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_hashicorp_consul_api",
    importpath = "github.com/hashicorp/consul/api",
    sum = "h1:HXNYlRkkM/t+Y/Yhxtwcy02dlYwIaoxzvxPnS+cqy78=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_hashicorp_consul_sdk",
    importpath = "github.com/hashicorp/consul/sdk",
    sum = "h1:UOxjlb4xVNF93jak1mzzoBatyFju9nrkxpVwIp/QqxQ=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_hashicorp_errwrap",
    importpath = "github.com/hashicorp/errwrap",
    sum = "h1:OxrOeh75EUXMY8TBjag2fzXGZ40LB6IKw45YeGUDY2I=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_hashicorp_go_cleanhttp",
    importpath = "github.com/hashicorp/go-cleanhttp",
    sum = "h1:035FKYIWjmULyFRBKPs8TBQoi0x6d9G4xc9neXJWAZQ=",
    version = "v0.5.2",
)

go_repository(
    name = "com_github_hashicorp_go_immutable_radix",
    importpath = "github.com/hashicorp/go-immutable-radix",
    sum = "h1:DKHmCUm2hRBK510BaiZlwvpD40f8bJFeZnpfm2KLowc=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_hashicorp_go_msgpack",
    importpath = "github.com/hashicorp/go-msgpack",
    sum = "h1:i9R9JSrqIz0QVLz3sz+i3YJdT7TTSLcfLLzJi9aZTuI=",
    version = "v0.5.5",
)

go_repository(
    name = "com_github_hashicorp_go_multierror",
    importpath = "github.com/hashicorp/go-multierror",
    sum = "h1:H5DkEtf6CXdFp0N0Em5UCwQpXMWke8IA0+lD48awMYo=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_hashicorp_go_net",
    importpath = "github.com/hashicorp/go.net",
    sum = "h1:sNCoNyDEvN1xa+X0baata4RdcpKwcMS6DH+xwfqPgjw=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_hashicorp_go_rootcerts",
    importpath = "github.com/hashicorp/go-rootcerts",
    sum = "h1:jzhAVGtqPKbwpyCPELlgNWhE1znq+qwJtW5Oi2viEzc=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_hashicorp_go_sockaddr",
    importpath = "github.com/hashicorp/go-sockaddr",
    sum = "h1:ztczhD1jLxIRjVejw8gFomI1BQZOe2WoVOu0SyteCQc=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_hashicorp_go_syslog",
    importpath = "github.com/hashicorp/go-syslog",
    sum = "h1:KaodqZuhUoZereWVIYmpUgZysurB1kBLX2j0MwMrUAE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_uuid",
    importpath = "github.com/hashicorp/go-uuid",
    sum = "h1:cfejS+Tpcp13yd5nYHWDI6qVCny6wyX2Mt5SGur2IGE=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_hashicorp_go_version",
    importpath = "github.com/hashicorp/go-version",
    sum = "h1:3vNe/fWF5CBgRIguda1meWhsZHy3m8gCJ5wx+dIzX/E=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    importpath = "github.com/hashicorp/golang-lru",
    sum = "h1:YDjusn29QI/Das2iO9M0BHnIbxPeyuCHsjMW+lJfyTc=",
    version = "v0.5.4",
)

go_repository(
    name = "com_github_hashicorp_logutils",
    importpath = "github.com/hashicorp/logutils",
    sum = "h1:dLEQVugN8vlakKOUE3ihGLTZJRB4j+M2cdTm/ORI65Y=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_mdns",
    importpath = "github.com/hashicorp/mdns",
    sum = "h1:WhIgCr5a7AaVH6jPUwjtRuuE7/RDufnUvzIr48smyxs=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_memberlist",
    importpath = "github.com/hashicorp/memberlist",
    sum = "h1:gkyML/r71w3FL8gUi74Vk76avkj/9lYAY9lvg0OcoGs=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_hashicorp_serf",
    importpath = "github.com/hashicorp/serf",
    sum = "h1:MWYcmct5EtKz0efYooPcL0yNkem+7kWxqXDi/UIh+8k=",
    version = "v0.8.3",
)

go_repository(
    name = "com_github_hudl_fargo",
    importpath = "github.com/hudl/fargo",
    sum = "h1:0U6+BtN6LhaYuTnIJq4Wyq5cpn6O2kWrxAtcqBmYY6w=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    importpath = "github.com/inconshreveable/mousetrap",
    sum = "h1:Z8tu5sraLXCXIcARxBp/8cbvlwVa7Z1NHg9XEKhtSvM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_influxdata_influxdb1_client",
    importpath = "github.com/influxdata/influxdb1-client",
    sum = "h1:/WZQPMZNsjZ7IlCpsLGdQBINg5bxKQ1K1sh6awxLtkA=",
    version = "v0.0.0-20191209144304-8bf82d3c094d",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    importpath = "github.com/jmespath/go-jmespath",
    sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_jonboulle_clockwork",
    importpath = "github.com/jonboulle/clockwork",
    sum = "h1:UOGuzwb1PwsrDAObMuhUnj0p5ULPj8V/xJ7Kx9qUBdQ=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_jpillora_backoff",
    importpath = "github.com/jpillora/backoff",
    sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
    version = "v4.20.0+incompatible",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_knetic_govaluate",
    importpath = "github.com/Knetic/govaluate",
    sum = "h1:1G1pk05UrOh0NlF1oeaaix1x8XzrfjIDK47TY0Zehcw=",
    version = "v3.0.1-0.20171022003610-9aa49832a739+incompatible",
)

go_repository(
    name = "com_github_lightstep_lightstep_tracer_common_golang_gogo",
    importpath = "github.com/lightstep/lightstep-tracer-common/golang/gogo",
    sum = "h1:143Bb8f8DuGWck/xpNUOckBVYfFbBTnLevfRZ1aVVqo=",
    version = "v0.0.0-20190605223551-bc2310a04743",
)

go_repository(
    name = "com_github_lightstep_lightstep_tracer_go",
    importpath = "github.com/lightstep/lightstep-tracer-go",
    sum = "h1:vi1F1IQ8N7hNWytK9DpJsUfQhGuNSc19z330K6vl4zk=",
    version = "v0.18.1",
)

go_repository(
    name = "com_github_lyft_protoc_gen_validate",
    importpath = "github.com/lyft/protoc-gen-validate",
    sum = "h1:KNt/RhmQTOLr7Aj8PsJ7mTronaFyx80mRTT9qF261dA=",
    version = "v0.0.13",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    importpath = "github.com/mattn/go-colorable",
    sum = "h1:c1ghPdyEDarC70ftn0y+A/Ee++9zz8ljHG1b13eJ0s8=",
    version = "v0.1.8",
)

go_repository(
    name = "com_github_mitchellh_cli",
    importpath = "github.com/mitchellh/cli",
    sum = "h1:iGBIsUe3+HZ/AD/Vd7DErOt5sU9fa8Uj7A2s1aggv1Y=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    importpath = "github.com/mitchellh/go-homedir",
    sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_mitchellh_go_testing_interface",
    importpath = "github.com/mitchellh/go-testing-interface",
    sum = "h1:jrgshOhYAUVNMAJiKbEu7EqAwgJJ2JqpQmpLJOu07cU=",
    version = "v1.14.1",
)

go_repository(
    name = "com_github_mitchellh_gox",
    importpath = "github.com/mitchellh/gox",
    sum = "h1:lfGJxY7ToLJQjHHwi0EX6uYBdK78egf954SQl13PQJc=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_mitchellh_iochan",
    importpath = "github.com/mitchellh/iochan",
    sum = "h1:C+X3KsSTLFVBr/tK1eYN/vs4rJcvsiLU338UhYPJWeY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    sum = "h1:OVowDSCllw/YjdLkam3/sm7wEtOy59d8ndGgCcyj8cs=",
    version = "v1.4.3",
)

go_repository(
    name = "com_github_nats_io_jwt",
    importpath = "github.com/nats-io/jwt",
    sum = "h1:+RB5hMpXUUA2dfxuhBTEkMOrYmM+gKIZYS1KjSostMI=",
    version = "v0.3.2",
)

go_repository(
    name = "com_github_nats_io_nats_go",
    importpath = "github.com/nats-io/nats.go",
    sum = "h1:ik3HbLhZ0YABLto7iX80pZLPw/6dx3T+++MZJwLnMrQ=",
    version = "v1.9.1",
)

go_repository(
    name = "com_github_nats_io_nats_server_v2",
    importpath = "github.com/nats-io/nats-server/v2",
    sum = "h1:i2Ly0B+1+rzNZHHWtD4ZwKi+OU5l+uQo1iDHZ2PmiIc=",
    version = "v2.1.2",
)

go_repository(
    name = "com_github_nats_io_nkeys",
    importpath = "github.com/nats-io/nkeys",
    sum = "h1:6JrEfig+HzTH85yxzhSVbjHRJv9cn0p6n3IngIcM5/k=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_nats_io_nuid",
    importpath = "github.com/nats-io/nuid",
    sum = "h1:5iA8DT8V7q8WK2EScv2padNa/rTESc1KdnPw4TC2paw=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_oklog_oklog",
    importpath = "github.com/oklog/oklog",
    sum = "h1:wVfs8F+in6nTBMkA7CbRw+zZMIB7nNM825cM1wuzoTk=",
    version = "v0.3.2",
)

go_repository(
    name = "com_github_oklog_run",
    importpath = "github.com/oklog/run",
    sum = "h1:GEenZ1cK0+q0+wsJew9qUg/DyD8k3JzYsZAi5gYi2mA=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_olekukonko_tablewriter",
    importpath = "github.com/olekukonko/tablewriter",
    sum = "h1:P2Ga83D34wi1o9J6Wh1mRuqd4mF/x/lgBS7N7AbDhec=",
    version = "v0.0.5",
)

go_repository(
    name = "com_github_op_go_logging",
    importpath = "github.com/op/go-logging",
    sum = "h1:lDH9UUVJtmYCjyT0CI4q8xvlXPxeZ0gYCVvWbmPlp88=",
    version = "v0.0.0-20160315200505-970db520ece7",
)

go_repository(
    name = "com_github_opentracing_basictracer_go",
    importpath = "github.com/opentracing/basictracer-go",
    sum = "h1:YyUAhaEfjoWXclZVJ9sGoNct7j4TVk7lZWlQw5UXuoo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_opentracing_contrib_go_observer",
    importpath = "github.com/opentracing-contrib/go-observer",
    sum = "h1:lM6RxxfUMrYL/f8bWEUqdXrANWtrL7Nndbm9iFN0DlU=",
    version = "v0.0.0-20170622124052-a52f23424492",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    sum = "h1:uEJPy/1a5RIPAJ0Ov+OIO8OxWu77jEv+1B0VhjKrZUs=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_openzipkin_contrib_zipkin_go_opentracing",
    importpath = "github.com/openzipkin-contrib/zipkin-go-opentracing",
    sum = "h1:ZCnq+JUrvXcDVhX/xRolRBZifmabN1HcS1wrPSvxhrU=",
    version = "v0.4.5",
)

go_repository(
    name = "com_github_openzipkin_zipkin_go",
    importpath = "github.com/openzipkin/zipkin-go",
    sum = "h1:nY8Hti+WKaP0cRsSeQ026wU03QsM762XBeCXBb9NAWI=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_pact_foundation_pact_go",
    importpath = "github.com/pact-foundation/pact-go",
    sum = "h1:OYkFijGHoZAYbOIb1LWXrwKQbMMRUv1oQ89blD2Mh2Q=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_pascaldekloe_goe",
    importpath = "github.com/pascaldekloe/goe",
    sum = "h1:cBOtyMzM9HTpWjXfbbunk26uA6nG3a8n06Wieeh0MwY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_pborman_uuid",
    importpath = "github.com/pborman/uuid",
    sum = "h1:J7Q5mO4ysT1dv8hyrUGHb9+ooztCXu1D8MY8DZYsu3g=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_performancecopilot_speed",
    importpath = "github.com/performancecopilot/speed",
    sum = "h1:2WnRzIquHa5QxaJKShDkLM+sc0JPuwhXzK8OYOyt3Vg=",
    version = "v3.0.0+incompatible",
)

go_repository(
    name = "com_github_pierrec_lz4",
    importpath = "github.com/pierrec/lz4",
    sum = "h1:Ix9yFKn1nSPBLFl/yZknTp8TU5G4Ps0JDmguYK6iH1A=",
    version = "v2.6.0+incompatible",
)

go_repository(
    name = "com_github_pkg_profile",
    importpath = "github.com/pkg/profile",
    sum = "h1:F++O52m40owAmADcojzM+9gyjmMOY/T4oYJkgFDH8RE=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_posener_complete",
    importpath = "github.com/posener/complete",
    sum = "h1:GqpA1/5oN1NgsxoSA4RH0YWTaqvUlQNeOpHXD/JRbOQ=",
    version = "v1.2.2-0.20190308074557-af07aa5181b3",
)

go_repository(
    name = "com_github_rcrowley_go_metrics",
    importpath = "github.com/rcrowley/go-metrics",
    sum = "h1:N/ElC8H3+5XpJzTSTfLsJV/mx9Q9g7kxmchpfZyxgzM=",
    version = "v0.0.0-20201227073835-cf1acfcdf475",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    importpath = "github.com/rogpeppe/fastuuid",
    sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_russross_blackfriday_v2",
    importpath = "github.com/russross/blackfriday/v2",
    sum = "h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_ryanuber_columnize",
    importpath = "github.com/ryanuber/columnize",
    sum = "h1:j1Wcmh8OrK4Q7GXY+V7SVSY8nUWQxHW5TkBe7YUl+2s=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_samuel_go_zookeeper",
    importpath = "github.com/samuel/go-zookeeper",
    sum = "h1:p3Vo3i64TCLY7gIfzeQaUJ+kppEO5WQG3cL8iE8tGHU=",
    version = "v0.0.0-20190923202752-2cc03de413da",
)

go_repository(
    name = "com_github_sean_seed",
    importpath = "github.com/sean-/seed",
    sum = "h1:nn5Wsu0esKSJiIVhscUtVbo7ada43DJhG55ua/hjS5I=",
    version = "v0.0.0-20170313163322-e2103e2c3529",
)

go_repository(
    name = "com_github_shopify_sarama",
    importpath = "github.com/Shopify/sarama",
    sum = "h1:lOi3SfE6OcFlW9Trgtked2aHNZ2BIG/d6Do+PEUAqqM=",
    version = "v1.28.0",
)

go_repository(
    name = "com_github_shopify_toxiproxy",
    importpath = "github.com/Shopify/toxiproxy",
    sum = "h1:TKdv8HiTLgE5wdJuEML90aBgNWsokNbMijUGhmcoBJc=",
    version = "v2.1.4+incompatible",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    importpath = "github.com/smartystreets/assertions",
    sum = "h1:UVQPSSmc3qtTi+zPPkCXvZX9VvW/xT/NsRvKfwY81a8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_smartystreets_goconvey",
    importpath = "github.com/smartystreets/goconvey",
    sum = "h1:fv0U8FUIMPNf1L9lnHLvLhgicrIVChEkdzIKYqbNC9s=",
    version = "v1.6.4",
)

go_repository(
    name = "com_github_soheilhy_cmux",
    importpath = "github.com/soheilhy/cmux",
    sum = "h1:jjzc5WVemNEDTLwv9tlmemhC73tI08BNOIGwBOo10Js=",
    version = "v0.1.5",
)

go_repository(
    name = "com_github_sony_gobreaker",
    importpath = "github.com/sony/gobreaker",
    sum = "h1:oMnRNZXX5j85zso6xCPRNPtmAycat+WcoKbklScLDgQ=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    sum = "h1:+KmjbUw1hriSNMF55oPrkZcb27aECyrj8V2ytv7kWDw=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_streadway_amqp",
    importpath = "github.com/streadway/amqp",
    sum = "h1:WhxRHzgeVGETMlmVfqhRn8RIeeNoPr2Czh33I4Zdccw=",
    version = "v0.0.0-20190827072141-edfb9018d271",
)

go_repository(
    name = "com_github_streadway_handy",
    importpath = "github.com/streadway/handy",
    sum = "h1:AhmOdSHeswKHBjhsLs/7+1voOxT+LLrSk/Nxvk35fug=",
    version = "v0.0.0-20190108123426-d5acb3125c2a",
)

go_repository(
    name = "com_github_tmc_grpc_websocket_proxy",
    importpath = "github.com/tmc/grpc-websocket-proxy",
    sum = "h1:uruHq4dN7GR16kFc5fp3d1RIYzJW5onx8Ybykw2YQFA=",
    version = "v0.0.0-20201229170055-e5319fda7802",
)

go_repository(
    name = "com_github_urfave_cli",
    importpath = "github.com/urfave/cli",
    sum = "h1:9ic0a+f2TCJ5tSbVRX/FSSCIHJacFLYxcuNexNMJF8Q=",
    version = "v1.22.8",
)

go_repository(
    name = "com_github_xiang90_probing",
    importpath = "github.com/xiang90/probing",
    sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
    version = "v0.0.0-20190116061207-43a291ad63a2",
)

go_repository(
    name = "com_github_yuin_goldmark",
    importpath = "github.com/yuin/goldmark",
    sum = "h1:zHPiabbIRssZOI0MAzJDHsyvG4MXCGqVaMOwR+HeoQQ=",
    version = "v1.4.8",
)

go_repository(
    name = "com_sourcegraph_sourcegraph_appdash",
    importpath = "sourcegraph.com/sourcegraph/appdash",
    sum = "h1:ucqkfpjg9WzSUubAO62csmucvxl4/JeW3F4I4909XkM=",
    version = "v0.0.0-20190731080439-ebfcffb1b5c0",
)

go_repository(
    name = "in_gopkg_cheggaaa_pb_v1",
    importpath = "gopkg.in/cheggaaa/pb.v1",
    sum = "h1:n1tBJnnK2r7g9OW2btFH91V92STTUevLXYFb8gy9EMk=",
    version = "v1.0.28",
)

go_repository(
    name = "in_gopkg_gcfg_v1",
    importpath = "gopkg.in/gcfg.v1",
    sum = "h1:m8OOJ4ccYHnx2f4gQwpno8nAX5OGOh7RLaaz0pj3Ogs=",
    version = "v1.2.3",
)

go_repository(
    name = "in_gopkg_resty_v1",
    importpath = "gopkg.in/resty.v1",
    sum = "h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
    version = "v1.12.0",
)

go_repository(
    name = "in_gopkg_warnings_v0",
    importpath = "gopkg.in/warnings.v0",
    sum = "h1:wFXVbFY8DY5/xOe1ECiWdKCzZlxgshcYVNkBHstARME=",
    version = "v0.1.2",
)

go_repository(
    name = "io_etcd_go_bbolt",
    importpath = "go.etcd.io/bbolt",
    sum = "h1:/ecaJf0sk1l4l6V4awd65v2C3ILy7MSj+s/x1ADCIMU=",
    version = "v1.3.6",
)

go_repository(
    name = "io_etcd_go_etcd",
    importpath = "go.etcd.io/etcd",
    sum = "h1:5aomL5mqoKHxw6NG+oYgsowk8tU8aOalo2IdZxdWHkw=",
    version = "v3.3.18+incompatible",
)

go_repository(
    name = "io_k8s_sigs_yaml",
    importpath = "sigs.k8s.io/yaml",
    sum = "h1:kr/MCeFWJWTwyaHoR9c8EjH9OumOmoF9YGiZd7lFm/Q=",
    version = "v1.2.0",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    sum = "h1:gqCw0LfLxScz8irSi8exQc7fyQ0fKQU/qnC/X8+V/1M=",
    version = "v0.23.0",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    sum = "h1:67zQnAE0T2rB0A3CwLSas0K+SbVzSxP+zTLkQLexeiw=",
    version = "v0.70.0",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:w43yiav+6bVFTBQFZX0r7ipe9JQ1QsbMgHwbBziscLw=",
    version = "v1.28.0",
)

go_repository(
    name = "com_github_cncf_udpa_go",
    importpath = "github.com/cncf/udpa/go",
    sum = "h1:hzAQntlaYRkVSFEfj9OTWlVV1H155FMD8BTKktLv0QI=",
    version = "v0.0.0-20210930031921-04548b0d99d4",
)

go_repository(
    name = "com_github_gin_contrib_sse",
    importpath = "github.com/gin-contrib/sse",
    sum = "h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_gin_gonic_gin",
    importpath = "github.com/gin-gonic/gin",
    sum = "h1:ahKqKTFpO5KTPHxWZjEdPScmYaGtLo8Y4DMHoEsnp14=",
    version = "v1.6.3",
)

go_repository(
    name = "com_github_go_playground_assert_v2",
    importpath = "github.com/go-playground/assert/v2",
    sum = "h1:MsBgLAaY856+nPRTKrp3/OZK38U/wa0CcBYNjji3q3A=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_go_playground_locales",
    importpath = "github.com/go-playground/locales",
    sum = "h1:HyWk6mgj5qFqCT5fjGBuRArbVDfE4hi8+e8ceBS/t7Q=",
    version = "v0.13.0",
)

go_repository(
    name = "com_github_go_playground_universal_translator",
    importpath = "github.com/go-playground/universal-translator",
    sum = "h1:icxd5fm+REJzpZx7ZfpaD876Lmtgy7VtROAbHHXk8no=",
    version = "v0.17.0",
)

go_repository(
    name = "com_github_go_playground_validator_v10",
    importpath = "github.com/go-playground/validator/v10",
    sum = "h1:KgJ0snyC2R9VXYN2rneOtQcw5aHQB1Vv0sFl1UcHBOY=",
    version = "v10.2.0",
)

go_repository(
    name = "com_github_klauspost_compress",
    importpath = "github.com/klauspost/compress",
    sum = "h1:xqfchp4whNFxn5A4XFyyYtitiWI8Hy5EW59jEwcyL6U=",
    version = "v1.15.0",
)

go_repository(
    name = "com_github_leodido_go_urn",
    importpath = "github.com/leodido/go-urn",
    sum = "h1:hpXL4XnriNwQ/ABnpepYM/1vCLWNDfUNts8dX3xTG6Y=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_ugorji_go",
    importpath = "github.com/ugorji/go",
    sum = "h1:/68gy2h+1mWMrwZFeD1kQialdSzAb432dtpeJ42ovdo=",
    version = "v1.1.7",
)

go_repository(
    name = "com_github_ugorji_go_codec",
    importpath = "github.com/ugorji/go/codec",
    sum = "h1:2SvQaVZ1ouYrrKKwoSk2pzd4A9evlKJb9oTL+OaLUSs=",
    version = "v1.1.7",
)

go_repository(
    name = "com_github_agl_ed25519",
    importpath = "github.com/agl/ed25519",
    sum = "h1:w1UutsfOrms1J05zt7ISrnJIXKzwaspym5BTKGx93EI=",
    version = "v0.0.0-20170116200512-5312a6153412",
)

go_repository(
    name = "com_github_asaskevich_govalidator",
    importpath = "github.com/asaskevich/govalidator",
    sum = "h1:Byv0BzEl3/e6D5CLfI0j/7hiIEtvGVFPCZ7Ei2oq8iQ=",
    version = "v0.0.0-20210307081110-f21760c49a8d",
)

go_repository(
    name = "com_github_bifurcation_mint",
    importpath = "github.com/bifurcation/mint",
    sum = "h1:fUjoj2bT6dG8LoEe+uNsKk8J+sLkDbQkJnB6Z1F02Bc=",
    version = "v0.0.0-20180715133206-93c51c6ce115",
)

go_repository(
    name = "com_github_cheekybits_genny",
    importpath = "github.com/cheekybits/genny",
    sum = "h1:uGGa4nei+j20rOSeDeP5Of12XVm7TGUd4dJA9RDitfE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_coreos_go_iptables",
    importpath = "github.com/coreos/go-iptables",
    sum = "h1:is9qnZMPYjLd8LYqmm/qlE+wwEgJIkTYdhV3rfZo4jk=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_dchest_siphash",
    importpath = "github.com/dchest/siphash",
    sum = "h1:9DFz8tQwl9pTVt5iok/9zKyzA1Q6bRGiF3HPiEEVr9I=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_docker_libcontainer",
    importpath = "github.com/docker/libcontainer",
    sum = "h1:++SbbkCw+X8vAd4j2gOCzZ2Nn7s2xFALTf7LZKmM1/0=",
    version = "v2.2.1+incompatible",
)

go_repository(
    name = "com_github_ginuerzh_gosocks4",
    importpath = "github.com/ginuerzh/gosocks4",
    sum = "h1:ojDKUyz+uaEeRm2usY1cyQiXTqJqrKxfeE6SVBXq4m0=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_ginuerzh_gosocks5",
    importpath = "github.com/ginuerzh/gosocks5",
    sum = "h1:K0Ua23U9LU3BZrf3XpGDcs0mP8DiEpa6PJE4TA/MU3s=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_ginuerzh_gost",
    importpath = "github.com/ginuerzh/gost",
    replace = "github.com/btwiuse/gost",
    sum = "h1:z71l9MRGQmnoFQ1qUTv2lfUlVlgdbTdDbSEKLEN/+OQ=",
    version = "v0.0.0-20220430110229-7bf437d68fdb",
)

go_repository(
    name = "com_github_ginuerzh_tls_dissector",
    importpath = "github.com/ginuerzh/tls-dissector",
    sum = "h1:VPXbYRvZUzTemsI7u0FzOnEuHeHwQuMTPXApAu8aeX4=",
    version = "v0.0.2-0.20200224064855-24ab2b3a3796",
)

go_repository(
    name = "com_github_go_gost_relay",
    importpath = "github.com/go-gost/relay",
    sum = "h1:itaaJhQJ19kUXEB4Igb0EbY8m+1Py2AaNNSBds/9gk4=",
    version = "v0.1.1-0.20211123134818-8ef7fd81ffd7",
)

go_repository(
    name = "com_github_go_log_log",
    importpath = "github.com/go-log/log",
    sum = "h1:z8i91GBudxD5L3RmF0KVpetCbcGWAV7q1Tw1eRwQM9Q=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_gobwas_glob",
    importpath = "github.com/gobwas/glob",
    sum = "h1:A4xDbljILXROh+kObIiy5kIaPYD8e96x1tgBhUI5J+Y=",
    version = "v0.2.3",
)

go_repository(
    name = "com_github_google_gopacket",
    importpath = "github.com/google/gopacket",
    sum = "h1:ves8RnFZPGiFnTS0uPQStjwru6uO6h+nlr9j6fL7kF8=",
    version = "v1.1.19",
)

go_repository(
    name = "com_github_klauspost_cpuid",
    importpath = "github.com/klauspost/cpuid",
    sum = "h1:5JNjFYYQrZeKRJ0734q51WCEEn2huer72Dc7K+R/b6s=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_klauspost_reedsolomon",
    importpath = "github.com/klauspost/reedsolomon",
    sum = "h1:g2erWKD2M6rgnPf89fCji6jNlhMKMdXcuNHMW1SYCIo=",
    version = "v1.9.15",
)

go_repository(
    name = "com_github_liamhaworth_go_tproxy",
    importpath = "github.com/LiamHaworth/go-tproxy",
    sum = "h1:eqa6queieK8SvoszxCu0WwH7lSVeL4/N/f1JwOMw1G4=",
    version = "v0.0.0-20190726054950-ef7efd7f24ed",
)

go_repository(
    name = "com_github_lucas_clemente_aes12",
    importpath = "github.com/lucas-clemente/aes12",
    sum = "h1:sSeNEkJrs+0F9TUau0CgWTTNEwF23HST3Eq0A+QIx+A=",
    version = "v0.0.0-20171027163421-cd47fb39b79f",
)

go_repository(
    name = "com_github_lucas_clemente_quic_go",
    importpath = "github.com/lucas-clemente/quic-go",
    sum = "h1:v6WY87q9zD4dKASbG8hy/LpzAVNzEQzw8sEIeloJsc4=",
    version = "v0.27.0",
)

go_repository(
    name = "com_github_lucas_clemente_quic_go_certificates",
    importpath = "github.com/lucas-clemente/quic-go-certificates",
    sum = "h1:zqEC1GJZFbGZA0tRyNZqRjep92K5fujFtFsu5ZW7Aug=",
    version = "v0.0.0-20160823095156-d2f86524cced",
)

go_repository(
    name = "com_github_milosgajdos83_tenus",
    importpath = "github.com/milosgajdos83/tenus",
    replace = "github.com/milosgajdos/tenus",
    sum = "h1:jmaJzwaY1DUyYVD0lM4U+uvP2kkEg1VahDqRFxIkVBE=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_mmcloughlin_avo",
    importpath = "github.com/mmcloughlin/avo",
    sum = "h1:ExoghBBFY7A3RzgkAOq0XbHs9zaT/bHq7xysgyp3z3Q=",
    version = "v0.0.0-20210104032911-599bdd1269f4",
)

go_repository(
    name = "com_github_riobard_go_bloom",
    importpath = "github.com/riobard/go-bloom",
    sum = "h1:f/FNXud6gA3MNr8meMVVGxhp+QBTqY91tM8HjEuMjGg=",
    version = "v0.0.0-20200614022211-cdc8013cb5b3",
)

go_repository(
    name = "com_github_ryanuber_go_glob",
    importpath = "github.com/ryanuber/go-glob",
    sum = "h1:iQh3xXAumdQ+4Ufa5b25cRpC5TYKlno6hsv6Cb3pkBk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_shadowsocks_go_shadowsocks2",
    importpath = "github.com/shadowsocks/go-shadowsocks2",
    sum = "h1:PDSQv9y2S85Fl7VBeOMF9StzeXZyK1HakRm86CUbr28=",
    version = "v0.1.5",
)

go_repository(
    name = "com_github_shadowsocks_shadowsocks_go",
    importpath = "github.com/shadowsocks/shadowsocks-go",
    sum = "h1:XU9hik0exChEmY92ALW4l9WnDodxLVS9yOSNh2SizaQ=",
    version = "v0.0.0-20200409064450-3e585ff90601",
)

go_repository(
    name = "com_github_songgao_water",
    importpath = "github.com/songgao/water",
    sum = "h1:TG/diQgUe0pntT/2D9tmUCz4VNwm9MfrtPr0SU2qSX8=",
    version = "v0.0.0-20200317203138-2b4b6d7c09d8",
)

go_repository(
    name = "com_github_templexxx_cpufeat",
    importpath = "github.com/templexxx/cpufeat",
    sum = "h1:89CEmDvlq/F7SJEOqkIdNDGJXrQIhuIx9D2DBXjavSU=",
    version = "v0.0.0-20180724012125-cef66df7f161",
)

go_repository(
    name = "com_github_templexxx_xor",
    importpath = "github.com/templexxx/xor",
    sum = "h1:fj5tQ8acgNUr6O8LEplsxDhUIe2573iLkJc+PqnzZTI=",
    version = "v0.0.0-20191217153810-f85b25db303b",
)

go_repository(
    name = "com_github_tjfoc_gmsm",
    importpath = "github.com/tjfoc/gmsm",
    sum = "h1:aMe1GlZb+0bLjn+cKTPEvvn9oUEBlJitaZiiBwsbgho=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_xtaci_tcpraw",
    importpath = "github.com/xtaci/tcpraw",
    sum = "h1:VDlqo0op17JeXBM6e2G9ocCNLOJcw9mZbobMbJjo0vk=",
    version = "v1.2.25",
)

go_repository(
    name = "com_github_yawning_chacha20",
    importpath = "github.com/Yawning/chacha20",
    sum = "h1:I6/SJSN9wJMJ+ZyQaCHUlzoTA4ypU5Bb44YWR1wTY/0=",
    version = "v0.0.0-20170904085104-e3b1f968fc63",
)

go_repository(
    name = "in_gopkg_gorilla_websocket_v1",
    importpath = "gopkg.in/gorilla/websocket.v1",
    sum = "h1:lREme3ezAGPCpxSHwjGkHhAJX+ed2B6vzAJ+kaqBEIM=",
    version = "v1.4.0",
)

go_repository(
    name = "in_gopkg_xtaci_kcp_go_v4",
    importpath = "gopkg.in/xtaci/kcp-go.v4",
    sum = "h1:S9IF+L55Ugzl/hVA6wvuL3SuAtTUzH2cBBC88MXQxnE=",
    version = "v4.3.2",
)

go_repository(
    name = "io_rsc_pdf",
    importpath = "rsc.io/pdf",
    sum = "h1:k1MczvYDUvJBe93bYd7wrZLLUEcLZAuF824/I4e5Xr4=",
    version = "v0.1.1",
)

go_repository(
    name = "org_golang_x_arch",
    importpath = "golang.org/x/arch",
    sum = "h1:YcaRWbSY2VfP0/k25uHKKrk3Vs3C7mo03vq103Ire8E=",
    version = "v0.0.0-20190909030613-46d78d1859ac",
)

go_repository(
    name = "org_torproject_git_pluggable_transports_goptlib_git",
    importpath = "git.torproject.org/pluggable-transports/goptlib.git",
    sum = "h1:0qRF7Dw5qXd0FtZkjWUiAh5GTutRtDGL4GXUDJ4qMHs=",
    version = "v1.2.0",
)

go_repository(
    name = "org_torproject_git_pluggable_transports_obfs4_git",
    importpath = "git.torproject.org/pluggable-transports/obfs4.git",
    sum = "h1:c8h60PKrRxEB5debIHBmP7T+s/EUNXTklXqlmJfYiJQ=",
    version = "v0.0.0-20181103133120-08f4d470188e",
)

go_repository(
    name = "com_github_abiosoft_ishell",
    importpath = "github.com/abiosoft/ishell",
    sum = "h1:zpwIuEHc37EzrsIYah3cpevrIc8Oma7oZPxr03tlmmw=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_abiosoft_readline",
    importpath = "github.com/abiosoft/readline",
    sum = "h1:CjPUSXOiYptLbTdr1RceuZgSFDQ7U15ITERUGrUORx8=",
    version = "v0.0.0-20180607040430-155bce2042db",
)

go_repository(
    name = "com_github_chzyer_logex",
    importpath = "github.com/chzyer/logex",
    sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_chzyer_test",
    importpath = "github.com/chzyer/test",
    sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
    version = "v0.0.0-20180213035817-a1ea475d72b1",
)

go_repository(
    name = "com_github_flynn_archive_go_shlex",
    importpath = "github.com/flynn-archive/go-shlex",
    sum = "h1:BMXYYRWTLOJKlh+lOBt6nUQgXAfB7oVIQt5cNreqSLI=",
    version = "v0.0.0-20150515145356-3f9db97f8568",
)

local_repository(
    name = "starlark",
    path = "starlark",
)

load("@starlark//:defs.bzl", "print_seq")

print_seq()

# https://github.com/google/bazel_rules_install
git_repository(
    name = "com_github_google_rules_install",
    branch = "main",
    # commit = "e93a17ed42a8a622a78fbf4737309e583f4b3cb4",
    remote = "https://github.com/google/bazel_rules_install.git",
)

load("@com_github_google_rules_install//:deps.bzl", "install_rules_dependencies")

install_rules_dependencies()

load("@com_github_google_rules_install//:setup.bzl", "install_rules_setup")

install_rules_setup()

# install bbox
git_repository(
    name = "bbox",
    # commit = "16d0642dda469579fecf2d2e1efff544e30a60c1",
    branch = "master",
    remote = "https://github.com/btwiuse/bbox.git",
)

load("@bbox//:package.bzl", "register_repositories")

register_repositories()
