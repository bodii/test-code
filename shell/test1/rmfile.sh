#!/bin/bash

# 使用命令替换$(command)的方法使得命令command的输出作为命令行的参数
# 删除当前目录下所有文件名以.rm结尾的文件

rm -i $(find . -name '*.rm1')

# 另一种命令替换的方法
rm -i `find . -name '*.rm2'`

exit 0
