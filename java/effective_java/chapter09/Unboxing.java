package chapter09;

/**
    这个程序运行起来比预计的要慢一些，因为它不小心将一个局部变量(sum)声明为是装箱
    基本类型Long,而不是基本类型long.程序编译起来没有错误或者警告，变量被反复地装箱
    和拆箱，导致明显的性能下降。
 */
public class Unboxing {
    public static void main(String[] args) {
        Long sum = 0L;
        for (long i = 0; i < Integer.MAX_VALUE; i++) 
            sum += i;

        System.out.println(sum);
    }
}