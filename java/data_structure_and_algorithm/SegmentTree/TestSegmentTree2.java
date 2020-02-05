package segment_tree;

import segment_tree.SegmentTree;
import segment_tree.Merger;

/**
 * 线段树测试类2
 */
public class TestSegmentTree2 {
    public static void main(String[] args) {
        Integer[] arr = {-2, 0, 3, -5, 2, -1};
        SegmentTree<Integer> segment = new SegmentTree<>(arr, (a, b) -> a + b);

        System.out.println(segment);
    }
}