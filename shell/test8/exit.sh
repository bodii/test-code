#!/bin/bash

# 这个回调函数会在脚本退出以前执行
quit_handler()
{
	echo
	echo "In quit_headler():"
	echo "Script will exit."
	echo "Bye."
	echo
}

# 注册伪信号0的回调函数，使得在脚本退出以前可以做一些事情
# 使用EXIT和0效果是一样的
# trap quit_handler EXIT
trap quit_handler 0

echo "These lines prints before the 'trap'--"
echo "Even though the script see the 'trap' first."
echo

# 即使注释exit命令，伪信号0也会产生
exit 0
