// 编写一段英语单词学习程序，根据显示的月份数值1~12,输入其英语表达

import java.util.Random;
import java.util.Scanner;

class TestMonthArrays02 {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);
		Random rand = new Random();

		String[] monthString = {
			"Januar", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"
		};

		System.out.println("请输入月份的英语表达。\n另外，首字母大写，之后的字母都小写。");
		boolean next = true;
		int m = 1 + rand.nextInt(11);
		while (next) {
			System.out.printf("%d月: ", m + 1); String a = stdIn.next();
			if (monthString[m].equals(a)) {
			Input: {
				System.out.print("回答正确。再来一次? 1...Yes/ 0...No: ");
				int nextInt = stdIn.nextInt();
				if ( nextInt == 0 ) 
					next = false;
				else if (nextInt == 1)
					m = 1 + rand.nextInt(11);
				else 
					break Input;
			}
			} else {
				System.out.println("回答错误。");
			}
		}
	}
}
