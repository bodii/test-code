import java.util.*;

/**
 * This program demonstrates a <code> for </code> loop.
 * @version 1.10 2019-08-11
 * @author wong
 */

public class LotteryOdds {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("How many numbers do you need to draw? "); 
		int k = stdIn.nextInt();

		System.out.print("Wath is the highest number you can draw? ");
		int n = stdIn.nextInt();

		/*
		 * compute binomial coefficient n*(n-1)*(n-2)*..*(n-k+1) / (1*2*3...*k)
		 */
		int lotteryOdds = 1;
		for (int i = 1; i <= k; i++) 
			lotteryOdds = lotteryOdds * (n - i + 1) / i;

		System.out.println("Your odds are 1 in  " + lotteryOdds + ". Good lock!");
	}
}
