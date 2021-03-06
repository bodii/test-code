package chapter15;

import java.util.ArrayList;
import java.util.Collection;
import java.util.EmptyStackException;
import java.util.List;

public class Stack2<E> {
    private List<E> elements;
    private int size = 0;
    private static final int DEFAULT_INITIAL_CAPACITY = 16;

    public Stack2() {
        elements = new ArrayList<>(DEFAULT_INITIAL_CAPACITY);
    }

    public void push(E e) {
        elements.add(size++, e);
    }

    public void pushAll(Iterable<? extends E> src) {
        for (E e : src)
            push(e);
    }

    public E pop() {
        if (size == 0)
            throw new EmptyStackException();

        return elements.remove(--size);
    }

    public void popAll(Collection<? super E> dst) {
        while (!isEmpty())
            dst.add(pop());
    }

    public boolean isEmpty() {
        return elements.isEmpty();
    }
}