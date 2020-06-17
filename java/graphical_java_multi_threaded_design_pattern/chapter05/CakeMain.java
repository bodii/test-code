package chapter05;

public class CakeMain {
    public static void main(String[] args) {
        Table table = new Table(3);
        new MakerThread("MakerThread-1", table, 31415).start();
        new MakerThread("MakerThread-2", table, 92653).start();
        new MakerThread("MakerThread-3", table, 56754).start();
        new EaterThread("EaterThread-1", table, 32232).start();
        new EaterThread("EaterThread-2", table, 64333).start();
        new EaterThread("EaterThread-3", table, 34554).start();
    }
}
