// 显示对int型数值左移和右移后的值

import java.util.Scanner;

class ShiftOperation {
	// -- 显示int型的位构成 --
	static void printBits(int x) {
		for (int i = 31; i >= 0; i--) 
			System.out.print(((x >>> i & 1) == 1) ? '1' : '0');
	}

	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("整数: "); int x = stdIn.nextInt();
		System.out.print("移位的位数: "); int n = stdIn.nextInt();

		System.out.print("整数 = "); printBits(x);
		System.out.print("\nx << n = "); printBits(x << n);
		System.out.print("\nx >> n = "); printBits(x >> n);
		System.out.print("\nx >>> n = "); printBits(x >>> n);
		System.out.println();
	}
}
