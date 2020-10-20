#!/bin/bash

path=""
path=`pwd`/${path}
file=${path}file.txt

sed -n '10p' $file