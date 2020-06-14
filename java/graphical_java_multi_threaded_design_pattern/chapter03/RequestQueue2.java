package chapter03;

import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;

public class RequestQueue2 {
    private final BlockingQueue<Request> queue = new LinkedBlockingQueue<>();

    public Request getRequest() {
        Request req = null;
        try {
            req = queue.take();
        } catch (InterruptedException e) {

        }

        return req;
    }

    public void putRqeust(Request request) {
        try {
            queue.put(request);
        } catch (InterruptedException e) {
            
        }
    }
}