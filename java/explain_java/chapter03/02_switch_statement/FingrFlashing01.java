// 根据读入的值显示对应的猜拳手势(if)

import java.util.Scanner;

class FingerFlashing01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("请选择手势(0--石头/1--剪刀/2--布): ");
		int hand = stdIn.nextInt();

		switch (hand) {
		case 0:
			System.out.println("石头");
			break;
		case 1:
			System.out.println("剪刀");
			break;
		case 2:			
			System.out.println("布");
			break;
		}

	}
}
