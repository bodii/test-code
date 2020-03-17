package hash_code;

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

}