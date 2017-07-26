#!/bin/bash

# 本脚本需要至少一个命令行参数
if [ "$#" -eq 0 ]
then
	echo "Usage:"
	echo -e "\t`basename $0` [FILE]..."
	exit 1
fi

# 处理每一个命令行参数
for FILE in "$@"
do
	if [ -f "$FILE" ] && [ -r "$FILE" ] && [ -s "$FILE" ]
	then
		echo "$FILE:"
		awk '
			# 统计所有的行数
			/.*/{ total_lines += 1; }
			# {total_lines +=1;}

			# 统计所有的空白行数
			/^[\s]*$/{ empty_lines += 1; next; }

			# 统计所有带注释的行
			/^[.]*#.*$/{ comment_lines += 1; }

			# 在主循环结束后输出结果
			END{
				printf "\tTotal Lines:%s\t\t",total_lines;
				printf "\tEmpty Lines:%s\t\t",empty_lines;
				printf "\tComment Lines:%s\t\t",comment_lines;
			}
		' "$FILE"
	else
		echo "$FILE can not be handled!"
	fi
done

echo "Done."

exit 0
