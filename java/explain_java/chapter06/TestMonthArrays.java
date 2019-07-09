import java.util.Random;
import java.util.Scanner;

class TestMonthArrays {
	public static void main(String[] args) {
		Random rand = new Random();
		Scanner stdIn = new Scanner(System.in);

		String[] monthString = {
			"Januar", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"
		};


		int r = 1 + rand.nextInt(11);
		int m = 0;
		while (r != m ){
			System.out.print("输入月份: "); m = stdIn.nextInt();
			if (r == m)
				break;
			System.out.println("回答错误。");
		}

		System.out.println("回答正确。");
		System.out.println(monthString[r]);

	}
}
