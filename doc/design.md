Architecture:

```
(1) Session ----- NodeAgent ---\         /--- (1) Browser
                /               \       /
(2) Session ---/  NodeAgent ----- [Hub] ----- (2) Cmdline
                                /       \
(3) Session ----- NodeAgent ---/         \--- (3) Browser
```

GOAL:

- [DONE] control http client from server
- [DONE] switch between connected clients
- [DONE] stream client command output

Requirements:

- [DONE] HTTP(S) -> TCP
- [DONE] TCP -> RPC
- [DONE] TCP -> gRPC
- [DONE] gRPC <=> Websocket

Similar stuff:

- https://en.wikipedia.org/wiki/TCP_Gender_Changer
- https://github.com/rancher/remotedialer
- https://github.com/rofl0r/nat-tunnel
- https://github.com/TreeHacks/botnet-hackpack
- https://github.com/peoples-cloud/pc

Potential use cases:

- distcc agent (cpu/io intensive)
- anonymous proxy endpoint (io intensive)
- crypto mining worker (cpu intensive)
- ddos attack agent (generate SYN flood, ICMP, UDP, TCP traffic)
- distributed web scraper (cpu/io intensive)
- podman/buildah/skopeo build log streaming with dind (Docker in Docker)

References:

- https://en.wikipedia.org/wiki/Botnet
- https://www.malwaretech.com/2013/12/peer-to-peer-botnets-for-beginners.html
- https://umbrella.cisco.com/blog/2014/05/23/cnc-botnet-english/
- https://ops.tips/gists/example-go-rpc-client-and-server/
- https://github.com/topics/botnet?l=Go
- https://rosettacode.org/wiki/Distributed_programming
