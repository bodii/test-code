package chapter18;

public class GenerateSerialNumber {
    private static long nextSerialNumber = 0;

    public static synchronized long generateSerialNumber() {
        return nextSerialNumber ++;
    }
}