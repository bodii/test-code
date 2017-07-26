#!/bin/bash

# 定义一个叫做usage的函数，用来显示脚本的使用方法
usage()
{
	# 在终端输出脚本的使用方法，需要两个参数，分别为operator 和 file
	echo "Usage: `basename $0` operator file."

	# 由于脚本执行失败，所以以错误代码1退出
	exit 1
}

# 检查执行本脚本时所指定的命令行参数个数
if [ $# -ne 2 ]
then
	echo "Error.Not specify arguments correctly6."

	# 由于命令行指定的参数个数不符合脚本运行的要求，所以调用函数usage
	usage
else
	# 此处只是模拟脚本成功执行所做的操作
	echo "My shell scripting is running."
	exit 0
fi
