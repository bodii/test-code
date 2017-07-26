#!/bin/bash

# 保存sed 编辑命令的文件
SED_SCRIPT=webpage.modify

# 处理命令行的每一个文件
for file in "$@"
do
	if [ -f "$file" ]
	then
		echo "**processing $file**"
		sed -f $SED_SCRIPT "$file" > "$file.$$"

		if [ -f "$file.$$" ]
		then
			mv -f "$file.$$" "$file"
		fi
	fi
done

echo "All Done"

exit 0
