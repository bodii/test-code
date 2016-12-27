#!/bin/bash

echo 
echo "Create two files in current directory:"
echo "  'separate.out' for standard output,"
echo "  'separate.err' for standard output."
echo
echo " Please check their contents."


# 同时列出存在的和不存在的文件， 并把结果重定向到不同的文件
# 如果没有指定文件描述符，默认使用文件描述符1 表示标准输出
ls /bin/bash /bin/ls /bin/dd /bin/this_file_ont_exist > separate.out 2>separate.err
echo 

exit 0
