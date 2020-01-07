package chapter09;

import java.util.Random;

public class ForRandom {

    // 现在选择随机数生成器时，大多使用ThreadLocalRandom.
    // 它会产生高质量的随机数，并且速度非常快。
    public static Random rnd = new Random();

    public static void main(String[] args) {
        int n = 2 * (Integer.MAX_VALUE / 3);
        int low = 0;

        for (int i = 0; i < 1000000; i ++) 
            if (random(n) < n / 2)
                low++;
        
        System.out.println(low);
    }

    public static int random(int n) {
        return Math.abs(rnd.nextInt()) % n;
    }
}