package chapter02;

public class Synch {
    private final String name = "Synch";

    @Override
    public synchronized String toString() {
        return "[ " + name + " ]";
    }
}