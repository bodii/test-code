#!/bin/bash

echo 
echo "重定向文件描述符练习"
echo "把exec.out文件重定向的文件描述符4上，如果不存在则创建"
echo "之后文件描述符4得到的所有内容都会被指向exec.out文件"
exec 4>exec2.out
echo 

echo "把文件符4重定向5"
echo "之后文件描述符5所得到的内容，都会被指向文件描述符4,最终指向exec.out"
exec 5>&4

echo "这是第一句话，将会写入文件描述符5,文件描述符5又指向文件描述符4,最终
指向文件exec2.out" 1>&5
echo "默认的文件描述符是1，所以可以省略" >&5

echo "这是第二句话，将会写入文件描述符4,最终指向文件exec2.out" >&4

echo "关闭文件描述符4"
exec 4>&-
echo "关闭文件描述符5"
exec 5>&-
echo 

exit 0

