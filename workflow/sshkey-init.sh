#/bin/bash


#$1 is password   $2 is ip list   $3 is softdir   $4is currentdir   $5 is option


if [ -f $4"/workflow/"$5".inventory" ];then
  rm -rf $4"/workflow/"$5".inventory"
fi

if [ "$5" == "install" ];then
  mkdir -p $3
fi

if [ ! -f "/root/.ssh/id_rsa" ];then
  ssh-keygen -t rsa -P "" -f /root/.ssh/id_rsa
fi

for sship in $2;
do
  sshpass -p $1 ssh-copy-id -p 22 root@$sship >/dev/null 2>&1
done




