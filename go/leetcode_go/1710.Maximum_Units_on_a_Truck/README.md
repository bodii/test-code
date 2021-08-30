### 1710. Maximum Units on a Truck
You are assigned to put some amount of boxes onto one truck. You are given a 2D array boxTypes, where boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi]:

* numberOfBoxesi is the number of boxes of type i.
* numberOfUnitsPerBoxi is the number of units in each box of the type i.

You are also given an integer truckSize, which is the maximum number of boxes that can be put on the truck. You can choose any boxes to put on the truck as long as the number of boxes does not exceed truckSize.

Return the maximum total number of units that can be put on the truck.



Example 1:

	Input: boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
	Output: 8
	Explanation: There are:
	- 1 box of the first type that contains 3 units.
	- 2 boxes of the second type that contain 2 units each.
	- 3 boxes of the third type that contain 1 unit each.
	You can take all the boxes of the first and second types, and one box of the third type.
	The total number of units will be = (1 * 3) + (2 * 2) + (1 * 1) = 8.

Example 2:

	Input: boxTypes = [[5,10],[2,5],[4,7],[3,9]], truckSize = 10
	Output: 91



Constraints:

* 1 <= boxTypes.length <= 1000
* 1 <= numberOfBoxesi, numberOfUnitsPerBoxi <= 1000
* 1 <= truckSize <= 106

----

### 1710. 卡车上的最大单元数
请你将一些箱子装在 一辆卡车 上。给你一个二维数组 boxTypes ，其中 boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi] ：

* numberOfBoxesi 是类型 i 的箱子的数量。
* numberOfUnitsPerBoxi 是类型 i 每个箱子可以装载的单元数量。

整数 truckSize 表示卡车上可以装载 箱子 的 最大数量 。只要箱子数量不超过 truckSize ，你就可以选择任意箱子装到卡车上。

返回卡车可以装载 单元 的 最大 总数。



示例 1：

	输入：boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
	输出：8
	解释：箱子的情况如下：
	- 1 个第一类的箱子，里面含 3 个单元。
	- 2 个第二类的箱子，每个里面含 2 个单元。
	- 3 个第三类的箱子，每个里面含 1 个单元。
	可以选择第一类和第二类的所有箱子，以及第三类的一个箱子。
	单元总数 = (1 * 3) + (2 * 2) + (1 * 1) = 8

示例 2：

	输入：boxTypes = [[5,10],[2,5],[4,7],[3,9]], truckSize = 10
	输出：91



提示：

* 1 <= boxTypes.length <= 1000
* 1 <= numberOfBoxesi, numberOfUnitsPerBoxi <= 1000
* 1 <= truckSize <= 106

