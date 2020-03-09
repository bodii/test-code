package avl_tree;

import avl_tree.AVLTree;

public class TestAVLTree {
    public static void main(String[] args) {
        AVLTree<Integer, Integer> avlt = new AVLTree<>();
        
        int[] nums = {5, 3, 6, 8, 4, 2};
        for (int i = 0; i < nums.length; i ++)
        	avlt.add(nums[i], i);

        System.out.println("size: " + avlt.getSize());
        System.out.println("isBlanceTree: " + avlt.isBalanced());
        System.out.println("isBinarySearchTree: " + avlt.isBinarySearchTree());

        System.out.println("get min value: " + avlt.min());
        System.out.println("get max value: " + avlt.max());

    }
}