package map;

/**
 * Mapable interface
 * 映射接口类
 * 
 * @param <K>
 * @param <V>
 */
public interface Mapable<K, V> {
    /**
     * 添加映射值
     * 
     * @param k 键
     * @param v 值
     */
    void add(K k, V v);

    /**
     * 删除映射值
     * 
     * @param k 要删除映射值的键
     * @return 删除后的值
     */
    V remove(K k);

    /**
     * 查看指的键是否存在
     * 
     * @param k 键
     * @return 是否包含
     */
    boolean contains(K k);

    /**
     * 获取指定键对应的值
     * 
     * @param k 键
     * @return 值
     */
    V get(K k);

    /**
     * 设置指定的键的值
     * 
     * @param k 键
     * @param v 值
     */
    void set(K k, V v);

    /**
     * 获取映射的大小
     * 
     * @return 大小
     */
    int getSize();

    /**
     * 检测映射是否为空
     * 
     * @return 是否为空
     */
    boolean isEmpty();
}