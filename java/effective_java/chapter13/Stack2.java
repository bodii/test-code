package chapter13;

import java.util.Arrays;
import java.util.EmptyStackException;

public class Stack2 {
    private Object[] elements;
    private int size = 0;
    private static final int DEFAULT_INTIAL_CAPACITY = 16;

    public Stack2() {
        this.elements = new Object[DEFAULT_INTIAL_CAPACITY];
    }

    public void push(Object e) {
        ensureCapacity();
        elements[++size] = e;
    }

    public Object pop() {
        if (size == 0)
            throw new EmptyStackException();
        Object result = elements[--size];
        elements[size] = null;

        return result;
    }

    private void ensureCapacity() {
        if (elements.length == size) 
            elements = Arrays.copyOf(elements, size * 2 + 1);
    }

    // Clone method for class with references to mutable state
    @Override public Stack2 clone() {
        try {
            Stack2 result = (Stack2) super.clone();
            result.elements = elements.clone();
            return result;
        } catch (CloneNotSupportedException e) {
            throw new AssertionError();
        }
    }
}