
# Version 0.8.0 Feature Preview:

<br>

We will add the following features in 0.8.0 version:

* The web platform supports username and password login.
* Add the `-logs` flag to support viewing logs on the command line.
* Starting to support `kubernetes v1.24`.
* Since `kubernetes v1.24` has forcibly deleted `dockershim`, we will use `containerd` as runtime by default in `kubernetes v1.24`.
* In `kubernetes v1.18` to `kubernetes v1.23`, we use `docker` as the runtime by default, but `docker` is deprecated. 
* `kubernetes v1.17` is deprecated, it will not be integrated by default. If you want to use `kubernetes v1.17`, you can use the <a href="https://github.com/cloudnativer/kube-install/releases/tag/v0.7.4">kube-install v0.7.*</a>.
* Add dashboard `metrics-scraper` to solve the problem of `kube-dashboard` monitoring chart.
* Set the deployment of `kube-dashboard` as an option, and users can choose not to deploy it.
* Solve the problem of `secret` "harbor-secret" not found.
* Support non-standard `SSH` port for installation.
* Fix DOS security vulnerability of `go-yaml`.
* Upgrade `kube-dashboard` to v2.4.0.
* Upgrade `pause` image to v3.5.
* Upgrade `helm` to v3.7.2.
* ...


<br>
