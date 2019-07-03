// 读入两个变量a、b的值，按如下所示的内容显示它们之间的大小关系
// “a更大”、“b更大”、“a和b相等"

import java.util.Scanner;

class Test01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		
		System.out.print("整数a: "); int a = stdIn.nextInt();
		System.out.print("整数b: "); int b = stdIn.nextInt();
		
		if (a > b) 
			System.out.println("a更大。");
		else if (b > a) 
			System.out.println("b更大。");
		else 
			System.out.println("a和b相等。");

	}
}
