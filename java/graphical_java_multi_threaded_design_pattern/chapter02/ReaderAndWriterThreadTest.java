package chapter02;

import java.util.List;
import java.util.ArrayList;
import java.util.Collections;

/**
 * java.util.ArryaList是非线程安全的类，但如果使用Collections.synchronizedList方法
 * 时行同步，就能够得到线程安全的实例。
 */
public class ReaderAndWriterThreadTest {
    public static void main(String[] args) {
        final List<Integer> list = 
            Collections.synchronizedList(new ArrayList<Integer>());
        new WriterThread(list).start();
        new ReaderThread(list).start();
    }
}