import java.util.*;

/**
 * This program demonstrates array manipulation.
 * @version 1.10 2019-08-11
 * @author wong
 */

public class LotteryDrawing {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		System.out.print("How many numbers do you need to draw? ");
		int k = stdIn.nextInt();

		System.out.print("What is the heighest number you can draw? ");
		int n = stdIn.nextInt();

		// fill an array with numbers 1 2 3 ... n
		int[] numbers = new int[n];
		for (int i = 0; i < n; i++) 
			numbers[i] = i + 1;

		int[] result = new int[k];
		for (int i = 0; i < result.length; i++) {
			// make a random index between 0 and n -1
			int r = (int) (Math.random() * n);

			// pick the element at the random location
			result[i] = numbers[r];

			// move the last element into the random location
			numbers[r] = numbers[n - 1];
			n--;
		}

		// print the sorted array
		Arrays.sort(result);
		System.out.println("Bet the following combination. It'll make you rich!");
		for (int r : result)
			System.out.println(r);
	}
}
