#!/bin/bash

# 使用当前脚本运行的进程PID创建一个唯一的文件
TMPFILE=tmpfile.$$

# 创建临时文件
echo
echo "Creating temporary file $TMPFILE..."
echo `date` > $TMPFILE

echo
# 模拟一些工作
echo "Script is running..."
echo -n "      "

# 打印进度条
TIME=15
until [ "$TIME" -eq 0 ]
do
	echo -n "###"
	# 每打印一次#号就睡眠一秒
	sleep 1
	# 让TIME的值减一,最终等于0循环结束
	let TIME-=1
done

# 只有脚本执行到这里时，才会清除临时文件
echo
echo "Cleaning Up termporary file $TMPFILE..."
rm -f $TMPFILE 2> /dev/null
echo

exit 0
