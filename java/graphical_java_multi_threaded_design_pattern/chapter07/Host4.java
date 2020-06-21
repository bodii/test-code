package chapter07;

import java.util.concurrent.Executor;

public class Host4 {
    private final Helper helper = new Helper();
    private final Executor executor;

    public Host4 (Executor exec) {
        this.executor = exec;
    }

    public void request(final int count, final char c) {
        System.out.println("    request(" + count + ", " + c + ") BEGIN");

        executor.execute(
            new Runnable(){
                @Override
                public void run() {
                   helper.handle(count, c);
                }
            }
        );

        System.out.println("    request(" + count + ", " + c + ") END");
    }
}