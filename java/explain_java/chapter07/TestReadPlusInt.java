// 显示出输入的正整数值，如果输入0或者负数，则提示再次输入

import java.util.Scanner;

class TestReadPlusInt {

	static Scanner stdIn = new Scanner(System.in);

	static int readPlusInt() {
		System.out.print("正整数值: "); 
		return TestReadPlusInt.stdIn.nextInt(); 
	}

	public static void main(String[] args) {
		while (readPlusInt() <= 0) {
		}
	}
}
