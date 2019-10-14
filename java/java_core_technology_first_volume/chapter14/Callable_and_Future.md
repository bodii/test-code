#### Callable与Future
Runnable封装一个异步运行的任务，可以把它想象成为一个没有参数和返回值的异步方法。Callable与
Runnable类似，但是有返回值。Callable接口是一个参数化的类型，只有一个方法call。
```java
public interface Callable<V> {
	V call() throws Exception;
}
```
类型参数是返回值的类型。 
例如， Callable<Integer>表示一个最终返回Integer对象的异步计算。

Future保存异步计算的结果。可以启动一个计算，将Future对象交给某个线程，然后忘掉它。
Future对象的所有者在结果计算好之后就可以获得它。

Future接口具有下面的方法:
```java
public interface Future<V> {
	V get() throws ...;
	V get(long timeout, TimeUnit unit) throws ...;
	void cancel(boolean mayInterrupt);
	boolean isCanceled();
	boolean isDone();
}
```
第一个get方法的调用被阻塞，直到计算完成。如果在计算完成之前，第二个方法的调用超时，抛出一
个TimeoutException异常。如果运行该计算的线程被中断，两个方法都将抛出InterruptedException。
如果计算已经完成，那么get方法立即返回。
如果计算还在进行，isDone方法返回false; 如果完成了，则返回true.
可以用cancel方法取消该计算。如果计算还没有开始，它被取消且不再开始。如果计算处于运行之中，
那么如果mayInterrupt参数为true, 它就被中断。
FutrueTask包装器是一种非常便利的机制，可将Callable转换成Future和Runnable, 它同时实现二者的
接口。例如:
```java
Callable<Integer> myComputation = ...;
FutureTask<Integer> task = new FutureTask<Integer>(myComputation);
Thread t = new Thread(task);
t.start();
...
Integer result = task.get(); // it's Future
```

