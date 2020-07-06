package chapter12;

public class ActiveObject2Main {
    public static void main(String[] args) {
        ActiveObject2 activeObject = ActiveObjectFactory2.createActiveObject();
        try {
            new MakerClientThread2("Alice", activeObject).start();
            new MakerClientThread2("Bobby", activeObject).start();
            new DisplayClientThread2("Chris", activeObject).start();
            Thread.sleep(5000);
        } catch (InterruptedException e) {
        } finally {
            System.out.println("*** shutdown ***");
            activeObject.shutdown();
        }

    }
}