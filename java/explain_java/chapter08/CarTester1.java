// 汽车类的使用

class CarTester1 {
	public static void main(String[] args) {
		Car vitz = new Car("威姿", 1660, 1500, 640, 40.0);
		Car march = new Car("玛驰", 1660, 1525, 3695, 41.0);

		vitz.putSpec();  // 显示型号
		System.out.println();
		march.putSpec(); // 显示型号
	}
}
