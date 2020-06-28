package chapter09;

import java.util.concurrent.Callable;

public class Host2 {
    public FutureData2 request(final int count, final char c) {
        System.out.println("    request(" + count + ", " + c + ") BEGIN");

        // 创建FutureData的实例
        //  向构造函数中传递Callable<RealData>)
        FutureData2 future = new FutureData2(
            new Callable<RealData>() {
                public RealData call() {
                    return new RealData(count, c);
                }
            }
        );

        // 启动一个新线程，用于创建RealData的实例
        new Thread(future).start();

        System.out.println("    request(" + count + ", " + c + ") END");
       

        // 返回FutureData的实例
        return future;
    }
}