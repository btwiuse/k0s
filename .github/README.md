# Kubernot - Not Kubernetes

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/k0s.io?tab=doc)
[![Go 1.18+](https://img.shields.io/github/go-mod/go-version/btwiuse/k0s)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/btwiuse/k0s?color=%23000&style=flat-round)](https://github.com/btwiuse/k0s/blob/master/LICENSE)
[![Visitors](https://visitor-badge.glitch.me/badge?page_id=btwiuse.k0s)](#)
[![DockerHub](https://img.shields.io/docker/pulls/btwiuse/k0s.svg)](https://hub.docker.com/r/btwiuse/k0s)
[![@kubernot](https://img.shields.io/twitter/url/https/twitter.com/kubernot.svg?style=social&label=Follow%20%40kubernot)](https://twitter.com/kubernot)

Kubernot is a modern SSH alternative with a Kubernetes-like interface, where

- `scp` becomes `kubectl cp`
- `ssh` becomes `kubectl exec`
- `ssh -L` becomes `kubectl port-forward`
- `~/.ssh/config` becomes `~/.kube/config`

Unlike SSH, Kubernot adopts the client-agent-server architecture from Kubernetes to allow accessing nodes behind a firewall.

||client|agent|server|
|:--:|:--:|:--:|:--:|
|SSH|`ssh`|-|`sshd`|
|Kubernot|`kubectl`|`k0s pode`|`k0s apiserver`|
|Kubernetes|`kubectl`|`kubelet`|`kube-apiserver`|

Kubernot emulates a Kubernetes cluster by implementing a minimal subset of the [Kubernetes API](https://kubernetes.io/docs/concepts/overview/kubernetes-api/).

The apiserver is modelled after [kube-apiserver](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/), therefore compatible with the familiar `kubectl` CLI.

Currently supported commands:

- [ ] kubectl cp
- [ ] kubectl exec
- [ ] kubectl port-forward
- [x] kubectl version

The name __Kubernot__ is inspired by many:

- [@alexellisuk](https://twitter.com/alexellisuk):

  [![image](https://user-images.githubusercontent.com/54848194/187806938-53ad18cd-b122-4690-9adb-8ea5cf194fe5.png)](https://twitter.com/alexellisuk/status/1366849550305140737)

- [@aevavoom](https://twitter.com/aevavoom):

  [![image](https://user-images.githubusercontent.com/54848194/187808142-748181f8-07f6-48c7-bb8e-786071e539c2.png)](https://twitter.com/aevavoom/status/1283146942738952193)

- [@mknz](https://twitter.com/mknz):

  [![image](https://user-images.githubusercontent.com/54848194/187809711-df63a8ef-9745-4992-9bd6-f9f168f39797.png)](https://twitter.com/mknz/status/1306608104201572357)
