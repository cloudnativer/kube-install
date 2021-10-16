
# Run as a systemd service

<br>

You can run Kube-Install in systemd service mode. 

<br>

## Prepare systemd service file

Prepare `/etc/systemd/system/kube-install.service` file as follows, <a href="../kube-install.service">here is a sample file for reference</a>.

```
[Unit]
Description=kube-install One click fast installation of highly available kubernetes cluster.
Documentation=https://cloudnativer.github.io/
After=sshd.service
Requires=sshd.service
  
[Service]
Environment="USER=root"
ExecStart=/var/lib/kube-install/kube-install -daemon -listen 0.0.0.0:9080
User=root
PrivateTmp=true
LimitNOFILE=65536
TimeoutStartSec=5
RestartSec=10
Restart=always

[Install]
WantedBy=multi-user.target

```

<br>

Notice: Please fill in the actual full path of `kube-install` binary file after `ExecStart=` parameter.
Kube-install web service listens to `TCP 9080` by default. If you want to modify the listening address, you can set it by modifying the `kube-install -daemon -listen ip:port` parameter in the `/etc/systemd/system/kube-install.service` file.

<br>

## Start the service

Start the service using the `systemctl start kube-install` command:

```
# systemctl start kube-install.service
#
# systemctl status kube-install.service
  ● kube-install.service - kube-install One click fast installation of highly available kubernetes cluster.
     Loaded: loaded (/etc/systemd/system/kube-install.service; disabled; vendor preset: disabled)
     Active: active (running) since Fri 2021-08-20 14:30:55 CST; 21min ago
       Docs: https://cloudnativer.github.io/
   Main PID: 2768 (kube-install)
     CGroup: /system.slice/kube-install.service
             └─2768 /go/src/kube-install/kube-install -daemon
   ...

```
<br>

## Set the service startup

To set the service startup, you can execute the following commands:

```
# systemctl enable kube-install.service
```

<br>
<br>
<br>
