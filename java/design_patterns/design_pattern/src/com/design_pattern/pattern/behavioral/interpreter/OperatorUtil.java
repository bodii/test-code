package com.design_pattern.pattern.behavioral.interpreter;

public class OperatorUtil {
    public static boolean isOperator(String symbol) {
        return (symbol.equals("+") || symbol.equals("*"));
    }

    public static Interpreter getExpressionOjbect(Interpreter first, Interpreter second, String symbol) {
        if (symbol.equals("+"))
            return new PlusInterpreter(first, second);
        else if (symbol.equals("*"))
            return new MultiplyInterpreter(first, second);

        return null;
    }
}
