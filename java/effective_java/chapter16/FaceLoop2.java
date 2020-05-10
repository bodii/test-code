package chapter16;

import java.util.Collection;
import java.util.EnumSet;
import java.util.Iterator;

import chapter16.Face;


public class FaceLoop2 {
    public static void main(String[] args) {
    
        Collection<Face> faces = EnumSet.allOf(Face.class);

        for (Face i : faces)
            for (Face j : faces)
                System.out.println(i + " " + j);
    }
}