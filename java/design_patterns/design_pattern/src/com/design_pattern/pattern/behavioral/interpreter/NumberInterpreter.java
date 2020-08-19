package com.design_pattern.pattern.behavioral.interpreter;

public class NumberInterpreter implements Interpreter {
    private int number;

    public NumberInterpreter(Integer number) {
        this.number = number;
    }

    public NumberInterpreter(String number) {
        this.number = Integer.parseInt(number);
    }

    @Override
    public int interpreter() {
        return number;
    }
}
