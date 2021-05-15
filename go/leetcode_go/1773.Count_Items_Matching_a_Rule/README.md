### 1773. Count Items Matching a Rule
You are given an array items, where each items[i] = [typei, colori, namei] describes the type, color, and name of the ith item. You are also given a rule represented by two strings, ruleKey and ruleValue.

The ith item is said to match the rule if one of the following is true:

* ruleKey == "type" and ruleValue == typei.
* ruleKey == "color" and ruleValue == colori.
* ruleKey == "name" and ruleValue == namei.

Return the number of items that match the given rule.



Example 1:

	Input: items = [["phone","blue","pixel"],["computer","silver","lenovo"],["phone","gold","iphone"]], ruleKey = "color", ruleValue = "silver"
	Output: 1
	Explanation: There is only one item matching the given rule, which is ["computer","silver","lenovo"].

Example 2:

	Input: items = [["phone","blue","pixel"],["computer","silver","phone"],["phone","gold","iphone"]], ruleKey = "type", ruleValue = "phone"
	Output: 2
	Explanation: There are only two items matching the given rule, which are ["phone","blue","pixel"] and ["phone","gold","iphone"]. Note that the item ["computer","silver","phone"] does not match.



Constraints:

* 1 <= items.length <= 104
* 1 <= typei.length, colori.length, namei.length, ruleValue.length <= 10
* ruleKey is equal to either "type", "color", or "name".
* All strings consist only of lowercase letters.

----

### 1773. 统计匹配检索规则的物品数量
给你一个数组 items ，其中 items[i] = [typei, colori, namei] ，描述第 i 件物品的类型、颜色以及名称。

另给你一条由两个字符串 ruleKey 和 ruleValue 表示的检索规则。

如果第 i 件物品能满足下述条件之一，则认为该物品与给定的检索规则 匹配 ：

* ruleKey == "type" 且 ruleValue == typei 。
* ruleKey == "color" 且 ruleValue == colori 。
* ruleKey == "name" 且 ruleValue == namei 。

统计并返回 匹配检索规则的物品数量 。



示例 1：

	输入：items = [["phone","blue","pixel"],["computer","silver","lenovo"],["phone","gold","iphone"]], ruleKey = "color", ruleValue = "silver"
	输出：1
	解释：只有一件物品匹配检索规则，这件物品是 ["computer","silver","lenovo"] 。

示例 2：

	输入：items = [["phone","blue","pixel"],["computer","silver","phone"],["phone","gold","iphone"]], ruleKey = "type", ruleValue = "phone"
	输出：2
	解释：只有两件物品匹配检索规则，这两件物品分别是 ["phone","blue","pixel"] 和 ["phone","gold","iphone"] 。注意，["computer","silver","phone"] 未匹配检索规则。



提示：

* 1 <= items.length <= 104
* 1 <= typei.length, colori.length, namei.length, ruleValue.length <= 10
* ruleKey 等于 "type"、"color" 或 "name"
* 所有字符串仅由小写字母组成

