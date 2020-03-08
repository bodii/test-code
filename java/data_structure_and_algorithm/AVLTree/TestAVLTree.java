package avl_tree;

import avl_tree.AVLTree;

public class TestAVLTree {
    public static void main(String[] args) {
        AVLTree<Integer, Integer> avlt = new AVLTree<>();
        
        int[] nums = {5, 3, 6, 8, 4, 2};
        for (int i = 0; i < nums.length; i ++)
        	avlt.add(i, nums[i]);

        System.out.println("size: " + avlt.getSize());
        System.out.println("isBlanceTree: " + avlt.isBalanced());

    }
}