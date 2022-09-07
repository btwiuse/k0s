# https://cirrus-ci.com/github/btwiuse/k0s

# FROM btwiuse/arch:bazel AS builder-bazel
FROM btwiuse/k0s:devcontainer AS builder-bazel
COPY . /k0s.io
WORKDIR /k0s.io
RUN make bazel-build

# FROM btwiuse/arch:golang AS builder-golang
# COPY . /k0s.io
# WORKDIR /k0s.io
# RUN make build

FROM btwiuse/arch
COPY --from=builder-bazel /k0s.io/bin/k0s_static /usr/bin/k0s
# COPY --from=builder-golang /k0s.io/bin/k0s /usr/bin/k0s
ENTRYPOINT ["/usr/bin/k0s"]
