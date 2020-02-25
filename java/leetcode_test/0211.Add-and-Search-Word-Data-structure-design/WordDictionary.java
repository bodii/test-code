package word_dictionary;

import java.util.TreeMap;

class WordDictionary {
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

    private Node root;
    private int size;

    /** Initialize your data structure here. */
    public WordDictionary() {
        root = new Node();
        size = 0;
    }
    
    /** Adds a word into the data structure. */
    public void addWord(String word) {
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
    
    /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
    public boolean search(String word) {
        Node current = root;
        StringBuilder prefix = new StringBuilder();
        for (int i = 0; i < word.length(); i ++) {
            char c = word.charAt(i);
            if (c == '.') {
                char j = 'a';
                for(;  j <= 'z';  j ++) {
                    prefix.append(j);
                    if (isPrefix(prefix.toString()))
                        break;
                    else 
                        prefix.delete(prefix.length() -1, prefix.length());
                } 
                current = current.next.get(j);
            } else {
                if (current.next.get(c) == null)
                    return false;
                prefix.append(c);
                current = current.next.get(c);
            }
        }

        return current.isWord;
    }

    /**
     * 查询当前词段是否是字典树的前缀树
     * 
     * @param word 要查询的单词段
     * @return 是否在字典树中是前缀树
     */
    private boolean isPrefix(String word) {
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

/**
 * Your WordDictionary object will be instantiated and called as such:
 * WordDictionary obj = new WordDictionary();
 * obj.addWord(word);
 * boolean param_2 = obj.search(word);
 */

class Test {
    public static void main(String[] args) {
        WordDictionary wordDict = new WordDictionary();

        wordDict.addWord("bad");
        wordDict.addWord("dad");
        wordDict.addWord("mad");
        System.out.println(wordDict.search("pad"));
        System.out.println(wordDict.search("bad"));
        System.out.println(wordDict.search(".ad"));
        System.out.println(wordDict.search("b.."));
    }
}