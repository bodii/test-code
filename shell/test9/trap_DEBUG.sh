#!/bin/bash

# 每执行一个命令，shell都将发送一个DEBUG伪信号
trap 'echo "before execute line:$LINENO,a=$a,b=$b,c=$c"' DEBUG

# 赋值语句产生DEBUG伪信号
a=1

# 命令[产生DEBUG伪信号
if [ "$a" -eq 1 ]
then
	# 赋值语句产生DEBUG伪信号
	b=2
else
	b=1
fi

# 赋值产生DEBUG伪信号
c=3

echo "end"
