FROM golang AS builder
RUN apt install -y make
# RUN git clone https://github.com/btwiuse/k0s /go/src/github.com/btwiuse/k0s
COPY . /go/src/github.com/btwiuse/k0s
WORKDIR /go/src/github.com/btwiuse/k0s
RUN make build && make install

FROM alpine:latest
RUN apk add curl
COPY --from=builder /usr/bin/k0s /usr/bin/k0s
ENTRYPOINT ["/usr/bin/k0s"]
