package com.design_pattern.pattern.behavioral.interpreter;

import java.util.Stack;

public class CalcExpressionParser {
    private Stack<Interpreter> stack = new Stack<>();

    public int parse(String str) {
        String[] strItemArray = str.split(" ");
        for (String symbol : strItemArray) {
            if (!OperatorUtil.isOperator(symbol)) {
                Interpreter numberInterpreter = new NumberInterpreter(symbol);
                stack.push(numberInterpreter);
                System.out.printf("入栈： %d\n", numberInterpreter.interpreter());
            } else {
                Interpreter first = stack.pop();
                Interpreter second = stack.pop();
                System.out.printf("出栈： %d 和 %d\n", first.interpreter(), second.interpreter());

                Interpreter operator = OperatorUtil.getExpressionOjbect(first, second, symbol);
                System.out.printf("应用运算符: %s\n", symbol);
                int result = operator != null ? operator.interpreter() : 0;

                stack.push(new NumberInterpreter(result));
                System.out.printf("阶段结果入栈： %d\n", result);
            }
        }
        Interpreter resultInterpreter = stack.pop();
        int result = resultInterpreter.interpreter();
        System.out.printf("结果是: %d\n", result);
        return result;
    }
}
