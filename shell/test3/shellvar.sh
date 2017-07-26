#!/bin/bash

# 主目录
echo -e "HOME directory:\t\t$HOME"

# 当前shell
echo -e "Current Shell:\t\t$SHELL"

# 登录名
echo -e "User Name;\t\t$USER"

# 登录用户UID
echo -e "User ID:\t\t$UID"

# shell的版本
echo -e "Bash Version:\t\t$BASH_VERSION"

# 当前工作目录
echo -e "Current Directory:\t\t$PWD"

# 修改当前工作目录到主目录
echo 
echo "***change current directory to '$HOME'***"
cd
echo -e "Current directory:\t\t$PWD"

# 之前的工作目录
echo -e "Previous Working Directory:\t$OLDPWD"
echo

# 睡眠10秒
echo "***Will sleep 10 seconds...***"
sleep 10s

# 从脚本被执行到现在一共多长时间
echo -e "Seconds since invoked:\t\t$SECONDS seconds"

# 打印bash级别
echo -e "Bash Level:\t\t$SHLVL"

# 打印10个随机数
echo -en "10 Random Numbers:\t\t"

i=10
until [ $i -lt 1 ]
do
	echo -n "$RANDOM "
	let i=$i-1
done

echo
echo

# 修改内部字段分隔符IFS为横杠
echo "***Change \$IFS to '-'***"
IFS=-
echo "'2012-11-11' will be retreated as 3 strings:"
DATE=2012-11-11
printf "%s\n" $DATE

echo 
echo -e "Command Search In:\t\t$PATH"
echo

exit 0

