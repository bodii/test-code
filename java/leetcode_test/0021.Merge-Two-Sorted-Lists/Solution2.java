package merge_two_sorted_lists;

import java.util.ArrayList;
import java.util.Collections;

import merge_two_sorted_lists.LinkedList;
import merge_two_sorted_lists.ListNode;

class Solution2 {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if (l1 == null)
            return l2;
        
        if (l2 == null)
            return l1;

        if (l1.val < l2.val) {
            l1.next = mergeTwoLists(l1.next, l2);
            return l1;
        } 

        l2.next = mergeTwoLists(l1, l2.next);
        return l2;
    }
}


class Test2 {
    public static void main(String[] args) {
        Solution2 solution2 = new Solution2();
        int[] n1 = {1,2,4}, n2 = {1,3,4};
        ListNode l1 = new ListNode(n1[0]);
        ListNode l2 = new ListNode(n2[0]);

        for (int i = 1;i < n1.length; i ++)
            l1 = LinkedList.add(l1, n1[i]);
        for (int i = 1; i < n2.length; i ++)
            l2 = LinkedList.add(l2, n2[i]);

        // LinkedList.print(l1);

        ListNode merge = solution2.mergeTwoLists(l1, l2);
        LinkedList.print(merge);
    }
}