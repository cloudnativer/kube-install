#/bin/bash

# example: ./sys/0x00000000000ssh/sshkey-init.sh "192.168.122.11 192.168.122.12 192.168.122.13 192.168.122.14 192.168.122.15" "123456789"
# $1 is ip list
# $2 is user
# $3 is password  


if [ ! -f "/root/.ssh/id_rsa" ];then
  ssh-keygen -t rsa -P "" -f /root/.ssh/id_rsa
fi

if [ "$3" == "" ];then
    ssh-copy-id -p 22 $2@$1 >/dev/null 2>&1
else
    for sship in $1;
    do
        sshpass -p $3 ssh-copy-id -p 22 $2@$sship >/dev/null 2>&1
    done
fi



