package chapter16;

import java.util.Set;
import java.util.TreeSet;
import java.util.List;
import java.util.ArrayList;

public class SetList2 {
    public static void main(String[] args) {
        Set<Integer> set = new TreeSet<>();
        List<Integer> list = new ArrayList<>();

        for (int i = -3; i < 3; i ++) {
            set.add(i);
            list.add(i);
        }

        for (int i = 0; i < 3; i ++) {
            set.remove(i);
            list.remove(Integer.valueOf(i));
            // list.remove((Integer) i)
        }

        System.out.println(set + " " + list);
    }
}