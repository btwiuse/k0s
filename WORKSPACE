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
    version = "1.16.4",
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
    branch = "master",
    remote = "https://github.com/protocolbuffers/protobuf",
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

load("@rules_rust//rust:repositories.bzl", "rust_repositories")

# https://bazelbuild.github.io/rules_rust/
# rust_repositories(
#     edition = "2018",
#     rustfmt_version = "1.50.0",
#     version = "1.50.0",
# )
rust_repositories(
    edition = "2018",
    iso_date = "2021-05-23",
    version = "nightly",
)

# https://docs.rs/crate/cargo-raze/0.0.19
load("//cargo:crates.bzl", "raze_fetch_remote_crates")

raze_fetch_remote_crates()

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
    sum = "h1:6hMiXswcleXj5oNfcJc+DXS8Vj36XX2LaX98udog6Kc=",
    version = "v1.5.15",
)

go_repository(
    name = "com_github_felixge_httpsnoop",
    importpath = "github.com/felixge/httpsnoop",
    sum = "h1:lvB5Jl89CsZtGIWuTcDM1E/vkVs49/Ml7JJe07l8SPQ=",
    version = "v1.0.1",
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
    sum = "h1:yUdfgN0XgIJw7foRItutHYUIhlcKzcSf5vDpdhQAKTc=",
    version = "v0.0.0-20180725130230-947c36da3153",
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
    sum = "h1:kLcOMZeuLAJvL2BPWLMIj5oaZQobrkAqrL+WFZwQses=",
    version = "v4.9.0+incompatible",
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
    sum = "h1:K7/B1jt6fIBQVd4Owv2MqGQClcgf0R266+7C/QjRcLc=",
    version = "v0.4.0",
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
    sum = "h1:41CTEDOoUXp+FxbPYuEhth5dE/s+NT1cRuhSoqhBQ1E=",
    version = "v0.0.0-20210122040257-d980be63207e",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    importpath = "github.com/googleapis/gax-go/v2",
    sum = "h1:sjZBwGj9Jlw33ImPtvFviGYvseOtDM7hkSKB7+Tv3SM=",
    version = "v2.0.5",
)

go_repository(
    name = "com_github_googleapis_gnostic",
    importpath = "github.com/googleapis/gnostic",
    sum = "h1:DLJCy1n/vrD4HPjOvYcT8aYQXpPIzoRZONaYwyycI+I=",
    version = "v0.4.1",
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
    sum = "h1:UauaLniWCFHWd+Jp9oCEkTBj8VO/9DKg3PV3VCNMDIg=",
    version = "v0.3.9",
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
    sum = "h1:kQ6Cb7aHOHTSzNVNEhmp8EcWKLb4CbiMW9h9VyIhO4E=",
    version = "v3.0.0",
)

go_repository(
    name = "com_github_sergi_go_diff",
    importpath = "github.com/sergi/go-diff",
    sum = "h1:we8PVUC3FE2uYfodKH/nBHMSetSfHDR6scGdBi+erh0=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    sum = "h1:5jhuqJyZCZf2JRofRvN/nIFgIWNzPa3/Vz8mYylgbWc=",
    version = "v1.2.2",
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
    sum = "h1:fMFkbjo6ElaL3POE3ctxLuX3aR+bEzpo2q1R5Aq2JhE=",
    version = "v0.9.2",
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
    sum = "h1:JApXBKYyB7l9xx+DK7/+mFjC7A9Bt5A93FPvFD0HIFE=",
    version = "v0.0.0-20201113003025-83324d819ded",
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
    sum = "h1:qqqz049vc4dT+0UBxt4BGzvcGr7iVbNiGRYo5NMz1Yc=",
    version = "v2.7.0",
)

go_repository(
    name = "io_k8s_kube_openapi",
    importpath = "k8s.io/kube-openapi",
    sum = "h1:sOHNzJIkytDF6qadMNKhhDRpc6ODik8lVC6nOur7B2c=",
    version = "v0.0.0-20201113171705-d219536bb9fd",
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
    sum = "h1:YHQV7Dajm86OuqnIR6zAelnDWBRjo+YhYV9PmGrh1s8=",
    version = "v4.0.2",
)

go_repository(
    name = "io_k8s_utils",
    importpath = "k8s.io/utils",
    sum = "h1:u5rPykqiCpL+LBfjRkXvnK71gOgIdmq3eHUEkPrbeTI=",
    version = "v0.0.0-20210305010621-2afb4311ab10",
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
    sum = "h1:Xkwi/a1rcvNg1PPYe5vI8GbeBY/jrVuDX5ASuANWTrk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_google_martian_v3",
    importpath = "github.com/google/martian/v3",
    sum = "h1:wCKgOCHuUEVfsaQLpPSJb7VdYCdTVZQAuOdYm1yc/60=",
    version = "v3.1.0",
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
    sum = "h1:SZxvLBoTP5yHO3Frd4z4vrF+DBX9vMVanchswa69toE=",
    version = "v0.0.0-20210220032956-6a3ed077a48d",
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
    sum = "h1:x3zkuE2lUk/RIekyAJ3XRqSCP4zwWDfcw/YJCuCAACg=",
    version = "v0.8.2",
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
    sum = "h1:HD8gA2tkByhMAwYaFAX9w2l7vxvBQ5NMoxDrkhqhtn4=",
    version = "v0.0.0-20190306092124-e2d15f34fcf9",
)

go_repository(
    name = "com_github_anmitsu_go_shlex",
    importpath = "github.com/anmitsu/go-shlex",
    sum = "h1:kFOfPq6dUM1hTo4JG6LR5AXSUEsOjtdm0kw0FtQtMJA=",
    version = "v0.0.0-20161002113705-648efa622239",
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
    sum = "h1:NpkT8MjJJMcgPJ5Q9E66QUgY9QRyxqM8MFx2P29uQZ4=",
    version = "v52.4.0+incompatible",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_azure_auth",
    importpath = "github.com/Azure/go-autorest/autorest/azure/auth",
    sum = "h1:8DQB8yl7aLQuP+nuR5e2RO6454OvFlSTXXaNHshc16s=",
    version = "v0.5.7",
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
    sum = "h1:15vMO4y76dehZSq7pAaOLQxC6dZYsSrj2GQpflyM/L4=",
    version = "v0.2.0",
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
    sum = "h1:s1TvRnXwL2xJRaccrdcBQMZxq6X7DvsMogtmJeHDdrc=",
    version = "v1.0.0",
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
    name = "com_github_caddyserver_caddy_v2",
    importpath = "github.com/caddyserver/caddy/v2",
    sum = "h1:fnrqJLa3G5vfxcxmOH/+kJOcunPLhSBnjgIvjXV/QTA=",
    version = "v2.3.0",
)

go_repository(
    name = "com_github_caddyserver_certmagic",
    importpath = "github.com/caddyserver/certmagic",
    sum = "h1:gpjCX6/8hHRgVXxy1v2AQdoAX6XRXIA8fBUZtEpnVg0=",
    version = "v0.12.1-0.20201215190346-201f83a06067",
)

go_repository(
    name = "com_github_caddyserver_forwardproxy",
    importpath = "github.com/caddyserver/forwardproxy",
    replace = "github.com/klzgrad/forwardproxy",
    sum = "h1:HGs00fbTpkCmIaH20H7BXDRvIJL/3pdawnOJHLWAVb4=",
    version = "v0.0.0-20210120121422-9b4a5a242bd6",
)

