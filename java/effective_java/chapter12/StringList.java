package chapter12;

import java.io.Serializable;

/**
 * 当一个对象的物理表示法与它的逻辑数据内容有实质性的区别时，使用默认序列化形式会有以下
 * 4个缺点：
 * 1. 它使这个类的导出API永远地束缚在该类的内部表示法上。
 * 2. 它会消耗过多的空间。
 * 3. 它会消耗过多的时间。
 * 4. 它会引起栈溢出。
 */
public final class StringList implements Serializable {
    private int size = 0;
    private Entry head = null;

    private static class Entry implements Serializable {
        String data;
        Entry next;
        Entry previous;
    }
}