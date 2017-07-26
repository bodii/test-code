#!/bin/bash

# 使用空字符串或冒号屏蔽信号2 SIGINT
# trap '' 2
# trap : 2
trap '' INT

echo
echo "Doing Critical Operation,Cannot be interrupted..."

# 睡眠30秒模拟某些操作
echo -en "\t"
COUNT=30
while [ $COUNT -gt 0 ]
do
	# 打印进度条
	echo -n "##"
	sleep 1
	let COUNT-=1
done

echo
echo "Done."

exit 0
