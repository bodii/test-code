#### Fork-Join框架
有些应用使用了大量线程，但其中大多数是空闲的。举例来说，一个Web服务器可能会为每个连接分别
使用一个线程。另外一些应用可能对每个处理器内核分别使用一个线程，来完成计算密集型任务，如图
像或视频处理。
Java SE 7中新引入了fork-join框架，专门用来支持后一类应用。假设有一个处理任务，它可以很自然
地分解为子任务，如:
```java
if (problemSize < threshold)
	solve problem directly
else {
	break problem into subproblems
	recursively solve each subproblem
	combine the results
}
```
图像处理就是这样一例子。要增强一个图像，可以变换上半部分和下部部分。如果有足够多空闲的处理
器，这些操作可以并行运行。(除了分解为两部分外，还需要做一些额外的工作，不过这属于技术细节）

采用框架可用的一种方式完成这种递归计算，需要提供一个扩展RecursiveTask<T>的类，
或提供一个扩展RecursiveAction的类。再覆盖compute方法来生成并调用子任务，然后合并其结果。
```java
class Counter extends RecursiveTask<Integer> {
	protected Integer compute() {
		if (to - from < THRESHOLD) {
			solve problem directly
		}
		else {
			int mid = (from + to) / 2;
			Counter first = new Counter(values, from, mid, filter);
			Counter second = new Counter(values, mid, to, filter);
			invokeAll(first, second);
			return first.join() + second.join();
		}
	}
}
```
在这里，invokeAll方法接收到很多任务并阻塞，直到所有这些任务都已经完成。join方法将生成结果。
我们对每个子任务应用了join,并返回其总和。

