### 717. 1-bit and 2-bit Characters
We have two special characters:

* The first character can be represented by one bit 0.
* The second character can be represented by two bits (10 or 11).

Given a binary array bits that ends with 0, return true if the last character must be a one-bit character.



Example 1:

	Input: bits = [1,0,0]
	Output: true
	Explanation: The only way to decode it is two-bit character and one-bit character.
	So the last character is one-bit character.

Example 2:

	Input: bits = [1,1,1,0]
	Output: false
	Explanation: The only way to decode it is two-bit character and two-bit character.
	So the last character is not one-bit character.



Constraints:

* 1 <= bits.length <= 1000
* bits[i] is either 0 or 1.

----
### 717. 1比特与2比特字符
有两种特殊字符。第一种字符可以用一比特0来表示。第二种字符可以用两比特(10 或 11)来表示。

现给一个由若干比特组成的字符串。问最后一个字符是否必定为一个一比特字符。给定的字符串总是由0结束。

示例 1:

	输入:
	bits = [1, 0, 0]
	输出: True
	解释:
	唯一的编码方式是一个两比特字符和一个一比特字符。所以最后一个字符是一比特字符。

示例 2:

	输入:
	bits = [1, 1, 1, 0]
	输出: False
	解释:
	唯一的编码方式是两比特字符和两比特字符。所以最后一个字符不是一比特字符。

注意:

* 1 <= len(bits) <= 1000.
* bits[i] 总是0 或 1.

