// 随机生成显示一位数的负整数(即-9 ~ -1的值)

import java.util.Random;

class LuckyNoTest02 {
	public static void main(String[] args) {
		Random rand = new Random();

		System.out.println("随机生成显示一位数的负整数(即-9 ~ -1的值: " + (rand.nextInt(9) - 9));
	}
}
