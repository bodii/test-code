// final数组

class FinalArray {
	public static void main(String[] args) {
		final int[] a = new int[5];
		a[0] = 10; // ok
		// a = null; // err
		// a = new int[10]; // err
	}
}
