290. Word Pattern

Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

 

Example 1:
```
Input: pattern = "abba", s = "dog cat cat dog"
Output: true
```
Example 2:
```
Input: pattern = "abba", s = "dog cat cat fish"
Output: false
```
Example 3:
```
Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false
```
Example 4:
```
Input: pattern = "abba", s = "dog dog dog dog"
Output: false
```
 

Constraints:

* 1 <= pattern.length <= 300
* pattern contains only lower-case English letters.
* 1 <= s.length <= 3000
* s contains only lower-case English letters and spaces ' '.
* s does not contain any leading or trailing spaces.
* All the words in s are separated by a single space.

-----------------------------------------------

290. 单词规律

给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:
```
输入: pattern = "abba", str = "dog cat cat dog"
输出: true
```
示例 2:
```
输入:pattern = "abba", str = "dog cat cat fish"
输出: false
```
示例 3:
```
输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
```
示例 4:
```
输入: pattern = "abba", str = "dog dog dog dog"
输出: false
```
说明:
* 你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。    