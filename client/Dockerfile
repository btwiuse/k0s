FROM golang as builder
ENV CGO_ENABLED=0
RUN git clone https://github.com/navigaid/hijack
RUN go build -o /bin/client hijack/client.go

FROM archlinux/base:latest
COPY --from=builder /bin/client /bin/client
ENTRYPOINT ["/bin/client"]
