package chapter18;

import java.util.concurrent.atomic.AtomicLong;

/**
 * 使用AtomicLong类，它是java.util.conurrent.atomic的组成部分
 * 这个包为在单个变量上进行免锁定、线程安全的编程提供了基本类型。
 * 虽然volatile只提供了同步的通信效果，但这个包还提供了原子性。
 */
public class GenerateSerialNumber2 {
    private static final AtomicLong nextSerialNum = new AtomicLong();

    public static long generateserialNumber() {
        return nextSerialNum.getAndIncrement();
    }
}
