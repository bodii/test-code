// 读入两个整数值，显示其中较小的值和较大的值。不过，如果两个整数值相等，则显示
// “两个值相等”

import java.util.Scanner;

class TestMinMax02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数a: "); int a = stdIn.nextInt();
		System.out.print("整数b: "); int b = stdIn.nextInt();

		String s;
		if (a > b) {
			s = "大数:" + a + ", 小数:" + b; 
		} else if (a < b) {
			s = "大数:" + b + ", 小数:" + a; 
		} else { 
			s = "两个值相等"; 
		}

		System.out.println(s);

	}
}
