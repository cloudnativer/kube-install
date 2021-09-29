バイナリ方式では、オフラインで利用可能な高い kubernetes クラスタをインストールするとともに、クベルネテツノードの追加、ノードの削除、ホストの廃棄、ホストの再構築、クラスタのアンマウントなどをサポートします。
<br>
（目的のホストにソフトウェアをインストールする必要はなく、純粋な裸マシンがあればオフラインで高利用可能な kubernetes  クラスタの配置を完了することができます。）
<br>

![kube-install](docs/images/kube-install-logo.jpg)

<br>

言語を切り替え: <a href="README0.7.md">English Documents</a> | <a href="README0.7-zh-hk.md">繁体中文文档</a> | <a href="README0.7-zh.md">简体中文文档</a> | <a href="README0.7-jp.md">日本語の文書</a>

<br>

# [1] 互換性

<br>
互換性説明:

<table>
<tr><td><b>kube-installバージョン</b></td><td><b>Kubernetesバージョン</b></td><td><b>オペレーティングシステム</b></td><td><b>関連文書</b></td></tr>
<tr><td> kube-install v0.7.* </td><td> kubernetes v1.23, v1.22, v1.20, v1.19, v1.18, v1.17 </td><td> CentOS 7 , RHEL 7 , CentOS 8 , RHEL 8 , SUSE Linux 15 , Ubuntu 20 </td><td><a href="README0.7.md">README0.7.md</a></td></tr>
<tr><td> kube-install v0.6.* </td><td> kubernetes v1.22, v1.21, v1.20, v1.19, v1.18, v1.17, v1.16, v1.15, v1.14 </td><td> CentOS 7 , RHEL 7 , CentOS 8 , RHEL 8 , SUSE Linux 15 </td><td><a href="README0.6.md">README0.6.md</a></td></tr>
<tr><td> kube-install v0.5.* </td><td> kubernetes v1.21, v1.20, v1.19, v1.18, v1.17, v1.16, v1.15, v1.14 </td><td> CentOS 7 , RHEL 7 </td><td><a href="README0.5.md">README0.5.md</a></td></tr>
<tr><td> kube-install v0.4.* </td><td> kubernetes v1.21, v1.20, v1.19, v1.18, v1.17, v1.16, v1.15, v1.14 </td><td> CentOS 7 , RHEL 7 </td><td><a href="README0.4.md">README0.4.md</a></td></tr>
<tr><td> kube-install v0.3.* </td><td> kubernetes v1.18, v1.17, v1.16, v1.15, v1.14 </td><td>CentOS 7</td><td><a href="README0.3.md">README0.3.md</a></td></tr>
<tr><td> kube-install v0.2.* </td><td> kubernetes v1.14 </td><td> CentOS 7 </td><td><a href="README0.2.md">README0.2.md</a></td></tr>
<tr><td> kube-install v0.1.* </td><td> kubernetes v1.14 </td><td> CentOS 7 </td><td><a href="README0.1.md">README0.1.md</a></td></tr>
</table>

<br>
注意：kube-installはCentOS 7、CentOS 8、SUSE 15、RHEL 7、RHEL 8のオペレーティングシステム環境に対応しています。<a href="docs/os-support.md">ここをクリックして、kube-installがサポートするオペレーティングシステムの発行版のリストを調べます<a>。
<br>
<br>
<br>

# [2] 速やかにKubernetesクラスタをインストールする

<br>

完璧の中で、期待してください ...

<br>
<br>


