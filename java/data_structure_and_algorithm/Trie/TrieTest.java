package trie;

import trie.Trie;

public class TrieTest {
    public static void main(String[] args) {
        Trie trie = new Trie();
        trie.add("add");
        trie.add("append");
        System.out.println("trie test contains 'add': " + trie.contains("add"));
        System.out.println("trie test isPrefix 'd': " + trie.isPrefix("d")); 
    }
}