#!/bin/bash

if [ $# != 1 ]
then
	echo
	echo "Usage:"
	echo -e "\t `basename $0` [directory]"
	echo 
	exit 1
fi

ls -lR "$1" | awk '
	# 定义找出最大数量的值
	function max(arr){
		max_number = 0;
		for(i in arr){
			if (arr[i] > max_number){
				max_number = arr[i];
			}
		}
		return max_number
	}

	# 打印#数量函数
	function progress_length(curentlength, maxlength){
		return curentlength / maxlength * 50
	}

	BEGIN{
		printf "%-10.10s %8s\n", "User Name", "Numbers"
	}

	# 主程序
	# 获取每一行用户数量统计
	/^-/{
		result[$3]++
	}

	# 结果程序
	END {
		max_length = max(result);
		for(user in result){
			printf "%-10.10s [%8d]",user,result[user];
			for(i=1;i<progress_length(result[user],max_length);i++){
				printf "#";
			}
			printf "\n";
		}
	}
	' 
	exit 0
