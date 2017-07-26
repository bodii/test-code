#!/bin/bash

# 首先在脚本程序中包含进库文件library.lib,从而可以调用其中定义的函数
#source ./library.lib 与下面作用相同
. ./library.lib

# 调用函数库中的函数test_platfrom
test_platform

# 输出函数test_platform中定义的全局变量PLATFORM
echo
echo "Our running platform is $PLATFORM"

exit 0
