package pair1;

/**
 * @version 1.1 2019-09-01
 * @author wong
 */
public class PairTest1 {
    public static void main(String[] args) {
        String[] words = { "Mary", "had", "a", "little", "lamb" };
        Pair<String> mm = ArrayAlg.minmax(words);
        System.out.println("min = " + mm.getFirst());
        System.out.println("max = " + mm.getSecond());
    }
}

class ArrayAlg {
    /**
     * Gets the minimum and maximum of an array of strings.
     * @param a an array of strings
     * @return a pair with the main and max value, or null if a is null or empty
     */
    public static Pair<String> minmax(String[] a) {
        if ( a == null || a.length == 0) return null;
        String min = a[0];
        String max = a[1];
        for (int i = 1; i < a.length; i++) {
            if (min.compareTo(a[i]) > 0) min = a[i];
            if (max.compareTo(a[i]) < 0) max = a[i];
        }

        return new Pair<>(min, max);
    }
}

class Pair<T> {
    private T first;
    private T second;

    public Pair() { first = null; second = null; }
    public Pair(T first, T second) {this.first = first; this.second = second; }

    public T getFirst() { return first; };
    public T getSecond() { return second; }

    public void setFirst(T newFirst) { first = newFirst; }
    public void setSecond(T newSecond) { second = newSecond; }
}