package chapter16;

import java.util.ArrayList;
import java.util.Collections;

/**
 * Lambda expression as function object(replaces anonymous class)
 */
public class TestLambdaExpression {
    public static void main(String[] args) {
        String[] w = {"a", "ab", "b", "ce"};
        ArrayList<String> words = new ArrayList<>();
        for (String d : w)
            words.add(d);

        Collections.sort(
            words, 
            (a, b) -> Integer.compare(a.length(), b.length())
        );

        for (String d : words)
            System.out.println(d);
    }
}