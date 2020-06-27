package chapter09;

public class Host {
    public Data request(final int count, final char c) {
        System.out.println("    request(" + count + ", " + c + ") BEGIN");

        // 创建FutureData的实例
        final FutureData future = new FutureData();

        // 启动一个新线程，用于创建RealData的实例
        new Thread() {
            @Override
            public void run () {
                RealData realdata = new RealData(count, c);
                future.setRealData(realdata);
            }
        }.start();

        // 返回FutureData的实例
        return future;
    }
}