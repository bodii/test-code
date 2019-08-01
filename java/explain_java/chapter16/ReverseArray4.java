import java.util.Scanner;

class ReverseArray4 {

	// -- 交换数组对应索引的位置
	static void swap(int[] a, int idx1, int idx2) {
		int t = a[idx1];
		a[idx1] = a[idx2];
		a[idx2] = t;
	}

	// -- 排序
	static void reverse(int[] a) {
		try {
			for (int i = 0; i < a.length / 2; i++)
				swap(a, i, a.length - i); // a.length - i - 1
		} catch (ArrayIndexOutOfBoundsException e) {
			e.printStackTrace();
			System.exit(1);
		} catch (NullPointerException e) {
			e.printStackTrace();
			System.exit(1);
		}
	}

	public static void main(String[] args) {

		Scanner stdIn = new Scanner(System.in);
		System.out.print("元素个数: "); int n = stdIn.nextInt();

		int[] a = new int[n];
		for (int i = 0; i < n; i++) { 
			System.out.printf("a[%d] = ", i); 
			a[i] = stdIn.nextInt();
		}

		reverse(a);

		for (int i = 0; i < n; i++) 
			System.out.printf("a[%d] = %d\n",  i,  a[i]);
	}
}
