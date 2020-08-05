package com.easykotlin.test.chapter5;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.function.Predicate;
import java.util.stream.Stream;

public class FilterOddsDemo {
    public static void main(String[] args) {
        List<Integer> list = Arrays.asList(new Integer[]{1, 2, 3, 4, 5, 6, 7});
        System.out.println(filterOdds(list));
    }

    public static List<Integer> filterOdds(List<Integer> list) {
        List<Integer> result = new ArrayList<>();

        for(Integer i : list) {
            if (odds(i)) {
                result.add(i);
            }
        }

        return result;
    }

    public static boolean odds(Integer i) {
        return i % 2 != 0;
    }
}
