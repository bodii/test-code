package binary_search_tree;

import java.util.ArrayList;
import java.util.Random;

import binary_search_tree.BinarySearchTree;

/**
 * 二分搜索树测试类2
 * 用于测试删除最小和最大元素
 */
public class TestBinarySearchTree2 {
    public static void main(String[] args) {
        BinarySearchTree<Integer> bst = new BinarySearchTree<>();

        Random random = new Random();

        ArrayList<Integer> array = new ArrayList<>();

        for(int i = 0; i < 1000; i ++) 
            bst.add(random.nextInt(10000));

        while (!bst.isEmpty()) 
            array.add(bst.removeMin()/* removeMax 从大到小 */); 

        System.out.println(array);

    }
}