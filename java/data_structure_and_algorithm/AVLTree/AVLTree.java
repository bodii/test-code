package avl_tree;

import java.util.Queue;
import java.util.ArrayList;

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
        public int height;

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
    public AVLTree() {
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
        
        // 平衡维护
        // LL
        if (balanceFactor > 1 && getBalanceFactor(node.left) >= 0)
        	return rightRotate(node);

        // RR
        if (balanceFactor < -1 && getBalanceFactor(node.right) <= 0)
            return leftRotate(node);

        // LR
        if (balanceFactor > 1 && getBalanceFactor(node.left) < 0) {
            node.left = leftRotate(node.left);
            return rightRotate(node);
        }

        // RL
        if (balanceFactor < -1 && getBalanceFactor(node.right) > 0) {
            node.right = rightRotate(node.right);
            return leftRotate(node);
        }

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
     * 判断该二叉树是不是二分搜索树
     * 
     * @return 是否是二分搜索树
     */
    public boolean isBinarySearchTree() {
        ArrayList<K> keys = new ArrayList<>();
        inOrder(root, keys);
        for (int i = 1; i < keys.size(); i ++)
            if (keys.get(i - 1).compareTo(keys.get(i)) > 0) 
                return false;

        return true;
    }

    /**
     * 二叉树的中序遍历
     * 
     * @param node 当前节点
     * @param keys 一个ArrayList<K> 的数组
     */
    private void inOrder(Node node, ArrayList<K> keys) {
        if (node == null)
            return ;

        inOrder(node.left, keys);
        keys.add(node.key);
        inOrder(node.right, keys);
    }

    /**
     * 对节点y进行向右旋转操作，返回旋转后新的根节点x
     * 
     * @param y y节点
     * @return
     */
    private Node rightRotate(Node y) {
        Node x = y.left;
        Node T3 = x.right;

        // 向右旋转过程
        x.right = y;
        y.left = T3;

        // 更新hight
        y.height = Math.max(getHeight(y.left), getHeight(y.right)) + 1;
        x.height = Math.max(getHeight(x.left), getHeight(x.right)) + 1;

        return x;
    }
    
    /**
     * 对节点y进行向左旋转操作，返回旋转后新的根节点x
     * 
     * @param y y节点
     * @return
     */
    private Node leftRotate(Node y) {
    	Node x = y.right;
    	Node T2 = x.left;
    	
    	// 向左旋转过程
    	x.left = y;
    	y.right = T2;
    	
    	// 更新height
    	y.height = Math.max(getHeight(y.left), getHeight(y.right)) + 1;
    	x.height = Math.max(getHeight(x.left), getHeight(x.right)) + 1;
    	
    	return x;
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

    public K min() {
        return minNode(root).key;
    }

    public K max() {
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
     * 删除二分搜索树的某个key对应的节点
     * 
     * @param K 对应的key
     * 
     */
    public V remove(K key) {
        Node node = getNode(root, key);
        if (node != null) {
            root = remove(root, key);
            return node.value;
        }
        return null;
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

        Node retNode;
        if (key.compareTo(node.key) < 0) {
            node.left = remove(node.left, key);
            retNode = node;
        } else if (key.compareTo(node.key) > 0) {
            node.right = remove(node.right, key);
            retNode = node;
        } else {
            // 如果左子树为空，则将右子树返回
            if (node.left == null) {
                Node rightNode = node.right;
                node.right = null;
                size --;
                retNode = rightNode;
            } 

            // 如果右子树为空，则将左子树返回
            if (node.right == null) {
                Node leftNode = node.left;
                node.left = null;
                size --;
                retNode = leftNode;
            }

            // 如果左右子树均不为空
            // 将右子树的最小元素节点删除并作为新节点，替换掉要删除的节点
            Node newNode = minNode(node.right); // 获取当前节点右子树的最小元素节点(或最大值节点也行，也叫前驱)
            newNode.right = remove(node.right, newNode.key); // 将当前节点右子树的最小元素节点删除后返回的节点树，作为新节点的右子树
            newNode.left = node.left; // 将当前节点的左子树，赋给新节点
            node.left = node.right = null;  // 将当前节点的左右子树销毁
            retNode = newNode; // 返回新的节点树
        }

        if (retNode == null)
            return null;

        // 更新height
        retNode.height = 1 + Math.max(getHeight(retNode.left), getHeight(retNode.right));

        // 计算平衡
        int balanceFactor = getBalanceFactor(retNode);

        // 平衡维护
        // LL
        if (balanceFactor > 1 && getBalanceFactor(retNode.left) >= 0)
            return rightRotate(retNode);

        // RR
        if (balanceFactor < -1 && getBalanceFactor(retNode.right) <= 0)
            return leftRotate(retNode);

        // LR
        if (balanceFactor > 1 && getBalanceFactor(retNode.left) < 0) {
            retNode.left = leftRotate(retNode.left);
            return rightRotate(retNode);
        }

        // RL
        if (balanceFactor < -1  && getBalanceFactor(retNode.right) > 0) {
            retNode.right = rightRotate(retNode.right);
            return leftRotate(retNode);
        }

        return retNode;
   }

}