package chapter10;

public class CountupThreadMain2 {
    public static void main(String[] args) {
        System.out.println("main: BEGIN");

        try {
            // 启动线程
            CountupThread2 countup = new CountupThread2();
            countup.start();

            // 隔一段时间
            Thread.sleep(10000);

            // 线程终止请求
            System.out.println("main: shutdownRequest");
            countup.shutdownRequest();

            // 等待线程终止
            System.out.println("main: join");
            countup.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        } finally {
            System.out.println("main: END");
        }
    }
}