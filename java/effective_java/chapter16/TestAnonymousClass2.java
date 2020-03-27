package chapter16;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.function.Function;

/**
 * Anonymous class instance as a function object - obsolete!
 */
public class TestAnonymousClass2 {
    public static void main(String[] args) {
        String[] w = {"a", "ab", "b", "ce"};
        ArrayList<String> words = new ArrayList<>();
        for (String d : w)
            words.add(d);

        // Collections.sort(words, ComparingInt(String::length));
        words.sort(ComparingInt(String::length));
        for (String d : words)
            System.out.println(d);
    }
}