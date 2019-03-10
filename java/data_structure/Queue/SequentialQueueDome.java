/**
 * SequentialQueueDome 顺序队列测试类
 */

public class SequentialQueueDome {
    public static void main(String[] args) {
        SequentialQueue sq = new SequentialQueue(10);
        sq.push('a');
        sq.push('b');
        sq.push('c');
        System.out.println("size: " + sq.size());
        System.out.println("pop: " + sq.pop());
        System.out.println("size: " + sq.size());
        System.out.println("peek: " + sq.peek());
        System.out.println("size: " + sq.size());
        System.out.println("pop: " + sq.pop());
        System.out.println("pop: " + sq.pop());
        sq.push("d");
        System.out.println("peek: " + sq.peek());
        System.out.println("size: " + sq.size());
        // sq.push("e");
        // sq.push("f");
        // sq.push("j");
        // sq.push("h");
        // sq.push("g");
        // sq.push("k");
        // sq.push("l");
        // sq.push("m");
        // sq.push("n");
        // System.out.println("size: " + sq.size());
        // sq.push("o");
        // sq.pop();
        // sq.pop();
    }
}