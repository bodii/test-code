#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第八章：类与对象
Desc: 本章主要关注点的是和类定义有关的常见编程模型。包括让对象支持常见的
Python 特性、特殊方法的使用、类封装技术、继承、内存管理以及有用的设计模
式。

Title;             不用递归实现访问者模式
Issue:你使用访问者模式遍历一个很深的嵌套树形数据结构，并且因为超过嵌套层级限制
而失败。你想消除递归，并同时保持访问者编程模式。
Answer:通过巧妙的使用生成器可以在树遍历或搜索算法中消除递归。
"""

"""
在 8.21 小节中，我
们给出了一个访问者类。下面我们利用一个栈和生成器重新实现这个类
"""
import types

class Node:
    pass

class NodeVisitor:
    def visit(self, node):
        stack = [node]
        last_result = None
        while stack:
            try:
                last = stack[-1]
                if isinstance(last, types.GeneratorType):
                    stack.append(last.send(last_result))
                    last_result = None
                elif isinstance(last, Node):
                    stack.append(self._visit(stack.pop()))
                else:
                    last_result = stack.pop()
            except StopIteration:
                stack.pop()

    def _visit(self, node):
        methname = 'visit_' + type(node).__name__
        meth = getattr(self, methname, None)
        if meth is None:
            meth = self.generic_visit
        return meth(node)

    def generic_visit(self, node):
        raise RuntimeError('No {} method'.format('visit_' + types(node).__name__))

"""
如果你使用这个类，也能达到相同的效果。事实上你完全可以将它作为上一节中的
访问者模式的替代实现。考虑如下代码，遍历一个表达式的树：
"""

class UnaryOperator(Node):
    def __init__(self, operand):
        self.operand = operand

class BinaryOperator(Node):
    def __init__(self, left, right):
        self.left = left
        self.right = right

class Add(BinaryOperator):
    pass
class Sub(BinaryOperator):
    pass
class Mul(BinaryOperator):
    pass
class Div(BinaryOperator):
    pass
class Negate(UnaryOperator):
    pass

class Number(Node):
    def __init__(self, value):
        self.value = value

class Evaluator(NodeVisitor):
    def visit_Number(self, node):
        return node.value

    def visit_Add(self, node):
        return self.visit(node.left) + self.visit(node.right)

    def visit_Sub(self, node):
        return self.visit(node.left) - self.visit(node.right)

    def visit_Mul(self, node):
        return self.visit(node.left) * self.visit(node.right)

    def visit_Div(self, node):
        return self.visit(node.left) / self.visit(node.right)

    def visit_Negate(self, node):
        return -self.visit(node.operand)

if __name__ == '__main__':
    t1 = Sub(Number(3), Number(4))
    t2 = Mul(Number(2), t1)
    t3 = Div(t2, Number(5))
    t4 = Add(Number(1), t3)

    e = Evaluator()
#    print(e.visit(t4))

"""
现在我们稍微修改下上面的 Evaluator
"""
class Evaluator(NodeVisitor):
    def visit(self, node):
        return node.value

    def visit_Add(self, node):
        yield (yield node.left) + (yield node.right)

    def visit_Sub(self, node):
        yield (yield node.left) + (yield node.right)

    def visit_Mul(self, node):
        yield (yield node.left) * (yield node.right)

    def visit_Div(self, node):
        yield (yield node.left) / (yield node.right)

    def visit_Negate(self, node):
        yield - (yield node.operand)

a = Number(0)
print(a)
for n in range(1, 1000):
    a = Add(a, Number(n))



e = Evaluator()
e.visit(a)
"""
如果你还想添加其他自定义逻辑也没问题
"""

class Evaluator(NodeVisitor):
    def visit_Add(self, node):
        print('Add:', node)
        lhs = yield node.left
        print('left=', lhs)
        rhs = yield node.right
        print('right=', rhs)
        yield lhs + rhs

e = Evaluator()
e.visit(t4)
left = 1
right = -0.3
# 0.6
"""
这一小节我们演示了生成器和协程在程序控制流方面的强大功能。避免递归的一个
通常方法是使用一个栈或队列的数据结构。例如，深度优先的遍历算法，第一次碰到
一个节点时将其压入栈中，处理完后弹出栈。visit() 方法的核心思路就是这样。
另外一个需要理解的就是生成器中 yield 语句。当碰到 yield 语句时，生成器会返
回一个数据并暂时挂起。上面的例子使用这个技术来代替了递归。例如，之前我们是
这样写递归：

#value = self.visit(node.left)
# 现在换成“yield"语句：
#value = yield node.left
"""
"""
它会将 node.left 返回给 visti() 方法，然后 visti() 方法调用那个节点相应的
vist Name() 方法。yield 暂时将程序控制器让出给调用者，当执行完后，结果会赋值
给 value，
看完这一小节，你也许想去寻找其它没有 yield 语句的方案。但是这么做没有必要，
你必须处理很多棘手的问题。例如，为了消除递归，你必须要维护一个栈结构，如果
不使用生成器，代码会变得很臃肿，到处都是栈操作语句、回调函数等。实际上，使
用 yield 语句可以让你写出非常漂亮的代码，它消除了递归但是看上去又很像递归实
现，代码很简洁。
"""
