// 猜数字游戏(目标数字范围为0 ~ 99)

import java.util.*;

class Kazuate {
	public static void main(String[] args) {
		Random rand = new Random();
		Scanner stdIn = new Scanner(System.in);

		int no = rand.nextInt(100); // 目标数字：生成一个0 ~ 99 的随机数
		System.out.println("猜数字游戏开始!!");
		System.out.println("请猜一个0~99的数字。");

		int x;
		do {
			System.out.print("是多少呢: ");
			x = stdIn.nextInt();

			if (x > no)
				System.out.println("再小一点。");
			else if (x < no)
				System.out.println("再大一点。");
		} while (x != no);

		System.out.println("回答正确。");
	}
}
