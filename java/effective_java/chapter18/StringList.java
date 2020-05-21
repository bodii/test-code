package chapter18;

import java.io.IOError;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.io.Serializable;

public class StringList implements Serializable {
    private transient int size = 0;
    private transient Entry head = null;

    private static class Entry {
        String data;
        Entry next;
        Entry previous;
    }

    public final void add(String s) { ... }

    /**
     * Serialize this {@code StringList} instance.
     * 
     * @serialData The size of the list(the number of strings it
     * contains) is emitted ({@code int}), followed by all of 
     * its elements (each a {@code String}), in the proper sequence.
     * @param s
     * @throws IOException
     */
    private void writeObject(ObjectOutputStream s) throws IOException {
        s.defaultWriteObject();
        s.writeInt(size);

        for (Entry e = head; e != null; e = e.next)
            s.writeObject(e.data);
    }

    private void readObject(ObjectInputStream s) 
            throws IOException, ClassNotFoundException {
        s.defaultReadObject();
        int numElements = s.readInt();

        for (int i = 0; i < numElements; i ++)
            add((String) s.readObject());
    }
    
}