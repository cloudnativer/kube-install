
# Use the web platform to install

<br>You can also install the Kubernetes cluster on the Kube-Install web platform. 

<br>

## Run the web management service

First, you need to execute the 'kube-install -exec init' command to initialize the system environment (you can skip if you have already initialized earlier), and then execute the 'systemctl start kube-install' command to run the web management platform service of kube-install.

```
#
# ./kube-install -daemon
```
Now, you can open it with a web browser `http://kube-install_IP:9080`, view kube-nstall web platform.
<br>
Notice: kube-install web platform service listens to `TCP 9080` by default. If you want to modify the listening address, you can set it by modifying the `kube-install -daemon -listen IP:port` parameter in the `/etc/systemd/system/kube-install.service` file.

## Open the SSH password free channel to the target host

Before starting the installation, please open the SSH password free channel from localhost to the target host.You can use the `kube-install -exec sshcontrol` command to SSH through, or click the `Open SSH Channel of Host` button in the upper right corner to SSH through.
<br>
use the `kube-install -exec sshcontrol` command to SSH through.

```
kube-install -exec sshcontrol -sship "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -sshpass "cloudnativer"
```

Or click the `Open SSH Channel of Host` button in the upper right corner to SSH through.

![kube-dashboard](images/webssh001.jpg)

Here is the process of SSH connection, <a href="webssh0.7.md">click here to view more details</a> !<br>

<br>
<br>

## Fill in the installation parameters in the form

Then click the `Install Kubernetes` button in the upper right corner to SSH through.<br>

![kube-dashboard](images/webinstall001.jpg)

Fill in relevant installation parameters in the pop-up form:<br>

![kube-dashboard](images/webinstall003.png)

Select the version of kubernetes you want to install and the type of operating system according to your actual environment.<br>

The default is to start the installation immediately. You can also set an installation time for scheduled installation.

<br>
<br>

## Start the installation of kubernetes

Click the `Submit` button to start the automatic installation.<br>

![kube-dashboard](images/webinstall004.jpg)

Using the schedule widget, you can also view all installation task plan calendars, <a href="schedule0.7.md">click here to view more details</a> !<br>

![kube-dashboard](images/webinstall002.jpg)

You can also view the installation process log by clicking the `Install Log` button.<br>

Wait about 15 minutes and the installation will be completed automatically in the background. 

<br>
<br>
<br>
