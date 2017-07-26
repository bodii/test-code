#!/usr/bin/awk -f

# 在awk脚本中不要再使用单引号包围程序
BEGIN {
	printf("%-19s%-8s%-1s%-13s%-5s%-8s%-5s%s\n",
		   "FileName","Time","Date","Size",
		   "Group","User","Link","Right");
	printf("--------------------------------------------------\n");

}

# 如果第一列不是字符串"tatal"，则执行后面的命令
# 丢弃包含字符串“total”或“总容量”的行
$1 !~/(total)|(总用量)/{
	printf("%-19s%-8s%-1s%-13s%-5s%8s%8s%5s%s\n",
		  $9,$8,$6,$7"日",$5,$4,$3,$2,$1);
}


