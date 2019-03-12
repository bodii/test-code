/**
 * 线性表选择排序法
 */

public class SelectionSort {
    private static SequenceList sort(SequenceList s1, int placenum) {
        if (placenum >= s1.size()) {
            return s1;
        }

        int head = Integer.parseInt(s1.get(placenum).toString());
        int element = Integer.parseInt(s1.get(placenum).toString());
        int elementPlace = 0;
        for (int i = placenum; i < s1.size(); i++) {
            if (Integer.parseInt(s1.get(i).toString()) < element) {
                element = Integer.parseInt(s1.get(i).toString());
                elementPlace = i;
            }
        }

        if (elementPlace != 0) {
            s1.set(placenum, element);
            s1.set(elementPlace, head);
        }

        sort(s1, ++placenum);

        return s1;
    }

    public static void main(String[] args) {
        SequenceList s1 = new SequenceList();
        s1.add(23);
        s1.add(45);
        s1.add(31);
        s1.add(17);
        s1.add(26);
        s1.add(44);

        SequenceList st = sort(s1, 0);
        System.out.println("size: " + st.size());
        for (int i = 0; i < st.size(); i++) {
            System.out.println(st.get(i));
        }
    }
}