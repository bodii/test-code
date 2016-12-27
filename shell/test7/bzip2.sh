#!/bin/bash

echo
echo " Creating a new data file data.file..."

# 在当前目录下创建一个数据文件
dd bs=1024 count=1024 if=/dev/zero of=data.file 2>/dev/null

if [ $? -eq 0 ]
then

	# list the size of the file
	echo
	echo "This is our data file:"
	echo "  `ls -l data.file`"
fi

echo
echo "Now compressing the data.file..."
bzip2 -v data.file

if [ $? -eq 0 ]
then
	echo
	echo "Done"
	echo

	# 查看压缩后的文件大小
	echo "Our data file has been renamed to 'data.file.bz2', and has smaller size:"
	echo "   `file data.file.bz2`"
	echo "   `ls -l data.file.bz2`"
fi

exit 0
