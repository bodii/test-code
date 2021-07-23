### 1436. Destination City
You are given the array paths, where paths[i] = [cityA<sub>i</sub>, cityB<sub>i</sub>] means there exists a direct path going from cityA<sub>i</sub> to cityB<sub>i</sub>. Return the destination city, that is, the city without any path outgoing to another city.

It is guaranteed that the graph of paths forms a line without any loop, therefore, there will be exactly one destination city.



Example 1:

    Input: paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
    Output: "Sao Paulo"
    Explanation: Starting at "London" city you will reach "Sao Paulo" city which is the destination city. Your trip consist of: "London" -> "New York" -> "Lima" -> "Sao Paulo".

Example 2:

    Input: paths = [["B","C"],["D","B"],["C","A"]]
    Output: "A"
    Explanation: All possible trips are:
    "D" -> "B" -> "C" -> "A".
    "B" -> "C" -> "A".
    "C" -> "A".
    "A".
    Clearly the destination city is "A".

Example 3:

    Input: paths = [["A","Z"]]
    Output: "Z"



Constraints:

* 1 <= paths.length <= 100
* paths[i].length == 2
* 1 <= cityA<sub>i</sub>.length, cityB<sub>i</sub>.length <= 10
* cityA<sub>i</sub> != cityB<sub>i</sub>
* All strings consist of lowercase and uppercase English letters and the space character.

----

### 1436. 旅行终点站
给你一份旅游线路图，该线路图中的旅行线路用数组 paths 表示，其中 paths[i] = [cityA<sub>i</sub> , cityB<sub>i</sub> ] 表示该线路将会从 cityA<sub>i</sub>  直接前往 cityB<sub>i</sub>  。请你找出这次旅行的终点站，即没有任何可以通往其他城市的线路的城市。

题目数据保证线路图会形成一条不存在循环的线路，因此只会有一个旅行终点站。

 

示例 1：

    输入：paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
    输出："Sao Paulo" 
    解释：从 "London" 出发，最后抵达终点站 "Sao Paulo" 。本次旅行的路线是 "London" -> "New York" -> "Lima" -> "Sao Paulo" 。

示例 2：

    输入：paths = [["B","C"],["D","B"],["C","A"]]
    输出："A"
    解释：所有可能的线路是：
    "D" -> "B" -> "C" -> "A". 
    "B" -> "C" -> "A". 
    "C" -> "A". 
    "A". 
    显然，旅行终点站是 "A" 。

示例 3：

    输入：paths = [["A","Z"]]
    输出："Z"

提示：

* 1 <= paths.length <= 100
* paths[i].length == 2
* 1 <= cityA<sub>i</sub>.length, cityB<sub>i</sub>.length <= 10
* cityA<sub>i</sub> != cityB<sub>i</sub>
* 所有字符串均由大小写英文字母和空格字符组成。