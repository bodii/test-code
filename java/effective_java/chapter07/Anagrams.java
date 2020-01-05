package chapter07;

import java.io.IOException;
import java.util.Map;
import java.util.Scanner;
import java.util.Set;
import java.util.HashMap;
import java.util.Arrays;
import java.util.TreeSet;
import java.io.File;

public class Anagrams {
    public static void main(String[] args) throws IOException {
        File dictionary = new File(args[0]);
        int minGroupSize = Integer.parseInt(args[1]);

        Map<String, Set<String>> groups = new HashMap<>();
        try (Scanner s = new Scanner(dictionary)) {
            while (s.hasNext()) {
                String word = s.next();
                groups.computeIfAbsent(
                        alphabetize(word), 
                        (unused) -> new TreeSet<>()
                    ).add(word);
            }
        }
        for (Set<String> group : groups.values())
            if (group.size() >= minGroupSize)
                System.out.println(group.size() + ":" + group);
    }

    private static String alphabetize(String s) {
        char[] a = s.toCharArray();
        Arrays.sort(a);
        return new String(a);
    }
}