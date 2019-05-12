GOAL:
- [DONE] control http client from server 
- [DONE] switch between connected clients
- [TODO] stream client command output    

Requirements:
- [DONE] HTTP -> TCP                     

Similar stuff:
- https://en.wikipedia.org/wiki/TCP_Gender_Changer
- https://github.com/rancher/remotedialer
- https://github.com/rofl0r/nat-tunnel

Potential use cases:
- distcc agent (cpu/io intensive)
- anonymous proxy endpoint (io intensive)
- crypto mining worker (cpu intensive)
- ddos attack agent (generate SYN flood, ICMP, UDP, TCP traffic)
- distributed web scraper (cpu/io intensive)
- podman/buildah/skopeo build log streaming with dind (Docker in Docker)
