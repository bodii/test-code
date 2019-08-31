package stackTrace;

import java.util.*;

public class StackTraceTest {
    public static int factorial(int n) {
        System.out.printf("factorial(%d):\n", n);
        Throwable t = new Throwable();
        StackTraceElement[] frames = t.getStackTrace();
        for (StackTraceElement f :  frames)
            System.out.println(f);

        int r;
        if (n <= 1) r = 1;
        else r = n * factorial(n - 1);
        System.out.println("return " + r);
        return r;
    }

    public static void main(String[] args) {
        Scanner stdIn = new Scanner(System.in);
        System.out.print("Enter n: ");
        int n = stdIn.nextInt();
        factorial(n);
    }
}