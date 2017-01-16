#!/bin/bash

ls -l | awk '
# 过滤其他特殊文件和目录， 只对普通文件的行数据进行操作
($1 ~/^-/) && ($6=="12月"){
	printf "Filename:%-29s Size:%s\n",$9,$5;
	sum+=$5;
}

# 在主循环结束以后输出结果
END {
	print "-----------------------";
	print "The sum:",sum;
}'

exit 0
