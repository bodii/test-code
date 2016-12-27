#!/bin/bash

# 检查子进程中是否有父进程中变量LOCATION的拷贝
echo "---------------------------"
echo "The child is at '${LOCATION:-somewhere not defined}'."

# 如果子进程中拥有变量LOCATION的拷贝，就修改它，如果没有就定义它
LOCATION=CHINA

# 访问子进程中的变量LOCATION
echo
echo "The child is at '$LOCATION'."
echo "---------------------------"

exit 0

