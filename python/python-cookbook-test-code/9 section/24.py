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

Title;                解析与分析 Python 源码
Issue:你想写解析并分析 Python 源代码的程序。
Answer: Python 能够计算或执行字符串形式的源代码.尽管如此，ast 模块能被用来将 Python 源码编译成一个可被分析的抽象语法树
（AST）。
"""
x = 42
print(eval('2 + 3 * 4 + x'))
# 56
exec('for i in range(10): print(i)')
'''
0
1
2
3
4
5
6
7
8
9
'''
import ast

ex = ast.parse('2 + 3 * 4 + x', mode='eval')
print(ex)
# <_ast.Expression object at 0x0000023F3E072CC0>
print(ast.dump(ex))
'''
Expression(body=BinOp(left=BinOp(left=Num(n=2), op=Add(), right=BinOp(left=Num(n=3), op=Mult(), right=Num(n=4))), op=Add(), right=Name(id='x', ctx=Load())))
'''
top = ast.parse('for i in range(10): print(i)', mode='exec')
print(top)
# <_ast.Module object at 0x00000204AF2FB978>
print(ast.dump(top))
'''
Module(body=[For(target=Name(id='i', ctx=Store()), iter=Call(func=Name(id='range', ctx=Load()), args=[Num(n=10)], keywords=[]), body=[Expr(value=Call(func=Name(id='print', ctx=Load()), args=[Name(id='i', ctx=Load())], keywords=[]))], orelse=[])])
'''

"""
分析源码树需要你自己更多的学习，它是由一系列 AST 节点组成的。分析这
些节点最简单的方法就是定义一个访问者类，实现很多 visit NodeName() 方法，
NodeName() 匹配那些你感兴趣的节点。下面是这样一个类，记录了哪些名字被加载、
存储和删除的信息。
"""
class CodeAnalyzer(ast.NodeVisitor):
    def __init__(self):
        self.loaded = set()
        self.stored = set()
        self.deleted = set()

    def visit_Name(self, node):
        if isinstance(node.ctx, ast.Load):
            self.loaded.add(node.id)
        elif isinstance(node.ctx, ast.Store):
            self.stored.add(node.id)
        elif isinstance(node.ctx, ast.Del):
            self.deleted.add(node.id)

if __name__ == '__main__':
    code = '''for i in range(10):print(i)\ndel i'''

    top = ast.parse(code, mode='exec')

    c = CodeAnalyzer()
    c.visit(top)
    print('Loaded:', c.loaded)
    print('Stored:', c.stored)
    print('Deleted:', c.deleted)
'''
Loaded: {'range', 'i', 'print'}
Stored: {'i'}
Deleted: {'i'}
'''
"""
最后，AST 可以通过 compile() 函数来编译并执行。
"""
exec(compile(top, '<stdin>', 'exec'))
'''
0
1
2
3
4
5
6
7
8
9
'''

"""
当你能够分析源代码并从中获取信息的时候，你就能写很多代码分析、优化或验证
工具了。例如，相比盲目的传递一些代码片段到类似 exec() 函数中，你可以先将它转
换成一个 AST，然后观察它的细节看它到底是怎样做的。你还可以写一些工具来查看
某个模块的全部源码，并且在此基础上执行某些静态分析。
需要注意的是，如果你知道自己在干啥，你还能够重写 AST 来表示新的代码。下
面是一个装饰器例子，可以通过重新解析函数体源码、重写 AST 并重新创建函数代码
对象来将全局访问变量降为函数体作用范围，
"""
import ast
import inspect

class NameLower(ast.NodeVisitor):
    def __init__(self, lowered_names):
        self.lowered_names = lowered_names

    def visit_FunctionDef(self, node):
        code = '__globals = globals()\n'
        code += '\n'.join("{0} = __globals['{0}']".format(name)
                          for name in self.lowered_names)
        code_ast = ast.parse(code, mode='exec')
        node.body[:0] = code_ast.body
        self.func = node

def lower_names(*namelist):
    def lower(func):
        srclines = inspect.getsource(func).splitlines()
        for n, line in enumerate(srclines):
            if '@lower_names' in line:
                break

        src = '\n'.join(srclines[n+1:])
        if src.startswith((' ', '\t')):
            src = 'if 1:\n' + src
        top = ast.parse(src, mode='exec')

        c1 = NameLower(namelist)
        c1.visit(top)

        temp = {}
        exec(compile(top, '', 'exec'), temp, temp)
        func.__code__ = temp[func.__name__].__code__
        return func
    return lower

"""
为了使用这个代码，你可以像下面这样写：
"""
INCR = 1
@lower_names('INCR')
def countdown(n):
    while n > 0:
        n -= INCR
"""
装饰器会将 countdown() 函数重写为类似下面这样子：
"""
def countdown(n):
    __globals = globals()
    INCR = __globals['INCR']
    while n > 0:
        n -= INCR

"""
在性能测试中，它会让函数运行快 20%
现在，你是不是想为你所有的函数都加上这个装饰器呢？或许不会。但是，这却是
对于一些高级技术比如 AST 操作、源码操作等等的一个很好的演示说明
本节受另外一个在 ActiveState 中处理 Python 字节码的章节的启示。使用 AST
是一个更加高级点的技术，并且也更简单些。参考下面一节获得字节码的更多信息。
"""