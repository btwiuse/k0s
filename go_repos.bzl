load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "ag_pack_amqp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "pack.ag/amqp",
        sum = "h1:cuNDWLUTbKRtEZwhB0WQBXf9pGbm87pUBXQhvcFxBWg=",
        version = "v0.11.2",
    )
    go_repository(
        name = "cc_mvdan_interfacer",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "mvdan.cc/interfacer",
        sum = "h1:WX1yoOaKQfddO/mLzdV4wptyWgoH/6hwLs7QHTixo0I=",
        version = "v0.0.0-20180901003855-c20040233aed",
    )
    go_repository(
        name = "cc_mvdan_lint",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "mvdan.cc/lint",
        sum = "h1:DxJ5nJdkhDlLok9K6qO+5290kphDJbHOQO1DFFFTeBo=",
        version = "v0.0.0-20170908181259-adc824a0674b",
    )
    go_repository(
        name = "cc_mvdan_unparam",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "mvdan.cc/unparam",
        sum = "h1:kAREL6MPwpsk1/PQPFD3Eg7WAQR5mPTWZJaBiG5LDbY=",
        version = "v0.0.0-20200501210554-b37ab49443f7",
    )
    go_repository(
        name = "co_honnef_go_tools",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "honnef.co/go/tools",
        sum = "h1:UoveltGrhghAA7ePc+e+QYDHXrBps2PqFZiHkGR/xK8=",
        version = "v0.0.1-2020.1.4",
    )
    go_repository(
        name = "com_gitea_lunny_log",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitea.com/lunny/log",
        sum = "h1:r1en/D7xJmcY24VkHkjkcJFa+7ZWubVWPBrvsHkmHxk=",
        version = "v0.0.0-20190322053110-01b5df579c4e",
    )
    go_repository(
        name = "com_github_abbot_go_http_auth",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/abbot/go-http-auth",
        sum = "h1:QjmvZ5gSC7jm3Zg54DqWE/T5m1t2AfDu6QlXJT0EVT0=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_abiosoft_ishell",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/abiosoft/ishell",
        sum = "h1:zpwIuEHc37EzrsIYah3cpevrIc8Oma7oZPxr03tlmmw=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_abiosoft_readline",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/abiosoft/readline",
        sum = "h1:CjPUSXOiYptLbTdr1RceuZgSFDQ7U15ITERUGrUORx8=",
        version = "v0.0.0-20180607040430-155bce2042db",
    )
    go_repository(
        name = "com_github_activestate_termtest_conpty",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ActiveState/termtest/conpty",
        sum = "h1:JLUe6YDs4Jw4xNPCU+8VwTpniYOGeKzQg4SM2YHQNA8=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_adalogics_go_fuzz_headers",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/AdaLogics/go-fuzz-headers",
        sum = "h1:V8krnnfGj4pV65YLUm3C0/8bl7V5Nry2Pwvy3ru/wLc=",
        version = "v0.0.0-20210715213245-6c3934b029d8",
    )
    go_repository(
        name = "com_github_aead_chacha20",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aead/chacha20",
        sum = "h1:KjTM2ks9d14ZYCvmHS9iAKVt9AyzRSqNU1qabPih5BY=",
        version = "v0.0.0-20180709150244-8b13a72661da",
    )
    go_repository(
        name = "com_github_afex_hystrix_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/afex/hystrix-go",
        sum = "h1:rFw4nCn9iMW+Vajsk51NtYIcwSTkXr+JGrMd36kTDJw=",
        version = "v0.0.0-20180502004556-fa1af6a1f4f5",
    )
    go_repository(
        name = "com_github_ajg_form",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ajg/form",
        sum = "h1:t9c7v8JUKu/XxOGBU0yjNpaMloxGEJhUkqFRq0ibGeU=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_akamai_akamaiopen_edgegrid_golang",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/akamai/AkamaiOPEN-edgegrid-golang",
        sum = "h1:6rJvj+NXjjauunLeS7uGy891F1cuAwsWKa9iGzTjz1s=",
        version = "v0.9.8",
    )
    go_repository(
        name = "com_github_akavel_rsrc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/akavel/rsrc",
        sum = "h1:zjWn7ukO9Kc5Q62DOJCcxGpXC18RawVtYAGdz2aLlfw=",
        version = "v0.8.0",
    )
    go_repository(
        name = "com_github_alangpierce_go_forceexport",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alangpierce/go-forceexport",
        sum = "h1:3ILjVyslFbc4jl1w5TWuvvslFD/nDfR2H8tVaMVLrEY=",
        version = "v0.0.0-20160317203124-8f1d6941cd75",
    )
    go_repository(
        name = "com_github_alcortesm_tgz",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alcortesm/tgz",
        sum = "h1:uSoVVbwJiQipAclBbw+8quDsfcvFjOpI5iCf4p/cqCs=",
        version = "v0.0.0-20161220082320-9c5fe88206d7",
    )
    go_repository(
        name = "com_github_alecthomas_assert",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/assert",
        sum = "h1:smF2tmSOzy2Mm+0dGI2AIUHY+w0BUc+4tn40djz7+6U=",
        version = "v0.0.0-20170929043011-405dbfeb8e38",
    )
    go_repository(
        name = "com_github_alecthomas_chroma",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/chroma",
        sum = "h1:7XDcGkCQopCNKjZHfYrNLraA+M7e0fMiJ/Mfikbfjek=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_alecthomas_colour",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/colour",
        sum = "h1:JHZL0hZKJ1VENNfmXvHbgYlbUOvpzYzvy2aZU5gXVeo=",
        version = "v0.0.0-20160524082231-60882d9e2721",
    )
    go_repository(
        name = "com_github_alecthomas_kingpin",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/kingpin",
        sum = "h1:5svnBTFgJjZvGKyYBtMB0+m5wvrbUHiqye8wRJMlnYI=",
        version = "v2.2.6+incompatible",
    )
    go_repository(
        name = "com_github_alecthomas_kingpin_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/kingpin/v2",
        sum = "h1:f48lwail6p8zpO1bC4TxtqACaGqHYA22qkHjHpqDjYY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "com_github_alecthomas_kong",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/kong",
        sum = "h1:Y0ZBCHAvHhTHw7FFJ2FzCAAG4pkbTgA45nc7BpMhDNk=",
        version = "v0.2.4",
    )
    go_repository(
        name = "com_github_alecthomas_kong_hcl",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/kong-hcl",
        sum = "h1:atLL+K8Hg0e8863K2X+k7qu+xz3M2a/mWFIACAPf55M=",
        version = "v0.1.8-0.20190615233001-b21fea9723c8",
    )
    go_repository(
        name = "com_github_alecthomas_repr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/repr",
        sum = "h1:p9Sln00KOTlrYkxI1zYWl1QLnEqAqEARBEYa8FQnQcY=",
        version = "v0.0.0-20180818092828-117648cd9897",
    )
    go_repository(
        name = "com_github_alecthomas_template",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/template",
        sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
        version = "v0.0.0-20190718012654-fb15b899a751",
    )
    go_repository(
        name = "com_github_alecthomas_units",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/units",
        sum = "h1:s6gZFSlWYmbqAuRjVTiNNhvNRfY2Wxp9nhfyel4rklc=",
        version = "v0.0.0-20211218093645-b94a6e3cc137",
    )
    go_repository(
        name = "com_github_alexflint_go_filemutex",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alexflint/go-filemutex",
        sum = "h1:IAWuUuRYL2hETx5b8vCgwnD+xSdlsTQY6s2JjBsqLdg=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_alexpantyukhin_go_pattern_match",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/alexpantyukhin/go-pattern-match",
        sum = "h1:OtpOxIcYahMMr487L0NWOCSo9RyN4/6+kgD+s5XhRVc=",
        version = "v0.0.0-20230301210247-d84479c117d7",
    )
    go_repository(
        name = "com_github_aliyun_alibaba_cloud_sdk_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aliyun/alibaba-cloud-sdk-go",
        sum = "h1:E273ePcLllLIBGg5BHr3T0Fp1BJTvUyh5Y57ziSy81w=",
        version = "v1.61.112",
    )
    go_repository(
        name = "com_github_aliyun_aliyun_oss_go_sdk",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aliyun/aliyun-oss-go-sdk",
        sum = "h1:nWDRPCyCltiTsANwC/n3QZH7Vww33Npq9MKqlwRzI/c=",
        version = "v0.0.0-20190307165228-86c17b95fcd5",
    )
    go_repository(
        name = "com_github_andreasbriese_bbloom",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/AndreasBriese/bbloom",
        sum = "h1:cTp8I5+VIoKjsnZuH8vjyaysT/ses3EvZeaV/1UkF2M=",
        version = "v0.0.0-20190825152654-46b345b51c96",
    )
    go_repository(
        name = "com_github_andrew_d_go_termutil",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/andrew-d/go-termutil",
        sum = "h1:axBiC50cNZOs7ygH5BgQp4N+aYrZ2DNpWZ1KG3VOSOM=",
        version = "v0.0.0-20150726205930-009166a695a2",
    )
    go_repository(
        name = "com_github_andybalholm_brotli",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/andybalholm/brotli",
        sum = "h1:V7DdXeJtZscaqfNuAdSRuRFzuiKlHSC/Zh3zl9qY3JY=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_anmitsu_go_shlex",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/anmitsu/go-shlex",
        sum = "h1:9AeTilPcZAjCFIImctFaOjnTIavg87rW78vTPkQqLI8=",
        version = "v0.0.0-20200514113438-38f4b401e2be",
    )
    go_repository(
        name = "com_github_antihax_optional",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/antihax/optional",
        sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_antlr_antlr4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/antlr/antlr4",
        sum = "h1:0cEys61Sr2hUBEXfNV8eyQP01oZuBgoMeHunebPirK8=",
        version = "v0.0.0-20200503195918-621b933c7a7f",
    )
    go_repository(
        name = "com_github_antlr_antlr4_runtime_go_antlr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/antlr/antlr4/runtime/Go/antlr",
        sum = "h1:ue9pVfIcP+QMEjfgo/Ez4ZjNZfonGgR6NgjMaJMu1Cg=",
        version = "v0.0.0-20220418222510-f25a4f6275ed",
    )
    go_repository(
        name = "com_github_antlr_antlr4_runtime_go_antlr_v4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/antlr/antlr4/runtime/Go/antlr/v4",
        sum = "h1:7RFfzj4SSt6nnvCPbCqijJi1nWCd+TqAT3bYCStRC18=",
        version = "v4.0.0-20230305170008-8188dc5388df",
    )
    go_repository(
        name = "com_github_aokoli_goutils",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aokoli/goutils",
        sum = "h1:7fpzNGoJ3VA8qcrm++XEE1QUe0mIwNeLa02Nwq7RDkg=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_apache_beam",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/apache/beam",
        sum = "h1:iYCqXSPdvj8YW6LS3MleV080RZKf5ULS+L490PjJY7Y=",
        version = "v2.30.0+incompatible",
    )
    go_repository(
        name = "com_github_apache_thrift",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/apache/thrift",
        sum = "h1:5hryIiq9gtn+MiLVn0wP37kb/uTeRZgN08WoCsAhIhI=",
        version = "v0.13.0",
    )
    go_repository(
        name = "com_github_apex_log",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/apex/log",
        sum = "h1:1fyfbPvUwD10nMoh3hY6MXzvZShJQn9/ck7ATgAt5pA=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_apex_logs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/apex/logs",
        sum = "h1:KmEBVwfDUOTFcBO8cfkJYwdQ5487UZSN+GteOGPmiro=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_aphistic_golf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aphistic/golf",
        sum = "h1:2KLQMJ8msqoPHIPDufkxVcoTtcmE5+1sL9950m4R9Pk=",
        version = "v0.0.0-20180712155816-02c07f170c5a",
    )
    go_repository(
        name = "com_github_aphistic_sweet",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aphistic/sweet",
        sum = "h1:I4z+fAUqvKfvZV/CHi5dV0QuwbmIvYYFDjG0Ss5QpAs=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_armon_circbuf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/circbuf",
        sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
        version = "v0.0.0-20150827004946-bbbad097214e",
    )
    go_repository(
        name = "com_github_armon_consul_api",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/consul-api",
        sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
        version = "v0.0.0-20180202201655-eb2c6b5be1b6",
    )
    go_repository(
        name = "com_github_armon_go_metrics",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/go-metrics",
        sum = "h1:O2sNqxBdvq8Eq5xmzljcYzAORli6RWCvEym4cJf9m18=",
        version = "v0.3.9",
    )
    go_repository(
        name = "com_github_armon_go_radix",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/go-radix",
        sum = "h1:F4z6KzEeeQIMeLFa97iZU6vupzoecKdU5TX24SNppXI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_armon_go_socks5",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/go-socks5",
        sum = "h1:0CwZNZbxp69SHPdPJAN/hZIm0C4OItdklCFmMRWYpio=",
        version = "v0.0.0-20160902184237-e75332964ef5",
    )
    go_repository(
        name = "com_github_aryann_difflib",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aryann/difflib",
        sum = "h1:uUXgbcPDK3KpW29o4iy7GtuappbWT0l5NaMo9H9pJDw=",
        version = "v0.0.0-20210328193216-ff5ff6dc229b",
    )
    go_repository(
        name = "com_github_asaskevich_govalidator",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/asaskevich/govalidator",
        sum = "h1:Byv0BzEl3/e6D5CLfI0j/7hiIEtvGVFPCZ7Ei2oq8iQ=",
        version = "v0.0.0-20210307081110-f21760c49a8d",
    )
    go_repository(
        name = "com_github_asdine_storm",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/asdine/storm",
        sum = "h1:dczuIkyqwY2LrtXPz8ixMrU/OFgZp71kbKTHGrXYt/Q=",
        version = "v2.1.2+incompatible",
    )
    go_repository(
        name = "com_github_aws_aws_lambda_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-lambda-go",
        sum = "h1:SuCy7H3NLyp+1Mrfp+m80jcbi9KYWAs9/BXwppwRDzY=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go",
        sum = "h1:zJnJ8Y964DjyRE55UVoMKgOG4w5i88LpN6xSpBX7z84=",
        version = "v1.41.14",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2",
        sum = "h1:GzvOVAdTbWxhEMRK4FfiblkGverOkAT0UodDxC1jHQM=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_config",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/config",
        sum = "h1:ZAoq32boMzcaTW9bcUacBswAmHTbvlvDJICgHFZuECo=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_credentials",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/credentials",
        sum = "h1:2faRNX8JgZVy7dDxERkaGBqb/xo5Rgmc8JMPL5j1o58=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_feature_ec2_imds",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/feature/ec2/imds",
        sum = "h1:EtEU7WRaWliitZh2nmuxEXrN0Cb8EgPUFGIoTMeqbzI=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_configsources",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/configsources",
        sum = "h1:LZwqhOyqQ2w64PZk04V0Om9AEExtW8WMkCRoE1h9/94=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_endpoints_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2",
        sum = "h1:ObMfGNk0xjOWduPxsrRWVwZZia3e9fOcO6zlKCkt38s=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_ecr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/ecr",
        sum = "h1:onTF83DG9dsRv6UzuhYb7phiktjwQ++s/n+ZtNlTQnM=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_presigned_url",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url",
        sum = "h1:4AH9fFjUlVktQMznF+YN33aWNXaR4VgDXyP28qokJC0=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_route53",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/route53",
        sum = "h1:cKr6St+CtC3/dl/rEBJvlk7A/IN5D5F02GNkGzfbtVU=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sso",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sso",
        sum = "h1:37QubsarExl5ZuCBlnRP+7l1tNwZPBSTqpTBrPH98RU=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sts",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sts",
        sum = "h1:TJoIfnIFubCX0ACVeJ0w46HEH5MwjwYN4iFhuYIhfIY=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_aws_smithy_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/smithy-go",
        sum = "h1:c7FUdEqrQA1/UVKKCNDFQPNKGp4FQg3YW4Ck5SLTG58=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_aybabtme_rgbterm",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/aybabtme/rgbterm",
        sum = "h1:WWB576BN5zNSZc/M9d/10pqEx5VHNhaQ/yOVAkmj5Yo=",
        version = "v0.0.0-20170906152045-cc83f3b3ce59",
    )
    go_repository(
        name = "com_github_azure_azure_amqp_common_go_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-amqp-common-go/v2",
        sum = "h1:+QbFgmWCnPzdaRMfsI0Yb6GrRdBj5jVL8N3EXuEUcBQ=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_azure_azure_pipeline_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-pipeline-go",
        sum = "h1:6oiIS9yaG6XCCzhgAgKFfIWyo4LLCiDhZot6ltoThhY=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-sdk-for-go",
        sum = "h1:Cw16jiP4dI+CK761aq44ol4RV5dUiIIXky1+EKpoiVM=",
        version = "v58.0.0+incompatible",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_azcore",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/azcore",
        sum = "h1:qoVeMsc9/fh/yhxVaA0obYjVH/oI/ihrOoMwsLS9KSA=",
        version = "v0.21.1",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_internal",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/internal",
        sum = "h1:E+m3SkZCN0Bf5q7YdTs5lSm2CYY3CK4spn5OmUIiQtk=",
        version = "v0.8.3",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_storage_azblob",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob",
        sum = "h1:Px2UA+2RvSSvv+RvJNuUB6n7rs5Wsel4dXLe90Um2n4=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_azure_azure_service_bus_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-service-bus-go",
        sum = "h1:G1qBLQvHCFDv9pcpgwgFkspzvnGknJRR0PYJ9ytY/JA=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_azure_azure_storage_blob_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-storage-blob-go",
        sum = "h1:53qhf0Oxa0nOjgbDeeYPUeyiNmafAFEY95rZLK0Tj6o=",
        version = "v0.8.0",
    )
    go_repository(
        name = "com_github_azure_go_ansiterm",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-ansiterm",
        sum = "h1:UQHMgLO+TxOElx5B5HZ4hJQsoJ/PvUvKRhJHDQXO8P8=",
        version = "v0.0.0-20210617225240-d185dfc1b5a1",
    )
    go_repository(
        name = "com_github_azure_go_autorest",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest",
        sum = "h1:V5VMDjClD3GiElqLWO7mz2MxNAK/vTfRHdAubSIPRgs=",
        version = "v14.2.0+incompatible",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest",
        sum = "h1:90Y4srNYrwOtAgVo3ndrQkTYn6kf1Eg/AjTFJ8Is2aM=",
        version = "v0.11.18",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_adal",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/adal",
        sum = "h1:Mp5hbtOePIzM8pJVRa3YLrWWmZtoxRXqUEzCfJt3+/Q=",
        version = "v0.9.13",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_azure_auth",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/azure/auth",
        sum = "h1:TzPg6B6fTZ0G1zBf3T54aI7p3cAT6u//TOXGPmFMOXg=",
        version = "v0.5.8",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_azure_cli",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/azure/cli",
        sum = "h1:dMOmEJfkLKW/7JsokJqkyoYSgmR08hi9KrhjZb+JALY=",
        version = "v0.4.2",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_date",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/date",
        sum = "h1:7gUk1U5M/CQbp9WoqinNzJar+8KY+LPI6wiWrP/myHw=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_mocks",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/mocks",
        sum = "h1:K0laFcLE6VLTOwNgSxaGbUcLPuGXlNkbVvq4cW4nIHk=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_to",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/to",
        sum = "h1:oXVqrxakqqV1UZdSazDOPOLvOIz+XA683u8EctwboHk=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_validation",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/validation",
        sum = "h1:AgyqjAd94fwNAoTjl/WQXg4VvFeRFpO+UhNyRXqF1ac=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_azure_go_autorest_logger",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/logger",
        sum = "h1:IG7i4p/mDa2Ce4TRyAO8IHnVhAVF3RFU+ZtXWSmf4Tg=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_azure_go_autorest_tracing",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/tracing",
        sum = "h1:TYi4+3m5t6K48TGI9AUdb+IzbnSxvnvUMfuitfgcfuo=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_baiyubin_aliyun_sts_go_sdk",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/baiyubin/aliyun-sts-go-sdk",
        sum = "h1:ZNv7On9kyUzm7fvRZumSyy/IUiSC7AzL0I1jKKtwooA=",
        version = "v0.0.0-20180326062324-cfa1a18b161f",
    )
    go_repository(
        name = "com_github_beevik_ntp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/beevik/ntp",
        sum = "h1:/w5VhpW5BGKS37vFm1p9oVk/t4HnnkKZAZIubHM6F7Q=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_benbjohnson_clock",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/benbjohnson/clock",
        sum = "h1:Q92kusRqC1XV2MjkWETPvjJVqKetz1OzxZB7mHJLju8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_bgentry_speakeasy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bgentry/speakeasy",
        sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_bifurcation_mint",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bifurcation/mint",
        sum = "h1:fUjoj2bT6dG8LoEe+uNsKk8J+sLkDbQkJnB6Z1F02Bc=",
        version = "v0.0.0-20180715133206-93c51c6ce115",
    )
    go_repository(
        name = "com_github_bitly_go_simplejson",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bitly/go-simplejson",
        sum = "h1:6IH+V8/tVMab511d5bn4M7EwGXZf9Hj6i2xSwkNEM+Y=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_bits_and_blooms_bitset",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bits-and-blooms/bitset",
        sum = "h1:Kn4yilvwNtMACtf1eYDlG8H77R07mZSPbMjLyS07ChA=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_bketelsen_crypt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bketelsen/crypt",
        sum = "h1:+0HFd5KSZ/mm3JmhmrDukiId5iR6w4+BdFtfSy4yWIc=",
        version = "v0.0.3-0.20200106085610-5cbc8cc4026c",
    )
    go_repository(
        name = "com_github_blakesmith_ar",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/blakesmith/ar",
        sum = "h1:m935MPodAbYS46DG4pJSv7WO+VECIWUQ7OJYSoTrMh4=",
        version = "v0.0.0-20190502131153-809d4375e1fb",
    )
    go_repository(
        name = "com_github_blang_semver",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/blang/semver",
        sum = "h1:cQNTCjp13qL8KC3Nbxr/y2Bqb63oX6wdnnjpJbkM4JQ=",
        version = "v3.5.1+incompatible",
    )
    go_repository(
        name = "com_github_blang_semver_v4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/blang/semver/v4",
        sum = "h1:1PFHFE6yCCTv8C1TeyNNarDzntLi7wMI5i/pzqYIsAM=",
        version = "v4.0.0",
    )
    go_repository(
        name = "com_github_bmizerany_assert",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bmizerany/assert",
        sum = "h1:DDGfHa7BWjL4YnC6+E63dPcxHo2sUxDIu8g3QgEJdRY=",
        version = "v0.0.0-20160611221934-b7ed37b82869",
    )
    go_repository(
        name = "com_github_boltdb_bolt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/boltdb/bolt",
        sum = "h1:JQmyP4ZBrce+ZQu0dY660FMfatumYDLun9hBCUVIkF4=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_bombsimon_wsl_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bombsimon/wsl/v2",
        sum = "h1:/DdSteYCq4lPX+LqDg7mdoxm14UxzZPoDT0taYc3DTU=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_bombsimon_wsl_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bombsimon/wsl/v3",
        sum = "h1:E5SRssoBgtVFPcYWUOFJEcgaySgdtTNYzsSKDOY7ss8=",
        version = "v3.1.0",
    )
    go_repository(
        name = "com_github_boombuler_barcode",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/boombuler/barcode",
        sum = "h1:NDBbPmhS+EqABEs5Kg3n/5ZNjy73Pz7SIV+KCeqyXcs=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_bradfitz_go_smtpd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bradfitz/go-smtpd",
        sum = "h1:ckJgFhFWywOx+YLEMIJsTb+NV6NexWICk5+AMSuz3ss=",
        version = "v0.0.0-20170404230938-deb6d6237625",
    )
    go_repository(
        name = "com_github_bshuster_repo_logrus_logstash_hook",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bshuster-repo/logrus-logstash-hook",
        sum = "h1:pgAtgj+A31JBVtEHu2uHuEx0n+2ukqUJnS2vVe5pQNA=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_btcsuite_btcd_btcec_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/btcsuite/btcd/btcec/v2",
        sum = "h1:fzn1qaOt32TuLjFlkzYSsBC35Q3KUjT1SwPxiMSCF5k=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_btwiuse_bcrypt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/btwiuse/bcrypt",
        sum = "h1:je6QGuEHitc4zAnAL/2w25WyHcSu+Iz/e5t4TRlvG1s=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_btwiuse_bingo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/btwiuse/bingo",
        sum = "h1:HgTIAC+ZAQ61MJLtQuf+wfYdUIpgF51qo0dCNEeO08k=",
        version = "v0.0.0-20240130160828-410a95142901",
    )
    go_repository(
        name = "com_github_btwiuse_cadvisor",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/btwiuse/cadvisor",
        sum = "h1:TN7StIvTOGlPb0jSnp3UIUCXkXzD53UUGZL0ZhyGA+E=",
        version = "v0.0.0-20210312172035-34fddda41018",
    )
    go_repository(
        name = "com_github_btwiuse_dkg",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/btwiuse/dkg",
        sum = "h1:1tIJhP64piTclCgtmtZNwcX2WKCgQsPXV6jENiu/h0A=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_btwiuse_gods",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/btwiuse/gods",
        sum = "h1:LZs8rCGipp9xKSjRpENvVijtxq60aATn6UL357/bVvU=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_btwiuse_gost",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/btwiuse/gost",
        sum = "h1:Ftd0gcqPU2HeDUYPzY8LOx391ZKK/N/+ffbMCnt8OM4=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_btwiuse_multicall",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/btwiuse/multicall",
        sum = "h1:jsvjulx9kBJplbZ85UwSlMHuxzhO17VMzgBaTeYI10k=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_btwiuse_pretty",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/btwiuse/pretty",
        sum = "h1:P6/2XTs08jm4T/Vqgb2fyrVtO1zTUGpBX48+jjfs1tI=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_btwiuse_sse",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/btwiuse/sse",
        sum = "h1:tXLz7isF3VeSQGLZUQCWgaM5tZIbb78nn1QW9/nkI7s=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_btwiuse_version",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/btwiuse/version",
        sum = "h1:L0168/a3iSVD1EXf6cBXLmp3EtVvNUpc8yXM2RJ6Blc=",
        version = "v0.0.0-20240130171800-042c8be1e199",
    )
    go_repository(
        name = "com_github_btwiuse_wetty",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/btwiuse/wetty",
        sum = "h1:MSMG+C1wR+bgG5xB2TRZK+5yxYiOmGwe44c/3gpaGT0=",
        version = "v0.0.36",
    )
    go_repository(
        name = "com_github_buger_jsonparser",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/buger/jsonparser",
        sum = "h1:2PnMjfWD7wBILjqQbt530v576A/cAbQvEW9gGIpYMUs=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_bugsnag_bugsnag_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bugsnag/bugsnag-go",
        sum = "h1:rFt+Y/IK1aEZkEHchZRSq9OQbsSzIT/OrI8YFFmRIng=",
        version = "v0.0.0-20141110184014-b1d153021fcd",
    )
    go_repository(
        name = "com_github_bugsnag_osext",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bugsnag/osext",
        sum = "h1:otBG+dV+YK+Soembjv71DPz3uX/V/6MMlSyD9JBQ6kQ=",
        version = "v0.0.0-20130617224835-0dd3f918b21b",
    )
    go_repository(
        name = "com_github_bugsnag_panicwrap",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/bugsnag/panicwrap",
        sum = "h1:nvj0OLI3YqYXer/kZD8Ri1aaunCxIEsOst1BVJswV0o=",
        version = "v0.0.0-20151223152923-e2c28503fcd0",
    )
    go_repository(
        name = "com_github_buildkite_agent_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/buildkite/agent/v3",
        sum = "h1:eMqvu48qkv89wvKyllIZ7CZOn/bHsaTeh8ZlTxUKIYM=",
        version = "v3.27.0",
    )
    go_repository(
        name = "com_github_buildkite_bintest_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/buildkite/bintest/v3",
        sum = "h1:auJ22sFpyy7t6xR7p5FcqAwpNgkP0AyVhEMSRErFmk0=",
        version = "v3.1.0",
    )
    go_repository(
        name = "com_github_buildkite_interpolate",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/buildkite/interpolate",
        sum = "h1:k6UDF1uPYOs0iy1HPeotNa155qXRWrzKnqAaGXHLZCE=",
        version = "v0.0.0-20200526001904-07f35b4ae251",
    )
    go_repository(
        name = "com_github_buildkite_shellwords",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/buildkite/shellwords",
        sum = "h1:hiVSLk7s3yFKFOHF/huoShLqrj13RMguWX2yzfvy7es=",
        version = "v0.0.0-20180315084142-c3f497d1e000",
    )
    go_repository(
        name = "com_github_buildkite_yaml",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/buildkite/yaml",
        sum = "h1:q+sMKdA6L8LyGVudTkpGoC73h6ak2iWSPFiFo/pFOU8=",
        version = "v0.0.0-20181016232759-0caa5f0796e3",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:ksErzDEI1khOiGPgpwuI7x2ebx/uXQNw7xJpn9Eq1+I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_burntsushi_xgb",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/BurntSushi/xgb",
        sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
        version = "v0.0.0-20160522181843-27f122750802",
    )
    go_repository(
        name = "com_github_caarlos0_ctrlc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caarlos0/ctrlc",
        sum = "h1:2DtF8GSIcajgffDFJzyG15vO+1PuBWOMUdFut7NnXhw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_caddy_dns_alidns",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caddy-dns/alidns",
        sum = "h1:B2JSeQrx2S3xgO5bte7SANmBhyoFfSd9v/Blo631gRc=",
        version = "v1.0.23",
    )
    go_repository(
        name = "com_github_caddy_dns_azure",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caddy-dns/azure",
        sum = "h1:j38dM8EAtK0u3CzKUML736MomOsbvyylMQ2hhZ3gUS0=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_caddy_dns_cloudflare",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caddy-dns/cloudflare",
        sum = "h1:usrheD4ZMCLeI7CI+vHff7+I+D8wIL8Gr3dTlkSmkBA=",
        version = "v0.0.0-20210607183747-91cf700356a1",
    )
    go_repository(
        name = "com_github_caddy_dns_digitalocean",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caddy-dns/digitalocean",
        sum = "h1:5JXI9U73KPkSnEveF8DTiSpZiZkNOgOr0mMJZeDx4s8=",
        version = "v0.0.0-20210809220558-ac6e4fd9e135",
    )
    go_repository(
        name = "com_github_caddy_dns_dnspod",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caddy-dns/dnspod",
        sum = "h1:txebUF94u/Y8Wt6+hZCmBlD7UQFaNbUUMcAEq8PjgRk=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_caddy_dns_duckdns",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caddy-dns/duckdns",
        sum = "h1:SSqTt/kbtV9RmJ4wqys5j2ltDwfF35Tv+UU4qaOmTpY=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_caddy_dns_route53",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caddy-dns/route53",
        sum = "h1:msy8aKay1JwTgN2u4kFlqREjlavpUX8mjHVFS+ZDqEY=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_caddy_dns_vultr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caddy-dns/vultr",
        sum = "h1:L86RWlbq2cObphvPPnUbL4Ap+gcF72NBUga9i4iPV34=",
        version = "v0.0.0-20211122185502-733392841379",
    )
    go_repository(
        name = "com_github_caddyserver_caddy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caddyserver/caddy",
        sum = "h1:5B1Hs0UF2x2tggr2X9jL2qOZtDXbIWQb9YLbmlxHSuM=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_caddyserver_caddy_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caddyserver/caddy/v2",
        sum = "h1:eCJdLyEyAGzuQTa5Mh3gETnYWDClo1LjtQm2q9RNZrs=",
        version = "v2.5.2",
    )
    go_repository(
        name = "com_github_caddyserver_certmagic",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caddyserver/certmagic",
        sum = "h1:rdSnjcUVJojmL4M0efJ+yHXErrrijS4YYg3FuwRdJkI=",
        version = "v0.16.1",
    )
    go_repository(
        name = "com_github_caddyserver_forwardproxy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/caddyserver/forwardproxy",
        sum = "h1:wC4VlHOHWTjslaf/SBQ4nFaU1aVUZFRqyzZRnCJNV7k=",
        version = "v0.0.0-20211013034647-8c6ef2bd4a8f",
    )
    go_repository(
        name = "com_github_campoy_unique",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/campoy/unique",
        sum = "h1:V9a67dfYqPLAvzk5hMQOXYJlZ4SLIXgyKIE+ZiHzgGQ=",
        version = "v0.0.0-20180121183637-88950e537e7e",
    )
    go_repository(
        name = "com_github_casbin_caddy_authz_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/casbin/caddy-authz/v2",
        sum = "h1:vUUoBIbTa8k02/zF17Qk0EjJIsu3HSU0ee2lX1nKo54=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_casbin_casbin_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/casbin/casbin/v2",
        sum = "h1:JApbUmymvG33xIIYJ4G4ijj20tlZ9b8LY4ByTe+Oz+M=",
        version = "v2.8.6",
    )
    go_repository(
        name = "com_github_cavaliercoder_go_cpio",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cavaliercoder/go-cpio",
        sum = "h1:hHg27A0RSSp2Om9lubZpiMgVbvn39bsUmW9U5h0twqc=",
        version = "v0.0.0-20180626203310-925f9528c45e",
    )
    go_repository(
        name = "com_github_cenkalti_backoff",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cenkalti/backoff",
        sum = "h1:tNowT99t7UNflLxfYYSlKYsBpXdEet03Pg2g16Swow4=",
        version = "v2.2.1+incompatible",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cenkalti/backoff/v3",
        sum = "h1:ske+9nBpD9qZsTBoF41nW5L+AIuFBKMeze18XQ3eG1c=",
        version = "v3.0.0",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cenkalti/backoff/v4",
        sum = "h1:y4OZtCnogmCPw98Zjyt5a6+QwPLGkiQsYW5oUqylYbM=",
        version = "v4.2.1",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_certifi_gocertifi",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/certifi/gocertifi",
        sum = "h1:uH66TXeswKn5PW5zdZ39xEwfS9an067BirqA+P4QaLI=",
        version = "v0.0.0-20200922220541-2c3bb06c6054",
    )
    go_repository(
        name = "com_github_cespare_cp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cespare/cp",
        sum = "h1:SE+dxFebS7Iik5LK0tsi1k9ZCxEaFX4AjQmoyA+1dJk=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_cespare_xxhash",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cespare/xxhash",
        sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_chai2010_gettext_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/chai2010/gettext-go",
        sum = "h1:1Lwwip6Q2QGsAdl/ZKPCwTe9fe0CjlUbqj5bFNSjIRk=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_checkpoint_restore_go_criu_v4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/checkpoint-restore/go-criu/v4",
        sum = "h1:WW2B2uxx9KWF6bGlHqhm8Okiafwwx7Y2kcpn8lCpjgo=",
        version = "v4.1.0",
    )
    go_repository(
        name = "com_github_checkpoint_restore_go_criu_v5",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/checkpoint-restore/go-criu/v5",
        sum = "h1:wpFFOoomK3389ue2lAb0Boag6XPht5QYpipxmSNL4d8=",
        version = "v5.3.0",
    )
    go_repository(
        name = "com_github_cheekybits_genny",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cheekybits/genny",
        sum = "h1:uGGa4nei+j20rOSeDeP5Of12XVm7TGUd4dJA9RDitfE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_chromedp_cdproto",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/chromedp/cdproto",
        sum = "h1:aPflPkRFkVwbW6dmcVqfgwp1i+UWGFH6VgR1Jim5Ygc=",
        version = "v0.0.0-20230802225258-3cf4e6d46a89",
    )
    go_repository(
        name = "com_github_chromedp_chromedp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/chromedp/chromedp",
        sum = "h1:dKtNz4kApb06KuSXoTQIyUC2TrA0fhGDwNZf3bcgfKw=",
        version = "v0.9.2",
    )
    go_repository(
        name = "com_github_chromedp_sysutil",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/chromedp/sysutil",
        sum = "h1:+ZxhTpfpZlmchB58ih/LBHX52ky7w2VhQVKQMucy3Ic=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_chzyer_logex",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/chzyer/logex",
        sum = "h1:XHDu3E6q+gdHgsdTPH6ImJMIp436vR6MPtH8gP05QzM=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_chzyer_readline",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/chzyer/readline",
        sum = "h1:upd/6fQk4src78LMRzh5vItIt361/o4uq553V8B5sGI=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_chzyer_test",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/chzyer/test",
        sum = "h1:p3BQDXSxOhOG0P9z6/hGnII4LGiEPOYBhs8asl/fC04=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_cilium_ebpf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cilium/ebpf",
        sum = "h1:V8gS/bTCCjX9uUnkUFUpPsksM8n1lXBAvHcpiFk1X2Y=",
        version = "v0.11.0",
    )
    go_repository(
        name = "com_github_circonus_labs_circonus_gometrics",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/circonus-labs/circonus-gometrics",
        sum = "h1:C29Ae4G5GtYyYMm1aztcyj/J5ckgJm2zwdDajFbx1NY=",
        version = "v2.3.1+incompatible",
    )
    go_repository(
        name = "com_github_circonus_labs_circonusllhist",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/circonus-labs/circonusllhist",
        sum = "h1:TJH+oke8D16535+jHExHj4nQvzlZrj7ug5D7I/orNUA=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_clbanning_x2j",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/clbanning/x2j",
        sum = "h1:EdRZT3IeKQmfCSrgo8SZ8V3MEnskuJP0wCYNpe+aiXo=",
        version = "v0.0.0-20191024224557-825249438eec",
    )
    go_repository(
        name = "com_github_client9_misspell",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_cloudflare_cloudflare_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cloudflare/cloudflare-go",
        sum = "h1:gFqGlGl/5f9UGXAaKapCGUfaTCgRKKnzu2VvzMZlOFA=",
        version = "v0.14.0",
    )
    go_repository(
        name = "com_github_cncf_udpa_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cncf/udpa/go",
        sum = "h1:QQ3GSy+MqSHxm/d8nCtnAiZdYFd45cYZPs8vOOIYKfk=",
        version = "v0.0.0-20220112060539-c52dc94e7fbe",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:/inchEIKaYC1Akx+H+gqO04wryn5h75LSazbRlnya1k=",
        version = "v0.0.0-20230607035331-e9ce68804cb4",
    )
    go_repository(
        name = "com_github_cockroachdb_apd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/apd",
        sum = "h1:3LFP3629v+1aKXU5Q37mxmRxX/pIu1nijXydLShEq5I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_cockroachdb_datadriven",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/datadriven",
        sum = "h1:xD/lrqdvwsc+O2bjSSi3YqY73Ke3LAiSCx49aCesA0E=",
        version = "v0.0.0-20200714090401-bf6692d28da5",
    )
    go_repository(
        name = "com_github_cockroachdb_errors",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/errors",
        sum = "h1:Lap807SXTH5tri2TivECb/4abUkMZC9zRoLarvcKDqs=",
        version = "v1.2.4",
    )
    go_repository(
        name = "com_github_cockroachdb_logtags",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/logtags",
        sum = "h1:o/kfcElHqOiXqcou5a3rIlMc7oJbMQkeLk0VQJ7zgqY=",
        version = "v0.0.0-20190617123548-eb05cc24525f",
    )
    go_repository(
        name = "com_github_codahale_hdrhistogram",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/codahale/hdrhistogram",
        sum = "h1:qMd81Ts1T2OTKmB4acZcyKaMtRnY5Y44NuXGX2GFJ1w=",
        version = "v0.0.0-20161010025455-3a0bb77429bd",
    )
    go_repository(
        name = "com_github_codegangsta_cli",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/codegangsta/cli",
        sum = "h1:iX1FXEgwzd5+XN6wk5cVHOGQj6Q3Dcp20lUeS4lHNTw=",
        version = "v1.20.0",
    )
    go_repository(
        name = "com_github_consensys_gnark_crypto",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/consensys/gnark-crypto",
        sum = "h1:C43yEtQ6NIf4ftFXD/V55gnGFgPbMQobd//YlnLjUJ8=",
        version = "v0.4.1-0.20210426202927-39ac3d4b3f1f",
    )
    go_repository(
        name = "com_github_containerd_aufs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/aufs",
        sum = "h1:2oeJiwX5HstO7shSrPZjrohJZLzK36wvpdmzDRkL/LY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_btrfs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/btrfs",
        sum = "h1:osn1exbzdub9L5SouXO5swW4ea/xVdJZ3wokxN5GrnA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_cgroups",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/cgroups",
        sum = "h1:ADZftAkglvCiD44c77s5YmMqaP2pzVCFZvBmAlBdAP4=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_containerd_console",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/console",
        sum = "h1:lIr7SlA5PxZyMV30bDW0MGbiOPXwc63yRuCP0ARubLw=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_containerd_containerd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/containerd",
        sum = "h1:oa2uY0/0G+JX4X7hpGCYvkp9FjUancz56kSNnb1sG3o=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_github_containerd_continuity",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/continuity",
        sum = "h1:QSqfxcn8c+12slxwu00AtzXrsami0MJb/MQs9lOLHLA=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_containerd_fifo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/fifo",
        sum = "h1:6PirWBr9/L7GDamKr+XM0IeUFXu5mf3M/BPpH9gaLBU=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_go_cni",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/go-cni",
        sum = "h1:t0MQwrtM96SH71Md8tH0uKrVE9v+jxkDTbvFSm3B9VE=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_containerd_go_runc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/go-runc",
        sum = "h1:oU+lLv1ULm5taqgV/CJivypVODI4SUz1znWjv3nNYS0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_imgcrypt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/imgcrypt",
        sum = "h1:69UKRsA3Q/lAwo2eDzWshdjimqhmprrWXfNtBeO0fBc=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_containerd_nri",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/nri",
        sum = "h1:6QioHRlThlKh2RkRTR4kIT3PKAcrLo3gIWnjkM4dQmQ=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_containerd_stargz_snapshotter_estargz",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/stargz-snapshotter/estargz",
        sum = "h1:5e7heayhB7CcgdTkqfZqrNaNv15gABwr3Q2jBTbLlt4=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_containerd_ttrpc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/ttrpc",
        sum = "h1:GbtyLRxb0gOLR0TYQWt3O6B0NvT8tMdorEHqIQo/lWI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_containerd_typeurl",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/typeurl",
        sum = "h1:Chlt8zIieDbzQFzXzAeBEF92KhExuE4p9p92/QmY7aY=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_containerd_zfs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/zfs",
        sum = "h1:cXLJbx+4Jj7rNsTiqVfm6i+RNLx6FFA2fMmDlEf+Wm8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containernetworking_cni",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containernetworking/cni",
        sum = "h1:9OIL/sZmMYDBe+G8svzILAlulUpaDTUjeAbtH/JNLBo=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_containernetworking_plugins",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containernetworking/plugins",
        sum = "h1:wwCfYbTCj5FC0EJgyzyjTXmqysOiJE9r712Z+2KVZAk=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_containers_ocicrypt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/containers/ocicrypt",
        sum = "h1:Ez+GAMP/4GLix5Ywo/fL7O0nY771gsBIigiqUm1aXz0=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_coreos_bbolt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/bbolt",
        sum = "h1:n6AiVyVRKQFNb6mJlwESEvvLoDyiTzXX7ORAUlkeBdY=",
        version = "v1.3.3",
    )
    go_repository(
        name = "com_github_coreos_etcd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/etcd",
        sum = "h1:Zz1aXgDrFFi1nadh58tA9ktt06cmPTwNNP3dXwIq1lE=",
        version = "v3.3.18+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_etcd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-etcd",
        sum = "h1:bXhRBIXoTm9BYHS3gE0TtQuyNZyeEMux2sDi4oo5YOo=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_iptables",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-iptables",
        sum = "h1:is9qnZMPYjLd8LYqmm/qlE+wwEgJIkTYdhV3rfZo4jk=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_coreos_go_oidc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-oidc",
        sum = "h1:mh48q/BqXqgjVHpy2ZY7WnWAbenxRjsz9N1i1YxjHAk=",
        version = "v2.2.1+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_semver",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-semver",
        sum = "h1:yi21YpKnrx1gt5R+la8n5WgS0kCrsPp33dmEyHReZr4=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_coreos_go_systemd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-systemd",
        sum = "h1:JOrtw2xFKzlg+cbHpyrpLDmnN1HqhBfnX7WDiW7eG2c=",
        version = "v0.0.0-20190719114852-fd7a80b32e1f",
    )
    go_repository(
        name = "com_github_coreos_go_systemd_v22",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-systemd/v22",
        sum = "h1:RrqgGjYQKalulkV8NGVIfkXQf6YYmOyiJKk8iXXhfZs=",
        version = "v22.5.0",
    )
    go_repository(
        name = "com_github_coreos_pkg",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/pkg",
        sum = "h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
        version = "v0.0.0-20180928190104-399ea9e2e55f",
    )
    go_repository(
        name = "com_github_corpix_uarand",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/corpix/uarand",
        sum = "h1:RMr1TWc9F4n5jiPDzFHtmaUXLKLNUFK0SgCLo4BhX/U=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_cppforlife_cobrautil",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cppforlife/cobrautil",
        sum = "h1:rPWcUBgMb1ox2eCohCuZ8gsZVe0aB5qBbYaBpdoxfCE=",
        version = "v0.0.0-20200514214827-bb86e6965d72",
    )
    go_repository(
        name = "com_github_cppforlife_color",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cppforlife/color",
        sum = "h1:mYQweUIBD+TBRjIeQnJmXr0GSVMpI6O0takyb/aaOgo=",
        version = "v1.9.1-0.20200716202919-6706ac40b835",
    )
    go_repository(
        name = "com_github_cppforlife_go_cli_ui",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cppforlife/go-cli-ui",
        sum = "h1:yVW0v4zDXzJo1i8G9G3vtvNpyzhvtLalO34BsN/K88E=",
        version = "v0.0.0-20200716203538-1e47f820817f",
    )
    go_repository(
        name = "com_github_cpu_goacmedns",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cpu/goacmedns",
        sum = "h1:hYAgjnPu7HogTgb8trqQouR/RrBgXq1TPBgmxbK9eRA=",
        version = "v0.0.2",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cpuguy83/go-md2man",
        sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
        version = "v1.0.10",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cpuguy83/go-md2man/v2",
        sum = "h1:p1EgwI/C7NhT0JmVkwCD2ZBK8j4aeHQX2pMHHBfMQ6w=",
        version = "v2.0.2",
    )
    go_repository(
        name = "com_github_creack_pty",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/creack/pty",
        sum = "h1:1/QdRyBaHHJP61QkWMXlOIBfsgdDeeKfK8SYVUWJKf0=",
        version = "v1.1.21",
    )
    go_repository(
        name = "com_github_creativeprojects_go_selfupdate",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/creativeprojects/go-selfupdate",
        sum = "h1:p+Mx6rCZGBMrYWa5XhHCHYsS+w9v21fSfItuEUlxIGo=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_cyberdelia_go_metrics_graphite",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cyberdelia/go-metrics-graphite",
        sum = "h1:M5QgkYacWj0Xs8MhpIK/5uwU02icXpEoSo9sM2aRCps=",
        version = "v0.0.0-20161219230853-39f87cc3b432",
    )
    go_repository(
        name = "com_github_cyphar_filepath_securejoin",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/cyphar/filepath-securejoin",
        sum = "h1:YX6ebbZCZP7VkM3scTTokDgBL2TY741X51MTk3ycuNI=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_d2g_dhcp4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/d2g/dhcp4",
        sum = "h1:Xo2rK1pzOm0jO6abTPIQwbAmqBIOj132otexc1mmzFc=",
        version = "v0.0.0-20170904100407-a1d1b6c41b1c",
    )
    go_repository(
        name = "com_github_d2g_dhcp4client",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/d2g/dhcp4client",
        sum = "h1:suYBsYZIkSlUMEz4TAYCczKf62IA2UWC+O8+KtdOhCo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_d2g_dhcp4server",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/d2g/dhcp4server",
        sum = "h1:+CpLbZIeUn94m02LdEKPcgErLJ347NUwxPKs5u8ieiY=",
        version = "v0.0.0-20181031114812-7d4a0a7f59a5",
    )
    go_repository(
        name = "com_github_d2g_hardwareaddr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/d2g/hardwareaddr",
        sum = "h1:itqmmf1PFpC4n5JW+j4BU7X4MTfVurhYRTjODoPb2Y8=",
        version = "v0.0.0-20190221164911-e7d9fbe030e4",
    )
    go_repository(
        name = "com_github_daaku_go_zipexe",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/daaku/go.zipexe",
        sum = "h1:wV4zMsDOI2SZ2m7Tdz1Ps96Zrx+TzaK15VbUaGozw0M=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_danwakefield_fnmatch",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/danwakefield/fnmatch",
        sum = "h1:y5HC9v93H5EPKqaS1UYVg1uYah5Xf51mBfIoWehClUQ=",
        version = "v0.0.0-20160403171240-cbb64ac3d964",
    )
    go_repository(
        name = "com_github_datadog_datadog_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/datadog-go",
        sum = "h1:o4QtYjBU/rG58VPh8Ne6F65YiMY5/v5q4WdY/HvRYMQ=",
        version = "v3.7.2+incompatible",
    )
    go_repository(
        name = "com_github_datadog_zstd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/zstd",
        sum = "h1:EndNeuB0l9syBZhut0wns3gV1hL8zX8LIu6ZiVHWLIQ=",
        version = "v1.4.5",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_daviddengcn_go_colortext",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/daviddengcn/go-colortext",
        sum = "h1:ANqDyC0ys6qCSvuEK7l3g5RaehL/Xck9EX8ATG8oKsE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_davidmz_go_pageant",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/davidmz/go-pageant",
        sum = "h1:bPblRCh5jGU+Uptpz6LgMZGD5hJoOt7otgT454WvHn0=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_dchest_siphash",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dchest/siphash",
        sum = "h1:9DFz8tQwl9pTVt5iok/9zKyzA1Q6bRGiF3HPiEEVr9I=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_dchest_uniuri",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dchest/uniuri",
        sum = "h1:74lLNRzvsdIlkTgfDSMuaPjBr4cf6k7pwQQANm/yLKU=",
        version = "v0.0.0-20160212164326-8902c56451e9",
    )
    go_repository(
        name = "com_github_deckarep_golang_set",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/deckarep/golang-set",
        sum = "h1:sk9/l/KqpunDwP7pSjUg0keiOOLEnOBHzykLrsPppp4=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_decker502_dnspod_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/decker502/dnspod-go",
        sum = "h1:6dwhUFCYbC5bgpebLKn7PrI43e/5mn9tpUL9YcYCdTU=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_decred_dcrd_dcrec_secp256k1_v4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/decred/dcrd/dcrec/secp256k1/v4",
        sum = "h1:YLtO71vCjJRCBcrPMtQ9nqBsqpA1m5sE92cU+pd5Mcc=",
        version = "v4.0.1",
    )
    go_repository(
        name = "com_github_deepmap_oapi_codegen",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/deepmap/oapi-codegen",
        sum = "h1:SegyeYGcdi0jLLrpbCMoJxnUUn8GBXHsvr4rbzjuhfU=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_github_denisbrodbeck_machineid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/denisbrodbeck/machineid",
        sum = "h1:geKr9qtkB876mXguW2X6TU4ZynleN6ezuMSRhl4D7AQ=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_dennwc_btrfs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dennwc/btrfs",
        sum = "h1:ue4Es4Xzz255hWQ7NAWzZxuXG+YOV7URzzusLLSe0zU=",
        version = "v0.0.0-20230312211831-a1f570bd01a1",
    )
    go_repository(
        name = "com_github_dennwc_ioctl",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dennwc/ioctl",
        sum = "h1:DsWAAjIxRqNcLn9x6mwfuf2pet3iB7aK90K4tF16rLg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_denverdino_aliyungo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/denverdino/aliyungo",
        sum = "h1:p6poVbjHDkKa+wtC8frBMwQtT3BmqGYBjzMwJ63tuR4=",
        version = "v0.0.0-20190125010748-a747050bb1ba",
    )
    go_repository(
        name = "com_github_devigned_tab",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/devigned/tab",
        sum = "h1:3mD6Kb1mUOYeLpJvTVSDwSg5ZsfSxfvxGRTxRsJsITA=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_dgraph_io_badger",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dgraph-io/badger",
        sum = "h1:mNw0qs90GVgGGWylh0umH5iag1j6n/PeJtNvL6KY/x8=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_github_dgraph_io_badger_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dgraph-io/badger/v2",
        sum = "h1:TRWBQg8UrlUhaFdco01nO2uXwzKS7zd+HVdwV/GHc4o=",
        version = "v2.2007.4",
    )
    go_repository(
        name = "com_github_dgraph_io_ristretto",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dgraph-io/ristretto",
        sum = "h1:KoJOtZf+6wpQaDTuOWGuo61GxcPBIfhwRxRTaTWGCTc=",
        version = "v0.0.4-0.20200906165740-41ebdbffecfd",
    )
    go_repository(
        name = "com_github_dgrijalva_jwt_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dgrijalva/jwt-go",
        sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_dgryski_go_farm",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dgryski/go-farm",
        sum = "h1:fAjc9m62+UWV/WAFKLNi6ZS0675eEUC9y3AlwSbQu1Y=",
        version = "v0.0.0-20200201041132-a6ae2369ad13",
    )
    go_repository(
        name = "com_github_dgryski_go_metro",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dgryski/go-metro",
        sum = "h1:BS21ZUJ/B5X2UVUbczfmdWH7GapPWAhxcMsDnjJTU1E=",
        version = "v0.0.0-20200812162917-85c65e2d0165",
    )
    go_repository(
        name = "com_github_dgryski_go_sip13",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dgryski/go-sip13",
        sum = "h1:RMLoZVzv4GliuWafOuPuQDKSm1SJph7uCRnnS61JAn4=",
        version = "v0.0.0-20181026042036-e10d5fee7954",
    )
    go_repository(
        name = "com_github_digitalocean_godo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/digitalocean/godo",
        sum = "h1:WYy7MIVVhTMZUNB+UA3irl2V9FyDJeDttsifYyn7jYA=",
        version = "v1.41.0",
    )
    go_repository(
        name = "com_github_dimchansky_utfbom",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dimchansky/utfbom",
        sum = "h1:vV6w1AhK4VMnhBno/TPVCoK9U/LP0PkLCS9tbxHdi/U=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_disintegration_imaging",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/disintegration/imaging",
        sum = "h1:w1LecBlG2Lnp8B3jk5zSuNqd7b4DXhcjwek1ei82L+c=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_github_distribution_reference",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/distribution/reference",
        sum = "h1:/FUIFXtfc/x2gpa5/VGfiGLuOIdYa1t65IKK2OFGvA0=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_djarvur_go_err113",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Djarvur/go-err113",
        sum = "h1:uCRZZOdMQ0TZPHYTdYpoC0bLYJKPEHPUJ8MeAa51lNU=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_dlclark_regexp2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dlclark/regexp2",
        sum = "h1:Izz0+t1Z5nI16/II7vuEo/nHjodOg0p7+OiDpjX5t1E=",
        version = "v1.4.1-0.20201116162257-a2a8dda75c91",
    )
    go_repository(
        name = "com_github_dnaeon_go_vcr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dnaeon/go-vcr",
        sum = "h1:r8L/HqC0Hje5AXMu1ooW8oyQyOFv4GxqpL0nRP7SLLY=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_dnsimple_dnsimple_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dnsimple/dnsimple-go",
        sum = "h1:N+q+ML1CZGf+5r4udu9Opy7WJNtOaFT9aM86Af9gLhk=",
        version = "v0.60.0",
    )
    go_repository(
        name = "com_github_docker_cli",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/cli",
        sum = "h1:OJ7YkwQA+k2Oi51lmCojpjiygKpi76P7bg91b2eJxYU=",
        version = "v20.10.9+incompatible",
    )
    go_repository(
        name = "com_github_docker_distribution",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/distribution",
        sum = "h1:Q50tZOPR6T/hjNsyc9g8/syEs6bk8XXApsHjKukMl68=",
        version = "v2.8.1+incompatible",
    )
    go_repository(
        name = "com_github_docker_docker",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/docker",
        sum = "h1:k5TYd5rIVQRSqcTwCID+cyVA0yRg86+Pcrz1ls0/frA=",
        version = "v25.0.1+incompatible",
    )
    go_repository(
        name = "com_github_docker_docker_credential_helpers",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/docker-credential-helpers",
        sum = "h1:zI2p9+1NQYdnG6sMU26EX4aVGlqbInSQxQXLvzJ4RPQ=",
        version = "v0.6.3",
    )
    go_repository(
        name = "com_github_docker_go_connections",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/go-connections",
        sum = "h1:El9xVISelRB7BuFusrZozjnkIM5YnzCViNKohAFqRJQ=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_docker_go_events",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/go-events",
        sum = "h1:+pKlWGMw7gf6bQ+oDZB4KHQFypsfjYlq/C4rfL7D3g8=",
        version = "v0.0.0-20190806004212-e31b211e4f1c",
    )
    go_repository(
        name = "com_github_docker_go_metrics",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/go-metrics",
        sum = "h1:AgB/0SvBxihN0X8OR4SjsblXkbMvalQ8cjmtKQ2rQV8=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_docker_go_units",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/go-units",
        sum = "h1:3uh0PgVws3nIA0Q+MwDC8yjEPf9zjRfZZWXZYDct3Tw=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_docker_libtrust",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/libtrust",
        sum = "h1:ZClxb8laGDf5arXfYcAtECDFgAgHklGI8CxgjHnXKJ4=",
        version = "v0.0.0-20150114040149-fa567046d9b1",
    )
    go_repository(
        name = "com_github_docker_spdystream",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/spdystream",
        sum = "h1:cenwrSVm+Z7QLSV/BsnenAOcDXdX4cMv4wP0B/5QbPg=",
        version = "v0.0.0-20160310174837-449fdfce4d96",
    )
    go_repository(
        name = "com_github_docopt_docopt_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/docopt/docopt-go",
        sum = "h1:bWDMxwH3px2JBh6AyO7hdCn/PkvCZXii8TGj7sbtEbQ=",
        version = "v0.0.0-20180111231733-ee0de3bc6815",
    )
    go_repository(
        name = "com_github_dop251_goja",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dop251/goja",
        sum = "h1:Yt+4K30SdjOkRoRRm3vYNQgR+/ZIy0RmeUDZo7Y8zeQ=",
        version = "v0.0.0-20220405120441-9037c2b61cbf",
    )
    go_repository(
        name = "com_github_dreamacro_clash",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Dreamacro/clash",
        sum = "h1:ZQe/7G+JclA1vvyAn8MtaEBvQK73mWR6lV3BceDaJoY=",
        version = "v1.11.4",
    )
    go_repository(
        name = "com_github_dsnet_compress",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dsnet/compress",
        sum = "h1:PlZu0n3Tuv04TzpfPbrnI0HW/YwodEXDS+oPKahKF0Q=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_dsnet_golib",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dsnet/golib",
        sum = "h1:tFh1tRc4CA31yP6qDcu+Trax5wW5GuMxvkIba07qVLY=",
        version = "v0.0.0-20171103203638-1ea166775780",
    )
    go_repository(
        name = "com_github_dustin_go_humanize",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/dustin/go-humanize",
        sum = "h1:GzkhY7T5VNhEkwH0PVJgjz+fX1rhBrR7pRT3mDkpeCY=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_eapache_go_resiliency",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/eapache/go-resiliency",
        sum = "h1:v7g92e/KSN71Rq7vSThKaWIq68fL4YHvWyiUKorFR1Q=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_eapache_go_xerial_snappy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/eapache/go-xerial-snappy",
        sum = "h1:YEetp8/yCZMuEPMUDHG0CW/brkkEp8mzqk2+ODEitlw=",
        version = "v0.0.0-20180814174437-776d5712da21",
    )
    go_repository(
        name = "com_github_eapache_queue",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/eapache/queue",
        sum = "h1:YOEu7KNc61ntiQlcEeUIoDTJ2o8mQznoNvUhiigpIqc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_ebfe_bcrypt_pbkdf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ebfe/bcrypt_pbkdf",
        sum = "h1:YtdtTUN1iH97s+6PUjLnaiKSQj4oG1/EZ3N9bx6g4kU=",
        version = "v0.0.0-20140212075826-3c8d2dcb253a",
    )
    go_repository(
        name = "com_github_ebi_yade_altsvc_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ebi-yade/altsvc-go",
        sum = "h1:HmZDNb5ZOPlkyXhi34LnRckawFCux7yPYw+dtInIixo=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_edsrzf_mmap_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/edsrzf/mmap-go",
        sum = "h1:CEBF7HpRnUCSJgGUb5h1Gm7e3VkmVDrR8lvWVLtrOFw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_elazarl_goproxy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/elazarl/goproxy",
        sum = "h1:pEtiCjIXx3RvGjlUJuCNxNOw0MNblyR9Wi+vJGBFh+8=",
        version = "v0.0.0-20191011121108-aa519ddbe484",
    )
    go_repository(
        name = "com_github_elithrar_simple_scrypt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/elithrar/simple-scrypt",
        sum = "h1:KIlOlxdoQf9JWKl5lMAJ28SY2URB0XTRDn2TckyzAZg=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_ema_qdisc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ema/qdisc",
        sum = "h1:EHLG08FVRbWLg8uRICa3xzC9Zm0m7HyMHfXobWFnXYg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_emicklei_go_restful",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/emicklei/go-restful",
        sum = "h1:spTtZBk5DYEvbxMVutUuTyh1Ao2r4iyvLdACqsl/Ljk=",
        version = "v2.9.5+incompatible",
    )
    go_repository(
        name = "com_github_emicklei_go_restful_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/emicklei/go-restful/v3",
        sum = "h1:rAQeMHw1c7zTmncogyy8VvRZwtkmkZ4FxERmMY4rD+g=",
        version = "v3.11.0",
    )
    go_repository(
        name = "com_github_emirpasic_gods",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/emirpasic/gods",
        sum = "h1:QAUIPSaCu4G+POclxeqb3F+WPpdKqFGlw36+yOzGlrg=",
        version = "v1.12.0",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:wSUXTlLfiAQRWs2F+p+EKOY9rUyis1MyGqJ2DIk5HpM=",
        version = "v0.11.1",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:QkIBuU5k+x7/QXPvPPnWXWlCdaBFApVqftFV6k087DA=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_etcd_io_gofail",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/etcd-io/gofail",
        sum = "h1:Y2I0lxOttdUKz+hNaIdG3FtjuQrTmwXun1opRV65IZc=",
        version = "v0.0.0-20190801230047-ad7f989257ca",
    )
    go_repository(
        name = "com_github_ethereum_go_ethereum",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ethereum/go-ethereum",
        sum = "h1:75IW830ClSS40yrQC1ZCMZCt5I+zU16oqId2SiQwdQ4=",
        version = "v1.10.20",
    )
    go_repository(
        name = "com_github_euank_go_kmsg_parser",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/euank/go-kmsg-parser",
        sum = "h1:cHD53+PLQuuQyLZeriD1V/esuG4MuU0Pjs5y6iknohY=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_evanphx_json_patch",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/evanphx/json-patch",
        sum = "h1:4onqiflcdA9EOZ4RxV643DvftH5pOlLGNtQ5lPWQu84=",
        version = "v4.12.0+incompatible",
    )
    go_repository(
        name = "com_github_evanphx_json_patch_v5",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/evanphx/json-patch/v5",
        sum = "h1:bAmFiUJ+o0o2B4OiTFeE3MqCOtyo+jjPP9iZ0VRxYUc=",
        version = "v5.5.0",
    )
    go_repository(
        name = "com_github_exoscale_egoscale",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/exoscale/egoscale",
        sum = "h1:1FNZVk8jHUx0AvWhOZxLEDNlacTU0chMXUUNkm9EZaI=",
        version = "v0.18.1",
    )
    go_repository(
        name = "com_github_exponent_io_jsonpath",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/exponent-io/jsonpath",
        sum = "h1:105gxyaGwCFad8crR9dcMQWvV9Hvulu6hwUh4tWPJnM=",
        version = "v0.0.0-20151013193312-d6023ce2651d",
    )
    go_repository(
        name = "com_github_fanliao_go_promise",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fanliao/go-promise",
        sum = "h1:0eU/faU2oDIB2BkQVM02hgRLJjGzzUuRf19HUhp0394=",
        version = "v0.0.0-20141029170127-1890db352a72",
    )
    go_repository(
        name = "com_github_fasthttp_contrib_websocket",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fasthttp-contrib/websocket",
        sum = "h1:DddqAaWDpywytcG8w/qoQ5sAN8X12d3Z3koB0C3Rxsc=",
        version = "v0.0.0-20160511215533-1f3b11f56072",
    )
    go_repository(
        name = "com_github_fatih_camelcase",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fatih/camelcase",
        sum = "h1:hxNvNX/xYBp0ovncs8WyWZrOrpBNub/JfaMvbURyft8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_fatih_color",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fatih/color",
        sum = "h1:8LOYc1KYPPmyKMuN8QV2DNRWNbLo6LZ0iLs8+mlH53w=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_github_fatih_structs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fatih/structs",
        sum = "h1:Q7juDM0QtcnhCpeyLGQKyg4TOIghuNXrkL32pHAUMxo=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_felixge_httpsnoop",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/felixge/httpsnoop",
        sum = "h1:s/nj+GCswXYzN5v2DpNMuMQYe+0DDwt5WVCU6CWBdXk=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_filebrowser_filebrowser_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/filebrowser/filebrowser/v2",
        sum = "h1:VnaNNDFI7lNzceTDRuusgsiVhwdqmark30ySAly1zuo=",
        version = "v2.11.0",
    )
    go_repository(
        name = "com_github_fjl_gencodec",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fjl/gencodec",
        sum = "h1:CndMRAH4JIwxbW8KYq6Q+cGWcGHz0FjGR3QqcInWcW0=",
        version = "v0.0.0-20220412091415-8bb9e558978c",
    )
    go_repository(
        name = "com_github_fjl_memsize",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fjl/memsize",
        sum = "h1:FtmdgXiUlNeRsoNMFlKLDt+S+6hbjVMEW6RGQ7aUf7c=",
        version = "v0.0.0-20190710130421-bcb5799ab5e5",
    )
    go_repository(
        name = "com_github_flynn_archive_go_shlex",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/flynn-archive/go-shlex",
        sum = "h1:BMXYYRWTLOJKlh+lOBt6nUQgXAfB7oVIQt5cNreqSLI=",
        version = "v0.0.0-20150515145356-3f9db97f8568",
    )
    go_repository(
        name = "com_github_flynn_go_shlex",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/flynn/go-shlex",
        sum = "h1:BHsljHzVlRcyQhjrss6TZTdY2VfCqZPbv5k3iBFa2ZQ=",
        version = "v0.0.0-20150515145356-3f9db97f8568",
    )
    go_repository(
        name = "com_github_flynn_noise",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/flynn/noise",
        sum = "h1:DlTHqmzmvcEiKj+4RYo/imoswx/4r6iBlCMfVtrMXpQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_form3tech_oss_jwt_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/form3tech-oss/jwt-go",
        sum = "h1:7ZaBxOI7TMoYBfyA3cQHErNNyAWIKUMIwqxEtgHOs5c=",
        version = "v3.2.3+incompatible",
    )
    go_repository(
        name = "com_github_fortytw2_leaktest",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fortytw2/leaktest",
        sum = "h1:u8491cBMTQ8ft8aeV+adlcytMZylmA5nnwwkRZjI8vw=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_francoispqt_gojay",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/francoispqt/gojay",
        sum = "h1:d2m3sFjloqoIUQU3TsHBgj6qg/BVGlTBeHDUmyJnXKk=",
        version = "v1.2.13",
    )
    go_repository(
        name = "com_github_franela_goblin",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/franela/goblin",
        sum = "h1:gb2Z18BhTPJPpLQWj4T+rfKHYCHxRHCtRxhKKjRidVw=",
        version = "v0.0.0-20200105215937-c9ffbefa60db",
    )
    go_repository(
        name = "com_github_franela_goreq",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/franela/goreq",
        sum = "h1:a9ENSRDFBUPkJ5lCgVZh26+ZbGyoVJG7yb5SSzF5H54=",
        version = "v0.0.0-20171204163338-bcd34c9993f8",
    )
    go_repository(
        name = "com_github_frankban_quicktest",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/frankban/quicktest",
        sum = "h1:dfYrrRyLtiqT9GyKXgdh+k4inNeTvmGbuSgZ3lx3GhA=",
        version = "v1.14.5",
    )
    go_repository(
        name = "com_github_freman_caddy2_reauth",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/freman/caddy2-reauth",
        sum = "h1:JBn9TEuc5i0FPckHVOL8NI8Q+R0kNBOBLvBLadgkZ3o=",
        version = "v0.0.0-20200518130136-6064aa96b1a8",
    )
    go_repository(
        name = "com_github_fsnotify_fsnotify",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fsnotify/fsnotify",
        sum = "h1:8JEhPFa5W2WU7YfeZzPNqzMP6Lwt7L2715Ggo0nosvA=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_fullsailor_pkcs7",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fullsailor/pkcs7",
        sum = "h1:RDBNVkRviHZtvDvId8XSGPu3rmpmSe+wKRcEWNgsfWU=",
        version = "v0.0.0-20190404230743-d7302db945fa",
    )
    go_repository(
        name = "com_github_fullstorydev_grpcurl",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fullstorydev/grpcurl",
        sum = "h1:Pp648wlTTg3OKySeqxM5pzh8XF6vLqrm8wRq66+5Xo0=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_fvbommel_sortorder",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fvbommel/sortorder",
        sum = "h1:fUmoe+HLsBTctBDoaBwpQo5N+nrCp8g/BjKb/6ZQmYw=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_fxamacker_cbor_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/fxamacker/cbor/v2",
        sum = "h1:aM45YGMctNakddNNAezPxDUpv38j44Abh+hifNuqXik=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_g07cha_defender",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/g07cha/defender",
        sum = "h1:gWvniJ4GbFfkf700kykAImbLiEMU0Q3QN9hQ26Js1pU=",
        version = "v0.0.0-20180505193036-5665c627c814",
    )
    go_repository(
        name = "com_github_garslo_gogen",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/garslo/gogen",
        sum = "h1:IZqZOB2fydHte3kUgxrzK5E1fW7RQGeDwE8F/ZZnUYc=",
        version = "v0.0.0-20170306192744-1d203ffc1f61",
    )
    go_repository(
        name = "com_github_garyburd_redigo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/garyburd/redigo",
        sum = "h1:yE/pwKCrbLpLpQICzYTeZ7JsTA/C53wFTJHaEtRqniM=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_github_gavv_httpexpect",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gavv/httpexpect",
        sum = "h1:1X9kcRshkSKEjNJJxX9Y9mQ5BRfbxU5kORdjhlA1yX8=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_gballet_go_libpcsclite",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gballet/go-libpcsclite",
        sum = "h1:tY80oXqGNY4FhTFhk+o9oFHGINQ/+vhlm8HFzi6znCI=",
        version = "v0.0.0-20190607065134-2772fd86a8ff",
    )
    go_repository(
        name = "com_github_gdamore_encoding",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gdamore/encoding",
        sum = "h1:+7OoQ1Bc6eTm5niUzBa0Ctsh6JbMW6Ra+YNuAtDBdko=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gdamore_tcell",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gdamore/tcell",
        sum = "h1:vUnHwJRvcPQa3tzi+0QI4U9JINXYJlOz9yiaiPQ2wMU=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_geertjohan_go_incremental",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/GeertJohan/go.incremental",
        sum = "h1:7AH+pY1XUgQE4Y1HcXYaMqAI0m9yrFqo/jt0CW30vsg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_geertjohan_go_rice",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/GeertJohan/go.rice",
        sum = "h1:KkI6O9uMaQU3VEKaj01ulavtF7o1fWT7+pk/4voiMLQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_getsentry_raven_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/getsentry/raven-go",
        sum = "h1:no+xWJRb5ZI7eE8TWgIq1jLulQiIoLG0IfYxv5JYMGs=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_ghodss_yaml",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ghodss/yaml",
        sum = "h1:Mn26/9ZMNWSw9C9ERFA1PUxfmGpolnw2v0bKOREu5ew=",
        version = "v1.0.1-0.20190212211648-25d852aebe32",
    )
    go_repository(
        name = "com_github_gliderlabs_ssh",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gliderlabs/ssh",
        sum = "h1:6zsha5zo/TWhRhwqCD3+EarCAgZ2yN28ipRnGPnwkI0=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_go_acme_lego",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-acme/lego",
        sum = "h1:5fNN9yRQfv8ymH3DSsxla+4aYeQt2IgfZqHKVnK8f0s=",
        version = "v2.5.0+incompatible",
    )
    go_repository(
        name = "com_github_go_acme_lego_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-acme/lego/v3",
        sum = "h1:qC5/8/CbltyAE8fGLE6bGlqucj7pXc/vBxiLwLOsmAQ=",
        version = "v3.7.0",
    )
    go_repository(
        name = "com_github_go_asn1_ber_asn1_ber",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-asn1-ber/asn1-ber",
        sum = "h1:gvPdv/Hr++TRFCl0UbPFHC54P9N9jgsRPnmnr419Uck=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_go_chi_chi",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chi/chi",
        sum = "h1:fGFk2Gmi/YKXk0OmGfBh0WgmN3XB8lVnEyNz34tQRec=",
        version = "v4.1.2+incompatible",
    )
    go_repository(
        name = "com_github_go_chi_chi_v5",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chi/chi/v5",
        sum = "h1:rDTPXLDHGATaeHvVlLcR4Qe0zftYethFucbjVQ1PxU8=",
        version = "v5.0.7",
    )
    go_repository(
        name = "com_github_go_chi_cors",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chi/cors",
        sum = "h1:xEC8UT3Rlp2QuWNEr4Fs/c2EAGVKBwy/1vHx3bppil4=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_go_chi_render",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chi/render",
        sum = "h1:4/5tis2cKaNdnv9zFLfXzcquC9HbeZgCnxGnKrltBS8=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_go_cmd_cmd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-cmd/cmd",
        sum = "h1:IK23uTRWxq6UJnNWp8nKO7mVCwnPfbaxA2lhzEKfNj0=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_go_critic_go_critic",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-critic/go-critic",
        sum = "h1:sGEEdiuvLV0OC7/yC6MnK3K6LCPBplspK45B0XVdFAc=",
        version = "v0.4.3",
    )
    go_repository(
        name = "com_github_go_errors_errors",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-errors/errors",
        sum = "h1:J6MZopCL4uSllY1OfXM374weqZFFItUbrImctkmUxIA=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_go_fed_httpsig",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-fed/httpsig",
        sum = "h1:9M+hb0jkEICD8/cAiNqEB66R87tTINszBRTjwjQzWcI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_go_git_gcfg",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-git/gcfg",
        sum = "h1:Q5ViNfGF8zFgyJWPqYwA7qGFoMTEiBmdlkcfRmpIMa4=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_go_git_go_billy_v5",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-git/go-billy/v5",
        sum = "h1:4pl5BV4o7ZG/lterP4S6WzJ6xr49Ba5ET9ygheTYahk=",
        version = "v5.1.0",
    )
    go_repository(
        name = "com_github_go_git_go_git_v5",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-git/go-git/v5",
        sum = "h1:8WKMtJR2j8RntEXR/uvTKagfEt4GYlwQ7mntE4+0GWc=",
        version = "v5.3.0",
    )
    go_repository(
        name = "com_github_go_gl_glfw",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-gl/glfw",
        sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
        version = "v0.0.0-20190409004039-e6da0acd62b1",
    )
    go_repository(
        name = "com_github_go_gl_glfw_v3_3_glfw",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-gl/glfw/v3.3/glfw",
        sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
        version = "v0.0.0-20200222043503-6f7a984d4dc4",
    )
    go_repository(
        name = "com_github_go_gost_gosocks4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-gost/gosocks4",
        sum = "h1:+k1sec8HlELuQV7rWftIkmy8UijzUt2I6t+iMPlGB2s=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_go_gost_gosocks5",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-gost/gosocks5",
        sum = "h1:Hkmp9YDRBSCJd7xywW6dBPT6B9aQTkuWd+3WCheJiJA=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_go_gost_relay",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-gost/relay",
        sum = "h1:itaaJhQJ19kUXEB4Igb0EbY8m+1Py2AaNNSBds/9gk4=",
        version = "v0.1.1-0.20211123134818-8ef7fd81ffd7",
    )
    go_repository(
        name = "com_github_go_gost_tls_dissector",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-gost/tls-dissector",
        sum = "h1:xj8gUZGYO3nb5+6Bjw9+tsFkA9sYynrOvDvvC4uDV2I=",
        version = "v0.0.2-0.20220408131628-aac992c27451",
    )
    go_repository(
        name = "com_github_go_ini_ini",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-ini/ini",
        sum = "h1:8+SRbfpRFlIunpSum4BEf1ClTtVjOgKzgBv9pHFkI6w=",
        version = "v1.44.0",
    )
    go_repository(
        name = "com_github_go_kit_kit",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kit/kit",
        sum = "h1:dXFJfIHVvUcpSgDOV+Ne6t7jXri8Tfv2uOLHUZ2XNuo=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_go_kit_log",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kit/log",
        sum = "h1:MRVx0/zhvdseW+Gza6N9rVzU/IVzaeE1SFI4raAhmBU=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_go_ldap_ldap_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-ldap/ldap/v3",
        sum = "h1:7WsKqasmPThNvdl0Q5GPpbTDD/ZD98CfuawrMIuh7qQ=",
        version = "v3.1.10",
    )
    go_repository(
        name = "com_github_go_lintpack_lintpack",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-lintpack/lintpack",
        sum = "h1:DI5mA3+eKdWeJ40nU4d6Wc26qmdG8RCi/btYq0TuRN0=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_go_log_log",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-log/log",
        sum = "h1:z8i91GBudxD5L3RmF0KVpetCbcGWAV7q1Tw1eRwQM9Q=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_go_logfmt_logfmt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-logfmt/logfmt",
        sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-logr/logr",
        sum = "h1:2y3SDp0ZXuc6/cjLSZ+Q3ir+QB9T/iG5yYRXqsagWSY=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_go_logr_stdr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-logr/stdr",
        sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_go_logr_zapr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-logr/zapr",
        sum = "h1:a9vnzlIBPQBBkeaR9IuMUfmVOrQlkoC4YfPoFkX3T7A=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_github_go_ole_go_ole",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-ole/go-ole",
        sum = "h1:/Fpf6oFPoeFik9ty7siob0G6Ke8QvQEuVcuChpwXzpY=",
        version = "v1.2.6",
    )
    go_repository(
        name = "com_github_go_openapi_jsonpointer",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/jsonpointer",
        sum = "h1:eCs3fxoIi3Wh6vtgmLTOjdhSpiqphQ+DaPn38N2ZdrE=",
        version = "v0.19.6",
    )
    go_repository(
        name = "com_github_go_openapi_jsonreference",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/jsonreference",
        sum = "h1:3sVjiK66+uXK/6oQ8xgcRKcFgQ5KXa2KvnJRumpMGbE=",
        version = "v0.20.2",
    )
    go_repository(
        name = "com_github_go_openapi_spec",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/spec",
        sum = "h1:0XRyw8kguri6Yw4SxhsQA/atC88yqrk0+G4YhI2wabc=",
        version = "v0.19.3",
    )
    go_repository(
        name = "com_github_go_openapi_swag",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/swag",
        sum = "h1:yMBqmnQ0gyZvEb/+KzuWZOXgllrXT4SADYbvDaXHv/g=",
        version = "v0.22.3",
    )
    go_repository(
        name = "com_github_go_piv_piv_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-piv/piv-go",
        sum = "h1:rfjdFdASfGV5KLJhSjgpGJ5lzVZVtRWn8ovy/H9HQ/U=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_go_redis_redis",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-redis/redis",
        sum = "h1:K0pv1D7EQUjfyoMql+r/jZqCLizCGKFlFgcHWWmHQjg=",
        version = "v6.15.9+incompatible",
    )
    go_repository(
        name = "com_github_go_sourcemap_sourcemap",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-sourcemap/sourcemap",
        sum = "h1:W1iEw64niKVGogNgBN3ePyLFfuisuzeidWPMPWmECqU=",
        version = "v2.1.3+incompatible",
    )
    go_repository(
        name = "com_github_go_sql_driver_mysql",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-sql-driver/mysql",
        sum = "h1:BCTh4TKNUYmOmMUcQ3IipzF5prigylS7XXjEkfCHuOE=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_go_stack_stack",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-stack/stack",
        sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_go_task_slim_sprig",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-task/slim-sprig",
        sum = "h1:tfuBGBXKqDEevZMzYi5KSi8KkcZtzBcTgAUUtapy0OI=",
        version = "v0.0.0-20230315185526-52ccab3ef572",
    )
    go_repository(
        name = "com_github_go_test_deep",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-test/deep",
        sum = "h1:onZX1rnHT3Wv6cqNgYyFOOlgVKJrksuCMCRvJStbMYw=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_go_toolsmith_astcast",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-toolsmith/astcast",
        sum = "h1:JojxlmI6STnFVG9yOImLeGREv8W2ocNUM+iOhR6jE7g=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_go_toolsmith_astcopy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-toolsmith/astcopy",
        sum = "h1:OMgl1b1MEpjFQ1m5ztEO06rz5CUd3oBv9RF7+DyvdG8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_go_toolsmith_astequal",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-toolsmith/astequal",
        sum = "h1:4zxD8j3JRFNyLN46lodQuqz3xdKSrur7U/sr0SDS/gQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_go_toolsmith_astfmt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-toolsmith/astfmt",
        sum = "h1:A0vDDXt+vsvLEdbMFJAUBI/uTbRw1ffOPnxsILnFL6k=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_go_toolsmith_astinfo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-toolsmith/astinfo",
        sum = "h1:wP6mXeB2V/d1P1K7bZ5vDUO3YqEzcvOREOxZPEu3gVI=",
        version = "v0.0.0-20180906194353-9809ff7efb21",
    )
    go_repository(
        name = "com_github_go_toolsmith_astp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-toolsmith/astp",
        sum = "h1:alXE75TXgcmupDsMK1fRAy0YUzLzqPVvBKoyWV+KPXg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_go_toolsmith_pkgload",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-toolsmith/pkgload",
        sum = "h1:4DFWWMXVfbcN5So1sBNW9+yeiMqLFGl1wFLTL5R0Tgg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_go_toolsmith_strparse",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-toolsmith/strparse",
        sum = "h1:Vcw78DnpCAKlM20kSbAyO4mPfJn/lyYA4BJUDxe2Jb4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_go_toolsmith_typep",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-toolsmith/typep",
        sum = "h1:8xdsa1+FSIH/RhEkgnD1j2CJOy5mNllW1Q9tRiYwvlk=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_go_xmlfmt_xmlfmt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-xmlfmt/xmlfmt",
        sum = "h1:khEcpUM4yFcxg4/FHQWkvVRmgijNXRfzkIDHh23ggEo=",
        version = "v0.0.0-20191208150333-d5b6f63a941b",
    )
    go_repository(
        name = "com_github_gobwas_glob",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobwas/glob",
        sum = "h1:A4xDbljILXROh+kObIiy5kIaPYD8e96x1tgBhUI5J+Y=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_gobwas_httphead",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobwas/httphead",
        sum = "h1:exrUm0f4YX0L7EBwZHuCF4GDp8aJfVeBrlLQrs6NqWU=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gobwas_pool",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobwas/pool",
        sum = "h1:xfeeEhW7pwmX8nuLVlqbzVc7udMDrwetjEv+TZIz1og=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_gobwas_ws",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobwas/ws",
        sum = "h1:F2aeBZrm2NDsc7vbovKrWSogd4wvfAxg0FQ89/iqOTk=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_godbus_dbus",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/godbus/dbus",
        sum = "h1:BWhy2j3IXJhjCbC68FptL43tDKIq8FladmaTs3Xs7Z8=",
        version = "v0.0.0-20190422162347-ade71ed3457e",
    )
    go_repository(
        name = "com_github_godbus_dbus_v5",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/godbus/dbus/v5",
        sum = "h1:4KLkAxT3aOY8Li4FRJe/KvhoNFFxo0m6fNuFUO8QJUk=",
        version = "v5.1.0",
    )
    go_repository(
        name = "com_github_gofrs_flock",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gofrs/flock",
        sum = "h1:ekuhfTjngPhisSjOJ0QWKpPQE8/rbknHaes6WVJj5Hw=",
        version = "v0.0.0-20190320160742-5135e617513b",
    )
    go_repository(
        name = "com_github_gofrs_uuid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gofrs/uuid",
        sum = "h1:yyYWMnhkhrKwwr8gAOcOCYxOOscHgDS9yZgBrnJfGa0=",
        version = "v4.2.0+incompatible",
    )
    go_repository(
        name = "com_github_gogo_googleapis",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gogo/googleapis",
        sum = "h1:zgVt4UpGxcqVOw97aRGxT4svlcmdK35fynLNctY32zI=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_goji_httpauth",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/goji/httpauth",
        sum = "h1:lBXNCxVENCipq4D1Is42JVOP4eQjlB8TQ6H69Yx5J9Q=",
        version = "v0.0.0-20160601135302-2da839ab0f4d",
    )
    go_repository(
        name = "com_github_golang_glog",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/glog",
        sum = "h1:/d3pCKDPWNnvIWe0vVUpNP32qc8U3PDVxySP/y360qE=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/groupcache",
        sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
        version = "v0.0.0-20210331224755-41bb18bfe9da",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt_v4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang-jwt/jwt/v4",
        sum = "h1:7cYmW1XlMY7h7ii7UhUyChSgS5wUJEnm9uZVTGqOWzg=",
        version = "v4.5.0",
    )
    go_repository(
        name = "com_github_golang_lint",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/lint",
        sum = "h1:2hRPrmiwPrp3fQX967rNJIhQPtiGXdlQWAxKbKw3VHA=",
        version = "v0.0.0-20180702182130-06c8688daad7",
    )
    go_repository(
        name = "com_github_golang_mock",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/mock",
        sum = "h1:ErTB+efbowRARo13NNdxyJji2egdxLGQhRaY+DUumQc=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/protobuf",
        sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_github_golang_snappy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/snappy",
        sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_golangci_check",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/check",
        sum = "h1:23T5iq8rbUYlhpt5DB4XJkc6BU31uODLD1o1gKvZmD0=",
        version = "v0.0.0-20180506172741-cfe4005ccda2",
    )
    go_repository(
        name = "com_github_golangci_dupl",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/dupl",
        sum = "h1:w8hkcTqaFpzKqonE9uMCefW1WDie15eSP/4MssdenaM=",
        version = "v0.0.0-20180902072040-3e9179ac440a",
    )
    go_repository(
        name = "com_github_golangci_errcheck",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/errcheck",
        sum = "h1:YYWNAGTKWhKpcLLt7aSj/odlKrSrelQwlovBpDuf19w=",
        version = "v0.0.0-20181223084120-ef45e06d44b6",
    )
    go_repository(
        name = "com_github_golangci_go_misc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/go-misc",
        sum = "h1:9kfjN3AdxcbsZBf8NjltjWihK2QfBBBZuv91cMFfDHw=",
        version = "v0.0.0-20180628070357-927a3d87b613",
    )
    go_repository(
        name = "com_github_golangci_go_tools",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/go-tools",
        sum = "h1:/7detzz5stiXWPzkTlPTzkBEIIE4WGpppBJYjKqBiPI=",
        version = "v0.0.0-20190318055746-e32c54105b7c",
    )
    go_repository(
        name = "com_github_golangci_goconst",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/goconst",
        sum = "h1:pe9JHs3cHHDQgOFXJJdYkK6fLz2PWyYtP4hthoCMvs8=",
        version = "v0.0.0-20180610141641-041c5f2b40f3",
    )
    go_repository(
        name = "com_github_golangci_gocyclo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/gocyclo",
        sum = "h1:pXTK/gkVNs7Zyy7WKgLXmpQ5bHTrq5GDsp8R9Qs67g0=",
        version = "v0.0.0-20180528144436-0a533e8fa43d",
    )
    go_repository(
        name = "com_github_golangci_gofmt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/gofmt",
        sum = "h1:iR3fYXUjHCR97qWS8ch1y9zPNsgXThGwjKPrYfqMPks=",
        version = "v0.0.0-20190930125516-244bba706f1a",
    )
    go_repository(
        name = "com_github_golangci_golangci_lint",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/golangci-lint",
        sum = "h1:VYLx63qb+XJsHdZ27PMS2w5JZacN0XG8ffUwe7yQomo=",
        version = "v1.27.0",
    )
    go_repository(
        name = "com_github_golangci_gosec",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/gosec",
        sum = "h1:fUdgm/BdKvwOHxg5AhNbkNRp2mSy8sxTXyBVs/laQHo=",
        version = "v0.0.0-20190211064107-66fb7fc33547",
    )
    go_repository(
        name = "com_github_golangci_ineffassign",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/ineffassign",
        sum = "h1:gLLhTLMk2/SutryVJ6D4VZCU3CUqr8YloG7FPIBWFpI=",
        version = "v0.0.0-20190609212857-42439a7714cc",
    )
    go_repository(
        name = "com_github_golangci_lint_1",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/lint-1",
        sum = "h1:MfyDlzVjl1hoaPzPD4Gpb/QgoRfSBR0jdhwGyAWwMSA=",
        version = "v0.0.0-20191013205115-297bf364a8e0",
    )
    go_repository(
        name = "com_github_golangci_maligned",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/maligned",
        sum = "h1:kNY3/svz5T29MYHubXix4aDDuE3RWHkPvopM/EDv/MA=",
        version = "v0.0.0-20180506175553-b1d89398deca",
    )
    go_repository(
        name = "com_github_golangci_misspell",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/misspell",
        sum = "h1:pLzmVdl3VxTOncgzHcvLOKirdvcx/TydsClUQXTehjo=",
        version = "v0.3.5",
    )
    go_repository(
        name = "com_github_golangci_prealloc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/prealloc",
        sum = "h1:leSNB7iYzLYSSx3J/s5sVf4Drkc68W2wm4Ixh/mr0us=",
        version = "v0.0.0-20180630174525-215b22d4de21",
    )
    go_repository(
        name = "com_github_golangci_revgrep",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/revgrep",
        sum = "h1:XQKc8IYQOeRwVs36tDrEmTgDgP88d5iEURwpmtiAlOM=",
        version = "v0.0.0-20180812185044-276a5c0a1039",
    )
    go_repository(
        name = "com_github_golangci_unconvert",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangci/unconvert",
        sum = "h1:zwtduBRr5SSWhqsYNgcuWO2kFlpdOZbP0+yRjmvPGys=",
        version = "v0.0.0-20180507085042-28b1c447d1f4",
    )
    go_repository(
        name = "com_github_golangplus_bytes",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangplus/bytes",
        sum = "h1:YQKBijBVMsBxIiXT4IEhlKR2zHohjEqPole4umyDX+c=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_golangplus_fmt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangplus/fmt",
        sum = "h1:FnUKtw86lXIPfBMc3FimNF3+ABcV+aH5F17OOitTN+E=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_golangplus_testing",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/golangplus/testing",
        sum = "h1:+ZeeiKZENNOMkTTELoSySazi+XaEhVO0mb+eanrSEUQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_google_btree",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/btree",
        sum = "h1:gK4Kx5IaGY9CD5sPJ36FHiBJ6ZXl0kilRiiCj+jdYp4=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_google_cadvisor",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/cadvisor",
        sum = "h1:jai6dmBP9QAYluNGqU18yVUTw6uuyAW0AqtZIjvl8Qg=",
        version = "v0.39.0",
    )
    go_repository(
        name = "com_github_google_cel_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/cel-go",
        sum = "h1:6ebJFzu1xO2n7TLtN+UBqShGBhlD85bhvglh5DpcfqQ=",
        version = "v0.17.7",
    )
    go_repository(
        name = "com_github_google_cel_spec",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/cel-spec",
        sum = "h1:hWEzw+1L1UNxfHAbKXYbirsPGlG8ArXNcTnBKvBqRJ0=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_google_certificate_transparency_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/certificate-transparency-go",
        sum = "h1:BwvtrCU6nQ3cw0bAbwClSGFQvZpY2S5yt8amttqFSnQ=",
        version = "v1.1.2-0.20210623111010-a50f74f4ce95",
    )
    go_repository(
        name = "com_github_google_gnostic_models",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/gnostic-models",
        sum = "h1:yo/ABAfM5IMRsS1VnXjTBvUb61tFIHozhlYvRgGre9I=",
        version = "v0.6.8",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-cmp",
        sum = "h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_google_go_containerregistry",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-containerregistry",
        sum = "h1:/+mFTs4AlwsJ/mJe8NDtKb7BxLtbZFpcn8vDsneEkwQ=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_google_go_github",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-github",
        sum = "h1:N0LgJ1j65A7kfXrZnUDaYCs/Sf4rEjNlfyDHW9dolSY=",
        version = "v17.0.0+incompatible",
    )
    go_repository(
        name = "com_github_google_go_github_v28",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-github/v28",
        sum = "h1:kORf5ekX5qwXO2mGzXXOjMe/g6ap8ahVe0sBEulhSxo=",
        version = "v28.1.1",
    )
    go_repository(
        name = "com_github_google_go_github_v30",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-github/v30",
        sum = "h1:VLDx+UolQICEOKu2m4uAoMti1SxuEBAl7RSEG16L+Oo=",
        version = "v30.1.0",
    )
    go_repository(
        name = "com_github_google_go_licenses",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-licenses",
        sum = "h1:JtmsUf+m+KdwCOgLG578T0Mvd0+l+dezPrJh5KYnXZg=",
        version = "v0.0.0-20210329231322-ce1d9163b77d",
    )
    go_repository(
        name = "com_github_google_go_querystring",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-querystring",
        sum = "h1:AnCroh3fv4ZBgVIf1Iwtovgjaw/GiKJo8M8yD/fhyJ8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_google_go_replayers_grpcreplay",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-replayers/grpcreplay",
        sum = "h1:eNb1y9rZFmY4ax45uEEECSa8fsxGRU+8Bil52ASAwic=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_go_replayers_httpreplay",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-replayers/httpreplay",
        sum = "h1:AX7FUb4BjrrzNvblr/OlgwrmFiep6soj5K2QSDW7BGk=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_gofuzz",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/gofuzz",
        sum = "h1:xRy4A+RhZaiKjJ1bPfwQ8sedCA+YS2YcCHW6ec7JMi0=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_google_gopacket",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/gopacket",
        sum = "h1:ves8RnFZPGiFnTS0uPQStjwru6uO6h+nlr9j6fL7kF8=",
        version = "v1.1.19",
    )
    go_repository(
        name = "com_github_google_licenseclassifier",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/licenseclassifier",
        sum = "h1:EfzlPF5MRmoWsCGvSkPZ1Nh9uVzHf4FfGnDQ6CXd2NA=",
        version = "v0.0.0-20210325184830-bb04aff29e72",
    )
    go_repository(
        name = "com_github_google_martian",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/martian",
        sum = "h1:xmapqc1AyLoB+ddYT6r04bD9lIjlOqGaREovi0SzFaE=",
        version = "v2.1.1-0.20190517191504-25dcb96d9e51+incompatible",
    )
    go_repository(
        name = "com_github_google_martian_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/martian/v3",
        sum = "h1:d8MncMlErDFTwQGBK1xhv026j9kqhvw1Qv9IbWT1VLQ=",
        version = "v3.2.1",
    )
    go_repository(
        name = "com_github_google_monologue",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/monologue",
        sum = "h1:0L/piDwninh6sjZ+vCZI7c6RA0R71ET8v1yinZzC9k8=",
        version = "v0.0.0-20191220140058-35abc9683a6c",
    )
    go_repository(
        name = "com_github_google_pprof",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/pprof",
        sum = "h1:pDhu5sgp8yJlEF/g6osliIIpF9K4F5jvkULXa4daRDQ=",
        version = "v0.0.0-20230821062121-407c9e7a662f",
    )
    go_repository(
        name = "com_github_google_renameio",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/renameio",
        sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_rpmpack",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/rpmpack",
        sum = "h1:BW6OvS3kpT5UEPbCZ+KyX/OB4Ks9/MNMhWjqPPkZxsE=",
        version = "v0.0.0-20191226140753-aa36bfddb3a0",
    )
    go_repository(
        name = "com_github_google_s2a_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/s2a-go",
        sum = "h1:1kZ/sQM3srePvKs3tXAvQzo66XfcReoqFpIpIccE7Oc=",
        version = "v0.1.4",
    )
    go_repository(
        name = "com_github_google_shlex",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/shlex",
        sum = "h1:El6M4kTTCOh6aBiKaUGG7oYTSPP8MxqL4YI3kZKwcP4=",
        version = "v0.0.0-20191202100458-e7afc7fbc510",
    )
    go_repository(
        name = "com_github_google_subcommands",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/subcommands",
        sum = "h1:/eqq+otEXm5vhfBrbREPCSVQbvofip6kIz+mX5TUH7k=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_google_trillian",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/trillian",
        sum = "h1:hXrMoOL/sEOsFF2Snw1pmE1fn8FH9qktivGvjf4y/wU=",
        version = "v1.3.14-0.20210622121126-870e0cdde059",
    )
    go_repository(
        name = "com_github_google_trillian_examples",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/trillian-examples",
        sum = "h1:dv2J28D109qglM6VfNzAXZ7VddBojviT5oMSs1yeDUY=",
        version = "v0.0.0-20190603134952-4e75ba15216c",
    )
    go_repository(
        name = "com_github_google_uuid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/uuid",
        sum = "h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_google_wire",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/wire",
        sum = "h1:kXcsA/rIGzJImVqPdhfnr6q0xsS9gU0515q1EPpJ9fE=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_googleapis_enterprise_certificate_proxy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/enterprise-certificate-proxy",
        sum = "h1:yk9/cqRKtT9wXZSsRH9aurXEpJX+U6FLtpYTdC3R06k=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_googleapis_gax_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/gax-go",
        sum = "h1:silFMLAnr330+NRuag/VjIGF7TLp/LBrV2CJKFLWEww=",
        version = "v2.0.2+incompatible",
    )
    go_repository(
        name = "com_github_googleapis_gax_go_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/gax-go/v2",
        sum = "h1:9V9PWXEsWnPpQhu/PeQIkS4eGzMlTLGgt80cUUI8Ki4=",
        version = "v2.11.0",
    )
    go_repository(
        name = "com_github_googleapis_gnostic",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/gnostic",
        sum = "h1:9fHAtK0uDfpveeqqo1hkEZJcFvYXAiCN3UutL8F9xHw=",
        version = "v0.5.5",
    )
    go_repository(
        name = "com_github_googlecloudplatform_cloudsql_proxy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/GoogleCloudPlatform/cloudsql-proxy",
        sum = "h1:sTOp2Ajiew5XIH92YSdwhYc+bgpUX5j5TKK/Ac8Saw8=",
        version = "v0.0.0-20191009163259-e802c2cb94ae",
    )
    go_repository(
        name = "com_github_googlecloudplatform_k8s_cloud_provider",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/GoogleCloudPlatform/k8s-cloud-provider",
        sum = "h1:N7lSsF+R7wSulUADi36SInSQA3RvfO/XclHQfedr0qk=",
        version = "v0.0.0-20190822182118-27a4ced34534",
    )
    go_repository(
        name = "com_github_gookit_color",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gookit/color",
        sum = "h1:xOYBan3Fwlrqj1M1UN2TlHOCRiek3bGzWf/vPnJ1roE=",
        version = "v1.2.4",
    )
    go_repository(
        name = "com_github_gophercloud_gophercloud",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gophercloud/gophercloud",
        sum = "h1:6sjpKIpVwRIIwmcEGp+WwNovNsem+c+2vm6oxshRpL8=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_gopherjs_gopherjs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gopherjs/gopherjs",
        sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
        version = "v0.0.0-20181017120253-0766667cb4d1",
    )
    go_repository(
        name = "com_github_goproxyio_goproxy_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/goproxyio/goproxy/v2",
        sum = "h1:GXJVqfe8nwaIUYuMwUIoCJXU1a+Hq8MjhMWGjsWLbYI=",
        version = "v2.0.5",
    )
    go_repository(
        name = "com_github_goproxyio_windows",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/goproxyio/windows",
        sum = "h1:geXWJzWrEzC5S4MjTi8N+fhzuSYgdf4FwcYb4OI/YGU=",
        version = "v0.0.0-20191126033816-f4a809841617",
    )
    go_repository(
        name = "com_github_gordonklaus_ineffassign",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gordonklaus/ineffassign",
        sum = "h1:vc7Dmrk4JwS0ZPS6WZvWlwDflgDTA26jItmbSj83nug=",
        version = "v0.0.0-20200309095847-7953dde2c7bf",
    )
    go_repository(
        name = "com_github_goreleaser_goreleaser",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/goreleaser/goreleaser",
        sum = "h1:Z+7XPrfGK11s/Sp+a06sx2FzGuCjTBdxN2ubpGvQbjY=",
        version = "v0.136.0",
    )
    go_repository(
        name = "com_github_goreleaser_nfpm",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/goreleaser/nfpm",
        sum = "h1:BPwIomC+e+yuDX9poJowzV7JFVcYA0+LwGSkbAPs2Hw=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_gorilla_context",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/context",
        sum = "h1:AWwleXJkX/nhcU9bZSnZoi3h/qGYqQAGhq6zZe/aQW8=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gorilla_csrf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/csrf",
        sum = "h1:60oN1cFdncCE8tjwQ3QEkFND5k37lQPcRjnlvm7CIJ0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_gorilla_handlers",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/handlers",
        sum = "h1:cLTUSsNkgcwhgRqvCNmdbRWG0A3N4F+M2nWKdScwyEE=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_gorilla_mux",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/mux",
        sum = "h1:TuBL49tXwgrFYWhqrNgrUNEY92u81SPhu7sTdzQEiWY=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_gorilla_rpc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/rpc",
        sum = "h1:WvvdC2lNeT1SP32zrIce5l0ECBfbAlmrmSBsuc57wfk=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_gorilla_securecookie",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/securecookie",
        sum = "h1:miw7JPhV+b/lAHSXz4qd/nN9jRiAFV5FwjeKyCS8BvQ=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gorilla_sessions",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/sessions",
        sum = "h1:DHd3rPN5lE3Ts3D8rKkQ8x/0kqfeNmBAaiSi+o7FsgI=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_gorilla_websocket",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/websocket",
        sum = "h1:PPwGk2jz7EePpoHN/+ClbZu8SPxiqlu12wZP/3sWmnc=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_gostaticanalysis_analysisutil",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gostaticanalysis/analysisutil",
        sum = "h1:iwp+5/UAyzQSFgQ4uR2sni99sJ8Eo9DEacKWM5pekIg=",
        version = "v0.0.3",
    )
    go_repository(
        name = "com_github_graph_gophers_graphql_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/graph-gophers/graphql-go",
        sum = "h1:Eb9x/q6MFpCLz7jBCiP/WTxjSDrYLR1QY41SORZyNJ0=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_greenpau_caddy_auth_jwt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/greenpau/caddy-auth-jwt",
        sum = "h1:CqM0yyAmLbD5vQvl5XaQkyMyq/AZnjk9ZneakxE8Xzo=",
        version = "v1.2.6",
    )
    go_repository(
        name = "com_github_greenpau_caddy_trace",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/greenpau/caddy-trace",
        sum = "h1:Hi5vBKCzUS9XKjjy5TGzjr+tQef0NvPE+8lzVs3Uvsw=",
        version = "v1.1.8",
    )
    go_repository(
        name = "com_github_gregjones_httpcache",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/gregjones/httpcache",
        sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
        version = "v0.0.0-20180305231024-9cad4c3443a7",
    )
    go_repository(
        name = "com_github_groob_finalizer",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/groob/finalizer",
        sum = "h1:5ikpG9mYCMFiZX0nkxoV6aU2IpCHPdws3gCNgdZeEV0=",
        version = "v0.0.0-20170707115354-4c2ed49aabda",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_middleware",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
        sum = "h1:+9834+KizmvFV7pXQGSXQTsaWhq2GjuNUt0aUU0YBYw=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_prometheus",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
        sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/grpc-ecosystem/grpc-gateway/v2",
        sum = "h1:YBftPWNWd4WwGqtY2yeZL2ef8rHAxPBD8KFhJpmcqms=",
        version = "v2.16.0",
    )
    go_repository(
        name = "com_github_h12w_go_socks5",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/h12w/go-socks5",
        sum = "h1:5XxdakFhqd9dnXoAZy1Mb2R/DZ6D1e+0bGC/JhucGYI=",
        version = "v0.0.0-20200522160539-76189e178364",
    )
    go_repository(
        name = "com_github_h2non_parth",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/h2non/parth",
        sum = "h1:2VTzZjLZBgl62/EtslCrtky5vbi9dd7HrQPQIx6wqiw=",
        version = "v0.0.0-20190131123155-b4df798d6542",
    )
    go_repository(
        name = "com_github_hashicorp_consul_api",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/consul/api",
        sum = "h1:HXNYlRkkM/t+Y/Yhxtwcy02dlYwIaoxzvxPnS+cqy78=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_hashicorp_consul_sdk",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/consul/sdk",
        sum = "h1:UOxjlb4xVNF93jak1mzzoBatyFju9nrkxpVwIp/QqxQ=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_hashicorp_cronexpr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/cronexpr",
        sum = "h1:NJZDd87hGXjoZBdvyCF9mX4DCq5Wy7+A/w+A7q0wn6c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_errwrap",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/errwrap",
        sum = "h1:OxrOeh75EUXMY8TBjag2fzXGZ40LB6IKw45YeGUDY2I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_bexpr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-bexpr",
        sum = "h1:9kuI5PFotCboP3dkDYFr/wi0gg0QVbSNz5oFRpxn4uE=",
        version = "v0.1.10",
    )
    go_repository(
        name = "com_github_hashicorp_go_cleanhttp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-cleanhttp",
        sum = "h1:035FKYIWjmULyFRBKPs8TBQoi0x6d9G4xc9neXJWAZQ=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_envparse",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-envparse",
        sum = "h1:bE++6bhIsNCPLvgDZkYqo3nA+/PFI51pkrHdmPSDFPY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_hclog",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-hclog",
        sum = "h1:K4ev2ib4LdQETX5cSZBG0DVLk1jwGqSPXBjdah3veNs=",
        version = "v0.16.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_immutable_radix",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-immutable-radix",
        sum = "h1:DKHmCUm2hRBK510BaiZlwvpD40f8bJFeZnpfm2KLowc=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_kms_wrapping_entropy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-kms-wrapping/entropy",
        sum = "h1:xuTi5ZwjimfpvpL09jDE71smCBRpnF5xfo871BSX4gs=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_msgpack",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-msgpack",
        sum = "h1:i9R9JSrqIz0QVLz3sz+i3YJdT7TTSLcfLLzJi9aZTuI=",
        version = "v0.5.5",
    )
    go_repository(
        name = "com_github_hashicorp_go_multierror",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-multierror",
        sum = "h1:H5DkEtf6CXdFp0N0Em5UCwQpXMWke8IA0+lD48awMYo=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_net",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go.net",
        sum = "h1:sNCoNyDEvN1xa+X0baata4RdcpKwcMS6DH+xwfqPgjw=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_plugin",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-plugin",
        sum = "h1:DXmvivbWD5qdiBts9TpBC7BYL1Aia5sxbRgQB+v6UZM=",
        version = "v1.4.3",
    )
    go_repository(
        name = "com_github_hashicorp_go_retryablehttp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-retryablehttp",
        sum = "h1:ZQgVdpTdAL7WpMIwLzCfbalOcSUdkDZnpUv3/+BxzFA=",
        version = "v0.7.4",
    )
    go_repository(
        name = "com_github_hashicorp_go_rootcerts",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-rootcerts",
        sum = "h1:jzhAVGtqPKbwpyCPELlgNWhE1znq+qwJtW5Oi2viEzc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_secure_stdlib_base62",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-secure-stdlib/base62",
        sum = "h1:6KMBnfEv0/kLAz0O76sliN5mXbCDcLfs2kP7ssP7+DQ=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_secure_stdlib_mlock",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-secure-stdlib/mlock",
        sum = "h1:cCRo8gK7oq6A2L6LICkUZ+/a5rLiRXFMf1Qd4xSwxTc=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_secure_stdlib_parseutil",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-secure-stdlib/parseutil",
        sum = "h1:78ki3QBevHwYrVxnyVeaEz+7WtifHhauYF23es/0KlI=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_secure_stdlib_password",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-secure-stdlib/password",
        sum = "h1:6JzmBqXprakgFEHwBgdchsjaA9x3GyjdI568bXKxa60=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_secure_stdlib_strutil",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-secure-stdlib/strutil",
        sum = "h1:nd0HIW15E6FG1MsnArYaHfuw9C2zgzM8LxkG5Ty/788=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_secure_stdlib_tlsutil",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-secure-stdlib/tlsutil",
        sum = "h1:Yc026VyMyIpq1UWRnakHRG01U8fJm+nEfEmjoAb00n8=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_sockaddr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-sockaddr",
        sum = "h1:ztczhD1jLxIRjVejw8gFomI1BQZOe2WoVOu0SyteCQc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_syslog",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-syslog",
        sum = "h1:KaodqZuhUoZereWVIYmpUgZysurB1kBLX2j0MwMrUAE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_uuid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-uuid",
        sum = "h1:cfejS+Tpcp13yd5nYHWDI6qVCny6wyX2Mt5SGur2IGE=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_version",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-version",
        sum = "h1:feTTfFNnjP967rlCxM/I9g701jU+RN74YKx2mOkIeek=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/golang-lru",
        sum = "h1:dg1dEPuWpEqDnvIw251EVy4zlP8gWbsGj4BsUKCRpYs=",
        version = "v0.5.5-0.20210104140557-80c98217689d",
    )
    go_repository(
        name = "com_github_hashicorp_hcl",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/hcl",
        sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_logutils",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/logutils",
        sum = "h1:dLEQVugN8vlakKOUE3ihGLTZJRB4j+M2cdTm/ORI65Y=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_mdns",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/mdns",
        sum = "h1:WhIgCr5a7AaVH6jPUwjtRuuE7/RDufnUvzIr48smyxs=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_memberlist",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/memberlist",
        sum = "h1:gkyML/r71w3FL8gUi74Vk76avkj/9lYAY9lvg0OcoGs=",
        version = "v0.1.4",
    )
    go_repository(
        name = "com_github_hashicorp_nomad_api",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/nomad/api",
        sum = "h1:Y/0SMGEylkIwKtWZfKA+SVSIbOAKLMQPz/BjJG4vwHo=",
        version = "v0.0.0-20220211135303-4afc67b7002e",
    )
    go_repository(
        name = "com_github_hashicorp_serf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/serf",
        sum = "h1:MWYcmct5EtKz0efYooPcL0yNkem+7kWxqXDi/UIh+8k=",
        version = "v0.8.3",
    )
    go_repository(
        name = "com_github_hashicorp_vault_api",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/vault/api",
        sum = "h1:pkDkcgTh47PRjY1NEFeofqR4W/HkNUi9qIakESO2aRM=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_hashicorp_vault_api_auth_approle",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/vault/api/auth/approle",
        sum = "h1:R5yA+xcNvw1ix6bDuWOaLOq2L4L77zDCVsethNw97xQ=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_vault_sdk",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/vault/sdk",
        sum = "h1:kR3dpxNkhh/wr6ycaJYqp6AFT/i2xaftbfnwZduTKEY=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_hashicorp_yamux",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/yamux",
        sum = "h1:b5rjCoWHc7eqmAS4/qyk21ZsHyb6Mxv/jykxvNTkU4M=",
        version = "v0.0.0-20180604194846-3520598351bb",
    )
    go_repository(
        name = "com_github_hodgesds_perf_utils",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hodgesds/perf-utils",
        sum = "h1:7KlHGMuig4FRH5fNw68PV6xLmgTe7jKs9hgAcEAbioU=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_holiman_bloomfilter_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/holiman/bloomfilter/v2",
        sum = "h1:73e0e/V0tCydx14a0SCYS/EWCxgwLZ18CZcZKVu0fao=",
        version = "v2.0.3",
    )
    go_repository(
        name = "com_github_holiman_uint256",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/holiman/uint256",
        sum = "h1:gpSYcPLWGv4sG43I2mVLiDZCNDh/EpGjSk8tmtxitHM=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_hpcloud_tail",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hpcloud/tail",
        sum = "h1:Ysi1UhrSyBltF8f+3RAt4UaqHc+53JJ0jyl0pY0sfck=",
        version = "v1.0.1-0.20180514194441-a1dbeea552b7",
    )
    go_repository(
        name = "com_github_huandu_xstrings",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/huandu/xstrings",
        sum = "h1:L18LIDzqlW6xN2rEkpdV8+oL/IXWJ1APd+vsdYy4Wdw=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_hudl_fargo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hudl/fargo",
        sum = "h1:0U6+BtN6LhaYuTnIJq4Wyq5cpn6O2kWrxAtcqBmYY6w=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_hugelgupf_socketpair",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/hugelgupf/socketpair",
        sum = "h1:/jC7qQFrv8CrSJVmaolDVOxTfS9kc36uB6H40kdbQq8=",
        version = "v0.0.0-20190730060125-05d35a94e714",
    )
    go_repository(
        name = "com_github_huin_goupnp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/huin/goupnp",
        sum = "h1:N8No57ls+MnjlB+JPiCVSOyy/ot7MJTqlo7rn+NYSqQ=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_ianlancetaylor_demangle",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ianlancetaylor/demangle",
        sum = "h1:BA4a7pe6ZTd9F8kXETBoijjFJ/ntaa//1wiH9BZu4zU=",
        version = "v0.0.0-20230524184225-eabc099b10ab",
    )
    go_repository(
        name = "com_github_icrowley_fake",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/icrowley/fake",
        sum = "h1:Mo9W14pwbO9VfRe+ygqZ8dFbPpoIK1HFrG/zjTuQ+nc=",
        version = "v0.0.0-20180203215853-4178557ae428",
    )
    go_repository(
        name = "com_github_iij_doapi",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/iij/doapi",
        sum = "h1:MZf03xP9WdakyXhOWuAD5uPK3wHh96wCsqe3hCMKh8E=",
        version = "v0.0.0-20190504054126-0bbf12d6d7df",
    )
    go_repository(
        name = "com_github_illumos_go_kstat",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/illumos/go-kstat",
        sum = "h1:hk4LPqXIY/c9XzRbe7dA6qQxaT6Axcbny0L/G5a4owQ=",
        version = "v0.0.0-20210513183136-173c9b0a9973",
    )
    go_repository(
        name = "com_github_imdario_mergo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/imdario/mergo",
        sum = "h1:b6R2BslTbIEToALKP7LxUvijTsNI9TAe80pLWN2g/HU=",
        version = "v0.3.12",
    )
    go_repository(
        name = "com_github_imkira_go_interpol",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/imkira/go-interpol",
        sum = "h1:KIiKr0VSG2CUW1hl1jpiyuzuJeKUUpC8iM1AIE7N1Vk=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_inconshreveable_mousetrap",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/inconshreveable/mousetrap",
        sum = "h1:wN+x4NVGpMsO7ErUn/mUI3vEoE6Jt13X2s0bqwp9tc8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_influxdata_influxdb",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/influxdata/influxdb",
        sum = "h1:WEypI1BQFTT4teLM+1qkEcvUi0dAvopAI/ir0vAiBg8=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_github_influxdata_influxdb1_client",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/influxdata/influxdb1-client",
        sum = "h1:/WZQPMZNsjZ7IlCpsLGdQBINg5bxKQ1K1sh6awxLtkA=",
        version = "v0.0.0-20191209144304-8bf82d3c094d",
    )
    go_repository(
        name = "com_github_influxdata_influxdb_client_go_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/influxdata/influxdb-client-go/v2",
        sum = "h1:HGBfZYStlx3Kqvsv1h2pJixbCl/jhnFtxpKFAv9Tu5k=",
        version = "v2.4.0",
    )
    go_repository(
        name = "com_github_influxdata_line_protocol",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/influxdata/line-protocol",
        sum = "h1:vilfsDSy7TDxedi9gyBkMvAirat/oRcL0lFdJBf6tdM=",
        version = "v0.0.0-20210311194329-9aa0e372d097",
    )
    go_repository(
        name = "com_github_influxdb_influxdb",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/influxdb/influxdb",
        sum = "h1:+8XPwrWoNps4WbLfhNhH0ct8LUJAP1q+faViiEpSHYc=",
        version = "v0.9.6-0.20151125225445-9eab56311373",
    )
    go_repository(
        name = "com_github_infobloxopen_go_trees",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/infobloxopen/go-trees",
        sum = "h1:w66aaP3c6SIQ0pi3QH1Tb4AMO3aWoEPxd1CNvLphbkA=",
        version = "v0.0.0-20200715205103-96a057b8dfb9",
    )
    go_repository(
        name = "com_github_insomniacslk_dhcp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/insomniacslk/dhcp",
        sum = "h1:l1QCwn715k8nYkj4Ql50rzEog3WnMdrd4YYMMwemxEo=",
        version = "v0.0.0-20220504074936-1ca156eafb9f",
    )
    go_repository(
        name = "com_github_intel_goresctrl",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/intel/goresctrl",
        sum = "h1:JyZjdMQu9Kl/wLXe9xA6s1X+tF6BWsQPFGJMEeCfWzE=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_j_keck_arping",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/j-keck/arping",
        sum = "h1:hlLhuXgQkzIJTZuhMigvG/CuSkaspeaD9hRDk2zuiMI=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_jackc_chunkreader",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/chunkreader",
        sum = "h1:4s39bBR8ByfqH+DKm8rQA3E1LHZWB9XWcrz8fqaZbe0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_chunkreader_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/chunkreader/v2",
        sum = "h1:i+RDz65UE+mmpjTfyz0MoVTnzeYxroil2G82ki7MGG8=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_jackc_pgconn",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgconn",
        sum = "h1:DzdIHIjG1AxGwoEEqS+mGsURyjt4enSmqzACXvVzOT8=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_jackc_pgio",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgio",
        sum = "h1:g12B9UwVnzGhueNavwioyEEpAmqMe1E/BN9ES+8ovkE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgmock",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgmock",
        sum = "h1:DadwsjnMwFjfWc9y5Wi/+Zz7xoE5ALHsRQlOctkOiHc=",
        version = "v0.0.0-20210724152146-4ad1a8207f65",
    )
    go_repository(
        name = "com_github_jackc_pgpassfile",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgpassfile",
        sum = "h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgproto3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgproto3",
        sum = "h1:FYYE4yRw+AgI8wXIinMlNjBbp/UitDJwfj5LqqewP1A=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_jackc_pgproto3_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgproto3/v2",
        sum = "h1:r7JypeP2D3onoQTCxWdTpCtJ4D+qpKr0TxvoyMhZ5ns=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_jackc_pgservicefile",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgservicefile",
        sum = "h1:C8S2+VttkHFdOOCXJe+YGfa4vHYwlt4Zx+IVXQ97jYg=",
        version = "v0.0.0-20200714003250-2b9c44734f2b",
    )
    go_repository(
        name = "com_github_jackc_pgtype",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgtype",
        sum = "h1:/SH1RxEtltvJgsDqp3TbiTFApD3mey3iygpuEGeuBXk=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_jackc_pgx_v4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgx/v4",
        sum = "h1:TgdrmgnM7VY72EuSQzBbBd4JA1RLqJolrw9nQVZABVc=",
        version = "v4.14.0",
    )
    go_repository(
        name = "com_github_jackc_puddle",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/puddle",
        sum = "h1:DNDKdn/pDrWvDWyT2FYvpZVE81OAhWrjCv19I9n108Q=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_jackpal_go_nat_pmp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackpal/go-nat-pmp",
        sum = "h1:KzKSgb7qkJvOUTqYl9/Hg/me3pWgBmERKrTGD7BdWus=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_jarcoal_httpmock",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jarcoal/httpmock",
        sum = "h1:cHtVEcTxRSX4J0je7mWPfc9BpDpqzXSJ5HbymZmyHck=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_jaypipes_ghw",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jaypipes/ghw",
        sum = "h1:L4Bg16C1brIyEOo025bILWbtJmJMj7Nvv36nbnaTK3M=",
        version = "v0.8.1-0.20220131141055-fb0598ce62c8",
    )
    go_repository(
        name = "com_github_jaypipes_pcidb",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jaypipes/pcidb",
        sum = "h1:VIM7GKVaW4qba30cvB67xSCgJPTzkG8Kzw/cbs5PHWU=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_jbenet_go_context",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jbenet/go-context",
        sum = "h1:BQSFePA1RWJOlocH6Fxy8MmwDt+yVQYULKfN0RoTN8A=",
        version = "v0.0.0-20150711004518-d14ea06fba99",
    )
    go_repository(
        name = "com_github_jcmturner_aescts_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/aescts/v2",
        sum = "h1:9YKLH6ey7H4eDBXW8khjYslgyqG2xZikXP0EQFKrle8=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_dnsutils_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/dnsutils/v2",
        sum = "h1:lltnkeZGL0wILNvrNiVCR6Ro5PGU/SeBvVO/8c/iPbo=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_gofork",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/gofork",
        sum = "h1:J7uCkflzTEhUZ64xqKnkDxq3kzc96ajM1Gli5ktUem8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_goidentity_v6",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/goidentity/v6",
        sum = "h1:VKnZd2oEIMorCTsFBnJWbExfNN7yZr3EhJAxwOkZg6o=",
        version = "v6.0.1",
    )
    go_repository(
        name = "com_github_jcmturner_gokrb5_v8",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/gokrb5/v8",
        sum = "h1:6ZIM6b/JJN0X8UM43ZOM6Z4SJzla+a/u7scXFJzodkA=",
        version = "v8.4.2",
    )
    go_repository(
        name = "com_github_jcmturner_rpc_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/rpc/v2",
        sum = "h1:7FXXj8Ti1IaVFpSAziCZWNzbNuZmnvw/i6CqLNdWfZY=",
        version = "v2.0.3",
    )
    go_repository(
        name = "com_github_jedisct1_go_minisign",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jedisct1/go-minisign",
        sum = "h1:UvSe12bq+Uj2hWd8aOlwPmoZ+CITRFrdit+sDGfAg8U=",
        version = "v0.0.0-20190909160543-45766022959e",
    )
    go_repository(
        name = "com_github_jellevandenhooff_dkim",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jellevandenhooff/dkim",
        sum = "h1:ujPKutqRlJtcfWk6toYVYagwra7HQHbXOaS171b4Tg8=",
        version = "v0.0.0-20150330215556-f50fe3d243e1",
    )
    go_repository(
        name = "com_github_jessevdk_go_flags",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jessevdk/go-flags",
        sum = "h1:4IU2WS7AumrZ/40jfhf4QVDMsQwqA7VEHozFRrGARJA=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_jhump_protoreflect",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jhump/protoreflect",
        sum = "h1:k2xE7wcUomeqwY0LDCYA16y4WWfyTcMx5mKhk0d4ua0=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_github_jimstudt_http_authentication",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jimstudt/http-authentication",
        sum = "h1:BcF8coBl0QFVhe8vAMMlD+CV8EISiu9MGKLoj6ZEyJA=",
        version = "v0.0.0-20140401203705-3eca13d6893a",
    )
    go_repository(
        name = "com_github_jingyugao_rowserrcheck",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jingyugao/rowserrcheck",
        sum = "h1:GmsqmapfzSJkm28dhRoHz2tLRbJmqhU86IPgBtN3mmk=",
        version = "v0.0.0-20191204022205-72ab7603b68a",
    )
    go_repository(
        name = "com_github_jirfag_go_printf_func_name",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jirfag/go-printf-func-name",
        sum = "h1:KA9BjwUk7KlCh6S9EAGWBt1oExIUv9WyNCiRz5amv48=",
        version = "v0.0.0-20200119135958-7558a9eaa5af",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jmespath/go-jmespath",
        sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath_internal_testify",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jmespath/go-jmespath/internal/testify",
        sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_jmoiron_sqlx",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jmoiron/sqlx",
        sum = "h1:lrdPtrORjGv1HbbEvKWDUAy97mPpFm4B8hp77tcCUJY=",
        version = "v1.2.1-0.20190826204134-d7d95172beb5",
    )
    go_repository(
        name = "com_github_joefitzgerald_rainbow_reporter",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/joefitzgerald/rainbow-reporter",
        sum = "h1:AuMG652zjdzI0YCCnXAqATtRBpGXMcAnrajcaTrSeuo=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_joho_godotenv",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/joho/godotenv",
        sum = "h1:Zjp+RcGpHhGlrMbJzXTrZZPrWj+1vfm90La1wgB6Bhc=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_jonboulle_clockwork",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jonboulle/clockwork",
        sum = "h1:UOGuzwb1PwsrDAObMuhUnj0p5ULPj8V/xJ7Kx9qUBdQ=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_josharian_intern",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/josharian/intern",
        sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_josharian_native",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/josharian/native",
        sum = "h1:uuaP0hAbW7Y4l0ZRQ6C9zfb7Mg1mbFKry/xzDAfmtLA=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_jpillora_ansi",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/ansi",
        sum = "h1:+Ei5HCAH0xsrQRCT2PDr4mq9r4Gm4tg+arNdXRkB22s=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_jpillora_backoff",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/backoff",
        sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jpillora_chisel",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/chisel",
        sum = "h1:eLbzoX+ekDhVmF5CpSJD01NtH/w7QMYeaFCIFbzn9ns=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_github_jpillora_cookieauth",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/cookieauth",
        sum = "h1:1BYnSG4c+/5kutOsY+7/Ba+rRhUfYv61Jrf605CxfRw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jpillora_eventsource",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/eventsource",
        sum = "h1:iMFSHw9uUmQyNWKHylAS9HoK9R9ps2NWqsjzKniCFus=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jpillora_go_echo_server",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/go-echo-server",
        sum = "h1:NXP4foXo1U0oMCrjMTmiG1v6PoHxVhQpb9lEK9xe0zw=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_jpillora_ipfilter",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/ipfilter",
        sum = "h1:K5zGPjyjgf2MPB+iTULZ7Hl4zXPWOb4JwgxMdogKq20=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_jpillora_opts",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/opts",
        sum = "h1:H8vWooV3P9nsqmCcPgxNZyIa7GPOWA1KQFsfAzIkCtE=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_jpillora_requestlog",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/requestlog",
        sum = "h1:bg++eJ74T7DYL3DlIpiwknrtfdUA9oP/M4fL+PpqnyA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jpillora_sizestr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/sizestr",
        sum = "h1:4tr0FLxs1Mtq3TnsLDV+GYUWG7Q26a6s+tV5Zfw2ygw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jpillora_velox",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/velox",
        sum = "h1:96BPWo+zw2ywf4T/3C9kEKhzGABsjJ4vWSOkf93el/g=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_jpillora_webproc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/webproc",
        sum = "h1:0BmX+F+HDs0uTYzjy3CSU/7oM8yTTuPkqsJim+b5Ndg=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_jsimonetti_rtnetlink",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jsimonetti/rtnetlink",
        sum = "h1:hVlNQNRlLDGZz31gBPicsG7Q53rnlsz1l1Ix/9XlpVA=",
        version = "v1.3.5",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/json-iterator/go",
        sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
        version = "v1.1.12",
    )
    go_repository(
        name = "com_github_jstemmer_go_junit_report",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jstemmer/go-junit-report",
        sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_jsternberg_zap_logfmt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jsternberg/zap-logfmt",
        sum = "h1:1v+PK4/B48cy8cfQbxL4FmmNZrjnIMr2BsnyEmXqv2o=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_jtolds_gls",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/jtolds/gls",
        sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
        version = "v4.20.0+incompatible",
    )
    go_repository(
        name = "com_github_juju_ansiterm",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/juju/ansiterm",
        sum = "h1:FaWFmfWdAUKbSCtOU2QjDaorUexogfaMgbipgYATUMU=",
        version = "v0.0.0-20180109212912-720a0952cc2a",
    )
    go_repository(
        name = "com_github_juju_ratelimit",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/juju/ratelimit",
        sum = "h1:+7AIFJVQ0EQgq/K9+0Krm7m530Du7tIz0METWzN0RgY=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_julienschmidt_httprouter",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/julienschmidt/httprouter",
        sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_k0kubun_colorstring",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/k0kubun/colorstring",
        sum = "h1:uC1QfSlInpQF+M0ao65imhwqKnz3Q2z/d8PWZRMQvDM=",
        version = "v0.0.0-20150214042306-9440f1994b88",
    )
    go_repository(
        name = "com_github_karalabe_usb",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/karalabe/usb",
        sum = "h1:M6QQBNxF+CQ8OFvxrT90BA0qBOXymndZnk5q235mFc4=",
        version = "v0.0.2",
    )
    go_repository(
        name = "com_github_kardianos_service",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kardianos/service",
        sum = "h1:bGuZ/epo3vrt8IPC7mnKQolqFeYJb7Cs8Rk4PSOBB/g=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_karrick_godirwalk",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/karrick/godirwalk",
        sum = "h1:DynhcF+bztK8gooS0+NDJFrdNZjJ3gzVzC545UNA9iw=",
        version = "v1.16.1",
    )
    go_repository(
        name = "com_github_kataras_basicauth",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kataras/basicauth",
        sum = "h1:Do2epS3AA2HJewna/xVbiOz0eqPfDl3BdWT+oV3/PFw=",
        version = "v0.0.3",
    )
    go_repository(
        name = "com_github_kballard_go_shellquote",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kballard/go-shellquote",
        sum = "h1:Z9n2FFNUXsshfwJMBgNA0RU6/i7WVaAegv3PtuIHPMs=",
        version = "v0.0.0-20180428030007-95032a82bc51",
    )
    go_repository(
        name = "com_github_kevinburke_ssh_config",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kevinburke/ssh_config",
        sum = "h1:DowS9hvgyYSX4TO5NpyC606/Z4SxnNYbT+WX27or6Ck=",
        version = "v0.0.0-20201106050909-4977a11b4351",
    )
    go_repository(
        name = "com_github_kisielk_errcheck",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kisielk/errcheck",
        sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_kisielk_gotool",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kisielk/gotool",
        sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_klauspost_compress",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/klauspost/compress",
        sum = "h1:wKRjX6JRtDdrE9qwa4b/Cip7ACOshUI4smpCQanqjSY=",
        version = "v1.15.9",
    )
    go_repository(
        name = "com_github_klauspost_cpuid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/klauspost/cpuid",
        sum = "h1:5JNjFYYQrZeKRJ0734q51WCEEn2huer72Dc7K+R/b6s=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_klauspost_cpuid_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/klauspost/cpuid/v2",
        sum = "h1:1XxvOiqXZ8SULZUKim/wncr3wZ38H4yCuVDvKdK9OGs=",
        version = "v2.0.13",
    )
    go_repository(
        name = "com_github_klauspost_reedsolomon",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/klauspost/reedsolomon",
        sum = "h1:g2erWKD2M6rgnPf89fCji6jNlhMKMdXcuNHMW1SYCIo=",
        version = "v1.9.15",
    )
    go_repository(
        name = "com_github_knetic_govaluate",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Knetic/govaluate",
        sum = "h1:1G1pk05UrOh0NlF1oeaaix1x8XzrfjIDK47TY0Zehcw=",
        version = "v3.0.1-0.20171022003610-9aa49832a739+incompatible",
    )
    go_repository(
        name = "com_github_koding_websocketproxy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/koding/websocketproxy",
        sum = "h1:N7A4JCA2G+j5fuFxCsJqjFU/sZe0mj8H0sSoSwbaikw=",
        version = "v0.0.0-20181220232114-7ed82d81a28c",
    )
    go_repository(
        name = "com_github_kolo_xmlrpc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kolo/xmlrpc",
        sum = "h1:TrxPzApUukas24OMMVDUMlCs1XCExJtnGaDEiIAR4oQ=",
        version = "v0.0.0-20190717152603-07c4ee3fd181",
    )
    go_repository(
        name = "com_github_konsorten_go_windows_terminal_sequences",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/konsorten/go-windows-terminal-sequences",
        sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_kr_fs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/fs",
        sum = "h1:Jskdu9ieNAYnjxsi0LbQp1ulIKZV1LAFgK1tWhpZgl8=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_kr_logfmt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/logfmt",
        sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
        version = "v0.0.0-20140226030751-b84e30acd515",
    )
    go_repository(
        name = "com_github_kr_pretty",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/pretty",
        sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_kr_pty",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/pty",
        sum = "h1:AkaSdXYQOWeaO3neb8EM634ahkXXe3jYbVh/F9lq+GI=",
        version = "v1.1.8",
    )
    go_repository(
        name = "com_github_kr_text",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_kylelemons_godebug",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/kylelemons/godebug",
        sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_labbsr0x_bindman_dns_webhook",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/labbsr0x/bindman-dns-webhook",
        sum = "h1:I7ITbmQPAVwrDdhd6dHKi+MYJTJqPCK0jE6YNBAevnk=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_labbsr0x_goh",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/labbsr0x/goh",
        sum = "h1:97aBJkDjpyBZGPbQuOK5/gHcSFbcr5aRsq3RSRJFpPk=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_letsencrypt_pkcs11key",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/letsencrypt/pkcs11key",
        sum = "h1:GfzE+uq7odDW7nOmp1QWuilLEK7kJf8i84XcIfk3mKA=",
        version = "v2.0.1-0.20170608213348-396559074696+incompatible",
    )
    go_repository(
        name = "com_github_letsencrypt_pkcs11key_v4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/letsencrypt/pkcs11key/v4",
        sum = "h1:qLc/OznH7xMr5ARJgkZCCWk+EomQkiNTOoOF5LAgagc=",
        version = "v4.0.0",
    )
    go_repository(
        name = "com_github_liamhaworth_go_tproxy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/LiamHaworth/go-tproxy",
        sum = "h1:eqa6queieK8SvoszxCu0WwH7lSVeL4/N/f1JwOMw1G4=",
        version = "v0.0.0-20190726054950-ef7efd7f24ed",
    )
    go_repository(
        name = "com_github_lib_pq",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lib/pq",
        sum = "h1:AqzbZs4ZoCBp+GtejcpCpcxM3zlSMx29dXbUSeVtJb8=",
        version = "v1.10.2",
    )
    go_repository(
        name = "com_github_libdns_alidns",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/libdns/alidns",
        sum = "h1:KwuC1AmihOJjoFWXFRFaQx1PcD/jRpY04jo7yNK5zfk=",
        version = "v1.0.2-x2",
    )
    go_repository(
        name = "com_github_libdns_azure",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/libdns/azure",
        sum = "h1:SVYG+iMKtSpSJZBZ0hjETAMNscPoWRMJI7nnlLonwD4=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_libdns_cloudflare",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/libdns/cloudflare",
        sum = "h1:93WkJaGaiXCe353LHEP36kAWCUw0YjFqwhkBkU2/iic=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_libdns_digitalocean",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/libdns/digitalocean",
        sum = "h1:JYi/h0UEECrxY2JCi5FIfZEDFuUJvwihUWdm3bnDu2A=",
        version = "v0.0.0-20210310230526-186c4ebd2215",
    )
    go_repository(
        name = "com_github_libdns_dnspod",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/libdns/dnspod",
        sum = "h1:xJHDIujgLjvZnpB8/rMoCHUqA/KxSGBqRUXxSIzNzAA=",
        version = "v0.0.3",
    )
    go_repository(
        name = "com_github_libdns_duckdns",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/libdns/duckdns",
        sum = "h1:wkgu98DkwpjduH2fxC2YkCiNkNnfQiHaYskMkEyRZ28=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_libdns_libdns",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/libdns/libdns",
        sum = "h1:Wu59T7wSHRgtA0cfxC+n1c/e+O3upJGWytknkmFEDis=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_libdns_route53",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/libdns/route53",
        sum = "h1:etUVkopzG9xGEt34xfmYbpz6rTgAnv+n0vcV/1Xdc7c=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_libdns_vultr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/libdns/vultr",
        sum = "h1:ds9Nu9RwQWIHXM/e7264RiUfdyAgNdoZkiWXt0xSIzY=",
        version = "v0.0.0-20211122184636-cd4cb5c12e51",
    )
    go_repository(
        name = "com_github_liggitt_tabwriter",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/liggitt/tabwriter",
        sum = "h1:9TO3cAIGXtEhnIaL+V+BEER86oLrvS+kWobKpbJuye0=",
        version = "v0.0.0-20181228230101-89fcab3d43de",
    )
    go_repository(
        name = "com_github_lightstep_lightstep_tracer_common_golang_gogo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lightstep/lightstep-tracer-common/golang/gogo",
        sum = "h1:143Bb8f8DuGWck/xpNUOckBVYfFbBTnLevfRZ1aVVqo=",
        version = "v0.0.0-20190605223551-bc2310a04743",
    )
    go_repository(
        name = "com_github_lightstep_lightstep_tracer_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lightstep/lightstep-tracer-go",
        sum = "h1:vi1F1IQ8N7hNWytK9DpJsUfQhGuNSc19z330K6vl4zk=",
        version = "v0.18.1",
    )
    go_repository(
        name = "com_github_linode_linodego",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/linode/linodego",
        sum = "h1:AMdb82HVgY8o3mjBXJcUv9B+fnJjfDMn2rNRGbX+jvM=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_linuxkit_virtsock",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/linuxkit/virtsock",
        sum = "h1:jUp75lepDg0phMUJBCmvaeFDldD2N3S1lBuPwUTszio=",
        version = "v0.0.0-20201010232012-f8cee7dfc7a3",
    )
    go_repository(
        name = "com_github_liquidweb_liquidweb_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/liquidweb/liquidweb-go",
        sum = "h1:vIj1I/Wf97fUnyirD+bi6Y63c0GiXk9nKI1+sFFl3G0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_lithammer_dedent",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lithammer/dedent",
        sum = "h1:VNzHMVCBNG1j0fh3OrsFRkVUwStdDArbgBWoPAffktY=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_logrusorgru_aurora",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/logrusorgru/aurora",
        sum = "h1:9MlwzLdW7QSDrhDjFlsEYmxpFyIoXmYRon3dt0io31k=",
        version = "v0.0.0-20181002194514-a7b3b318ed4e",
    )
    go_repository(
        name = "com_github_lucas_clemente_aes12",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lucas-clemente/aes12",
        sum = "h1:sSeNEkJrs+0F9TUau0CgWTTNEwF23HST3Eq0A+QIx+A=",
        version = "v0.0.0-20171027163421-cd47fb39b79f",
    )
    go_repository(
        name = "com_github_lucas_clemente_quic_clients",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lucas-clemente/quic-clients",
        sum = "h1:/P9n0nICT/GnQJkZovtBqridjxU0ao34m7DpMts79qY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_lucas_clemente_quic_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lucas-clemente/quic-go",
        sum = "h1:9eXVRgIkMQQyiyorz/dAaOYIx3TFzXsIFkNFz4cxuJM=",
        version = "v0.28.0",
    )
    go_repository(
        name = "com_github_lucas_clemente_quic_go_certificates",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lucas-clemente/quic-go-certificates",
        sum = "h1:zqEC1GJZFbGZA0tRyNZqRjep92K5fujFtFsu5ZW7Aug=",
        version = "v0.0.0-20160823095156-d2f86524cced",
    )
    go_repository(
        name = "com_github_lucasb_eyer_go_colorful",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lucasb-eyer/go-colorful",
        sum = "h1:QIbQXiugsb+q10B+MI+7DI1oQLdmnep86tWFlaaUAac=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_lufia_iostat",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lufia/iostat",
        sum = "h1:tnCdZBIglgxD47RyD55kfWQcJMGzO+1QBziSQfesf2k=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_lukesampson_figlet",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lukesampson/figlet",
        sum = "h1:UtyD+eBVdLYSj5/pjfSR6mtnzMgIiOVcFT024G2l4CY=",
        version = "v0.0.0-20190211215653-8a3ef4a6ac42",
    )
    go_repository(
        name = "com_github_lunixbochs_vtclean",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lunixbochs/vtclean",
        sum = "h1:xu2sLAri4lGiovBDQKxl5mrXyESr3gUr5m5SM5+LVb8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_lunny_tango",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lunny/tango",
        sum = "h1:QeUe+2ksZ3LScC+SKhDbS1wbS/ctuyRnZ3fAsL10J4M=",
        version = "v0.5.6",
    )
    go_repository(
        name = "com_github_lxn_walk",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lxn/walk",
        sum = "h1:NVRJ0Uy0SOFcXSKLsS65OmI1sgCCfiDUPj+cwnH7GZw=",
        version = "v0.0.0-20210112085537-c389da54e794",
    )
    go_repository(
        name = "com_github_lxn_win",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lxn/win",
        sum = "h1:H+t6A/QJMbhCSEH5rAuRxh+CtW96g0Or0Fxa9IKr4uc=",
        version = "v0.0.0-20210218163916-a377121e959e",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_validate",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/lyft/protoc-gen-validate",
        sum = "h1:KNt/RhmQTOLr7Aj8PsJ7mTronaFyx80mRTT9qF261dA=",
        version = "v0.0.13",
    )
    go_repository(
        name = "com_github_magiconair_properties",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/magiconair/properties",
        sum = "h1:ZC2Vc7/ZFkGmsVC9KvOjumD+G5lXy2RtTKyzRKO2BQ4=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_mailru_easyjson",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mailru/easyjson",
        sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
        version = "v0.7.7",
    )
    go_repository(
        name = "com_github_makenowjust_heredoc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/MakeNowJust/heredoc",
        sum = "h1:cXCdzVdstXyiTqTvfqk9SDHpKNjxuom+DOlyEeQ4pzQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_manifoldco_promptui",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/manifoldco/promptui",
        sum = "h1:3V4HzJk1TtXW1MTZMP7mdlwbBpIinw3HztaIlYthEiA=",
        version = "v0.9.0",
    )
    go_repository(
        name = "com_github_maratori_testpackage",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/maratori/testpackage",
        sum = "h1:QtJ5ZjqapShm0w5DosRjg0PRlSdAdlx+W6cCKoALdbQ=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_marstr_guid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/marstr/guid",
        sum = "h1:/M4H/1G4avsieL6BbUwCOBzulmoeKVP5ux/3mQNnbyI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_marten_seemann_chacha20",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/marten-seemann/chacha20",
        sum = "h1:f40vqzzx+3GdOmzQoItkLX5WLvHgPgyYqFFIO5Gh4hQ=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_marten_seemann_qpack",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/marten-seemann/qpack",
        sum = "h1:jvTsT/HpCn2UZJdP+UUB53FfUUgeOyG5K1ns0OJOGVs=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_marten_seemann_qtls",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/marten-seemann/qtls",
        sum = "h1:ECsuYUKalRL240rRD4Ri33ISb7kAQ3qGDlrrl55b2pc=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_marten_seemann_qtls_go1_15",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/marten-seemann/qtls-go1-15",
        sum = "h1:RehYMOyRW8hPVEja1KBVsFVNSm35Jj9Mvs5yNoZZ28A=",
        version = "v0.1.4",
    )
    go_repository(
        name = "com_github_marten_seemann_qtls_go1_16",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/marten-seemann/qtls-go1-16",
        sum = "h1:o9JrYPPco/Nukd/HpOHMHZoBDXQqoNtUCmny98/1uqQ=",
        version = "v0.1.5",
    )
    go_repository(
        name = "com_github_marten_seemann_qtls_go1_17",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/marten-seemann/qtls-go1-17",
        sum = "h1:JADBlm0LYiVbuSySCHeY863dNkcpMmDR7s0bLKJeYlQ=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_marten_seemann_qtls_go1_18",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/marten-seemann/qtls-go1-18",
        sum = "h1:JH6jmzbduz0ITVQ7ShevK10Av5+jBEKAHMntXmIV7kM=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_marten_seemann_qtls_go1_19",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/marten-seemann/qtls-go1-19",
        sum = "h1:7m/WlWcSROrcK5NxuXaxYD32BZqe/LEEnBrWcH/cOqQ=",
        version = "v0.1.0-beta.1",
    )
    go_repository(
        name = "com_github_maruel_natural",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/maruel/natural",
        sum = "h1:PEhRT94KBTY4E0KdCYmhvDGWjSFBxc68j2M6PMRix8U=",
        version = "v0.0.0-20180416170133-dbcb3e2e8cf1",
    )
    go_repository(
        name = "com_github_marusama_semaphore_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/marusama/semaphore/v2",
        sum = "h1:Y29DhhFMvreVgoqF9EtaSJAF9t2E7Sk7i5VW81sqB8I=",
        version = "v2.4.1",
    )
    go_repository(
        name = "com_github_mastercactapus_proxyprotocol",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mastercactapus/proxyprotocol",
        sum = "h1:WpDMFKCYdF8NsoA6OrXyNKyZrzMURqqOP1PE7297RCE=",
        version = "v0.0.3",
    )
    go_repository(
        name = "com_github_masterminds_glide",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Masterminds/glide",
        sum = "h1:M5MOH04TyRiMBVeWHbifqTpnauxWINIubTCOkhXh+2g=",
        version = "v0.13.2",
    )
    go_repository(
        name = "com_github_masterminds_goutils",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Masterminds/goutils",
        sum = "h1:5nUrii3FMTL5diU80unEVvNevw1nH4+ZV4DSLVJLSYI=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_masterminds_semver",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Masterminds/semver",
        sum = "h1:H65muMkzWKEuNDnfl9d70GUjFniHKHRbFPGBuZ3QEww=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_masterminds_semver_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Masterminds/semver/v3",
        sum = "h1:RN9w6+7QoMeJVGyfmbcgs28Br8cvmnucEXnY0rYXWg0=",
        version = "v3.2.1",
    )
    go_repository(
        name = "com_github_masterminds_sprig",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Masterminds/sprig",
        sum = "h1:z4yfnGrZ7netVz+0EDJ0Wi+5VZCSYp4Z0m2dk6cEM60=",
        version = "v2.22.0+incompatible",
    )
    go_repository(
        name = "com_github_masterminds_sprig_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Masterminds/sprig/v3",
        sum = "h1:17jRggJu518dr3QaafizSXOjKYp94wKfABxUmyxvxX8=",
        version = "v3.2.2",
    )
    go_repository(
        name = "com_github_masterminds_vcs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Masterminds/vcs",
        sum = "h1:USF5TvZGYgIpcbNAEMLfFhHqP08tFZVlUVrmTSpqnyA=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_github_matoous_godox",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/matoous/godox",
        sum = "h1:RHba4YImhrUVQDHUCe2BNSOz4tVy2yGyXhvYDvxGgeE=",
        version = "v0.0.0-20190911065817-5d6d842e92eb",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:sqDoxXbdeALODt0DAeJCVp38ps9ZogZEAXjus69YV3U=",
        version = "v0.1.9",
    )
    go_repository(
        name = "com_github_mattn_go_ieproxy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-ieproxy",
        sum = "h1:qiyop7gCflfhwCzGyeT0gro3sF9AIg9HU98JORTkqfI=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:xfD0iDuEKnDkl03q4limB+vH+GxLEtL/jb4xVJSWWEY=",
        version = "v0.0.20",
    )
    go_repository(
        name = "com_github_mattn_go_runewidth",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-runewidth",
        sum = "h1:lTGmDsbAYt5DmK6OnoV7EuIF1wEIFAcxld6ypU4OSgU=",
        version = "v0.0.13",
    )
    go_repository(
        name = "com_github_mattn_go_shellwords",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-shellwords",
        sum = "h1:M2zGm7EW6UQJvDeQxo4T51eKPurbeFbe8WtebGE2xrk=",
        version = "v1.0.12",
    )
    go_repository(
        name = "com_github_mattn_go_sqlite3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-sqlite3",
        sum = "h1:jbhqpg7tQe4SupckyijYiy0mJJ/pRyHvXf7JdWK860o=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_mattn_go_tty",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-tty",
        sum = "h1:8TGB3DFRNl06DB1Q6zBX+I7FDoCUZY2fmMS9WGUIIpw=",
        version = "v0.0.0-20180219170247-931426f7535a",
    )
    go_repository(
        name = "com_github_mattn_go_xmlrpc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-xmlrpc",
        sum = "h1:Y6WEMLEsqs3RviBrAa1/7qmbGB7DVD3brZIbqMbQdGY=",
        version = "v0.0.3",
    )
    go_repository(
        name = "com_github_mattn_go_zglob",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-zglob",
        sum = "h1:xsEx/XUoVlI6yXjqBK062zYhRTZltCNmYPx6v+8DNaY=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_mattn_goveralls",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/goveralls",
        sum = "h1:7eJB6EqsPhRVxvwEXGnqdO2sJI0PTsrWoTMXEk9/OQc=",
        version = "v0.0.2",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        sum = "h1:mmDVorXM7PCGKw94cs5zkfA9PSy5pEvNWRP0ET0TIVo=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/matttproud/golang_protobuf_extensions/v2",
        sum = "h1:jWpvCLoY8Z/e3VKvlsiIGKtc+UG6U5vzxaoagmhXfyg=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_maxbrunsfeld_counterfeiter_v6",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/maxbrunsfeld/counterfeiter/v6",
        sum = "h1:g+4J5sZg6osfvEfkRZxJ1em0VT95/UOZgi/l7zi1/oE=",
        version = "v6.2.2",
    )
    go_repository(
        name = "com_github_mdlayher_ethernet",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mdlayher/ethernet",
        sum = "h1:lez6TS6aAau+8wXUP3G9I3TGlmPFEq2CTxBaRqY6AGE=",
        version = "v0.0.0-20190606142754-0394541c37b7",
    )
    go_repository(
        name = "com_github_mdlayher_ethtool",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mdlayher/ethtool",
        sum = "h1:XAWHsmKhyPOo42qq/yTPb0eFBGUKKTR1rE0dVrWVQ0Y=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_mdlayher_genetlink",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mdlayher/genetlink",
        sum = "h1:KdrNKe+CTu+IbZnm/GVUMXSqBBLqcGpRDa0xkQy56gw=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_mdlayher_netlink",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mdlayher/netlink",
        sum = "h1:/UtM3ofJap7Vl4QWCPDGXY8d3GIY2UGSDbK+QWmY8/g=",
        version = "v1.7.2",
    )
    go_repository(
        name = "com_github_mdlayher_raw",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mdlayher/raw",
        sum = "h1:aFkJ6lx4FPip+S+Uw4aTegFMct9shDvP+79PsSxpm3w=",
        version = "v0.0.0-20191009151244-50f2db8cc065",
    )
    go_repository(
        name = "com_github_mdlayher_socket",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mdlayher/socket",
        sum = "h1:eM9y2/jlbs1M615oshPQOHZzj6R6wMT7bX5NPiQvn2U=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_mdlayher_wifi",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mdlayher/wifi",
        sum = "h1:y8wYRUXwok5CtUZOXT3egghYesX0O79E3ALl+SIDm9Q=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_mesos_mesos_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mesos/mesos-go",
        sum = "h1:jMp9+W3zLu46g8EuP2su2Sjj7ipBh4N/g65c0kzGl/8=",
        version = "v0.0.11",
    )
    go_repository(
        name = "com_github_mgutz_ansi",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mgutz/ansi",
        sum = "h1:5PJl274Y63IEHC+7izoQE9x6ikvDFZS2mDVS3drnohI=",
        version = "v0.0.0-20200706080929-d51e80ef957d",
    )
    go_repository(
        name = "com_github_mholt_acmez",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mholt/acmez",
        sum = "h1:C8wsEBIUVi6e0DYoxqCcFuXtwc4AWXL/jgcDjF7mjVo=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_mholt_archiver",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mholt/archiver",
        sum = "h1:1dCVxuqs0dJseYEhi5pl7MYPH9zDa1wBi7mF09cbNkU=",
        version = "v3.1.1+incompatible",
    )
    go_repository(
        name = "com_github_mholt_caddy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mholt/caddy",
        sum = "h1:KI6RPGih2GFzWRPG8s9clKK28Ns4ZlVMKR/v7mxq6+c=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mholt_caddy_l4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mholt/caddy-l4",
        sum = "h1:APtotvfGzmVDwL1YKOKcJX2OSg3/bKOhqAiTX13rolk=",
        version = "v0.0.0-20220420174601-aec6535658b1",
    )
    go_repository(
        name = "com_github_mholt_certmagic",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mholt/certmagic",
        sum = "h1:JOUiX9IAZbbgyjNP2GY6v/6lorH+9GkZsc7ktMpGCSo=",
        version = "v0.8.3",
    )
    go_repository(
        name = "com_github_microcosm_cc_bluemonday",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/microcosm-cc/bluemonday",
        sum = "h1:SIYunPjnlXcW+gVfvm0IlSeR5U3WZUOLfVmqg85Go44=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_micromdm_scep_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/micromdm/scep/v2",
        sum = "h1:2fS9Rla7qRR266hvUoEauBJ7J6FhgssEiq2OkSKXmaU=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_microsoft_go_winio",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Microsoft/go-winio",
        sum = "h1:aPJp2QD7OOrhO5tQXqQoGSJc+DjDtWTGLOmNyAm6FgY=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_microsoft_hcsshim",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Microsoft/hcsshim",
        sum = "h1:wB06W5aYFfUB3IvootYAY2WnOmIdgPGfqSI6tufQNnY=",
        version = "v0.9.2",
    )
    go_repository(
        name = "com_github_microsoft_hcsshim_test",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Microsoft/hcsshim/test",
        sum = "h1:4FA+QBaydEHlwxg0lMN3rhwoDaQy6LKhVWR4qvq4BuA=",
        version = "v0.0.0-20210227013316-43a75bb4edd3",
    )
    go_repository(
        name = "com_github_miekg_dns",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/miekg/dns",
        sum = "h1:ca2Hdkz+cDg/7eNF6V56jjzuZ4aCAE+DbVkILdQWG/4=",
        version = "v1.1.58",
    )
    go_repository(
        name = "com_github_miekg_pkcs11",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/miekg/pkcs11",
        sum = "h1:iMwmD7I5225wv84WxIG/bmxz9AXjWvTWIbM/TYHvWtw=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_mindprince_gonvml",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mindprince/gonvml",
        sum = "h1:PS1dLCGtD8bb9RPKJrc8bS7qHL6JnW1CZvwzH9dPoUs=",
        version = "v0.0.0-20190828220739-9ebdce4bb989",
    )
    go_repository(
        name = "com_github_mistifyio_go_zfs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mistifyio/go-zfs",
        sum = "h1:aKW/4cBs+yK6gpqU3K/oIwk9Q/XICqd3zOX/UFuvqmk=",
        version = "v2.1.2-0.20190413222219-f784269be439+incompatible",
    )
    go_repository(
        name = "com_github_mitchellh_cli",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/cli",
        sum = "h1:iGBIsUe3+HZ/AD/Vd7DErOt5sU9fa8Uj7A2s1aggv1Y=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mitchellh_copystructure",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/copystructure",
        sum = "h1:vpKXTN4ewci03Vljg/q9QvCGUDttBOGBIa15WveJJGw=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_homedir",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/go-homedir",
        sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_ps",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/go-ps",
        sum = "h1:i6ampVEEF4wQFF+bkYfwYgY+F/uYJDktmvLPf7qIgjc=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_testing_interface",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/go-testing-interface",
        sum = "h1:jrgshOhYAUVNMAJiKbEu7EqAwgJJ2JqpQmpLJOu07cU=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_github_mitchellh_go_vnc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/go-vnc",
        sum = "h1:FI2NIv6fpef6BQl2u3IZX/Cj20tfypRF4yd+uaHOMtI=",
        version = "v0.0.0-20150629162542-723ed9867aed",
    )
    go_repository(
        name = "com_github_mitchellh_go_wordwrap",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/go-wordwrap",
        sum = "h1:TLuKupo69TCn6TQSyGxwI1EblZZEsQ0vMlAFQflz0v0=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_mitchellh_gox",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/gox",
        sum = "h1:lfGJxY7ToLJQjHHwi0EX6uYBdK78egf954SQl13PQJc=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_mitchellh_iochan",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/iochan",
        sum = "h1:C+X3KsSTLFVBr/tK1eYN/vs4rJcvsiLU338UhYPJWeY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:OVowDSCllw/YjdLkam3/sm7wEtOy59d8ndGgCcyj8cs=",
        version = "v1.4.3",
    )
    go_repository(
        name = "com_github_mitchellh_osext",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/osext",
        sum = "h1:2+myh5ml7lgEU/51gbeLHfKGNfgEQQIWrlbdaOsidbQ=",
        version = "v0.0.0-20151018003038-5e2d6d41470f",
    )
    go_repository(
        name = "com_github_mitchellh_pointerstructure",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/pointerstructure",
        sum = "h1:O+i9nHnXS3l/9Wu7r4NrEdwA2VFTicjUEN1uBnDo34A=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_mitchellh_reflectwalk",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/reflectwalk",
        sum = "h1:G2LzWKi524PWgd3mLHV8Y5k7s6XUvT0Gef6zxSIeXaQ=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_mmcloughlin_avo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mmcloughlin/avo",
        sum = "h1:ULR/QWMgcgRiZLUjSSJMU+fW+RDMstRdmnDWj9Q+AsA=",
        version = "v0.0.0-20200803215136-443f81d77104",
    )
    go_repository(
        name = "com_github_moby_locker",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/locker",
        sum = "h1:fOXqR41zeveg4fFODix+1Ch4mj/gT0NE1XJbp/epuBg=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_moby_spdystream",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/spdystream",
        sum = "h1:cjW1zVyyoiM0T7b6UoySUFqzXMoqRckQtXwGPiBhOM8=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_moby_sys_mountinfo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/sys/mountinfo",
        sum = "h1:2Ks8/r6lopsxWi9m58nlwjaeSzUX9iiL1vj5qB/9ObI=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_moby_sys_signal",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/sys/signal",
        sum = "h1:aDpY94H8VlhTGa9sNYUFCFsMZIUh5wm0B6XkIoJj/iY=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_moby_sys_symlink",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/sys/symlink",
        sum = "h1:tk1rOM+Ljp0nFmfOIBtlV3rTDlWOwFRhjEeAhZB0nZc=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_moby_term",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/term",
        sum = "h1:HfkjXDfhgVaN5rmueG8cL8KKeFNecRCXFhaJ2qZ5SKA=",
        version = "v0.0.0-20221205130635-1aeaba878587",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_mohae_deepcopy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mohae/deepcopy",
        sum = "h1:RWengNIwukTxcDr9M+97sNutRR1RKhG96O6jWumTTnw=",
        version = "v0.0.0-20170929034955-c48cc78d4826",
    )
    go_repository(
        name = "com_github_monochromegane_go_gitignore",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/monochromegane/go-gitignore",
        sum = "h1:n6/2gBQ3RWajuToeY6ZtZTIKv2v7ThUy5KKusIT0yc0=",
        version = "v0.0.0-20200626010858-205db1a8cc00",
    )
    go_repository(
        name = "com_github_morikuni_aec",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/morikuni/aec",
        sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_moul_http2curl",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/moul/http2curl",
        sum = "h1:dRMWoAtb+ePxMlLkrCbAqh4TlPHXvoGUSQ323/9Zahs=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mozilla_tls_observatory",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mozilla/tls-observatory",
        sum = "h1:1xJ+Xi9lYWLaaP4yB67ah0+548CD3110mCPWhVVjFkI=",
        version = "v0.0.0-20200317151703-4fa42e1c2dee",
    )
    go_repository(
        name = "com_github_mreiferson_go_httpclient",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mreiferson/go-httpclient",
        sum = "h1:oKIteTqeSpenyTrOVj5zkiyCaflLa8B+CD0324otT+o=",
        version = "v0.0.0-20160630210159-31f0106b4474",
    )
    go_repository(
        name = "com_github_mrunalp_fileutils",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mrunalp/fileutils",
        sum = "h1:NKzVxiH7eSk+OQ4M+ZYW1K6h27RUV3MI6NUTsHhU6Z4=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_munnerz_goautoneg",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/munnerz/goautoneg",
        sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
        version = "v0.0.0-20191010083416-a7dc8b61c822",
    )
    go_repository(
        name = "com_github_mwitkow_go_conntrack",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mwitkow/go-conntrack",
        sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
        version = "v0.0.0-20190716064945-2f068394615f",
    )
    go_repository(
        name = "com_github_mwitkow_go_proto_validators",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mwitkow/go-proto-validators",
        sum = "h1:F6LFfmgVnfULfaRsQWBbe7F7ocuHCr9+7m+GAeDzNbQ=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_mxk_go_flowrate",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/mxk/go-flowrate",
        sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
        version = "v0.0.0-20140419014527-cca7078d478f",
    )
    go_repository(
        name = "com_github_nakabonne_nestif",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nakabonne/nestif",
        sum = "h1:+yOViDGhg8ygGrmII72nV9B/zGxY188TYpfolntsaPw=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_namedotcom_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/namedotcom/go",
        sum = "h1:o6uBwrhM5C8Ll3MAAxrQxRHEu7FkapwTuI2WmL1rw4g=",
        version = "v0.0.0-20180403034216-08470befbe04",
    )
    go_repository(
        name = "com_github_naoina_go_stringutil",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/naoina/go-stringutil",
        sum = "h1:rCUeRUHjBjGTSHl0VC00jUPLz8/F9dDzYI70Hzifhks=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_naoina_toml",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/naoina/toml",
        sum = "h1:shk/vn9oCoOTmwcouEdwIeOtOGA/ELRUw/GwvxwfT+0=",
        version = "v0.1.2-0.20170918210437-9fafd6967416",
    )
    go_repository(
        name = "com_github_nats_io_jwt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/jwt",
        sum = "h1:+RB5hMpXUUA2dfxuhBTEkMOrYmM+gKIZYS1KjSostMI=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_nats_io_nats_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/nats.go",
        sum = "h1:ik3HbLhZ0YABLto7iX80pZLPw/6dx3T+++MZJwLnMrQ=",
        version = "v1.9.1",
    )
    go_repository(
        name = "com_github_nats_io_nats_server_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/nats-server/v2",
        sum = "h1:i2Ly0B+1+rzNZHHWtD4ZwKi+OU5l+uQo1iDHZ2PmiIc=",
        version = "v2.1.2",
    )
    go_repository(
        name = "com_github_nats_io_nkeys",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/nkeys",
        sum = "h1:6JrEfig+HzTH85yxzhSVbjHRJv9cn0p6n3IngIcM5/k=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_nats_io_nuid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/nuid",
        sum = "h1:5iA8DT8V7q8WK2EScv2padNa/rTESc1KdnPw4TC2paw=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_nbio_st",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nbio/st",
        sum = "h1:W6apQkHrMkS0Muv8G/TipAy/FJl/rCYT0+EuS8+Z0z4=",
        version = "v0.0.0-20140626010706-e9e8d9816f32",
    )
    go_repository(
        name = "com_github_nbrownus_go_metrics_prometheus",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nbrownus/go-metrics-prometheus",
        sum = "h1:8dM0ilqKL0Uzl42GABzzC4Oqlc3kGRILz0vgoff7nwg=",
        version = "v0.0.0-20210712211119-974a6260965f",
    )
    go_repository(
        name = "com_github_nbutton23_zxcvbn_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nbutton23/zxcvbn-go",
        sum = "h1:AREM5mwr4u1ORQBMvzfzBgpsctsbQikCVpvC+tX285E=",
        version = "v0.0.0-20180912185939-ae427f1e4c1d",
    )
    go_repository(
        name = "com_github_ncw_swift",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ncw/swift",
        sum = "h1:4DQRPj35Y41WogBxyhOXlrI37nzGlyEcsforeudyYPQ=",
        version = "v1.0.47",
    )
    go_repository(
        name = "com_github_neelance_astrewrite",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/neelance/astrewrite",
        sum = "h1:D6paGObi5Wud7xg83MaEFyjxQB1W5bz5d0IFppr+ymk=",
        version = "v0.0.0-20160511093645-99348263ae86",
    )
    go_repository(
        name = "com_github_neelance_sourcemap",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/neelance/sourcemap",
        sum = "h1:eFXv9Nu1lGbrNbj619aWwZfVF5HBrm9Plte8aNptuTI=",
        version = "v0.0.0-20151028013722-8c68805598ab",
    )
    go_repository(
        name = "com_github_newrelic_go_agent",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/newrelic/go-agent",
        sum = "h1:IB0Fy+dClpBq9aEoIrLyQXzU34JyI1xVTanPLB/+jvU=",
        version = "v2.15.0+incompatible",
    )
    go_repository(
        name = "com_github_ngdinhtoan_glide_cleanup",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ngdinhtoan/glide-cleanup",
        sum = "h1:kN4sV+0tp2F1BvwU+5SfNRMDndRmvIfnI3kZ7B8Yv4Y=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_niemeyer_pretty",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/niemeyer/pretty",
        sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
        version = "v0.0.0-20200227124842-a10e7caefd8e",
    )
    go_repository(
        name = "com_github_nightlyone_lockfile",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nightlyone/lockfile",
        sum = "h1:+2OJrU8cmOstEoh0uQvYemRGVH1O6xtO2oANUWHFnP0=",
        version = "v0.0.0-20180618180623-0ad87eef1443",
    )
    go_repository(
        name = "com_github_nishanths_predeclared",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nishanths/predeclared",
        sum = "h1:3f0nxAmdj/VoCGN/ijdMy7bj6SBagaqYg1B0hu8clMA=",
        version = "v0.0.0-20200524104333-86fad755b4d3",
    )
    go_repository(
        name = "com_github_nkovacs_streamquote",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nkovacs/streamquote",
        sum = "h1:E2B8qYyeSgv5MXpmzZXRNp8IAQ4vjxIjhpAf5hv/tAg=",
        version = "v0.0.0-20170412213628-49af9bddb229",
    )
    go_repository(
        name = "com_github_nrdcg_auroradns",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nrdcg/auroradns",
        sum = "h1:m/kBq83Xvy3cU261MOknd8BdnOk12q4lAWM+kOdsC2Y=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_nrdcg_dnspod_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nrdcg/dnspod-go",
        sum = "h1:c/jn1mLZNKF3/osJ6mz3QPxTudvPArXTjpkmYj0uK6U=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_nrdcg_goinwx",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nrdcg/goinwx",
        sum = "h1:AJnjoWPELyCtofhGcmzzcEMFd9YdF2JB/LgutWsWt/s=",
        version = "v0.6.1",
    )
    go_repository(
        name = "com_github_nrdcg_namesilo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nrdcg/namesilo",
        sum = "h1:kLjCjsufdW/IlC+iSfAqj0iQGgKjlbUUeDJio5Y6eMg=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_nwaples_rardecode",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nwaples/rardecode",
        sum = "h1:r7vGuS5akxOnR4JQSkko62RJ1ReCMXxQRPtxsiFMBOs=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_nxadm_tail",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/nxadm/tail",
        sum = "h1:nPr65rt6Y5JFSKQO7qToXr7pePgD6Gwiw05lkbyAQTE=",
        version = "v1.4.8",
    )
    go_repository(
        name = "com_github_nytimes_gziphandler",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/NYTimes/gziphandler",
        sum = "h1:ZUDjpQae29j0ryrS0u/B8HZfJBtBQHjqw2rQ2cqUQ3I=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_oklog_oklog",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/oklog/oklog",
        sum = "h1:wVfs8F+in6nTBMkA7CbRw+zZMIB7nNM825cM1wuzoTk=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_oklog_run",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/oklog/run",
        sum = "h1:Ru7dDtJNOyC66gQ5dQmaCa0qIsAUFY3sFpK1Xk8igrw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_oklog_ulid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/oklog/ulid",
        sum = "h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_oleiade_reflections",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/oleiade/reflections",
        sum = "h1:I6mXuorHlvwNDFelz7a+j0HaGYSzX7+Gq60DqLVypfc=",
        version = "v0.0.0-20160817071559-0e86b3c98b2f",
    )
    go_repository(
        name = "com_github_olekukonko_tablewriter",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/olekukonko/tablewriter",
        sum = "h1:P2Ga83D34wi1o9J6Wh1mRuqd4mF/x/lgBS7N7AbDhec=",
        version = "v0.0.5",
    )
    go_repository(
        name = "com_github_oneofone_xxhash",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/OneOfOne/xxhash",
        sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_onsi_ginkgo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/onsi/ginkgo",
        sum = "h1:29JGrr5oVBm5ulCWet69zQkzWipVXIol6ygQUe/EzNc=",
        version = "v1.16.4",
    )
    go_repository(
        name = "com_github_onsi_ginkgo_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/onsi/ginkgo/v2",
        sum = "h1:0jY9lJquiL8fcf3M4LAXN5aMlS/b2BV86HFFPCPMgE4=",
        version = "v2.13.0",
    )
    go_repository(
        name = "com_github_onsi_gomega",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/onsi/gomega",
        sum = "h1:KIA/t2t5UBzoirT4H9tsML45GEbo3ouUnBHsCfD2tVg=",
        version = "v1.29.0",
    )
    go_repository(
        name = "com_github_op_go_logging",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/op/go-logging",
        sum = "h1:lDH9UUVJtmYCjyT0CI4q8xvlXPxeZ0gYCVvWbmPlp88=",
        version = "v0.0.0-20160315200505-970db520ece7",
    )
    go_repository(
        name = "com_github_opencontainers_go_digest",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/opencontainers/go-digest",
        sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_opencontainers_image_spec",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/opencontainers/image-spec",
        sum = "h1:9yCKha/T5XdGtO0q9Q9a6T5NUCsTn/DrBg0D7ufOcFM=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_opencontainers_runc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/opencontainers/runc",
        sum = "h1:O9+X96OcDjkmmZyfaG996kV7yq8HsoU2h1XRRQcefG8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_opencontainers_runtime_spec",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/opencontainers/runtime-spec",
        sum = "h1:3snG66yBm59tKhhSPQrQ/0bCrv1LQbKt40LnUPiUxdc=",
        version = "v1.0.3-0.20210326190908-1c3f411f0417",
    )
    go_repository(
        name = "com_github_opencontainers_runtime_tools",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/opencontainers/runtime-tools",
        sum = "h1:H7DMc6FAjgwZZi8BRqjrAAHWoqEr5e5L6pS4V0ezet4=",
        version = "v0.0.0-20181011054405-1d69bd0f9c39",
    )
    go_repository(
        name = "com_github_opencontainers_selinux",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/opencontainers/selinux",
        sum = "h1:+5Zbo97w3Lbmb3PeqQtpmTkMwsW5nRI3YaLpt7tQ7oU=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_github_opendns_vegadns2client",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/OpenDNS/vegadns2client",
        sum = "h1:xPMsUicZ3iosVPSIP7bW5EcGUzjiiMl1OYTe14y/R24=",
        version = "v0.0.0-20180418235048-a3fa4a771d87",
    )
    go_repository(
        name = "com_github_openpeedeep_depguard",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/OpenPeeDeeP/depguard",
        sum = "h1:VlW4R6jmBIv3/u1JNlawEvJMM4J+dPORPaZasQee8Us=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_opentracing_basictracer_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/opentracing/basictracer-go",
        sum = "h1:YyUAhaEfjoWXclZVJ9sGoNct7j4TVk7lZWlQw5UXuoo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_opentracing_contrib_go_observer",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/opentracing-contrib/go-observer",
        sum = "h1:lM6RxxfUMrYL/f8bWEUqdXrANWtrL7Nndbm9iFN0DlU=",
        version = "v0.0.0-20170622124052-a52f23424492",
    )
    go_repository(
        name = "com_github_opentracing_opentracing_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/opentracing/opentracing-go",
        sum = "h1:uEJPy/1a5RIPAJ0Ov+OIO8OxWu77jEv+1B0VhjKrZUs=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_openzipkin_contrib_zipkin_go_opentracing",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/openzipkin-contrib/zipkin-go-opentracing",
        sum = "h1:ZCnq+JUrvXcDVhX/xRolRBZifmabN1HcS1wrPSvxhrU=",
        version = "v0.4.5",
    )
    go_repository(
        name = "com_github_openzipkin_zipkin_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/openzipkin/zipkin-go",
        sum = "h1:nY8Hti+WKaP0cRsSeQ026wU03QsM762XBeCXBb9NAWI=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_oracle_oci_go_sdk",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/oracle/oci-go-sdk",
        sum = "h1:oj5ESjXwwkFRdhZSnPlShvLWYdt/IZ65RQxveYM3maA=",
        version = "v7.0.0+incompatible",
    )
    go_repository(
        name = "com_github_orcaman_concurrent_map",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/orcaman/concurrent-map",
        sum = "h1:lNCW6THrCKBiJBpz8kbVGjC7MgdCGKwuvBgc7LoD6sw=",
        version = "v0.0.0-20190826125027-8c72a8bb44f6",
    )
    go_repository(
        name = "com_github_oschwald_geoip2_golang",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/oschwald/geoip2-golang",
        sum = "h1:JW1r5AKi+vv2ujSxjKthySK3jo8w8oKWPyXsw+Qs/S8=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_oschwald_maxminddb_golang",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/oschwald/maxminddb-golang",
        sum = "h1:tIk4nv6VT9OiPyrnDAfJS1s1xKDQMZOsGojab6EjC1Y=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_otiai10_copy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/otiai10/copy",
        sum = "h1:HvG945u96iNadPoG2/Ja2+AUJeW5YuFQMixq9yirC+k=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_otiai10_curr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/otiai10/curr",
        sum = "h1:TJIWdbX0B+kpNagQrjgq8bCMrbhiuX73M2XwgtDMoOI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_otiai10_mint",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/otiai10/mint",
        sum = "h1:BCmzIS3n71sGfHB5NMNDB3lHYPz8fWSkCAErHed//qc=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_ovh_go_ovh",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ovh/go-ovh",
        sum = "h1:37VE5TYj2m/FLA9SNr4z0+A0JefvTmR60Zwf8XSEV7c=",
        version = "v0.0.0-20181109152953-ba5adb4cf014",
    )
    go_repository(
        name = "com_github_p4gefau1t_trojan_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/p4gefau1t/trojan-go",
        sum = "h1:KhXCKH7Rpd+1YYABK0qcN7s8146Kfatzk5SCy8nAFAI=",
        version = "v0.8.2",
    )
    go_repository(
        name = "com_github_pact_foundation_pact_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pact-foundation/pact-go",
        sum = "h1:OYkFijGHoZAYbOIb1LWXrwKQbMMRUv1oQ89blD2Mh2Q=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_pascaldekloe_goe",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pascaldekloe/goe",
        sum = "h1:cBOtyMzM9HTpWjXfbbunk26uA6nG3a8n06Wieeh0MwY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_patrickmn_go_cache",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/patrickmn/go-cache",
        sum = "h1:HRMgzkcYKYpi3C8ajMPV8OFXaaRUnok+kx1WdO15EQc=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "com_github_pborman_uuid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pborman/uuid",
        sum = "h1:J7Q5mO4ysT1dv8hyrUGHb9+ooztCXu1D8MY8DZYsu3g=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_pelletier_go_buffruneio",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pelletier/go-buffruneio",
        sum = "h1:U4t4R6YkofJ5xHm3dJzuRpPZ0mr5MMCoAWooScCR7aA=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_pelletier_go_toml",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pelletier/go-toml",
        sum = "h1:zeC5b1GviRUyKYd6OJPvBU/mcVDVoL1OhT17FCt5dSQ=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_github_performancecopilot_speed",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/performancecopilot/speed",
        sum = "h1:2WnRzIquHa5QxaJKShDkLM+sc0JPuwhXzK8OYOyt3Vg=",
        version = "v3.0.0+incompatible",
    )
    go_repository(
        name = "com_github_peterbourgon_diskv",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/peterbourgon/diskv",
        sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
        version = "v2.0.1+incompatible",
    )
    go_repository(
        name = "com_github_peterh_liner",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/peterh/liner",
        sum = "h1:oYW+YCJ1pachXTQmzR3rNLYGGz4g/UgFcjb28p/viDM=",
        version = "v1.1.1-0.20190123174540-a2c9a5303de7",
    )
    go_repository(
        name = "com_github_petermattis_goid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/petermattis/goid",
        sum = "h1:q2e307iGHPdTGp0hoxKjt1H5pDo6utceo3dQVK3I5XQ=",
        version = "v0.0.0-20180202154549-b0b1615b78e5",
    )
    go_repository(
        name = "com_github_phayes_checkstyle",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/phayes/checkstyle",
        sum = "h1:CdDQnGF8Nq9ocOS/xlSptM1N3BbrA6/kmaep5ggwaIA=",
        version = "v0.0.0-20170904204023-bfd46e6a821d",
    )
    go_repository(
        name = "com_github_phayes_freeport",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/phayes/freeport",
        sum = "h1:JhzVVoYvbOACxoUmOs6V/G4D5nPVUW73rKvXxP4XUJc=",
        version = "v0.0.0-20180830031419-95f893ade6f2",
    )
    go_repository(
        name = "com_github_philhofer_fwd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/philhofer/fwd",
        sum = "h1:UbZqGr5Y38ApvM/V/jEljVxwocdweyH+vmYvRPBnbqQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_phuslu_geoip",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/phuslu/geoip",
        sum = "h1:pap5n0dO6f2HUOXKGW0OrG0Y9OlxN0uC+XKMvziUm6g=",
        version = "v1.0.20200217",
    )
    go_repository(
        name = "com_github_pierrec_lz4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pierrec/lz4",
        sum = "h1:Ix9yFKn1nSPBLFl/yZknTp8TU5G4Ps0JDmguYK6iH1A=",
        version = "v2.6.0+incompatible",
    )
    go_repository(
        name = "com_github_pires_go_proxyproto",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pires/go-proxyproto",
        sum = "h1:A4Jv4ZCaV3AFJeGh5mGwkz4iuWUYMlQ7IoO/GTuSuLo=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_pkg_diff",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/diff",
        sum = "h1:aoZm08cpOy4WuID//EZDgcC4zIxODThtZNPirFr42+A=",
        version = "v0.0.0-20210226163009-20ebb0f2a09e",
    )
    go_repository(
        name = "com_github_pkg_errors",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pkg_profile",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/profile",
        sum = "h1:F++O52m40owAmADcojzM+9gyjmMOY/T4oYJkgFDH8RE=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_pkg_sftp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/sftp",
        sum = "h1:VasscCm72135zRysgrJDKsntdmPN+OuU3+nnHYA9wyc=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_portainer_agent",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/portainer/agent",
        sum = "h1:0tLPXwCJFT6twnJ/tUGQCBk9ME4eAPpBbhg1tm9iDgI=",
        version = "v0.0.0-20220708094808-9c896747b4ff",
    )
    go_repository(
        name = "com_github_portainer_docker_compose_wrapper",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/portainer/docker-compose-wrapper",
        sum = "h1:GFTn2e5AyIoBuK6hXbdVNkuV2m450DQnYmgQDZRU3x8=",
        version = "v0.0.0-20220708023447-a69a4ebaa021",
    )
    go_repository(
        name = "com_github_portainer_libcrypto",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/portainer/libcrypto",
        sum = "h1:qY8TbocN75n5PDl16o0uVr5MevtM5IhdwSelXEd4nFM=",
        version = "v0.0.0-20210422035235-c652195c5c3a",
    )
    go_repository(
        name = "com_github_portainer_libhelm",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/portainer/libhelm",
        sum = "h1:5e8KAnDa2G3cEHK7aV/ue8lOaoQwBZUzoALslwWkR04=",
        version = "v0.0.0-20210929000907-825e93d62108",
    )
    go_repository(
        name = "com_github_portainer_libhttp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/portainer/libhttp",
        sum = "h1:GMIjRVV2LADpJprPG2+8MdRH6XvrFgC7wHm7dFUdOpc=",
        version = "v0.0.0-20211208103139-07a5f798eb3f",
    )
    go_repository(
        name = "com_github_portainer_portainer_api",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/portainer/portainer/api",
        sum = "h1:gUe/VH30JLUrr9cc6wM6ByWB6rZJVBXql/E+A9e25EM=",
        version = "v0.0.0-20220303203420-547d9c2fde15",
    )
    go_repository(
        name = "com_github_posener_complete",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/posener/complete",
        sum = "h1:GqpA1/5oN1NgsxoSA4RH0YWTaqvUlQNeOpHXD/JRbOQ=",
        version = "v1.2.2-0.20190308074557-af07aa5181b3",
    )
    go_repository(
        name = "com_github_pquerna_cachecontrol",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pquerna/cachecontrol",
        sum = "h1:yJMy84ti9h/+OEWa752kBTKv4XC30OtVVHYv/8cTqKc=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_pquerna_ffjson",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pquerna/ffjson",
        sum = "h1:xoIK0ctDddBMnc74udxJYBqlo9Ylnsp1waqjLsnef20=",
        version = "v0.0.0-20190930134022-aa0246cd15f7",
    )
    go_repository(
        name = "com_github_pquerna_otp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pquerna/otp",
        sum = "h1:TBZrpfnzVbgmpYhiYBK+bJ4Ig0+ye+GGNMe2pTrvxCo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_prashantv_gostub",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/prashantv/gostub",
        sum = "h1:BTyx3RfQjRHnUWaGF9oQos79AlQ5k8WNktv7VGvVH4g=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_prometheus_client_golang",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/client_golang",
        sum = "h1:HzFfmkOzH5Q8L8G+kSJKUx5dtG87sewO+FoDDqP5Tbk=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:VQw1hfvPvk3Uv6Qf29VrPF32JB6rtbgI6cYPYQjL0Qw=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_prometheus_common",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/common",
        sum = "h1:doXzt5ybi1HBKpsZOL0sSkaNHJJqkyfEWZGGqqScV0Y=",
        version = "v0.46.0",
    )
    go_repository(
        name = "com_github_prometheus_community_go_runit",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus-community/go-runit",
        sum = "h1:uTWEj/Fn2RoLdfg/etSqwzgYNOYPrARx1BHUN052tGA=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_prometheus_exporter_toolkit",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/exporter-toolkit",
        sum = "h1:yOAzZTi4M22ZzVxD+fhy1URTuNRj/36uQJJ5S8IPza8=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_prometheus_node_exporter",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/node_exporter",
        sum = "h1:7MVpSdfWrThNo0SlldhUyAVFZ7LWbC9+QJRzB4QmkE8=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:jluTpSng7V9hY0O2R9DzzJHYb2xULk9VTR1V1R/k6Bo=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_github_prometheus_tsdb",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/tsdb",
        sum = "h1:YZcsG11NqnK4czYLrWd9mpEuAJIHVQLwdrleYfszMAA=",
        version = "v0.7.1",
    )
    go_repository(
        name = "com_github_pseudomuto_protoc_gen_doc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pseudomuto/protoc-gen-doc",
        sum = "h1:aNTZq0dy0Pq2ag2v7bhNKFNgBBA8wMCoJSChhd7RciE=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_pseudomuto_protokit",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/pseudomuto/protokit",
        sum = "h1:hlnBDcy3YEDXH7kc9gV+NLaN0cDzhDvD1s7Y6FZ8RpM=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_puerkitobio_purell",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/PuerkitoBio/purell",
        sum = "h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_puerkitobio_urlesc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/PuerkitoBio/urlesc",
        sum = "h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=",
        version = "v0.0.0-20170810143723-de5bf2ad4578",
    )
    go_repository(
        name = "com_github_qri_io_jsonpointer",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/qri-io/jsonpointer",
        sum = "h1:C8RRfIlExwwrXw28G8LkrpWiHUVT4uLowfvjUYJ2Iec=",
        version = "v0.0.0-20180309164927-168dd9e45cf2",
    )
    go_repository(
        name = "com_github_qri_io_jsonschema",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/qri-io/jsonschema",
        sum = "h1:vwTGeGWCew89DI4ZwKCaobGAN7ExvZiBzgn4LZHMVOc=",
        version = "v0.0.0-20180607150648-d0d3b10ec792",
    )
    go_repository(
        name = "com_github_quasilyte_go_consistent",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/quasilyte/go-consistent",
        sum = "h1:JoUA0uz9U0FVFq5p4LjEq4C0VgQ0El320s3Ms0V4eww=",
        version = "v0.0.0-20190521200055-c6f3937de18c",
    )
    go_repository(
        name = "com_github_quasilyte_go_ruleguard",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/quasilyte/go-ruleguard",
        sum = "h1:DvnesvLtRPQOvaUbfXfh0tpMHg29by0H7F2U+QIkSu8=",
        version = "v0.1.2-0.20200318202121-b00d7a75d3d8",
    )
    go_repository(
        name = "com_github_quic_go_qpack",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/quic-go/qpack",
        sum = "h1:Cr9BXA1sQS2SmDUWjSofMPNKmvF6IiIfDRmgU0w1ZCo=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_quic_go_qtls_go1_20",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/quic-go/qtls-go1-20",
        sum = "h1:MfFAPULvst4yoMgY9QmtpYmfij/em7O8UUi+bNVm7Cg=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_quic_go_quic_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/quic-go/quic-go",
        sum = "h1:aD8MmHfgqTURWNJy48IYFg2OnxwHT3JL7ahGs73lb4k=",
        version = "v0.41.0",
    )
    go_repository(
        name = "com_github_quic_go_webtransport_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/quic-go/webtransport-go",
        sum = "h1:CvNsKqc4W2HljHJnoT+rMmbRJybShZ0YPFDD3NxaZLY=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_rainycape_memcache",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rainycape/memcache",
        sum = "h1:dq90+d51/hQRaHEqRAsQ1rE/pC1GUS4sc2rCbbFsAIY=",
        version = "v0.0.0-20150622160815-1031fa0ce2f2",
    )
    go_repository(
        name = "com_github_rakyll_statik",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rakyll/statik",
        sum = "h1:OF3QCZUuyPxuGEP7B4ypUa7sB/iHtqOTDYZXGM8KOdQ=",
        version = "v0.1.7",
    )
    go_repository(
        name = "com_github_rancher_dapper",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rancher/dapper",
        sum = "h1:HtlWLY7MIbBsr+1ei4m4eixIV0+TXS7Kwx+c8Y7T+mE=",
        version = "v0.5.5",
    )
    go_repository(
        name = "com_github_rcrowley_go_metrics",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rcrowley/go-metrics",
        sum = "h1:N/ElC8H3+5XpJzTSTfLsJV/mx9Q9g7kxmchpfZyxgzM=",
        version = "v0.0.0-20201227073835-cf1acfcdf475",
    )
    go_repository(
        name = "com_github_refraction_networking_utls",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/refraction-networking/utls",
        sum = "h1:vIkvetWOJZSADSKCF9MLTsQNW2httdBmYz47dQQteP8=",
        version = "v0.0.0-20200601200209-ada0bb9b38a0",
    )
    go_repository(
        name = "com_github_remyoudompheng_bigfft",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/remyoudompheng/bigfft",
        sum = "h1:/NRJ5vAYoqz+7sG51ubIDHXeWO8DlTSrToPu6q11ziA=",
        version = "v0.0.0-20170806203942-52369c62f446",
    )
    go_repository(
        name = "com_github_rican7_retry",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Rican7/retry",
        sum = "h1:FqK94z34ly8Baa6K+G8Mmza9rYWTKOJk+yckIBB5qVk=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_riobard_go_bloom",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/riobard/go-bloom",
        sum = "h1:f/FNXud6gA3MNr8meMVVGxhp+QBTqY91tM8HjEuMjGg=",
        version = "v0.0.0-20200614022211-cdc8013cb5b3",
    )
    go_repository(
        name = "com_github_rivo_uniseg",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rivo/uniseg",
        sum = "h1:S1pD9weZBuJdFmowNwbpi7BJ8TNftyUImj/0WQi72jY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_rjeczalik_notify",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rjeczalik/notify",
        sum = "h1:CLCKso/QK1snAlnhNR/CNvNiFU2saUtjV0bx3EwNeCE=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_rkl_digest",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rkl-/digest",
        sum = "h1:rDj3WeO+TiWyxfcydUnKegWAZoR5kQsnW0wzhggdOrw=",
        version = "v0.0.0-20180419075440-8316caa4a777",
    )
    go_repository(
        name = "com_github_robfig_cron_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/robfig/cron/v3",
        sum = "h1:WdRxkvbJztn8LMz/QEvLN5sBU+xKpSqwwUO1Pjr4qDs=",
        version = "v3.0.1",
    )
    go_repository(
        name = "com_github_rogpeppe_fastuuid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rogpeppe/fastuuid",
        sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:TMyTOH3F/DB16zRVcYyreMH6GnZZrwQVAoYjRBZyWFQ=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_rs_cors",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rs/cors",
        sum = "h1:L0uuZVXIKlI1SShY2nhFfo44TYvDPQ1w4oFkUJNfhyo=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_rs_xid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rs/xid",
        sum = "h1:mhH9Nq+C1fY2l1XIpgxIiUOfNpRBYH1kKcr+qfKgjRc=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_rs_zerolog",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rs/zerolog",
        sum = "h1:uPRuwkWF4J6fGsJ2R0Gn2jB1EQiav9k3S6CSdygQJXY=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_github_rubiojr_go_vhd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/rubiojr/go-vhd",
        sum = "h1:ht7N4d/B7Ezf58nvMNVF3OlvDlz9pp+WHVcRNS0nink=",
        version = "v0.0.0-20160810183302-0bfd3b39853c",
    )
    go_repository(
        name = "com_github_russross_blackfriday",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/russross/blackfriday",
        sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_russross_blackfriday_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/russross/blackfriday/v2",
        sum = "h1:JIOH55/0cWyOuilr9/qlrm0BSXldqnqwMsf35Ld67mk=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_ryancurrah_gomodguard",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ryancurrah/gomodguard",
        sum = "h1:DWbye9KyMgytn8uYpuHkwf0RHqAYO6Ay/D0TbCpPtVU=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_ryanuber_columnize",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ryanuber/columnize",
        sum = "h1:j1Wcmh8OrK4Q7GXY+V7SVSY8nUWQxHW5TkBe7YUl+2s=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "com_github_ryanuber_go_glob",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ryanuber/go-glob",
        sum = "h1:iQh3xXAumdQ+4Ufa5b25cRpC5TYKlno6hsv6Cb3pkBk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_sacloud_libsacloud",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/sacloud/libsacloud",
        sum = "h1:td3Kd7lvpSAxxHEVpnaZ9goHmmhi0D/RfP0Rqqf/kek=",
        version = "v1.26.1",
    )
    go_repository(
        name = "com_github_safchain_ethtool",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/safchain/ethtool",
        sum = "h1:gimQJpsI6sc1yIqP/y8GYgiXn/NjgvpM0RNoWLVVmP0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_samfoo_ansi",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/samfoo/ansi",
        sum = "h1:CmSpbxmewNQbzqztaY0bke1qzHhyNyC29wYgh17Gxfo=",
        version = "v0.0.0-20160124022901-b6bd2ded7189",
    )
    go_repository(
        name = "com_github_samuel_go_zookeeper",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/samuel/go-zookeeper",
        sum = "h1:p3Vo3i64TCLY7gIfzeQaUJ+kppEO5WQG3cL8iE8tGHU=",
        version = "v0.0.0-20190923202752-2cc03de413da",
    )
    go_repository(
        name = "com_github_sasha_s_go_deadlock",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/sasha-s/go-deadlock",
        sum = "h1:T7hUw7pBSINuHQyWwMdfIWZZH5M3ju4yXIbuV/Upp+4=",
        version = "v0.0.0-20180226215254-237a9547c8a5",
    )
    go_repository(
        name = "com_github_sassoftware_go_rpmutils",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/sassoftware/go-rpmutils",
        sum = "h1:+gCnWOZV8Z/8jehJ2CdqB47Z3S+SREmQcuXkRFLNsiI=",
        version = "v0.0.0-20190420191620-a8f1baeba37b",
    )
    go_repository(
        name = "com_github_satori_go_uuid",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/satori/go.uuid",
        sum = "h1:0uYX9dsZ2yD7q2RtLRtPSdGDWzjeM3TbMJP9utgA0ww=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_sclevine_agouti",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/sclevine/agouti",
        sum = "h1:8IBJS6PWz3uTlMP3YBIR5f+KAldcGuOeFkFbUWfBgK4=",
        version = "v3.0.0+incompatible",
    )
    go_repository(
        name = "com_github_sclevine_spec",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/sclevine/spec",
        sum = "h1:1Jwdf9jSfDl9NVmt8ndHqbTZ7XCCPbh1jI3hkDBHVYA=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_sean_seed",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/sean-/seed",
        sum = "h1:nn5Wsu0esKSJiIVhscUtVbo7ada43DJhG55ua/hjS5I=",
        version = "v0.0.0-20170313163322-e2103e2c3529",
    )
    go_repository(
        name = "com_github_seandolphin_bqschema",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/SeanDolphin/bqschema",
        sum = "h1:iCYFd5Qsw6caM2k5/SsITSL9+3kQCr+oz6pnNjWTq90=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_seccomp_libseccomp_golang",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/seccomp/libseccomp-golang",
        sum = "h1:58EBmR2dMNL2n/FnbQewK3D14nXr0V9CObDSvMJLq+Y=",
        version = "v0.9.2-0.20210429002308-3879420cc921",
    )
    go_repository(
        name = "com_github_securego_gosec",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/securego/gosec",
        sum = "h1:rq2/kILQnPtq5oL4+IAjgVOjh5e2yj2aaCYi7squEvI=",
        version = "v0.0.0-20200401082031-e946c8c39989",
    )
    go_repository(
        name = "com_github_securego_gosec_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/securego/gosec/v2",
        sum = "h1:y/9mCF2WPDbSDpL3QDWZD3HHGrSYw0QSHnCqTfs4JPE=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_seiflotfy_cuckoofilter",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/seiflotfy/cuckoofilter",
        sum = "h1:pqy40B3MQWYrza7YZXOXgl0Nf0QGFqrOC0BKae1UNAA=",
        version = "v0.0.0-20201222105146-bc6005554a0c",
    )
    go_repository(
        name = "com_github_sereal_sereal",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Sereal/Sereal",
        sum = "h1:3trCIB5GsAOIY8NxlfMztCYIhVsW9V5sZ+brsecjaiI=",
        version = "v0.0.0-20190430203904-6faf9605eb56",
    )
    go_repository(
        name = "com_github_sergi_go_diff",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/sergi/go-diff",
        sum = "h1:XU+rvMAioB0UC3q1MFrIQy4Vo5/4VsRDQQXHsEya6xQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_shadowsocks_go_shadowsocks2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shadowsocks/go-shadowsocks2",
        sum = "h1:PDSQv9y2S85Fl7VBeOMF9StzeXZyK1HakRm86CUbr28=",
        version = "v0.1.5",
    )
    go_repository(
        name = "com_github_shadowsocks_shadowsocks_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shadowsocks/shadowsocks-go",
        sum = "h1:XU9hik0exChEmY92ALW4l9WnDodxLVS9yOSNh2SizaQ=",
        version = "v0.0.0-20200409064450-3e585ff90601",
    )
    go_repository(
        name = "com_github_shirou_gopsutil",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shirou/gopsutil",
        sum = "h1:Bn1aCHHRnjv4Bl16T8rcaFjYSrGrIZvpiGO6P3Q4GpU=",
        version = "v3.21.4-0.20210419000835-c7a38de76ee5+incompatible",
    )
    go_repository(
        name = "com_github_shirou_w32",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shirou/w32",
        sum = "h1:udFKJ0aHUL60LboW/A+DfgoHVedieIzIXE8uylPue0U=",
        version = "v0.0.0-20160930032740-bb4de0191aa4",
    )
    go_repository(
        name = "com_github_shopify_logrus_bugsnag",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Shopify/logrus-bugsnag",
        sum = "h1:UrqY+r/OJnIp5u0s1SbQ8dVfLCZJsnvazdBP5hS4iRs=",
        version = "v0.0.0-20171204204709-577dee27f20d",
    )
    go_repository(
        name = "com_github_shopify_sarama",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Shopify/sarama",
        sum = "h1:lOi3SfE6OcFlW9Trgtked2aHNZ2BIG/d6Do+PEUAqqM=",
        version = "v1.28.0",
    )
    go_repository(
        name = "com_github_shopify_toxiproxy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/Shopify/toxiproxy",
        sum = "h1:TKdv8HiTLgE5wdJuEML90aBgNWsokNbMijUGhmcoBJc=",
        version = "v2.1.4+incompatible",
    )
    go_repository(
        name = "com_github_shopspring_decimal",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shopspring/decimal",
        sum = "h1:abSATXmQEYyShuxI4/vyW3tV1MrKAJzCZ/0zLUXYbsQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_shurcool_component",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/component",
        sum = "h1:Fth6mevc5rX7glNLpbAMJnqKlfIkcTjZCSHEeqvKbcI=",
        version = "v0.0.0-20170202220835-f88ec8f54cc4",
    )
    go_repository(
        name = "com_github_shurcool_events",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/events",
        sum = "h1:vabduItPAIz9px5iryD5peyx7O3Ya8TBThapgXim98o=",
        version = "v0.0.0-20181021180414-410e4ca65f48",
    )
    go_repository(
        name = "com_github_shurcool_github_flavored_markdown",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/github_flavored_markdown",
        sum = "h1:qb9IthCFBmROJ6YBS31BEMeSYjOscSiG+EO+JVNTz64=",
        version = "v0.0.0-20181002035957-2122de532470",
    )
    go_repository(
        name = "com_github_shurcool_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/go",
        sum = "h1:MZM7FHLqUHYI0Y/mQAt3d2aYa0SiNms/hFqC9qJYolM=",
        version = "v0.0.0-20180423040247-9e1955d9fb6e",
    )
    go_repository(
        name = "com_github_shurcool_go_goon",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/go-goon",
        sum = "h1:llrF3Fs4018ePo4+G/HV/uQUqEI1HMDjCeOf2V6puPc=",
        version = "v0.0.0-20170922171312-37c2f522c041",
    )
    go_repository(
        name = "com_github_shurcool_gofontwoff",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/gofontwoff",
        sum = "h1:Yoy/IzG4lULT6qZg62sVC+qyBL8DQkmD2zv6i7OImrc=",
        version = "v0.0.0-20180329035133-29b52fc0a18d",
    )
    go_repository(
        name = "com_github_shurcool_gopherjslib",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/gopherjslib",
        sum = "h1:UOk+nlt1BJtTcH15CT7iNO7YVWTfTv/DNwEAQHLIaDQ=",
        version = "v0.0.0-20160914041154-feb6d3990c2c",
    )
    go_repository(
        name = "com_github_shurcool_highlight_diff",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/highlight_diff",
        sum = "h1:vYEG87HxbU6dXj5npkeulCS96Dtz5xg3jcfCgpcvbIw=",
        version = "v0.0.0-20170515013008-09bb4053de1b",
    )
    go_repository(
        name = "com_github_shurcool_highlight_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/highlight_go",
        sum = "h1:7pDq9pAMCQgRohFmd25X8hIH8VxmT3TaDm+r9LHxgBk=",
        version = "v0.0.0-20181028180052-98c3abbbae20",
    )
    go_repository(
        name = "com_github_shurcool_home",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/home",
        sum = "h1:MPblCbqA5+z6XARjScMfz1TqtJC7TuTRj0U9VqIBs6k=",
        version = "v0.0.0-20181020052607-80b7ffcb30f9",
    )
    go_repository(
        name = "com_github_shurcool_htmlg",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/htmlg",
        sum = "h1:crYRwvwjdVh1biHzzciFHe8DrZcYrVcZFlJtykhRctg=",
        version = "v0.0.0-20170918183704-d01228ac9e50",
    )
    go_repository(
        name = "com_github_shurcool_httperror",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/httperror",
        sum = "h1:eHRtZoIi6n9Wo1uR+RU44C247msLWwyA89hVKwRLkMk=",
        version = "v0.0.0-20170206035902-86b7830d14cc",
    )
    go_repository(
        name = "com_github_shurcool_httpfs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/httpfs",
        sum = "h1:SWV2fHctRpRrp49VXJ6UZja7gU9QLHwRpIPBN89SKEo=",
        version = "v0.0.0-20171119174359-809beceb2371",
    )
    go_repository(
        name = "com_github_shurcool_httpgzip",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/httpgzip",
        sum = "h1:fxoFD0in0/CBzXoyNhMTjvBZYW6ilSnTw7N7y/8vkmM=",
        version = "v0.0.0-20180522190206-b1c53ac65af9",
    )
    go_repository(
        name = "com_github_shurcool_issues",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/issues",
        sum = "h1:T4wuULTrzCKMFlg3HmKHgXAF8oStFb/+lOIupLV2v+o=",
        version = "v0.0.0-20181008053335-6292fdc1e191",
    )
    go_repository(
        name = "com_github_shurcool_issuesapp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/issuesapp",
        sum = "h1:Y+TeIabU8sJD10Qwd/zMty2/LEaT9GNDaA6nyZf+jgo=",
        version = "v0.0.0-20180602232740-048589ce2241",
    )
    go_repository(
        name = "com_github_shurcool_notifications",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/notifications",
        sum = "h1:TQVQrsyNaimGwF7bIhzoVC9QkKm4KsWd8cECGzFx8gI=",
        version = "v0.0.0-20181007000457-627ab5aea122",
    )
    go_repository(
        name = "com_github_shurcool_octicon",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/octicon",
        sum = "h1:bu666BQci+y4S0tVRVjsHUeRon6vUXmsGBwdowgMrg4=",
        version = "v0.0.0-20181028054416-fa4f57f9efb2",
    )
    go_repository(
        name = "com_github_shurcool_reactions",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/reactions",
        sum = "h1:LneqU9PHDsg/AkPDU3AkqMxnMYL+imaqkpflHu73us8=",
        version = "v0.0.0-20181006231557-f2e0b4ca5b82",
    )
    go_repository(
        name = "com_github_shurcool_sanitized_anchor_name",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/sanitized_anchor_name",
        sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_shurcool_users",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/users",
        sum = "h1:YGaxtkYjb8mnTvtufv2LKLwCQu2/C7qFB7UtrOlTWOY=",
        version = "v0.0.0-20180125191416-49c67e49c537",
    )
    go_repository(
        name = "com_github_shurcool_webdavfs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/webdavfs",
        sum = "h1:JtcyT0rk/9PKOdnKQzuDR+FSjh7SGtJwpgVpfZBRKlQ=",
        version = "v0.0.0-20170829043945-18c3829fa133",
    )
    go_repository(
        name = "com_github_siebenmann_go_kstat",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/siebenmann/go-kstat",
        sum = "h1:GfSdC6wKfTGcgCS7BtzF5694Amne1pGCSTY252WhlEY=",
        version = "v0.0.0-20210513183136-173c9b0a9973",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:trlNQbNUG3OdDrDil03MCb1H2o9nJ1x4/5LYw7byDE0=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_skip2_go_qrcode",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/skip2/go-qrcode",
        sum = "h1:MRM5ITcdelLK2j1vwZ3Je0FKVCfqOLp5zO6trqMLYs0=",
        version = "v0.0.0-20200617195104-da1b6568686e",
    )
    go_repository(
        name = "com_github_skratchdot_open_golang",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/skratchdot/open-golang",
        sum = "h1:fyKiXKO1/I/B6Y2U8T7WdQGWzwehOuGIrljPtt7YTTI=",
        version = "v0.0.0-20160302144031-75fb7ed4208c",
    )
    go_repository(
        name = "com_github_slackhq_nebula",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/slackhq/nebula",
        sum = "h1:wuIOHsOnrNw3rQx8yPxXiGu8wAtAxxtUI/K8W7Vj7EI=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_smallstep_assert",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/smallstep/assert",
        sum = "h1:unQFBIznI+VYD1/1fApl1A+9VcBk+9dcqGfnePY87LY=",
        version = "v0.0.0-20200723003110-82e2b9b3b262",
    )
    go_repository(
        name = "com_github_smallstep_certificates",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/smallstep/certificates",
        sum = "h1:wW344Q/QpupjKKFKa4PqzEXfwgeq/54dkU/HNvGnwQQ=",
        version = "v0.19.0",
    )
    go_repository(
        name = "com_github_smallstep_certinfo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/smallstep/certinfo",
        sum = "h1:XO3f9krMkKFDRG4CJFdb3vVoLaXMDAhCn3g0padk9EI=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_smallstep_cli",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/smallstep/cli",
        sum = "h1:BslbUHuMfj/LbVHxuZ4Hv1sL+vAHHidqia4JRoCBwXs=",
        version = "v0.18.0",
    )
    go_repository(
        name = "com_github_smallstep_nosql",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/smallstep/nosql",
        sum = "h1:Go3WYwttUuvwqMtFiiU4g7kBIlY+hR0bIZAqVdakQ3M=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_smallstep_truststore",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/smallstep/truststore",
        sum = "h1:JUTkQ4oHr40jHTS/A2t0usEhteMWG+45CDD2iJA/dIk=",
        version = "v0.11.0",
    )
    go_repository(
        name = "com_github_smallstep_zcrypto",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/smallstep/zcrypto",
        sum = "h1:q0IDrQpquiWcU1nJmjwEremPwG+pT2AAGsDatPgg3Kw=",
        version = "v0.0.0-20210924233136-66c2600f6e71",
    )
    go_repository(
        name = "com_github_smallstep_zlint",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/smallstep/zlint",
        sum = "h1:mZxeNDk2+xg4eTEuR4y6z2+i6HBvcWjtdscDiUmBRzc=",
        version = "v0.0.0-20180727184541-d84eaafe274f",
    )
    go_repository(
        name = "com_github_smartystreets_assertions",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/smartystreets/assertions",
        sum = "h1:UVQPSSmc3qtTi+zPPkCXvZX9VvW/xT/NsRvKfwY81a8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_smartystreets_go_aws_auth",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/smartystreets/go-aws-auth",
        sum = "h1:hp2CYQUINdZMHdvTdXtPOY2ainKl4IoMcpAXEf2xj3Q=",
        version = "v0.0.0-20180515143844-0c1422d1fdb9",
    )
    go_repository(
        name = "com_github_smartystreets_goconvey",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/smartystreets/goconvey",
        sum = "h1:fv0U8FUIMPNf1L9lnHLvLhgicrIVChEkdzIKYqbNC9s=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_github_smartystreets_gunit",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/smartystreets/gunit",
        sum = "h1:RyPDUFcJbvtXlhJPk7v+wnxZRY2EUokhEYl2EJOPToI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_soheilhy_cmux",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/soheilhy/cmux",
        sum = "h1:jjzc5WVemNEDTLwv9tlmemhC73tI08BNOIGwBOo10Js=",
        version = "v0.1.5",
    )
    go_repository(
        name = "com_github_songgao_water",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/songgao/water",
        sum = "h1:TG/diQgUe0pntT/2D9tmUCz4VNwm9MfrtPr0SU2qSX8=",
        version = "v0.0.0-20200317203138-2b4b6d7c09d8",
    )
    go_repository(
        name = "com_github_sony_gobreaker",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/sony/gobreaker",
        sum = "h1:oMnRNZXX5j85zso6xCPRNPtmAycat+WcoKbklScLDgQ=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_sourcegraph_annotate",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/sourcegraph/annotate",
        sum = "h1:yKm7XZV6j9Ev6lojP2XaIshpT4ymkqhMeSghO5Ps00E=",
        version = "v0.0.0-20160123013949-f4cad6c6324d",
    )
    go_repository(
        name = "com_github_sourcegraph_go_diff",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/sourcegraph/go-diff",
        sum = "h1:lhIKJ2nXLZZ+AfbHpYxTn0pXpNTTui0DX7DO3xeb1Zs=",
        version = "v0.5.3",
    )
    go_repository(
        name = "com_github_sourcegraph_syntaxhighlight",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/sourcegraph/syntaxhighlight",
        sum = "h1:qpG93cPwA5f7s/ZPBJnGOYQNK/vKsaDaseuKT5Asee8=",
        version = "v0.0.0-20170531221838-bd320f5d308e",
    )
    go_repository(
        name = "com_github_spaolacci_murmur3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/spaolacci/murmur3",
        sum = "h1:7c1g84S4BPRrfL5Xrdp6fOJ206sU9y293DDHaoy0bLI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_spf13_afero",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/afero",
        sum = "h1:p5gZEKLYoL7wh8VrJesMaYeNxdEd1v3cb4irOk9zB54=",
        version = "v1.3.3",
    )
    go_repository(
        name = "com_github_spf13_cast",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/cast",
        sum = "h1:s0hze+J0196ZfEMTs80N7UlFt0BDuQ7Q+JDnHiMWKdA=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_spf13_cobra",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/cobra",
        sum = "h1:hyqWnYt1ZQShIddO5kBpj3vu05/++x6tJ6dg8EC572I=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_spf13_jwalterweatherman",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/jwalterweatherman",
        sum = "h1:ue6voC5bR5F8YxI5S67j9i582FU4Qvo2bmqnqMYADFk=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/pflag",
        sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_spf13_viper",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/viper",
        sum = "h1:xVKxvI7ouOI5I+U9s2eeiUfMaWBVoXA3AWskkrqK0VM=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_src_d_gcfg",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/src-d/gcfg",
        sum = "h1:xXbNR5AlLSA315x2UO+fTSSAXCDf+Ar38/6oyGbDKQ4=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_stackexchange_wmi",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/StackExchange/wmi",
        sum = "h1:VIkavFPXSjcnS+O8yTq7NI32k0R5Aj+v39y29VYDOSA=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_status_im_keycard_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/status-im/keycard-go",
        sum = "h1:Gb2Tyox57NRNuZ2d3rmvB3pcmbu7O1RS3m8WRx7ilrg=",
        version = "v0.0.0-20190316090335-8537d3370df4",
    )
    go_repository(
        name = "com_github_stefanberger_go_pkcs11uri",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/stefanberger/go-pkcs11uri",
        sum = "h1:lIOOHPEbXzO3vnmx2gok1Tfs31Q8GQqKLc8vVqyQq/I=",
        version = "v0.0.0-20201008174630-78d3cae3a980",
    )
    go_repository(
        name = "com_github_stoewer_go_strcase",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/stoewer/go-strcase",
        sum = "h1:Z2iHWqGXH00XYgqDmNgQbIBxf3wrNq0F3feEy0ainaU=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_streadway_amqp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/streadway/amqp",
        sum = "h1:WhxRHzgeVGETMlmVfqhRn8RIeeNoPr2Czh33I4Zdccw=",
        version = "v0.0.0-20190827072141-edfb9018d271",
    )
    go_repository(
        name = "com_github_streadway_handy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/streadway/handy",
        sum = "h1:AhmOdSHeswKHBjhsLs/7+1voOxT+LLrSk/Nxvk35fug=",
        version = "v0.0.0-20190108123426-d5acb3125c2a",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/stretchr/objx",
        sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/stretchr/testify",
        sum = "h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=",
        version = "v1.8.4",
    )
    go_repository(
        name = "com_github_subosito_gotenv",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/subosito/gotenv",
        sum = "h1:Slr1R9HxAlEKefgq5jn9U+DnETlIUa6HfgEzj0g5d7s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_supranational_blst",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/supranational/blst",
        sum = "h1:m+8fKfQwCAy1QjzINvKe/pYtLjo2dl59x2w9YSEJxuY=",
        version = "v0.3.8-0.20220526154634-513d2456b344",
    )
    go_repository(
        name = "com_github_syndtr_gocapability",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/syndtr/gocapability",
        sum = "h1:kdXcSzyDtseVEc4yCz2qF8ZrQvIDBJLl4S1c3GCXmoI=",
        version = "v0.0.0-20200815063812-42c35b437635",
    )
    go_repository(
        name = "com_github_syndtr_goleveldb",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/syndtr/goleveldb",
        sum = "h1:1ur3QoCqvE5fl+nylMaIr9PVV1w343YRDtsy+Rwu7XI=",
        version = "v1.0.1-0.20220614013038-64ee5596c38a",
    )
    go_repository(
        name = "com_github_tailscale_tscert",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tailscale/tscert",
        sum = "h1:xwMw7LFhV9dbvot9A7NLClP9udqbjrQlIwWMH8e7uiQ=",
        version = "v0.0.0-20220316030059-54bbcb9f74e2",
    )
    go_repository(
        name = "com_github_tango_contrib_basicauth",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tango-contrib/basicauth",
        sum = "h1:tfB+xuTYq48HTLSXD5V99WAn+W+4nqM4KqZAv2ABmfY=",
        version = "v0.0.0-20170526072948-7fbc19aece86",
    )
    go_repository(
        name = "com_github_tarm_serial",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tarm/serial",
        sum = "h1:UyzmZLoiDWMRywV4DUYb9Fbt8uiOSooupjTq10vpvnU=",
        version = "v0.0.0-20180830185346-98f6abe2eb07",
    )
    go_repository(
        name = "com_github_tchap_go_patricia",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tchap/go-patricia",
        sum = "h1:JvoDL7JSoIP2HDE8AbDH3zC8QBPxmzYe32HHy5yQ+Ck=",
        version = "v2.2.6+incompatible",
    )
    go_repository(
        name = "com_github_tdakkota_asciicheck",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tdakkota/asciicheck",
        sum = "h1:HxLVTlqcHhFAz3nWUcuvpH7WuOMv8LQoCWmruLfFH2U=",
        version = "v0.0.0-20200416200610-e657995f937b",
    )
    go_repository(
        name = "com_github_templexxx_cpu",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/templexxx/cpu",
        sum = "h1:pUEZn8JBy/w5yzdYWgx+0m0xL9uk6j4K91C5kOViAzo=",
        version = "v0.0.7",
    )
    go_repository(
        name = "com_github_templexxx_xorsimd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/templexxx/xorsimd",
        sum = "h1:iUZcywbOYDRAZUasAs2eSCUW8eobuZDy0I9FJiORkVg=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_tetafro_godot",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tetafro/godot",
        sum = "h1:Dib7un+rYJFUi8vN0Bk6EHheKy6fv6ZzFURHw75g6m8=",
        version = "v0.4.2",
    )
    go_repository(
        name = "com_github_tetratelabs_wazero",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tetratelabs/wazero",
        sum = "h1:gewLFbKZlapSc6+a7qnS27wAm6L3gdRvv0P76m+236M=",
        version = "v1.0.0-beta1",
    )
    go_repository(
        name = "com_github_thales_e_security_pool",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/thales-e-security/pool",
        sum = "h1:RAPs4q2EbWsTit6tpzuvTFlgFRJ3S8Evf5gtvVDbmPg=",
        version = "v0.0.2",
    )
    go_repository(
        name = "com_github_thalesignite_crypto11",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ThalesIgnite/crypto11",
        sum = "h1:3MebRK/U0mA2SmSthXAIZAdUA9w8+ZuKem2O6HuR1f8=",
        version = "v1.2.4",
    )
    go_repository(
        name = "com_github_thomasrooney_gexpect",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ThomasRooney/gexpect",
        sum = "h1:CjexZrggt4RldpEUXFZf52vSO3cnmFaqW6B4wADj05Q=",
        version = "v0.0.0-20161231170123-5482f0350944",
    )
    go_repository(
        name = "com_github_tidwall_gjson",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/gjson",
        sum = "h1:6aeJ0bzojgWLa82gDQHcx3S0Lr/O51I9bJ5nv6JFx5w=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_github_tidwall_match",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/match",
        sum = "h1:+Ho715JplO36QYgwN9PGYNhgZvoUSc9X2c80KVTi+GA=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_tidwall_pretty",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/pretty",
        sum = "h1:RWIZEg2iJ8/g6fDDYzMpobmaoGh5OLl4AXtGUGPcqCs=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_timakin_bodyclose",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/timakin/bodyclose",
        sum = "h1:ig99OeTyDwQWhPe2iw9lwfQVF1KB3Q4fpP3X7/2VBG8=",
        version = "v0.0.0-20200424151742-cb6215831a94",
    )
    go_repository(
        name = "com_github_timewasted_linode",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/timewasted/linode",
        sum = "h1:CpHxIaZzVy26GqJn8ptRyto8fuoYOd1v0fXm9bG3wQ8=",
        version = "v0.0.0-20160829202747-37e84520dcf7",
    )
    go_repository(
        name = "com_github_tinylib_msgp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tinylib/msgp",
        sum = "h1:gWmO7n0Ys2RBEb7GPYB9Ujq8Mk5p2U08lRnmMcGy6BQ=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_tj_assert",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tj/assert",
        sum = "h1:Rw8kxzWo1mr6FSaYXjQELRe88y2KdfynXdnK72rdjtA=",
        version = "v0.0.0-20171129193455-018094318fb0",
    )
    go_repository(
        name = "com_github_tj_go_elastic",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tj/go-elastic",
        sum = "h1:eGaGNxrtoZf/mBURsnNQKDR7u50Klgcf2eFDQEnc8Bc=",
        version = "v0.0.0-20171221160941-36157cbbebc2",
    )
    go_repository(
        name = "com_github_tj_go_kinesis",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tj/go-kinesis",
        sum = "h1:m74UWYy+HBs+jMFR9mdZU6shPewugMyH5+GV6LNgW8w=",
        version = "v0.0.0-20171128231115-08b17f58cb1b",
    )
    go_repository(
        name = "com_github_tj_go_spin",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tj/go-spin",
        sum = "h1:lhdWZsvImxvZ3q1C5OIB7d72DuOwP4O2NdBg9PyzNds=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_tjfoc_gmsm",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tjfoc/gmsm",
        sum = "h1:aMe1GlZb+0bLjn+cKTPEvvn9oUEBlJitaZiiBwsbgho=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_tklauser_go_sysconf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tklauser/go-sysconf",
        sum = "h1:uu3Xl4nkLzQfXNsWn15rPc/HQCJKObbt1dKJeWp3vU4=",
        version = "v0.3.5",
    )
    go_repository(
        name = "com_github_tklauser_numcpus",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tklauser/numcpus",
        sum = "h1:oyhllyrScuYI6g+h/zUvNXNp1wy7x8qQy3t/piefldA=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_tmc_grpc_websocket_proxy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tmc/grpc-websocket-proxy",
        sum = "h1:6fotK7otjonDflCTK0BCfls4SPy3NcCVb5dqqmbRknE=",
        version = "v0.0.0-20220101234140-673ab2c3ae75",
    )
    go_repository(
        name = "com_github_tomasen_realip",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tomasen/realip",
        sum = "h1:fb190+cK2Xz/dvi9Hv8eCYJYvIGUTN2/KLq1pT6CjEc=",
        version = "v0.0.0-20180522021738-f0c99a92ddce",
    )
    go_repository(
        name = "com_github_tommy_muehle_go_mnd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tommy-muehle/go-mnd",
        sum = "h1:RC4maTWLKKwb7p1cnoygsbKIgNlJqSYBeAFON3Ar8As=",
        version = "v1.3.1-0.20200224220436-e6f9a994e8fa",
    )
    go_repository(
        name = "com_github_traefik_yaegi",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/traefik/yaegi",
        sum = "h1:Dg6hYcDtxWY3L9jP2OjoK+LSX59wF0sckJXT/cYNqA4=",
        version = "v0.9.14",
    )
    go_repository(
        name = "com_github_transip_gotransip",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/transip/gotransip",
        sum = "h1:clyOmELPZd2LuFEyuo1mP6RXpbAW75PwD+RfDj4kBm0=",
        version = "v0.0.0-20190812104329-6d8d9179b66f",
    )
    go_repository(
        name = "com_github_transip_gotransip_v6",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/transip/gotransip/v6",
        sum = "h1:rOCMY607PYF+YvMHHtJt7eZRd0mx/uhyz6dsXWPmn+4=",
        version = "v6.0.2",
    )
    go_repository(
        name = "com_github_tv42_httpunix",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tv42/httpunix",
        sum = "h1:u6SKchux2yDvFQnDHS3lPnIRmfVJ5Sxy3ao2SIdysLQ=",
        version = "v0.0.0-20191220191345-2ba4b9c3382c",
    )
    go_repository(
        name = "com_github_txthinking_runnergroup",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/txthinking/runnergroup",
        sum = "h1:vlDgnShahmE2XLslpr0hnzxfAmSj3JLX2CYi8Xct7G4=",
        version = "v0.0.0-20200327135940-540a793bb997",
    )
    go_repository(
        name = "com_github_txthinking_socks5",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/txthinking/socks5",
        sum = "h1:yu9Bs+KssCJNxu/9fzRag6QgzOnxoH1Q6TvIiD4L6rQ=",
        version = "v0.0.0-20200531111549-252709fcb919",
    )
    go_repository(
        name = "com_github_txthinking_x",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/txthinking/x",
        sum = "h1:ngJOce33YJJT1PFTfC9ao7S27AfrUh11Dr3Bc+ooBdM=",
        version = "v0.0.0-20200330144832-5ad2416896a9",
    )
    go_repository(
        name = "com_github_tyler_smith_go_bip39",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/tyler-smith/go-bip39",
        sum = "h1:wHSqTBrZW24CsNJDfeh9Ex6Pm0Rcpc7qrgKBiL44vF4=",
        version = "v1.0.1-0.20181017060643-dbb3b84ba2ef",
    )
    go_repository(
        name = "com_github_u_root_uio",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/u-root/uio",
        sum = "h1:BFvcl34IGnw8yvJi8hlqLFo9EshRInwWBs2M5fGWzQA=",
        version = "v0.0.0-20210528114334-82958018845c",
    )
    go_repository(
        name = "com_github_uber_go_atomic",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/uber-go/atomic",
        sum = "h1:Azu9lPBWRNKzYXSIwRfgRuDuS0YKsK4NFhiQv98gkxo=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_ugorji_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ugorji/go",
        sum = "h1:j4s+tAvLfL3bZyefP2SEWmhBzmuIlH/eqNuPdFPgngw=",
        version = "v1.1.4",
    )
    go_repository(
        name = "com_github_ugorji_go_codec",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ugorji/go/codec",
        sum = "h1:3SVOIvH7Ae1KRYyQWRjXWJEA9sS/c/pjvH++55Gr648=",
        version = "v0.0.0-20181204163529-d75b2dcb6bc8",
    )
    go_repository(
        name = "com_github_ulikunitz_xz",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ulikunitz/xz",
        sum = "h1:kpFauv27b6ynzBNT/Xy+1k+fK4WswhN/6PN5WhFAGw8=",
        version = "v0.5.11",
    )
    go_repository(
        name = "com_github_ultraware_funlen",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ultraware/funlen",
        sum = "h1:Av96YVBwwNSe4MLR7iI/BIa3VyI7/djnto/pK3Uxbdo=",
        version = "v0.0.2",
    )
    go_repository(
        name = "com_github_ultraware_whitespace",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/ultraware/whitespace",
        sum = "h1:If7Va4cM03mpgrNH9k49/VOicWpGoG70XPBFFODYDsg=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_urfave_cli",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/urfave/cli",
        sum = "h1:lNq9sAHXK2qfdI8W+GRItjCEkI+2oR4d+MEHy1CKXoU=",
        version = "v1.22.5",
    )
    go_repository(
        name = "com_github_urfave_cli_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/urfave/cli/v2",
        sum = "h1:x3p8awjp/2arX+Nl/G2040AZpOCHS/eMJJ1/a+mye4Y=",
        version = "v2.10.2",
    )
    go_repository(
        name = "com_github_uudashr_gocognit",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/uudashr/gocognit",
        sum = "h1:MoG2fZ0b/Eo7NXoIwCVFLG5JED3qgQz5/NEE+rOsjPs=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_v2fly_v2ray_core_v4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/v2fly/v2ray-core/v4",
        sum = "h1:JrA3N5BdIc7+rtWgmP9zqw341BmimYmEtuU4lVBFgHY=",
        version = "v4.36.2",
    )
    go_repository(
        name = "com_github_v2fly_vsign",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/v2fly/VSign",
        sum = "h1:p1UzXK6VAutXFFQMnre66h7g1BjRKUnLv0HfmmRoz7w=",
        version = "v0.0.0-20201108000810-e2adc24bf848",
    )
    go_repository(
        name = "com_github_valyala_bytebufferpool",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/valyala/bytebufferpool",
        sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_valyala_fasthttp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/valyala/fasthttp",
        sum = "h1:Z7kVhKP9NZz+tCSY7AVhCMPPAk7b+e5fq0l/BfdTlFc=",
        version = "v1.13.1",
    )
    go_repository(
        name = "com_github_valyala_fasttemplate",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/valyala/fasttemplate",
        sum = "h1:tY9CJiPnMXf1ERmG2EyK7gNUd+c6RKGD0IfU8WdUSz8=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_valyala_quicktemplate",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/valyala/quicktemplate",
        sum = "h1:BaO1nHTkspYzmAjPXj0QiDJxai96tlcZyKcI9dyEGvM=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_valyala_tcplisten",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/valyala/tcplisten",
        sum = "h1:0R4NLDRDZX6JcmhJgXi5E4b8Wg84ihbmUKp/GvSPEzc=",
        version = "v0.0.0-20161114210144-ceec8f93295a",
    )
    go_repository(
        name = "com_github_vdemeester_k8s_pkg_credentialprovider",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/vdemeester/k8s-pkg-credentialprovider",
        sum = "h1:czKEIG2Q3YRTgs6x/8xhjVMJD5byPo6cZuostkbTM74=",
        version = "v1.17.4",
    )
    go_repository(
        name = "com_github_viant_assertly",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/viant/assertly",
        sum = "h1:5x1GzBaRteIwTr5RAGFVG14uNeRFxVNbXPWrK2qAgpc=",
        version = "v0.4.8",
    )
    go_repository(
        name = "com_github_viant_toolbox",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/viant/toolbox",
        sum = "h1:6TteTDQ68CjgcCe8wH3D3ZhUQQOJXMTbj/D9rkk2a1k=",
        version = "v0.24.0",
    )
    go_repository(
        name = "com_github_victoriametrics_fastcache",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/VictoriaMetrics/fastcache",
        sum = "h1:C/3Oi3EiBCqufydp1neRZkqcwmEiuRT9c3fqvvgKm5o=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_vishvananda_netlink",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/vishvananda/netlink",
        sum = "h1:+UB2BJA852UkGH42H+Oee69djmxS3ANzl2b/JtT1YiA=",
        version = "v1.1.1-0.20210330154013-f5de75959ad5",
    )
    go_repository(
        name = "com_github_vishvananda_netns",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/vishvananda/netns",
        sum = "h1:gga7acRE695APm9hlsSMoOoE65U4/TcqNj90mc69Rlg=",
        version = "v0.0.0-20211101163701-50045581ed74",
    )
    go_repository(
        name = "com_github_vito_go_interact",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/vito/go-interact",
        sum = "h1:Klu98tQ9Z1t23gvC7p7sCmvxkZxLhBHLNyrUPsWsYFg=",
        version = "v0.0.0-20171111012221-fa338ed9e9ec",
    )
    go_repository(
        name = "com_github_vividcortex_gohistogram",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/VividCortex/gohistogram",
        sum = "h1:6+hBz+qvs0JOrrNhhmR7lFxo5sINxBCGXrdtl/UvroE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_vmihailenco_msgpack",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/vmihailenco/msgpack",
        sum = "h1:dSLoQfGFAo3F6OoNhwUmLwVgaUXK79GlxNBwueZn0xI=",
        version = "v4.0.4+incompatible",
    )
    go_repository(
        name = "com_github_vmware_govmomi",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/vmware/govmomi",
        sum = "h1:gpw/0Ku+6RgF3jsi7fnCLmlcikBHfKBCUcu1qgc16OU=",
        version = "v0.20.3",
    )
    go_repository(
        name = "com_github_vojtechvitek_yaml_cli",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/VojtechVitek/yaml-cli",
        sum = "h1:jsovPZX2Yau3+g0/xpwFXkCIfGRehtyecM0vB+Y6pFo=",
        version = "v0.0.5",
    )
    go_repository(
        name = "com_github_vultr_govultr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/vultr/govultr",
        sum = "h1:UnNMixYFVO0p80itc8PcweoVENyo1PasfvwKhoasR9U=",
        version = "v0.1.4",
    )
    go_repository(
        name = "com_github_vultr_govultr_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/vultr/govultr/v2",
        sum = "h1:cvxH1mC/VTs2oRx3njhIhpypg2EBE70rlxAKKtRU/Zg=",
        version = "v2.11.0",
    )
    go_repository(
        name = "com_github_webteleport_auth",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/webteleport/auth",
        sum = "h1:toIPVsYsj4Zi00JAJ7Rv4YfcTg9zjBQFeyk8pcrO0/g=",
        version = "v0.0.6",
    )
    go_repository(
        name = "com_github_webteleport_utils",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/webteleport/utils",
        sum = "h1:Q/+PwgR8z2y6Ha3n7JUAmj+9y56iu2DCQFf+En4Y7yc=",
        version = "v0.2.5",
    )
    go_repository(
        name = "com_github_webteleport_webteleport",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/webteleport/webteleport",
        sum = "h1:mi9qjOwJ3HA1qGKTP4l/JhKP4RwWxG81JwBAGx/KfI0=",
        version = "v0.3.8",
    )
    go_repository(
        name = "com_github_webteleport_wtf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/webteleport/wtf",
        sum = "h1:I7cDZkw8OjzKDJxiV52WQ6c3Zg/hIQg6VcJPn+U9+Qk=",
        version = "v0.0.8",
    )
    go_repository(
        name = "com_github_weppos_publicsuffix_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/weppos/publicsuffix-go",
        sum = "h1:YSnfg3V65LcCFKtIGKGoBhkyKolEd0hlipcXaOjdnQw=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_wi2l_jsondiff",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/wI2L/jsondiff",
        sum = "h1:dE00WemBa1uCjrzQUUTE/17I6m5qAaN0EMFOg2Ynr/k=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_willf_bitset",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/willf/bitset",
        sum = "h1:N7Z7E9UvjW+sGsEl7k/SJrvY2reP1A07MrGuCjIOjRE=",
        version = "v1.1.11",
    )
    go_repository(
        name = "com_github_x448_float16",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/x448/float16",
        sum = "h1:qLwI1I70+NjRFUR3zs1JPUCgaCXSh3SW62uAKT1mSBM=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_github_xanzy_go_gitlab",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xanzy/go-gitlab",
        sum = "h1:4p0IirHqEp5f0baK/aQqr4TR57IsD+8e4fuyAA1yi88=",
        version = "v0.95.2",
    )
    go_repository(
        name = "com_github_xanzy_ssh_agent",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xanzy/ssh-agent",
        sum = "h1:wUMzuKtKilRgBAD1sUb8gOwwRr2FGoBVumcjoOACClI=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_xdg_scram",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xdg/scram",
        sum = "h1:u40Z8hqBAAQyv+vATcGgV0YCnDjqSL7/q/JyPhhJSPk=",
        version = "v0.0.0-20180814205039-7eeb5667e42c",
    )
    go_repository(
        name = "com_github_xdg_stringprep",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xdg/stringprep",
        sum = "h1:d9X0esnoa3dFsV0FG35rAT0RIhYFlPq7MiP+DW89La0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonpointer",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xeipuuv/gojsonpointer",
        sum = "h1:J9EGpcZtP0E/raorCMxlFGSTBrsSlaDGf3jU/qvAE2c=",
        version = "v0.0.0-20180127040702-4e3ac2762d5f",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonreference",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xeipuuv/gojsonreference",
        sum = "h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=",
        version = "v0.0.0-20180127040603-bd5ef7bd5415",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonschema",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xeipuuv/gojsonschema",
        sum = "h1:LhYJRs+L4fBtjZUfuSZIKGeVu0QRy8e5Xi7D17UxZ74=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_xhit_go_str2duration_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xhit/go-str2duration/v2",
        sum = "h1:lxklc02Drh6ynqX+DdPyp5pCKLUQpRT8bp8Ydu2Bstc=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_xi2_xz",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xi2/xz",
        sum = "h1:nIPpBwaJSVYIxUFsDv3M8ofmx9yWTog9BfvIu0q41lo=",
        version = "v0.0.0-20171230120015-48954b6210f8",
    )
    go_repository(
        name = "com_github_xiang90_probing",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xiang90/probing",
        sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
        version = "v0.0.0-20190116061207-43a291ad63a2",
    )
    go_repository(
        name = "com_github_xiaq_persistent",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xiaq/persistent",
        sum = "h1:HxX+USVm4XyGwvWS0eJy+GMttkfSRdFcrZ46WtAs5RQ=",
        version = "v0.0.0-20200820214153-3175cfb92e14",
    )
    go_repository(
        name = "com_github_xlab_treeprint",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xlab/treeprint",
        sum = "h1:HzHnuAF1plUN2zGlAFHbSQP2qJ0ZAD3XF5XD7OesXRQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_xordataexchange_crypt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xordataexchange/crypt",
        sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
        version = "v0.0.3-0.20170626215501-b2862e3d0a77",
    )
    go_repository(
        name = "com_github_xrash_smetrics",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xrash/smetrics",
        sum = "h1:bAn7/zixMGCfxrRTfdpNzjtPYqr8smhKouy9mxVdGPU=",
        version = "v0.0.0-20201216005158-039620a65673",
    )
    go_repository(
        name = "com_github_xtaci_kcp_go_v5",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xtaci/kcp-go/v5",
        sum = "h1:Pwn0aoeNSPF9dTS7IgiPXn0HEtaIlVb6y5UKWPsx8bI=",
        version = "v5.6.1",
    )
    go_repository(
        name = "com_github_xtaci_lossyconn",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xtaci/lossyconn",
        sum = "h1:EWU6Pktpas0n8lLQwDsRyZfmkPeRbdgPtW609es+/9E=",
        version = "v0.0.0-20200209145036-adba10fffc37",
    )
    go_repository(
        name = "com_github_xtaci_smux",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xtaci/smux",
        sum = "h1:FBPYOkW8ZTjLKUM4LI4xnnuuDC8CQ/dB04HD519WoEk=",
        version = "v1.5.16",
    )
    go_repository(
        name = "com_github_xtaci_tcpraw",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/xtaci/tcpraw",
        sum = "h1:VDlqo0op17JeXBM6e2G9ocCNLOJcw9mZbobMbJjo0vk=",
        version = "v1.2.25",
    )
    go_repository(
        name = "com_github_yalp_jsonpath",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/yalp/jsonpath",
        sum = "h1:6fRhSjgLCkTD3JnJxvaJ4Sj+TYblw757bqYgZaOq5ZY=",
        version = "v0.0.0-20180802001716-5cc68e5049a0",
    )
    go_repository(
        name = "com_github_yudai_gojsondiff",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/yudai/gojsondiff",
        sum = "h1:27cbfqXLVEJ1o8I6v3y9lg8Ydm53EKqHXAOMxEGlCOA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_yudai_golcs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/yudai/golcs",
        sum = "h1:BHyfKlQyqbsFN5p3IfnEUduWvb9is428/nNb5L3U01M=",
        version = "v0.0.0-20170316035057-ecda9a501e82",
    )
    go_repository(
        name = "com_github_yudai_pp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/yudai/pp",
        sum = "h1:Q4//iY4pNF6yPLZIigmvcl7k/bPgrcTPIFIcmawg5bI=",
        version = "v2.0.1+incompatible",
    )
    go_repository(
        name = "com_github_yuin_goldmark",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
        version = "v1.4.13",
    )
    go_repository(
        name = "com_github_yuin_goldmark_highlighting",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/yuin/goldmark-highlighting",
        sum = "h1:yHfZyN55+5dp1wG7wDKv8HQ044moxkyGq12KFFMFDxg=",
        version = "v0.0.0-20220208100518-594be1970594",
    )
    go_repository(
        name = "com_github_yvasiyarov_go_metrics",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/yvasiyarov/go-metrics",
        sum = "h1:+lm10QQTNSBd8DVTNGHx7o/IKu9HYDvLMffDhbyLccI=",
        version = "v0.0.0-20140926110328-57bccd1ccd43",
    )
    go_repository(
        name = "com_github_yvasiyarov_gorelic",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/yvasiyarov/gorelic",
        sum = "h1:hlE8//ciYMztlGpl/VA+Zm1AcTPHYkHJPbHqE6WJUXE=",
        version = "v0.0.0-20141212073537-a9bba5b9ab50",
    )
    go_repository(
        name = "com_github_yvasiyarov_newrelic_platform_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/yvasiyarov/newrelic_platform_go",
        sum = "h1:ERexzlUfuTvpE74urLSbIQW0Z/6hF9t8U4NsJLaioAY=",
        version = "v0.0.0-20140908184405-b21fdbd4370f",
    )
    go_repository(
        name = "com_github_zenazn_goji",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/zenazn/goji",
        sum = "h1:RSQQAbXGArQ0dIDEq+PI6WqN6if+5KHu6x2Cx/GXLTQ=",
        version = "v0.9.0",
    )
    go_repository(
        name = "com_github_zmap_rc2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/zmap/rc2",
        sum = "h1:Nzukz5fNOBIHOsnP+6I79kPx3QhLv8nBy2mfFhBRq30=",
        version = "v0.0.0-20190804163417-abaa70531248",
    )
    go_repository(
        name = "com_github_zmap_zcertificate",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/zmap/zcertificate",
        sum = "h1:hxHelFG6LEcCsUyu6oKo4P7ZkmzLLeQhOZlVtaUymBk=",
        version = "v0.0.0-20190521191901-30e388164f71",
    )
    go_repository(
        name = "com_github_zmap_zcrypto",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/zmap/zcrypto",
        sum = "h1:EoQIHS1co8tkbljRLMADWiRAWLcKI02M/ZtPrAUxjHc=",
        version = "v0.0.0-20190329181646-dff83107394d",
    )
    go_repository(
        name = "com_github_zmap_zlint",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "github.com/zmap/zlint",
        sum = "h1:174pnZ4WOF6mGuOJy7Qm6V3cmWn61CfhAWMxvPhqwmc=",
        version = "v0.0.0-20190516161541-9047d02cf65a",
    )
    go_repository(
        name = "com_gitlab_mjwhitta_cli",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitlab.com/mjwhitta/cli",
        sum = "h1:rTSnmWaOItnQZ3f2NcPVjS1JSZcJFvwt9ctfWTBfut4=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_gitlab_mjwhitta_errors",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitlab.com/mjwhitta/errors",
        sum = "h1:rivvgv4HCdIJ6Y6kFfPcirzKnBRewkjgsOktR0Uaj8g=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_gitlab_mjwhitta_hilighter",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitlab.com/mjwhitta/hilighter",
        sum = "h1:meBsY/zGSEgxAFY092AZc4e16Zi7FlEzLFv+ragPLsc=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_gitlab_mjwhitta_jq",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitlab.com/mjwhitta/jq",
        sum = "h1:XEjsAfYAP+4Fc3FWgd5O58Gsak3WuOsVE5lH6EsxpII=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_gitlab_mjwhitta_jsoncfg",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitlab.com/mjwhitta/jsoncfg",
        sum = "h1:GxrNiulOGmWkG8Of09jjpkUWo7odZwhWgEorKnWGvpU=",
        version = "v1.6.7",
    )
    go_repository(
        name = "com_gitlab_mjwhitta_log",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitlab.com/mjwhitta/log",
        sum = "h1:f/LWphvfLt93tjABlnTVLn+2bf7DjYG8AEE/pP1q3zQ=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_gitlab_mjwhitta_pathname",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitlab.com/mjwhitta/pathname",
        sum = "h1:s8KM/t85B9jTLlnvssOaa1et2dWoMe+Jf3h/9RSCxRk=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_gitlab_mjwhitta_safety",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitlab.com/mjwhitta/safety",
        sum = "h1:P1rukXKs6A5HPCxVsQJxQiL+LDk/rQoG/Whrv0cbY9s=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_gitlab_mjwhitta_sysinfo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitlab.com/mjwhitta/sysinfo",
        sum = "h1:aq5wCbBDR/ItorhpEYDZYXSJ3R/T+kwRWNd6Q6jPefw=",
        version = "v1.4.8",
    )
    go_repository(
        name = "com_gitlab_mjwhitta_where",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitlab.com/mjwhitta/where",
        sum = "h1:Vnrk4rQE2vk1kgINJhE8jPdOeGdi7w+OzqnkVIQwK6I=",
        version = "v1.2.5",
    )
    go_repository(
        name = "com_gitlab_yawning_edwards25519_extra_git",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitlab.com/yawning/edwards25519-extra.git",
        sum = "h1:qRSZHsODmAP5qDvb3YsO7Qnf3TRiVbGxNG/WYnlM4/o=",
        version = "v0.0.0-20211229043746-2f91fcc9fbdb",
    )
    go_repository(
        name = "com_gitlab_yawning_obfs4_git",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gitlab.com/yawning/obfs4.git",
        sum = "h1:tJ8F7ABaQ3p3wjxwXiWSktVDgjZEXkvaRawd2rIq5ws=",
        version = "v0.0.0-20220204003609-77af0cba934d",
    )
    go_repository(
        name = "com_google_cloud_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go",
        sum = "h1:8uYAkj3YHTP/1iwReuHPxLSbdcyc+dSBbzFMrVwDR6Q=",
        version = "v0.110.6",
    )
    go_repository(
        name = "com_google_cloud_go_accessapproval",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/accessapproval",
        sum = "h1:/5YjNhR6lzCvmJZAnByYkfEgWjfAKwYP6nkuTk6nKFE=",
        version = "v1.7.1",
    )
    go_repository(
        name = "com_google_cloud_go_accesscontextmanager",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/accesscontextmanager",
        sum = "h1:WIAt9lW9AXtqw/bnvrEUaE8VG/7bAAeMzRCBGMkc4+w=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_aiplatform",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/aiplatform",
        sum = "h1:M5davZWCTzE043rJCn+ZLW6hSxfG1KAx4vJTtas2/ec=",
        version = "v1.48.0",
    )
    go_repository(
        name = "com_google_cloud_go_analytics",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/analytics",
        sum = "h1:TFBC1ZAqX9/jL56GEXdLrVe5vT3I22bDVWyDwZX4IEg=",
        version = "v0.21.3",
    )
    go_repository(
        name = "com_google_cloud_go_apigateway",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/apigateway",
        sum = "h1:aBSwCQPcp9rZ0zVEUeJbR623palnqtvxJlUyvzsKGQc=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeconnect",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/apigeeconnect",
        sum = "h1:6u/jj0P2c3Mcm+H9qLsXI7gYcTiG9ueyQL3n6vCmFJM=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeregistry",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/apigeeregistry",
        sum = "h1:hgq0ANLDx7t2FDZDJQrCMtCtddR/pjCqVuvQWGrQbXw=",
        version = "v0.7.1",
    )
    go_repository(
        name = "com_google_cloud_go_appengine",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/appengine",
        sum = "h1:J+aaUZ6IbTpBegXbmEsh8qZZy864ZVnOoWyfa1XSNbI=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_area120",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/area120",
        sum = "h1:wiOq3KDpdqXmaHzvZwKdpoM+3lDcqsI2Lwhyac7stss=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_artifactregistry",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/artifactregistry",
        sum = "h1:k6hNqab2CubhWlGcSzunJ7kfxC7UzpAfQ1UPb9PDCKI=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_google_cloud_go_asset",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/asset",
        sum = "h1:vlHdznX70eYW4V1y1PxocvF6tEwxJTTarwIGwOhFF3U=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_google_cloud_go_assuredworkloads",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/assuredworkloads",
        sum = "h1:yaO0kwS+SnhVSTF7BqTyVGt3DTocI6Jqo+S3hHmCwNk=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_google_cloud_go_automl",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/automl",
        sum = "h1:iP9iQurb0qbz+YOOMfKSEjhONA/WcoOIjt6/m+6pIgo=",
        version = "v1.13.1",
    )
    go_repository(
        name = "com_google_cloud_go_baremetalsolution",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/baremetalsolution",
        sum = "h1:0Ge9PQAy6cZ1tRrkc44UVgYV15nw2TVnzJzYsMHXF+E=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_google_cloud_go_batch",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/batch",
        sum = "h1:uE0Q//W7FOGPjf7nuPiP0zoE8wOT3ngoIO2HIet0ilY=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_google_cloud_go_beyondcorp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/beyondcorp",
        sum = "h1:VPg+fZXULQjs8LiMeWdLaB5oe8G9sEoZ0I0j6IMiG1Q=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_google_cloud_go_bigquery",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/bigquery",
        sum = "h1:K3wLbjbnSlxhuG5q4pntHv5AEbQM1QqHKGYgwFIqOTg=",
        version = "v1.53.0",
    )
    go_repository(
        name = "com_google_cloud_go_billing",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/billing",
        sum = "h1:1iktEAIZ2uA6KpebC235zi/rCXDdDYQ0bTXTNetSL80=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_binaryauthorization",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/binaryauthorization",
        sum = "h1:cAkOhf1ic92zEN4U1zRoSupTmwmxHfklcp1X7CCBKvE=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_google_cloud_go_certificatemanager",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/certificatemanager",
        sum = "h1:uKsohpE0hiobx1Eak9jNcPCznwfB6gvyQCcS28Ah9E8=",
        version = "v1.7.1",
    )
    go_repository(
        name = "com_google_cloud_go_channel",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/channel",
        sum = "h1:dqRkK2k7Ll/HHeYGxv18RrfhozNxuTJRkspW0iaFZoY=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_cloudbuild",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/cloudbuild",
        sum = "h1:YBbAWcvE4x6xPWTyS+OU4eiUpz5rCS3VCM/aqmfddPA=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_clouddms",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/clouddms",
        sum = "h1:rjR1nV6oVf2aNNB7B5uz1PDIlBjlOiBgR+q5n7bbB7M=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_google_cloud_go_cloudtasks",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/cloudtasks",
        sum = "h1:cMh9Q6dkvh+Ry5LAPbD/U2aw6KAqdiU6FttwhbTo69w=",
        version = "v1.12.1",
    )
    go_repository(
        name = "com_google_cloud_go_compute",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/compute",
        sum = "h1:tP41Zoavr8ptEqaW6j+LQOnyBBhO7OkOMAGrgLopTwY=",
        version = "v1.23.0",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:mg4jlk7mCAj6xXp9UJ4fjI9VUI5rubuGBW5aJ7UnBMY=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_google_cloud_go_contactcenterinsights",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/contactcenterinsights",
        sum = "h1:YR2aPedGVQPpFBZXJnPkqRj8M//8veIZZH5ZvICoXnI=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_container",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/container",
        sum = "h1:N51t/cgQJFqDD/W7Mb+IvmAPHrf8AbPx7Bb7aF4lROE=",
        version = "v1.24.0",
    )
    go_repository(
        name = "com_google_cloud_go_containeranalysis",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/containeranalysis",
        sum = "h1:SM/ibWHWp4TYyJMwrILtcBtYKObyupwOVeceI9pNblw=",
        version = "v0.10.1",
    )
    go_repository(
        name = "com_google_cloud_go_datacatalog",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datacatalog",
        sum = "h1:qVeQcw1Cz93/cGu2E7TYUPh8Lz5dn5Ws2siIuQ17Vng=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataflow",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataflow",
        sum = "h1:VzG2tqsk/HbmOtq/XSfdF4cBvUWRK+S+oL9k4eWkENQ=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_google_cloud_go_dataform",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataform",
        sum = "h1:xcWso0hKOoxeW72AjBSIp/UfkvpqHNzzS0/oygHlcqY=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_datafusion",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datafusion",
        sum = "h1:eX9CZoyhKQW6g1Xj7+RONeDj1mV8KQDKEB9KLELX9/8=",
        version = "v1.7.1",
    )
    go_repository(
        name = "com_google_cloud_go_datalabeling",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datalabeling",
        sum = "h1:zxsCD/BLKXhNuRssen8lVXChUj8VxF3ofN06JfdWOXw=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_dataplex",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataplex",
        sum = "h1:yoBWuuUZklYp7nx26evIhzq8+i/nvKYuZr1jka9EqLs=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataproc_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataproc/v2",
        sum = "h1:4OpSiPMMGV3XmtPqskBU/RwYpj3yMFjtMLj/exi425Q=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_google_cloud_go_dataqna",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataqna",
        sum = "h1:ITpUJep04hC9V7C+gcK390HO++xesQFSUJ7S4nSnF3U=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_datastore",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datastore",
        sum = "h1:ktbC66bOQB3HJPQe8qNI1/aiQ77PMu7hD4mzE6uxe3w=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_datastream",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datastream",
        sum = "h1:ra/+jMv36zTAGPfi8TRne1hXme+UsKtdcK4j6bnqQiw=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_deploy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/deploy",
        sum = "h1:A+w/xpWgz99EYzB6e31gMGAI/P5jTZ2UO7veQK5jQ8o=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_dialogflow",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dialogflow",
        sum = "h1:sCJbaXt6ogSbxWQnERKAzos57f02PP6WkGbOZvXUdwc=",
        version = "v1.40.0",
    )
    go_repository(
        name = "com_google_cloud_go_dlp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dlp",
        sum = "h1:tF3wsJ2QulRhRLWPzWVkeDz3FkOGVoMl6cmDUHtfYxw=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_google_cloud_go_documentai",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/documentai",
        sum = "h1:dW8ex9yb3oT9s1yD2+yLcU8Zq15AquRZ+wd0U+TkxFw=",
        version = "v1.22.0",
    )
    go_repository(
        name = "com_google_cloud_go_domains",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/domains",
        sum = "h1:rqz6KY7mEg7Zs/69U6m6LMbB7PxFDWmT3QWNXIqhHm0=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_google_cloud_go_edgecontainer",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/edgecontainer",
        sum = "h1:zhHWnLzg6AqzE+I3gzJqiIwHfjEBhWctNQEzqb+FaRo=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_google_cloud_go_errorreporting",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/errorreporting",
        sum = "h1:kj1XEWMu8P0qlLhm3FwcaFsUvXChV/OraZwA70trRR0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_essentialcontacts",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/essentialcontacts",
        sum = "h1:OEJ0MLXXCW/tX1fkxzEZOsv/wRfyFsvDVNaHWBAvoV0=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_google_cloud_go_eventarc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/eventarc",
        sum = "h1:xIP3XZi0Xawx8DEfh++mE2lrIi5kQmCr/KcWhJ1q0J4=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_filestore",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/filestore",
        sum = "h1:Eiz8xZzMJc5ppBWkuaod/PUdUZGCFR8ku0uS+Ah2fRw=",
        version = "v1.7.1",
    )
    go_repository(
        name = "com_google_cloud_go_firestore",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/firestore",
        sum = "h1:PPgtwcYUOXV2jFe1bV3nda3RCrOa8cvBjTOn2MQVfW8=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_google_cloud_go_functions",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/functions",
        sum = "h1:LtAyqvO1TFmNLcROzHZhV0agEJfBi+zfMZsF4RT/a7U=",
        version = "v1.15.1",
    )
    go_repository(
        name = "com_google_cloud_go_gkebackup",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gkebackup",
        sum = "h1:lgyrpdhtJKV7l1GM15YFt+OCyHMxsQZuSydyNmS0Pxo=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_gkeconnect",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gkeconnect",
        sum = "h1:a1ckRvVznnuvDWESM2zZDzSVFvggeBaVY5+BVB8tbT0=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_gkehub",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gkehub",
        sum = "h1:2BLSb8i+Co1P05IYCKATXy5yaaIw/ZqGvVSBTLdzCQo=",
        version = "v0.14.1",
    )
    go_repository(
        name = "com_google_cloud_go_gkemulticloud",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gkemulticloud",
        sum = "h1:MluqhtPVZReoriP5+adGIw+ij/RIeRik8KApCW2WMTw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_google_cloud_go_gsuiteaddons",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gsuiteaddons",
        sum = "h1:mi9jxZpzVjLQibTS/XfPZvl+Jr6D5Bs8pGqUjllRb00=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_google_cloud_go_iam",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/iam",
        sum = "h1:lW7fzj15aVIXYHREOqjRBV9PsH0Z6u8Y46a1YGvQP4Y=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_google_cloud_go_iap",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/iap",
        sum = "h1:X1tcp+EoJ/LGX6cUPt3W2D4H2Kbqq0pLAsldnsCjLlE=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_ids",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/ids",
        sum = "h1:khXYmSoDDhWGEVxHl4c4IgbwSRR+qE/L4hzP3vaU9Hc=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_google_cloud_go_iot",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/iot",
        sum = "h1:yrH0OSmicD5bqGBoMlWG8UltzdLkYzNUwNVUVz7OT54=",
        version = "v1.7.1",
    )
    go_repository(
        name = "com_google_cloud_go_kms",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/kms",
        sum = "h1:xYl5WEaSekKYN5gGRyhjvZKM22GVBBCzegGNVPy+aIs=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_language",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/language",
        sum = "h1:3MXeGEv8AlX+O2LyV4pO4NGpodanc26AmXwOuipEym0=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_google_cloud_go_lifesciences",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/lifesciences",
        sum = "h1:axkANGx1wiBXHiPcJZAE+TDjjYoJRIDzbHC/WYllCBU=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_google_cloud_go_logging",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/logging",
        sum = "h1:CJYxlNNNNAMkHp9em/YEXcfJg+rPDg7YfwoRpMU+t5I=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_longrunning",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/longrunning",
        sum = "h1:Fr7TXftcqTudoyRJa113hyaqlGdiBQkp0Gq7tErFDWI=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_google_cloud_go_managedidentities",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/managedidentities",
        sum = "h1:2/qZuOeLgUHorSdxSQGtnOu9xQkBn37+j+oZQv/KHJY=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_google_cloud_go_maps",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/maps",
        sum = "h1:PdfgpBLhAoSzZrQXP+/zBc78fIPLZSJp5y8+qSMn2UU=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_mediatranslation",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/mediatranslation",
        sum = "h1:50cF7c1l3BanfKrpnTCaTvhf+Fo6kdF21DG0byG7gYU=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_memcache",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/memcache",
        sum = "h1:7lkLsF0QF+Mre0O/NvkD9Q5utUNwtzvIYjrOLOs0HO0=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_google_cloud_go_metastore",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/metastore",
        sum = "h1:+9DsxUOHvsqvC0ylrRc/JwzbXJaaBpfIK3tX0Lx8Tcc=",
        version = "v1.12.0",
    )
    go_repository(
        name = "com_google_cloud_go_monitoring",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/monitoring",
        sum = "h1:65JhLMd+JiYnXr6j5Z63dUYCuOg770p8a/VC+gil/58=",
        version = "v1.15.1",
    )
    go_repository(
        name = "com_google_cloud_go_networkconnectivity",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/networkconnectivity",
        sum = "h1:LnrYM6lBEeTq+9f2lR4DjBhv31EROSAQi/P5W4Q0AEc=",
        version = "v1.12.1",
    )
    go_repository(
        name = "com_google_cloud_go_networkmanagement",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/networkmanagement",
        sum = "h1:/3xP37eMxnyvkfLrsm1nv1b2FbMMSAEAOlECTvoeCq4=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_networksecurity",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/networksecurity",
        sum = "h1:TBLEkMp3AE+6IV/wbIGRNTxnqLXHCTEQWoxRVC18TzY=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_google_cloud_go_notebooks",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/notebooks",
        sum = "h1:CUqMNEtv4EHFnbogV+yGHQH5iAQLmijOx191innpOcs=",
        version = "v1.9.1",
    )
    go_repository(
        name = "com_google_cloud_go_optimization",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/optimization",
        sum = "h1:pEwOAmO00mxdbesCRSsfj8Sd4rKY9kBrYW7Vd3Pq7cA=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_google_cloud_go_orchestration",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/orchestration",
        sum = "h1:KmN18kE/xa1n91cM5jhCh7s1/UfIguSCisw7nTMUzgE=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_orgpolicy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/orgpolicy",
        sum = "h1:I/7dHICQkNwym9erHqmlb50LRU588NPCvkfIY0Bx9jI=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_google_cloud_go_osconfig",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/osconfig",
        sum = "h1:dgyEHdfqML6cUW6/MkihNdTVc0INQst0qSE8Ou1ub9c=",
        version = "v1.12.1",
    )
    go_repository(
        name = "com_google_cloud_go_oslogin",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/oslogin",
        sum = "h1:LdSuG3xBYu2Sgr3jTUULL1XCl5QBx6xwzGqzoDUw1j0=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_google_cloud_go_phishingprotection",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/phishingprotection",
        sum = "h1:aK/lNmSd1vtbft/vLe2g7edXK72sIQbqr2QyrZN/iME=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_policytroubleshooter",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/policytroubleshooter",
        sum = "h1:XTMHy31yFmXgQg57CB3w9YQX8US7irxDX0Fl0VwlZyY=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_privatecatalog",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/privatecatalog",
        sum = "h1:B/18xGo+E0EMS9LOEQ0zXz7F2asMgmVgTYGSI89MHOA=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_google_cloud_go_pubsub",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/pubsub",
        sum = "h1:6SPCPvWav64tj0sVX/+npCBKhUi/UjJehy9op/V3p2g=",
        version = "v1.33.0",
    )
    go_repository(
        name = "com_google_cloud_go_pubsublite",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/pubsublite",
        sum = "h1:pX+idpWMIH30/K7c0epN6V703xpIcMXWRjKJsz0tYGY=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_recaptchaenterprise_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/recaptchaenterprise/v2",
        sum = "h1:IGkbudobsTXAwmkEYOzPCQPApUCsN4Gbq3ndGVhHQpI=",
        version = "v2.7.2",
    )
    go_repository(
        name = "com_google_cloud_go_recommendationengine",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/recommendationengine",
        sum = "h1:nMr1OEVHuDambRn+/y4RmNAmnR/pXCuHtH0Y4tCgGRQ=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_recommender",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/recommender",
        sum = "h1:UKp94UH5/Lv2WXSQe9+FttqV07x/2p1hFTMMYVFtilg=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_google_cloud_go_redis",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/redis",
        sum = "h1:YrjQnCC7ydk+k30op7DSjSHw1yAYhqYXFcOq1bSXRYA=",
        version = "v1.13.1",
    )
    go_repository(
        name = "com_google_cloud_go_resourcemanager",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/resourcemanager",
        sum = "h1:QIAMfndPOHR6yTmMUB0ZN+HSeRmPjR/21Smq5/xwghI=",
        version = "v1.9.1",
    )
    go_repository(
        name = "com_google_cloud_go_resourcesettings",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/resourcesettings",
        sum = "h1:Fdyq418U69LhvNPFdlEO29w+DRRjwDA4/pFamm4ksAg=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_google_cloud_go_retail",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/retail",
        sum = "h1:gYBrb9u/Hc5s5lUTFXX1Vsbc/9BEvgtioY6ZKaK0DK8=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_google_cloud_go_run",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/run",
        sum = "h1:kHeIG8q+N6Zv0nDkBjSOYfK2eWqa5FnaiDPH/7/HirE=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_google_cloud_go_scheduler",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/scheduler",
        sum = "h1:yoZbZR8880KgPGLmACOMCiY2tPk+iX4V/dkxqTirlz8=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_google_cloud_go_secretmanager",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/secretmanager",
        sum = "h1:cLTCwAjFh9fKvU6F13Y4L9vPcx9yiWPyWXE4+zkuEQs=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_google_cloud_go_security",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/security",
        sum = "h1:jR3itwycg/TgGA0uIgTItcVhA55hKWiNJxaNNpQJaZE=",
        version = "v1.15.1",
    )
    go_repository(
        name = "com_google_cloud_go_securitycenter",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/securitycenter",
        sum = "h1:XOGJ9OpnDtqg8izd7gYk/XUhj8ytjIalyjjsR6oyG0M=",
        version = "v1.23.0",
    )
    go_repository(
        name = "com_google_cloud_go_servicedirectory",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/servicedirectory",
        sum = "h1:pBWpjCFVGWkzVTkqN3TBBIqNSoSHY86/6RL0soSQ4z8=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_google_cloud_go_shell",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/shell",
        sum = "h1:aHbwH9LSqs4r2rbay9f6fKEls61TAjT63jSyglsw7sI=",
        version = "v1.7.1",
    )
    go_repository(
        name = "com_google_cloud_go_spanner",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/spanner",
        sum = "h1:aqiMP8dhsEXgn9K5EZBWxPG7dxIiyM2VaikqeU4iteg=",
        version = "v1.47.0",
    )
    go_repository(
        name = "com_google_cloud_go_speech",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/speech",
        sum = "h1:MCagaq8ObV2tr1kZJcJYgXYbIn8Ai5rp42tyGYw9rls=",
        version = "v1.19.0",
    )
    go_repository(
        name = "com_google_cloud_go_storage",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/storage",
        sum = "h1:STgFzyU5/8miMl0//zKh2aQeTyeaUH3WN9bSUiJ09bA=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_storagetransfer",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/storagetransfer",
        sum = "h1:+ZLkeXx0K0Pk5XdDmG0MnUVqIR18lllsihU/yq39I8Q=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_talent",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/talent",
        sum = "h1:j46ZgD6N2YdpFPux9mc7OAf4YK3tiBCsbLKc8rQx+bU=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_google_cloud_go_texttospeech",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/texttospeech",
        sum = "h1:S/pR/GZT9p15R7Y2dk2OXD/3AufTct/NSxT4a7nxByw=",
        version = "v1.7.1",
    )
    go_repository(
        name = "com_google_cloud_go_tpu",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/tpu",
        sum = "h1:kQf1jgPY04UJBYYjNUO+3GrZtIb57MfGAW2bwgLbR3A=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_google_cloud_go_trace",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/trace",
        sum = "h1:EwGdOLCNfYOOPtgqo+D2sDLZmRCEO1AagRTJCU6ztdg=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_google_cloud_go_translate",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/translate",
        sum = "h1:PQHamiOzlehqLBJMnM72lXk/OsMQewZB12BKJ8zXrU0=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_google_cloud_go_video",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/video",
        sum = "h1:BRyyS+wU+Do6VOXnb8WfPr42ZXti9hzmLKLUCkggeK4=",
        version = "v1.19.0",
    )
    go_repository(
        name = "com_google_cloud_go_videointelligence",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/videointelligence",
        sum = "h1:MBMWnkQ78GQnRz5lfdTAbBq/8QMCF3wahgtHh3s/J+k=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_google_cloud_go_vision_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/vision/v2",
        sum = "h1:ccK6/YgPfGHR/CyESz1mvIbsht5Y2xRsWCPqmTNydEw=",
        version = "v2.7.2",
    )
    go_repository(
        name = "com_google_cloud_go_vmmigration",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/vmmigration",
        sum = "h1:gnjIclgqbEMc+cF5IJuPxp53wjBIlqZ8h9hE8Rkwp7A=",
        version = "v1.7.1",
    )
    go_repository(
        name = "com_google_cloud_go_vmwareengine",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/vmwareengine",
        sum = "h1:qsJ0CPlOQu/3MFBGklu752v3AkD+Pdu091UmXJ+EjTA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_google_cloud_go_vpcaccess",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/vpcaccess",
        sum = "h1:ram0GzjNWElmbxXMIzeOZUkQ9J8ZAahD6V8ilPGqX0Y=",
        version = "v1.7.1",
    )
    go_repository(
        name = "com_google_cloud_go_webrisk",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/webrisk",
        sum = "h1:Ssy3MkOMOnyRV5H2bkMQ13Umv7CwB/kugo3qkAX83Fk=",
        version = "v1.9.1",
    )
    go_repository(
        name = "com_google_cloud_go_websecurityscanner",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/websecurityscanner",
        sum = "h1:CfEF/vZ+xXyAR3zC9iaC/QRdf1MEgS20r5UR17Q4gOg=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_google_cloud_go_workflows",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/workflows",
        sum = "h1:2akeQ/PgtRhrNuD/n1WvJd5zb7YyuDZrlOanBj2ihPg=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_shuralyov_dmitri_app_changes",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "dmitri.shuralyov.com/app/changes",
        sum = "h1:hJiie5Bf3QucGRa4ymsAUOxyhYwGEz1xrsVk0P8erlw=",
        version = "v0.0.0-20180602232624-0a106ad413e3",
    )
    go_repository(
        name = "com_shuralyov_dmitri_gpu_mtl",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "dmitri.shuralyov.com/gpu/mtl",
        sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
        version = "v0.0.0-20190408044501-666a987793e9",
    )
    go_repository(
        name = "com_shuralyov_dmitri_html_belt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "dmitri.shuralyov.com/html/belt",
        sum = "h1:SPOUaucgtVls75mg+X7CXigS71EnsfVUK/2CgVrwqgw=",
        version = "v0.0.0-20180602232347-f7d459c86be0",
    )
    go_repository(
        name = "com_shuralyov_dmitri_service_change",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "dmitri.shuralyov.com/service/change",
        sum = "h1:GvWw74lx5noHocd+f6HBMXK6DuggBB1dhVkuGZbv7qM=",
        version = "v0.0.0-20181023043359-a85b471d5412",
    )
    go_repository(
        name = "com_shuralyov_dmitri_state",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "dmitri.shuralyov.com/state",
        sum = "h1:ivON6cwHK1OH26MZyWDCnbTRZZf0IhNsENoNAKFS1g4=",
        version = "v0.0.0-20180228185332-28bcc343414c",
    )
    go_repository(
        name = "com_sourcegraph_sourcegraph_appdash",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "sourcegraph.com/sourcegraph/appdash",
        sum = "h1:ucqkfpjg9WzSUubAO62csmucvxl4/JeW3F4I4909XkM=",
        version = "v0.0.0-20190731080439-ebfcffb1b5c0",
    )
    go_repository(
        name = "com_sourcegraph_sourcegraph_go_diff",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "sourcegraph.com/sourcegraph/go-diff",
        sum = "h1:eTiIR0CoWjGzJcnQ3OkhIl/b9GJovq4lSAVRt0ZFEG8=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_sourcegraph_sqs_pbtypes",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "sourcegraph.com/sqs/pbtypes",
        sum = "h1:f7lAwqviDEGvON4kRv0o5V7FT/IQK+tbkF664XMbP3o=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_sslmate_software_src_go_pkcs12",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "software.sslmate.com/src/go-pkcs12",
        sum = "h1:SqYE5+A2qvRhErbsXFfUEUmpWEKxxRSMgGLkvRAFOV4=",
        version = "v0.0.0-20210415151418-c5206de65a78",
    )
    go_repository(
        name = "com_v2ray_core",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "v2ray.com/core",
        sum = "h1:JWoYsYlCpFOJX5KcmSkAMHOqOjzux+wx/HtgMBkUvSg=",
        version = "v4.19.1+incompatible",
    )
    go_repository(
        name = "com_zx2c4_golang_wintun",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.zx2c4.com/wintun",
        sum = "h1:Ug9qvr1myri/zFN6xL17LSCBGFDnphBBhzmILHsM5TY=",
        version = "v0.0.0-20211104114900-415007cec224",
    )
    go_repository(
        name = "com_zx2c4_golang_wireguard_windows",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.zx2c4.com/wireguard/windows",
        sum = "h1:OnYw96PF+CsIMrqWo5QP3Q59q5hY1rFErk/yN3cS+JQ=",
        version = "v0.5.1",
    )
    go_repository(
        name = "dev_gocloud",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gocloud.dev",
        sum = "h1:EDRyaRAnMGSq/QBto486gWFxMLczAfIYUmusV7XLNBM=",
        version = "v0.19.0",
    )
    go_repository(
        name = "in_gopkg_airbrake_gobrake_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/airbrake/gobrake.v2",
        sum = "h1:7z2uVWwn7oVeeugY1DtlPAy5H+KYgB1KeKTnqjNatLo=",
        version = "v2.0.9",
    )
    go_repository(
        name = "in_gopkg_alecthomas_kingpin_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/alecthomas/kingpin.v2",
        sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
        version = "v2.2.6",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/check.v1",
        sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
        version = "v1.0.0-20201130134442-10cb98267c6c",
    )
    go_repository(
        name = "in_gopkg_cheggaaa_pb_v1",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/cheggaaa/pb.v1",
        sum = "h1:n1tBJnnK2r7g9OW2btFH91V92STTUevLXYFb8gy9EMk=",
        version = "v1.0.28",
    )
    go_repository(
        name = "in_gopkg_datadog_dd_trace_go_v1",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/DataDog/dd-trace-go.v1",
        sum = "h1:EmglUJuykRsTwsQDcKaAo3CmOunWU6Dqk7U2lo7Pjss=",
        version = "v1.28.0",
    )
    go_repository(
        name = "in_gopkg_errgo_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/errgo.v2",
        sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
        version = "v2.1.0",
    )
    go_repository(
        name = "in_gopkg_fsnotify_fsnotify_v1",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/fsnotify/fsnotify.v1",
        sum = "h1:XNNYLJHt73EyYiCZi6+xjupS9CpvmiDgjPTAjrBlQbo=",
        version = "v1.4.7",
    )
    go_repository(
        name = "in_gopkg_fsnotify_v1",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/fsnotify.v1",
        sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
        version = "v1.4.7",
    )
    go_repository(
        name = "in_gopkg_gcfg_v1",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/gcfg.v1",
        sum = "h1:m8OOJ4ccYHnx2f4gQwpno8nAX5OGOh7RLaaz0pj3Ogs=",
        version = "v1.2.3",
    )
    go_repository(
        name = "in_gopkg_gemnasium_logrus_airbrake_hook_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/gemnasium/logrus-airbrake-hook.v2",
        sum = "h1:OAj3g0cR6Dx/R07QgQe8wkA9RNjB2u4i700xBkIT4e0=",
        version = "v2.1.2",
    )
    go_repository(
        name = "in_gopkg_h2non_gock_v1",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/h2non/gock.v1",
        sum = "h1:SzLqcIlb/fDfg7UvukMpNcWsu7sI5tWwL+KCATZqks0=",
        version = "v1.0.15",
    )
    go_repository(
        name = "in_gopkg_inconshreveable_log15_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/inconshreveable/log15.v2",
        sum = "h1:RlWgLqCMMIYYEVcAR5MDsuHlVkaIPDAF+5Dehzg8L5A=",
        version = "v2.0.0-20180818164646-67afb5ed74ec",
    )
    go_repository(
        name = "in_gopkg_inf_v0",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/inf.v0",
        sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
        version = "v0.9.1",
    )
    go_repository(
        name = "in_gopkg_ini_v1",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/ini.v1",
        sum = "h1:DPMeDvGTM54DXbPkVIZsp19fp/I2K7zwA/itHYHKo8Y=",
        version = "v1.56.0",
    )
    go_repository(
        name = "in_gopkg_mcuadros_go_syslog_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/mcuadros/go-syslog.v2",
        sum = "h1:60g8zx1BijSVSgLTzLCW9UC4/+i1Ih9jJ1DR5Tgp9vE=",
        version = "v2.2.1",
    )
    go_repository(
        name = "in_gopkg_natefinch_lumberjack_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/natefinch/lumberjack.v2",
        sum = "h1:bBRl1b0OH9s/DuPhuXpNl+VtCaJXFZ5/uEFST95x9zc=",
        version = "v2.2.1",
    )
    go_repository(
        name = "in_gopkg_natefinch_npipe_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/natefinch/npipe.v2",
        sum = "h1:+JknDZhAj8YMt7GC73Ei8pv4MzjDUNPHgQWJdtMAaDU=",
        version = "v2.0.0-20160621034901-c1b8fa8bdcce",
    )
    go_repository(
        name = "in_gopkg_ns1_ns1_go_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/ns1/ns1-go.v2",
        sum = "h1:GAcf+t0o8gdJAdSFYdE9wChu4bIyguMVqz0RHiFL5VY=",
        version = "v2.0.0-20190730140822-b51389932cbc",
    )
    go_repository(
        name = "in_gopkg_olivere_elastic_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/olivere/elastic.v2",
        sum = "h1:7cpl3MW8ysa4GYFBXklpo5mspe4NK0rpZTdyZ+QcD4U=",
        version = "v2.0.61",
    )
    go_repository(
        name = "in_gopkg_resty_v1",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/resty.v1",
        sum = "h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
        version = "v1.12.0",
    )
    go_repository(
        name = "in_gopkg_square_go_jose_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/square/go-jose.v2",
        sum = "h1:NGk74WTnPKBNUhNzQX7PYcTLUjoq7mzKk2OKbvwk2iI=",
        version = "v2.6.0",
    )
    go_repository(
        name = "in_gopkg_src_d_go_billy_v4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/src-d/go-billy.v4",
        sum = "h1:0SQA1pRztfTFx2miS8sA97XvooFeNOmvUenF4o0EcVg=",
        version = "v4.3.2",
    )
    go_repository(
        name = "in_gopkg_src_d_go_git_fixtures_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/src-d/go-git-fixtures.v3",
        sum = "h1:ivZFOIltbce2Mo8IjzUHAFoq/IylO9WHhNOAJK+LsJg=",
        version = "v3.5.0",
    )
    go_repository(
        name = "in_gopkg_src_d_go_git_v4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/src-d/go-git.v4",
        sum = "h1:SRtFyV8Kxc0UP7aCHcijOMQGPxHSmMOPrzulQWolkYE=",
        version = "v4.13.1",
    )
    go_repository(
        name = "in_gopkg_tomb_v1",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/tomb.v1",
        sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
        version = "v1.0.0-20141024135613-dd632973f1e7",
    )
    go_repository(
        name = "in_gopkg_warnings_v0",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/warnings.v0",
        sum = "h1:wFXVbFY8DY5/xOe1ECiWdKCzZlxgshcYVNkBHstARME=",
        version = "v0.1.2",
    )
    go_repository(
        name = "in_gopkg_yaml_v1",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/yaml.v1",
        sum = "h1:POO/ycCATvegFmVuPpQzZFJ+pGZeX22Ufu6fibxDVjU=",
        version = "v1.0.0-20140924161607-9f9df34309c0",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "io_etcd_go_bbolt",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/bbolt",
        sum = "h1:xs88BrvEv273UsB79e0hcVrlUWmS0a8upikMFhSyAtA=",
        version = "v1.3.8",
    )
    go_repository(
        name = "io_etcd_go_etcd",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd",
        sum = "h1:5aomL5mqoKHxw6NG+oYgsowk8tU8aOalo2IdZxdWHkw=",
        version = "v3.3.18+incompatible",
    )
    go_repository(
        name = "io_etcd_go_etcd_api_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/api/v3",
        sum = "h1:szRajuUUbLyppkhs9K6BRtjY37l66XQQmw7oZRANE4k=",
        version = "v3.5.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_pkg_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/client/pkg/v3",
        sum = "h1:kfYIdQftBnbAq8pUWFXfpuuxFSKzlmM5cSn76JByiT0=",
        version = "v3.5.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/client/v2",
        sum = "h1:MrmRktzv/XF8CvtQt+P6wLUlURaNpSDJHFZhe//2QE4=",
        version = "v2.305.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/client/v3",
        sum = "h1:W9TXNZ+oB3MCd/8UjxHTWK5J9Nquw9fQBLJd5ne5/Ao=",
        version = "v3.5.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_etcdctl_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/etcdctl/v3",
        sum = "h1:i8DGjR9gBRoS6NEHF3XBxxh7QwL1DyilXMCkHpyy6zM=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_etcd_go_etcd_etcdutl_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/etcdutl/v3",
        sum = "h1:orNfs85GWmiOl0p23Yi9YRfHNb3Qfdlt0wVFkPTRVxQ=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_etcd_go_etcd_pkg_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/pkg/v3",
        sum = "h1:WPR8K0e9kWl1gAhB5A7gEa5ZBTNkT9NdNWrR8Qpo1CM=",
        version = "v3.5.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_raft_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/raft/v3",
        sum = "h1:cgNAYe7xrsrn/5kXMSaH8kM/Ky8mAdMqGOxyYwpP0LA=",
        version = "v3.5.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_server_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/server/v3",
        sum = "h1:4NOGyOwD5sUZ22PiWYKmfxqoeh72z6EhYjNosKGLmZg=",
        version = "v3.5.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_tests_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/tests/v3",
        sum = "h1:+uMuHYKKlLUzbW322XrQXbaGM9qiV7vUL+AEPT/KYY4=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_etcd_go_etcd_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/v3",
        sum = "h1:fs7tB+L/xRDi/+p9qKuaPGCtMX6vkovLRXTqvEE98Ek=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_etcd_go_gofail",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/gofail",
        sum = "h1:XItAMIhOojXFQMgrxjnd2EIIHun/d5qL0Pf7FzVTkFg=",
        version = "v0.1.0",
    )
    go_repository(
        name = "io_filippo_edwards25519",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "filippo.io/edwards25519",
        sum = "h1:iJoUgXvhagsNMrJrvavw7vu1eG8+hm6jLOxlLFcoODw=",
        version = "v1.0.0-rc.1.0.20210721174708-390f27c3be20",
    )
    go_repository(
        name = "io_gitea_code_sdk_gitea",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "code.gitea.io/sdk/gitea",
        sum = "h1:3jCPOG2ojbl8AcfaUCRYLT5MUcBMFwS0OSK2mA5Zok8=",
        version = "v0.17.1",
    )
    go_repository(
        name = "io_h12_socks",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "h12.io/socks",
        sum = "h1:cZhhbV8+DE0Y1kotwhr1a3RC3kFO7AtuZ4GLr3qKSc8=",
        version = "v1.0.2",
    )
    go_repository(
        name = "io_k8s_api",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/api",
        sum = "h1:DAjwWX/9YT7NQD4INu49ROJuZAAAP/Ijki48GUPzxqw=",
        version = "v0.29.1",
    )
    go_repository(
        name = "io_k8s_apimachinery",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/apimachinery",
        sum = "h1:KY4/E6km/wLBguvCZv8cKTeOwwOBqFNjwJIdMkMbbRc=",
        version = "v0.29.1",
    )
    go_repository(
        name = "io_k8s_apiserver",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/apiserver",
        sum = "h1:e2wwHUfEmMsa8+cuft8MT56+16EONIEK8A/gpBSco+g=",
        version = "v0.29.1",
    )
    go_repository(
        name = "io_k8s_cli_runtime",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/cli-runtime",
        sum = "h1:By3WVOlEWYfyxhGko0f/IuAOLQcbBSMzwSaDren2JUs=",
        version = "v0.29.1",
    )
    go_repository(
        name = "io_k8s_client_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/client-go",
        sum = "h1:19B/+2NGEwnFLzt0uB5kNJnfTsbV8w6TgQRz9l7ti7A=",
        version = "v0.29.1",
    )
    go_repository(
        name = "io_k8s_cloud_provider",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/cloud-provider",
        sum = "h1:ELMIQwweSNu8gfVEnLDypxd9034S1sZJg6QcdWJOvMI=",
        version = "v0.17.4",
    )
    go_repository(
        name = "io_k8s_code_generator",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/code-generator",
        sum = "h1:8ba8BdtSmAVHgAMpzThb/fuyQeTRtN7NtN7VjMcDLew=",
        version = "v0.29.1",
    )
    go_repository(
        name = "io_k8s_component_base",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/component-base",
        sum = "h1:MUimqJPCRnnHsskTTjKD+IC1EHBbRCVyi37IoFBrkYw=",
        version = "v0.29.1",
    )
    go_repository(
        name = "io_k8s_component_helpers",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/component-helpers",
        sum = "h1:54MMEDu6xeJmMtAKztsPwu0kJKr4+jCUzaEIn2UXRoc=",
        version = "v0.29.1",
    )
    go_repository(
        name = "io_k8s_cri_api",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/cri-api",
        sum = "h1:0DHL/hpTf4Fp+QkUXFefWcp1fhjXr9OlNdY9X99c+O8=",
        version = "v0.23.1",
    )
    go_repository(
        name = "io_k8s_csi_translation_lib",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/csi-translation-lib",
        sum = "h1:bP9yGfCJDknP7tklCwizZtwgJNRePMVcEaFIfeA11ho=",
        version = "v0.17.4",
    )
    go_repository(
        name = "io_k8s_gengo",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/gengo",
        sum = "h1:pWEwq4Asjm4vjW7vcsmijwBhOr1/shsbSYiWXmNGlks=",
        version = "v0.0.0-20230829151522-9cce18d56c01",
    )
    go_repository(
        name = "io_k8s_klog",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/klog",
        sum = "h1:Pt+yjF5aB1xDSVbau4VsWe+dQNzA0qv1LlXdC2dF6Q8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "io_k8s_klog_v2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/klog/v2",
        sum = "h1:U/Af64HJf7FcwMcXyKm2RPM22WZzyR7OSpYj5tg3cL0=",
        version = "v2.110.1",
    )
    go_repository(
        name = "io_k8s_kms",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/kms",
        sum = "h1:6dMOaxllwiAZ8p3Hys65b78MDG+hONpBBpk1rQsaEtk=",
        version = "v0.29.1",
    )
    go_repository(
        name = "io_k8s_kube_openapi",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/kube-openapi",
        sum = "h1:aVUu9fTY98ivBPKR9Y5w/AuzbMm96cd3YHRTU83I780=",
        version = "v0.0.0-20231010175941-2dd684a91f00",
    )
    go_repository(
        name = "io_k8s_kubectl",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/kubectl",
        sum = "h1:rWnW3hi/rEUvvg7jp4iYB68qW5un/urKbv7fu3Vj0/s=",
        version = "v0.29.1",
    )
    go_repository(
        name = "io_k8s_kubernetes",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/kubernetes",
        sum = "h1:qTfB+u5M92k2fCCCVP2iuhgwwSOv1EkAkvQY1tQODD8=",
        version = "v1.13.0",
    )
    go_repository(
        name = "io_k8s_legacy_cloud_providers",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/legacy-cloud-providers",
        sum = "h1:VvFqJGiYAr2gIdoNuqbeZLEdxIFeN4Yt6OLJS9l2oIE=",
        version = "v0.17.4",
    )
    go_repository(
        name = "io_k8s_metrics",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/metrics",
        sum = "h1:qutc3aIPMCniMuEApuLaeYX47rdCn8eycVDx7R6wMlQ=",
        version = "v0.29.1",
    )
    go_repository(
        name = "io_k8s_sigs_apiserver_network_proxy_konnectivity_client",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/apiserver-network-proxy/konnectivity-client",
        sum = "h1:TgtAeesdhpm2SGwkQasmbeqDo8th5wOBA5h/AjTKA4I=",
        version = "v0.28.0",
    )
    go_repository(
        name = "io_k8s_sigs_json",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/json",
        sum = "h1:EDPBXCAspyGV4jQlpZSudPeMmr1bNJefnuqLsRAsHZo=",
        version = "v0.0.0-20221116044647-bc3834ca7abd",
    )
    go_repository(
        name = "io_k8s_sigs_kustomize_api",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/kustomize/api",
        sum = "h1:XX3Ajgzov2RKUdc5jW3t5jwY7Bo7dcRm+tFxT+NfgY0=",
        version = "v0.13.5-0.20230601165947-6ce0bf390ce3",
    )
    go_repository(
        name = "io_k8s_sigs_kustomize_cmd_config",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/kustomize/cmd/config",
        sum = "h1:YyoHHbxxsLUts/gWLGgIQkdT82ekp3zautbpcml54vc=",
        version = "v0.11.2",
    )
    go_repository(
        name = "io_k8s_sigs_kustomize_kustomize_v5",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/kustomize/kustomize/v5",
        sum = "h1:vq2TtoDcQomhy7OxXLUOzSbHMuMYq0Bjn93cDtJEdKw=",
        version = "v5.0.4-0.20230601165947-6ce0bf390ce3",
    )
    go_repository(
        name = "io_k8s_sigs_kustomize_kyaml",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/kustomize/kyaml",
        sum = "h1:W6cLQc5pnqM7vh3b7HvGNfXrJ/xL6BDMS0v1V/HHg5U=",
        version = "v0.14.3-0.20230601165947-6ce0bf390ce3",
    )
    go_repository(
        name = "io_k8s_sigs_structured_merge_diff",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/structured-merge-diff",
        sum = "h1:zD2IemQ4LmOcAumeiyDWXKUI2SO0NYDe3H6QGvPOVgU=",
        version = "v1.0.1-0.20191108220359-b1b620dd3f06",
    )
    go_repository(
        name = "io_k8s_sigs_structured_merge_diff_v4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/structured-merge-diff/v4",
        sum = "h1:150L+0vs/8DA78h1u02ooW1/fFq/Lwr+sGiqlzvrtq4=",
        version = "v4.4.1",
    )
    go_repository(
        name = "io_k8s_sigs_yaml",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/yaml",
        sum = "h1:a2VclLzOGrwOHDiV8EfBGhvjHvP46CtW5j6POvhYGGo=",
        version = "v1.3.0",
    )
    go_repository(
        name = "io_k8s_utils",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/utils",
        sum = "h1:sgn3ZU783SCgtaSJjpcVVlRqd6GSnlTLKgpAAttJvpI=",
        version = "v0.0.0-20230726121419-3b25d923346b",
    )
    go_repository(
        name = "io_nhooyr_websocket",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "nhooyr.io/websocket",
        sum = "h1:mv4p+MnGrLDcPlBoWsvPP7XCzTYMXP9F9eIGoKbgx7Q=",
        version = "v1.8.10",
    )
    go_repository(
        name = "io_opencensus_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opencensus.io",
        sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
        version = "v0.24.0",
    )
    go_repository(
        name = "io_opencensus_go_contrib_exporter_aws",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "contrib.go.opencensus.io/exporter/aws",
        sum = "h1:YsbWYxDZkC7x2OxlsDEYvvEXZ3cBI3qBgUK5BqkZvRw=",
        version = "v0.0.0-20181029163544-2befc13012d0",
    )
    go_repository(
        name = "io_opencensus_go_contrib_exporter_ocagent",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "contrib.go.opencensus.io/exporter/ocagent",
        sum = "h1:TKXjQSRS0/cCDrP7KvkgU6SmILtF/yV2TOs/02K/WZQ=",
        version = "v0.5.0",
    )
    go_repository(
        name = "io_opencensus_go_contrib_exporter_stackdriver",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "contrib.go.opencensus.io/exporter/stackdriver",
        sum = "h1:hAIKGfweGma95MG+PZZTN2rD2Vc1DKmUmPeRYlG6j7c=",
        version = "v0.13.7",
    )
    go_repository(
        name = "io_opencensus_go_contrib_integrations_ocsql",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "contrib.go.opencensus.io/integrations/ocsql",
        sum = "h1:kfg5Yyy1nYUrqzyfW5XX+dzMASky8IJXhtHe0KTYNS4=",
        version = "v0.1.4",
    )
    go_repository(
        name = "io_opencensus_go_contrib_resource",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "contrib.go.opencensus.io/resource",
        sum = "h1:4r2CANuYhKGmYWP02+5E94rLRcS/YeD+KlxSrOsMxk0=",
        version = "v0.1.1",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib",
        sum = "h1:ubFQUn0VCZ0gPwIoJfBJVpeBlyRMxu8Mm/huKWYd9p0=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
        sum = "h1:ZOLJc06r4CB42laIXg/7udr0pbZyuAihN10A/XuiQRY=",
        version = "v0.42.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp",
        sum = "h1:KfYpVmrjI7JuToy5k8XV3nkapjWx48k4E4JOtVstzQI=",
        version = "v0.44.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel",
        sum = "h1:MuS/TNf4/j4IXsZuJegVzI1cwut7Qc00344rgH7p8bs=",
        version = "v1.19.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp",
        sum = "h1:PTNgq9MRmQqqJY0REVbZFvwkYOA85vbdQU/nVfxDyqg=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_internal_retry",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/internal/retry",
        sum = "h1:j7AwzDdAQBJjcqayAaYbvpYeZzII7cEe5qJTu+De6UY=",
        version = "v1.4.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace",
        build_directives = [
            "gazelle:resolve go go.opentelemetry.io/otel/exporters/otlp/internal @io_opentelemetry_go_otel//exporters/otlp/internal",
            "gazelle:resolve go go.opentelemetry.io/otel/exporters/otlp/internal/envconfig @io_opentelemetry_go_otel//exporters/otlp/internal/envconfig",
        ],
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace",
        sum = "h1:Mne5On7VWdx7omSrSSZvM4Kw7cS7NQkOOmLcgscI51U=",
        version = "v1.19.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracegrpc",
        build_directives = [
            "gazelle:resolve go go.opentelemetry.io/otel/exporters/otlp/internal @io_opentelemetry_go_otel//exporters/otlp/internal",
            "gazelle:resolve go go.opentelemetry.io/otel/exporters/otlp/internal/envconfig @io_opentelemetry_go_otel//exporters/otlp/internal/envconfig",
        ],
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc",
        sum = "h1:3d+S281UTjM+AbF31XSOYn1qXn3BgIdWl8HNEpx08Jk=",
        version = "v1.19.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracehttp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp",
        sum = "h1:Ydage/P0fRrSPpZeCVxzjqGcI6iVmG2xb43+IR8cjqM=",
        version = "v1.3.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_internal_metric",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/internal/metric",
        sum = "h1:9dAVGAfFiiEq5NVB9FUJ5et+btbDQAUIJehJ+ikyryk=",
        version = "v0.27.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_metric",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/metric",
        sum = "h1:aTzpGtV0ar9wlV4Sna9sdJyII5jTVJEvKETPiOKwvpE=",
        version = "v1.19.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_oteltest",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/oteltest",
        sum = "h1:HiITxCawalo5vQzdHfKeZurV8x7ljcqAgiWzF6Vaeaw=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/sdk",
        sum = "h1:6USY6zH+L8uMH8L3t1enZPR3WFEmSTADlqldyHtJi3o=",
        version = "v1.19.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk_export_metric",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/sdk/export/metric",
        sum = "h1:c5VRjxCXdQlx1HjzwGdQHzZaVI82b5EbBgOu2ljD92g=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk_metric",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/sdk/metric",
        sum = "h1:7ao1wpzHRVKf0OQ7GIxiQJA6X7DLX9o14gmVon7mMK8=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_trace",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/trace",
        sum = "h1:DFVQmlVbfVeOuBRrwdtaehRrWiL1JoVs9CPIQ1Dzxpg=",
        version = "v1.19.0",
    )
    go_repository(
        name = "io_opentelemetry_go_proto_otlp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/proto/otlp",
        sum = "h1:T0TX0tmXU8a3CbNXzEKGeU5mIVOdf0oykP+u2lIVU/I=",
        version = "v1.0.0",
    )
    go_repository(
        name = "io_robpike_ivy",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "robpike.io/ivy",
        sum = "h1:C0QJH9Il77S2MMd+8Fkc2VKv0B9lXd0yexiRzQrFUIE=",
        version = "v0.3.4",
    )
    go_repository(
        name = "io_rsc_binaryregexp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "rsc.io/binaryregexp",
        sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
        version = "v0.2.0",
    )
    go_repository(
        name = "io_rsc_pdf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "rsc.io/pdf",
        sum = "h1:k1MczvYDUvJBe93bYd7wrZLLUEcLZAuF824/I4e5Xr4=",
        version = "v0.1.1",
    )
    go_repository(
        name = "io_rsc_quote_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "rsc.io/quote/v3",
        sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
        version = "v3.1.0",
    )
    go_repository(
        name = "io_rsc_sampler",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "rsc.io/sampler",
        sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
        version = "v1.3.0",
    )
    go_repository(
        name = "net_howett_plist",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "howett.net/plist",
        sum = "h1:7CrbWYbPPO/PyNy38b2EB/+gYbjCe2DXBxgtOOZbSQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "net_starlark_go",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.starlark.net",
        sum = "h1:VdD38733bfYv5tUZwEIskMM93VanwNIi5bIKnDrJdEY=",
        version = "v0.0.0-20230525235612-a134d8f9ddca",
    )
    go_repository(
        name = "org_apache_git_thrift_git",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "git.apache.org/thrift.git",
        sum = "h1:OR8VhtwhcAI3U48/rzBsVOuHi0zDPzYI1xASVcdSgR8=",
        version = "v0.0.0-20180902110319-2566ecd5d999",
    )
    go_repository(
        name = "org_bazil_fuse",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "bazil.org/fuse",
        sum = "h1:SRsZGA7aFnCZETmov57jwPrWuTmaZK6+4R4v5FUe1/c=",
        version = "v0.0.0-20200407214033-5883e5a4b512",
    )
    go_repository(
        name = "org_bitbucket_creachadair_shell",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "bitbucket.org/creachadair/shell",
        sum = "h1:reJflDbKqnlnqb4Oo2pQ1/BqmY/eCWcNGHrIUO8qIzc=",
        version = "v0.0.6",
    )
    go_repository(
        name = "org_go4",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go4.org",
        sum = "h1:+hE86LblG4AyDgwMCLTE6FOlM9+qjHSYS+rKqxUVdsM=",
        version = "v0.0.0-20180809161055-417644f6feb5",
    )
    go_repository(
        name = "org_go4_grpc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "grpc.go4.org",
        sum = "h1:tmXTu+dfa+d9Evp8NpJdgOy6+rt8/x4yG7qPBrtNfLY=",
        version = "v0.0.0-20170609214715-11d0a25b4919",
    )
    go_repository(
        name = "org_golang_google_api",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/api",
        sum = "h1:q4GJq+cAdMAC7XP7njvQ4tvohGLiSlytuL4BQxbIZ+o=",
        version = "v0.126.0",
    )
    go_repository(
        name = "org_golang_google_appengine",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/appengine",
        sum = "h1:IhEN5q69dyKagZPYMSdIjS2HqprW324FRQZJcGqPAsM=",
        version = "v1.6.8",
    )
    go_repository(
        name = "org_golang_google_cloud",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/cloud",
        sum = "h1:Cpp2P6TPjujNoC5M2KHY6g7wfyLYfIWRZaSdIKfDasA=",
        version = "v0.0.0-20151119220103-975617b05ea8",
    )
    go_repository(
        name = "org_golang_google_genproto",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/genproto",
        sum = "h1:L6iMMGrtzgHsWofoFcihmDEMYeDR9KN/ThbPWGrh++g=",
        version = "v0.0.0-20230803162519-f966b187b2e5",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:z3vDksarJxsAKM5dmEGv0GHwE2hKJ096wZra71Vs4sw=",
        version = "v0.0.0-20230726155614-23370e0ffb3e",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_bytestream",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/genproto/googleapis/bytestream",
        sum = "h1:g3hIDl0jRNd9PPTs2uBzYuaD5mQuwOkZY0vSc0LR32o=",
        version = "v0.0.0-20230530153820-e85fd2cbaebc",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:uvYuEyMHKNt+lT4K3bN6fGswmK8qSvcreM3BwjDh+y4=",
        version = "v0.0.0-20230822172742-b8732ec3820d",
    )
    go_repository(
        name = "org_golang_google_grpc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc",
        sum = "h1:BjnpXut1btbtgN/6sp+brB2Kbm2LjNXnidYujAVbSoQ=",
        version = "v1.58.3",
    )
    go_repository(
        name = "org_golang_google_grpc_cmd_protoc_gen_go_grpc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
        sum = "h1:M1YKkFIboKNieVO5DLUEVzQfGwJD30Nv2jfUgzb5UcE=",
        version = "v1.1.0",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/protobuf",
        sum = "h1:pPC6BG5ex8PDFnkbrGU3EixyhKcQ2aDuBS36lqK/C7I=",
        version = "v1.32.0",
    )
    go_repository(
        name = "org_golang_x_arch",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/arch",
        sum = "h1:YcaRWbSY2VfP0/k25uHKKrk3Vs3C7mo03vq103Ire8E=",
        version = "v0.0.0-20190909030613-46d78d1859ac",
    )
    go_repository(
        name = "org_golang_x_build",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/build",
        sum = "h1:E2M5QgjZ/Jg+ObCQAudsXxuTsLj7Nl5RV/lZcQZmKSo=",
        version = "v0.0.0-20190111050920-041ab4dc3f9d",
    )
    go_repository(
        name = "org_golang_x_crypto",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/crypto",
        sum = "h1:PGVlW0xEltQnzFZ55hkuX5+KLyrMYhHld1YHO4AKcdc=",
        version = "v0.18.0",
    )
    go_repository(
        name = "org_golang_x_exp",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/exp",
        sum = "h1:Q8/wZp0KX97QFTc2ywcOE0YRjZPVIx+MXInMzdvQqcA=",
        version = "v0.0.0-20240119083558-1b970713d09a",
    )
    go_repository(
        name = "org_golang_x_image",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/image",
        sum = "h1:hVwzHzIUGRjiF7EcUjqNxk3NCfkPxbDKRdnNE1Rpg0U=",
        version = "v0.0.0-20191009234506-e7c1f5e7dbb8",
    )
    go_repository(
        name = "org_golang_x_lint",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/lint",
        sum = "h1:VLliZ0d+/avPrXXH+OakdXhpJuEoBZuwh1m2j7U6Iug=",
        version = "v0.0.0-20210508222113-6edffad5e616",
    )
    go_repository(
        name = "org_golang_x_mobile",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/mobile",
        sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
        version = "v0.0.0-20190719004257-d2bd2a29d028",
    )
    go_repository(
        name = "org_golang_x_mod",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/mod",
        sum = "h1:dGoOF9QVLYng8IHTm7BAyWqCqSheQ5pYWGhzW00YJr0=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_net",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/net",
        sum = "h1:aCL9BSgETF1k+blQaYUBx9hJ9LOGP3gAVemcZlf1Kpo=",
        version = "v0.20.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/oauth2",
        sum = "h1:aDkGMBSYxElaoP81NpoUoz2oo2R2wHdZpGToUxfyQrQ=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_golang_x_perf",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/perf",
        sum = "h1:xYq6+9AtI+xP3M4r0N1hCkHrInHDBohhquRgx9Kk6gI=",
        version = "v0.0.0-20180704124530-6e6d33e29852",
    )
    go_repository(
        name = "org_golang_x_sync",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/sync",
        sum = "h1:5BMeUDZ7vkXGfEr1x9B4bRcTH4lpkTkpdh0T/J+qjbQ=",
        version = "v0.6.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/sys",
        sum = "h1:xWw16ngr6ZMtmxDyKyIgsE93KNKz5HKmMa3b8ALHidU=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_golang_x_term",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/term",
        sum = "h1:m+B6fahuftsE9qjo0VWp2FW0mB3MTJvR0BaMQrq0pmE=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_golang_x_text",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/text",
        sum = "h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_time",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/time",
        sum = "h1:rg5rLMjNzMS1RkNLzCG38eapWhnYLFYXDXj2gOlr8j4=",
        version = "v0.3.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/tools",
        sum = "h1:FvmRgNOcs3kOa+T20R1uhfP9F6HgG2mfxDv1vrx1Htc=",
        version = "v0.17.0",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/xerrors",
        sum = "h1:H2TDz8ibqkAF6YGhCdN3jS9O0/s90v0rJh3X/OLHEUk=",
        version = "v0.0.0-20220907171357-04be3eba64a2",
    )
    go_repository(
        name = "org_gonum_v1_gonum",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gonum.org/v1/gonum",
        sum = "h1:OB/uP/Puiu5vS5QMRPrXCDWUPb+kt8f1KW8oQzFejQw=",
        version = "v0.0.0-20190331200053-3d26580ed485",
    )
    go_repository(
        name = "org_gonum_v1_netlib",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gonum.org/v1/netlib",
        sum = "h1:jRyg0XfpwWlhEV8mDfdNGBeSJM2fuyh9Yjrnd8kF2Ts=",
        version = "v0.0.0-20190331212654-76723241ea4e",
    )
    go_repository(
        name = "org_modernc_cc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "modernc.org/cc",
        sum = "h1:nPibNuDEx6tvYrUAtvDTTw98rx5juGsa5zuDnKwEEQQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_golex",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "modernc.org/golex",
        sum = "h1:wWpDlbK8ejRfSyi0frMyhilD3JBvtcx2AdGDnU+JtsE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_httpfs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "modernc.org/httpfs",
        sum = "h1:AAgIpFZRXuYnkjftxTAZwMIiwEqAfk8aVB2/oA6nAeM=",
        version = "v1.0.6",
    )
    go_repository(
        name = "org_modernc_mathutil",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "modernc.org/mathutil",
        sum = "h1:93vKjrJopTPrtTNpZ8XIovER7iCIH1QU7wNbOQXC60I=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_strutil",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "modernc.org/strutil",
        sum = "h1:XVFtQwFVwc02Wk+0L/Z/zDDXO81r5Lhe6iMKmGX3KhE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_xc",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "modernc.org/xc",
        sum = "h1:7ccXrupWZIS3twbUGrtKmHS2DXY6xegFua+6O3xgAFU=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_mozilla_go_pkcs7",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.mozilla.org/pkcs7",
        sum = "h1:CCriYyAfq1Br1aIYettdHZTy8mBTIPo7We18TuO/bak=",
        version = "v0.0.0-20210826202110-33d05740a352",
    )
    go_repository(
        name = "org_torproject_git_pluggable_transports_goptlib_git",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "git.torproject.org/pluggable-transports/goptlib.git",
        sum = "h1:0qRF7Dw5qXd0FtZkjWUiAh5GTutRtDGL4GXUDJ4qMHs=",
        version = "v1.2.0",
    )
    go_repository(
        name = "org_uber_go_atomic",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/atomic",
        sum = "h1:9qC72Qh0+3MqyJbAn8YU5xVq1frD8bn3JtD2oXtafVQ=",
        version = "v1.10.0",
    )
    go_repository(
        name = "org_uber_go_automaxprocs",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/automaxprocs",
        sum = "h1:e1YG66Lrk73dn4qhg8WFSvhF0JuFQF0ERIp4rpuV8Qk=",
        version = "v1.5.1",
    )
    go_repository(
        name = "org_uber_go_goleak",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/goleak",
        sum = "h1:NBol2c7O1ZokfZ0LEU9K6Whx/KnwvepVetCUhtKja4A=",
        version = "v1.2.1",
    )
    go_repository(
        name = "org_uber_go_mock",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/mock",
        sum = "h1:3mUxI1No2/60yUYax92Pt8eNOEecx2D3lcXZh2NEZJo=",
        version = "v0.3.0",
    )
    go_repository(
        name = "org_uber_go_multierr",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/multierr",
        sum = "h1:blXXJkSxSSfBVBlC76pxqeO+LN3aDfLQo+309xJstO0=",
        version = "v1.11.0",
    )
    go_repository(
        name = "org_uber_go_ratelimit",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/ratelimit",
        sum = "h1:d9qaMM+ODpCq+9We41//fu/sHsTnXcrqd1en3x+GKy4=",
        version = "v0.0.0-20180316092928-c15da0234277",
    )
    go_repository(
        name = "org_uber_go_tools",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/tools",
        sum = "h1:0mgffUl7nfd+FpvXMVz4IDEaUSmT1ysygQC7qYo7sG4=",
        version = "v0.0.0-20190618225709-2cfd321de3ee",
    )
    go_repository(
        name = "org_uber_go_zap",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/zap",
        sum = "h1:WefMeulhovoZ2sYXz7st6K0sLj7bBhpiFaud4r4zST8=",
        version = "v1.21.0",
    )
    go_repository(
        name = "sh_elv_src",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "src.elv.sh",
        sum = "h1:stqxghJ8lnlVknA37nlryttw+uizKuVJpZdp0hljNaw=",
        version = "v0.14.1-0.20210218105754-53593c3ab79f",
    )
    go_repository(
        name = "sm_step_go_cli_utils",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.step.sm/cli-utils",
        sum = "h1:2GvY5Muid1yzp7YQbfCCS+gK3q7zlHjjLL5Z0DXz8ds=",
        version = "v0.7.0",
    )
    go_repository(
        name = "sm_step_go_crypto",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.step.sm/crypto",
        sum = "h1:4mnZk21cSxyMGxsEpJwZKKvJvDu1PN09UVrWWFNUBdk=",
        version = "v0.16.1",
    )
    go_repository(
        name = "sm_step_go_linkedca",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "go.step.sm/linkedca",
        sum = "h1:lEkGRDY+u7FudGKt8yEo7nBy5OzceO9s3rl+/sZVL5M=",
        version = "v0.15.0",
    )
    go_repository(
        name = "tools_gotest",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gotest.tools",
        sum = "h1:VsBPFP1AI068pPrMxtb/S8Zkgf9xEmTLJjfM+P5UIEo=",
        version = "v2.2.0+incompatible",
    )
    go_repository(
        name = "tools_gotest_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gotest.tools/v3",
        sum = "h1:4AuOwCGf4lLR9u3YOe2awrHygurzhO/HeQ6laiA6Sx0=",
        version = "v3.0.3",
    )
    go_repository(
        name = "xyz_gomodules_jsonpatch_v3",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gomodules.xyz/jsonpatch/v3",
        sum = "h1:Te7hKxV52TKCbNYq3t84tzKav3xhThdvSsSp/W89IyI=",
        version = "v3.0.1",
    )
    go_repository(
        name = "xyz_gomodules_orderedmap",
        build_file_generation = "on",
        build_file_proto_mode = "disable",
        importpath = "gomodules.xyz/orderedmap",
        sum = "h1:fM/+TGh/O1KkqGR5xjTKg6bU8OKBkg7p0Y+x/J9m8Os=",
        version = "v0.1.0",
    )
