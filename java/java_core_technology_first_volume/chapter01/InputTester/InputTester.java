import java.util.*;

/**
 * This program demonstrates console input.
 * @version 1.10 2019-08-11
 * @auth wong
 */

public class InputTester {
	public static void main(String[] args) {
		Scanner stdIn = new Scanner(System.in);

		// get first input
		System.out.print("What is your name? ");
		String name = stdIn.nextLine();

		// get second input
		System.out.print("How old are you ? ");
		int age = stdIn.nextInt();

		// display output on console
		System.out.println("Hello, " + name + ".Next year, you'll be " + (age + 1));
	}
}
