// 显示读入的整数值和实数值

import java.util.Scanner;

class DecimalFloat {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		
		System.out.print("整数x: "); int x = stdIn.nextInt();
		System.out.print("整数y: "); double y = stdIn.nextDouble();
		
		System.out.printf("x =%3d\ty=%6.2f\n", x, y);
	}
}
