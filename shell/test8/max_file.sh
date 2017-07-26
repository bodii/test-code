#!/bin/bash

echo

read -p "Find the biggest file in which dirctory(or quit):"

until [ "$REPLY" = "quit" ] || [ "$REPLY" = "Quit" ] || [ "$REPLY" = "q" ]
do
	# 可处理的目录要具有可读和可执行权限
	if [ -d "$REPLY" ] && [ -r "$REPLY" ] && [ -x "$REPLY" ]
	then
		# 通过管道让awk程序处理ls命令的输出
		ls -lR "$REPLY" | awk '
			$1 !~/总用量/{
				# 如果找到一个更大的文件，就记录下它的大小和文件名
				if ( $5 > max_size)
				{
					max_size = $5;
					max_file = $9;
				}
			}

			# 主循环结束以后，打印出结果
			END {
				printf("In directory %s,\nFile %s is the biggest,has %d bytes.\n",dir,max_file,max_size);
			}

			# 把要操作的目录传进awk脚本
			' dir=$REPLY

		else
			echo "$REPLY can not be accessed!"
		fi
		echo
		read -p "Find the biggest file in which directory(or quit):"
	done

	echo "Bye."

	exit 0
