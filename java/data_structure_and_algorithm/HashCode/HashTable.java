import java.util.TreeMap;

/**
 * HashTable class
 * 哈希表（散列表）泛型类
 * 
 * @param <K>
 * @param <V>
 */
public class HashTable<K, V> {
    private TreeMap<K, V>[] hashtable; // 存储元素的TreeMap
    private int size; // 元素个数
    private int M; // 计算哈希值的取模数，哈希表的长度大小

    public HashTable(int M) {
        this.M = M;
        size = 0;
        hashtable = new TreeMap[M];

        for (int i = 0; i < M; i ++)
            hashtable[i] = new TreeMap<>();
    }

    public HashTable() {
        this(97);
    }

    /**
     * 根据传入的key计算hash值键，作为哈希表的键
     * 
     * @param key 要计算hash值的键
     * @return hash值
     */
    public int hash(K key) {
        return (key.hashCode() & 0x7fffffff) % M;
    }

    /**
     * 获取当前哈希表元素的个数
     * 
     * @return 元素的个数
     */
    public int getSize() {
        return size;
    }

    /**
     * 通过过传入的key计算哈希值，然后存入的对应的哈希列中
     * 
     * @param key 要添加的key
     * @param value 要添加的value 
     */
    public void add(K key, V value) {
        TreeMap<K, V> map = hashtable[hash(key)];
        if (!map.containsKey(key))
            size ++;
        map.put(key, value);
    }

    /**
     * 删除指定key的哈希单元
     * 
     * @param key 要删除的单元key
     * @return 删除后有value值
     */
    public V remove(K key) {
        V result = null;
        TreeMap<K, V> map = hashtable[hash(key)];
        if (map.containsKey(key)) {
            result = map.remove(key);
            size --;
        }

        return result;
    }

    /**
     * 设置指定key对应的value值
     * 
     * @param key 要设置的key
     * @param value 新的value值
     */
    public void set(K key, V value) {
        TreeMap<K, V> map = hashtable[hash(key)];
        if (!map.containsKey(key))
            throw new IllegalArgumentException(key + " doesn't exist!");
        map.put(key, value);
    }

    /**
     * 检查指定key是否被包含
     * 
     * @param key 指定的key
     * @return 是否包含
     */
    public boolean contains(K key) {
        return hashtable[hash(key)].containsKey(key);
    }

    /**
     * 获取指定key的value值
     * 
     * @param key 指定的key
     * @return value值
     */
    public V get(K key) {
        return hashTable[hash(key)].get(key);
    }
}