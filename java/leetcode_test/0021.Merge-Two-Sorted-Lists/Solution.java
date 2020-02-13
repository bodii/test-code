package merge_two_sorted_lists;

import java.util.ArrayList;
import java.util.Collections;

class Solution {
    private ListNode add(ListNode node, int ele) {
        if (node.next == null)
            node.next = new ListNode(ele);
        else    
            node.next = add(node.next, ele);

        return node;
    }

    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if (l1 == null && l2 == null)
            return null;
            
        ArrayList<Integer> arr = new ArrayList<>();
        for (ListNode l = l1; l != null; l = l.next) 
            arr.add(l.val);

        for (ListNode l = l2; l != null; l = l.next) 
            arr.add(l.val);

        if (arr.isEmpty())
            return null;

        Collections.sort(arr);

        ListNode result = new ListNode(arr.get(0));
        for (int i = 1; i < arr.size(); i ++)
            result = add(result, arr.get(i));

       return result;
    }
}


// Definition for singly-linked list.
class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class LinkedList {
    public static ListNode add(ListNode node, int ele) {
        if (node.next == null)
            node.next = new ListNode(ele);
        else    
            node.next = add(node.next, ele);

        return node;
    }

    public static void print(ListNode node) {
        if (node == null)
            return;

        System.out.print(node.val);
        if (node.next == null) 
            return;
        System.out.print("->");

        print(node.next);
    }
}

class Test {
    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] n1 = {1,2,4}, n2 = {1,3,4};
        ListNode l1 = new ListNode(n1[0]);
        ListNode l2 = new ListNode(n2[0]);

        for (int i = 1;i < n1.length; i ++)
            l1 = LinkedList.add(l1, n1[i]);
        for (int i = 1; i < n2.length; i ++)
            l2 = LinkedList.add(l2, n2[i]);

        // LinkedList.print(l1);

        ListNode merge = solution.mergeTwoLists(l1, l2);
        LinkedList.print(merge);
    }
}