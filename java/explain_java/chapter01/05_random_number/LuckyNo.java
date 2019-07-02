// 随机生成并显示一个0 ~ 9的幸运数字

import java.util.Random;

class LuckyNo {
	public static void main(String[] args) {
		Random rand = new Random();

		int lucky = rand.nextInt(10);  // 0 ~ 9的随机数

		System.out.println("今天的幸运数字是" + lucky + "。");
	}
}
