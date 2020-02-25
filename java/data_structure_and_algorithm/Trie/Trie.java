package trie;

import java.util.TreeMap;

/**
 * 字典树
 */
public class Trie {
    // 内部节点类
    private class Node {
        public boolean isWord; // 是否是单词
        public TreeMap<Character, Node> next; // 下一个节点

        public Node(boolean isWord) {
            this.isWord = isWord;
            next = new TreeMap<>();
        }

        public Node() {
            this(false);
        }
    }

    private Node root; // 整个字典树的根节点
    private int size; // 字典树存储了多少个单词

    public Trie() {
         root = new Node();
         size = 0;
    }

    /**
     * 获取当前字典树中单词的个数
     * @return
     */
    public int getSize() {
        return size;
    }

    /**
     * 向字典树中添加一个单词
     * 
     * @param word 向添加的单词
     */
    public void add(String word) {
        Node current = root;
        // 遍历这个单词
        for (int i = 0; i < word.length(); i ++) {
            char c = word.charAt(i);
            if (current.next.get(c) == null) // 如果单词的下一个节点为空
                current.next.put(c, new Node()); // 将下一个中添加字符映射
            
            current = current.next.get(c); // 向下移动一个指针
        }

        // 如果遍历完成后该单词没有添加过
        if (!current.isWord) {
            current.isWord = true; // 设置当前节点的是单词的结尾
            size ++; // 单词数加1
        }
    }

    /**
     * 查询当前单词是否被字典树包含
     * 
     * @param word 要查询的单词
     * @return  是否包含
     */
    public boolean contains(String word) {
        Node current = root;
        for (int i = 0; i < word.length(); i ++) {
            char c = word.charAt(i);
            if (current.next.get(c) == null)
                return false;
            current = current.next.get(c);
        }

        return current.isWord;
    }

    /**
     * 查询当前词段是否是字典树的前缀树
     * 
     * @param word 要查询的单词段
     * @return 是否在字典树中是前缀树
     */
    public boolean isPrefix(String word) {
        Node current = root;
        for (int i = 0; i < word.length(); i ++) {
            char c = word.charAt(i);
            if (current.next.get(c) == null)
                return false;

            current = current.next.get(c);
        }

        return true;
    }

}