#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第七章：函数
Desc: 使用 def 语句定义函数是所有程序的基础。本章的目标是讲解一些更加高级和不
常见的函数定义与使用模式。涉及到的内容包括默认参数、任意数量参数、强制关键
字参数、注解和闭包。另外，一些高级的控制流和利用回调函数传递数据的技术在这
里也会讲解到。

Title;       给函数参数增加元信息
Issue: 你写好了一个函数，然后想为这个函数的参数增加一些额外的信息，这样的话其他
使用者就能清楚的知道这个函数应该怎么使用。
Answer: 使用函数参数注解是一个很好的办法，它能提示程序员应该怎样正确使用这个函
数。例如，下面有一个被注解了的函数
"""
def add(x:int, y:int) -> int:
    return x + y

"""
python 解释器不会对这些注解添加任何的语义。它们不会被类型检查，运行时跟
没有加注解之前的效果也没有任何差距。然而，对于那些阅读源码的人来讲就很有帮
助啦。第三方工具和框架可能会对这些注解添加语义。同时它们也会出现在文档中
"""

print(help(add))

"""
尽管你可以使用任意类型的对象给函数添加注解 (例如数字，字符串，对象实例等
等)，不过通常来讲使用类或着字符串会比较好点。
"""

"""
函数注解只存储在函数的 annotations 属性中。
"""
print(add.__annotations__)
# {'x': <class 'int'>, 'return': <class 'int'>, 'y': <class 'int'>}

"""
尽管注解的使用方法可能有很多种，但是它们的主要用途还是文档。因为 python
并没有类型声明，通常来讲仅仅通过阅读源码很难知道应该传递什么样的参数给这个
函数。这时候使用注解就能给程序员更多的提示，让他们可以争取的使用函数。
"""

def sta(a:str, b:str) -> str:
    return a+b

print(sta("a", "b"))