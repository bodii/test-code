package chapter07;

import java.io.IOException;
import java.nio.file.Paths;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.stream.Stream;

public class Anagrams2 {
    public static void main (String[] args) throws IOException {
        Path dictionary = Paths.get(args[0]);
        int minGroupSize = Integer.parseInt(args[1]);

        try (Stream<String> words = Files.lines(dictionary)) {
            words.collect(
                groupingBy(
                    word ->word.chars().sorted()
                        .collect(
                            StringBuilder::new,
                            (sb, c) -> sb.append((char) c
                        ), 
                    StringBuilder::append).toString()
                ).values()
                .stream().filter(group -> group.size() >= minGroupSize)
            .map(group  -> group.size() + ":" +  group)
            .forEach(System.out::println);
        }
    }
}