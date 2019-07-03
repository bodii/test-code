// 读入两个整数值，并显示它们的差值
 
import java.util.Scanner;

class TestDiff01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("a的值："); int a = stdIn.nextInt();
		System.out.print("b的值："); int b = stdIn.nextInt();

		int diff = a - b;
		System.out.println("a和b的差值是：" + diff + "。");
	}
}
