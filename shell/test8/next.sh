#!/bin/bash

# 如果第一条规则成功地打印出用户名，就执行netx命令跳过第二第规则
awk -F: '
	$7 =="/bin/bash" {print $1;next;}
	$3 >= 1000 {print $1;}
	' /etc/passwd
