package chapter06;

public class ReadWriteLock {
    private int readingReaders = 0;  // {A} ... 实际正在读取中的线程个数
    private int waitingWriters = 0;  // {B} ... 正在等待写入的线程个数
    private int writingWriters = 0;  // {C} -- 实际正在写入中的线程个数
    private boolean preferWriter = true; // 若写入优先，则为true

    public synchronized void readLock() throws InterruptedException {
        while (writingWriters > 0 || (preferWriter && waitingWriters > 0)) 
            wait();

        readingReaders ++;
    }

    public synchronized void readUnlock() {
        readingReaders --;
        preferWriter = true;
        notifyAll();
    }

    public synchronized void writeLock() throws InterruptedException {
        waitingWriters ++;
        try {
            while (readingReaders > 0 || writingWriters > 0)
                wait();
        } finally {
            waitingWriters --;
        }
    }

    public synchronized void writeUnlock() {
        writingWriters --;
        preferWriter = false;
        notifyAll();
    }
}