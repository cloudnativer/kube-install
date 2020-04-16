#/bin/bash


#$1是path

mkdir -p $1

rpm -U /$2/pkg/sshops/*.rpm --force >/dev/null 2>&1

sed -i '/StrictHostKeyChecking/s/^#//; /StrictHostKeyChecking/s/ask/no/' /etc/ssh/ssh_config

systemctl daemon-reload

systemctl restart sshd


