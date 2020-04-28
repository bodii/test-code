package chapter16;

import java.util.Random;

public class RandomCommon {
    public static Random rand = new Random();
    public static void main(String[] args) {
        int n = 2 * (Integer.MAX_VALUE / 3);
        int low = 0;
        for (int i = 0; i < 100000; i ++)
            if (random(n) < n/2)
                low++;
        System.out.println(low);
    }

    public static int random(int n) {
        return Math.abs(rand.nextInt()) % n;
    }
}