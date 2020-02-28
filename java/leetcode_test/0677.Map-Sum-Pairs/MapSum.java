package map_sum;

import java.util.TreeMap;

/**
 * 键值映射类
 * MapSum class
 */
class MapSum {
    /**
     * 内部节点类
     */
    private class Node {
        public int val;
        public boolean isWord;
        public TreeMap<Character, Node> next; 

        
        public Node(int val, boolean isWord) {
            this.val = val;
            this.isWord = isWord;
            next = new TreeMap<>();
        }

        public Node() {
            this(0, false);
        }
    } 

    private Node root; // 前缀树的根节点
    private int size; // 已添加的组合词组的个数

    /** Initialize your data structure here. */
    public MapSum() {
        root = new Node();
        size = 0;
    }
    
    /**
     * 插入元素和值
     * 
     * @param key 元素
     * @param val 数值
     */
    public void insert(String key, int val) {
        Node current = root;
        for (int i = 0; i < key.length(); i ++) {
            char c = key.charAt(i);
            if (current.next.get(c) == null)
                current.next.put(c, new Node());

            current = current.next.get(c);
        }

        current.val = val;
        if (!current.isWord) {
            current.isWord = true;
            size ++;
        }
    }

    /**
     * 计算当前字符串在
     * @param prefix 要计算的字符串
     * @return 返回计算后的值
     */
    public int sum(String prefix) {
        int result = 0;
        Node current = root;
        // 如果要查询的字符串为空则直接返回0
        if (prefix.isEmpty()) 
            return result;

        // 是否是前缀
        boolean isPrefix = isPrefix(prefix);
        if (!isPrefix) // 如果前缀树中不包含当前字符串，则返回0
            return result;

        // 是否是单词
        boolean isWord = isWord(prefix);
        for (int i = 0; i < prefix.length(); i ++) {
            char c = prefix.charAt(i);
            current = current.next.get(c);
            // 并且当前键对应的值小于等于0, 跳过
            if (current.val <= 0)
                continue;

            // 如果要查询的字符串不是单词边界，则累加对应的值
            if (!isWord)
                result += current.val;
        }

        // 如果要查询的字符串是单词边界，则以当前键对应的值作为初始值，
        // 并向后查看是否结束，如果没有结束累加后面的值
        if (isWord)
            result = current.val;

        // 如果当前查询的字符串只是前缀树的前缀字符，并遍历累加后面的值
        for (char nextChar : current.next.keySet())
            result += getVal(current, nextChar, 0);

        return result;
    }

