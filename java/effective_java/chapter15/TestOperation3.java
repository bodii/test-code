package chapter15;

public class TestOperation3 {
    public static void main(String[] args) {
        double x = Double.parseDouble(args[0]);
        double y = Double.parseDouble(args[1]);
        for (Operation3 p : Operation3.values())
            System.out.printf("%f %s %f = %f%n", x, p, y, p.apply(x, y));
    }
}