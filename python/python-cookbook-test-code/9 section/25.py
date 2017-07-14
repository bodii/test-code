#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第九章：元编程
Desc: 软件开发领域中最经典的口头禅就是“don’t repeat yourself”。也就是说，任何时
候当你的程序中存在高度重复 (或者是通过剪切复制) 的代码时，都应该想想是否有更
好的解决方案。在 Python 当中，通常都可以通过元编程来解决这类问题。简而言之，
元编程就是关于创建操作源代码 (比如修改、生成或包装原来的代码) 的函数和类。主
要技术是使用装饰器、类装饰器和元类。不过还有一些其他技术，包括签名对象、使
用 exec() 执行代码以及对内部函数和类的反射技术等。本章的主要目的是向大家介绍
这些元编程技术，并且给出实例来演示它们是怎样定制化你的源代码行为的。

Title;                拆解 Python 字节码
Issue:你想通过将你的代码反编译成低级的字节码来查看它底层的工作机制。
Answer: dis 模块可以被用来输出任何 Python 函数的反编译结果。
"""
def countdown(n):
    while n > 0:
        print('T-minus', n)
        n -= 1
    print('Blastoff!')

import dis
print(dis.dis(countdown))
'''
 19           0 SETUP_LOOP              39 (to 42)
        >>    3 LOAD_FAST                0 (n)
              6 LOAD_CONST               1 (0)
              9 COMPARE_OP               4 (>)
             12 POP_JUMP_IF_FALSE       41

 20          15 LOAD_GLOBAL              0 (print)
             18 LOAD_CONST               2 ('T-minus')
             21 LOAD_FAST                0 (n)
             24 CALL_FUNCTION            2 (2 positional, 0 keyword pair)
             27 POP_TOP

 21          28 LOAD_FAST                0 (n)
             31 LOAD_CONST               3 (1)
             34 INPLACE_SUBTRACT
             35 STORE_FAST               0 (n)
             38 JUMP_ABSOLUTE            3
        >>   41 POP_BLOCK

 22     >>   42 LOAD_GLOBAL              0 (print)
             45 LOAD_CONST               4 ('Blastoff!')
             48 CALL_FUNCTION            1 (1 positional, 0 keyword pair)
             51 POP_TOP
             52 LOAD_CONST               0 (None)
             55 RETURN_VALUE
None
'''

"""
当你想要知道你的程序底层的运行机制的时候， dis 模块是很有用的。比如如果你
想试着理解性能特征。被 dis() 函数解析的原始字节码如下所示：
"""
print(countdown.__code__.co_code)
'''
b"x'\x00|\x00\x00d\x01\x00k\x04\x00r)\x00t\x00\x00d\x02\x00|\x00\x00\x83\x02\x00\x01|\x00\x00d\x03\x008}\x00\x00q\x03\x00Wt\x00\x00d\x04\x00\x83\x01\x00\x01d\x00\x00S"
'''
"""
如果你想自己解释这段代码，你需要使用一些在 opcode 模块中定义的常量。
"""
c = countdown.__code__.co_code
import opcode
print(opcode.opname[c[0]])
# SETUP_LOOP
print(opcode.opname[c[3]])
# LOAD_FAST

"""
奇怪的是，在 dis 模块中并没有函数让你以编程方式很容易的来处理字节码。不
过，下面的生成器函数可以将原始字节码序列转换成 opcodes 和参数。
"""
import opcode

def generate_opcodes(codebytes):
    extended_arg = 0
    i = 0
    n = len(codebytes)
    while i < n:
        op = codebytes[i]
        i += 1
        if op >= opcode.HAVE_ARGUMENT:
            oparg = codebytes[i] + codebytes[i+1]*256 + extended_arg
            extended_arg = 0
            i += 2
            if op == opcode.EXTENDED_ARG:
                extended_arg = oparg * 65536
                continue
            else:
                oparg = None
            yield(op, oparg)

for op, oparg in generate_opcodes(countdown.__code__.co_code):
    print(op, opcode.opname[op], oparg)

'''
120 SETUP_LOOP None
124 LOAD_FAST None
100 LOAD_CONST None
107 COMPARE_OP None
114 POP_JUMP_IF_FALSE None
116 LOAD_GLOBAL None
100 LOAD_CONST None
124 LOAD_FAST None
131 CALL_FUNCTION None
124 LOAD_FAST None
100 LOAD_CONST None
125 STORE_FAST None
113 JUMP_ABSOLUTE None
116 LOAD_GLOBAL None
100 LOAD_CONST None
131 CALL_FUNCTION None
100 LOAD_CONST None
'''

"""
这种方式很少有人知道，你可以利用它替换任何你想要替换的函数的原始字节码。
下面我们用一个示例来演示整个过程：
"""
def add(x, y):
    return x + y

c = add.__code__
print(c)
#<code object add at 0x0000017348C35AE0, file "", line 124>
print(c.co_code)
# b'|\x00\x00|\x01\x00\x17S'
import types
newbytecode = b'xxxxxxx'
nc = types.CodeType(c.co_argcount, c.co_kwonlyargcount,
                    c.co_nlocals, c.co_stacksize, c.co_filename,
                    c.co_name,c.co_firstlineno, c.co_lnotab
                    )
print(nc)

#add.__code__ = nc
print(add(2, 3))
"""
你可以像这样耍大招让解释器奔溃。但是，对于编写更高级优化和元编程工具的程
序员来讲，他们可能真的需要重写字节码。本节最后的部分演示了这个是怎样做到的。
你还可以参考另外一个类似的例子： this code on ActiveState
"""