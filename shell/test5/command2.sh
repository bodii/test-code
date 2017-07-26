#!/bin/bash

echo date

echo "`date +%Y-%m-%d`"

DAY=$(date +%A) # 星期

echo -n "Today is $DAY"

if [ "$DAY" = 'Saturday' -o "$DAY" = 'Sunday' ]
then
	echo "Have a good weekend!"
fi
echo

echo " Next year will be `expr \`date +%Y\` + 1`"

echo

echo "Previous year was $(expr $(date +%Y) - 1)"

echo
# 显示登录用户 
echo -n "Crrrently login users:"
echo $( who |cut -c1-4 |sort -u )
echo

exit 0
