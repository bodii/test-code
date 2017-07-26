#!/bin/bash

# 创建临时目录

mp3_dir='./tmp/songs'
if [ ! -d $mp3_dir ]
then
	mkdir -p $mp3_dir
fi

if [ "$?" -eq "0" ]
then
	# 在主目录下找到所有的mp3文件
	# 然后把找到的这些文件都移到 /tmp/songs 文件夹
	find ./ -iname "*.mp3" -exec mv -f '{}' $mp3_dir/ \; >/dev/null 2>&1 
fi

echo 
echo "All mp3 files have been moved to directory '$mp3_dir'."
echo

exit 0
