#!/bin/bash

echo 
echo -n "This Program add line numbers for a text file.Specify
a text file path: "

# 读取用户输入的文件路径
read file

echo 
echo " Processing each line of file $file..."
echo 

count=0

# 从完整路径中获取文件名
filename=`basename $file`

# 每成功读取一行数据， while循环就会继续执行
while read line
do
	# 行号加1
	count=$((count+1))

	# 在原始每一行数据的前面添加行号
	echo $count:$line

	# 把文件中的数据作为while语句中read命令的标准输入
	# 同时把while循环的输出重定向到文件filename.lined中
done < $file > $filename.lned

echo "Output file is $filename.lined."
echo 

exit 0


