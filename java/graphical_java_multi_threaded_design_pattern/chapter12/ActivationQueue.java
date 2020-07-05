package chapter12;

public class ActivationQueue {
    private static final int MAX_METHOD_REQUEST = 100;
    private final MethodRequest<Object>[] requestQueue;
    private int tail; // 下次putRequest的位置
    private int head; // 下次getRequest的位置
    private int count; // MethodRequest的数量

    public ActivationQueue() {
        this.requestQueue = new MethodRequest[MAX_METHOD_REQUEST];
        this.head = 0;
        this.tail = 0;
        this.count = 0;
    }

    public synchronized void putRequest(MethodRequest<Object> request) {
        while (count >= requestQueue.length) {
            try {
                wait();
            } catch (InterruptedException e) {

            }
        }

        requestQueue[tail] = request;
        tail = (tail + 1) % requestQueue.length;
        count ++;
        notifyAll();
    }

    public synchronized MethodRequest<Object> takeRequest() {
        while (count <= 0) {
            try {
                wait();
            } catch (InterruptedException e) {

            }
        }

        MethodRequest request = requestQueue[head];
        head = (head + 1) % requestQueue.length;
        count --;
        notifyAll();

        return request;
    }
}