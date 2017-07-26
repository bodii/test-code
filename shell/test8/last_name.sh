#!/bin/bash

awk '
	BEGIN {
		# 设置输入字符分隔符Input Field Separator 为换行符
		FS="\n";

		# 设置输入记录分隔符Input Record Separator为空字符串
		RS="";

		# 设置输出字段Output Field Separator为换行符
		OFS="\n";

	}

	# 使用正则表达式^N过滤所有无关数据
	# 只对以字母N开始的记录块进行处理
	/^N/ {
		# 输出以字母N开始的开发者的全名
		print $1;
	}

	# 第一个awk程序的处理结果会继续发送给第二个awk程序进行处理
	' CREDIT | ask '
		BEGIN {
			print "The Last Name of Developers:"
			print "----------------------------"
		}

		/^N/ {
			# 打印行号和最后一个字段
			print NR,$NF;
		}

		'
exit 0
