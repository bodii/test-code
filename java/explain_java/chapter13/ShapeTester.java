// 图形类群使用示例

class ShapeTester {
	public static void main(String[] args)  {
		
		Shape[] a = new Shape[2];
		a[0] = new Point();				// 点
		a[1] = new Rectangle(4, 3);		// 长方形

		for (Shape s : a) {
			s.draw();
			System.out.println();
		}
	}
}