go_repository(
    name = "com_github_caddyserver_nginx_adapter",
    importpath = "github.com/caddyserver/nginx-adapter",
    sum = "h1:n3jhK8Tp51oN8V0lrcYhOPTOWOcgJsRRSGnQwP27NXU=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_caddyserver_ntlm_transport",
    importpath = "github.com/caddyserver/ntlm-transport",
    sum = "h1:E7CFpPD7vDOYfaX6qUrmP9r2VOzhCVExOSPM1J4fOyY=",
    version = "v0.1.1-0.20200409193839-5d99ab17e974",
)

go_repository(
    name = "com_github_cenkalti_backoff_v4",
    importpath = "github.com/cenkalti/backoff/v4",
    sum = "h1:JIufpQLbh4DkbQoii76ItQIUFzevQSqOLZca4eamEDs=",
    version = "v4.0.2",
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
    sum = "h1:3oxKN3wbHibqx897utPC2LTQU4J+IHWWJO+glkAkpFM=",
    version = "v1.4.1",
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
    sum = "h1:5oWIuRvwn93cie+OSt1zSnkaIQ1JFQM8bGlIv6O6Sts=",
    version = "v1.5.3",
)

go_repository(
    name = "com_github_dgraph_io_badger_v2",
    importpath = "github.com/dgraph-io/badger/v2",
    sum = "h1:n8KbImHW5qZCXv1y3tHjz5yz418/eTxeRJZ2ZuDm1ZU=",
    version = "v2.0.1-rc1.0.20200413122845-09dd2e1a4195",
)

go_repository(
    name = "com_github_dgraph_io_ristretto",
    importpath = "github.com/dgraph-io/ristretto",
    sum = "h1:MQLRM35Pp0yAyBYksjbj1nZI/w6eyRY/mWoM1sFf4kU=",
    version = "v0.0.2-0.20200115201040-8f368f2f2ab3",
)

go_repository(
    name = "com_github_dgryski_go_farm",
    importpath = "github.com/dgryski/go-farm",
    sum = "h1:tdlZCpZ/P9DhczCTSixgIKmwPv6+wP5DGjqLYw5SUiA=",
    version = "v0.0.0-20190423205320-6a90982ecee2",
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
    sum = "h1:8sAhBGEM0dRWogWqWyQeIJnxjWO6oIjl8FKqREDsGfk=",
    version = "v1.2.0",
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
    name = "com_github_freman_caddy2_reauth",
    importpath = "github.com/freman/caddy2-reauth",
    sum = "h1:JBn9TEuc5i0FPckHVOL8NI8Q+R0kNBOBLvBLadgkZ3o=",
    version = "v0.0.0-20200518130136-6064aa96b1a8",
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
    sum = "h1:j3L6gSLQalDETeEg/Jg0mGY0/y/N6zI2xX1978P0Uqw=",
    version = "v0.1.1",
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
    sum = "h1:gvPdv/Hr++TRFCl0UbPFHC54P9N9jgsRPnmnr419Uck=",
    version = "v1.3.1",
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
    sum = "h1:LUHzmkK3GUKUrL/1gfBUxAHzcev3apQlezX/+O7ma6w=",
    version = "v1.0.1",
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
    sum = "h1:7WsKqasmPThNvdl0Q5GPpbTDD/ZD98CfuawrMIuh7qQ=",
    version = "v3.1.10",
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
    sum = "h1:nNBDSCOigTSiarFpYE9J/KtEA1IOW4CNeqT9TQDqCxI=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_go_piv_piv_go",
    importpath = "github.com/go-piv/piv-go",
    sum = "h1:UtHPfrJsZKY+Z3UIjmJLh6DY+KtmNOl/9b/zt4N81pM=",
    version = "v1.5.0",
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
    sum = "h1:y12jRkkFxsd7GpqdSZ+/KCs/fJbqpEXSGd4+jfEaewE=",
    version = "v3.2.0+incompatible",
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
    sum = "h1:Li+angxmgvzlwDsPuFc1/nbqnq3gc4K/X7NrWjOADFI=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_google_cel_spec",
    importpath = "github.com/google/cel-spec",
    sum = "h1:HktvAjyBrKbDEZzD3oJQJ2khwAL1CEE1P7a5BNdVOMU=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_google_certificate_transparency_go",
    importpath = "github.com/google/certificate-transparency-go",
    sum = "h1:10MlrYzh5wfkToxWI4yJzffsxLfxcEDlOATMx/V9Kzw=",
    version = "v1.1.0",
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
    sum = "h1:Bn71r2jt5ObayLNUtMlCzNlKiw7o59esC9sz9ENjSe0=",
    version = "v1.2.2-0.20190612132142-05461f4df60a",
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
    sum = "h1:Ab7+HwSRL0dmsCYPvfSdiJ/iSC8y7Wf6wfHocry541I=",
    version = "v1.1.5",
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
    sum = "h1:4jgBlKK6tLKFvO8u5pmYjG91cqytmDCDvGh7ECVFfFs=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_iamd3vil_caddy_yaml_adapter",
    importpath = "github.com/iamd3vil/caddy_yaml_adapter",
    sum = "h1:5eTxtJy0pyxzY5a1N3bOap7JonTWkuRjrIEs9sK7ciE=",
    version = "v0.0.0-20200503183711-d479c29b475a",
)

go_repository(
    name = "com_github_iancoleman_strcase",
    importpath = "github.com/iancoleman/strcase",
    sum = "h1:VHgatEHNcBFEB7inlalqfNqw65aNkM1lGX2yt3NmbS8=",
    version = "v0.0.0-20191112232945-16388991a334",
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
    sum = "h1:g0I61F2K2DjRHz1cnxlkNSBIaePVoJIjjnHui8QHbiw=",
    version = "v2.0.4",
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
    sum = "h1:LXpIM/LZ5xGFhOpXAQUIMM1HdyqzVYM13zNdjCEEcA0=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_libdns_libdns",
    importpath = "github.com/libdns/libdns",
    sum = "h1:ewg3ByWrdUrxrje8ChPVMBNcotg7H9LQYg+u5De2RzI=",
    version = "v0.2.0",
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
    sum = "h1:ZC2Vc7/ZFkGmsVC9KvOjumD+G5lXy2RtTKyzRKO2BQ4=",
    version = "v1.8.1",
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
    sum = "h1:3l11YT8tm9MnwGFQ4kETwkzpAwY2Jt9lCrumCUW4+z4=",
    version = "v0.7.0",
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
    sum = "h1:LIH6K34bPVttyXnUWixk0bzH6/N07VxbSabxn5A5gZQ=",
    version = "v0.1.1",
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
    sum = "h1:zukEsf/1JZwCMgHiK3GZftabmxiCw4apj3a28RPBiVg=",
    version = "v1.1.0",
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
    sum = "h1:Y2lUDsFKVRSYGojLJ1yLxSXdMmMYTYls0rCvoqmMUQk=",
    version = "v3.1.0",
)

go_repository(
    name = "com_github_masterminds_sprig_v3",
    importpath = "github.com/Masterminds/sprig/v3",
    sum = "h1:j7GpgZ7PdFqNsmncycTHsLmVPf5/3wJtlgW9TNDYD9Y=",
    version = "v3.1.0",
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
    sum = "h1:KQODCqk+hBn3O7qfCRPj6L96uG65T5BSS95FKNEqtdA=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_mholt_caddy_webdav",
    importpath = "github.com/mholt/caddy-webdav",
    sum = "h1:8ATFx5TsqSf+PzTmQLAAAld7LRLwjKGjYCZ7a+hZvOc=",
    version = "v0.0.0-20200916200058-c949b3226234",
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
    sum = "h1:qkLXKzb1QoVatRyd/YlXZ/Kg0m5K3SPuoD82jjSOaBc=",
    version = "v0.4.15",
)

