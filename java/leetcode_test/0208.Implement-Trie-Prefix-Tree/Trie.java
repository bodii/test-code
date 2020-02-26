package prefix_tree;

import java.util.TreeMap;

class Trie {
    private class Node {
        public boolean isWord;
        public TreeMap<Character, Node> next;

        public Node(boolean isWord) {
            this.isWord = isWord;
            next = new TreeMap<>();
        }

        public Node() {
            this(false);
        }
    }

    private Node root;
    private int size;

    /** Initialize your data structure here. */
    public Trie() {
        root = new Node();
        size = 0;
    }

    private int getSize() {
        return size;
    }

    /** Inserts a word into the trie. */
    public void insert(String word) {
        Node current = root;
        for (int i = 0; i < word.length(); i ++) {
            char c = word.charAt(i);
            if (current.next.get(c) == null)
                current.next.put(c, new Node());

            current = current.next.get(c);
        }

        if (!current.isWord) {
            current.isWord = true;
            size ++;
        }
    }

    /** Returns if the word is in the trie. */
    public boolean search(String word) {
        Node current = root;
        for (int i = 0; i < word.length(); i ++) {
            char c = word.charAt(i);

            if (current.next.get(c) == null)
                return false;

            current = current.next.get(c);
        }

        return current.isWord;
    }

    /** Returns if there is any word in the trie that starts with the given prefix. */
    public boolean startsWith(String prefix) {
        Node current = root;
        for (int i = 0; i < prefix.length(); i ++) {
            char c = prefix.charAt(i);

            if (current.next.get(c) == null)
                return false;

            current = current.next.get(c);
        }

        return true;
    }
}


class Test {
    public static void main(String[] args) {
        Trie trie = new Trie();

        trie.insert("apple");
        System.out.println("prefix tree is contains 'apple' : " + trie.search("apple"));   // 返回 true
        System.out.println("prefix tree is contains 'app' : " + trie.search("app"));     // 返回 false
        System.out.println(" 'app' prefix whether in prefix tree : " + trie.startsWith("app")); // 返回 true
        trie.insert("app");   
        System.out.println(" prefix tree is contains 'app': " + trie.search("app"));     // 返回 true

    }
}