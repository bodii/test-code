package deleteDuplicates;

import java.util.ArrayList;


// Definition for singly-linked list.
class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }
}

class Solution {
    public ListNode deleteDuplicates(ListNode head) {
        ListNode current = head;
        while (current != null && current.next != null) {
            if (current.val == current.next.val) {
                current.next = current.next.next;
            } else {
                current = current.next;
            }
        }
        return head;
    }

    public static void main(String[] args) {
        ListNode head = new ListNode(0);
        int[] vals = { 1, 1, 2, 3, 3 };
        for (int l = vals.length - 1; l > 0; l--) {
            ListNode node = new ListNode(vals[l]);
            node.next = head;
            head = node;
        }

        Solution s = new Solution();
        ListNode newHead = s.deleteDuplicates(head);

        while (newHead.next != null) {
            System.out.println(newHead.val);
            newHead = newHead.next;
        }
    }

}