go_repository(
    name = "com_github_miekg_pkcs11",
    importpath = "github.com/miekg/pkcs11",
    sum = "h1:CIBkOawOtzJNE0B+EpRiUBzuVW7JEQAwdwhSS6YhIeg=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_mitchellh_copystructure",
    importpath = "github.com/mitchellh/copystructure",
    sum = "h1:Laisrj+bAB6b/yJwB5Bt3ITZhGJdqmxquMKeZ+mmkFQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_go_ps",
    importpath = "github.com/mitchellh/go-ps",
    sum = "h1:9+ke9YJ9KGWw5ANXK6ozjoK47uI3uNbXv4YVINBnGm8=",
    version = "v0.0.0-20190716172923-621e5597135b",
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
    sum = "h1:9D+8oIskB4VJBN5SFlmc27fSlIBZaov1Wpk/IfikLNY=",
    version = "v1.0.0",
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
    sum = "h1:eFXv9Nu1lGbrNbj619aWwZfVF5HBrm9Plte8aNptuTI=",
    version = "v0.0.0-20151028013722-8c68805598ab",
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
    sum = "h1:DQuhQpB1tVlglWS2hLQ5OV6B5r8aGxSrPc5Qo6uTN78=",
    version = "v1.4.4",
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
    sum = "h1:Keo9qb7iRJs2voHvunFtuuYFsbWeOBh8/P9v/kVMFtw=",
    version = "v1.8.0",
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
    sum = "h1:MZM7FHLqUHYI0Y/mQAt3d2aYa0SiNms/hFqC9qJYolM=",
    version = "v0.0.0-20180423040247-9e1955d9fb6e",
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
    sum = "h1:SWV2fHctRpRrp49VXJ6UZja7gU9QLHwRpIPBN89SKEo=",
    version = "v0.0.0-20171119174359-809beceb2371",
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
    sum = "h1:bBGb2GqrQ8wKHVOhcUfgRlaTHsO1S5rqsq/z93/mRSc=",
    version = "v0.15.4",
)

go_repository(
    name = "com_github_smallstep_certinfo",
    importpath = "github.com/smallstep/certinfo",
    sum = "h1:DAW0OxOdVIdAgsF/JN5Ro4ZApnTVfYqfaNYK7u5/obg=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_smallstep_cli",
    importpath = "github.com/smallstep/cli",
    sum = "h1:bOrYD1w0Vu82XN3r7mHuXoEI9RyczHyzKjzDHQ7V7EE=",
    version = "v0.15.2",
)

go_repository(
    name = "com_github_smallstep_nosql",
    importpath = "github.com/smallstep/nosql",
    sum = "h1:V1X5vfDsDt89499h3jZFUlR4VnnsYYs5tXaQZ0w8z5U=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_smallstep_truststore",
    importpath = "github.com/smallstep/truststore",
    sum = "h1:vNzEJmaJL0XOZD8uouXLmYu4/aP1UQ/wHUopH3qKeYA=",
    version = "v0.9.6",
)

go_repository(
    name = "com_github_smallstep_zcrypto",
    importpath = "github.com/smallstep/zcrypto",
    sum = "h1:cZMeOsWtXts169P82wfUkTCjReFOUTHVWKu/o5rv3YQ=",
    version = "v0.0.0-20200203191936-fbc32cf76bce",
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
    sum = "h1:nFm6S0SMdyzrzcmThSipiEubIDy8WEXKNZ0UOgiRpng=",
    version = "v1.3.1",
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
    sum = "h1:xVKxvI7ouOI5I+U9s2eeiUfMaWBVoXA3AWskkrqK0VM=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_stackexchange_wmi",
    importpath = "github.com/StackExchange/wmi",
    sum = "h1:G0m3OIz70MZUWq3EgK3CesDbo8upS2Vm9/P3FtgI+Jk=",
    version = "v0.0.0-20190523213315-cbe66965904d",
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
    sum = "h1:VWSxtAiQNh3zgHJpdpkpVYjTPqRE3P6UZCOPa1nRDio=",
    version = "v0.0.0-20200307114337-60d527fdb691",
)

go_repository(
    name = "com_github_zenazn_goji",
    importpath = "github.com/zenazn/goji",
    sum = "h1:mXV20Aj/BdWrlVzIn1kXFa+Tq62INlUi0cFFlztTaK0=",
    version = "v0.9.1-0.20160507202103-64eb34159fe5",
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
    sum = "h1:DPMeDvGTM54DXbPkVIZsp19fp/I2K7zwA/itHYHKo8Y=",
    version = "v1.56.0",
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
    sum = "h1:7odma5RETjNHWJnR32wx8t+Io4djHE1PqxCFx3iiZ2w=",
    version = "v2.5.1",
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
    sum = "h1:Dll2uFfOVI3fa8UzsHyP6z0M6fEc9ZTAMo+Y3z282Xg=",
    version = "v0.12.1",
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
    sum = "h1:AQkaJpH+/FmqRjmXZPELom5zIERYZfwTjnHpfoVMQEc=",
    version = "v0.0.0-20200419221736-3b63eb3a43b5",
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
    sum = "h1:xYq6+9AtI+xP3M4r0N1hCkHrInHDBohhquRgx9Kk6gI=",
    version = "v0.0.0-20180704124530-6e6d33e29852",
)

go_repository(
    name = "sm_step_go_crypto",
    importpath = "go.step.sm/crypto",
    sum = "h1:fbGUG5VJmDetC+RQ/T0tb6Sx0wCOgqKZcZYzTpUa7eo=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_caddyserver_caddy",
    importpath = "github.com/caddyserver/caddy",
    sum = "h1:i9gRhBgvc5ifchwWtSe7pDpsdS9+Q0Rw9oYQmYUTw1w=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_casbin_caddy_authz",
    importpath = "github.com/casbin/caddy-authz",
    sum = "h1:8gr0V4XIKWAkrGz/QEuDGD0uGSVh2GeV+6s2AWhSS28=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_casbin_casbin",
    importpath = "github.com/casbin/casbin",
    sum = "h1:ucjbS5zTrmSLtH4XogqOG920Poe6QatdXtz1FEbApeM=",
    version = "v1.9.1",
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
    name = "com_github_mholt_certmagic",
    importpath = "github.com/mholt/certmagic",
    sum = "h1:xKE9kZ5C8gelJC3+BNM6LJs1x21rivK7yxfTZMAuY2s=",
    version = "v0.6.2-0.20190624175158-6a42ef9fe8c2",
)

go_repository(
    name = "in_gopkg_mcuadros_go_syslog_v2",
    importpath = "gopkg.in/mcuadros/go-syslog.v2",
    sum = "h1:60g8zx1BijSVSgLTzLCW9UC4/+i1Ih9jJ1DR5Tgp9vE=",
    version = "v2.2.1",
)

go_repository(
    name = "com_github_hairyhenderson_caddy_teapot_module",
    importpath = "github.com/hairyhenderson/caddy-teapot-module",
    sum = "h1:POfr7XzYFIUEzAR5/f1+LZ9sN33GqG7XyscWljyFu5Y=",
    version = "v0.0.2",
)

