#### 可完成Future
假设我们希望从一个Web页面抽取所有链接来建立一个网络爬虫。下面假设有这样一个方法：
```java
public void CompletableFutrue<String> readPage(URL url)
```
Web页面可用时这会生成这个页面的文本。如果方法:
```java
public static List<URL> getLinks(String page)
```
生成一个HTML页面中的URL，可以调度当页面可用时再调用这个方法:
```java
CompletableFuture<String> contents = readPage(url);
CompletableFuture<List<URL>> links = contents.thenApply(Parser::getLinks);
```
thenApply方法不会阻塞。它会返回另一个future。第一个future完成时，其结果会提供给getLinks方法，
这个方法的返回值就是最终的结果。
利用可完成future,可以指定你希望做什么，以及希望以什么顺序执行这些工作。当然，这不会立即发生
，不过重要的是所有代码都放在一处。
