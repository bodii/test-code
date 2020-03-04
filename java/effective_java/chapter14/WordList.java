package chapter14;

import java.util.Set;
import java.util.TreeSet;
import java.util.Collections;


public class WordList {
    public static void main(String[] args) {
        Set<String> s = new TreeSet<>();
        Collections.addAll(s, args);
        System.out.println(s);
    }
}