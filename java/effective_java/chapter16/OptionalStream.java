import java.util.Comparator;

public class OptionalStream {
    public static <E extends Comparable<E>> 
        Optional<E> max(Collection<E> c) {
            return c.stream().max(Comparator.naturalOrder());
    }
}