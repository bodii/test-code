package red_black_tree;

import red_black_tree.RedBlackTree;
import red_black_tree.Mapable;

public class RedBlackSetTree<K extends Comparable<K>, V> implements  Mapable<K, V> {
    private RedBlackTree<K, V> root;

    public RedBlackSetTree() {
        root = new RedBlackTree<>();
    }

    @Override public void add(K k, V v) {
        root.add(k, v);
    }

    @Override public void remove(K k) {
        root.remove(k);
    }

    @Override public boolean contains(K k) {
        return root.contains(k);
    }

    @Override public int getSize() {
        return root.getSize();
    }

    @Override public boolean isEmpty() {
        return root.isEmpty();
    }
}