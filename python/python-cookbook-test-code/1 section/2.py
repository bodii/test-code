#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:   解压可迭代对象赋值给多个变量
Issue: 如果一个可迭代对象的元素个数超过变量个数时，会抛出一个 ValueError 。那么
怎样才能从这个可迭代对象中解压出 N 个元素出来？
answer: Python 的星号表达式可以用来解决这个问题。比如，你在学习一门课程，在学期
末的时候，你想统计下家庭作业的平均成绩，但是排除掉第一个和最后一个分数。如
果只有四个分数，你可能就直接去简单的手动赋值，但如果有 24 个呢？这时候星号表
达式就派上用场了
"""
# drop_first_last function
def drop_first_last(grades):
    (first,*middle,last) = grades
    return middle.avg()

record = ('Dave', 'dave@example.com', '773-555-1212', '847-555-1212')

name,email,*phone_numbers = record
print(name) # Dave
print(email) # dave@example.com
print(phone_numbers) # ['773-555-1212', '847-555-1212']

'''
星号表达式也能用在列表的开始部分。比如，你有一个公司前 8 个月销售数据的序
列，但是你想看下最近一个月数据和前面7个月的平均值的对比。

*trailing_qtrs, current_qtr = sales_record
trailing_avg = sum(trailing_qtrs) / len(trailing_qtrs)
return avg_comparison(trailing_avg, current_qtr)
'''
*trailing, current = [10, 8, 7, 1, 9, 5, 10, 3]
print(trailing) # [10, 8, 7, 1, 9, 5, 10]
print(current) # 3

'''
扩展的迭代解压语法是专门为解压不确定个数或任意个数元素的可迭代对象而设计
的。通常，这些可迭代对象的元素结构有确定的规则（比如第 1 个元素后面都是电话
号码），星号表达式让开发人员可以很容易的利用这些规则来解压出元素来。而不是通
过一些比较复杂的手段去获取这些关联的的元素值。
'''

records = [
    ('foo', 1, 2),
    ('bar', 'hello'),
    ('foo', 3, 4),
]

def do_foo(x,y):
    print('foo', x, y)

def do_bar(s):
    print('bar', s)

for tag, *args in records:
    if tag == 'foo':
        do_foo(*args)
    elif tag == 'bar':
        do_bar(*args)

#
        '''
        foo 1 2
        bar hello
        foo 3 4
        '''
#

# 星号解压语法在字符串操作的时候也会很有用，比如字符串的分割。
line = 'nobody:*:-2:-2:Unprivileged User:/var/empty:/usr/bin/false'
uname, *fields, homedir, sh = line.split(':')
print(uname) #nobody
print(homedir) #/var/empty
print(sh) # /usr/bin/false
print(fields) # ['*', '-2', '-2', 'Unprivileged User']

#有时候，你想解压一些元素后丢弃它们，你不能简单就使用 * ，但是你可以使用一个普通的废弃名称，比如或者ign 。
record = ('ABCD', 50, 100, (12,18,2020))
name, *_, (*_,year) = record
print(name) # ABCD
print(year) # 2020

#在很多函数式语言中，星号解压语法跟列表处理有许多相似之处。比如，如果你有一个列表，你可以很容易的将它分割成前后两部分
items = [1, 10, 7, 4, 5, 9]
head, *tail = items
print(head) # 1
print(tail) # [10, 7, 4, 5, 9]

# 还能用这种分割语法去巧妙的实现递归算法。
def sum(items):
    head, *tail = items
    return head + sum(tail) if tail else head

print(sum(items)) # 36
