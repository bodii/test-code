package red_black_tree;

import red_black_tree.RedBlackTree;

public class TestRedBlackTree {
    public static void main(String[] args) {
        RedBlackTree<Integer, Integer> rbt = new RedBlackTree<>();
        
        int[] nums = {5, 3, 6, 8, 4, 2};
        for (int i = 0; i < nums.length; i ++)
            rbt.add(nums[i], i);

        System.out.println("contains 8: " + rbt.contains(8));
        System.out.println(rbt);
    }
}