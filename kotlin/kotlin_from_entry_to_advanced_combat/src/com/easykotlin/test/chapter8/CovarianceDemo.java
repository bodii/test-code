package com.easykotlin.test.chapter8;

import java.util.ArrayList;
import java.util.List;

public class CovarianceDemo {
    public static void main(String[] args) {
        Integer[] ints = new Integer[3];
        ints[0] = 1;
        ints[1] = 2;
        ints[2] = 3;

        Number[] numbers = new Number[3];
        numbers = ints;
        for (Number n : numbers)
            System.out.println(n);


        List<Integer> integerList = new ArrayList<>();
        integerList.add(0);
        integerList.add(1);
        integerList.add(2);
        List<Number> numberList = new ArrayList<>();
//        numberList = integerList; // error

        List<? extends Number> list = new ArrayList<>();
//        list.add(new Integer(1)); // error

        List<? super Number> list3 = new ArrayList<Number>();
        List<? super Number> list4 = new ArrayList<Object>();
        list3.add(new Integer(3));
        list4.add(new Integer(4));
    }
}
