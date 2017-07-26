#!/bin/bash

echo 
echo "Create two files in current directory:"
echo "  'separate.out' for standard output, "
echo "  'separate.err' fro standard error output."
echo 
echo "Please check their contents."

#列出存在和不存在的文件，并把标准输出和标准错误输出重定向到同一个文件中
ls /bin/bash /bin/ls /bin/dd /bin/this_file_not_exist &> same1.out

#重定向标准错误输出到标准输出
ls /bin/bash /bin/ls /bin/dd /bin/this_file_not_exist > same2.out 2>&1

#重定向标准输出到标准错误输出
ls /bin/bash /bin/ls /bin/dd /bin/this_file_not_exist 2> same3.out 1>&2

echo
