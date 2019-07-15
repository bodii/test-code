// 汽车类－第2版的使用

class CarTester03 {
	public static void main(String[] args) {
		Car02 myCar = new Car02("爱车", 1885, 1490, 5220, 90.0, new Day02(2000, 11, 18));

		myCar.putSpec();
		System.out.println("购买日期: " + myCar.getPurchaseDay().toString());
	}
}
