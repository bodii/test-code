#!/bin/bash

# 临时目录
TMPDIR=globbing

# 创建临时目录
mkdir -p $TMPDIR 

# 进入创建的临时目录
cd $TMPDIR

# 创建一些文件来显示
touch a.txt b.txt c.txt

count=1
while [ $count -lt 4 ]
do
	touch "tmpfile.$count"
	let count+=1
done

echo 

echo "OUTPUT from ls *:"
ls *

echo
echo "OUTPUT from ls [ab]*:"
ls [ab]*

echo
echo "OUTPUT from ls ?.txt"
ls ?.txt

echo 
echo "OUTPUT from ls [^ab]*"
ls [^ab]*

echo
echo "OUTPUT from ls [!ab]*"
ls [!ab]*

echo 
echo "OUTPUT from ls [a-c]*"
ls [a-c]*

echo
echo "OUTPUT from ls *[1-4]"
ls *[1-4]

# 返回之前的目录
cd - > /dev/null

# 清除临时目录
echo 
echo "Clean op tmp dir"
rm -rf $TMPDIR
exit 0

