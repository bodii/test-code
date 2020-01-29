package map;

import map.BinarySearchTreeMap;

public class TestBinarySearchTreeMap{
    public static void main(String[] args) {
        BinarySearchTreeMap<Integer, String> llm = new BinarySearchTreeMap<>();
        
        StringBuilder words = new StringBuilder();
        words.append("JavaTM and NetbeansTM screenshots ©2017 by Oracle Corporation, all rights reserved. Reprinted with permission.");
        System.out.println(words.toString());

        System.out.println("\nString is size: " + words.length());
        for (int i = 0; i < words.length(); i++) {
            llm.add(i, words.charAt(i) + "");
        }

        System.out.println("\nlinked list map is size: " + llm.getSize());

        System.out.println("\nlinked list map: " + llm);

        System.out.println("\nremove k is 2 of linked list map: " + llm.remove(2));

        System.out.println("\nlinked list map: " + llm);

        System.out.println("\nremove k is 109 of linked list map: " + llm.remove(109));

        System.out.println("\nlinked list map: " + llm);

        System.out.println("\nset k is 7 the new value 'gogogo': " );
        llm.set(7, "gogogo");
        System.out.println("\nlinked list map: " + llm);

        System.out.println("\nget k is 17 the value: " + llm.get(17) );

    }
}