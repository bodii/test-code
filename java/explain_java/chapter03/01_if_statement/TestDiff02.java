// 读入两个整数值，如果它们的差值小于等于10，则显示“它们的差值小于10”，
// 否则显示“它们的差值大于等于11“

import java.util.Scanner;

class TestDiff02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("a的值： "); int a = stdIn.nextInt();
		System.out.print("b的值： "); int b = stdIn.nextInt();

		int diff = a - b;
		if (10 >= diff) 
			System.out.println("它们的差值小于等于10");
		else
			System.out.println("它们的差值大于等于11");
	}
}
