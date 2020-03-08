package binary_search_tree;

import binary_search_tree.AVLTree;

public class TestAVLTree {
    public static void main(String[] args) {
        AVLTree<Integer, Integer> bst = new AVLTree<>();
        
        int[] nums = {5, 3, 6, 8, 4, 2};
        for (int i = 0; i < nums.length; i ++)
            bst.add(i, nums[i]);

        System.out.println("size: " + bst.getSize());

        System.out.println("\n递归前序遍历: ");
        bst.preOrder();

        System.out.println("\n递归中序遍历: ");
        bst.inOrder();

        System.out.println("\n递归后序遍历: ");
        bst.postOrder();

        System.out.println("\n将当前二分搜索树转换成字符串: ");
        System.out.println(bst);

        System.out.println("\n二分搜索树的层序遍历: ");
        bst.levelOrder();

        System.out.println("\n二分搜索树的最小(最左)元素是："  + bst.min());
        System.out.println("\n二分搜索树的最大(最右)元素是："  + bst.max());

        System.out.println("\n删除二分搜索树的最小(最左)："  + bst.removeMin());
        System.out.println("\n删除二分搜索树的最大(最右)元素是："  + bst.removeMax());

        System.out.println("\n删除当前二分搜索树最小和最大元素后: ");
        System.out.println(bst);



    }
}