package chapter06;

import chapter06.IOperation;

public enum ExtendedOperation implements IOperation {
    EXP("^") {
        public double apply(double x, double y) { return Math.pow(x, y); }
    },
    REMAINDER("%") {
        public double apply(double x, double y) { return x % y; }
    };

    private final String symbol;

    ExtendedOperation(String symbol) {
        this.symbol = symbol;
    }

    @Override public String toString () {
        return symbol;
    }
}