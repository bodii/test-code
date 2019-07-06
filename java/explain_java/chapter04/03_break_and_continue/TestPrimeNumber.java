// 写一段程序、判断读入的正整数值是否是质数。
// 质数，就是不可以被大于等于2且小于n中的任何一个数整除的整数n.

import java.util.Scanner;

class TestPrimeNumber {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("正整数: "); int n = stdIn.nextInt();
		for (int i = 2; i < n; i++) {
			if (n % i == 0) 
				System.out.println(i + "不是" + n + "的质数");
			else
				System.out.println(i + "是" + n + "的质数");
		}

	}
}
