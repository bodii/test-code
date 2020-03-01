package union_find;

/**
 * 并查集接口类
 */
public interface UnionFindable {

    /**
     * 查询两个元素是否是合并的
     * 
     * @param p 元素
     * @param q 元素
     * @return 是否是合并的
     */
    boolean isConnected(int p, int q);

    /**
     * 将两个元素合并起来
     * @param p 元素
     * @param q 元素
     */
    void unionElements(int p, int q); 

    /**
     * 获取并查集元素的个数
     * 
     * @return 个数
     */
    int getSize();
}