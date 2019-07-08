// float型变量从0.001为单位从0.0递增至1.0的情形

class TestFloatSum01 {
	public static void main(String[] args) {
		float x = 0.0f;

		for (int i = 1; i < 1000; i++) {
			x += (float)i * 0.001;
			System.out.println("值: " + ((float)i * 0.001));
		}

		System.out.println("合计: " + x);
	}
}
