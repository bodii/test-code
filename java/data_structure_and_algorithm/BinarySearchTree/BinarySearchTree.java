package binary_search_tree;

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
        add(e, root);
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
    //     } else if (node.e.compareTo(e) < 0 && node.left == null ) {
    //         node.left = new Node(e);
    //         size ++;
    //         return;
    //     } else if (node.e.compareTo(e) > 0 && node.right == null) {
    //         node.right = new Node(e);
    //         size ++;
    //         return;
    //     } else if (node.e.equals(e))
    //         return;

    //     if (node.e.compareTo(e) < 0) 
    //         add(e, node.left);
    //     else if (node.e.compareTo(e) > 0)
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

        if (node.e.compareTo(e) < 0)
            node.left = add(e, node.left);
        else if (node.e.compareTo(e) > 0)
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

        if (node.e.compareTo(e) == 0)
            return true;
        else if (node.e.compareTo(e) < 0)
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
}