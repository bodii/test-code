package map;

import map.Mapable;

/**
 * 基于二分搜索树的映射类
 */
public class BinarySearchTreeMap<K extends Comparable<K>, V> implements Mapable<K, V> {

    /**
     * 二分搜索树映射类的内部节点类
     */
    private class MapNode {
        public K k;
        public V v;
        public MapNode left,  right;

        /**
         * 二分搜索树映射类内部节点类的构造函数(全部参数)
         * 
         * @param k 节点键
         * @param v 节点值
         * @param left 左孩子节点
         * @param right 右孩子节点
         */
        public MapNode(K k, V v, MapNode left, MapNode right) {
            this.k = k;
            this.v = v;
            this.left = left;
            this.right = right;
        }

        /**
         * 二分搜索树映射类内部节点类的构造函数
         * 
         * @param k 节点键
         * @param v 节点值
         */
        public MapNode(K k, V v) {
            this(k, v, null, null);
        }

        /**
         * 二分搜索树映射类内部节点类的构造函数
         * 
         * @param k 节点键
         */
        public MapNode(K k) {
            this(k, null);
        }

        /**
         * 将当前节点的键和值转换成字符串
         * 
         * @return 转换后的字符串
         */
        @Override 
        public String toString() {
            return k + " : " + v;
        }
    }

    private MapNode root; // 二分搜索树映射的根节点
    public int size; // 节点大小

    /**
     * 二分搜索树映射的构造函数
     */
    public BinarySearchTreeMap() {
        size = 0;
        root = null;
    }

    /**
     * 二分搜索树映射的添加元素
     * 
     * @param k 要添加元素的键
     * @param v 要添加元素的值
     */
    @Override
    public void add(K k, V v) {
        root = add(k, v, root);
    }

    /**
     * 内部递归方法，二分搜索树映射的添加元素
     * 
     * 
     * @param k 键
     * @param v 值
     * @param node 下一个节点
     * @return 返回一个添加元素后的新节点
     */
    private MapNode add(K k, V v, MapNode node) {
        if (node == null) { // 如果是未使用的节点
            node = new MapNode(k, v);
            size ++;
            
            return node;
        } 

        if (k.compareTo(node.k) < 0) //  传入的键比当前节点的键小
            node.left = add(k, v, node.left);
        else if (k.compareTo(node.k) > 0) //  传入的键比当前节点的键大
            node.right = add(k, v, node.right);
        else if (k.compareTo(node.k) == 0) // 传入的键与当前节点的键相等
            set(k, v);  // 更新
            
        return node;
    }

    /**
     * 二分搜索树映射删除元素
     * 
     * @param k 要删除元素对应的键
     * @return 删除的元素
     */
    @Override
    public V remove(K k) {
        return remove(k, root).v;
    }

    private MapNode remove(K k, MapNode node) {
        if (node == null)
            return null;

        if (k.compareTo(node.k) < 0) {
            node.left = remove(k, node.left);
            return node;
        } else if (k.compareTo(node.k) > 0) {
            node.right = remove(k, node.right);
            return node;
        } else {
            if (node.left == null) {
                MapNode rightNode = node.right;
                node.right = null;
                size --;
                return rightNode;
            }

            if (node.right == null) {
                MapNode leftNode = node.left;
                node.left = null;
                size --;
                return leftNode;
            }

            MapNode minNode = getMinNode(node.right);
            minNode.right = removeMin(node.right);
            minNode.left = node.left;
            node.left = node.right = null;
            size --;

            return minNode;
        }
    }

    /**
     * 删除最小元素节点，并返回最小元素节点的值
     * 
     * @return 最小元素节点的值
     */
    public V removeMin() {
        MapNode minNode = getMinNode(root);
        root = removeMin(root);

        return minNode.v;
    }

