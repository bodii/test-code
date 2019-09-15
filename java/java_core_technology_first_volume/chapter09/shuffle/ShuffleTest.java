package shuffle;

import java.util.*;

/**
 * This program deonstrates the random shutffle and sort algorithms.
 * @version 1.1 2019-09
 * @author wong
 */
public class ShuffleTest {
	public static void main(String[] args) {
		List<Integer> numbers = new ArrayList<>();
		for (int i = 1; i <= 49; i++)
			numbers.add(i);
		Collections.shuffle(numbers);
		List<Integer> winningCombination = numbers.subList(0, 6);
		Collections.sort(winningCombination);
		System.out.println(winningCombination);
	}
}
