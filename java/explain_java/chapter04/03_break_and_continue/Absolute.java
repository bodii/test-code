// 计算负整数的绝对值(带标签的break语句的使用示例)

import java.util.Scanner;

class Absolute {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		a: {
			   System.out.print("负整数：");
			   int t = stdIn.nextInt();
			   if (t >= 0) break a;
			   t = -t;
			   System.out.println("绝对值是:" + t + "。");
		}
	}
}
