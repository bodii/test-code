// 图形类群
class Shape04Tester {
	public static void main(String[] args) {
		Shape04[] p = new Shape04[4];
		
		p[0] = new Point04();				// 点
		p[1] = new HorzLine04(5);			// 横线
		p[2] = new VertLine04(3);			// 竖线
		p[3] = new Rectangle04(4, 3);		// 长方形

		for (Shape04 s : p) {
			s.print();
			System.out.println();
		}
	}
}
