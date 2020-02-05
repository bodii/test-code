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
        // 查询测试
        System.out.println("查询0-2区间的值:" + segment.query(0, 2));

        // 更新测试
        segment.set(1, 15);
        System.out.println("将索引为1的值更新为15后: \n" + segment);
        System.out.println("查询0-3区间的值:" + segment.query(0, 3));
    }
}