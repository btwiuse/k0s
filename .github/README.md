# 🕸️ k0s

kubotnetes, botnet as a service, pronounced like 'chaos'

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/k0s.io?tab=doc)
[![Go 1.18+](https://img.shields.io/github/go-mod/go-version/btwiuse/k0s)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/btwiuse/k0s?color=%23000&style=flat-round)](https://github.com/btwiuse/k0s/blob/master/LICENSE)
[![DockerHub](https://img.shields.io/docker/pulls/btwiuse/k0s.svg)](https://hub.docker.com/r/btwiuse/k0s)
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/btwiuse/k0s)

## Why

As a long-term Chrome OS user, I'm really into the idea of doing everything in the browser.

Chromebooks, however, are quite limited in their computing power.

When compiling large Rust projects, I still need to resort to my Ryzen 9 5950x home server.

k0s was created to solve the problem.

It works by exposing my home server on the web, allowing me to connect to it from anywhere.

Sounds like a botnet, huh? Quite close, except that it's operated by you on your own infrastructure.

You can try it at [k0s.io](https://k0s.io).

Some targets are not protected by a password, they act as honeypots for the purpose of demonstration.

If you destroy them, they will be recovered later.


## DISCLAIMER

This project is not to be confused with [k0s - The Zero Friction Kubernetes by Team Lens](https://github.com/k0sproject/k0s).

I started k0s as a hobby project and registered [k0s.io](https://k0s.io) in 2019.

The Zero Friction Kubernetes project was launched with the domain name [k0sproject.io](https://k0sproject.io) in late 2020.

Some people may think that by using the package name `k0s`, I am pretending to be the other project and performing a supply chain attack, which is not true.

Several Chinese sources recently even flagged the npm installer as malware.

Now I must make it explicit:

* This project is for demonstration and education purposes only

* It is not a distro of Kubernetes

* It is not malware and contains no malicious code

* It was created in hopes that it would be useful to others, not harmful

* By using it you acknowlege that your are at your own risk

* Author is not responsible for misuse

