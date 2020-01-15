package queue;

import queue.Queue;

public class TestQueue {
    public static void main(String[] args) {
        Queue<Integer> queue = new Queue<>();

        System.out.println("在队列中添加一个元素: 3");
        queue.enqueue(3);
        System.out.println("在队列中添加一个元素: 5");
        queue.enqueue(5);
        System.out.println("在队列中添加一个元素: 7");
        queue.enqueue(7);
        System.out.println("在队列中去掉首部一个元素" + queue.dequeue());
        System.out.println("在队列中取出一个元素" + queue.getFront());
        System.out.println("打印当前队列的元素" + queue);
    }
}