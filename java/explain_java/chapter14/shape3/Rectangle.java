// -- 长方形
public class Rectangle extends Shape implements Plane2D {
	private int width; 
	private int height;

	// 计算面积
	public int getArea() { return width * height; }

	public void draw() {
		for(int i = 0; i < height; i++) {
			for(int j = 0; j < width; j++) System.out.print('*');
			System.out.println();
		}
	}
}
