#!/bin/bash

if [ -f ./data.file.bz2 ]
then
	echo 
	echo "This is our compressed data file:"
	# 查看压缩文件的大小
	echo " `ls -l data.file.bz2`"
	echo

	# 解压缩文件data.file.bz2,可以得到文件data.file
	echo
	echo "Now uncompress the file 'data.file.bz2' back to 'data.file'..."
	bunzip2 data.file.bz2

	echo
	echo "After uncompressing,file data.file.bz2 no longer exists by default."
	echo
	
	echo "This is our original data file:"
	# 检查文件data.file的类型和大小
	echo " `file data.file`"
	echo " `ls -l data.file`"
	echo "The size fo 'data.file equals to the original file size."
else
	echo
	echo " There is no compressed filename 'data.file.bz2' in current directory."
	echo "Please run bzip2.sh first,then ren this script again to uncompress 'data.file.bz2' file!"
fi

echo

exit 0
