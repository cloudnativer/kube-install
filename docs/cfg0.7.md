# Use configuration file to install kubernetes cluster

First, you need to write a configuration file in the following format:


```
master = 192.168.56.81,192.168.56.82,192.168.56.83
node = 192.168.56.81,192.168.56.82,192.168.56.83
sshpwd = 123456789
ostype = centos7
softdir = /data/kube-install
```

Then you can execute `./kube-install -opt install -cfg your.cfg` to install kubernetes cluster.

<br>
<br>

# Use configuration file to add kubernetes node

First, you need to write a configuration file in the following format:


```
node = 192.168.56.85
sshpwd = 123456789
ostype = centos7
softdir = /data/kube-install
```

Then you can execute `./kube-install -opt addnode -cfg your.cfg` to add kubernetes node.


<br>
<br>

# Use configuration file to ***

to do ...




