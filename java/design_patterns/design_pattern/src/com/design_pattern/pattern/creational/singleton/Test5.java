package com.design_pattern.pattern.creational.singleton;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;

public class Test5 {
    public static void main(String[] args) {
        HungrySingleton instance = HungrySingleton.getInstance();
        HungrySingleton newInstance = null;
        
        try {
            ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream("singleton_file"));
            oos.writeObject(instance);

            File file = new File("singleton_file");
            ObjectInputStream ois = new ObjectInputStream(new FileInputStream(file));
            newInstance = (HungrySingleton) ois.readObject();
        } catch (IOException | ClassNotFoundException e) {
            e.printStackTrace();
        }

        System.out.println("instance: " + instance);
        System.out.println("newinstance: " + newInstance);
        System.out.println(instance == newInstance);

    }
}