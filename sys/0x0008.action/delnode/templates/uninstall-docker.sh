#!/bin/bash


for dockersoftware in {docker-ce-19,docker-ce-cli,containerd.io,container-selinux};
do
  rpm -e $(rpm -qa | grep $dockersoftware)
done

if [ -d "/etc/docker/" ];then
  rm -rf /etc/docker/
fi

if [ -d "/etc/containerd/" ];then
  rm -rf /etc/containerd/
fi

for dockerinstallfile in {/usr/bin/docker,/usr/bin/dockerd,/usr/bin/docker-init,/usr/bin/docker-proxy,/usr/bin/runc,/usr/bin/containerd,/usr/bin/containerd-shim,/lib/systemd/system/docker.service,/lib/systemd/system/docker.socket,/lib/systemd/system/containerd.service,/lib/systemd/system/container-getty@.service,/etc/systemd/system/docker.service,/etc/systemd/system/docker.socket,/etc/systemd/system/containerd.service,/etc/systemd/system/container-getty@.service,/lib/systemd/system/kubelet.service,/lib/systemd/system/kube-proxy.service,/etc/systemd/system/kubelet.service,/etc/systemd/system/kube-proxy.service};
do
  if [ -f $dockerinstallfile ];then
    rm -rf $dockerinstallfile
  fi
done


