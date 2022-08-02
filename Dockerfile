# https://cirrus-ci.com/github/btwiuse/k0s

# FROM btwiuse/bazel AS builder
# COPY . /k0s.io
# WORKDIR /k0s.io
# RUN make bazel-build

FROM golang:1.19 AS builder-go
COPY . /k0s.io
WORKDIR /k0s.io
RUN make build

FROM alpine
RUN apk add curl
# COPY --from=builder /k0s.io/bin/k0s_static /usr/bin/k0s
COPY --from=builder-go /k0s.io/bin/k0s /usr/bin/k0s
ENTRYPOINT ["/usr/bin/k0s"]
