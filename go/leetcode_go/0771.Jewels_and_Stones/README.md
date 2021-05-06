### 771. Jewels and Stones
You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have. Each character in stones is a type of stone you have. You want to know how many of the stones you have are also jewels.

Letters are case sensitive, so "a" is considered a different type of stone from "A".



Example 1:

	Input: jewels = "aA", stones = "aAAbbbb"
	Output: 3

Example 2:

	Input: jewels = "z", stones = "ZZ"
	Output: 0



Constraints:

* 1 <= jewels.length, stones.length <= 50
* jewels and stones consist of only English letters.
* All the characters of jewels are unique.


----

### 771. 宝石与石头
给定字符串J 代表石头中宝石的类型，和字符串 S代表你拥有的石头。 S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。

J 中的字母不重复，J 和 S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。

示例 1:

	输入: J = "aA", S = "aAAbbbb"
	输出: 3

示例 2:

	输入: J = "z", S = "ZZ"
	输出: 0

注意:
* S 和 J 最多含有50个字母。
* J 中的字符不重复。

