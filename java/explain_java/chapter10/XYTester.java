// 带标识编号的XY类

class XY {
	private static int counter = 0;
	private int id;

	private int x = 0;
	private int y = 0;

	{			// 实例初始化器
		id = ++counter;
	}

	public XY() { }
	public XY(int x) { this.x = x; }
	public XY(int x, int y) { this.x = x; this.y = y; }

	public String toString() {
		return "No." + id + " ... (" + x + ", " + y + ")";
	}

}


public class XYTester {
	public static void main(String[] args) {
		XY a = new XY();
		XY b = new XY(10);
		XY c = new XY(20, 30);
		
		System.out.println("a = " + a);
		System.out.println("b = " + b);
		System.out.println("c = " + c);
	}
}
