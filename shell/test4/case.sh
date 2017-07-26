#!/bin/bash

# 得到当前的操作系统平台
PLATFORM=`uname -s`

# 为各种操作系统定义相应的空变量
FREEBSD=
SUNOS=
MAC=
AIX=
MINIX=
LINUX=

echo 
echo "Identifying the platform on which this installer is running on..."
echo 

# case 依次进行平台测试
case "$PLATFORM" in
	"FreeBSD")
		echo "This is FreeBSD system."
		FREEBSD=1
		;;
	"SunOS")
		echo "This is SOlaris system."
		SUNOS=1
		;;
	"Darwin")
		echo "This is Mac OSX system."
		MAC=1
		;;
	"AIX")
		echo "This is AIX system."
		AIX=1
		;;
	"MINIX")
		echo "This is MINIX systme."
		MINIX=1
		;;
	"Linux")
		echo "This is Linux system."
		LINUX=1
		;;
	*)
		echo "Failed to identify this Operating System.Abort!"
		exit 1
		;;
esac

echo "Starting to install the software..."
echo 

exit 0


