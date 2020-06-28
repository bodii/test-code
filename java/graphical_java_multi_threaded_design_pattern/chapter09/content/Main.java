package chapter09.content;

import java.io.FileOutputStream;
import java.io.IOException;

public class Main {
    public static void main(String[] args) {
        long start = System.currentTimeMillis();

        Content content1 = Retriever.retrieve("http://163.com/");
        Content content2 = Retriever.retrieve("http://cn.bing.com/");
        Content content3 = Retriever.retrieve("http://baidu.com/");

        saveToFile("163.html", content1);
        saveToFile("bing.html", content2);
        saveToFile("baidu.html", content3);

        long end = System.currentTimeMillis();

        System.out.println("Elapsed time = " + (end - start) + "msec.");
    }

    /**
     * 将content中的内容写入名为filename的文件中
     * 
     * @param filename
     * @param content
     */
    private static void saveToFile(String filename, Content content) {
        byte[] bytes = content.getBytes();
        try {
            System.out.println(Thread.currentThread().getName() + ": Saving to " + filename);
            FileOutputStream out = new FileOutputStream(filename);
            for (int i = 0; i < bytes.length; i++) {
                out.write(bytes[i]);
            }
            out.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}