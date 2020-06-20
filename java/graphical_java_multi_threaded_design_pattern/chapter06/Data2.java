package chapter06;

import java.util.concurrent.locks.ReadWriteLock;
import java.util.concurrent.locks.ReentrantReadWriteLock;
import java.util.concurrent.locks.Lock;

public class Data2 {
    private final char[] buffer;

    private final ReadWriteLock lock = new ReentrantReadWriteLock(true /* fair */);
    private final Lock readLock = lock.readLock();
    private final Lock writeLock = lock.writeLock();

    public Data2(int size) {
        buffer = new char[size];
        for (int i = 0; i < buffer.length; i++) {
            buffer[i] = '*';
        }
    }

    public char[] read() throws InterruptedException {
        readLock.lock();
        try {
            return doRead();
        } finally {
            readLock.unlock();
        }

    }

    public void write(char c) throws InterruptedException {
        writeLock.lock();
        try {
            doWrite(c);
        } finally {
            writeLock.unlock();
        }
    }

    private char[] doRead() {
        char[] newBuf = new char[buffer.length];
        for (int i = 0; i < newBuf.length; i++) {
            newBuf[i] = buffer[i];
        }
        slowly();

        return newBuf;
    }

    private void doWrite(char c) {
        for (int i = 0; i < buffer.length; i++) {
            buffer[i] = c;
            slowly();
        }
    }

    private void slowly() {
        try {
            Thread.sleep(50);
        } catch (InterruptedException e) {

        }
    }
}