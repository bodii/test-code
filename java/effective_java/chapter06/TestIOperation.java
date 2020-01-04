package chapter06;

import chapter06.IOperation;
import chapter06.ExtendedOperation;

class TestIOpration {
    public static void main(String[] args) {
        double x = Double.parseDouble(args[0]);
        double y = Double.parseDouble(args[1]);
        test(ExtendedOperation.class, x, y);
    }

    private static <T extends Enum<T> & IOperation> void test (
        Class<T> opEnumType, double x, double y
    ) {
        for (IOperation op : opEnumType.getEnumConstants())
            System.out.printf("%f %s %f = %f%n", x, op, y, op.apply(x, y));
    }
}