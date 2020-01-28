package linked_list_set;

import linked_list_set.LinkedListSet;

public class TestLinkedListSet {
    public static void main (String[] args) {
        LinkedListSet<Integer> set = new LinkedListSet<>();

        int[] nums = {5, 3, 6, 8, 4, 2};
        for (int num : nums)
            set.add(num);

        System.out.println("生成的链表集合：" + set);

        System.out.println("删除元素:");
        set.remove(6);
        System.out.println("删除元素后的链表集合: " + set);
        set.remove(5);
        System.out.println("删除元素后的链表集合: " + set);
    }
}