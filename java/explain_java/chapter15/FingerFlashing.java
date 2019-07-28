// 猜拳

import java.util.Scanner;
import java.util.Random;

class FingerFlashing {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		Random rand = new Random();
		String[] hands = { "石头", "剪刀", "布" };

		int retry; // 重新开始

		do {
			// 生成0, 1, 2的随机数，作为计算机的手势
			int comp = rand.nextInt(3);

			// 读入玩家的手势0, 1, 2
			int user;
			do {
				System.out.print("石头剪刀布");
				for (int i = 0; i < 3; i++) 
					System.out.printf("(%d)%s ", i, hands[i]);
				System.out.print(": ");
				user =  stdIn.nextInt();
			} while( user < 0 || user > 2);

			// 显示双方的手势
			System.out.println("我出" + hands[comp] + ", 你出" + hands[user] + "。");

			// 判断
			int judge = (user - comp + 3) % 3;
			switch (judge) {
				case 0: System.out.println("平局。");   break;
				case 1: System.out.println("你输了。");   break;
				case 2: System.out.println("你赢了。");   break;
			}

			// 确认是否再来一次
			do {
				System.out.print("再来一次? [0]否 [1]是： ");
				retry = stdIn.nextInt();
			} while(retry != 0 && retry != 1);
		} while(retry == 1);
	}
}
