// 写一个"猜数字游戏", 目标数字为2位的整数值(10 ~ 99)

import java.util.Scanner;
import java.util.Random;

class TestGuessNumber {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		Random rand = new Random();
		
		// 随机数0 ~ 99
		int no = rand.nextInt(100);
		System.out.println("游戏开始>>\n");

		System.out.print("请输入一个数值:"); 
		int x; 		
		do {
			x = stdIn.nextInt();
			if ( x > no ) 
				System.out.println("比这个数字小哟。\n");
			else if (x < no)
				System.out.println("比这个数字大哟。\n");
			else 
				System.out.println("回答正确。");

		} while(no != x);

	}
}
