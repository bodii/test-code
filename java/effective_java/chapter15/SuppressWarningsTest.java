package chapter15;

import java.util.Arrays;

public class SuppressWarningsTest<T> {
    private int size;
    private T[] elements;

    public <T> T[] toArray(T[] a) {
        if (a.length < size) {
            // 每当使用SuppressWarnings("unchecked")注解时，都要添加一条注释，说明
            // 为什么这么做是安全的
            @SuppressWarnings("unchecked") T[] result = 
                (T[]) Arrays.copyOf(elements, size, a.getClass());

            return result;
        }
        System.arraycopy(elements, 0, a, 0, size);
        if (a.length > size)
            a[size] = null;

        return a;
    }
}