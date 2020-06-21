package chapter07;

public class HostT {
    private final Helper helper = new Helper();

    public void request(final int count, final char c){
        System.out.println("    reqeust(" + count + ", " + c + ") BEGIN");
        /**
         * run()是Thread的一个普通方法，当调用run方法时还是在主线程内执行。
         * 
         * start()方法可以启动线程，当调用start()方法时，线程处理就绪状态（可运行），
         * 它包含要执行的这个线程的内容，Run方法运行结束，此线程随即终止。
         */
        new Thread(){
            public void run() {
                helper.handle(count, c);
            }
        }.run();

        System.out.println("    reqeust(" + count + ", " + c + ") END");
    }
}