go_repository(
    name = "net_starlark_go",
    importpath = "go.starlark.net",
    sum = "h1:4cM8hRkSWxRrHV7AQowcBrM91A8KJcUp2z0sfQOGMY0=",
    version = "v0.0.0-20210305151048-6a590ae7f4eb",
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
    sum = "h1:stqxghJ8lnlVknA37nlryttw+uizKuVJpZdp0hljNaw=",
    version = "v0.14.1-0.20210218105754-53593c3ab79f",
)

go_repository(
    name = "io_robpike_ivy",
    importpath = "robpike.io/ivy",
    sum = "h1:oCFPKLI9pk0tRme8d9GXdXFUat+Mb5TzvJ5M2LCfc8Q=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_traefik_yaegi",
    importpath = "github.com/traefik/yaegi",
    sum = "h1:Dg6hYcDtxWY3L9jP2OjoK+LSX59wF0sckJXT/cYNqA4=",
    version = "v0.9.14",
)

go_repository(
    name = "com_github_rancher_dapper",
    importpath = "github.com/rancher/dapper",
    sum = "h1:HtlWLY7MIbBsr+1ei4m4eixIV0+TXS7Kwx+c8Y7T+mE=",
    version = "v0.5.5",
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
    sum = "h1:GXJVqfe8nwaIUYuMwUIoCJXU1a+Hq8MjhMWGjsWLbYI=",
    version = "v2.0.5",
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
    sum = "h1:7UCwP93aiSfvWpapti8g88vVVGp2qqtGyePsSuDafo4=",
    version = "v1.0.0",
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
    sum = "h1:KhXCKH7Rpd+1YYABK0qcN7s8146Kfatzk5SCy8nAFAI=",
    version = "v0.8.2",
)

go_repository(
    name = "com_github_refraction_networking_utls",
    importpath = "github.com/refraction-networking/utls",
    sum = "h1:vIkvetWOJZSADSKCF9MLTsQNW2httdBmYz47dQQteP8=",
    version = "v0.0.0-20200601200209-ada0bb9b38a0",
)

go_repository(
    name = "com_github_txthinking_runnergroup",
    importpath = "github.com/txthinking/runnergroup",
    sum = "h1:vlDgnShahmE2XLslpr0hnzxfAmSj3JLX2CYi8Xct7G4=",
    version = "v0.0.0-20200327135940-540a793bb997",
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
    sum = "h1:cZhhbV8+DE0Y1kotwhr1a3RC3kFO7AtuZ4GLr3qKSc8=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_mholt_caddy_l4",
    importpath = "github.com/mholt/caddy-l4",
    sum = "h1:M3Ss1Vq4NdnvTLY2xH6/lR8ixvgAd9Iy6gQzKwGhCdk=",
    version = "v0.0.0-20210209073014-d1d54b015e34",
)

go_repository(
    name = "com_github_abiosoft_caddy_json_schema",
    importpath = "github.com/abiosoft/caddy-json-schema",
    sum = "h1:8pK1q4nCIgA98a/vG7gu24rdc8NTrztVN0ZOih7Meg0=",
    version = "v0.0.0-20200527180432-2d0cb96ed8ea",
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
    sum = "h1:PlZu0n3Tuv04TzpfPbrnI0HW/YwodEXDS+oPKahKF0Q=",
    version = "v0.0.1",
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
    sum = "h1:VnaNNDFI7lNzceTDRuusgsiVhwdqmark30ySAly1zuo=",
    version = "v2.11.0",
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
    sum = "h1:r7vGuS5akxOnR4JQSkko62RJ1ReCMXxQRPtxsiFMBOs=",
    version = "v1.0.0",
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
    sum = "h1:YvTNdFzX6+W5m9msiYg/zpkSURPPtOlzbqYjrFn7Yt4=",
    version = "v0.5.7",
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
    sum = "h1:w+iIsaOQNcT7OZ575w+acHgRric5iCyQh+xv+KJ4HB8=",
    version = "v0.0.0-20170929234023-d6e3b3328b78",
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
    sum = "h1:+0HFd5KSZ/mm3JmhmrDukiId5iR6w4+BdFtfSy4yWIc=",
    version = "v0.0.3-0.20200106085610-5cbc8cc4026c",
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
    sum = "h1:xvUo53O5MRZhVMJAxWCJcS5HHrqAiAG9SJ1LpMu6aAI=",
    version = "v0.0.0-20191021191039-0944d244cd40",
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
    sum = "h1:kq/SbG2BCKLkDKkjQf5OWwKWUKj1lgs3lFI4PxnR5lg=",
    version = "v22.1.0",
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
    sum = "h1:TcekIExNqud5crz4xD2pavyTgWiPvpYe4Xau31I0PRk=",
    version = "v3.2.2+incompatible",
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
    sum = "h1:ZqHaoEF7TBzh4jzPmqVhE/5A1z9of6orkAe5uHoAeME=",
    version = "v5.0.3",
)

go_repository(
    name = "com_github_google_addlicense",
    importpath = "github.com/google/addlicense",
    sum = "h1:JypWNzPMSgH5yL0NvFoAIsDRlKFgL0AsS3GO5bg4Pto=",
    version = "v0.0.0-20200906110928-a0294312aa76",
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
    sum = "h1:rzf0wL0CHVc8CEsgyygG0Mn9CNCCPZqOPaz8RiiHYQk=",
    version = "v0.0.0-20201216013528-df9cb8a40635",
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
    sum = "h1:TihvEz9MPj2u0KWds6E2OBUXfwaL4qRJ33c7HGiJpqk=",
    version = "v0.0.14",
)

go_repository(
    name = "io_k8s_sigs_metrics_server",
    importpath = "sigs.k8s.io/metrics-server",
    sum = "h1:7kHiDF+sJEZ+ZySHw/v3eIyNhhar7NPuGkc7jhqRl7M=",
    version = "v0.4.2",
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
    name = "com_github_mathetake_gasm",
    importpath = "github.com/mathetake/gasm",
    sum = "h1:GOar/82ZJdJ9COsZmNfSKgI/h5eYlBj6WvmUkuFEuDE=",
    version = "v0.0.0-20200928142744-80e74517647c",
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
    sum = "h1:rtRG4N6Ct7GNssATwgpvMGfnjnwfjnu/Zs9W3Ikzq+M=",
    version = "v1.4.4",
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
    sum = "h1:DO0qK9hESxkOTWyd/93hjYBRL7MdVSFqaXdcR7n4pVY=",
    version = "v0.7.0",
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
    sum = "h1:ipVmcsJcz90+0u0rkn07wKwcXNiw1W/3E2w3rxIPpWM=",
    version = "v1.7.6",
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
    sum = "h1:hjx7cr3VzWXr5UxEPG80S05kAAhjjACXHGJjpZJFl6E=",
    version = "v0.0.0-20210129020346-65fa249df79a",
)

go_repository(
    name = "com_github_portainer_libcrypto",
    importpath = "github.com/portainer/libcrypto",
    sum = "h1:ScnP50yr/HhfMizrFrDzlp4nflb9Y357YQlIUDMEOCI=",
    version = "v0.0.0-20201216185936-7a703a1ea452",
)

