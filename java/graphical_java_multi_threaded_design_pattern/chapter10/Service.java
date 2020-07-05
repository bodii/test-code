package chapter10;

public class Service {
    private volatile static boolean isShutdown = false;
    // 开始运行服务
    public static void service() {
        System.out.print("service");
        for (int i = 0; i < 50; i++) {
            if (!isShutdown) {
                System.out.print(".");
                try {
                    Thread.sleep(100);
                } catch (InterruptedException e) {

                }
            }
        }
        System.out.println("done.");
    }

    // 停止服务
    public static void cancel() {
        try {
            Thread.currentThread().join();
        } catch (InterruptedException e) {

        }
    }
}