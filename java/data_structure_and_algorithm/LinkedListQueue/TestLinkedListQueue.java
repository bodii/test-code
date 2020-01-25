package linked_list_queue;

import linked_list_queue.LinkedListQueue;

public class TestLinkedListQueue {
    public static void main(String[] args) {
        LinkedListQueue<Integer> lq = new LinkedListQueue<>();

        for (int i = 0; i < 10; i ++) {
            lq.enqueue(i);
            System.out.println(lq);

            if (i %3 == 2) {
                System.out.println("取出元素：" + lq.dequeue());
            }
        }

        System.out.println("\n获取头部元素的值：" + lq.getFront());
    }
}