    /**
     * 递归删除最小元素节点
     * 
     * @param node
     * @return
     */
    private MapNode removeMin(MapNode node) {
        if (node.left == null) {
            MapNode rightNode = node.right;
            node.right = null;
            size --;

            return rightNode;
        }

        node.left = removeMin(node.left);

        return node;
    }

    /**
     * 删除最大元素节点
     * 
     * @return 被删除的节点元素的值
     */
    public V removeMax() {
        MapNode maxNode = getMaxNode(root);
        root = removeMax(root);

        return maxNode.v;
    }

    /**
     * 递归删除最大元素的节点，并返回最大元素
     * 
     * @param node 节点
     * @return 下个节点或最大节点
     */
    private MapNode removeMax(MapNode node) {
        if (node.right == null) {
            MapNode leftNode = node.left;
            node.left = null;
            size --;

            return leftNode;
        }

        node.right = removeMax(node.right);

        return node;
    }

    /**
     * 递归获取当前二分搜索树映射的最小元素的节点
     * 
     * @param node 节点
     * @return 获取到的最小节点
     */
    private MapNode getMinNode(MapNode node) {
        if (node.left == null)
            return node;

        return getMinNode(node.left);
    }

    /**
     * 递归获取当前二分搜索树映射的最大元素的节点
     * 
     * @param node 节点
     * @return 获取到的最大节点
     */
    private MapNode getMaxNode(MapNode node) {
        if (node.right == null)
            return node;

        return getMaxNode(node.right);
    }

    /**
     * 内部递归方法，二分搜索树映射通过键查询k对应的节点
     * 
     * @param k 键
     * @param node 节点
     * @return 相等k节点或不相等k的下一个节点
     */
    private MapNode getEqualsKToNode(K k, MapNode node) {
        if (node == null) 
            return null;

        if (k.compareTo(node.k) == 0) 
            return node;
        else if (k.compareTo(node.k) < 0) 
            return getEqualsKToNode(k, node.left);
        else //(k.compareTo(node.k) > 0)
            return getEqualsKToNode(k, node.right);
    }

    /**
     * 查看指的键是否包含在二分搜索树映射中
     * 
     * @param k 要查找的键
     * @return 返回是或否
     */
    @Override
    public boolean contains(K k) {
        return getEqualsKToNode(k, root) != null;
    }

    /**
     * 获取键对应的值
     * 
     * @param k 键
     * @return 返回键对应值
     */
    @Override
    public V get(K k) {
        MapNode node = getEqualsKToNode(k, root);

        return node != null ? node.v : null;
    }

    /**
     * 设置二分搜索树的键对应的新值
     * 
     * @param k 键
     * @param v 新值
     */
    @Override
    public void set(K k, V v) {
        MapNode node = getEqualsKToNode(k, root);

        if (node == null)
            throw new IllegalArgumentException(k + " doesn't exists!");

        node.v = v;
    }

    /**
     * 获取二分搜索树的元素节点个数(尺寸大小)
     * 
     * @return 元素节点个数
     */
    @Override
    public int getSize() {
        return size;
    }

    /**
     * 检测当前二分搜索树映射的元素数量是否为空
     * 
     * @return 是否为空
     */
    @Override
    public boolean isEmpty() {
        return size == 0;
    }

    /**
     * 将当前二分搜索树映射转换成字符串
     * 
     * @return 转换后的字符串
     */
    @Override 
    public String toString() {
        StringBuilder result = new StringBuilder();

        result.append(String.format("Binary search tree map size is %d%n", size));
        result.append("{");
        generateBSTMString(root, result);
        result.replace(result.length() - 2, result.length(), "}");

        return result.toString();
    }

    /**
     * 将二分搜索树映射中的所有元素转换成字符串
     * 
     * @param node 当前节点
     * @param result 向该字符串中添加
     */
    private void generateBSTMString(MapNode node, StringBuilder result) {
        if (node == null) 
            return;

        result.append(node + ", ");
        generateBSTMString(node.left, result);
        generateBSTMString(node.right, result);
    }

} 