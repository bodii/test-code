package stack;

/**
 * Stack class
 * 栈的泛型类
 * 
 * @param <E>
 */
public class Stack<E> implements Stackable<E> {
    Array<E> array;

    /**
     * 栈的构造函数
     */
    public Stack() {
        array = new Array<>();
    }

    /**
     * 获取栈的大小
     * @return int  size
     */
    @Override public int getSize() {
        return array.getSize();
    }

    /**
     * 判断当前栈是否为空
     * @return boolean 是否为空
     */
    @Override public boolean isEmpty() {
        return array.isEmpty();
    }

    /**
     * 获取当前栈的容量
     * 
     * @return int 容量
     */
    public int getCapacity() {
        return array.getCapacity();
    }

    /**
     * 在栈尾压入一个元素
     * @param E 元素
     */
    @Override public void push(E e) {
        array.addLast(e);
    }

    /**
     * 从栈尾弹出一个元素
     * @return E 元素
     */
    @Override public E pop() {
        return array.removeLast();
    }

    /**
     * 查看栈尾元素的值
     * @return E 元素
     */
    @Override public E peek() {
        return array.get(getSize() - 1);
    }

    /**
     * 将当前栈转换成字符串
     * @return String 返回生成后的字符串
     */
    @Override public String toString() {
        StringBuilder result = new StringBuilder();
        result.append("Stack: [");
        for (int i = 0; i < getSize(); i++) 
            result.append(array.get(i) + ", ");
        result.replace(result.length() -2, result.length(), "] top");

        return result.toString();
    }
}