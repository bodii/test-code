// 写一个"猜数字游戏", 目标数字为2位的整数值(10 ~ 99)

import java.util.Scanner;
import java.util.Random;

class TestGuessNumber02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		Random rand = new Random();

		int no = rand.nextInt(90);
		no += 10; 
		int x;

		System.out.println("游戏开始>>\n");

		do {
			System.out.print("请输入数值: ");
			x = stdIn.nextInt(); // 输入的数值

			String s; // 接收显示结果的字符串

			if (x > no)
				s = "比这个数字小哟。\n";
			else if (x < no)
				s = "比这个数字大哟。\n";
			else
				s = "回答正确。";

			System.out.println(s);

		}while (no != x);
	}
}
