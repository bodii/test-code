package chapter15;

public class TestOperation4 {
    public static void main(String[] args) {
        double x = Double.parseDouble(args[0]);
        double y = Double.parseDouble(args[1]);
        test(ExtendedOperation4.class, x, y);
    }

    private static <T extends Enum<T> & Operation4> void test(
        Class<T> opEnumType, double x, double y
    ) {
        for (Operation4 op : opEnumType.getEnumConstants())
            System.out.printf("%f %s %f = %f%n", x, op, y, op.apply(x, y));
    }
}