package com.easykotlin.test.chapter6;

import java.util.ArrayList;
import java.util.List;

public class ListUtilTest<T> {
    public List<T> filter(List<T> list, Predicate<T> p) {
        List<T> result = new ArrayList<>();
        for (T t : list) {
            if (p.predicate(t))
                result.add(t);
        }

        return result;
    }
}

interface Predicate<T> {
    Boolean predicate(T t);
}
