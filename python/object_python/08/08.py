#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' YAML文件的格式化 '''

import yaml

with open('some_destination.yaml', 'w', encoding='UTF-8') as target:
    yaml.dump( some_collection, target )

# 读取YAML文件时，会使用类似的技巧
with open('some_source.yaml', 'r', encoding='UTF-8') as source:
    objects = yaml.load( source )


"""
有一些格式化的方式，如下：

explicit_start          如果是true, 则在每个对象前写一个---标记

explicit_end            如果是true,则在每个对象前写一个...标记。当我们将一个YAML文档
的序列转储到一个文件中并且操作是串行的时候，可以使用它或者explicit_start

version                 指定一个整数对（x, y),在文件头输出%YAML x.y, 这应该是版本=
（1, 2）

tags                    指定一种映射，在文件头使用不同的标签缩写输出一个 YAML %TAG

canonical               如果为true，则每块数据包含一个标签，如果为false，则认为包含了很多标签

indent                  如果设定一个数字，就会改变块之间的缩进

width                   如果设定一个数字，当项太长至于显示为多行，缩进行，这个设置会改变行宽度
                        
allow_unicode           如果设为true，将支持完整的、没有包含转义符的Unicode编码。否则，在ASCII
自已外部的字符就会包含使用了转义符的字符

line_break              使用一种不同的行结束符，默认是换行符


！！ 以上这些选项中，explicit_end和allow_unicode可能是最常用的。
"""


