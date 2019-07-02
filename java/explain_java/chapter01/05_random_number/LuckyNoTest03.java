// 随机生成并显示两位数的正整数(即10 ~ 99 的值)

import java.util.Random;

class LuckyNoTest03 {
	public static void main(String[] args) {
		Random rand = new Random();

		System.out.println("显示两位数的正整数:" + (rand.nextInt(90) + 10));
	}
}
