#/bin/bash


#$1是密码   $2是所有ip列表   $3是软件目录 $4是当前目录 $5是执行的操作选项


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




