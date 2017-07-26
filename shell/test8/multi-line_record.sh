#!/bin/bash

awk '
	BEGIN {
		# 设置分隔符为换行符\n
		FS="\n";

		# 设置输入记录分隔符Input Record Separator 为空字符串
		RS="";

		# 设置输出字段分隔符Output Field Separator为换行符
		OFS="\n";

		# 设置输出记录分隔符Output Record Separator 为两个换行符
		ORS="\n\n";

		# 打印标题
		print "The developers of Linux Kernel:";
		print "(N)name (E)email";
		print " ------------------------ "
	}

	# 使用正则表达式^N过滤所有无关数据
	# 只对以字母N起始的记录块进行处理
	/^N/ {
		print $1, $2;
	}

' CREDITS

exit 0
