#!/usr/bin/env bash

for bin in {conntroll,agent,hub,client}; do
  docker build -t conntroll/${bin} -f Dockerfile.${bin} .
done

for distro in {alpine,ubuntu,debian,centos,fedora,gentoo,arch}; do
  docker build -t conntroll/agent:${distro} -f Dockerfile.agent.${distro} .
done
