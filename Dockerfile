FROM btwiuse/bazel AS builder
COPY . /src
WORKDIR /src
RUN make bazel-build

FROM alpine:latest
RUN apk add curl
COPY --from=builder /src/bin/k0s_static /usr/bin/k0s
ENTRYPOINT ["/usr/bin/k0s"]
