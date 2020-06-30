package chapter10;

public class CountupThreadMain {
    public static void main(String[] args) {
        System.out.println("main: BEGIN");

        try {
            // 启动线程
            CountupThread countup = new CountupThread();
            countup.start();

            // 稍微间隔一段时间 
            Thread.sleep(10000);

            // 线程的终止请求
            System.out.println("main: shutdownRequest");
            countup.shutdownRequest();

            System.out.println("main: join");

            // 等待线程终止
            countup.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println("main: END");
    }
}