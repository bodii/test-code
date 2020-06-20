package chapter06;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReadWriteLock;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class Database<K, V> {

    private final Map<K, V> map = new HashMap<>();
    private final ReadWriteLock lock = new ReentrantReadWriteLock(true /* fair */);
    private final Lock readLock = lock.readLock();
    private final Lock writeLock = lock.writeLock();

    /**
     * 全部清除
     */
    public void clear() {
        writeLock.lock();
        readLock.lock();

        try {
            verySlowly();
            map.clear();
        } finally {
            writeLock.unlock();
            readLock.unlock();
        }
    }

    /**
     * 给key 分配value
     * @param key
     * @param value
     */
    public void assign(K key, V value) {
        readLock.lock();
       
        try {
            verySlowly();
            map.put(key, value);
        } finally {
            readLock.unlock();
        }
        
    }

    /**
     * 获取给key分配的值
     * 
     * @param key
     * @return
     */
    public V retrieve(K key) {
        writeLock.lock();

        try {
            slowly();
            return map.get(key);
        } finally {
            writeLock.unlock();
        }
    }

    private void slowly() {
        try {
            Thread.sleep(50);
        } catch (InterruptedException e) {

        }
    }

    private void verySlowly() {
        try {
            Thread.sleep(500);
        } catch (InterruptedException e) {

        }
    }


    
}