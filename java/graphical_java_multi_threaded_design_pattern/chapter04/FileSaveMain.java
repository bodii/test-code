package chapter04;

public class FileSaveMain {
    public static void main(String[] args) {
        Data data = new Data("data.txt", "(empty)");
        new ChangerThread("ChangerThread", data).start();
        new SaverThread("ChangerThread", data).start();
    }
}