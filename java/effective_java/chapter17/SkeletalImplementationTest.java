package chapter17;

import java.util.ListIterator;
import java.util.NoSuchElementException;

/**
 * 异常转译：
 *  更高层的实现应该捕获低层的异常，同时抛出可以按照高层抽象进行解释的异常。
 */
public class SkeletalImplementationTest {
    /**
     * Returns the element at the specified position in this list.
     * @throws IndexOutOfBoundsException if the index is out of range
     *          ({@code index < 0 || index >= size()}).
     * @param index
     * @return
     */
    public E get(int index) {
        ListIterator<E> i = ListIterator(index);
        try {
            return i.next();
        } catch (NoSuchElementException e) {
            throw new IndexOutOfBoundsException("Index: " + index);
        }
    }
}