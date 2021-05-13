#/bin/bash


#$1 and $2 is path , $3 is ostype

mkdir -p $1

rpm -U /$2/pkg/sshops-$3/*.rpm --force --nodeps >/dev/null 2>&1

sed -i '/StrictHostKeyChecking/s/^#//; /StrictHostKeyChecking/s/ask/no/' /etc/ssh/ssh_config

systemctl daemon-reload

systemctl restart sshd


