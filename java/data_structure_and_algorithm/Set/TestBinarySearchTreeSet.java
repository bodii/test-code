package set;

import set.BinarySearchTreeSet;

public class TestBinarySearchTreeSet {
    public static void main(String[] args) {
        BinarySearchTreeSet<String> bsts = new BinarySearchTreeSet<>();
        
        StringBuilder words = new StringBuilder();
        words.append("JavaTM and NetbeansTM screenshots ©2017 by Oracle Corporation, all rights reserved. Reprinted with permission.");
        System.out.println(words.toString());

        System.out.println("\nString is size: " + words.length());
        for (int i = 0; i < words.length(); i++) {
            bsts.add(words.charAt(i) + "");
        }

        System.out.println("\nBinary search tree set is size: " + bsts.getSize());
    }
}