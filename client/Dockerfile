FROM golang as builder
ENV CGO_ENABLED=0
RUN git clone https://github.com/navigaid/hijack
RUN go get -d ./hijack/client/...
RUN go build -o /bin/client ./hijack/client

FROM docker:dind
RUN apk add --no-cache bash git vim jq moreutils openssl curl
COPY --from=builder /bin/client /bin/client
ENTRYPOINT ["/bin/client"]
