#!/usr/bin/env bash
for bin in {conntroll,agent,hub,client}; do
  docker --config ~/.docker/ push conntroll/${bin}
done

for distro in {alpine,ubuntu,debian,centos,fedora,gentoo,arch}; do
  docker --config ~/.docker/ push conntroll/agent:${distro}
done
