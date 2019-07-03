// 读入一个正整数值，如果它是10的倍数，则显示"该值是10的倍数"，否则显示"该值不是10的倍数"

import java.util.Scanner;

class Test03 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数值: "); int i = stdIn.nextInt();

		if ((i % 10) == 0)
			System.out.println("该值是10的倍数");
		else 
			System.out.println("该值不是10的倍数");
			
	}
}
