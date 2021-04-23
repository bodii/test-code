#!/bin/bash

while getopts "v:" opt;
do
	case "$opt" in
		v) vcode="$OPTARG" ;;
	esac
done

case "$vcode" in
	'18.04') Codename='bionic' ;;
	'20.04') Codename='focal' ;;
	'20.10') Codename='groovy' ;;
	'21.04') Codename='hirsute' ;;
	*)
	echo -e "[-v] version number parameters are emtpy! e.g.: 20.04";
	exit 1;;
esac

upsource="
deb http://mirrors.aliyun.com/ubuntu/ $Codename main multiverse restricted universe\n
deb http://mirrors.aliyun.com/ubuntu/ $Codename-backports main multiverse restricted universe\n
deb http://mirrors.aliyun.com/ubuntu/ $Codename-proposed main multiverse restricted universe\n
deb http://mirrors.aliyun.com/ubuntu/ $Codename-security main multiverse restricted universe\n
deb http://mirrors.aliyun.com/ubuntu/ $Codename-updates main multiverse restricted universe\n
deb-src http://mirrors.aliyun.com/ubuntu/ $Codename main multiverse restricted universe\n
deb-src http://mirrors.aliyun.com/ubuntu/ $Codename-backports main multiverse restricted universe\n
deb-src http://mirrors.aliyun.com/ubuntu/ $Codename-proposed main multiverse restricted universe\n
deb-src http://mirrors.aliyun.com/ubuntu/ $Codename-security main multiverse restricted universe\n
deb-src http://mirrors.aliyun.com/ubuntu/ $Codename-updates main multiverse restricted universe\n
"

echo -e $upsource > "sources_$Codename.list"
