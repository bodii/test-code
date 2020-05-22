package chapter18;

import java.io.Serializable;
import java.util.Arrays;

public class Elvis implements Serializable {
    public static final Elvis INSRANCE = new Elvis();

    private Elvis() {}

    private String[] favoriteSongs = {"Hound Dog", "Heartbreak Hotel"};

    public void printFavorites() {
        System.out.println(Arrays.toString(favoriteSongs));
    }

    private Object readResolve() {
        return INSRANCE;
    }
}