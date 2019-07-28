// 一个三个猜拳程序，由计算机担任其中两个角色

import java.util.Scanner;
import java.util.Random;

class FingerFlashing2 {
	public static Scanner stdIn = new Scanner(System.in);
	public static String[] hands = { "石头", "剪刀", "布" };

	public static int punch() {
		int user;
		do {
			System.out.print("石头剪刀布");
			for (int i = 0; i < 3; i++)
				System.out.printf("(%d)%s ", i, hands[i]);
			System.out.print(": ");
			user = stdIn.nextInt();
		} while (user < 0 || user > 2); 
		return user;
	}

	public static void main(String[] args) {
		Random rand = new Random();

		int retry;  // 重新开始

		do {
			int comp1 = rand.nextInt(3), comp2 = rand.nextInt(3);
			int user = punch();

			// 显示三方的手势
			System.out.printf(
					"电脑comp1 出 %s, comp2 出 %s, 你出的 %s\n",
					hands[comp1], hands[comp2], hands[user]
					);

			int judge = (user - comp1 + 3) % 3;
			System.out.println("judge: " + judge);
			if (judge == 1) 
				System.out.println("你输了。"); 
			else if (judge == 2 || judge == 0) { 
				judge = (user - comp2 + 3) % 3;
				// 判断
				switch (judge) {
					case 0: System.out.println("平局。"); break;
					case 1: System.out.println("你输了。"); break;
					case 2: System.out.println("你赢了。"); break;
				}
			}

			do {
				System.out.print("再来一次? [0]否 [1]是");
				retry = stdIn.nextInt();
			} while (retry != 0 && retry != 1); 
		
		} while (retry == 1);

	}
}
