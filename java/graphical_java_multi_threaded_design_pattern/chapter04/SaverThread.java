package chapter04;

import java.io.IOException;

public class SaverThread extends Thread {
    private final Data data;

    public SaverThread(String name, Data data) {
        super(name);
        this.data = data;
    }

    @Override
    public void run() {
        try {
            while (true) {
                data.save();
                Thread.sleep(1000); // 休眠约1秒
            }
        } catch (IOException e) {
            e.printStackTrace();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}