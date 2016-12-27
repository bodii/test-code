#!/bin/bash

# 使用变量替换定义若干个变量
TIME=`date`
NAME=`uname -n`
KERNEL=`uname -s`
VERSION=`uname -r`
ARCH=`uname -m`
OS=`uname -o`

echo 
echo "  Inof about your computer"
echo " =================================="
echo " Current Time:             $TIME"
echo " Host Name:                $NAME"
echo " Operating System:         $OS"
echo " Computer ARCH:            $ARCH"
echo " Kernel Version:           $KERNEL $VERSION"
echo

exit 0
