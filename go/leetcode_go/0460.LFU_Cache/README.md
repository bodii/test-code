### 460. LFU Cache
Design and implement a data structure for a Least Frequently Used (LFU) cache.

Implement the LFUCache class:
* LFUCache(int capacity) Initializes the object with the capacity of the data structure.
* int get(int key) Gets the value of the key if the key exists in the cache. Otherwise, returns -1.
* void put(int key, int value) Update the value of the key if present, or inserts the key if not already present. When the cache reaches its capacity, it should invalidate and remove the least frequently used key before inserting a new item. For this problem, when there is a tie (i.e., two or more keys with the same frequency), the least recently used key would be invalidated.

To determine the least frequently used key, a use counter is maintained for each key in the cache. The key with the smallest use counter is the least frequently used key.

When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation). The use counter for a key in the cache is incremented either a get or put operation is called on it.



Example 1:

	Input
	["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
	[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
	Output
	[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

	Explanation
	// cnt(x) = the use counter for key x
	// cache=[] will show the last used order for tiebreakers (leftmost element is  most recent)
	LFUCache lfu = new LFUCache(2);
	lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
	lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
	lfu.get(1);      // return 1
					 // cache=[1,2], cnt(2)=1, cnt(1)=2
	lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
					 // cache=[3,1], cnt(3)=1, cnt(1)=2
	lfu.get(2);      // return -1 (not found)
	lfu.get(3);      // return 3
					 // cache=[3,1], cnt(3)=2, cnt(1)=2
	lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
					 // cache=[4,3], cnt(4)=1, cnt(3)=2
	lfu.get(1);      // return -1 (not found)
	lfu.get(3);      // return 3
					 // cache=[3,4], cnt(4)=1, cnt(3)=3
	lfu.get(4);      // return 4
                 // cache=[3,4], cnt(4)=2, cnt(3)=3



Constraints:
* 0 <= capacity, key, value <= 104
* At most 105 calls will be made to get and put.


Follow up: Could you do both operations in O(1) time complexity?


--------------------------------------------

### 460. LFU 缓存
请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。

实现 LFUCache 类：
* LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
* int get(int key) - 如果键存在于缓存中，则获取键的值，否则返回 -1。
* void put(int key, int value) - 如果键已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量时，则应该在插入新项之前，使最不经常使用的项无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。

注意「项的使用次数」就是自插入该项以来对其调用 get 和 put 函数的次数之和。使用次数会在对应项被移除后置为 0 。

为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。



示例：

	输入：
	["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
	[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
	输出：
	[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

	解释：
	// cnt(x) = 键 x 的使用计数
	// cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
	LFUCache lFUCache = new LFUCache(2);
	lFUCache.put(1, 1);   // cache=[1,_], cnt(1)=1
	lFUCache.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
	lFUCache.get(1);      // 返回 1
						  // cache=[1,2], cnt(2)=1, cnt(1)=2
	lFUCache.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
						  // cache=[3,1], cnt(3)=1, cnt(1)=2
	lFUCache.get(2);      // 返回 -1（未找到）
	lFUCache.get(3);      // 返回 3
						  // cache=[3,1], cnt(3)=2, cnt(1)=2
	lFUCache.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
						  // cache=[4,3], cnt(4)=1, cnt(3)=2
	lFUCache.get(1);      // 返回 -1（未找到）
	lFUCache.get(3);      // 返回 3
						  // cache=[3,4], cnt(4)=1, cnt(3)=3
	lFUCache.get(4);      // 返回 4
						  // cache=[3,4], cnt(4)=2, cnt(3)=3



提示：
* 0 <= capacity, key, value <= 104
* 最多调用 105 次 get 和 put 方法



进阶：你可以为这两种操作设计时间复杂度为 O(1) 的实现吗？
