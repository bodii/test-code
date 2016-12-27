#!/bin/bash

# 显示当前日期
DAY=$(date +%A)
echo "Date:`date +%Y-%m-%d`"
echo -n "Today is $DAY."

if [ "$DAY" = "Saturday" -o "$DAY" = "Sunday" ]
then
	echo "Have a good weekend!"
fi
echo

# 显示明年和去年
echo " Next year will be `expr \`date +%Y\` + 1`"
echo 
echo " Previous year was $(expr $(date +%Y) - 1)"

echo 
# 显示登录用户
echo -n "Currently login users:"
echo $(who | cut -c1-8 | sort -u)
echo 

exit 0
