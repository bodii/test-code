package com.design_pattern.pattern.behavioral.interpreter;

public class PlusInterpreter implements Interpreter {
    private Interpreter firstExpression, secondExpression;

    public PlusInterpreter(Interpreter firstExpression, Interpreter secondExpression) {
        this.firstExpression = firstExpression;
        this.secondExpression = secondExpression;
    }

    @Override
    public int interpreter() {
        return firstExpression.interpreter() + secondExpression.interpreter();
    }

    @Override
    public String toString() {
        return "+";
    }
}
