#!/bin/bash

VERSION=2.0
# 声明变量VARSION 是一个只读变量
readonly VERSION

echo 
echo -e "Defined readonly variable:\n VERSION=$VERSION"
echo


# 尝试修改一个只读变量，会失败
echo "Try to modify readonly variable VERSION."
VERSION=3.0
echo 
echo -e "Readonly variable VERSION doesn't change,\n VERSION=$VERSION"

echo "------------------------------------"
# 定义一个只读变量
readonly PATCHLEVEL=1
echo 
echo -e "Defined another readonly variable:\n PATCHLEVEL=$PATCHLEVEL"
echo

# 尝试修改只读变量，会失败
echo "Try to modify readonlhy variable PATCHLEVEL."
PATCHLEVEL=3
echo 
echo -e "Readonly variable PATCHLEVEL doesn't change,\n PATCHLEVEL=$PATCHLEVEL"

echo 
# 尝试清空只读变量，会失败
echo "Try to unset readonly variable PATCHEVEL."
unset PATCHLEVEL

echo 
echo -e "Readonly variable PATCHLEVEL doesn't change,\n PATCHLEVEL=$PATCHLEVEL"
echo 

exit 0
