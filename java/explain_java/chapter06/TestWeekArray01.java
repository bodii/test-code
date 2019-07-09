// 根据显示的星期，输入其英语表达

import java.util.Scanner;
import java.util.Random;

class TestWeekArray01 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		Random rand = new Random();

		String[] weeks =  {"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"};
		String[] weeks_zh =  {"星期－", "星期二", "星期三", "星期四", "星期五", "星期六", "星期七" };
		System.out.println("请用小写输入英文的星期名。");

		int w = rand.nextInt(7);

		boolean next = true;
		while (next) {
			System.out.printf("%s: ", weeks_zh[w]); String a = stdIn.next();
			if (weeks[w].equals(a)) {
			Input: {
				System.out.print("回答正确。再来一次？1...Yes / 0...No: "); int t = stdIn.nextInt();
				if (t == 1)
					w = rand.nextInt(7);
				else if (t == 0)
					next = false;
				else 
					break Input;
			}

			} else {
				System.out.println("回答错误。");
			}
		}
	}
}
