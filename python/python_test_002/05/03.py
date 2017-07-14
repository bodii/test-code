#!/usr/bin/env python3
# -*- coding:utf-8 -*-

'''
lambda 

lambda只能是一个简单函数，并且不能包含其他语句，比如为变量创建一个名称。在
lambda内部，只能执行有限的操作，比如测试相等性、将两个数相乘或者以特定的方
式使用其他已经存在的函数。但是不能使用if...:elif...:else: 这样的构造，甚至
不能为变量创建新的名称！只能使用传递到lambda函数的参数。然而，可以通过使用
and和or操作，执行简单声明语句复杂一些的操作。但是仍然应该记住lambda只有非
常有限的用法。
lambda主要用在内置函数map和filter中。通过使用lambda，这些函数能够以紧凑的
方式执行一些大的操作，而不必使用循环。
'''

# Map: 短路循环
''' 匿名函数经常用到map函数中。map是一个特殊的函数，它用于需要对列表中的
每个元素执行一个指定的操作的情形。它在实现这种操作时不必编写一个循环。
如果向map中传递一个列表（或无组，map可以接受任何类型的序列作为参数）的列
表，则它期待得到列表。主列表中的每个序列应该有相同数目的元素
'''



# Now map gets to be run in the simple case
map_me = [ 'a', 'b', 'c', 'd', 'e', 'f', 'g' ]
result = map(lambda x: "The letter is %s\n" % x, map_me)
print(*result)