go_repository(
    name = "com_github_portainer_libhttp",
    importpath = "github.com/portainer/libhttp",
    sum = "h1:aHpxpXbYFEQY+qKufNNC5U09rfvfo0MLPCgxt/hJlx4=",
    version = "v0.0.0-20201216185909-d20481a3da82",
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
    sum = "h1:Fv93L3KKckEcEHR3oApXVzyBTDA8WAm6VXhPE00N3f8=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_containerd_console",
    importpath = "github.com/containerd/console",
    sum = "h1:u7SFAJyRqWcG6ogaMAx3KjSTy1e3hT9QxqX7Jco7dRc=",
    version = "v1.0.1",
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
    sum = "h1:PvuK4E3D5S5q6IqsPDCy928FhP0LUIGcmZ/Yhgp5Djw=",
    version = "v1.0.1",
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
    sum = "h1:8sXhOn0uLys67V8EsXLc6eszDs8VXWxL3iRvebPhedY=",
    version = "v1.11.3",
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
    sum = "h1:1KInV3Huv18akCu58V7lzNlt+jFmqlu1EaErnEHE/VM=",
    version = "v0.4.0",
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
    sum = "h1:x2UMpOOVf3kQ8arv/EsDGwim8PTNqzL1/EYDr/+scOM=",
    version = "v1.0.0-rc93",
)

go_repository(
    name = "com_github_opencontainers_runtime_spec",
    importpath = "github.com/opencontainers/runtime-spec",
    sum = "h1:pNa8metDkwZjb9g4T8s+krQ+HRgZAkqnXml+wNir/+s=",
    version = "v1.0.3-0.20200929063507-e6143ca7d51d",
)

go_repository(
    name = "com_github_opencontainers_selinux",
    importpath = "github.com/opencontainers/selinux",
    sum = "h1:+77ba4ar4jsCbL1GLbFL8fFM57w6suPfSS9PDLDY7KM=",
    version = "v1.8.0",
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
    sum = "h1:1iyaYNBLmP6L0220aDnYQpo1QEV4t4hJ+xEEhhJH8j0=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_vishvananda_netns",
    importpath = "github.com/vishvananda/netns",
    sum = "h1:4hwBBUfQCFe3Cym0ZtKyq7L16eZUtYKs+BaHDN6mAns=",
    version = "v0.0.0-20200728191858-db3c7e526aae",
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
    sum = "h1:qWj4qVYZ95vLWwqyNJCQg7rDsG5wPdze0UaPolH7DUk=",
    version = "v0.0.0-20200531161412-0dbf7f05ba59",
)

go_repository(
    name = "com_github_containerd_continuity",
    importpath = "github.com/containerd/continuity",
    sum = "h1:TP+534wVlf61smEIq1nwLLAjQVEK2EADoW3CX9AuT+8=",
    version = "v0.0.0-20190426062206-aaeac12a7ffc",
)

go_repository(
    name = "com_github_containerd_fifo",
    importpath = "github.com/containerd/fifo",
    sum = "h1:PUD50EuOMkXVcpBIA/R95d56duJR9VxhwncsFbNnxW4=",
    version = "v0.0.0-20190226154929-a9fb20d87448",
)

go_repository(
    name = "com_github_containerd_go_runc",
    importpath = "github.com/containerd/go-runc",
    sum = "h1:esQOJREg8nw8aXj6uCN5dfW5cKUBiEJ/+nni1Q/D/sw=",
    version = "v0.0.0-20180907222934-5a6d9f37cfa3",
)

go_repository(
    name = "com_github_containernetworking_cni",
    importpath = "github.com/containernetworking/cni",
    sum = "h1:BT9lpgGoH4jw3lFC7Odz2prU5ruiYKcgAjMCbgybcKI=",
    version = "v0.8.0",
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
    sum = "h1:1xpVY4dSUSbW3PcSGxZJhI8Z+CJiqbd933kM7HIinTc=",
    version = "v0.8.10-0.20200715222032-5eafd1556990",
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
    sum = "h1:gPeOspTx01shAQuFhae+O7ZfU0WKlD7RE5L0pooU0g8=",
    version = "v1.20.4",
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
    sum = "h1:A4Jv4ZCaV3AFJeGh5mGwkz4iuWUYMlQ7IoO/GTuSuLo=",
    version = "v0.5.0",
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
    sum = "h1:JrA3N5BdIc7+rtWgmP9zqw341BmimYmEtuU4lVBFgHY=",
    version = "v4.36.2",
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
    sum = "h1:yVW0v4zDXzJo1i8G9G3vtvNpyzhvtLalO34BsN/K88E=",
    version = "v0.0.0-20200716203538-1e47f820817f",
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
    sum = "h1:2HQmlpI3yI9deH18Q6xiSOIjXD4sLI55Y/gfpa8/558=",
    version = "v0.0.0-20191017083524-a8ff7f821017",
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
    sum = "h1:CG6TE5H9/JXsFWJCfoIVpKFIkFe6ysEuHirp4DxCsHI=",
    version = "v0.9.2",
)

go_repository(
    name = "com_github_hashicorp_go_retryablehttp",
    importpath = "github.com/hashicorp/go-retryablehttp",
    sum = "h1:92lWxgpa+fF3FozM4B3UZtHZMJX8T5XT+TFdCxsPyWs=",
    version = "v0.6.8",
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
    sum = "h1:j7+1HpAFS1zy5+Q4qx1fWh90gTKwiN4QCGoY9TWyyO4=",
    version = "v0.0.0-20170206155736-9520e82c474b",
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
    sum = "h1:Klu98tQ9Z1t23gvC7p7sCmvxkZxLhBHLNyrUPsWsYFg=",
    version = "v0.0.0-20171111012221-fa338ed9e9ec",
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
    sum = "h1:NZs4qyJdoFpRTX4Lg+jYfqSNE7xJskzEl3kJxgPAjDc=",
    version = "v1.0.3-0.20210322055141-7532f6617482",
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
    sum = "h1:nvV9bUm0PD4axZ0qwVvv995HvsKf1joidN1uSyI4J6s=",
    version = "v0.0.0-20210105070211-eda8e5aa2223",
)

go_repository(
    name = "com_github_caddy_dns_digitalocean",
    importpath = "github.com/caddy-dns/digitalocean",
    sum = "h1:hj7y6xn90dsHLSzHqXOAm+V/xXLPbYeHkatts+r3oEA=",
    version = "v0.0.0-20210408173619-385f9346b5ac",
)

go_repository(
    name = "com_github_caddy_dns_dnspod",
    importpath = "github.com/caddy-dns/dnspod",
    sum = "h1:ZQcuP+4tqevUUQAb6/FyF01cxWsA1rdLHns841p3XlM=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_caddy_dns_duckdns",
    importpath = "github.com/caddy-dns/duckdns",
    sum = "h1:NJ/8bF1o9KnYSRvpf5dD7uGjNWBqdYqyodBZC7KKheA=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_caddy_dns_route53",
    importpath = "github.com/caddy-dns/route53",
    sum = "h1:0D6tC6TZf60F+EiG5Mr8XKQwSWUwKhyiFwyI4zdDe3M=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_caddy_dns_vultr",
    importpath = "github.com/caddy-dns/vultr",
    sum = "h1:+9RmeGYfnjoIf4O2EwLLPOs46k/CcWsNaGpXK8BP4vI=",
    version = "v0.0.0-20210105121117-3162aa6b9c27",
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
    sum = "h1:hOqjF6rSvJQVaImoZE30bnzB+wgxFtkGXmVqnqcjIqY=",
    version = "v1.0.1",
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
    sum = "h1:Jx0AoxHtj2NMwxHByM8VmcqvGMa3lEu28xVDArhSi7E=",
    version = "v0.0.0-20200528144945-97886e7873b1",
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
    sum = "h1:/nd9nPOZ66kwSIodPZ2GWgtkJ6jMzbAA4RKggYosodA=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_libdns_duckdns",
    importpath = "github.com/libdns/duckdns",
    sum = "h1:GFDYvMtLuOZlNlTeUHAh5ZjtDdBHbdHKJNdWurygPtw=",
    version = "v0.0.0-20201229185916-cd405ff644b9",
)

