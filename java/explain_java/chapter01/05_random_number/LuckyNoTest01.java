// 随机生成并显示一位数的正整数(即1 ~ 9的值)

import java.util.Random;

class LuckyNoTest01 {
	public static void main(String[] args) {
		Random rand = new Random();

		System.out.println("一位数的正整数:" + (rand.nextInt(9) + 1));
	}
}
