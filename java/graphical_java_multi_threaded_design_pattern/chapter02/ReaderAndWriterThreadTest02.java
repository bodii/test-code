package chapter02;

import java.util.List;
import java.util.concurrent.CopyOnWriteArrayList;

/**
 * java.util.concurrent.CopyOnWriteArrayList类是线程安全的。与使用Collections.synchronizedList
 * 方法进行同步不同，CopyOnWriteArrayList类是采用copy-on-write技术来避免读写冲突的。
 */
public class ReaderAndWriterThreadTest02 {
    public static void main(String[] args) {
        final List<Integer> list = new
            CopyOnWriteArrayList<Integer>();
        new WriterThread(list).start();
        new ReaderThread(list).start();
    }
}

/*
所谓copy-on-write，就是“写时复制”的意思。如果使用copy-on-write，当对集合执行“写”操作时，
内部已确保安全的数组就会被整体复制。复制之后，就无需在使用迭代器依次读取元素时担心元素会被修改了。
因此，CopyOnWriteArrayList类(及迭代器)绝对不会抛出ConcurrentModificationException异常。
*/