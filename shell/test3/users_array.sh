#!/bin/bash

#数组元素从0开始
index=0

# 从passwd文件中取出所有用户
# 指定分隔符为冒号，使用cut命令从passwd文件的每一行中提取出用户名
for i in `cut -f 1 -d : /etc/passwd`
do
	#把每一个用户名赋值给user数组的元素
	user[$index]=$i
	#每赋值一个用户以后，把索引值加一
	let index=$index+1
done

# 重复使用变量index作为行号
index=1


# 使用字符@分隔所有的数组元素
# 如果没有使用双引号引用，使用星号${user[*]}也可以
for name in "${user[@]}"
do
	echo "$index: $name"
	# 行号加一
	let index=$index+1
done

echo "-----------------------------------"
echo "Print all users in one line:"
echo 

#把所有的数组元素作为一个整体进行打印
echo "${user[*]}"

#把数组元素作为单个的个体进行打印
#echo ${user[@]}
#echo ${user[*]}
echo

echo "-----------------------------------"
# 对user数组重新赋值，前50个元素将会丢失
# 即使没有明确地覆盖它们

echo "Reassign the user array, the user names will be lost:"
user=([50]=a1,a2,a3)
echo 
echo ${user[*]}
echo 

exit 0

