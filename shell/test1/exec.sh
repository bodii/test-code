#!/bin/bash

# 打开一个新的文件描述符4,它被关联到文件exec.out
echo 
echo "Open file descriptor 4, which is associated with file exec.out"
exec 4>exec.out
echo

echo "Open another file descriptor 5, which is associated with file descriptor4."
# 打开另一个文件描述符5, 它被关联到文件描述符4
exec 5>&4

echo "Sending some date...."

# 重定向标准输出到文件描述符5
echo "Data stream from file descriptor 1 STDOUT," >&5
echo "readirected to file descriptor 5," >&5
echo "Then data will go to file descriptor 4, which is associated with file
exec.out." >&5
echo

# 任务完成以后关闭文件描述符4和5
echo "closing fd 4..."
echo 4>&-
echo "closing fd 5..."
echo 5>&-
echo 

exit 0
