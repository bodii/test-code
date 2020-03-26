package chapter15;


public enum BasicOperation4 implements Operation4 {
    PLUS("+") {
        public double apply(double x, double y) { return x + y; }
    },
    MINUS("－") {
        public double apply(double x, double y) { return x - y; }
    },
    TIMES("*") {
        public double apply(double x, double y) { return x * y; }
    },
    DIVIDE("/") {
        public double apply(double x, double y) { return x / y; }
    };

    private final String symbol;

    BasicOperation4(String symbol) {
        this.symbol = symbol;
    }

    @Override public String toString() {
        return symbol;
    }
}