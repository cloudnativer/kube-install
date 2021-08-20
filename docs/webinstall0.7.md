
# Use the web platform to install

<br>You can also install the Kubernetes cluster through the Kube-Install web platform. 

<br>

## Run the web management service

First run the web management service with the `kube-install -daemon` command, and then open `http://your_Kube-Install_host_IP:9080` with a browser.
```
#
# ./kube-install -daemon
```
Notice: The web service listens to `TCP 9080` port by default. You can also use the `-listen` parameter to modify the port number that the web service listens to.

## Open the SSH password free channel to the target host

Before starting the installation, please open the SSH password free channel from localhost to the target host.You can use the `kube-install -exec sshcontrol` command to SSH through, or click the `Open SSH Channel of Host` button in the upper right corner to SSH through.
<br>
use the `kube-install -exec sshcontrol` command to SSH through.

```
kube-install -exec sshcontrol -sship "192.168.1.11,192.168.1.12,192.168.1.13,192.168.1.14" -sshpass "cloudnativer"
```

Or click the `Open SSH Channel of Host` button in the upper right corner to SSH through.

![kube-dashboard](images/webinstall001.jpg)

Here is the process of SSH connection, <a href="webssh0.7.md">click here to view more details</a> !<br>

<br>
<br>

## Fill in the installation parameters in the form

The kubernetes node is added through the pre written configuration file.
<br>
Fill in relevant installation parameters in the pop-up form:

![kube-dashboard](images/webinstall003.png)

Select the version of kubernetes you want to install and the type of operating system according to your actual environment.

<br>
<br>

## Start the installation of kubernetes

The default is to start the installation immediately. You can also set an installation time for scheduled installation.

![kube-dashboard](images/webinstall004.jpg)

Click the `Submit` button to start the automatic installation.  Through the schedule widget, you can also view all installation task plan calendars, <a href="schedule0.7.md">click here to view more details</a> !<br>

![kube-dashboard](images/webinstall002.jpg)

Wait about 15 minutes and the installation will be completed automatically in the background. You can also view the installation process log by clicking the `Install Log` button.

<br>
<br>
<br>
