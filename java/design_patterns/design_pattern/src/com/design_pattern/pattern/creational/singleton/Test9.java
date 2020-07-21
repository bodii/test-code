package com.design_pattern.pattern.creational.singleton;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;

public class Test9 {
    public static void main(String[] args) throws FileNotFoundException, IOException, ClassNotFoundException {
        EnumInstance instance = EnumInstance.getInstance();

        ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream("enumInstance"));
        oos.writeObject(instance);

        File file = new File("enumInstance");
        ObjectInputStream ois = new ObjectInputStream(new FileInputStream(file));
        EnumInstance enumInstance = (EnumInstance) ois.readObject();

        System.out.println(instance);
        System.out.println(enumInstance);
        System.out.println(instance == enumInstance);
    }
}