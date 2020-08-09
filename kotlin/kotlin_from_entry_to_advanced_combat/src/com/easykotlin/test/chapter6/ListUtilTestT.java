package com.easykotlin.test.chapter6;

import java.util.Arrays;
import java.util.List;

public class ListUtilTestT {
    public static void main(String[] args) {
        List<Integer> list = Arrays.asList(new Integer[] { 1, 2, 3, 4, 5, 6, 7});
        ListUtilTest<Integer> listUtilTest = new ListUtilTest<>();
        List<Integer> result = listUtilTest.filter(list, (it) -> it % 2 == 1);

        System.out.println(result);
    }
}
