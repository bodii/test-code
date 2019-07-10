// 显示int型整数的按位与、按位或、按位异或、按位取反

import java.util.Scanner;

class BitwiseOperation {
	// -- 显示int型的位构成 --
	static void printBits(int x) {
		for (int i = 31; i >= 0; i--)
			System.out.print(((x >>> i & 1) == 1) ? '1' : '0');
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("请输入两个整数。");
		System.out.print("a: "); int a = stdIn.nextInt();
		System.out.print("b: "); int b = stdIn.nextInt();

		System.out.print("a\t= "); printBits(a);
		System.out.print("\nb\t= "); printBits(b);
		System.out.print("\na & b\t= "); printBits(a & b);
		System.out.print("\na | b\t= "); printBits(a | b);
		System.out.print("\na ^ b\t= "); printBits(a ^ b);
		System.out.print("\n~a\t= "); printBits(~a);
		System.out.print("\n~b\t= "); printBits(~b);
		System.out.println();
	}
}
