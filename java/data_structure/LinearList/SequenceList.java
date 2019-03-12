/**
 * 线性表顺序存储操作实现
 */
public class SequenceList implements LinearList {
    private Object[] sList;
    private int size;

    public SequenceList(int length) {
        if (length == 0) {
            this.sList = new Object[10];
        }
        this.sList = new Object[length];
    }

    public SequenceList() {
        this(10);
    }

    public boolean isEmpty() {
        if (size <= 0) {
            return true;
        } else {
            return false;
        }
    }

    public int size() {
        return size;
    }

    public Object get(int index) {
        checkIndex(index);

        return (Object) sList[index];
    }

    public void set(int index, Object element) {
        checkIndex(index);

        sList[index] = element;
    }

    public boolean add(Object element) {
        return add(size, element);
    }

    public boolean add(int index, Object element) {
        checkIndex(index);
        checkSize();

        for (int i = size - 1; i >= index; i--) {
            sList[i + 1] = sList[i];
        }
        sList[index] = element;
        size++;

        return true;
    }

    public Object remove(int index) {
        checkIndex(index);

        Object old = (Object) sList[index];
        for (int i = index; i < size - 1; i++) {
            sList[i] = sList[i + 1];
        }
        sList[--size] = null;
        return old;
    }

    public void clear() {
        for (int i = 0; i < size; i++) {
            sList[i] = null;
        }

        size = 0;
    }

    private void checkIndex(int index) {
        if (index > size || index < 0) {
            throw new IndexOutOfBoundsException("index:  " + index + "  size:" + size);
        }
    }

    private void checkSize() {
        if (size == sList.length) {
            Object[] temp = sList;
            this.sList = new Object[temp.length * 2];
            for (int i = 0; i < temp.length; i++) {
                this.sList[i] = temp[i];
            }
        }
    }
}