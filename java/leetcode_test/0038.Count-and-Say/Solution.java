package count_and_say;

import java.util.TreeMap;

/**
 * 下一个数是对上一个数的描述，比方说 1211 里有 “ 1 个 1 ， 1 个 2 ， 2 个 1 ” ，
 * 那么 111221 就是它的下一个数。通常我们把这个数列叫做“外观数列”。
 * 
 * * !! 这个方法包含已生成的所有外观数列
 */
class Solution {
    public String countAndSay(int n) {
        if (1 > n  || 30 < n) return "";

        int i = 1; // 初始化一个递进的索引
        String current = "1"; // 初始化1的外观数列
        TreeMap<Integer, String> map =  new TreeMap<>(); // 存放外观数列的容器
        map.put(i, current); // 将1添加到容器中

        while (i < n) { // 遍历添加指定数量的外观数列
            current = nextAll(map.get(i));
            map.put(++i, current);
        }

        // 返回指定数量的最后一个外观数列
        return current;
    }

    /**
     * 根据当前外观数列，获取它的外观数列
     * 
     * @param current 当前外观数列
     * @return 新生成的外观数列
     */
    private String nextAll(String current) {
        StringBuilder result = new StringBuilder();
        for (int i = 0; i < current.length();) {
            int step = step(current, i);
            result.append(step);
            result.append(current.charAt(i));
            i += step;
        }
        return result.toString();
    }

    /**
     *  获取指定索引的元素在当前索引位向后数相邻的有几位
     * 
     * @param current 当前外观数列
     * @param index 当前索引
     * @return 有几位与之相邻并值相同的数
     */
    private int step(String current, int index) {
        char v = current.charAt(index);
        int step = 1;
        while (++index < current.length()) {
            if (current.charAt(index) != v)
                break;
            step ++;
        }

        return step;
    }
}


class Test {
    public static void main(String[] args) {
        Solution solution = new Solution();
        for(int i = 1; i <= 13; i ++)
            System.out.println(i + " :  " +solution.countAndSay(i));

    }
}