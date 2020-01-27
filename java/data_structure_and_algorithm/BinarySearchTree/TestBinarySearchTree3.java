package binary_search_tree;

import binary_search_tree.BinarySearchTree;

/**
 * 测试二分搜索树类3
 * 删除节点中的任意元素的节点
 */
public class TestBinarySearchTree3 {
    public static void main(String[] args) {
        BinarySearchTree<Integer> bst = new BinarySearchTree<>();
        
        int[] nums = {5, 3, 6, 8, 4, 2};
        for (int num : nums)
            bst.add(num);

        System.out.println("原始二分搜索树: ");
        System.out.println(bst);

        int removeElement = 3;
        System.out.printf("%n删除元素为%d的节点: %n", removeElement);
        bst.remove(removeElement);

        System.out.println("删除元素后的二分搜索树: ");
        System.out.println(bst);

    }
}