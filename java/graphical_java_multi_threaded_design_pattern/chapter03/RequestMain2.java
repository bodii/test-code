package chapter03;

public class RequestMain2 {
    public static void main(String[] args) {
        // 启动线程
        RequestQueue requestQueue = new RequestQueue();
        Thread alice = new ClientThread(requestQueue, "Alice", 314159L);
        Thread bobby = new ServerThread(requestQueue, "Bobby", 265358L);
        alice.start();
        bobby.start();

        // 等待约10秒
        try {
            Thread.sleep(10000);
        } catch (InterruptedException e) {

        }

        // 调用interrupt方法 需要在线程对象的run方法中设置 通过isInterrupted判断是否设置的中断
        System.out.println("***** calling interrupt ****");
        alice.interrupt();
        bobby.interrupt();
    }
}