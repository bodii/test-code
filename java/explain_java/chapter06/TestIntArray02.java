// 创建一个元素类型为int型的数组，将1~10的随机数赋给数组的全部元素
// 大于等于1小于等于10的数值

import java.util.Scanner;
import java.util.Random;

class TestIntArray02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		Random rand = new Random();

		System.out.print("元素个数: "); int n = stdIn.nextInt();

		int[] a = new int[n];
		for (int i = 0; i < n; i++) {
			a[i] = 1 + rand.nextInt(10);
			System.out.println("a[" + i + "] = " + a[i]);
		}
	}
}
