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

# 依次进行平台测试
if [ "$PLATFORM" = "FreeBSD" ]
then
	echo "This is FreeBSD system."
	# 如果测试为真，给相应的操作系统变量赋值1
	FREEBSD=1
# 如果之前的if条件不被满足， 继续进行elif语句中的测试
# 后面以此类推
elif [ "$PLATFORM" = "SunOS" ]
then
	echo "This is Solaris system."
	SUNOS=1
elif [ "$PLATFORM" = "Darwin" ]
then
	echo "This is Mac OSX system."
	MAC=1
elif [ "$PLATFORM" = "AIX" ]
then
	echo "This is AIX system."
	AIX=1
elif [ "$PLATFORM" = "MINIX" ]
then
	echo "This is MINIX system."
	MINIX=1
elif [ "$PLATFORM" = "Linux" ]
then
	echo "This is Linux system."
	LINUX=1
else
	echo "Failed to identify this Operating System, Abort!"
	exit 1
fi

echo "Starting to install the softWare..."
echo

# 后面的安装脚本可以检查各个操作系统变量的值是否为1
# 以此来执行各种不同的操作

echo
