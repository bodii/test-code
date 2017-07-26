#!/bin/bash

awk '
	BEGIN{
		# 在文循环开始以前指定分隔符为冒号
		FS=":";
	}

	# 通过表达式匹配所有的偶数行
	NR%2 == 0 {
		print NR, $0;
	}

	' /etc/passwd

exit 0
