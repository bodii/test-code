#!/bin/bash


echo 
echo "Backup Files"
echo "----------------"
# 根据本脚本所在的位置，有可能需要修改相对目录../
# 遍历上一级的所有子目录，备份其中的特殊文件
for file in ./*
do
	if [ -f $file ]
	then
		for extension in sh c txt h
		do
			# 备份文件
			echo -n "Backup \"`basename $file`\" to "
			echo "${PWD}/backup/`basename $file`  "
			 cp *.${extension} backup/ 2>/dev/null
			echo "Done."
		done
	else
		echo 
		echo "> it not file"
		continue
	fi
done

echo 
echo "Bye!"

exit 0
