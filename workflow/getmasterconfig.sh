#!/bin/bash

cat $1/workflow/install.inventory  | grep $2 | grep $3 | cut -d "=" -f 3


