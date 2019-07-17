// 测试带标识编号的三维坐标类

public class Point3DTester {
	public static void main(String[] args) {
		Point3D p1 = new Point3D();
		Point3D p2 = new Point3D(1);
		Point3D p3 = new Point3D(2, 3);
		Point3D p4 = new Point3D(4, 5, 6);

		System.out.println("最后赋的标识编号: " + Point3D.getCounter());
		System.out.println("p1 = " + p1.getId());
		System.out.println("p2 = " + p2.getId());
		System.out.println("p3 = " + p3.getId());
	}
}
