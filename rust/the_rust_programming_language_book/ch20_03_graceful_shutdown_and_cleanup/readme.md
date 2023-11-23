### 优雅停机与清理
为ThreadPool实现Drop trait对线程池中的每一个线程调用join，这样这些线程将会执行完他们的请求。接着会为ThreadPool实现一个告诉线程他们应该停止接收新请求并结束的方式。

#### 为ThreadPool实现Drop trait

