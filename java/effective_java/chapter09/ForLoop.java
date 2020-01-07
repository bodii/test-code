package chapter09;

import java.util.*;

import org.w3c.dom.Element;

public class ForLoop {
    public static void main(String[] args) {

        List<String> c = new ArrayList<>();
        List<String> c2 = new ArrayList<>();
        c.add(new String("a"));
        // 返回一个迭代器，并获取是否有下一个元素
        for (Iterator<String> i = c.iterator(); i.hasNext();) {
            String e = i.next();
        }

        for(Iterator<String>i2 = c2.iterator(); i2.hasNext();) {
            String e2 = i2.next();
        }

        Set<Integer> s = new HashSet<>();

        Iterator<Integer> i3  = s.iterator();

        Element[] elements = new Element[2];
        // 当见到冒号(:)时，可以把它读作“在......里面”。
        for (Element e : elements) {

        }

    }
}