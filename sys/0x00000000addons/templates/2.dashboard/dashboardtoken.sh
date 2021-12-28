#!/bin/bash

# Show ADMIN_SECRET

adminsecret=`/usr/sbin/kubectl --kubeconfig=/etc/kubernetes/ssl/kube-install.kubeconfig get secrets -n kube-system | grep dashboard-admin | awk '{print $1}'`


# Create DASHBOARD_LOGIN_TOKEN

token=`/usr/sbin/kubectl --kubeconfig=/etc/kubernetes/ssl/kube-install.kubeconfig describe secret -n kube-system $adminsecret | grep -E '^token' | awk '{print $2}'`


# Create dashboard_login_token files

echo $token >> /etc/kubernetes/dashboard_login_token.txt



