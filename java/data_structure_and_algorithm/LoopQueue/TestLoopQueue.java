package loop_queue;

import loop_queue.LoopQueue;

public class TestLoopQueue {
    public static void main(String[] args) {
        LoopQueue<Integer> lq = new LoopQueue<>();
        for (int i = 0; i < 10; i ++) {
            lq.enqueue(i);
            System.out.println("loop queue all element : " + lq);

            if (i % 3 == 2) {
                lq.dequeue();
                System.out.println("loop queue front element: " + lq.getFront());
            }
        }
    }
}