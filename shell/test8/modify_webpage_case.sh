#!/bin/bash

# sed  命令替换 不区分大小写
for file in "$@"
do
	if [ -f "$file" ]
	then
		echo "**processing $file**"
		sed -e 's/jerry/emma/gi' -e 's/2013/2014/gw 06.txt' "$file" > "$file.$$"

		if [ -f "$file.$$" ] && [ $? -eq 0 ] 
		then
			mv -f "$file.$$" "$file"
		fi
	fi
done

echo "All Done."

exit 0
