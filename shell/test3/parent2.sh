#!/bin/bash


if [ "$1" = "--export" ]
then
	export LOCATION='parent'
elif [ "$1" = "--not-export" ]
then
	LOCATION='parent'
else
	echo
	echo -e "`basename $0` --export\t使用父进程的的环境变量"
	echo -e "`basename $0` --not-export\t 不使用父进程的环境变量"
	exit 0
fi

echo 
echo "父进程的变量LOCATION='$LOCATION'"


# 进入子进程
./child2.sh

echo "父进程的变量LOCATION at: "
echo "$LOCATION"
