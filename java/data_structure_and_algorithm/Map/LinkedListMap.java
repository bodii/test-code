package map;

import java.nio.channels.FileChannel.MapMode;

import map.Mapable;

/**
 * 链表映射类
 * 
 * @param <K>
 * @param <V>
 */
public class LinkedListMap<K, V> implements Mapable<K, V> {

    /**
     * 链表映射类内部的映射节点类
     */
    private class MapNode {
        public K k; /* 键 */
        public V v; /* 值 */
        public MapNode next; /* 下一个映射节点 */

        /**
         * 映射节点类的构造方法
         * @param k
         * @param v
         * @param MapNode 下一个映射节点
         */
        public MapNode(K k, V v, MapNode next) {
            this.k = k;
            this.v = v;
            this.next = next;
        }

         /**
         * 映射节点类的构造方法
         * @param k
         * @param v
         */
        public MapNode(K k, V v) {
            this(k, v, null);
        }

        /**
         * 声明一个空映射节点的构造方法
         */
        public MapNode() {
            this(null, null, null);
        }

        /**
         * 将当前映射节点转成字符串
         * 
         * @return 生成的字符串
         */
        @Override public String toString() {
            return String.format("%s : %s", k.toString(), v.toString());
        }
    }

    private MapNode dummyHead; // 虚拟映射头节点
    private int size; // 映射的元素个数(大小)

    /**
     * 链表映射类的构造函数
     */
    public LinkedListMap() {
        dummyHead = new MapNode();
        size = 0;
    }

    /**
     * 获取指定键对应的节点
     * 
     * @param k 键
     * @return 对应的节点
     */
    private MapNode getNode(K k) {
        MapNode current = dummyHead.next;

        while(current != null) {
            if (current.k.equals(k)) 
                return current;
            current = current.next;
        }

        return null;
    }

    /**
     * 添加映射值
     * 
     * @param k 键
     * @param v 值
     */
    @Override
    public void add(K k, V v) {
        if (contains(k)) { // 说明当前键存在
            set(k, v); // 更新键对应的值
            return;
        }
        
        // dummyHead.next = new MapNode(k, v, dummyHead.next); // 将新元素添加到头部

        // 将新元素添加到尾部
        MapNode prev = dummyHead;

        while(prev != null) {
            if (prev.next == null) {
                prev.next = new MapNode(k, v);
                size ++;
                return;
            }

            prev = prev.next;
        }

    }

    /**
     * 删除映射的值
     * 
     * @param k 要删除值对应的键
     * @return 返回删除值
     */
    @Override
    public V remove(K k) {
        MapNode prev = dummyHead;

        while(prev != null && prev.next != null) {
            if (prev.next.k.equals(k)) {
                V delVal = prev.next.v;
                prev.next = prev.next.next;
                size --;
                return delVal;
            } 

            prev = prev.next;
        }

        return null;
    }

    /**
     * 查看指定的键是否存在
     * 
     * @param k 查找的值对应的键
     * @return 是否存在
     */
    @Override
    public boolean contains(K k) {
        MapNode current = dummyHead.next;

        while(current != null) {
            if (current.k.equals(k)) 
                return true;
            current = current.next;
        }

        return false;
    }

    /**
     * 获取指定键对应的值
     * 
     * @param k 键
     * @return 值
     */
    @Override
    public V get(K k) {
        MapNode current = dummyHead.next;

        while(current != null) {
            if (current.k.equals(k)) 
                return current.v;
            current = current.next;
        }

        return null;
    }

    /**
     * 设置指定的键的值
     * 
     * @param k 键
     * @param v 值
     */
    @Override
    public void set(K k, V v) {
        if (!contains(k))
            throw new IllegalArgumentException(k + "doesn't exist!");

        MapNode current = dummyHead.next;

        while(current != null) {
            if (current.k.equals(k)) 
                current.v = v;
            current = current.next;
        }
    }

    /**
     * 获取当前链表映射的大小
     * 
     * @return 大小
     */
    @Override
    public int getSize() {
        return size;
    }

    /**
     * 查看该链表映射是否为空
     * 
     * @return 是否为空
     */
    @Override
    public boolean isEmpty() {
        return size == 0;
    }

    /**
     * 将当前链表映射转成字符串
     * 
     * @return 返回生成的字符串
     */
    @Override
    public String toString() {
        StringBuilder result = new StringBuilder();
        MapNode current = dummyHead.next;

        result.append(String.format("linked list map is size = %d:%n", size));
        result.append("{");

        while(current != null) {
            result.append(current + ", ");
            current = current.next;
        }

        if (result.charAt(result.length() - 2) == ',') 
            result.replace(result.length() - 2, result.length(), "");

        result.append("}");

        return result.toString();
    }

}