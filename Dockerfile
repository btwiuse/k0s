# https://cirrus-ci.com/github/btwiuse/k0s

FROM btwiuse/bazel AS builder
COPY . /src/k0s.io/k0s
RUN make -C /src/k0s.io/k0s bazel-build

FROM alpine
RUN apk add curl
COPY --from=builder /src/k0s.io/k0s/bin/k0s_static /usr/bin/k0s
ENTRYPOINT ["/usr/bin/k0s"]
