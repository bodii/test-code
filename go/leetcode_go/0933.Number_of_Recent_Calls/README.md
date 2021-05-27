### 933. Number of Recent Calls
You have a RecentCounter class which counts the number of recent requests within a certain time frame.

Implement the RecentCounter class:

* RecentCounter() Initializes the counter with zero recent requests.
* int ping(int t) Adds a new request at time t, where t represents some time in milliseconds, and returns the number of requests that has happened in the past 3000 milliseconds (including the new request). Specifically, return the number of requests that have happened in the inclusive range [t - 3000, t].

It is guaranteed that every call to ping uses a strictly larger value of t than the previous call.



Example 1:

	Input
	["RecentCounter", "ping", "ping", "ping", "ping"]
	[[], [1], [100], [3001], [3002]]
	Output
	[null, 1, 2, 3, 3]

	Explanation
	RecentCounter recentCounter = new RecentCounter();
	recentCounter.ping(1);     // requests = [1], range is [-2999,1], return 1
	recentCounter.ping(100);   // requests = [1, 100], range is [-2900,100], return 2
	recentCounter.ping(3001);  // requests = [1, 100, 3001], range is [1,3001], return 3
	recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002], range is [2,3002], return 3



Constraints:

* 1 <= t <= 10<sup>9</sup>
* Each test case will call ping with strictly increasing values of t.
* At most 10<sup>4</sup> calls will be made to ping.

----

### 933. 最近的请求次数
写一个 RecentCounter 类来计算特定时间范围内最近的请求。

请你实现 RecentCounter 类：

* RecentCounter() 初始化计数器，请求数为 0 。
* int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 [t-3000, t] 内发生的请求数。

保证 每次对 ping 的调用都使用比之前更大的 t 值。



示例：

	输入：
	["RecentCounter", "ping", "ping", "ping", "ping"]
	[[], [1], [100], [3001], [3002]]
	输出：
	[null, 1, 2, 3, 3]

	解释：
	RecentCounter recentCounter = new RecentCounter();
	recentCounter.ping(1);     // requests = [1]，范围是 [-2999,1]，返回 1
	recentCounter.ping(100);   // requests = [1, 100]，范围是 [-2900,100]，返回 2
	recentCounter.ping(3001);  // requests = [1, 100, 3001]，范围是 [1,3001]，返回 3
	recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002]，范围是 [2,3002]，返回 3



提示：

* 1 <= t <= 10<sup>9</sup>
* 保证每次对 ping 调用所使用的 t 值都 严格递增
* 至多调用 ping 方法 10<sup>4</sup> 次
