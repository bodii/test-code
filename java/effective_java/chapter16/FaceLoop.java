package chapter16;

import java.util.Collection;
import java.util.EnumSet;
import java.util.Iterator;

import chapter16.Face;


public class FaceLoop {
    public static void main(String[] args) {
    
        Collection<Face> faces = EnumSet.allOf(Face.class);

        for (Iterator<Face> i = faces.iterator(); i.hasNext();)
            for (Iterator<Face> j = faces.iterator(); j.hasNext();)
                System.out.println(i.next() + " " + j.next());
    }
}