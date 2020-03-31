package chapter16;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;
import java.util.stream.Stream;

public class TestStream2 {
    public static void main(String[] args) throws FileNotFoundException {
        Map<String, Long> freq = new HashMap<>();
        File file = new File("words.txt");
        try (Stream<String> words = new Scanner(file).tokens()) {
            words.forEach(word -> {
                freq.merge(word.toLowerCase(), 1L, Long::sum);
            });
        }
    }
}