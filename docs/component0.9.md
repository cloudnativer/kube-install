
# Detailed component and version list

Installation component list information for kube-install v0.9：


| component                 | version                            | type                                                  | description                                              |
| ------------------------- | ---------------------------------- | ----------------------------------------------------- | -------------------------------------------------------- |
| kube-apiserver            | v1.24, v1.25, v1.26,  v1.27, v1.28 | k8s-master                                            | Different versions can be  selected for installation.    |
| kube-controller-manager   | k8s-master                         | Different versions can be  selected for installation. |                                                          |
| kube-scheduler            | k8s-master                         | Different versions can be  selected for installation. |                                                          |
| etcd                      | v3.5.13                            | k8s-master                                            |                                                          |
| kubelet                   | v1.24, v1.25, v1.26,  v1.27, v1.28 | k8s-node                                              | Different versions can be  selected for installation.    |
| kube-proxy                | k8s-node                           | Different versions can be  selected for installation. |                                                          |
| pause-amd64               | v3.5                               | k8s-node                                              |                                                          |
| calico-node               | v3.19.3                            | CNI-plugin                                            | Different CNI-plugins can be  selected for installation. |
| calico-pod2daemon-flexvol | v3.19.3                            | CNI-plugin                                            | Different CNI-plugins can be  selected for installation. |
| calico-cni                | v3.19.3                            | CNI-plugin                                            | Different CNI-plugins can be  selected for installation. |
| calico-kube-controllers   | v3.19.3                            | CNI-plugin                                            | Different CNI-plugins can be  selected for installation. |
| flannel-cni-plugin        | v1.0.0                             | CNI-plugin                                            | Different CNI-plugins can be  selected for installation. |
| flannel                   | v0.15.1                            | CNI-plugin                                            | Different CNI-plugins can be  selected for installation. |
| kube-router               | v1.3.2                             | CNI-plugin                                            | Different CNI-plugins can be  selected for installation. |
| weave-kube                | 2.8.1                              | CNI-plugin                                            | Different CNI-plugins can be  selected for installation. |
| weave-npc                 | 2.8.1                              | CNI-plugin                                            | Different CNI-plugins can be  selected for installation. |
| cilium                    | v1.9.0                             | CNI-plugin                                            | Different CNI-plugins can be  selected for installation. |
| cilium-operator-generic   | v1.9.0                             | CNI-plugin                                            | Different CNI-plugins can be  selected for installation. |
| kubernetes-dashboard      | v2.7.0                             | addons                                                |                                                          |
| metrics-server-amd64      | v0.5.0                             | addons                                                |                                                          |
| metrics-scraper           | v1.0.8                             | addons                                                |                                                          |
| heapster-amd64            | v1.5.4                             | addons                                                |                                                          |
| traefik                   | v2.0.7                             | addons                                                |                                                          |
| alpine                    | v3.6                               | addons                                                |                                                          |
| coredns                   | 1.3.1                              | addons                                                |                                                          |
| registry                  | v2.7.1                             | addons                                                |                                                          |






