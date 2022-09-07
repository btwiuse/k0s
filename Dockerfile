FROM btwiuse/arch:golang

RUN env GOBIN=/usr/local/bin/ go install .

CMD k3s server --disable-agent
