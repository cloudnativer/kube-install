#!/bin/bash

sed -i 's/auto/512M/g' /etc/default/grub
sed -i 's/crashkernel=512M rhgb quiet/crashkernel=512M rhgb quiet intel_idle.max_cstate=0 processor.max_cstate=1 intel_pstate=disable idle=poll nopti spectre_v2=off/g' /etc/default/grub
grub2-set-default $[`awk -F\' '$1=="menuentry " {print $2}' /etc/grub2.cfg | sed -n '/4.19.12/='`-1] 
grub2-mkconfig -o /boot/grub2/grub.cfg



