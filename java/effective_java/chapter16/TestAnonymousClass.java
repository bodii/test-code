package chapter16;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;

/**
 * Anonymous class instance as a function object - obsolete!
 */
public class TestAnonymousClass {
    public static void main(String[] args) {
        String[] w = {"a", "ab", "b", "ce"};
        ArrayList<String> words = new ArrayList<>();
        for (String d : w)
            words.add(d);

        Collections.sort(words, new Comparator<String>() {
            public int compare(String s1, String s2) {
                return Integer.compare(s1.length(), s2.length());
            }
        });

        for (String d : words)
            System.out.println(d);
    }
}