    /**
     * 检测当前字符串是否在前缀树中是一个已添加的组合词组 
     * 
     * @param word 要查询的字符串
     * @return 是否在前缀树中是一个已添加的组合词组
     */
    private boolean isWord(String word) {
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
     * 检测当前字符串是否是前缀树的前缀元素 
     * 
     * @param prefix 要查询的字符串
     * @return 是否是前缀树的前缀元素 
     */
    private boolean isPrefix(String prefix) {
        Node current = root;
        for (int i = 0; i < prefix.length(); i ++) {
            char c = prefix.charAt(i);

            if (current.next.get(c) == null)
                return false;

            current = current.next.get(c);
        }

        return true;
    }

    /**
     * 获取前缀字符串后面所有匹配前缀键的值
     * 
     * @param node 上一个节点
     * @param c 要查询的键
     * @param total 累计的值
     * @return 计算后的值
     */
    private int getVal(Node node, char c, int total) {
        if (node.next.get(c) == null)
            return total;

        node = node.next.get(c);
        if (node.val > 0)
            System.out.println("val: " + node.val);
        total += node.val;
        for (char nextChar : node.next.keySet())
            return getVal(node, nextChar, total);

        return total;
    }

    /**
     * 打印当前前缀树
     */
    public void print() {
        Node current = root;
        for (char nextChar : current.next.keySet())
            print(current, nextChar, 1);
    }

    /**
     * 递归打印当前前缀树内容
     * 
     * @param prevNode 上一个节点
     * @param c 下一个节点的key
     * @param level 第几层
     */
    private void print(Node prevNode, char c, int level) {
        System.out.println("level: " + level + "; val: " + prevNode.next.get(c).val + "  char: " + c);

        if (prevNode.next.get(c) == null)
            return;

        prevNode = prevNode.next.get(c);
        for (char nextChar : prevNode.next.keySet()) {
            print(prevNode, nextChar, level+1);
        }
    }
}

/**
 * Your MapSum object will be instantiated and called as such:
 * MapSum obj = new MapSum();
 * obj.insert(key,val);
 * int param_2 = obj.sum(prefix);
 */

class Test {
    public static void main(String[] args) {
        MapSum sum = new MapSum();
        // sum.insert("apple", 3);
        // sum.insert("a", 3);
        // // sum.print();
        // System.out.println(sum.sum("ap"));

        // sum.insert("b", 2);
        // System.out.println(sum.sum("a"));

        // sum.insert("app", 2);
        // System.out.println(sum.sum("ap"));

        // sum.insert("ap", 10);
        // System.out.println(sum.sum("ap"));

        // sum.insert("apdp", 10);
        // System.out.println(sum.sum("ap"));
        // sum.print();

        /*
        ["MapSum", "insert", "sum", "insert", "sum", "sum", "insert", "sum"]
        [[], ["aa",3], ["a"], ["aa",2], ["a"], ["aa"], ["aaa", 3], ["aaa"]]

        output:
        [null,null,3,null,2,2,null,3]
        */
        // sum.insert("aa", 3);
        // // sum.print();
        // System.out.println(sum.sum("a"));

        // sum.insert("aa", 2);
        // System.out.println(sum.sum("a"));
        // System.out.println(sum.sum("aa"));
        // sum.insert("aaa", 3);
        // System.out.println(sum.sum("aaa"));

        /*
        ["MapSum", "insert", "sum", "insert", "sum", "sum", "insert", "sum", "sum", "sum", "insert", "sum", "sum"]
        [[], ["aa",3], ["a"], ["aa",2], ["a"], ["aa"], ["aaa", 3], 
        ["aaa"], ["bbb"], ["ccc"], ["aewfwaefjeoawefjwoeajfowajfoewajfoawefjeowajfowaj", 111], ["aa"], ["a"]]

        output:
        [null,null,3,null,2,2,null,3,0,0,null,5,116]
        */
        // sum.insert("aa", 3);
        // System.out.println(sum.sum("a"));

        // sum.insert("aa", 2);
        // System.out.println(sum.sum("a"));
        // System.out.println(sum.sum("aa"));

        // sum.insert("aaa", 3);
        // System.out.println(sum.sum("aaa"));
        // System.out.println(sum.sum("bbb"));
        // System.out.println(sum.sum("ccc"));

        // sum.insert("aewfwaefjeoawefjwoeajfowajfoewajfoawefjeowajfowaj", 111);
        // System.out.println(sum.sum("aa"));// ?5
        // System.out.println(sum.sum("a"));


        /*
        ["MapSum", "insert", "sum", "insert", "sum", "sum", "insert", "sum", "sum", "sum", "insert", "sum", "sum", "sum", "sum", "sum", "insert", "insert", "insert", "sum", "sum", "sum", "sum", "sum", "sum", "insert", "sum", "sum", "sum", "sum"]
        [[], ["aa",3], ["a"], ["aa",2], ["a"], ["aa"], ["aaa", 3], ["aaa"], ["bbb"], ["ccc"], ["aewfwaefjeoawefjwoeajfowajfoewajfoawefjeowajfowaj", 111], ["aa"], ["a"], ["b"], ["c"], ["aewfwaefjeoawefjwoeajfowajfoewajfoawefjeowajfowaj"], 
        ["bb", 143], ["cc", 144], ["aew", 145], ["bb"], ["cc"], ["aew"], ["dddd"], ["cdcd"], ["aab"], ["aab", 33], ["aab"], ["ab"], ["aa"], ["a"]]

        output: 
        [null,null,3,null,2,2,null,3,0,0,null,5,116,0,0,111,null,null,null,143,144,256,0,0,0,null,33,0,38,294]
        */
        sum.insert("aa", 3);
        System.out.println(sum.sum("a"));// 3

        sum.insert("aa", 2);
        System.out.println(sum.sum("a")); // 2
        System.out.println(sum.sum("aa")); // 2

        sum.insert("aaa", 3);
        System.out.println(sum.sum("aaa")); // 3
        System.out.println(sum.sum("bbb")); // 0
        System.out.println(sum.sum("ccc")); // 0

        sum.insert("aewfwaefjeoawefjwoeajfowajfoewajfoawefjeowajfowaj", 111);
        System.out.println(sum.sum("aa")); // 5
        System.out.println(sum.sum("a")); // 116
        System.out.println(sum.sum("b")); // 0
        System.out.println(sum.sum("c")); // 0
        System.out.println(sum.sum("aewfwaefjeoawefjwoeajfowajfoewajfoawefjeowajfowaj")); // 111

        sum.insert("bb", 143);
        sum.insert("cc", 144);
        sum.insert("aew", 145);
        System.out.println(sum.sum("bb")); // 143
        System.out.println(sum.sum("cc")); // 144
        System.out.println(sum.sum("aew")); // 256
        System.out.println(sum.sum("dddd")); // 0
        System.out.println(sum.sum("cdcd")); // 0
        System.out.println(sum.sum("aab")); // 0

        sum.insert("aab", 33);
        System.out.println(sum.sum("aab")); // 33
        System.out.println(sum.sum("ab")); // 0
        System.out.println(sum.sum("aa")); // 38
        System.out.println(sum.sum("a")); // 294
    }
}