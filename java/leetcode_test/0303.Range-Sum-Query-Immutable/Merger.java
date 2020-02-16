package reange_sum_query_lmmutable;

/**
 * 线段树(区间树)融合类
 * @param <E>
 */
public interface Merger<E> {
    // 用于将两个元素计算融合成一个元素的值的方法
    E merge(E a, E b);
}