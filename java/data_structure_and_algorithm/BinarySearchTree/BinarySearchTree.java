package binary_search_tree;

import java.util.Stack;

/**
 * 二分搜索树泛型类(二叉树)
 * 二分搜索树的内部元素要具有可比性
 * 
 * @param <E>
 */
public class BinarySearchTree<E extends Comparable<E>> {

    /**
     * 内部节点类
     */
    private class Node {
        public E e;
        public Node left;
        public Node right;

        public Node(E e, Node left, Node right) {
            this.e = e;
            this.left = left;
            this.right = right;
        }

        public Node(E e, Node left) {
            this(e, left, null);
        }

        public Node(E e) {
            this(e, null, null);
        }

        public Node() {
            this(null, null, null);
        }

        @Override public String toString() {
            return e.toString();
        }
    }

    private int size;   // 树的节点大小
    private Node root; // 根树

    /**
     * 构造函数
     */
    public BinarySearchTree() {
        root = null;
        size = 0; 
    }

    /**
     * 获取树元素节点的长度
     * 
     * @return int 节点数
     */
    public int getSize() {
        return size;
    }

    /**
     * 检测树是否为空
     * 
     * @return boolean 是否为空
     */
    public boolean isEmpty() {
        return size == 0;
    }

    /**
     * 向二分搜索树中添加元素
     * 
     * @param e 元素
     */
    public void add(E e) {
        root = add(e, root);
    }

    /**
     * 向二分搜索树中添加元素，递归算法
     * 
     * @param e 元素
     * @param node 节点
     */
    // private void add(E e, Node node) {
    //     if (node.e == null) {
    //         node.e = e;
    //         size ++;
    //         return;
    //     } else if (e.compareTo(node.e)< 0 && node.left == null ) {
    //         node.left = new Node(e);
    //         size ++;
    //         return;
    //     } else if (e.compareTo(node.e)> 0 && node.right == null) {
    //         node.right = new Node(e);
    //         size ++;
    //         return;
    //     } else if (node.e.equals(e))
    //         return;

    //     if (e.compareTo(node.e)< 0) 
    //         add(e, node.left);
    //     else if (e.compareTo(node.e)> 0)
    //         add(e, node.right);
    // }

    /**
     * 向二分搜索树中添加元素，递归算法
     * 
     * @param e 元素
     * @param node 节点
     * 
     * @return 插入的新点
     */
    private Node add(E e, Node node) {
        if (node == null) {
            size ++;
            return new Node(e);
        }

        if (e.compareTo(node.e)< 0)
            node.left = add(e, node.left);
        else if (e.compareTo(node.e)> 0)
            node.right = add(e, node.right);

        return node;
    }

    /**
     * 查询某个元素是否包含在二叉搜索树中
     * 
     * @param e 元素
     * @return 是否包含元素
     */
    public boolean contains(E e) {
        return contains(e, root);
    }

    /**
     * 查询某个元素是否包含在二叉搜索树中(私有方法)
     * 
     * @param e 元素
     * @param node 节点
     * @return 是否包含元素
     */
    private boolean contains(E e, Node node) {
        if (node == null)
            return false;

        if (e.compareTo(node.e)== 0)
            return true;
        else if (e.compareTo(node.e)< 0)
            return contains(e, node.left);
        else
            return contains(e, node.right);
    }

    /**
     * 二分搜索树的前序遍历
     */
    public void preOrder() {
        preOrder(root);
    }

    /**
     * 二分搜索树的前序遍历(私有方法)
     * @param node
     */
    private void preOrder(Node node) {
        // if (node == null) // 终止条件
        //     return;

        if (node != null) { // 先判断不为空(终止条件)
            System.out.println(node);
            preOrder(node.left);
            preOrder(node.right);
        }
    }

    /**
     * 二分搜索树的中序遍历
     */
    public void inOrder() {
        inOrder(root);
    }

    /**
     *  二分搜索树的中序遍历(私有方法)
     * 
     * @param node
     */
    private void inOrder(Node node) {
        if (node == null)
            return;

        inOrder(node.left);
        System.out.println(node);
        inOrder(node.right);
    }

    /**
     * 二分搜索树的后序遍历
     */
    public void postOrder() {
        postOrder(root);
    }

    /**
     *  二分搜索树的后序遍历(私有方法)
     * 
     * @param node
     */
    private void postOrder(Node node) {
        if (node == null)
            return;

        postOrder(node.left);
        postOrder(node.right);
        System.out.println(node);
    }

    /**
     * 二分搜索树的递归前序遍历
     */
    public void preOrderNR() {
        Stack<Node> stack = new Stack<>();
        stack.push(root);
        while(!stack.isEmpty()) {
            Node current = stack.pop();
            System.out.println(current);

            if(current.right != null)
                stack.push(current.right);
            if (current.left != null)
                stack.push(current.left);
        }
    }

    /**
     * 将当前二分搜索树转换成字符串
     */
    @Override public String toString() {
        StringBuilder result = new StringBuilder();
        result.append(String.format("Binary search tree of size = %d\n", size));
        generateBSTString(root, 0, result);
        return result.toString();
    }

    /**
     * 生成二分搜索树节点与深度字符串
     * 
     * @param node 当前节点
     * @param depth 当前节点距离根节点的深度
     * @param result 生成的字符串
     */
    private void generateBSTString(Node node, int depth, StringBuilder result) {
        if (node == null ) {
            result.append(generateDepthString(depth) + "null\n");
            return;
        }

        result.append(generateDepthString(depth) + node.e + "\n");
        generateBSTString(node.left, depth + 1, result);
        generateBSTString(node.right, depth + 1, result);
    }

    /**
     * 生成二分搜索树深度的字符串
     * e.g. depth=1 "--"; depth=3 "------".
     * 
     * @param depth int 深度值
     * @return 拼接好的深度字符串
     */
    private String generateDepthString(int depth) {
        StringBuilder result = new StringBuilder();

        for (int i = 0; i < depth; i++)
            result.append("--");

        return result.toString();
    }
}