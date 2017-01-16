#!/bin/bash

# 把ls -l 的输出数据重定向给awk程序进行处理
ls -l | awk '

# 打印标题
BEGIN {
	printf("%-12s%-8s%-13s%-5s%-8s%-8s%-5s%s\n",
	"FileName","Time","Date","Size","Group","User","Link","Right");
	printf("-----------------------------------------------------------\n");
}

# 如果第一列不是字符串"total",则执行后面的命令
# 丢弃包含字符串”total"的行
$1 !~/total|总用量/{
	printf("%-12s%-8s%-13s%-5s%-8s%-8s%-5s%s\n"),$8,$7,$6,$5,$4,$3,$2,$1;
}'

exit 0
