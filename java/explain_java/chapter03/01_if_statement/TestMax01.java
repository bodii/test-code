// 读入两个实数值，并显示其中较大的值

import java.util.Scanner;

class TestMax01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("a的值: "); int a = stdIn.nextInt();
		System.out.print("b的值: "); int b = stdIn.nextInt();

		int max;
		if (a > b) 
			max = a;
		else
			max = b;

		System.out.println("较大的值是" + max + "。");
	}
}
