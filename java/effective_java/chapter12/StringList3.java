package chapter12;

import java.io.IOException;
import java.io.ObjectOutputStream;
import java.io.Serializable;

/**
 * transient修饰符表明这个实例域将从一个类的默认序列化形式中省略掉
 */
public final class StringList3 implements Serializable {
    private transient int size = 0;
    private transient Entry head = null;
    // 不管你选择哪种序列化形式，都要为自已编写的每个可序列化的类声明一个
    // 显式的序列版本UID.
    private static final long serialVersionUID = randomLongValue;

    private static class Entry {
        String data;
        Entry next;
        Entry previous;
    }

    public final void add(String s) {}

    private void writeObject(ObjectOutputStream s) throws IOException {
        s.defaultWriteObject();
        s.writeInt(size);

        for (Entry e = head; e != null; e = e.next)
            s.wirteObject(e.data);
    }

    private void readObject(ObjectOutputStream s) 
        throws IOException, ClassNotFoundException {
            s.defaultReadObject();
            int numElements = s.readInt();

            for (int i = 0; i < numElements; i++)
                add((String) s.readObject());
    }


}