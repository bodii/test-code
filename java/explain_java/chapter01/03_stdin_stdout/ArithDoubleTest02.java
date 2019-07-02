/**
 *
 * test
 * 请编写一段程序，读入三角形的底和高，并显示其面积
 *
 */

import java.util.Scanner;

class ArithDoubleTest02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.println("读入三角形的底和高，并显示其面积>>");

		System.out.print("底: ");
		double bottom = stdIn.nextDouble();

		System.out.print("高: ");
		double high = stdIn.nextDouble();

		System.out.println("面积: " + bottom * high / 2);
	}
}
