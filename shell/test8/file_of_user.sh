#!/bin/bash

# 本脚本在执行时需要指定一个参数目录，从而统计其中的文件
if [ $# != 1 ]
then
	echo 
	echo "Count how many files belong to which user."
	echo " Usage:"
	echo -e "\t `basename $0` [directory]"
	echo
	exit 1
fi

# 递归地显示指定目录下的所有文件信息
ls -lR "$1" | awk '
	# 定义函数max,用来找出数组result中的最大值
	function max(arr){
		max_number = 0;
		for(i in arr){
			if (arr[i] > max_number){
				max_number = arr[i];
			}
		}

		return max_number
	}

	# 定义函数process_length,用来计算打印多少个#字符
	function progress_length(currentlength, maxlength){
		return currentlength / maxlength * 50
	}

	BEGIN{
		printf "%-10.10s %8s\n", "User Name", "File Numbers";
	}

	# 主循环
	# 只处理以横杠起始的行，也就是只处理普通文件
	/^-/{
		# 每遇到一个普通文件，就把它的用户名作为索引值
		# 把相应数组对象的值加1,当主循环结束后就可以得到各个用户拥有的文件数
		result[$3]++
	}

	# 当主循环结束后，计算并打印出进度条
	END{
		max_length = max(result);

		# 使用for循环遍历数组result中的所有用户，并打印出它们拥有多少个文件
		for(user in result){
			printf "%-10.10s [%8d]:", user, result[user];
			for(i=1;i<progress_length(result[user],max_length);i++){
				printf "#";
			}
			printf "\n";
		}
	}
	'
	exit 0

