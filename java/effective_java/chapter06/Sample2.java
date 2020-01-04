package chapter06;

import java.lang.annotation.*;

public class Sample2 {
    @IExceptionTest(ArithmeticException.class)
    public static void m1() {
        int i = 0;
        i = i / i;
    }

    @IExceptionTest(ArithmeticException.class) 
    public static void m2() {
        int[] a = new int[0];
        int i = a[1];
    }

    @IExceptionTest(ArithmeticException.class) 
    public static void m3() {  }
}