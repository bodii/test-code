#!/bin/bash

# 
# 使用当前脚本运行的进程PID创建一个唯一的文件
TMPFILE=tmpfile.$$

# 定义回调函数
Init()
{
	echo
	echo "Receive SIGHUP,Reinitalize script..."
	echo "Deleting temprorav file.Creating new one..."
	rm -f $TMPFILE 2> /dev/null
	echo `date` > $TMPFILE
	echo "Done."
	echo -en "\t"
	TIME=15
}

# 定义回调函数
CleanUp()
{
	if [ -f "$TMPFILE" ]; then
		echo
		echo " Cleaning Up..."
		rm -f $TMPFILE 2> /dev/null
		echo "Done"
		echo
	fi

	# 从脚本退出
	exit 2
}

# 为信号2、3、15注册回调函数CleanUp
trap CleanUp 2 3 15
# 为信号1注册回调函数Init
trap Init 1


# 创建临时文件
echo 
echo "Creating temporary file $TMPFILE..."
echo `date` > $TMPFILE
echo

# 模拟一些工作
echo "Script is runing..."
echo -en "\t"


# 打印进度条
TIME=15
until [ "$TIME" -eq 0 ]
do

	echo -n "###"
	# 每打印一次###号就睡眠一秒
	sleep 0.4
	# 让TIME的值减一，最终等于零时循环结束
	let TIME-=1
done

# 执行到这里时，会清除临时文件
echo
echo "Cleaning Up termporary file $TMPFILE..."
rm -f $TMPFILE 2> /dev/null
echo

exit 0

