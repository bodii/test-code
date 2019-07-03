// 根据通过键盘输入的分数来判断伏/良/及格/不及格
// 0 ~ 59 不及格 / 60 ~ 69 及格 / 70 ~ 79 良 / 80 ~ 100 优

import java.util.Scanner;

class Test04 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("分数: "); int n = stdIn.nextInt();
		if (0 <= n && n <= 59)
			System.out.println("不及格");
		else if (60 <= n && n <= 69)
			System.out.println("及格");
		else if (70 <= n && n <= 79)
			System.out.println("良");
		else if (80 <= n && n <= 100)
			System.out.println("优");
		else 
			System.out.println("输入分数不合法");
	}
}
