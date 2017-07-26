#!/bin/bash

# 使用空字符串或冒号屏蔽信号2 SIGINT
# trap '' 2
# trap : 2
trap '' INT

echo
echo "Doing Critical Operation,Cannot be interrupted..."

# 睡眠30秒模拟某些操作
echo -en "\t"
COUNT=20
while [ $COUNT -gt 0 ]
do
	# 打印进度条
	echo -n "##"
	sleep 1
	let COUNT-=1
done

echo
echo "Critical OPeration if finished, can receive SIGINT now."

# 重新设置信号SIGINT的回调函数为默认行为
trap INT
echo "Press CTRL+C!"
echo -en "\t"
COUNT=20
while [ $COUNT -gt 0 ]
do
	echo -n "##"
	sleep 1
	let COUNT-=1
done

echo
echo "Done."

exit 0