go_repository(
    name = "com_github_libdns_route53",
    importpath = "github.com/libdns/route53",
    sum = "h1:LnMHXiOwoDXKDR8w2LcjGttalZTJeuaXx4lMoLrBCAo=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_libdns_vultr",
    importpath = "github.com/libdns/vultr",
    sum = "h1:i65uWz6ebW7T8bMqnJhzwj1B0HYB8mg/GXueGPxLrOs=",
    version = "v0.0.0-20201128180404-1d5ee21ea62f",
)

go_repository(
    name = "com_github_vultr_govultr_v2",
    importpath = "github.com/vultr/govultr/v2",
    sum = "h1:+lAtqfWy3g9VwL7tT2Fpyad8Vv4MxOhT/NU8O5dk+EQ=",
    version = "v2.0.0",
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
    sum = "h1:u/UEqS66A5ckRmS4yNpjmVH56sVtS/RfclBAYocb4as=",
    version = "v0.0.0-20180327030543-2492fe189ae6",
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
    sum = "h1:qsup4IcBdlmsnGfqyLl4Ntn3C2XCCuKAE7DwHpScyUo=",
    version = "v1.0.0",
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
    sum = "h1:UQZhZ2O0vMHr2cI+DC1Mbh0TJxzA3RcLoMsFw+aXw7E=",
    version = "v0.0.0-20190924025748-f65c72e2690d",
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
    sum = "h1:bO0cdWOtgJk7rlRCOObS/TqtbvAisBWSoWpmdJEOius=",
    version = "v0.0.0-20190401073227-519ff4ea1882",
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
    sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    importpath = "github.com/census-instrumentation/opencensus-proto",
    sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_cespare_xxhash_v2",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=",
    version = "v2.1.1",
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
    sum = "h1:07n33Z8lZxZ2qwegKbObQohDhXDQxiMMz1NOUGYlesw=",
    version = "v1.1.11",
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
    sum = "h1:vFgEHPqWBTp4pTjdLwjAA4bSo3gvIGOYwuJTlEjVBCw=",
    version = "v20.10.2+incompatible",
)

go_repository(
    name = "com_github_ema_qdisc",
    importpath = "github.com/ema/qdisc",
    sum = "h1:I3hLsM87FSASssIrIOGwJCio31dvLkvpYDKn2+r31ec=",
    version = "v0.0.0-20190904071900-b82c76788043",
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
    sum = "h1:EmNYJhPYy0pOFjCx2PrgtaBXmee0iUX9hLlxE1xHOJE=",
    version = "v0.9.9-0.20201210154907-fd9021fe5dad",
)

go_repository(
    name = "com_github_envoyproxy_protoc_gen_validate",
    importpath = "github.com/envoyproxy/protoc-gen-validate",
    sum = "h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:hsms1Qyu0jgnwNXIxa+/V/PDsU6CfLf6CNO8H7IWoS4=",
    version = "v1.4.9",
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
    sum = "h1:r35w0JBADPZCVQijYebl6YMWWtHRqVEGt7kL2eBADRM=",
    version = "v1.3.0",
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
    sum = "h1:TrB8swr/68K7m9CcGut2g3UOihhbcbiMAYiuTXdEih4=",
    version = "v0.5.0",
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
    sum = "h1:s+PDl6lozQ+dEUtUtQnO7+A2iPG3sK1pI4liU+jxn90=",
    version = "v0.0.0-20190402143921-271e53dc4968",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    sum = "h1:DqDEcV5aeaTmdFBePNpYsp3FlcVH/2ISVVM9Qf8PSls=",
    version = "v1.3.1",
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
    sum = "h1:jlYHihg//f7RRwuPfptm04yp4s7O6Kw8EZiVYIGcH0g=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:JjCZWpVbqXDqFVmTfYWEVTMIYrL/NPdPSCHPJ0T/raM=",
    version = "v1.4.3",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:Khx7svrCpmxxtHBq5j2mp/xVjsi8hQMfNLvJFAlrGgU=",
    version = "v0.5.5",
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
    sum = "h1:qJYtXnJRWmpe7m/3XlyhrsLrEURqHRM2kxzoxXqyUDs=",
    version = "v1.2.0",
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
    sum = "h1:+/TMaTYc4QFitKJxsQ7Yye35DkWvkdLcvGKqM+x0Ufc=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_hodgesds_perf_utils",
    importpath = "github.com/hodgesds/perf-utils",
    sum = "h1:6BT6cddpouM0G7eHhLFS+XcqtPvhrzWbPreyIvgFEcg=",
    version = "v0.0.8",
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
    sum = "h1:nwOc1YaOrYJ37sEBrtWZrdqzK22hiJs3GpDmP3sR2Yw=",
    version = "v0.0.0-20200117123717-f846d4f6c1f4",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    sum = "h1:Kz6Cvnvv2wGdaG/V8yMvfkmNiXq9Ya2KUv4rouJJr68=",
    version = "v1.1.10",
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
    sum = "h1:Fmg33tUaq4/8ym9TJN1x7sLJnHVwhP33CNkpYV/7rwI=",
    version = "v0.2.1",
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
    sum = "h1:mCMFu6PgSozg9tDNMMK3g18oJBX7oYGrC09mS6CXfO4=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_lufia_iostat",
    importpath = "github.com/lufia/iostat",
    sum = "h1:Z1wa4Hhxwi8uSKfgRsFc5RLtt3SuFPIOgkiPGkUtHDY=",
    version = "v1.1.0",
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
    sum = "h1:wuysRhFDzyxgEmMf5xjvJ2M9dZoWAXNNr5LSBS7uHXY=",
    version = "v0.0.12",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    sum = "h1:CoZ3S2P7pvtP45xOtBw+/mDL2z0RKI576gSkzRRpdGg=",
    version = "v0.0.10",
)

go_repository(
    name = "com_github_mattn_go_shellwords",
    importpath = "github.com/mattn/go-shellwords",
    sum = "h1:Y7Xqm8piKOO3v10Thp7Z36h4FYFjt5xB//6XvOrs2Gw=",
    version = "v1.0.10",
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
    sum = "h1:mpdLgm+brq10nI9zM1BpX1kpDbh3NLl3RSnVq6ZSkfg=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_mdlayher_wifi",
    importpath = "github.com/mdlayher/wifi",
    sum = "h1:hZDujBrW3ye2xxdKNFYT59D4yCH5Q0zLuNBNtysKtok=",
    version = "v0.0.0-20190303161829-b1436901ddee",
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
    sum = "h1:pyyPFfGMnciYUk/mXpKkVmeMQjfXqt3FAJ2hy7tPiLA=",
    version = "v1.1.40",
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
    sum = "h1:9f412s+6RmYXLWZSEzVVgPGK7C2PphHj5RJrvfx9AWI=",
    version = "v1.0.1",
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
    sum = "h1:2mOpI4JVVPBN+WQRa0WKH2eXR+Ey+uK4n7Zj0aYpIQA=",
    version = "v1.14.0",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    sum = "h1:o0+MgICZLuZ7xjH7Vx6zS/zcu93/BEp1VwkIW1mEXCE=",
    version = "v1.10.1",
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
    sum = "h1:Rrch9mh17XcxvEu9D9DEpb4isxjGBtcevQjKvxPRQIU=",
    version = "v1.9.0",
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
    sum = "h1:4fgOnadei3EZvgRwxJ7RMpG1k1pOZth5Pc13tyspaKM=",
    version = "v0.15.0",
)

