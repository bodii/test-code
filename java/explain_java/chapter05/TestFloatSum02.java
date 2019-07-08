// int型变更从0递增到1000的值再除以1000的情形

class TestFloatSum02 {
	public static void main(String[] args) {
		int x = 0;
		for (int i = 1; i <= 1000; i++) {
			x += i;
		}
		System.out.println("x: " + (x / 1000));
	}
}
