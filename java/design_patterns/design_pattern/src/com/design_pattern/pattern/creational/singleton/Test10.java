package com.design_pattern.pattern.creational.singleton;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;

public class Test10 {
    public static void main(String[] args) throws FileNotFoundException, IOException, ClassNotFoundException {
        EnumInstance instance = EnumInstance.getInstance();
        instance.setData(new Object());

        ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream("enumInstance2"));
        oos.writeObject(instance);

        ObjectInputStream ois = new ObjectInputStream(new FileInputStream(new File("enumInstance2")));
        EnumInstance  enumInstance = (EnumInstance) ois.readObject();
        
        System.out.println(instance.getData());
        System.out.println(enumInstance.getData());
        System.out.println(instance == enumInstance);
    }
}