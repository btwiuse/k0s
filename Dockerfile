# https://cirrus-ci.com/github/btwiuse/k0s

FROM btwiuse/bazel AS builder
COPY . /k0s.io/k0s
WORKDIR /k0s.io/k0s
RUN make bazel-build

FROM alpine
RUN apk add curl
COPY --from=builder /k0s.io/k0s/bin/k0s_static /usr/bin/k0s
ENTRYPOINT ["/usr/bin/k0s"]
