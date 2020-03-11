package chapter14;

import java.util.List;

public class TestInstrumentedHashSet {
    public static void main(String[] args) {
        InstrumentedHashSet<Integer> in = new InstrumentedHashSet<>();

        in.addAll(List.of(5, 8, 3));

        System.out.println("count: " + in.getAddCount());
    }
}