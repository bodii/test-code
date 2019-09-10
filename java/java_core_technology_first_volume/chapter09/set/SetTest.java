package set;

import java.util.*;

/**
 * This program uses a set to print all unique words in System.in.
 * @version 1.1 2019-09
 * @author wong
 */
public class SetTest {
	public static void main(String[] args) {
		Set<String> words = new HashSet<>();
		long totalTime = 0;

		try (Scanner stdIn = new Scanner(System.in)) {
			while (stdIn.hasNext() && !stdIn.hasNext("0")) {
				String word = stdIn.next();
				long callTime = System.currentTimeMillis();
				words.add(word);
				callTime = System.currentTimeMillis() - callTime;
				totalTime += callTime;
			}
		}

		Iterator<String> iter = words.iterator();
		for (int i = 1; i <= 20 && iter.hasNext(); i++)
			System.out.println(iter.next());
		System.out.println("...");
		System.out.println(words.size() + " distinct words. " + totalTime + " milliseconds.");
	}
}
