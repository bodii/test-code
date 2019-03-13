/**
 * Fibonacci 斐波那契数列
 * 
 * F(n) = F(n-1) + F(n-2)
 * 
 * 1 1 2 3 5 8 13 21 39
 */

public class Fibonacci {

    private static int F1(int n) {
        int res = 0;
        if (n < 1) {
            return res;
        }

        if (n < 3) {
            return 1;
        }

        int left = 1;
        int right = 1;
        for (int i = 3; i <= n; i++) {
            res = left + right;
            left = right;
            right = res;
            System.out.println("i: " + i + "   res: " + res);
        }

        return res;
    }

    public static int F2(int n) {
        int res = 0;
        if (n < 1) {
            return res;
        }

        if (n < 3) {
            return 1;
        }

        return F2(n - 1) + F2(n - 2);
    }

    public static void main(String[] args) {
        int v1 = F1(9);
        System.out.println("for loop  of [v1] value is: " + v1);

        int v2 = F2(9);
        System.out.println("recur of [v2] value is: " + v2);

        for (int i = 1; i < 10; i++) {
            System.out.print(F2(i) + "   ");
        }
    }
}