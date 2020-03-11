package chapter14;

import java.util.HashSet;
import java.util.Collection;

public class InstrumentedHashSet<E> extends HashSet<E> {
    /**
     *
     */
    private static final long serialVersionUID = 1L;
    private int addCount = 0;

    public InstrumentedHashSet() {

    }

    public InstrumentedHashSet(int initCap, float loadFactor) {
        super(initCap, loadFactor);
    }

    @Override public boolean add(E e) {
        addCount ++;
        return super.add(e);
    }

    @Override public boolean addAll(Collection<? extends E> c) {
        // addCount += c.size();
        return super.addAll(c);
    }

    public int getAddCount() {
        return addCount;
    }
}