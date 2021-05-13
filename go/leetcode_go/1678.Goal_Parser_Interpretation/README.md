### 1678. Goal Parser Interpretation
You own a Goal Parser that can interpret a string command. The command consists of an alphabet of "G", "()" and/or "(al)" in some order. The Goal Parser will interpret "G" as the string "G", "()" as the string "o", and "(al)" as the string "al". The interpreted strings are then concatenated in the original order.

Given the string command, return the Goal Parser's interpretation of command.



Example 1:

	Input: command = "G()(al)"
	Output: "Goal"
	Explanation: The Goal Parser interprets the command as follows:
	G -> G
	() -> o
	(al) -> al
	The final concatenated result is "Goal".

Example 2:

	Input: command = "G()()()()(al)"
	Output: "Gooooal"

Example 3:

	Input: command = "(al)G(al)()()G"
	Output: "alGalooG"



Constraints:

* 1 <= command.length <= 100
* command consists of "G", "()", and/or "(al)" in some order.

----

### 1678. 设计 Goal 解析器
请你设计一个可以解释字符串 command 的 Goal 解析器 。command 由 "G"、"()" 和/或 "(al)" 按某种顺序组成。Goal 解析器会将 "G" 解释为字符串 "G"、"()" 解释为字符串 "o" ，"(al)" 解释为字符串 "al" 。然后，按原顺序将经解释得到的字符串连接成一个字符串。

给你字符串 command ，返回 Goal 解析器 对 command 的解释结果。



示例 1：

	输入：command = "G()(al)"
	输出："Goal"
	解释：Goal 解析器解释命令的步骤如下所示：
	G -> G
	() -> o
	(al) -> al
	最后连接得到的结果是 "Goal"

示例 2：

	输入：command = "G()()()()(al)"
	输出："Gooooal"

示例 3：

	输入：command = "(al)G(al)()()G"
	输出："alGalooG"



提示：

* 1 <= command.length <= 100
* command 由 "G"、"()" 和/或 "(al)" 按某种顺序组成
