# Kubernot - Not Kubernetes

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/k0s.io?tab=doc)
[![Go 1.18+](https://img.shields.io/github/go-mod/go-version/btwiuse/k0s)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/btwiuse/k0s?color=%23000&style=flat-round)](https://github.com/btwiuse/k0s/blob/master/LICENSE)
[![Visitors](https://visitor-badge.glitch.me/badge?page_id=btwiuse.k0s)](#)
[![DockerHub](https://img.shields.io/docker/pulls/btwiuse/k0s.svg)](https://hub.docker.com/r/btwiuse/k0s)
[![@kubernot](https://img.shields.io/twitter/url/https/twitter.com/kubernot.svg?style=social&label=Follow%20%40kubernot)](https://twitter.com/kubernot)

Kubernot is an experimental SSH alternative with Kubernetes-like UX, where

- [ ] `scp` becomes `kubectl cp`
- [ ] `ssh` becomes `kubectl exec`
- [ ] `ssh -L` becomes `kubectl port-forward`

Unlike SSH, Kubernot adopts the client-agent-server model from Kubernetes that allows access to nodes behind a firewall.

||client|agent|server|
|:--:|:--:|:--:|:--:|
|SSH|`ssh`|-|`sshd`|
|Kubernot|`kubectl`|`knot`|`kuber`|
|Kubernetes|`kubectl`|`kubelet`|`kube-apiserver`|

Kubernot has two major components: __kuber__ and __knot__:

- Kuber implements a minimal subset of [The Kubernetes API](https://kubernetes.io/docs/concepts/overview/kubernetes-api/), and is therefore compatible with kubectl. 

  It is directly modelled after [kube-apiserver](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/).

- Knot is the equivalent of [kubelet](https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/) in Kubernetes.

  It runs on targets like Windows / BSD / Android / Chrome (via WASI), since it only deals with processes, sockets and files - no container support is required.

## Not Kubernetes

    Knowledge is power,

    France is bacon,

    but Kubernot is not Kubernetes:


    knot is not kubelet,

    kuber is not kube-apiserver,

    so k0s is not k8s.

The name __Kubernot__ is inspired by [many](#credits).

It should be shortened to __k0s__, mimicking the spelling of k8s, because Kubernot is made to resemble Kubernetes, while not being a distro of it.

The shortened form should be pronounced as "chaos" not "kay-zero-es" in order to distinguish from [Mirantis' k0s](https://www.mirantis.com/software/k0s/).

## Getting Started

`k0s` is the multicall binary of Kubernot. In addition to knot and kuber, it also bundles kubectl for convenience.

To install it, run:

```
$ go install k0s.io/cmd/k0s@latest
```

To start the server, run:

```
$ k0s kuber
```

then point kubectl to http://127.0.0.1:8000 and print the version:

```
$ k0s kubectl -s http://127.0.0.1:8000 version
```

## Credits

- [@alexellisuk](https://twitter.com/alexellisuk):

  [![image](https://user-images.githubusercontent.com/54848194/187806938-53ad18cd-b122-4690-9adb-8ea5cf194fe5.png)](https://twitter.com/alexellisuk/status/1366849550305140737)

- [@aevavoom](https://twitter.com/aevavoom):

  [![image](https://user-images.githubusercontent.com/54848194/187808142-748181f8-07f6-48c7-bb8e-786071e539c2.png)](https://twitter.com/aevavoom/status/1283146942738952193)

- [@mknz](https://twitter.com/mknz):

  [![image](https://user-images.githubusercontent.com/54848194/187809711-df63a8ef-9745-4992-9bd6-f9f168f39797.png)](https://twitter.com/mknz/status/1306608104201572357)
