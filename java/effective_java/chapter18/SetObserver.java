package chapter18;

import java.util.Set;

@FunctionalInterface public interface SetObserver<E> {
    void added(Set<E> set, E element);
}