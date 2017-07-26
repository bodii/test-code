#!/bin/bash

# 处理命令行的每一个文件，可以使用通配符
for file in "$@"
do
	echo "**processing $file**"
	
	# 保存结果到临时文件中
	sed 's/jerry@zulmma.com/emma@zulmma.com/g' "$file" > "$file.$$"

	if [ -f $file.$$ ]
	then
		# 覆盖原始文件
		mv -f "$file.$$" "$file"
	fi
done

echo "All Done."
exit 0

