package chapter18;

import java.util.HashSet;
import chapter18.ObservableSet;

public class TestObservableSet {
    public static void main(String[] args) {
        ObservableSet<Integer> set = 
            new ObservableSet<>(new HashSet<>());

        set.addObserver((s, e) -> System.out.println(e));

        for (int i = 0; i < 100; i ++)
            set.add(i);
    }
}