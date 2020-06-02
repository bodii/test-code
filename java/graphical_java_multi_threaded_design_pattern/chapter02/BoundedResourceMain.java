package chapter02;

public class BoundedResourceMain {
    public static void main(String[] args) {
        BoundedResource resource = new BoundedResource(3);

        // 10个线程使用资源
        for (int i = 0; i < 10; i ++) {
            new UseThreadResource(resource).start();
        }
    }
}