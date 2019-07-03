// 读入的整数是正值还是非正值呢

import java.util.Scanner;

class PositiveNot {
	public static void main(String[] args)
	{
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数值: ");
		int n = stdIn.nextInt();

		if (n > 0) 
			System.out.println("该值为正。");
		else
			System.out.println("该值为0或负。");
	}
}
