#!/bin/bash

# 打印用户主目录字段串从第7个字符开始的7个字符长的子串‘
name=$(expr substr $HOME 7 7)
echo $name

exit 0
