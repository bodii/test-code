package hash_code;

import java.util.Hashtable;
import java.util.TreeMap;

/**
 * 哈希表（散列表）泛型类2
 * 
 * @param <K>
 * @param <V>
 */
public class HashTable2<K, V> {

    private static final int initCapacity = 7;
    private static final int upperTol = 10;
    private static final int lowerTol = 2;

    private TreeMap<K, V>[] hashtable;
    private int size;
    private int M;

    public HashTable2(int M) {
        this.M = M;
        size = 0;
        hashtable = TreeMap[M];
        for (int i = 0; i < M; i ++)
            hashtable[i] = new TreeMap<>();
    }

    public HashTable2() {
        this(initCapacity);
    }

    /**
     * 根据传入的key生成hash值
     * 
     * @param key 传入的键
     * @return 生成后hash值
     */
    public int hash(K key) {
        return (key.hashCode() & 0x7fffffff) % M;
    }

    /**
     * 获取元素个数
     * 
     * @return 个数
     */
    public int getSize() {
        return size;
    }

    /**
     * 向hashtable中添加指定的key和值
     * 
     * @param key 传入的key
     * @param value 值
     */
    public void add(K key, V value) {
        TreeMap<K, V> map = hashtable[hash(key)];
        if (!map.containsKey(key))
            size ++;

        map.put(key, value);

        // 如果当前元素个数 大于等于 每个列的最大数 ＊ 总列数
        if (size >= upperTol * M)
            resize(2 * M);
    }
    
    /**
     * 移除指定key对应的值
     * 
     * @param key 要删除的key
     * @return 被删除的值
     */
    public V remove(K key) {
        V result = null;
        TreeMpa<K, V> map = hashtable[hash(key)];
        if (!map.contanisKey(key))
            return result;

        result = map.remove(key);
        size --;

        if (size < lowerTol * M && M / 2 >= initCapacity)
            resize(M / 2);

        return result;
    }

    /**
     * 设置指定key对应的值
     * 
     * @param key  指定的key
     * @param value 新值
     */
    public void set(K key, V value) {
        TreeMap<K, V> map = hashtable[hash(key)];
        if (!map.containsKey(key))
            throw new IllegalArgumentException(key + " doesn't exist!");

        map.put(key, value);
    }

    /**
     *  查询当前key是否存在
     * 
     * @param key 指定的key
     * @return 是否存在
     */
    public boolean contains(K key) {
        return hashtable[hash(key)].containsKey(key);
    }

    /**
     *  获取指定key对应的值
     * 
     * @param key 要获取的元素对应的key
     * @return 查询到的值
     */
    public V get(K key) {
        return hashtable[hash(key)].get(key);
    }

    /**
     * 重置当前哈希表的列长
     * 
     * @param newM 要设置长度的新值
     */
    private void resize(int newM) {
        TreeMap<K, V>[] newMap = new TreeMap[newM];
        for (int i = 0; i < newM; i++)
            newMap[i] = new TreeMpa<>();

        int oldM = M;
        this.M = newM;
        for (int i = 0; i < oldM; i ++) {
            TreeMap<K, V> map = hashtable[i];
            for (K key : map.keySet())
                newMap[hash(key)].put(key, map.get(key));
        }

        hashtable = newMap;
    }

}