go_repository(
    name = "com_github_prometheus_node_exporter",
    importpath = "github.com/prometheus/node_exporter",
    sum = "h1:xTBtauxuNCMhuF4FKiNo2wDCuCAWgS3PoTlVbXLzNO0=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_prometheus_procfs",
    importpath = "github.com/prometheus/procfs",
    sum = "h1:wH4vA7pcjKuZzjF7lM8awk4fnuJO6idemZXoKnULUx4=",
    version = "v0.2.0",
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
    sum = "h1:qLvObTrvO/XRCqmkKxUlOBc48bI3efyDuAZe25QiF0w=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_rs_cors",
    importpath = "github.com/rs/cors",
    sum = "h1:+88SsELBHx5r+hZ8TCkggzSstaWNbDvThkVK8H6f9ik=",
    version = "v1.7.0",
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
    sum = "h1:rRF7gJ7t0E1bfqNLwMqgb59eb273kgi+GgLE/yEiDzs=",
    version = "v0.0.0-20200303194639-4e8294f9e9d5",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:nfhvjKcUMhBMVqbKHJlk5RPrrfYr/NMo3692g0dwfWU=",
    version = "v1.8.0",
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
    sum = "h1:nwc3DEeHmmLAfoZucVR881uASk0Mfjw8xYJ99tb5CcY=",
    version = "v1.7.0",
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
    sum = "h1:yu9Bs+KssCJNxu/9fzRag6QgzOnxoH1Q6TvIiD4L6rQ=",
    version = "v0.0.0-20200531111549-252709fcb919",
)

go_repository(
    name = "com_github_txthinking_x",
    importpath = "github.com/txthinking/x",
    sum = "h1:ngJOce33YJJT1PFTfC9ao7S27AfrUh11Dr3Bc+ooBdM=",
    version = "v0.0.0-20200330144832-5ad2416896a9",
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
    sum = "h1:oKpsiyKMfVpwR3zSAkQixGzlVE5ovitBuO0qSmCf0bI=",
    version = "v0.78.0",
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
    replace = "github.com/btwiuse/websocket",
    sum = "h1:Vhtco9ZzTOgUTj3R+YMbUBERszZssZOoxRYldrUDWwc=",
    version = "v1.8.6",
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
    sum = "h1:hcskBH5qZCOa7WpTUFUFvoebnSFZBYpjykLtjIp9DVk=",
    version = "v0.0.0-20210303154014-9728d6b83eeb",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:o1bcQ6imQMIOpdrO3SWf2z5RV72WbDwdXuK0MDlc8As=",
    version = "v1.36.0",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:wSOdpTq0/eI46Ez/LkDwIsAKA71YP2SRKBODiRWM0as=",
    version = "v0.0.0-20210314154223-e6e6c4f2bb5b",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:FR+oGxGfbQu1d+jglI3rCkjAjUnhRSZcUxr+DqlDLNo=",
    version = "v0.0.0-20200331195152-e8c3332aa8e5",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:2M3HP5CCK1Si9FQhwnzYhXdG6DXeebvUHFpre8QvbyI=",
    version = "v0.0.0-20201208152925-83fdc39ff7b5",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:Kvvh58BN8Y9/lBi7hTekvtMpm07eUZ0ck5pRHpsMWrY=",
    version = "v0.4.1",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:qWPm9rbaAMKs8Bq/9LRpbMqxWRVUAQwMI9fVrssnTfw=",
    version = "v0.0.0-20210226172049-e18ecbb05110",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:Y2nxrNrrWOZn5yjDEEVU3R7V9HGW5SWsw6B6YL/ZRFw=",
    version = "v0.0.0-20210311163135-5366d9dc1934",
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
    sum = "h1:46ULzRKLh1CwgRq2dC5SlBzEqqNCi8rreOZnNrbqcIY=",
    version = "v0.0.0-20210309074719-68d13333faf2",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:i6eZZ+zk0SOf0xgBpEpPD18qWcJda6q1sxt3S0kzyUQ=",
    version = "v0.3.5",
)

go_repository(
    name = "org_golang_x_time",
    importpath = "golang.org/x/time",
    sum = "h1:Hir2P/De0WpUhtrKGGjvSb2YxUgyZ7EFOSLIcSSpiwE=",
    version = "v0.0.0-20201208040808-7e3f01d25324",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:po9/4sTYwZU9lPhi1tOrb4hCv3qrhiQ77LZfGa2OjwY=",
    version = "v0.1.0",
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
    sum = "h1:LtuKNg6JMiaBKVQHKd6Phhvk+2GFp+pUcmDQgRjrds0=",
    version = "v1.0.0",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    sum = "h1:ADUqmZGgLDDfbSL9ZmPxKTybcoEYHgpYfELNoN+7hsw=",
    version = "v1.7.0",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    sum = "h1:y6IPFStTAIT5Ytl7/XYmHvzXQ7S3g/IeZW9hyZ5thw4=",
    version = "v1.6.0",
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
    sum = "h1:uFRZXykJGK9lLY4HtgSw44DnIcAM+kRBP7x5m+NpAOM=",
    version = "v1.16.0",
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
    sum = "h1:8GUt8eRujhVEGZFFEjBj46YV4rDjvGrNxb0KMWYkL2I=",
    version = "v0.0.0-20180917152333-f0300d1749da",
)

go_repository(
    name = "com_github_armon_go_radix",
    importpath = "github.com/armon/go-radix",
    sum = "h1:BUAU3CGlLvorLI26FmByPp2eC2qla6E1Tw+scpcg/to=",
    version = "v0.0.0-20180808171621-7fddfc383310",
)

go_repository(
    name = "com_github_aryann_difflib",
    importpath = "github.com/aryann/difflib",
    sum = "h1:pv34s756C4pEXnjgPfGYgdhg/ZdajGhyOvzx8k+23nw=",
    version = "v0.0.0-20170710044230-e206f873d14a",
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
    sum = "h1:LRwl+97B4D69Z7tz+eRUxJ1C7baBaIYhgrn5eLtua+Q=",
    version = "v1.37.10",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2",
    importpath = "github.com/aws/aws-sdk-go-v2",
    sum = "h1:qZ+woO4SamnH/eEbjM2IDLhRNwIwND/RQyVlBLp3Jqg=",
    version = "v0.18.0",
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
    sum = "h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_gogo_googleapis",
    importpath = "github.com/gogo/googleapis",
    sum = "h1:kFkMAZBNAn4j7K0GiZr8cRYzejq68VbheufiV3YuyFI=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_golang_groupcache",
    importpath = "github.com/golang/groupcache",
    sum = "h1:1r7pUrabqp18hOBcwBwiTsbnFeTZHV9eER/QT5JVZxY=",
    version = "v0.0.0-20200121045136-8c9f03a8e57e",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    sum = "h1:aeE13tS0IiQgFjYdoL8qN3K1N2bXXtI6Vi51/y7BpMw=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_google_btree",
    importpath = "github.com/google/btree",
    sum = "h1:0udJVsspx3VBr5FwtLhQQtuAsVc79tTq0ocGIPAU6qo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
    version = "v0.0.0-20181017120253-0766667cb4d1",
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
    sum = "h1:FlFbCRLd5Jr4iYXZufAvgWN6Ao0JrI5chLINnUXDDr0=",
    version = "v1.2.2",
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
    sum = "h1:hLrqtEDnRye3+sgx6z4qVLNuviH3MR5aQ0ykNJa/UYA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_cleanhttp",
    importpath = "github.com/hashicorp/go-cleanhttp",
    sum = "h1:dH3aiDG9Jvb5r5+bYHsikaOUIpcM0xvgMXVoDkXMzJM=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_hashicorp_go_immutable_radix",
    importpath = "github.com/hashicorp/go-immutable-radix",
    sum = "h1:AKDB1HM5PWEA7i4nhcpwOrO2byshxBjXVn/J/3+z5/0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_msgpack",
    importpath = "github.com/hashicorp/go-msgpack",
    sum = "h1:zKjpN5BK/P5lMYrLmBHdBULWbJ0XpYR+7NGzqkZzoD4=",
    version = "v0.5.3",
)

