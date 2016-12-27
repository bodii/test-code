#!/bin/bash

# 设置行号
linenumber=1

#如果读到的是一个空行，则退出循环
while read oneline
do
	echo -en "$linenumber : $oneline\n"

	#行号加1
	linenumber=`expr $linenumber + 1`

done < /etc/passwd

exit 0
