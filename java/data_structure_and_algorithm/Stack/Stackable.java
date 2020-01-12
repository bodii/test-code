package stack;

/**
 * 栈的接口泛型类
 * 
 * @param <E>
 */
public interface Stackable<E> {
    void push(E e); // 压入元素

    E pop(); // 弹出末尾元素

    E peek(); // 查看末尾元素

    int getSize(); // 获取栈的大小

    boolean isEmpty(); // 判断栈是否为空
}