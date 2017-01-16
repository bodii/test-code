#!/bin/bash

# 定义捕捉ERR伪信号的函数
ERRTRAP(){
	echo "[LINE:$1] Error:Command or function exited with status $?"
}

fun_will_fail(){

	# 函数返回非零值表示错误
	return 1
}

fun_will_succed(){

	# 返回零表示函数执行成功
	return 0
}

# 命令或函数产生的每一个错误都将产生一个ERR伪信号
# 这里执行trap 命令是希望通过ERRTRAP捕捉这个ERR伪信号 
trap 'ERRTRAP $LINENO' ERR

# 模拟一个失败的命令
command_not_exist 2>/dev/null

# 模拟一个成功执行的函数
fun_will_succed

# 模拟一个失败了的函数
fun_will_fail

exit 0
