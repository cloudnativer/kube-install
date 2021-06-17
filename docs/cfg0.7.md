
# Operating with configuration files

<br>
You can perform various operations through pre written configuration files.
<br>

## Use configuration file to install kubernetes cluster

The cluster is installed through the pre written configuration file. 
<br>
First, you need to write a configuration file in the following format. For example, create a configuration file named `your.cfg`.

```
master = 192.168.1.11,192.168.1.12,192.168.1.13
node = 192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14
sshpwd = 123456789
ostype = centos7
softdir = /data/k8s
```

Then you can execute `./kube-install -opt install -cfg your.cfg` to install kubernetes cluster.

<br>
<br>

## Use configuration file to add kubernetes node

The kubernetes node is added through the pre written configuration file.
<br>
First, you need to write a configuration file in the following format. For example, create a configuration file named `your.cfg`.

```
node = 192.168.1.15,192.168.1.16,192.168.1.17
sshpwd = 123456789
ostype = centos7
softdir = /data/k8s
```

Then you can execute `kube-install -opt addnode -cfg your.cfg` to add kubernetes node.


<br>
<br>

## Use configuration file delete kubernetes node

The kubernetes node is deleted through the pre written configuration file.
<br>
First, you need to write a configuration file in the following format. For example, create a configuration file named `your.cfg`.

```
node = 192.168.1.14,192.168.1.16
sshpwd = 123456789
ostype = centos7
softdir = /data/k8s
```

Then you can execute `kube-install -opt delnode -cfg your.cfg` to del kubernetes node.

<br>
<br>

## Use configuration file delete kubernetes master

The kubernetes master is deleted through the pre written configuration file.
<br>
First, you need to write a configuration file in the following format. For example, create a configuration file named `your.cfg`.

```
master = 192.168.1.12
sshpwd = 123456789
ostype = centos7
softdir = /data/k8s
```

Then you can execute `kube-install -opt delmaster -cfg your.cfg` to del kubernetes master.

<br>
<br>

## Use configuration file rebuild kubernetes master

The kubernetes master is rebuilt through the pre written configuration file.
<br>
First, you need to write a configuration file in the following format. For example, create a configuration file named `your.cfg`.

```
master = 192.168.1.12
sshpwd = 123456789
ostype = centos7
softdir = /data/k8s
```

Then you can execute `kube-install -opt rebuildmaster -cfg your.cfg` to del kubernetes master.

<br>
<br>

## Use configuration file to uninstall kubernetes cluster

The cluster is uninstalled through the pre written configuration file.
<br>
First, you need to write a configuration file in the following format. For example, create a configuration file named `your.cfg`.

```
master = 192.168.1.11,192.168.1.12,192.168.1.13
node = 192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14,192.168.1.15,192.168.1.16,192.168.1.17
sshpwd = 123456789
ostype = centos7
softdir = /data/k8s
```

Then you can execute `kube-install -opt uninstall -cfg your.cfg` to install kubernetes cluster.

<br>
<br>

## Operating with flag parameters

You can also specify flag parameters to carry out various operation and maintenance operations. Click here to <a href="operation0.6.md">see more operation documents</a>.

<br>
<br>
<br>
