// 汽车类－第2版的使用

class CarTester02 {
	public static void main(String[] args) {
		Day02 d = new Day02(2010, 10, 15);
		Car02 myCar = new Car02("爱车", 1885, 1490, 5220, 90.0, d);

		Day02 p = myCar.getPurchaseDay();
		System.out.println("爱车的购买日期: " + p);

		p.set(1999, 12, 31);

		Day02 q = myCar.getPurchaseDay();
		System.out.println("爱车的购买日期: " + q);
	}
}
