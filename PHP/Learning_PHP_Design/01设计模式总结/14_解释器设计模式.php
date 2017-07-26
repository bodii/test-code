<?php

/*
     [ 解释器设计模式 ]

     Interpreter Pattern
     提供了评估语言的语法或表达式的方式，它属于行为型模式。这种模式实现了一个表达式
     接口，该接口解释一个特定的上下文。这种模式被用在SQL解析、符号处理引擎等

*/

// 抽象表达式
abstract class Expression
{
    // 任何表达式子类都应该有一种解析任务
    abstract public function interpreter($context);
}

// 抽象表达式是生成语法集合（语法树）的关键，每个语法集合完成指定语法解析任务
// 抽象表达式通过递归调用的方法，最终由最小语法单元进行解析完成

// 终结符表达式 通常指运算变量
class TerminalExpression extends Expression
{
    public function interpreter($context)
    {
        return null;
    }
}

class NonterminalExpression extends Expression
{
    public function interpreter($context)
    {
        return null;
    }
}