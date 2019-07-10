// 心算能力训练

import java.util.Scanner;
import java.util.Random;

class MentalArithmetic {
	static Scanner stdIn = new Scanner(System.in);

	// -- 确认是否继续 -- //
	static boolean confirmRetry () {
		int cont;
		do {
			System.out.print("再来一次? <Yes...1/No...0>: ");
			cont = stdIn.nextInt();
		}while (cont != 0 && cont != 1);

		return cont == 1;
	}

	public static void main(String[] args) {
		Random rand = new Random();

		System.out.println("心算能力训练!!");
		do {
			int x = rand.nextInt(900) + 100;
			int y = rand.nextInt(900) + 100;
			int z = rand.nextInt(900) + 100;

			while (true) {
				System.out.print(x + " + " + y + " + " + z + " = ");
				int k = stdIn.nextInt();
				if (k == x + y + z)
					break;
				System.out.println("回答错误!!");
			}
		} while (confirmRetry());
	}
}
