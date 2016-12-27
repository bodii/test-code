#!/bin/bash

echo
echo "This Program can help you clean up all *.txt files in a directory!"
echo

echo "Please enter a source code directory which contains *.txt files:"
echo -n "Type 'quit' to exit:"

# 读取用户输入的目录，在该目录下查找×.o文件
read SOURCE_DIR

while [ "$SOURCE_DIR" != "quit" ]
do
	if [ -d "$SOURCE_DIR" ]
	then
		echo 
		echo "Finding all *.txt files in '$SOURCE_DIR'..."
		echo
		echo "Directory '$SOURCE_DIR' contains the following *.txt files:"
		echo

		# 打印出所有子目录下的*.txt文件
		find "$SOURCE_DIR" -name \*.txt -print

		# 删除所有这些*.txt文件，并把这些文件记录在files.list文件中
		find "$SOURCE_DIR" -name \*.txt -print -exec rm '{}'\; > files.list

		echo 
		echo " All *.txt files have been removed.Check Removed files list in 'files.list'"
		echo

		# 可以使用单引号或双引号引用星号×及分号;,防止 shell处理这它们
		# find "$SOURCE_DIR" -name "*.txt" -print -exec rm '{}'';' > files.list

		echo "Clean up another souce code directory which contains *.txt files?"
		echo -n "Enter another directory or type 'quit' to exit:"

		read SOURCE_DIR
	else
		echo
		echo "'$SOURCE_DIR' is not a directory, input directory again!"
		echo -n " Or type 'quit' to exit:"
		read SOURCE_DIR
	fi
done

echo
echo " Bye!"
echo

exit 0
