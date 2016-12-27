#!/bin/bash

echo "Today is $(date +%Y-%m-%d)"

# 获得今天是这一年中的第几天
DAY=$(date +%j)

# 不能在DAY的前面添加$
echo "There is $(((365 - DAY) / 7 )) weeks before New Year."

# 在$DAY的两侧要有空格
echo "Have passed $(expr $DAY / 7) weeks in this year."

# let 命令的算术表达式不能包含任何空格， 并且不会用计算结果替换表达式
let weeks=(365 - $DAY)/7
echo "There is $weeks weeks before New Year."

exit 0
