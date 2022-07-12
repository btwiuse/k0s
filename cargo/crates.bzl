"""
@generated
cargo-raze generated Bazel file.

DO NOT EDIT! Replaced on runs of cargo-raze
"""

load("@bazel_tools//tools/build_defs/repo:git.bzl", "new_git_repository")  # buildifier: disable=load
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")  # buildifier: disable=load
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")  # buildifier: disable=load

def raze_fetch_remote_crates():
    """This function defines a collection of repos and should be called in a WORKSPACE file"""
    maybe(
        http_archive,
        name = "raze__actix__0_13_0",
        url = "https://crates.io/api/v1/crates/actix/0.13.0/download",
        type = "tar.gz",
        sha256 = "f728064aca1c318585bf4bb04ffcfac9e75e508ab4e8b1bd9ba5dfe04e2cbed5",
        strip_prefix = "actix-0.13.0",
        build_file = Label("//cargo/remote:BUILD.actix-0.13.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__actix_codec__0_5_0",
        url = "https://crates.io/api/v1/crates/actix-codec/0.5.0/download",
        type = "tar.gz",
        sha256 = "57a7559404a7f3573127aab53c08ce37a6c6a315c374a31070f3c91cd1b4a7fe",
        strip_prefix = "actix-codec-0.5.0",
        build_file = Label("//cargo/remote:BUILD.actix-codec-0.5.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__actix_http__3_2_1",
        url = "https://crates.io/api/v1/crates/actix-http/3.2.1/download",
        type = "tar.gz",
        sha256 = "6f9ffb6db08c1c3a1f4aef540f1a63193adc73c4fbd40b75a95fc8c5258f6e51",
        strip_prefix = "actix-http-3.2.1",
        build_file = Label("//cargo/remote:BUILD.actix-http-3.2.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__actix_macros__0_2_3",
        url = "https://crates.io/api/v1/crates/actix-macros/0.2.3/download",
        type = "tar.gz",
        sha256 = "465a6172cf69b960917811022d8f29bc0b7fa1398bc4f78b3c466673db1213b6",
        strip_prefix = "actix-macros-0.2.3",
        build_file = Label("//cargo/remote:BUILD.actix-macros-0.2.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__actix_router__0_5_0",
        url = "https://crates.io/api/v1/crates/actix-router/0.5.0/download",
        type = "tar.gz",
        sha256 = "eb60846b52c118f2f04a56cc90880a274271c489b2498623d58176f8ca21fa80",
        strip_prefix = "actix-router-0.5.0",
        build_file = Label("//cargo/remote:BUILD.actix-router-0.5.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__actix_rt__2_7_0",
        url = "https://crates.io/api/v1/crates/actix-rt/2.7.0/download",
        type = "tar.gz",
        sha256 = "7ea16c295198e958ef31930a6ef37d0fb64e9ca3b6116e6b93a8bdae96ee1000",
        strip_prefix = "actix-rt-2.7.0",
        build_file = Label("//cargo/remote:BUILD.actix-rt-2.7.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__actix_server__2_1_1",
        url = "https://crates.io/api/v1/crates/actix-server/2.1.1/download",
        type = "tar.gz",
        sha256 = "0da34f8e659ea1b077bb4637948b815cd3768ad5a188fdcd74ff4d84240cd824",
        strip_prefix = "actix-server-2.1.1",
        build_file = Label("//cargo/remote:BUILD.actix-server-2.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__actix_service__2_0_2",
        url = "https://crates.io/api/v1/crates/actix-service/2.0.2/download",
        type = "tar.gz",
        sha256 = "3b894941f818cfdc7ccc4b9e60fa7e53b5042a2e8567270f9147d5591893373a",
        strip_prefix = "actix-service-2.0.2",
        build_file = Label("//cargo/remote:BUILD.actix-service-2.0.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__actix_utils__3_0_0",
        url = "https://crates.io/api/v1/crates/actix-utils/3.0.0/download",
        type = "tar.gz",
        sha256 = "e491cbaac2e7fc788dfff99ff48ef317e23b3cf63dbaf7aaab6418f40f92aa94",
        strip_prefix = "actix-utils-3.0.0",
        build_file = Label("//cargo/remote:BUILD.actix-utils-3.0.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__actix_web__4_1_0",
        url = "https://crates.io/api/v1/crates/actix-web/4.1.0/download",
        type = "tar.gz",
        sha256 = "a27e8fe9ba4ae613c21f677c2cfaf0696c3744030c6f485b34634e502d6bb379",
        strip_prefix = "actix-web-4.1.0",
        build_file = Label("//cargo/remote:BUILD.actix-web-4.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__actix_web_actors__4_1_0",
        url = "https://crates.io/api/v1/crates/actix-web-actors/4.1.0/download",
        type = "tar.gz",
        sha256 = "31efe7896f3933ce03dd4710be560254272334bb321a18fd8ff62b1a557d9d19",
        strip_prefix = "actix-web-actors-4.1.0",
        build_file = Label("//cargo/remote:BUILD.actix-web-actors-4.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__actix_web_codegen__4_0_1",
        url = "https://crates.io/api/v1/crates/actix-web-codegen/4.0.1/download",
        type = "tar.gz",
        sha256 = "5f270541caec49c15673b0af0e9a00143421ad4f118d2df7edcb68b627632f56",
        strip_prefix = "actix-web-codegen-4.0.1",
        build_file = Label("//cargo/remote:BUILD.actix-web-codegen-4.0.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__actix_derive__0_6_0",
        url = "https://crates.io/api/v1/crates/actix_derive/0.6.0/download",
        type = "tar.gz",
        sha256 = "6d44b8fee1ced9671ba043476deddef739dd0959bf77030b26b738cc591737a7",
        strip_prefix = "actix_derive-0.6.0",
        build_file = Label("//cargo/remote:BUILD.actix_derive-0.6.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ahash__0_7_6",
        url = "https://crates.io/api/v1/crates/ahash/0.7.6/download",
        type = "tar.gz",
        sha256 = "fcb51a0695d8f838b1ee009b3fbf66bda078cd64590202a864a8f3e8c4315c47",
        strip_prefix = "ahash-0.7.6",
        build_file = Label("//cargo/remote:BUILD.ahash-0.7.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__aho_corasick__0_7_18",
        url = "https://crates.io/api/v1/crates/aho-corasick/0.7.18/download",
        type = "tar.gz",
        sha256 = "1e37cfd5e7657ada45f742d6e99ca5788580b5c529dc78faf11ece6dc702656f",
        strip_prefix = "aho-corasick-0.7.18",
        build_file = Label("//cargo/remote:BUILD.aho-corasick-0.7.18.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__atty__0_2_14",
        url = "https://crates.io/api/v1/crates/atty/0.2.14/download",
        type = "tar.gz",
        sha256 = "d9b39be18770d11421cdb1b9947a45dd3f37e93092cbf377614828a319d5fee8",
        strip_prefix = "atty-0.2.14",
        build_file = Label("//cargo/remote:BUILD.atty-0.2.14.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__autocfg__1_1_0",
        url = "https://crates.io/api/v1/crates/autocfg/1.1.0/download",
        type = "tar.gz",
        sha256 = "d468802bab17cbc0cc575e9b053f41e72aa36bfa6b7f55e3529ffa43161b97fa",
        strip_prefix = "autocfg-1.1.0",
        build_file = Label("//cargo/remote:BUILD.autocfg-1.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__backoff__0_4_0",
        url = "https://crates.io/api/v1/crates/backoff/0.4.0/download",
        type = "tar.gz",
        sha256 = "b62ddb9cb1ec0a098ad4bbf9344d0713fa193ae1a80af55febcff2627b6a00c1",
        strip_prefix = "backoff-0.4.0",
        build_file = Label("//cargo/remote:BUILD.backoff-0.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__base64__0_13_0",
        url = "https://crates.io/api/v1/crates/base64/0.13.0/download",
        type = "tar.gz",
        sha256 = "904dfeac50f3cdaba28fc6f57fdcddb75f49ed61346676a78c4ffe55877802fd",
        strip_prefix = "base64-0.13.0",
        build_file = Label("//cargo/remote:BUILD.base64-0.13.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__bitflags__1_3_2",
        url = "https://crates.io/api/v1/crates/bitflags/1.3.2/download",
        type = "tar.gz",
        sha256 = "bef38d45163c2f1dde094a7dfd33ccf595c92905c8f8f4fdc18d06fb1037718a",
        strip_prefix = "bitflags-1.3.2",
        build_file = Label("//cargo/remote:BUILD.bitflags-1.3.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__block_buffer__0_10_2",
        url = "https://crates.io/api/v1/crates/block-buffer/0.10.2/download",
        type = "tar.gz",
        sha256 = "0bf7fe51849ea569fd452f37822f606a5cabb684dc918707a0193fd4664ff324",
        strip_prefix = "block-buffer-0.10.2",
        build_file = Label("//cargo/remote:BUILD.block-buffer-0.10.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__bytes__0_5_6",
        url = "https://crates.io/api/v1/crates/bytes/0.5.6/download",
        type = "tar.gz",
        sha256 = "0e4cec68f03f32e44924783795810fa50a7035d8c8ebe78580ad7e6c703fba38",
        strip_prefix = "bytes-0.5.6",
        build_file = Label("//cargo/remote:BUILD.bytes-0.5.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__bytes__1_1_0",
        url = "https://crates.io/api/v1/crates/bytes/1.1.0/download",
        type = "tar.gz",
        sha256 = "c4872d67bab6358e59559027aa3b9157c53d9358c51423c17554809a8858e0f8",
        strip_prefix = "bytes-1.1.0",
        build_file = Label("//cargo/remote:BUILD.bytes-1.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__bytestring__1_1_0",
        url = "https://crates.io/api/v1/crates/bytestring/1.1.0/download",
        type = "tar.gz",
        sha256 = "86b6a75fd3048808ef06af5cd79712be8111960adaf89d90250974b38fc3928a",
        strip_prefix = "bytestring-1.1.0",
        build_file = Label("//cargo/remote:BUILD.bytestring-1.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cc__1_0_73",
        url = "https://crates.io/api/v1/crates/cc/1.0.73/download",
        type = "tar.gz",
        sha256 = "2fff2a6927b3bb87f9595d67196a70493f627687a71d87a0d692242c33f58c11",
        strip_prefix = "cc-1.0.73",
        build_file = Label("//cargo/remote:BUILD.cc-1.0.73.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cfg_if__0_1_10",
        url = "https://crates.io/api/v1/crates/cfg-if/0.1.10/download",
        type = "tar.gz",
        sha256 = "4785bdd1c96b2a846b2bd7cc02e86b6b3dbf14e7e53446c4f54c92a361040822",
        strip_prefix = "cfg-if-0.1.10",
        build_file = Label("//cargo/remote:BUILD.cfg-if-0.1.10.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cfg_if__1_0_0",
        url = "https://crates.io/api/v1/crates/cfg-if/1.0.0/download",
        type = "tar.gz",
        sha256 = "baf1de4339761588bc0619e3cbc0120ee582ebb74b53b4efbf79117bd2da40fd",
        strip_prefix = "cfg-if-1.0.0",
        build_file = Label("//cargo/remote:BUILD.cfg-if-1.0.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__chrono__0_4_19",
        url = "https://crates.io/api/v1/crates/chrono/0.4.19/download",
        type = "tar.gz",
        sha256 = "670ad68c9088c2a963aaa298cb369688cf3f9465ce5e2d4ca10e6e0098a1ce73",
        strip_prefix = "chrono-0.4.19",
        build_file = Label("//cargo/remote:BUILD.chrono-0.4.19.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__colored__2_0_0",
        url = "https://crates.io/api/v1/crates/colored/2.0.0/download",
        type = "tar.gz",
        sha256 = "b3616f750b84d8f0de8a58bda93e08e2a81ad3f523089b05f1dffecab48c6cbd",
        strip_prefix = "colored-2.0.0",
        build_file = Label("//cargo/remote:BUILD.colored-2.0.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__convert_case__0_4_0",
        url = "https://crates.io/api/v1/crates/convert_case/0.4.0/download",
        type = "tar.gz",
        sha256 = "6245d59a3e82a7fc217c5828a6692dbc6dfb63a0c8c90495621f7b9d79704a0e",
        strip_prefix = "convert_case-0.4.0",
        build_file = Label("//cargo/remote:BUILD.convert_case-0.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cpufeatures__0_2_2",
        url = "https://crates.io/api/v1/crates/cpufeatures/0.2.2/download",
        type = "tar.gz",
        sha256 = "59a6001667ab124aebae2a495118e11d30984c3a653e99d86d58971708cf5e4b",
        strip_prefix = "cpufeatures-0.2.2",
        build_file = Label("//cargo/remote:BUILD.cpufeatures-0.2.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__crossbeam_channel__0_5_5",
        url = "https://crates.io/api/v1/crates/crossbeam-channel/0.5.5/download",
        type = "tar.gz",
        sha256 = "4c02a4d71819009c192cf4872265391563fd6a84c81ff2c0f2a7026ca4c1d85c",
        strip_prefix = "crossbeam-channel-0.5.5",
        build_file = Label("//cargo/remote:BUILD.crossbeam-channel-0.5.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__crossbeam_utils__0_8_10",
        url = "https://crates.io/api/v1/crates/crossbeam-utils/0.8.10/download",
        type = "tar.gz",
        sha256 = "7d82ee10ce34d7bc12c2122495e7593a9c41347ecdd64185af4ecf72cb1a7f83",
        strip_prefix = "crossbeam-utils-0.8.10",
        build_file = Label("//cargo/remote:BUILD.crossbeam-utils-0.8.10.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__crypto_common__0_1_4",
        url = "https://crates.io/api/v1/crates/crypto-common/0.1.4/download",
        type = "tar.gz",
        sha256 = "5999502d32b9c48d492abe66392408144895020ec4709e549e840799f3bb74c0",
        strip_prefix = "crypto-common-0.1.4",
        build_file = Label("//cargo/remote:BUILD.crypto-common-0.1.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__darling__0_14_1",
        url = "https://crates.io/api/v1/crates/darling/0.14.1/download",
        type = "tar.gz",
        sha256 = "4529658bdda7fd6769b8614be250cdcfc3aeb0ee72fe66f9e41e5e5eb73eac02",
        strip_prefix = "darling-0.14.1",
        build_file = Label("//cargo/remote:BUILD.darling-0.14.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__darling_core__0_14_1",
        url = "https://crates.io/api/v1/crates/darling_core/0.14.1/download",
        type = "tar.gz",
        sha256 = "649c91bc01e8b1eac09fb91e8dbc7d517684ca6be8ebc75bb9cafc894f9fdb6f",
        strip_prefix = "darling_core-0.14.1",
        build_file = Label("//cargo/remote:BUILD.darling_core-0.14.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__darling_macro__0_14_1",
        url = "https://crates.io/api/v1/crates/darling_macro/0.14.1/download",
        type = "tar.gz",
        sha256 = "ddfc69c5bfcbd2fc09a0f38451d2daf0e372e367986a83906d1b0dbc88134fb5",
        strip_prefix = "darling_macro-0.14.1",
        build_file = Label("//cargo/remote:BUILD.darling_macro-0.14.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__derivative__2_2_0",
        url = "https://crates.io/api/v1/crates/derivative/2.2.0/download",
        type = "tar.gz",
        sha256 = "fcc3dd5e9e9c0b295d6e1e4d811fb6f157d5ffd784b8d202fc62eac8035a770b",
        strip_prefix = "derivative-2.2.0",
        build_file = Label("//cargo/remote:BUILD.derivative-2.2.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__derive_more__0_99_17",
        url = "https://crates.io/api/v1/crates/derive_more/0.99.17/download",
        type = "tar.gz",
        sha256 = "4fb810d30a7c1953f91334de7244731fc3f3c10d7fe163338a35b9f640960321",
        strip_prefix = "derive_more-0.99.17",
        build_file = Label("//cargo/remote:BUILD.derive_more-0.99.17.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__digest__0_10_3",
        url = "https://crates.io/api/v1/crates/digest/0.10.3/download",
        type = "tar.gz",
        sha256 = "f2fb860ca6fafa5552fb6d0e816a69c8e49f0908bf524e30a90d97c85892d506",
        strip_prefix = "digest-0.10.3",
        build_file = Label("//cargo/remote:BUILD.digest-0.10.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__dirs_next__2_0_0",
        url = "https://crates.io/api/v1/crates/dirs-next/2.0.0/download",
        type = "tar.gz",
        sha256 = "b98cf8ebf19c3d1b223e151f99a4f9f0690dca41414773390fc824184ac833e1",
        strip_prefix = "dirs-next-2.0.0",
        build_file = Label("//cargo/remote:BUILD.dirs-next-2.0.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__dirs_sys_next__0_1_2",
        url = "https://crates.io/api/v1/crates/dirs-sys-next/0.1.2/download",
        type = "tar.gz",
        sha256 = "4ebda144c4fe02d1f7ea1a7d9641b6fc6b580adcfa024ae48797ecdeb6825b4d",
        strip_prefix = "dirs-sys-next-0.1.2",
        build_file = Label("//cargo/remote:BUILD.dirs-sys-next-0.1.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__dyn_clone__1_0_6",
        url = "https://crates.io/api/v1/crates/dyn-clone/1.0.6/download",
        type = "tar.gz",
        sha256 = "140206b78fb2bc3edbcfc9b5ccbd0b30699cfe8d348b8b31b330e47df5291a5a",
        strip_prefix = "dyn-clone-1.0.6",
        build_file = Label("//cargo/remote:BUILD.dyn-clone-1.0.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__either__1_7_0",
        url = "https://crates.io/api/v1/crates/either/1.7.0/download",
        type = "tar.gz",
        sha256 = "3f107b87b6afc2a64fd13cac55fe06d6c8859f12d4b14cbcdd2c67d0976781be",
        strip_prefix = "either-1.7.0",
        build_file = Label("//cargo/remote:BUILD.either-1.7.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__encoding_rs__0_8_31",
        url = "https://crates.io/api/v1/crates/encoding_rs/0.8.31/download",
        type = "tar.gz",
        sha256 = "9852635589dc9f9ea1b6fe9f05b50ef208c85c834a562f0c6abb1c475736ec2b",
        strip_prefix = "encoding_rs-0.8.31",
        build_file = Label("//cargo/remote:BUILD.encoding_rs-0.8.31.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__firestorm__0_5_1",
        url = "https://crates.io/api/v1/crates/firestorm/0.5.1/download",
        type = "tar.gz",
        sha256 = "2c5f6c2c942da57e2aaaa84b8a521489486f14e75e7fa91dab70aba913975f98",
        strip_prefix = "firestorm-0.5.1",
        build_file = Label("//cargo/remote:BUILD.firestorm-0.5.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fnv__1_0_7",
        url = "https://crates.io/api/v1/crates/fnv/1.0.7/download",
        type = "tar.gz",
        sha256 = "3f9eec918d3f24069decb9af1554cad7c880e2da24a9afd88aca000531ab82c1",
        strip_prefix = "fnv-1.0.7",
        build_file = Label("//cargo/remote:BUILD.fnv-1.0.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__foreign_types__0_3_2",
        url = "https://crates.io/api/v1/crates/foreign-types/0.3.2/download",
        type = "tar.gz",
        sha256 = "f6f339eb8adc052cd2ca78910fda869aefa38d22d5cb648e6485e4d3fc06f3b1",
        strip_prefix = "foreign-types-0.3.2",
        build_file = Label("//cargo/remote:BUILD.foreign-types-0.3.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__foreign_types_shared__0_1_1",
        url = "https://crates.io/api/v1/crates/foreign-types-shared/0.1.1/download",
        type = "tar.gz",
        sha256 = "00b0228411908ca8685dba7fc2cdd70ec9990a6e753e89b6ac91a84c40fbaf4b",
        strip_prefix = "foreign-types-shared-0.1.1",
        build_file = Label("//cargo/remote:BUILD.foreign-types-shared-0.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__form_urlencoded__1_0_1",
        url = "https://crates.io/api/v1/crates/form_urlencoded/1.0.1/download",
        type = "tar.gz",
        sha256 = "5fc25a87fa4fd2094bffb06925852034d90a17f0d1e05197d4956d3555752191",
        strip_prefix = "form_urlencoded-1.0.1",
        build_file = Label("//cargo/remote:BUILD.form_urlencoded-1.0.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fuchsia_zircon__0_3_3",
        url = "https://crates.io/api/v1/crates/fuchsia-zircon/0.3.3/download",
        type = "tar.gz",
        sha256 = "2e9763c69ebaae630ba35f74888db465e49e259ba1bc0eda7d06f4a067615d82",
        strip_prefix = "fuchsia-zircon-0.3.3",
        build_file = Label("//cargo/remote:BUILD.fuchsia-zircon-0.3.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fuchsia_zircon_sys__0_3_3",
        url = "https://crates.io/api/v1/crates/fuchsia-zircon-sys/0.3.3/download",
        type = "tar.gz",
        sha256 = "3dcaa9ae7725d12cdb85b3ad99a434db70b468c09ded17e012d86b5c1010f7a7",
        strip_prefix = "fuchsia-zircon-sys-0.3.3",
        build_file = Label("//cargo/remote:BUILD.fuchsia-zircon-sys-0.3.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures__0_3_21",
        url = "https://crates.io/api/v1/crates/futures/0.3.21/download",
        type = "tar.gz",
        sha256 = "f73fe65f54d1e12b726f517d3e2135ca3125a437b6d998caf1962961f7172d9e",
        strip_prefix = "futures-0.3.21",
        build_file = Label("//cargo/remote:BUILD.futures-0.3.21.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_channel__0_3_21",
        url = "https://crates.io/api/v1/crates/futures-channel/0.3.21/download",
        type = "tar.gz",
        sha256 = "c3083ce4b914124575708913bca19bfe887522d6e2e6d0952943f5eac4a74010",
        strip_prefix = "futures-channel-0.3.21",
        build_file = Label("//cargo/remote:BUILD.futures-channel-0.3.21.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_core__0_3_21",
        url = "https://crates.io/api/v1/crates/futures-core/0.3.21/download",
        type = "tar.gz",
        sha256 = "0c09fd04b7e4073ac7156a9539b57a484a8ea920f79c7c675d05d289ab6110d3",
        strip_prefix = "futures-core-0.3.21",
        build_file = Label("//cargo/remote:BUILD.futures-core-0.3.21.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_executor__0_3_21",
        url = "https://crates.io/api/v1/crates/futures-executor/0.3.21/download",
        type = "tar.gz",
        sha256 = "9420b90cfa29e327d0429f19be13e7ddb68fa1cccb09d65e5706b8c7a749b8a6",
        strip_prefix = "futures-executor-0.3.21",
        build_file = Label("//cargo/remote:BUILD.futures-executor-0.3.21.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_io__0_3_21",
        url = "https://crates.io/api/v1/crates/futures-io/0.3.21/download",
        type = "tar.gz",
        sha256 = "fc4045962a5a5e935ee2fdedaa4e08284547402885ab326734432bed5d12966b",
        strip_prefix = "futures-io-0.3.21",
        build_file = Label("//cargo/remote:BUILD.futures-io-0.3.21.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_macro__0_3_21",
        url = "https://crates.io/api/v1/crates/futures-macro/0.3.21/download",
        type = "tar.gz",
        sha256 = "33c1e13800337f4d4d7a316bf45a567dbcb6ffe087f16424852d97e97a91f512",
        strip_prefix = "futures-macro-0.3.21",
        build_file = Label("//cargo/remote:BUILD.futures-macro-0.3.21.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_sink__0_3_21",
        url = "https://crates.io/api/v1/crates/futures-sink/0.3.21/download",
        type = "tar.gz",
        sha256 = "21163e139fa306126e6eedaf49ecdb4588f939600f0b1e770f4205ee4b7fa868",
        strip_prefix = "futures-sink-0.3.21",
        build_file = Label("//cargo/remote:BUILD.futures-sink-0.3.21.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_task__0_3_21",
        url = "https://crates.io/api/v1/crates/futures-task/0.3.21/download",
        type = "tar.gz",
        sha256 = "57c66a976bf5909d801bbef33416c41372779507e7a6b3a5e25e4749c58f776a",
        strip_prefix = "futures-task-0.3.21",
        build_file = Label("//cargo/remote:BUILD.futures-task-0.3.21.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_util__0_3_21",
        url = "https://crates.io/api/v1/crates/futures-util/0.3.21/download",
        type = "tar.gz",
        sha256 = "d8b7abd5d659d9b90c8cba917f6ec750a74e2dc23902ef9cd4cc8c8b22e6036a",
        strip_prefix = "futures-util-0.3.21",
        build_file = Label("//cargo/remote:BUILD.futures-util-0.3.21.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__generic_array__0_14_5",
        url = "https://crates.io/api/v1/crates/generic-array/0.14.5/download",
        type = "tar.gz",
        sha256 = "fd48d33ec7f05fbfa152300fdad764757cbded343c1aa1cff2fbaf4134851803",
        strip_prefix = "generic-array-0.14.5",
        build_file = Label("//cargo/remote:BUILD.generic-array-0.14.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__getrandom__0_2_7",
        url = "https://crates.io/api/v1/crates/getrandom/0.2.7/download",
        type = "tar.gz",
        sha256 = "4eb1a864a501629691edf6c15a593b7a51eebaa1e8468e9ddc623de7c9b58ec6",
        strip_prefix = "getrandom-0.2.7",
        build_file = Label("//cargo/remote:BUILD.getrandom-0.2.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__h2__0_3_13",
        url = "https://crates.io/api/v1/crates/h2/0.3.13/download",
        type = "tar.gz",
        sha256 = "37a82c6d637fc9515a4694bbf1cb2457b79d81ce52b3108bdeea58b07dd34a57",
        strip_prefix = "h2-0.3.13",
        build_file = Label("//cargo/remote:BUILD.h2-0.3.13.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__hashbrown__0_12_1",
        url = "https://crates.io/api/v1/crates/hashbrown/0.12.1/download",
        type = "tar.gz",
        sha256 = "db0d4cf898abf0081f964436dc980e96670a0f36863e4b83aaacdb65c9d7ccc3",
        strip_prefix = "hashbrown-0.12.1",
        build_file = Label("//cargo/remote:BUILD.hashbrown-0.12.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__hermit_abi__0_1_19",
        url = "https://crates.io/api/v1/crates/hermit-abi/0.1.19/download",
        type = "tar.gz",
        sha256 = "62b467343b94ba476dcb2500d242dadbb39557df889310ac77c5d99100aaac33",
        strip_prefix = "hermit-abi-0.1.19",
        build_file = Label("//cargo/remote:BUILD.hermit-abi-0.1.19.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__http__0_2_8",
        url = "https://crates.io/api/v1/crates/http/0.2.8/download",
        type = "tar.gz",
        sha256 = "75f43d41e26995c17e71ee126451dd3941010b0514a81a9d11f3b341debc2399",
        strip_prefix = "http-0.2.8",
        build_file = Label("//cargo/remote:BUILD.http-0.2.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__http_body__0_4_5",
        url = "https://crates.io/api/v1/crates/http-body/0.4.5/download",
        type = "tar.gz",
        sha256 = "d5f38f16d184e36f2408a55281cd658ecbd3ca05cce6d6510a176eca393e26d1",
        strip_prefix = "http-body-0.4.5",
        build_file = Label("//cargo/remote:BUILD.http-body-0.4.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__http_range_header__0_3_0",
        url = "https://crates.io/api/v1/crates/http-range-header/0.3.0/download",
        type = "tar.gz",
        sha256 = "0bfe8eed0a9285ef776bb792479ea3834e8b94e13d615c2f66d03dd50a435a29",
        strip_prefix = "http-range-header-0.3.0",
        build_file = Label("//cargo/remote:BUILD.http-range-header-0.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__httparse__1_7_1",
        url = "https://crates.io/api/v1/crates/httparse/1.7.1/download",
        type = "tar.gz",
        sha256 = "496ce29bb5a52785b44e0f7ca2847ae0bb839c9bd28f69acac9b99d461c0c04c",
        strip_prefix = "httparse-1.7.1",
        build_file = Label("//cargo/remote:BUILD.httparse-1.7.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__httpdate__1_0_2",
        url = "https://crates.io/api/v1/crates/httpdate/1.0.2/download",
        type = "tar.gz",
        sha256 = "c4a1e36c821dbe04574f602848a19f742f4fb3c98d40449f11bcad18d6b17421",
        strip_prefix = "httpdate-1.0.2",
        build_file = Label("//cargo/remote:BUILD.httpdate-1.0.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__hyper__0_14_19",
        url = "https://crates.io/api/v1/crates/hyper/0.14.19/download",
        type = "tar.gz",
        sha256 = "42dc3c131584288d375f2d07f822b0cb012d8c6fb899a5b9fdb3cb7eb9b6004f",
        strip_prefix = "hyper-0.14.19",
        build_file = Label("//cargo/remote:BUILD.hyper-0.14.19.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__hyper_openssl__0_9_2",
        url = "https://crates.io/api/v1/crates/hyper-openssl/0.9.2/download",
        type = "tar.gz",
        sha256 = "d6ee5d7a8f718585d1c3c61dfde28ef5b0bb14734b4db13f5ada856cdc6c612b",
        strip_prefix = "hyper-openssl-0.9.2",
        build_file = Label("//cargo/remote:BUILD.hyper-openssl-0.9.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__hyper_timeout__0_4_1",
        url = "https://crates.io/api/v1/crates/hyper-timeout/0.4.1/download",
        type = "tar.gz",
        sha256 = "bbb958482e8c7be4bc3cf272a766a2b0bf1a6755e7a6ae777f017a31d11b13b1",
        strip_prefix = "hyper-timeout-0.4.1",
        build_file = Label("//cargo/remote:BUILD.hyper-timeout-0.4.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ident_case__1_0_1",
        url = "https://crates.io/api/v1/crates/ident_case/1.0.1/download",
        type = "tar.gz",
        sha256 = "b9e0384b61958566e926dc50660321d12159025e767c18e043daf26b70104c39",
        strip_prefix = "ident_case-1.0.1",
        build_file = Label("//cargo/remote:BUILD.ident_case-1.0.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__idna__0_2_3",
        url = "https://crates.io/api/v1/crates/idna/0.2.3/download",
        type = "tar.gz",
        sha256 = "418a0a6fab821475f634efe3ccc45c013f742efe03d853e8d3355d5cb850ecf8",
        strip_prefix = "idna-0.2.3",
        build_file = Label("//cargo/remote:BUILD.idna-0.2.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__indexmap__1_9_1",
        url = "https://crates.io/api/v1/crates/indexmap/1.9.1/download",
        type = "tar.gz",
        sha256 = "10a35a97730320ffe8e2d410b5d3b69279b98d2c14bdb8b70ea89ecf7888d41e",
        strip_prefix = "indexmap-1.9.1",
        build_file = Label("//cargo/remote:BUILD.indexmap-1.9.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__instant__0_1_12",
        url = "https://crates.io/api/v1/crates/instant/0.1.12/download",
        type = "tar.gz",
        sha256 = "7a5bbe824c507c5da5956355e86a746d82e0e1464f65d862cc5e71da70e94b2c",
        strip_prefix = "instant-0.1.12",
        build_file = Label("//cargo/remote:BUILD.instant-0.1.12.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__iovec__0_1_4",
        url = "https://crates.io/api/v1/crates/iovec/0.1.4/download",
        type = "tar.gz",
        sha256 = "b2b3ea6ff95e175473f8ffe6a7eb7c00d054240321b84c57051175fe3c1e075e",
        strip_prefix = "iovec-0.1.4",
        build_file = Label("//cargo/remote:BUILD.iovec-0.1.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__itoa__1_0_2",
        url = "https://crates.io/api/v1/crates/itoa/1.0.2/download",
        type = "tar.gz",
        sha256 = "112c678d4050afce233f4f2852bb2eb519230b3cf12f33585275537d7e41578d",
        strip_prefix = "itoa-1.0.2",
        build_file = Label("//cargo/remote:BUILD.itoa-1.0.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__json_patch__0_2_6",
        url = "https://crates.io/api/v1/crates/json-patch/0.2.6/download",
        type = "tar.gz",
        sha256 = "f995a3c8f2bc3dd52a18a583e90f9ec109c047fa1603a853e46bcda14d2e279d",
        strip_prefix = "json-patch-0.2.6",
        build_file = Label("//cargo/remote:BUILD.json-patch-0.2.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__jsonpath_lib__0_3_0",
        url = "https://crates.io/api/v1/crates/jsonpath_lib/0.3.0/download",
        type = "tar.gz",
        sha256 = "eaa63191d68230cccb81c5aa23abd53ed64d83337cacbb25a7b8c7979523774f",
        strip_prefix = "jsonpath_lib-0.3.0",
        build_file = Label("//cargo/remote:BUILD.jsonpath_lib-0.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__k8s_openapi__0_15_0",
        url = "https://crates.io/api/v1/crates/k8s-openapi/0.15.0/download",
        type = "tar.gz",
        sha256 = "d2ae2c04fcee6b01b04e3aadd56bb418932c8e0a9d8a93f48bc68c6bdcdb559d",
        strip_prefix = "k8s-openapi-0.15.0",
        build_file = Label("//cargo/remote:BUILD.k8s-openapi-0.15.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__kernel32_sys__0_2_2",
        url = "https://crates.io/api/v1/crates/kernel32-sys/0.2.2/download",
        type = "tar.gz",
        sha256 = "7507624b29483431c0ba2d82aece8ca6cdba9382bff4ddd0f7490560c056098d",
        strip_prefix = "kernel32-sys-0.2.2",
        build_file = Label("//cargo/remote:BUILD.kernel32-sys-0.2.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__kube__0_74_0",
        url = "https://crates.io/api/v1/crates/kube/0.74.0/download",
        type = "tar.gz",
        sha256 = "a527a8001a61d8d470dab27ac650889938760c243903e7cd90faaf7c60a34bdd",
        strip_prefix = "kube-0.74.0",
        build_file = Label("//cargo/remote:BUILD.kube-0.74.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__kube_client__0_74_0",
        url = "https://crates.io/api/v1/crates/kube-client/0.74.0/download",
        type = "tar.gz",
        sha256 = "c0d48f42df4e8342e9f488c4b97e3759d0042c4e7ab1a853cc285adb44409480",
        strip_prefix = "kube-client-0.74.0",
        build_file = Label("//cargo/remote:BUILD.kube-client-0.74.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__kube_core__0_74_0",
        url = "https://crates.io/api/v1/crates/kube-core/0.74.0/download",
        type = "tar.gz",
        sha256 = "91f56027f862fdcad265d2e9616af416a355e28a1c620bb709083494753e070d",
        strip_prefix = "kube-core-0.74.0",
        build_file = Label("//cargo/remote:BUILD.kube-core-0.74.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__kube_derive__0_74_0",
        url = "https://crates.io/api/v1/crates/kube-derive/0.74.0/download",
        type = "tar.gz",
        sha256 = "66d74121eb41af4480052901f31142d8d9bbdf1b7c6b856da43bcb02f5b1b177",
        strip_prefix = "kube-derive-0.74.0",
        build_file = Label("//cargo/remote:BUILD.kube-derive-0.74.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__kube_runtime__0_74_0",
        url = "https://crates.io/api/v1/crates/kube-runtime/0.74.0/download",
        type = "tar.gz",
        sha256 = "8fdcf5a20f968768e342ef1a457491bb5661fccd81119666d626c57500b16d99",
        strip_prefix = "kube-runtime-0.74.0",
        build_file = Label("//cargo/remote:BUILD.kube-runtime-0.74.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__language_tags__0_3_2",
        url = "https://crates.io/api/v1/crates/language-tags/0.3.2/download",
        type = "tar.gz",
        sha256 = "d4345964bb142484797b161f473a503a434de77149dd8c7427788c6e13379388",
        strip_prefix = "language-tags-0.3.2",
        build_file = Label("//cargo/remote:BUILD.language-tags-0.3.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__lazy_static__1_4_0",
        url = "https://crates.io/api/v1/crates/lazy_static/1.4.0/download",
        type = "tar.gz",
        sha256 = "e2abad23fbc42b3700f2f279844dc832adb2b2eb069b2df918f455c4e18cc646",
        strip_prefix = "lazy_static-1.4.0",
        build_file = Label("//cargo/remote:BUILD.lazy_static-1.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__libc__0_2_126",
        url = "https://crates.io/api/v1/crates/libc/0.2.126/download",
        type = "tar.gz",
        sha256 = "349d5a591cd28b49e1d1037471617a32ddcda5731b99419008085f72d5a53836",
        strip_prefix = "libc-0.2.126",
        build_file = Label("//cargo/remote:BUILD.libc-0.2.126.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__linked_hash_map__0_5_6",
        url = "https://crates.io/api/v1/crates/linked-hash-map/0.5.6/download",
        type = "tar.gz",
        sha256 = "0717cef1bc8b636c6e1c1bbdefc09e6322da8a9321966e8928ef80d20f7f770f",
        strip_prefix = "linked-hash-map-0.5.6",
        build_file = Label("//cargo/remote:BUILD.linked-hash-map-0.5.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__linked_hash_set__0_1_4",
        url = "https://crates.io/api/v1/crates/linked_hash_set/0.1.4/download",
        type = "tar.gz",
        sha256 = "47186c6da4d81ca383c7c47c1bfc80f4b95f4720514d860a5407aaf4233f9588",
        strip_prefix = "linked_hash_set-0.1.4",
        build_file = Label("//cargo/remote:BUILD.linked_hash_set-0.1.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__local_channel__0_1_3",
        url = "https://crates.io/api/v1/crates/local-channel/0.1.3/download",
        type = "tar.gz",
        sha256 = "7f303ec0e94c6c54447f84f3b0ef7af769858a9c4ef56ef2a986d3dcd4c3fc9c",
        strip_prefix = "local-channel-0.1.3",
        build_file = Label("//cargo/remote:BUILD.local-channel-0.1.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__local_waker__0_1_3",
        url = "https://crates.io/api/v1/crates/local-waker/0.1.3/download",
        type = "tar.gz",
        sha256 = "e34f76eb3611940e0e7d53a9aaa4e6a3151f69541a282fd0dad5571420c53ff1",
        strip_prefix = "local-waker-0.1.3",
        build_file = Label("//cargo/remote:BUILD.local-waker-0.1.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__lock_api__0_4_7",
        url = "https://crates.io/api/v1/crates/lock_api/0.4.7/download",
        type = "tar.gz",
        sha256 = "327fa5b6a6940e4699ec49a9beae1ea4845c6bab9314e4f84ac68742139d8c53",
        strip_prefix = "lock_api-0.4.7",
        build_file = Label("//cargo/remote:BUILD.lock_api-0.4.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__log__0_4_17",
        url = "https://crates.io/api/v1/crates/log/0.4.17/download",
        type = "tar.gz",
        sha256 = "abb12e687cfb44aa40f41fc3978ef76448f9b6038cad6aef4259d3c095a2382e",
        strip_prefix = "log-0.4.17",
        build_file = Label("//cargo/remote:BUILD.log-0.4.17.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__matches__0_1_9",
        url = "https://crates.io/api/v1/crates/matches/0.1.9/download",
        type = "tar.gz",
        sha256 = "a3e378b66a060d48947b590737b30a1be76706c8dd7b8ba0f2fe3989c68a853f",
        strip_prefix = "matches-0.1.9",
        build_file = Label("//cargo/remote:BUILD.matches-0.1.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__memchr__2_5_0",
        url = "https://crates.io/api/v1/crates/memchr/2.5.0/download",
        type = "tar.gz",
        sha256 = "2dffe52ecf27772e601905b7522cb4ef790d2cc203488bbd0e2fe85fcb74566d",
        strip_prefix = "memchr-2.5.0",
        build_file = Label("//cargo/remote:BUILD.memchr-2.5.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mime__0_3_16",
        url = "https://crates.io/api/v1/crates/mime/0.3.16/download",
        type = "tar.gz",
        sha256 = "2a60c7ce501c71e03a9c9c0d35b861413ae925bd979cc7a4e30d060069aaac8d",
        strip_prefix = "mime-0.3.16",
        build_file = Label("//cargo/remote:BUILD.mime-0.3.16.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mio__0_6_23",
        url = "https://crates.io/api/v1/crates/mio/0.6.23/download",
        type = "tar.gz",
        sha256 = "4afd66f5b91bf2a3bc13fad0e21caedac168ca4c707504e75585648ae80e4cc4",
        strip_prefix = "mio-0.6.23",
        build_file = Label("//cargo/remote:BUILD.mio-0.6.23.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mio__0_8_4",
        url = "https://crates.io/api/v1/crates/mio/0.8.4/download",
        type = "tar.gz",
        sha256 = "57ee1c23c7c63b0c9250c339ffdc69255f110b298b901b9f6c82547b7b87caaf",
        strip_prefix = "mio-0.8.4",
        build_file = Label("//cargo/remote:BUILD.mio-0.8.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mio_named_pipes__0_1_7",
        url = "https://crates.io/api/v1/crates/mio-named-pipes/0.1.7/download",
        type = "tar.gz",
        sha256 = "0840c1c50fd55e521b247f949c241c9997709f23bd7f023b9762cd561e935656",
        strip_prefix = "mio-named-pipes-0.1.7",
        build_file = Label("//cargo/remote:BUILD.mio-named-pipes-0.1.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mio_uds__0_6_8",
        url = "https://crates.io/api/v1/crates/mio-uds/0.6.8/download",
        type = "tar.gz",
        sha256 = "afcb699eb26d4332647cc848492bbc15eafb26f08d0304550d5aa1f612e066f0",
        strip_prefix = "mio-uds-0.6.8",
        build_file = Label("//cargo/remote:BUILD.mio-uds-0.6.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__miow__0_2_2",
        url = "https://crates.io/api/v1/crates/miow/0.2.2/download",
        type = "tar.gz",
        sha256 = "ebd808424166322d4a38da87083bfddd3ac4c131334ed55856112eb06d46944d",
        strip_prefix = "miow-0.2.2",
        build_file = Label("//cargo/remote:BUILD.miow-0.2.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__miow__0_3_7",
        url = "https://crates.io/api/v1/crates/miow/0.3.7/download",
        type = "tar.gz",
        sha256 = "b9f1c5b025cda876f66ef43a113f91ebc9f4ccef34843000e0adf6ebbab84e21",
        strip_prefix = "miow-0.3.7",
        build_file = Label("//cargo/remote:BUILD.miow-0.3.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__net2__0_2_37",
        url = "https://crates.io/api/v1/crates/net2/0.2.37/download",
        type = "tar.gz",
        sha256 = "391630d12b68002ae1e25e8f974306474966550ad82dac6886fb8910c19568ae",
        strip_prefix = "net2-0.2.37",
        build_file = Label("//cargo/remote:BUILD.net2-0.2.37.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_integer__0_1_45",
        url = "https://crates.io/api/v1/crates/num-integer/0.1.45/download",
        type = "tar.gz",
        sha256 = "225d3389fb3509a24c93f5c29eb6bde2586b98d9f016636dff58d7c6f7569cd9",
        strip_prefix = "num-integer-0.1.45",
        build_file = Label("//cargo/remote:BUILD.num-integer-0.1.45.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_traits__0_2_15",
        url = "https://crates.io/api/v1/crates/num-traits/0.2.15/download",
        type = "tar.gz",
        sha256 = "578ede34cf02f8924ab9447f50c28075b4d3e5b269972345e7e0372b38c6cdcd",
        strip_prefix = "num-traits-0.2.15",
        build_file = Label("//cargo/remote:BUILD.num-traits-0.2.15.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_cpus__1_13_1",
        url = "https://crates.io/api/v1/crates/num_cpus/1.13.1/download",
        type = "tar.gz",
        sha256 = "19e64526ebdee182341572e50e9ad03965aa510cd94427a4549448f285e957a1",
        strip_prefix = "num_cpus-1.13.1",
        build_file = Label("//cargo/remote:BUILD.num_cpus-1.13.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_threads__0_1_6",
        url = "https://crates.io/api/v1/crates/num_threads/0.1.6/download",
        type = "tar.gz",
        sha256 = "2819ce041d2ee131036f4fc9d6ae7ae125a3a40e97ba64d04fe799ad9dabbb44",
        strip_prefix = "num_threads-0.1.6",
        build_file = Label("//cargo/remote:BUILD.num_threads-0.1.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__once_cell__1_13_0",
        url = "https://crates.io/api/v1/crates/once_cell/1.13.0/download",
        type = "tar.gz",
        sha256 = "18a6dbe30758c9f83eb00cbea4ac95966305f5a7772f3f42ebfc7fc7eddbd8e1",
        strip_prefix = "once_cell-1.13.0",
        build_file = Label("//cargo/remote:BUILD.once_cell-1.13.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__openssl__0_10_40",
        url = "https://crates.io/api/v1/crates/openssl/0.10.40/download",
        type = "tar.gz",
        sha256 = "fb81a6430ac911acb25fe5ac8f1d2af1b4ea8a4fdfda0f1ee4292af2e2d8eb0e",
        strip_prefix = "openssl-0.10.40",
        build_file = Label("//cargo/remote:BUILD.openssl-0.10.40.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__openssl_macros__0_1_0",
        url = "https://crates.io/api/v1/crates/openssl-macros/0.1.0/download",
        type = "tar.gz",
        sha256 = "b501e44f11665960c7e7fcf062c7d96a14ade4aa98116c004b2e37b5be7d736c",
        strip_prefix = "openssl-macros-0.1.0",
        build_file = Label("//cargo/remote:BUILD.openssl-macros-0.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__openssl_sys__0_9_74",
        url = "https://crates.io/api/v1/crates/openssl-sys/0.9.74/download",
        type = "tar.gz",
        sha256 = "835363342df5fba8354c5b453325b110ffd54044e588c539cf2f20a8014e4cb1",
        strip_prefix = "openssl-sys-0.9.74",
        build_file = Label("//cargo/remote:BUILD.openssl-sys-0.9.74.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ordered_float__2_10_0",
        url = "https://crates.io/api/v1/crates/ordered-float/2.10.0/download",
        type = "tar.gz",
        sha256 = "7940cf2ca942593318d07fcf2596cdca60a85c9e7fab408a5e21a4f9dcd40d87",
        strip_prefix = "ordered-float-2.10.0",
        build_file = Label("//cargo/remote:BUILD.ordered-float-2.10.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__parking_lot__0_12_1",
        url = "https://crates.io/api/v1/crates/parking_lot/0.12.1/download",
        type = "tar.gz",
        sha256 = "3742b2c103b9f06bc9fff0a37ff4912935851bee6d36f3c02bcc755bcfec228f",
        strip_prefix = "parking_lot-0.12.1",
        build_file = Label("//cargo/remote:BUILD.parking_lot-0.12.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__parking_lot_core__0_9_3",
        url = "https://crates.io/api/v1/crates/parking_lot_core/0.9.3/download",
        type = "tar.gz",
        sha256 = "09a279cbf25cb0757810394fbc1e359949b59e348145c643a939a525692e6929",
        strip_prefix = "parking_lot_core-0.9.3",
        build_file = Label("//cargo/remote:BUILD.parking_lot_core-0.9.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__paste__1_0_7",
        url = "https://crates.io/api/v1/crates/paste/1.0.7/download",
        type = "tar.gz",
        sha256 = "0c520e05135d6e763148b6426a837e239041653ba7becd2e538c076c738025fc",
        strip_prefix = "paste-1.0.7",
        build_file = Label("//cargo/remote:BUILD.paste-1.0.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pem__1_0_2",
        url = "https://crates.io/api/v1/crates/pem/1.0.2/download",
        type = "tar.gz",
        sha256 = "e9a3b09a20e374558580a4914d3b7d89bd61b954a5a5e1dcbea98753addb1947",
        strip_prefix = "pem-1.0.2",
        build_file = Label("//cargo/remote:BUILD.pem-1.0.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__percent_encoding__2_1_0",
        url = "https://crates.io/api/v1/crates/percent-encoding/2.1.0/download",
        type = "tar.gz",
        sha256 = "d4fd5641d01c8f18a23da7b6fe29298ff4b55afcccdf78973b24cf3175fee32e",
        strip_prefix = "percent-encoding-2.1.0",
        build_file = Label("//cargo/remote:BUILD.percent-encoding-2.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pin_project__1_0_11",
        url = "https://crates.io/api/v1/crates/pin-project/1.0.11/download",
        type = "tar.gz",
        sha256 = "78203e83c48cffbe01e4a2d35d566ca4de445d79a85372fc64e378bfc812a260",
        strip_prefix = "pin-project-1.0.11",
        build_file = Label("//cargo/remote:BUILD.pin-project-1.0.11.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pin_project_internal__1_0_11",
        url = "https://crates.io/api/v1/crates/pin-project-internal/1.0.11/download",
        type = "tar.gz",
        sha256 = "710faf75e1b33345361201d36d04e98ac1ed8909151a017ed384700836104c74",
        strip_prefix = "pin-project-internal-1.0.11",
        build_file = Label("//cargo/remote:BUILD.pin-project-internal-1.0.11.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pin_project_lite__0_1_12",
        url = "https://crates.io/api/v1/crates/pin-project-lite/0.1.12/download",
        type = "tar.gz",
        sha256 = "257b64915a082f7811703966789728173279bdebb956b143dbcd23f6f970a777",
        strip_prefix = "pin-project-lite-0.1.12",
        build_file = Label("//cargo/remote:BUILD.pin-project-lite-0.1.12.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pin_project_lite__0_2_9",
        url = "https://crates.io/api/v1/crates/pin-project-lite/0.2.9/download",
        type = "tar.gz",
        sha256 = "e0a7ae3ac2f1173085d398531c705756c94a4c56843785df85a60c1a0afac116",
        strip_prefix = "pin-project-lite-0.2.9",
        build_file = Label("//cargo/remote:BUILD.pin-project-lite-0.2.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pin_utils__0_1_0",
        url = "https://crates.io/api/v1/crates/pin-utils/0.1.0/download",
        type = "tar.gz",
        sha256 = "8b870d8c151b6f2fb93e84a13146138f05d02ed11c7e7c54f8826aaaf7c9f184",
        strip_prefix = "pin-utils-0.1.0",
        build_file = Label("//cargo/remote:BUILD.pin-utils-0.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pkg_config__0_3_25",
        url = "https://crates.io/api/v1/crates/pkg-config/0.3.25/download",
        type = "tar.gz",
        sha256 = "1df8c4ec4b0627e53bdf214615ad287367e482558cf84b109250b37464dc03ae",
        strip_prefix = "pkg-config-0.3.25",
        build_file = Label("//cargo/remote:BUILD.pkg-config-0.3.25.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ppv_lite86__0_2_16",
        url = "https://crates.io/api/v1/crates/ppv-lite86/0.2.16/download",
        type = "tar.gz",
        sha256 = "eb9f9e6e233e5c4a35559a617bf40a4ec447db2e84c20b55a6f83167b7e57872",
        strip_prefix = "ppv-lite86-0.2.16",
        build_file = Label("//cargo/remote:BUILD.ppv-lite86-0.2.16.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__proc_macro2__1_0_40",
        url = "https://crates.io/api/v1/crates/proc-macro2/1.0.40/download",
        type = "tar.gz",
        sha256 = "dd96a1e8ed2596c337f8eae5f24924ec83f5ad5ab21ea8e455d3566c69fbcaf7",
        strip_prefix = "proc-macro2-1.0.40",
        build_file = Label("//cargo/remote:BUILD.proc-macro2-1.0.40.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__quote__1_0_20",
        url = "https://crates.io/api/v1/crates/quote/1.0.20/download",
        type = "tar.gz",
        sha256 = "3bcdf212e9776fbcb2d23ab029360416bb1706b1aea2d1a5ba002727cbcab804",
        strip_prefix = "quote-1.0.20",
        build_file = Label("//cargo/remote:BUILD.quote-1.0.20.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rand__0_8_5",
        url = "https://crates.io/api/v1/crates/rand/0.8.5/download",
        type = "tar.gz",
        sha256 = "34af8d1a0e25924bc5b7c43c079c942339d8f0a8b57c39049bef581b46327404",
        strip_prefix = "rand-0.8.5",
        build_file = Label("//cargo/remote:BUILD.rand-0.8.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rand_chacha__0_3_1",
        url = "https://crates.io/api/v1/crates/rand_chacha/0.3.1/download",
        type = "tar.gz",
        sha256 = "e6c10a63a0fa32252be49d21e7709d4d4baf8d231c2dbce1eaa8141b9b127d88",
        strip_prefix = "rand_chacha-0.3.1",
        build_file = Label("//cargo/remote:BUILD.rand_chacha-0.3.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rand_core__0_6_3",
        url = "https://crates.io/api/v1/crates/rand_core/0.6.3/download",
        type = "tar.gz",
        sha256 = "d34f1408f55294453790c48b2f1ebbb1c5b4b7563eb1f418bcfcfdbb06ebb4e7",
        strip_prefix = "rand_core-0.6.3",
        build_file = Label("//cargo/remote:BUILD.rand_core-0.6.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__redox_syscall__0_2_13",
        url = "https://crates.io/api/v1/crates/redox_syscall/0.2.13/download",
        type = "tar.gz",
        sha256 = "62f25bc4c7e55e0b0b7a1d43fb893f4fa1361d0abe38b9ce4f323c2adfe6ef42",
        strip_prefix = "redox_syscall-0.2.13",
        build_file = Label("//cargo/remote:BUILD.redox_syscall-0.2.13.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__redox_users__0_4_3",
        url = "https://crates.io/api/v1/crates/redox_users/0.4.3/download",
        type = "tar.gz",
        sha256 = "b033d837a7cf162d7993aded9304e30a83213c648b6e389db233191f891e5c2b",
        strip_prefix = "redox_users-0.4.3",
        build_file = Label("//cargo/remote:BUILD.redox_users-0.4.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__regex__1_6_0",
        url = "https://crates.io/api/v1/crates/regex/1.6.0/download",
        type = "tar.gz",
        sha256 = "4c4eb3267174b8c6c2f654116623910a0fef09c4753f8dd83db29c48a0df988b",
        strip_prefix = "regex-1.6.0",
        build_file = Label("//cargo/remote:BUILD.regex-1.6.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__regex_syntax__0_6_27",
        url = "https://crates.io/api/v1/crates/regex-syntax/0.6.27/download",
        type = "tar.gz",
        sha256 = "a3f87b73ce11b1619a3c6332f45341e0047173771e8b8b73f87bfeefb7b56244",
        strip_prefix = "regex-syntax-0.6.27",
        build_file = Label("//cargo/remote:BUILD.regex-syntax-0.6.27.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rustc_version__0_4_0",
        url = "https://crates.io/api/v1/crates/rustc_version/0.4.0/download",
        type = "tar.gz",
        sha256 = "bfa0f585226d2e68097d4f95d113b15b83a82e819ab25717ec0590d9584ef366",
        strip_prefix = "rustc_version-0.4.0",
        build_file = Label("//cargo/remote:BUILD.rustc_version-0.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ryu__1_0_10",
        url = "https://crates.io/api/v1/crates/ryu/1.0.10/download",
        type = "tar.gz",
        sha256 = "f3f6f92acf49d1b98f7a81226834412ada05458b7364277387724a237f062695",
        strip_prefix = "ryu-1.0.10",
        build_file = Label("//cargo/remote:BUILD.ryu-1.0.10.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__schemars__0_8_10",
        url = "https://crates.io/api/v1/crates/schemars/0.8.10/download",
        type = "tar.gz",
        sha256 = "1847b767a3d62d95cbf3d8a9f0e421cf57a0d8aa4f411d4b16525afb0284d4ed",
        strip_prefix = "schemars-0.8.10",
        build_file = Label("//cargo/remote:BUILD.schemars-0.8.10.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__schemars_derive__0_8_10",
        url = "https://crates.io/api/v1/crates/schemars_derive/0.8.10/download",
        type = "tar.gz",
        sha256 = "af4d7e1b012cb3d9129567661a63755ea4b8a7386d339dc945ae187e403c6743",
        strip_prefix = "schemars_derive-0.8.10",
        build_file = Label("//cargo/remote:BUILD.schemars_derive-0.8.10.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__scopeguard__1_1_0",
        url = "https://crates.io/api/v1/crates/scopeguard/1.1.0/download",
        type = "tar.gz",
        sha256 = "d29ab0c6d3fc0ee92fe66e2d99f700eab17a8d57d1c1d3b748380fb20baa78cd",
        strip_prefix = "scopeguard-1.1.0",
        build_file = Label("//cargo/remote:BUILD.scopeguard-1.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__secrecy__0_8_0",
        url = "https://crates.io/api/v1/crates/secrecy/0.8.0/download",
        type = "tar.gz",
        sha256 = "9bd1c54ea06cfd2f6b63219704de0b9b4f72dcc2b8fdef820be6cd799780e91e",
        strip_prefix = "secrecy-0.8.0",
        build_file = Label("//cargo/remote:BUILD.secrecy-0.8.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__semver__1_0_12",
        url = "https://crates.io/api/v1/crates/semver/1.0.12/download",
        type = "tar.gz",
        sha256 = "a2333e6df6d6598f2b1974829f853c2b4c5f4a6e503c10af918081aa6f8564e1",
        strip_prefix = "semver-1.0.12",
        build_file = Label("//cargo/remote:BUILD.semver-1.0.12.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde__1_0_138",
        url = "https://crates.io/api/v1/crates/serde/1.0.138/download",
        type = "tar.gz",
        sha256 = "1578c6245786b9d168c5447eeacfb96856573ca56c9d68fdcf394be134882a47",
        strip_prefix = "serde-1.0.138",
        build_file = Label("//cargo/remote:BUILD.serde-1.0.138.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde_value__0_7_0",
        url = "https://crates.io/api/v1/crates/serde-value/0.7.0/download",
        type = "tar.gz",
        sha256 = "f3a1a3341211875ef120e117ea7fd5228530ae7e7036a779fdc9117be6b3282c",
        strip_prefix = "serde-value-0.7.0",
        build_file = Label("//cargo/remote:BUILD.serde-value-0.7.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde_derive__1_0_138",
        url = "https://crates.io/api/v1/crates/serde_derive/1.0.138/download",
        type = "tar.gz",
        sha256 = "023e9b1467aef8a10fb88f25611870ada9800ef7e22afce356bb0d2387b6f27c",
        strip_prefix = "serde_derive-1.0.138",
        build_file = Label("//cargo/remote:BUILD.serde_derive-1.0.138.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde_derive_internals__0_26_0",
        url = "https://crates.io/api/v1/crates/serde_derive_internals/0.26.0/download",
        type = "tar.gz",
        sha256 = "85bf8229e7920a9f636479437026331ce11aa132b4dde37d121944a44d6e5f3c",
        strip_prefix = "serde_derive_internals-0.26.0",
        build_file = Label("//cargo/remote:BUILD.serde_derive_internals-0.26.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde_json__1_0_82",
        url = "https://crates.io/api/v1/crates/serde_json/1.0.82/download",
        type = "tar.gz",
        sha256 = "82c2c1fdcd807d1098552c5b9a36e425e42e9fbd7c6a37a8425f390f781f7fa7",
        strip_prefix = "serde_json-1.0.82",
        build_file = Label("//cargo/remote:BUILD.serde_json-1.0.82.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde_urlencoded__0_7_1",
        url = "https://crates.io/api/v1/crates/serde_urlencoded/0.7.1/download",
        type = "tar.gz",
        sha256 = "d3491c14715ca2294c4d6a88f15e84739788c1d030eed8c110436aafdaa2f3fd",
        strip_prefix = "serde_urlencoded-0.7.1",
        build_file = Label("//cargo/remote:BUILD.serde_urlencoded-0.7.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde_yaml__0_8_24",
        url = "https://crates.io/api/v1/crates/serde_yaml/0.8.24/download",
        type = "tar.gz",
        sha256 = "707d15895415db6628332b737c838b88c598522e4dc70647e59b72312924aebc",
        strip_prefix = "serde_yaml-0.8.24",
        build_file = Label("//cargo/remote:BUILD.serde_yaml-0.8.24.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__sha1__0_10_1",
        url = "https://crates.io/api/v1/crates/sha1/0.10.1/download",
        type = "tar.gz",
        sha256 = "c77f4e7f65455545c2153c1253d25056825e77ee2533f0e41deb65a93a34852f",
        strip_prefix = "sha1-0.10.1",
        build_file = Label("//cargo/remote:BUILD.sha1-0.10.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__signal_hook_registry__1_4_0",
        url = "https://crates.io/api/v1/crates/signal-hook-registry/1.4.0/download",
        type = "tar.gz",
        sha256 = "e51e73328dc4ac0c7ccbda3a494dfa03df1de2f46018127f60c693f2648455b0",
        strip_prefix = "signal-hook-registry-1.4.0",
        build_file = Label("//cargo/remote:BUILD.signal-hook-registry-1.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__simple_logger__2_2_0",
        url = "https://crates.io/api/v1/crates/simple_logger/2.2.0/download",
        type = "tar.gz",
        sha256 = "166fea527c36d9b8a0a88c0c6d4c5077c699d9ffb5cf890b231a3f08c35f3d40",
        strip_prefix = "simple_logger-2.2.0",
        build_file = Label("//cargo/remote:BUILD.simple_logger-2.2.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__slab__0_4_6",
        url = "https://crates.io/api/v1/crates/slab/0.4.6/download",
        type = "tar.gz",
        sha256 = "eb703cfe953bccee95685111adeedb76fabe4e97549a58d16f03ea7b9367bb32",
        strip_prefix = "slab-0.4.6",
        build_file = Label("//cargo/remote:BUILD.slab-0.4.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__smallvec__1_9_0",
        url = "https://crates.io/api/v1/crates/smallvec/1.9.0/download",
        type = "tar.gz",
        sha256 = "2fd0db749597d91ff862fd1d55ea87f7855a744a8425a64695b6fca237d1dad1",
        strip_prefix = "smallvec-1.9.0",
        build_file = Label("//cargo/remote:BUILD.smallvec-1.9.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__socket2__0_4_4",
        url = "https://crates.io/api/v1/crates/socket2/0.4.4/download",
        type = "tar.gz",
        sha256 = "66d72b759436ae32898a2af0a14218dbf55efde3feeb170eb623637db85ee1e0",
        strip_prefix = "socket2-0.4.4",
        build_file = Label("//cargo/remote:BUILD.socket2-0.4.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__strsim__0_10_0",
        url = "https://crates.io/api/v1/crates/strsim/0.10.0/download",
        type = "tar.gz",
        sha256 = "73473c0e59e6d5812c5dfe2a064a6444949f089e20eec9a2e5506596494e4623",
        strip_prefix = "strsim-0.10.0",
        build_file = Label("//cargo/remote:BUILD.strsim-0.10.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__syn__1_0_98",
        url = "https://crates.io/api/v1/crates/syn/1.0.98/download",
        type = "tar.gz",
        sha256 = "c50aef8a904de4c23c788f104b7dddc7d6f79c647c7c8ce4cc8f73eb0ca773dd",
        strip_prefix = "syn-1.0.98",
        build_file = Label("//cargo/remote:BUILD.syn-1.0.98.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__thiserror__1_0_31",
        url = "https://crates.io/api/v1/crates/thiserror/1.0.31/download",
        type = "tar.gz",
        sha256 = "bd829fe32373d27f76265620b5309d0340cb8550f523c1dda251d6298069069a",
        strip_prefix = "thiserror-1.0.31",
        build_file = Label("//cargo/remote:BUILD.thiserror-1.0.31.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__thiserror_impl__1_0_31",
        url = "https://crates.io/api/v1/crates/thiserror-impl/1.0.31/download",
        type = "tar.gz",
        sha256 = "0396bc89e626244658bef819e22d0cc459e795a5ebe878e6ec336d1674a8d79a",
        strip_prefix = "thiserror-impl-1.0.31",
        build_file = Label("//cargo/remote:BUILD.thiserror-impl-1.0.31.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__time__0_3_11",
        url = "https://crates.io/api/v1/crates/time/0.3.11/download",
        type = "tar.gz",
        sha256 = "72c91f41dcb2f096c05f0873d667dceec1087ce5bcf984ec8ffb19acddbb3217",
        strip_prefix = "time-0.3.11",
        build_file = Label("//cargo/remote:BUILD.time-0.3.11.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__time_macros__0_2_4",
        url = "https://crates.io/api/v1/crates/time-macros/0.2.4/download",
        type = "tar.gz",
        sha256 = "42657b1a6f4d817cda8e7a0ace261fe0cc946cf3a80314390b22cc61ae080792",
        strip_prefix = "time-macros-0.2.4",
        build_file = Label("//cargo/remote:BUILD.time-macros-0.2.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tinyvec__1_6_0",
        url = "https://crates.io/api/v1/crates/tinyvec/1.6.0/download",
        type = "tar.gz",
        sha256 = "87cc5ceb3875bb20c2890005a4e226a4651264a5c75edb2421b52861a0a0cb50",
        strip_prefix = "tinyvec-1.6.0",
        build_file = Label("//cargo/remote:BUILD.tinyvec-1.6.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tinyvec_macros__0_1_0",
        url = "https://crates.io/api/v1/crates/tinyvec_macros/0.1.0/download",
        type = "tar.gz",
        sha256 = "cda74da7e1a664f795bb1f8a87ec406fb89a02522cf6e50620d016add6dbbf5c",
        strip_prefix = "tinyvec_macros-0.1.0",
        build_file = Label("//cargo/remote:BUILD.tinyvec_macros-0.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio__0_2_25",
        url = "https://crates.io/api/v1/crates/tokio/0.2.25/download",
        type = "tar.gz",
        sha256 = "6703a273949a90131b290be1fe7b039d0fc884aa1935860dfcbe056f28cd8092",
        strip_prefix = "tokio-0.2.25",
        build_file = Label("//cargo/remote:BUILD.tokio-0.2.25.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio__1_19_2",
        url = "https://crates.io/api/v1/crates/tokio/1.19.2/download",
        type = "tar.gz",
        sha256 = "c51a52ed6686dd62c320f9b89299e9dfb46f730c7a48e635c19f21d116cb1439",
        strip_prefix = "tokio-1.19.2",
        build_file = Label("//cargo/remote:BUILD.tokio-1.19.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio_io_timeout__1_2_0",
        url = "https://crates.io/api/v1/crates/tokio-io-timeout/1.2.0/download",
        type = "tar.gz",
        sha256 = "30b74022ada614a1b4834de765f9bb43877f910cc8ce4be40e89042c9223a8bf",
        strip_prefix = "tokio-io-timeout-1.2.0",
        build_file = Label("//cargo/remote:BUILD.tokio-io-timeout-1.2.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio_macros__0_2_6",
        url = "https://crates.io/api/v1/crates/tokio-macros/0.2.6/download",
        type = "tar.gz",
        sha256 = "e44da00bfc73a25f814cd8d7e57a68a5c31b74b3152a0a1d1f590c97ed06265a",
        strip_prefix = "tokio-macros-0.2.6",
        build_file = Label("//cargo/remote:BUILD.tokio-macros-0.2.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio_openssl__0_6_3",
        url = "https://crates.io/api/v1/crates/tokio-openssl/0.6.3/download",
        type = "tar.gz",
        sha256 = "c08f9ffb7809f1b20c1b398d92acf4cc719874b3b2b2d9ea2f09b4a80350878a",
        strip_prefix = "tokio-openssl-0.6.3",
        build_file = Label("//cargo/remote:BUILD.tokio-openssl-0.6.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio_util__0_7_3",
        url = "https://crates.io/api/v1/crates/tokio-util/0.7.3/download",
        type = "tar.gz",
        sha256 = "cc463cd8deddc3770d20f9852143d50bf6094e640b485cb2e189a2099085ff45",
        strip_prefix = "tokio-util-0.7.3",
        build_file = Label("//cargo/remote:BUILD.tokio-util-0.7.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tower__0_4_13",
        url = "https://crates.io/api/v1/crates/tower/0.4.13/download",
        type = "tar.gz",
        sha256 = "b8fa9be0de6cf49e536ce1851f987bd21a43b771b09473c3549a6c853db37c1c",
        strip_prefix = "tower-0.4.13",
        build_file = Label("//cargo/remote:BUILD.tower-0.4.13.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tower_http__0_3_4",
        url = "https://crates.io/api/v1/crates/tower-http/0.3.4/download",
        type = "tar.gz",
        sha256 = "3c530c8675c1dbf98facee631536fa116b5fb6382d7dd6dc1b118d970eafe3ba",
        strip_prefix = "tower-http-0.3.4",
        build_file = Label("//cargo/remote:BUILD.tower-http-0.3.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tower_layer__0_3_1",
        url = "https://crates.io/api/v1/crates/tower-layer/0.3.1/download",
        type = "tar.gz",
        sha256 = "343bc9466d3fe6b0f960ef45960509f84480bf4fd96f92901afe7ff3df9d3a62",
        strip_prefix = "tower-layer-0.3.1",
        build_file = Label("//cargo/remote:BUILD.tower-layer-0.3.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tower_service__0_3_2",
        url = "https://crates.io/api/v1/crates/tower-service/0.3.2/download",
        type = "tar.gz",
        sha256 = "b6bc1c9ce2b5135ac7f93c72918fc37feb872bdc6a5533a8b85eb4b86bfdae52",
        strip_prefix = "tower-service-0.3.2",
        build_file = Label("//cargo/remote:BUILD.tower-service-0.3.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tracing__0_1_35",
        url = "https://crates.io/api/v1/crates/tracing/0.1.35/download",
        type = "tar.gz",
        sha256 = "a400e31aa60b9d44a52a8ee0343b5b18566b03a8321e0d321f695cf56e940160",
        strip_prefix = "tracing-0.1.35",
        build_file = Label("//cargo/remote:BUILD.tracing-0.1.35.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tracing_attributes__0_1_22",
        url = "https://crates.io/api/v1/crates/tracing-attributes/0.1.22/download",
        type = "tar.gz",
        sha256 = "11c75893af559bc8e10716548bdef5cb2b983f8e637db9d0e15126b61b484ee2",
        strip_prefix = "tracing-attributes-0.1.22",
        build_file = Label("//cargo/remote:BUILD.tracing-attributes-0.1.22.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tracing_core__0_1_28",
        url = "https://crates.io/api/v1/crates/tracing-core/0.1.28/download",
        type = "tar.gz",
        sha256 = "7b7358be39f2f274f322d2aaed611acc57f382e8eb1e5b48cb9ae30933495ce7",
        strip_prefix = "tracing-core-0.1.28",
        build_file = Label("//cargo/remote:BUILD.tracing-core-0.1.28.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__treediff__3_0_2",
        url = "https://crates.io/api/v1/crates/treediff/3.0.2/download",
        type = "tar.gz",
        sha256 = "761e8d5ad7ce14bb82b7e61ccc0ca961005a275a060b9644a2431aa11553c2ff",
        strip_prefix = "treediff-3.0.2",
        build_file = Label("//cargo/remote:BUILD.treediff-3.0.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__try_lock__0_2_3",
        url = "https://crates.io/api/v1/crates/try-lock/0.2.3/download",
        type = "tar.gz",
        sha256 = "59547bce71d9c38b83d9c0e92b6066c4253371f15005def0c30d9657f50c7642",
        strip_prefix = "try-lock-0.2.3",
        build_file = Label("//cargo/remote:BUILD.try-lock-0.2.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__typenum__1_15_0",
        url = "https://crates.io/api/v1/crates/typenum/1.15.0/download",
        type = "tar.gz",
        sha256 = "dcf81ac59edc17cc8697ff311e8f5ef2d99fcbd9817b34cec66f90b6c3dfd987",
        strip_prefix = "typenum-1.15.0",
        build_file = Label("//cargo/remote:BUILD.typenum-1.15.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__unicode_bidi__0_3_8",
        url = "https://crates.io/api/v1/crates/unicode-bidi/0.3.8/download",
        type = "tar.gz",
        sha256 = "099b7128301d285f79ddd55b9a83d5e6b9e97c92e0ea0daebee7263e932de992",
        strip_prefix = "unicode-bidi-0.3.8",
        build_file = Label("//cargo/remote:BUILD.unicode-bidi-0.3.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__unicode_ident__1_0_1",
        url = "https://crates.io/api/v1/crates/unicode-ident/1.0.1/download",
        type = "tar.gz",
        sha256 = "5bd2fe26506023ed7b5e1e315add59d6f584c621d037f9368fea9cfb988f368c",
        strip_prefix = "unicode-ident-1.0.1",
        build_file = Label("//cargo/remote:BUILD.unicode-ident-1.0.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__unicode_normalization__0_1_21",
        url = "https://crates.io/api/v1/crates/unicode-normalization/0.1.21/download",
        type = "tar.gz",
        sha256 = "854cbdc4f7bc6ae19c820d44abdc3277ac3e1b2b93db20a636825d9322fb60e6",
        strip_prefix = "unicode-normalization-0.1.21",
        build_file = Label("//cargo/remote:BUILD.unicode-normalization-0.1.21.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__url__2_2_2",
        url = "https://crates.io/api/v1/crates/url/2.2.2/download",
        type = "tar.gz",
        sha256 = "a507c383b2d33b5fc35d1861e77e6b383d158b2da5e14fe51b83dfedf6fd578c",
        strip_prefix = "url-2.2.2",
        build_file = Label("//cargo/remote:BUILD.url-2.2.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__vcpkg__0_2_15",
        url = "https://crates.io/api/v1/crates/vcpkg/0.2.15/download",
        type = "tar.gz",
        sha256 = "accd4ea62f7bb7a82fe23066fb0957d48ef677f6eeb8215f372f52e48bb32426",
        strip_prefix = "vcpkg-0.2.15",
        build_file = Label("//cargo/remote:BUILD.vcpkg-0.2.15.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__version_check__0_9_4",
        url = "https://crates.io/api/v1/crates/version_check/0.9.4/download",
        type = "tar.gz",
        sha256 = "49874b5167b65d7193b8aba1567f5c7d93d001cafc34600cee003eda787e483f",
        strip_prefix = "version_check-0.9.4",
        build_file = Label("//cargo/remote:BUILD.version_check-0.9.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__want__0_3_0",
        url = "https://crates.io/api/v1/crates/want/0.3.0/download",
        type = "tar.gz",
        sha256 = "1ce8a968cb1cd110d136ff8b819a556d6fb6d919363c61534f6860c7eb172ba0",
        strip_prefix = "want-0.3.0",
        build_file = Label("//cargo/remote:BUILD.want-0.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__wasi__0_11_0_wasi_snapshot_preview1",
        url = "https://crates.io/api/v1/crates/wasi/0.11.0+wasi-snapshot-preview1/download",
        type = "tar.gz",
        sha256 = "9c8d87e72b64a3b4db28d11ce29237c246188f4f51057d65a7eab63b7987e423",
        strip_prefix = "wasi-0.11.0+wasi-snapshot-preview1",
        build_file = Label("//cargo/remote:BUILD.wasi-0.11.0+wasi-snapshot-preview1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi__0_2_8",
        url = "https://crates.io/api/v1/crates/winapi/0.2.8/download",
        type = "tar.gz",
        sha256 = "167dc9d6949a9b857f3451275e911c3f44255842c1f7a76f33c55103a909087a",
        strip_prefix = "winapi-0.2.8",
        build_file = Label("//cargo/remote:BUILD.winapi-0.2.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi__0_3_9",
        url = "https://crates.io/api/v1/crates/winapi/0.3.9/download",
        type = "tar.gz",
        sha256 = "5c839a674fcd7a98952e593242ea400abe93992746761e38641405d28b00f419",
        strip_prefix = "winapi-0.3.9",
        build_file = Label("//cargo/remote:BUILD.winapi-0.3.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi_build__0_1_1",
        url = "https://crates.io/api/v1/crates/winapi-build/0.1.1/download",
        type = "tar.gz",
        sha256 = "2d315eee3b34aca4797b2da6b13ed88266e6d612562a0c46390af8299fc699bc",
        strip_prefix = "winapi-build-0.1.1",
        build_file = Label("//cargo/remote:BUILD.winapi-build-0.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi_i686_pc_windows_gnu__0_4_0",
        url = "https://crates.io/api/v1/crates/winapi-i686-pc-windows-gnu/0.4.0/download",
        type = "tar.gz",
        sha256 = "ac3b87c63620426dd9b991e5ce0329eff545bccbbb34f3be09ff6fb6ab51b7b6",
        strip_prefix = "winapi-i686-pc-windows-gnu-0.4.0",
        build_file = Label("//cargo/remote:BUILD.winapi-i686-pc-windows-gnu-0.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi_x86_64_pc_windows_gnu__0_4_0",
        url = "https://crates.io/api/v1/crates/winapi-x86_64-pc-windows-gnu/0.4.0/download",
        type = "tar.gz",
        sha256 = "712e227841d057c1ee1cd2fb22fa7e5a5461ae8e48fa2ca79ec42cfc1931183f",
        strip_prefix = "winapi-x86_64-pc-windows-gnu-0.4.0",
        build_file = Label("//cargo/remote:BUILD.winapi-x86_64-pc-windows-gnu-0.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_sys__0_36_1",
        url = "https://crates.io/api/v1/crates/windows-sys/0.36.1/download",
        type = "tar.gz",
        sha256 = "ea04155a16a59f9eab786fe12a4a450e75cdb175f9e0d80da1e17db09f55b8d2",
        strip_prefix = "windows-sys-0.36.1",
        build_file = Label("//cargo/remote:BUILD.windows-sys-0.36.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_aarch64_msvc__0_36_1",
        url = "https://crates.io/api/v1/crates/windows_aarch64_msvc/0.36.1/download",
        type = "tar.gz",
        sha256 = "9bb8c3fd39ade2d67e9874ac4f3db21f0d710bee00fe7cab16949ec184eeaa47",
        strip_prefix = "windows_aarch64_msvc-0.36.1",
        build_file = Label("//cargo/remote:BUILD.windows_aarch64_msvc-0.36.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_i686_gnu__0_36_1",
        url = "https://crates.io/api/v1/crates/windows_i686_gnu/0.36.1/download",
        type = "tar.gz",
        sha256 = "180e6ccf01daf4c426b846dfc66db1fc518f074baa793aa7d9b9aaeffad6a3b6",
        strip_prefix = "windows_i686_gnu-0.36.1",
        build_file = Label("//cargo/remote:BUILD.windows_i686_gnu-0.36.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_i686_msvc__0_36_1",
        url = "https://crates.io/api/v1/crates/windows_i686_msvc/0.36.1/download",
        type = "tar.gz",
        sha256 = "e2e7917148b2812d1eeafaeb22a97e4813dfa60a3f8f78ebe204bcc88f12f024",
        strip_prefix = "windows_i686_msvc-0.36.1",
        build_file = Label("//cargo/remote:BUILD.windows_i686_msvc-0.36.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_x86_64_gnu__0_36_1",
        url = "https://crates.io/api/v1/crates/windows_x86_64_gnu/0.36.1/download",
        type = "tar.gz",
        sha256 = "4dcd171b8776c41b97521e5da127a2d86ad280114807d0b2ab1e462bc764d9e1",
        strip_prefix = "windows_x86_64_gnu-0.36.1",
        build_file = Label("//cargo/remote:BUILD.windows_x86_64_gnu-0.36.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_x86_64_msvc__0_36_1",
        url = "https://crates.io/api/v1/crates/windows_x86_64_msvc/0.36.1/download",
        type = "tar.gz",
        sha256 = "c811ca4a8c853ef420abd8592ba53ddbbac90410fab6903b3e79972a631f7680",
        strip_prefix = "windows_x86_64_msvc-0.36.1",
        build_file = Label("//cargo/remote:BUILD.windows_x86_64_msvc-0.36.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ws2_32_sys__0_2_1",
        url = "https://crates.io/api/v1/crates/ws2_32-sys/0.2.1/download",
        type = "tar.gz",
        sha256 = "d59cefebd0c892fa2dd6de581e937301d8552cb44489cdff035c6187cb63fa5e",
        strip_prefix = "ws2_32-sys-0.2.1",
        build_file = Label("//cargo/remote:BUILD.ws2_32-sys-0.2.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__yaml_rust__0_4_5",
        url = "https://crates.io/api/v1/crates/yaml-rust/0.4.5/download",
        type = "tar.gz",
        sha256 = "56c1936c4cc7a1c9ab21a1ebb602eb942ba868cbd44a99cb7cdc5892335e1c85",
        strip_prefix = "yaml-rust-0.4.5",
        build_file = Label("//cargo/remote:BUILD.yaml-rust-0.4.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__zeroize__1_5_6",
        url = "https://crates.io/api/v1/crates/zeroize/1.5.6/download",
        type = "tar.gz",
        sha256 = "20b578acffd8516a6c3f2a1bdefc1ec37e547bb4e0fb8b6b01a4cafc886b4442",
        strip_prefix = "zeroize-1.5.6",
        build_file = Label("//cargo/remote:BUILD.zeroize-1.5.6.bazel"),
    )
