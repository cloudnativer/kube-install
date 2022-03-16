
# kube-install cluster architecture description

<br>

The `./data/` directory is used to store all status information of `kube-install`, including configuration information of all kubernetes clusters. You can share the `./data/` directory through file storage, so as to realize active and standby or load balancing cluster.

<br>

## Active/standby architecture

<br>

You can use software such as `Keepalived` or `Heartbeat` to detect and switch between `active and standby`.

![architecture](images/architecture1.jpg)

<br>
<br>

## Load balancing architecture

<br>

You can use `LVS`, `Nginx` or `Haproxy` software to achieve `load balancing` and cluster switching.

![architecture](images/architecture2.jpg)

<br>
<br>
<br>
