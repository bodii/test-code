package chapter14;

import java.util.Iterator;
import java.util.Objects;
import java.util.function.Predicate;

public interface Predicatea<E> {
    default boolean removeIf(Predicate<? super E> filter) {
        Objects.requireNonNull(filter);
        boolean result = false;
        for (Iterator<E> it = iterator(); it.hasNext();) {
            if (filter.test(it.next())) {
                it.remove();
                result = true;
            }
        }

        return result;
    }

    Iterator<E> iterator();
}