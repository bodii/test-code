### 1108. Defanging an IP Address
Given a valid (IPv4) IP address, return a defanged version of that IP address.

A defanged IP address replaces every period "." with "[.]".



Example 1:

	Input: address = "1.1.1.1"
	Output: "1[.]1[.]1[.]1"

Example 2:

	Input: address = "255.100.50.0"
	Output: "255[.]100[.]50[.]0"



Constraints:

* The given address is a valid IPv4 address.

----

### 1108. IP 地址无效化
给你一个有效的 IPv4 地址 address，返回这个 IP 地址的无效化版本。

所谓无效化 IP 地址，其实就是用 "[.]" 代替了每个 "."。



示例 1：

	输入：address = "1.1.1.1"
	输出："1[.]1[.]1[.]1"

示例 2：

	输入：address = "255.100.50.0"
	输出："255[.]100[.]50[.]0"



提示：

* 给出的 address 是一个有效的 IPv4 地址
