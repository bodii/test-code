#!/bin/bash

path=""
path=`pwd`/${path}
file=${path}file.txt

awk 'NR==10{print $0;}' $file