FROM golang as builder
ENV CGO_ENABLED=0
RUN git clone https://github.com/invctrl/hijack
RUN go get -d ./hijack/client/...
RUN go build -o /bin/client -ldflags="-X 'main.commit=$(git -C hijack rev-list -1 HEAD)' -X 'main.built=$(date)' -X 'main.branch=$(git -C hijack rev-parse --abbrev-ref HEAD)'" ./hijack/client

FROM invctrl/client-base
COPY --from=builder /bin/client /bin/client
ENTRYPOINT ["/bin/client"]
