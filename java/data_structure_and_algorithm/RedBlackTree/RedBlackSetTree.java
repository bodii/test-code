package red_black_tree;

import red_black_tree.RedBlackTree;
import red_black_tree.Setable;

public class RedBlackSetTree<E extends Comparable<E>> implements  Setable<E> {
    private RedBlackTree<E, Object> root;

    public RedBlackSetTree() {
        root = new RedBlackTree<>();
    }

    @Override public void add(E e) {
        root.add(e, null);
    }

    @Override public void remove(E e) {
        root.remove(e);
    }

    @Override public boolean contains(E e) {
        return root.contains(e);
    }

    @Override public int getSize() {
        return root.getSize();
    }

    @Override public boolean isEmpty() {
        return root.isEmpty();
    }
}