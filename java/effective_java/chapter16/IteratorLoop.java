package chapter16;

import java.util.Iterator;

import javax.lang.model.element.Element;

public class IteratorLoop {
    public static void main(String[] args) {
        Iterable<Element> c = new Iterator<>()
        for (Iterator<Element> i = c.iterator(); i.hasNext()) {
            Element e = i.next();
        }
    }
}