package red_black_tree;

public class RedBlackTree<K extends Comparable<K>, V> {

    private  static enum Color { RED, BLACK };

    private class Node {
        public K key;
        public V value;
        public  Node left, right;
        public Color color;

        public Node(K key, V value) {
            this.key = key;
            this.value = value;
            left = null;
            right = null;
            color = Color.RED;
        }
    }

    private Node root;
    private int size;

    public RedBlackTree() {
        root = null;
        size = 0;
    }

    /**
     * 获取当前红黑树的元素节点个数
     * 
     * @return 节点数
     */
    public int getSize() {
        return size;
    }

    /**
     * 获取当前红黑树的元素个数是否为空
     * 
     * @return 是否为空
     */
    public boolean isEmpty() {
        return size == 0;
    }

    /**
     * 获取当前点的元素
     * 如果当前节点为null，则返回黑色
     * 
     * @param node
     * @return
     */
    private Color getColor(Node node) {
        if (node == null)
            return Color.BLACK;

        return node.color;
    }

    public boolean isRed(Node node) {
        return getColor(node) == Color.RED;
    }

    /**
     * 节点元素的颜色进行旋转
     * 
     * @param node 当前节点
     */
    private void flipColors(Node node) {
        node.color = Color.RED;
        node.left.color = Color.BLACK;
        node.right.color = Color.BLACK;
    }

    //   node                           x
    //  /   \     左旋转         /         \
    // T1   x   --------->   node   T3
    //        / \                       /   \
    //      T2 T3                T1   T2
    /**
     * 指定节点向左旋转
     * 
     * @param node 当前节点
     * @return 旋转后的节点
     */
    private Node leftRotate(Node node) {
        Node x = node.right;

        // 左旋转
        node.right = x.left;
        x.left = node;

        x.color = node.color;
        node.color = Color.RED;

        return x;
    }

    //    node                      x
    //    /   \     右旋转       /  \
    //   x    T2   ------->   y   node
    //  / \                           /  \
    // y  T1                     T1  T2
    /**
     * 指定节点向右旋转
     * 
     * @param node 当前指定的节点
     * @return 转换后的节点
     */
    private Node rightRotate(Node node) {
        Node x = node.left;

        // 右旋转
        node.left = x.right;
        x.right = node;

        x.color = node.color;
        node.color = Color.RED;

        return x;
    }

    /**
     * 添加节点元素
     * 
     * @param k 键
     * @param v 值
     */
    public void add(K k, V v) {
        root = add(root, k, v);
        root.color = Color.BLACK; // 最终节点为黑色
    }

    /**
     * 向节点中插入元素
     * 递归算法
     * 
     * @param node 当前节点
     * @param k 要插入的键
     * @param v 要插入的键对应的值
     * @return 新插入的节点为的红黑树
     */
    private Node add(Node node, K k, V v) {
        if (node == null) {
            size ++;
            return new Node(k, v);
        }

        if (k.compareTo(node.key)  < 0)
            node.left = add(node.left, k, v);
        else if (k.compareTo(node.key) > 0)
            node.right = add(node.right, k, v);
        else
            node.value = v;

        // 设置当前元素节点的颜色
        if (isRed(node.right) && !isRed(node.left)) 
            node = leftRotate(node);
        if (isRed(node.left) && isRed(node.left.left))
            node = rightRotate(node);
        if (isRed(node.left) && isRed(node.right))
            flipColors(node);

        return node;
    }

    /**
     * 递归查询当前红黑树节点开始是否包含指定的key
     * 
     * @param node 当前节点
     * @param k 查询的key
     * @return 查询成功的节点
     */
    private Node getNode(Node node, K k) {
        if (node == null)
            return null;

        if (k.compareTo(node.key) < 0) 
            node.left = getNode(node.left, k);
        else if (k.compareTo(node.key) > 0)
            node.right = getNode(node.right, k);

        return node;
    }

    /**
     * 查询当前有红黑树中是否包含指定的key
     * 
     * @param k 要查询的key
     * @return 是否包含
     */
    public boolean contains(K k) {
        return getNode(root, k) != null;
    }

    /**
     * 获取指定key对应的值
     * 
     * @param k key
     * @return 值
     */
    public V get(K k) {
        Node node = getNode(root, k);
        return node == null ? null : node.value;
    }

    /**
     * 设置指定的key对应的值
     * 
     * @param k 键
     * @param v 新值
     */
    public void set(K k, V v) {
        Node node = getNode(root, k);
        if (node == null)
            throw new IllegalArgumentException(k + " dosn't exist!");

        node.value = v;
    }

    /**
     * 递归查询最小键对应的节点
     * 
     * @param node 当前节点
     * @return 最小的key的节点
     */
    private Node minNode(Node node) {
        if (node.left == null)
            return node;

        return minNode(node.left);
    }

    /**
     * 递归删除当前节点开始的向后遍历中最小key对应的节点
     * 最左节点
     * 
     * @param node 当前节点
     * @return 删除后的节点
     */
    private Node removeMinNode(Node node) {
        if (node.left == null) {
            Node rightNode = node.right;
            node.right = null;
            size --;
            return rightNode;
        }

        node.left = removeMinNode(node.left);

        return node;
    }

    /**
     * 删除指定key对应的节点
     * 
     * @param k 指定的key
     * @return 删除后的节点树
     */
    public V remove(K k) {
        Node node = getNode(root, k);
        if (node !=null) {
            root = remove(root, k);
            return node.value;
        }

        return null;
    }

    /**
     * 递归删除当前节点以及之后对应的key的节点
     * 
     * @param node 当前节点位
     * @param k 指定的key
     * @return 被删除的节点
     */
    private Node remove(Node node, K k) {
        if (node == null)
            return null;

        if (k.compareTo(node.key) < 0)  {
            node.left = remove(node.left, k);
            return node;
        } else if (k.compareTo(node.key) > 0) {
            node.right = remove(node.right, k);
            return node;
        }

        // 当前节点位的key与指定的key相等
        // 
        if (node.left == null) {
            Node rightNode = node.right;
            node.right = null;
            size -- ;
            return rightNode;
        }

        if (node.right == null) {
            Node leftNode = node.left;
            node.left = null;
            size --;
            return leftNode;
        }

        // 待删除节点的左右子树切不为空的情况
        Node result = minNode(node.right);
        result.right = removeMinNode(node.right);
        result.left = node.left;

        node.left = node.right = null;
        return result;
    }

    @Override public String toString() {
        StringBuilder result = new StringBuilder();
        result.append(String.format("Red black tree of size = %d\n", size));
        generateRBTString(root, 0, result);
        return result.toString();
    }

    private void generateRBTString(Node node, int depth, StringBuilder result) {
        if (node == null) {
            result.append(generateDepthString(depth) + "null\n");
            return;
        }

        result.append(generateDepthString(depth) + "{"+ node.key + " : "+ node.value + "["+ node.color +"] }\n");
        generateRBTString(node.left, depth+1, result);
        generateRBTString(node.right, depth+1, result);
    }

    private String generateDepthString(int depth) {
        StringBuilder result = new StringBuilder();
        for (int i = 0; i < depth; i ++) 
            result.append("--");

        return result.toString();
    }

}