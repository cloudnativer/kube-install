#!/bin/bash

cat $1/config/install.inventory  | grep $2 | grep $3 | cut -d "=" -f 3,4 | head -1


