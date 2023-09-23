#!/bin/bash


# $1 is path , $2 is ostype


# Install the sshpass tool
if [ "$2" = "ubuntu20" ];then
    dpkg --force-depends -i /$1/pkg/ubuntu20/sshops-ubuntu20/*.deb >/dev/null 2>&1
else
    rpm -U /$1/pkg/$2/sshops-$2/*.rpm --force --nodeps >/dev/null 2>&1
fi

# Set StrictHostKeyChecking
sed -i "/PermitRootLogin/d" /etc/ssh/sshd_config
sh -c "echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config"
sed -i '/StrictHostKeyChecking/s/^#//; /StrictHostKeyChecking/s/ask/no/' /etc/ssh/ssh_config
systemctl daemon-reload
systemctl restart sshd

