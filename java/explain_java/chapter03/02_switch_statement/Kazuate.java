// 猜数字游戏(目标数字范围为0~99)

import java.util.Scanner;
import java.util.Random;

class Kazuate {
	public static void main(String[] args) {
		Random rand = new Random();
		Scanner stdIn = new Scanner(System.in);

		int no = rand.nextInt(100); // 目标数字： 生成一个0~99 的随机数
		System.out.println("猜数字游戏开始!");
		System.out.println("请猜一个0~99的数字。");

		int x;
		do {
			System.out.print("是多少呢: ");
			x = stdIn.nextInt();

			if (x > no) 
				System.out.println("比这个数字小哟。");
			else if (x < no)
				System.out.println("比这个数字大哟。");
		} while (x != no);

		System.out.println("回答正确。");
	}
}