go_repository(
    name = "com_github_hashicorp_go_multierror",
    importpath = "github.com/hashicorp/go-multierror",
    sum = "h1:B9UzwGQJehnUY1yNrnwREHc3fGbC2xefo8g4TbElacI=",
    version = "v1.1.0",
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
    sum = "h1:Rqb66Oo1X/eSV1x66xbDccZjhJigjg0+e82kpwzSwCI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_sockaddr",
    importpath = "github.com/hashicorp/go-sockaddr",
    sum = "h1:GeH6tui99pF4NJgfnhp+L6+FfobzVW3Ah46sLo0ICXs=",
    version = "v1.0.0",
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
    sum = "h1:EmmoJme1matNzb+hMpDuR/0sbJSUisxyqBGG676r31M=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_hashicorp_serf",
    importpath = "github.com/hashicorp/serf",
    sum = "h1:YZ7UKsJv+hKjqGVUUbtE3HNj79Eln2oQ75tniF6iPt0=",
    version = "v0.8.2",
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
    sum = "h1:reN85Pxc5larApoH1keMBiu2GWtPqXQ1nc9gx+jOU+E=",
    version = "v1.2.0",
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
    sum = "h1:fzU/JVNcaqHQEcVFAKeR41fkiLdIPrefOvVG1VZ96U0=",
    version = "v1.0.0",
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
    sum = "h1:CpVNEelQCZBooIPDn+AR3NpivK/TIKU8bDxdASFVQag=",
    version = "v1.4.1",
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
    sum = "h1:Lgl0gzECD8GnQ5QCWA8o6BtfL6mDH5rQgM4/fX3avOs=",
    version = "v0.0.0-20180627143212-57f6aae5913c",
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
    sum = "h1:UFr9zpz4xgTnIE5yIMtWAMngCdZ9p/+q6lTbgelo80M=",
    version = "v0.0.0-20160712163229-9b3edd62028f",
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
    sum = "h1:0HKaf1o97UwFjHH9o5XsHUOF+tqmdA7KEzXLpiyaw0E=",
    version = "v0.1.4",
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
    sum = "h1:xghbfqPkxzxP3C/f3n5DdpAbdKLj4ZE4BWQI362l53M=",
    version = "v1.1.3",
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
    sum = "h1:j6JEOq5QWFker+d7mFQYOhjTZonQ7YkLTHm56dbn+yM=",
    version = "v0.0.0-20200427203606-3cfed13b9966",
)

go_repository(
    name = "com_github_urfave_cli",
    importpath = "github.com/urfave/cli",
    sum = "h1:lNq9sAHXK2qfdI8W+GRItjCEkI+2oR4d+MEHy1CKXoU=",
    version = "v1.22.5",
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
    sum = "h1:ruQGxdhGHe7FWOJPT0mKs5+pD2Xs1Bm/kdGlHO04FmM=",
    version = "v1.2.1",
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
    sum = "h1:XAzx9gjCb0Rxj7EoqcClPD1d5ZBxZJk0jbuoPHenBt0=",
    version = "v1.3.5",
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
    sum = "h1:12aHIhhQCpWtd3Rcp2WwbboB5W72tJHcjzyA9MCoHAw=",
    version = "v0.41.0",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:Ejskq+SyPohKW+1uil0JJMtmHCgJPJ/qWTxr8qp+R4c=",
    version = "v1.25.0",
)

go_repository(
    name = "com_github_cncf_udpa_go",
    importpath = "github.com/cncf/udpa/go",
    sum = "h1:cqQfy1jclcSy/FwLjemeg3SR1yaINm74aQyupQ0Bl8M=",
    version = "v0.0.0-20201120205902-5459f2c99403",
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
    sum = "h1:0hzRabrMN4tSTvMfnL3SCv1ZGeAP23ynzodBgaHeMeg=",
    version = "v1.11.7",
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
    sum = "h1:46PFijGLmAjMPwCCCo7Jf0W6f9slllCkkv7vyc1yOSg=",
    version = "v0.0.0-20200907205600-7a23bdc65eef",
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
    sum = "h1:DpHb9vJrZQEFMcVLFKAAGMUVX0XoRC0ptCthinRYm38=",
    version = "v0.4.5",
)

go_repository(
    name = "com_github_dchest_siphash",
    importpath = "github.com/dchest/siphash",
    sum = "h1:YWOShuhvg0GqbQpMa60QlCGtEyf7O7HC1Jf0VjdQ60M=",
    version = "v1.2.0",
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
    sum = "h1:06t8Fpg0CH1vu8OuOaGi4K5oDJl7DKqga4uPBUmKgWQ=",
    version = "v0.0.0-20210218034515-4e5ef1691e9f",
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
    sum = "h1:UOf2YwAzzaUjY5mdpMuLfSw0vz62iIFYk7oJQkuhlGw=",
    version = "v0.1.0",
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
    sum = "h1:2JqaNE1hGdABW2YbA3TenkO7RiPFRvSWnEnGqWh9sHE=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_klauspost_reedsolomon",
    importpath = "github.com/klauspost/reedsolomon",
    sum = "h1:qCL7LZlv17xMixl55nq2/Oa1Y86nfO8EqDfv2GHND54=",
    version = "v1.9.9",
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
    sum = "h1:eCDQqvGBB+kCTkA0XrAFtNe81FMa0/fn4QSoeAbmiF4=",
    version = "v0.19.3",
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
    sum = "h1:p7xbxYTzzfXghR1kpsJDeoVVRRWAotKc8u7FP/N48rU=",
    version = "v0.0.0-20200213042214-218e1707c495",
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
    sum = "h1:orbzVR/BM+SW2K6n9Sy3u/7UneOw/KxpEbCz+hjCCOI=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_shadowsocks_shadowsocks_go",
    importpath = "github.com/shadowsocks/shadowsocks-go",
    sum = "h1:tJgNXb3S+RkB4kNPi6N5OmEWe3m+Y3Qs6LUMiNDAONM=",
    version = "v0.0.0-20170121203516-97a5c71f80ba",
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
    sum = "h1:7JVkAn5bvUJ7HtU08iW6UiD+UTmJTIToHCfeFzkcCxM=",
    version = "v1.3.2",
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
    sum = "h1:XmKBi9R6duxOB3lfc72wyrwiOY7X2Jl1wuI+RFOyMDE=",
    version = "v0.0.0-20201008161808-52c3e6f60cff",
)

go_repository(
    name = "org_torproject_git_pluggable_transports_goptlib_git",
    importpath = "git.torproject.org/pluggable-transports/goptlib.git",
    sum = "h1:PYcONLFUhr00kGrq7Mf14JRtoXHG7BOSKIfIha0Hu5Q=",
    version = "v0.0.0-20180321061416-7d56ec4f381e",
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
