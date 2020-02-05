package segment_tree;

import segment_tree.SegmentTree;
import segment_tree.Merger;

/**
 * 线段树测试类
 */
public class TestSegmentTree {
    public static void main(String[] args) {
        Integer[] arr = {-2, 0, 3, -5, 2, -1};
        SegmentTree<Integer> segment = new SegmentTree<>(arr, new Merger<Integer>() {
            @Override
            public Integer merger(Integer a, Integer b) {
                return a + b;
            }
        });

        System.out.println(segment);
    }
}