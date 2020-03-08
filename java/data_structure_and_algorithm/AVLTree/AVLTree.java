package binary_search_tree;

import java.util.Queue;
import java.util.Stack;
import java.util.LinkedList;

/**
 * 平衡二分搜索树泛型类(二叉树)
 * 平衡二分搜索树的内部元素要具有可比性
 * 维护一个平衡因子，其它左右两边最大和最小的高度差不大于1
 * 
 * @param <E>
 */
public class AVLTree<K extends Comparable<K>, V> {

    /**
     * 内部节点类
     */
    private class Node {
        public K key;
        public V value;
        public Node left;
        public Node right;

        public Node (K key, V value) {
            this.key = key;
            this.value = value;
            left = null;
            right = null;
            height = 1; // 节点的高度
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
     * 获取当前节点的高度
     * 
     * @param Node node 要查询的节点
     * @return 高度
     */
    private int getHeight(Node node) {
        if (node == null)
            return 0;
        return node.height;
    }

    /**
     * 向二分搜索树中添加元素
     * 
     * @param K 键
     * @param V 值
     */
    public void add(K key, V value) {
        root = add(root, key, value);
    }

    /**
     * 向二分搜索树中添加元素，递归算法
     * 
     * @param node 节点
     * @param K 键
     * @param V 值
     */
    private Node add(Node node, K key, V value) {
        if (node == null) {
            size ++;
            return new Node(key, value);
        } 
        
        if (key.compareTo(node.key) < 0)
            node.left =  add(node.left, key, value);
        else if (key.compareTo(node.key) > 0)
            node.right = add(node.right, key, value);
        else 
            node.value = value;

        // 更新当前height 的值
        node.height = 1 + Math.max(getHeight(node.left), getHeight(node.right));

        // 计算当前节点的平衡因子
        int balanceFactor = getBalanceFactor(node);

        // 说明i当前最多节点的平衡因子与最小节点的平衡因子相差大于1
        if (Math.abs(balanceFactor) > 1) 
            System.out.println("unbalanceFactor: " + balanceFactor);

        return node;
    }

    /**
     * 获取以Node为根节点的二分搜索树中，key所在的节点
     */
    private Node getNode(Node node, K key) {
        if (node == null)
            return null;
            
        if (key.compareTo(node.key) < 0)
            return getNode(node.left, key);
        else if (key.compareTo(node.key) > 0)
            return getNode(node.right, key);
        // key.compareTo(node.key) == 0
        return node;
    }

    /**
     * 设置指定key对应的值
     * 
     * @param key 
     * @param value
     */
    public void set(K key, V value) {
        Node node = getNode(root, key);
        if (node == null)
            throw new IllegalArgumentException(key + " doesn't exists!");

        node.value = value;
    }

    /**
     * 获取指定key对应的值
     * 
     * @param key
     * @return 值
     */
    public V get(K key) {
        Node node = getNode(root, key);
        return node == null ? null : node.value;
    }

    /**
     * 获取节点的平衡因子
     * 
     * @param node
     * @return 返回当前节点的平衡因子
     */
    private int getBalanceFactor(Node node) {
        if (node == null)
            return 0;

        return getHeight(node.left) - getHeight(node.right);
    }

    /**
     * 查询当前二叉树是不是平衡二叉树
     * 
     * @return 是不是平衡二叉树
     */
    public boolean isBalanced() {
        return isBalanced(root);
    }

    /**
     * 递归查询当前二叉树的所有节点是否满足平衡二叉树
     * 
     * @param Node 当前节点
     * @return 是否当前节点和后面的节点是否满足平衡
     */
    private boolean isBalanced(Node node) {
        if (node == null)
            return true;

        int balanceFactor = getBalanceFactor(root);
        if (Math.abs(balanceFactor) > 1)
            return false;

        return isBalanced(node.left) && isBalanced(node.right);
    }

    /**
     * 查询某个key键是否包含在二叉搜索树中
     * 
     * @param K 元素
     * @return 是否包含
     */
    public boolean contains(K key) {
        return getNode(root, key) != null;
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
     * 二分搜索树的层序遍历(广度优先)
     */
    public void levelOrder() {
        Queue<Node> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            Node current = queue.remove();
            System.out.println(current);

            if (current.left != null) 
                queue.add(current.left);
            if (current.right != null)
                queue.add(current.right);
        }
    }

    /**
     * 获取二分搜索树中最小的key(最左key)
     * 
     * @return K 最小key
     */
    public K min() {
        nonEmpty();

        return minNode(root).key;
    }

    /**
     * 获取二分搜索树中最小元素的节点(递归遍历)
     * 
     * @param node 当前节点
     * @return 遍历的当前左节点
     */
    private Node minNode(Node node) {
        if (node.left == null) 
            return node;

        return minNode(node.left);
    }

    /**
     * 获取二分搜索树中最大的key(最右元素)
     * 
     * @return E 最大key
     */
    public E max() {
        nonEmpty();

        return maxNode(root).key;
    }

    /**
     *  获取二分搜索树中最小元素的节点(递归遍历)
     * 
     * @param node 当前节点
     * @return  遍历的当前右节点
     */
    private Node maxNode(Node node) {
        if (node.right == null)
            return node;

        return maxNode(node.right);
    }

    /**
     * 删除二分搜索树中的最小元素
     * 
     * @return 被删除的元素
     */
    public E removeMin() {
        E min = min();
        root = removeMin(root);

        return min;
    }

    /**
     * 删除二分搜索树中的最左(最小)元素的节点
     * 
     * @param node 当前节点
     * @return 下一个节点
     */
    private Node removeMin(Node node) {
        if (node.left == null) {
            Node rightNode = node.right;
            node.right = null;
            size --;
            return rightNode;
        }

        node.left = removeMin(node.left);

        return node;
    }

    /**
     * 删除二分搜索树中最大的元素
     * 
     * @return 最大的元素
     */
    public E removeMax() {
        E max = max();
        root = removeMax(root);

        return max;
    }

    /**
     * 删除二分搜索树中的最右(最大)元素的节点
     * 
     * @param node 当前节点
     * @return 下一个节点
     */
    private Node removeMax(Node node) {
        if (node.right == null) {
            Node leftNode = node.left;
            node.right = null;
            size --;

            return leftNode;
        }

        node.right = removeMax(node.right);

        return node;
    }

    /**
     * 删除二分搜索树的某个key对应的节点
     * 
     * @param K 对应的key
     * 
     */
    public void remove(K key) {
        remove(root, key);
    }

    /**
     * 根据递归遍历查找并删除二分搜索树的指定key的节点
     * @param node
     * @param K 指定的key
     * @return
     */
    private Node remove(Node node, K key) {
        if (node == null)
            return null;

        if (key.compareTo(node.key) < 0) {
            node.left = remove(node.left, key);
            return node;
        } else if (key.compareTo(node.key) > 0) {
            node.right = remove(node.right, e);
            return node;
        } else {
            // 如果左子树为空，则将右子树返回
            if (node.left == null) {
                Node rightNode = node.right;
                node.right = null;
                size --;
                return rightNode;
            } 

            // 如果右子树为空，则将左子树返回
            if (node.right == null) {
                Node leftNode = node.left;
                node.left = null;
                size --;
                return leftNode;
            }

            // 如果左右子树均不为空
            // 将右子树的最小元素节点删除并作为新节点，替换掉要删除的节点
            Node newNode = minNode(node.right); // 获取当前节点右子树的最小元素节点(或最大值节点也行，也叫前驱)
            newNode.right = removeMin(node.right); // 将当前节点右子树的最小元素节点删除后返回的节点树，作为新节点的右子树
            newNode.left = node.left; // 将当前节点的左子树，赋给新节点
            node.left = node.right = null;  // 将当前节点的左右子树销毁
            return newNode; // 返回新的节点树
        }

    }

    /**
     * 检测当前二分搜索树是否不为空，则否抛出异常
     * 
     * @throws IllegalArgumentException
     */
    private void nonEmpty() {
        if (size ==0)
            throw new IllegalArgumentException("Binary search tree is empty.");
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

        result.append(generateDepthString(depth) + node.key + "\n");
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