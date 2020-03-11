package chapter14;

import java.util.List;
import java.util.TreeSet;

public class TestInstrumentedSet {
    public static void main(String[] args) {
        InstrumentedSet<Integer> in = new InstrumentedSet<>(new TreeSet<>());

        in.addAll(List.of(5, 8, 3));

        System.out.println("count: " + in.getAddCount());
    }
}