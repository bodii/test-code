package merge_two_sorted_lists;

class Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        return new ListNode(1);
    }
}


// Definition for singly-linked list.
public class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

class Test {
    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] n1 = {1,2,4}, n2 = {1,3,4};
        ListNode l1, l2;
        for (int i = 0; i < n1.length; i ++, l1 = l1.next)
            l1 = new ListNode(n1[i]);

        for (int i = 0;i < n2.length; i ++, l2 = l2.next)
            l2 = new ListNode(n2[i]);

        NodeList result = solution.mergeTwoLists(l1, l2);
        

    }
}