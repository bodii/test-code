#### 使用Swing工作线程
SwingWorker类使后台任务的实现不那么繁琐。
每当要在工作器线程中做一些工作时，构建一个新的工作器(每一个工作器对象只能被使用一次)。
然后调用execute方法。典型的方式是在事件分配线程中调用execute,但没有这样的需求。

在doInBackground方法中，读取一个文件，每次一行。在读取每一行之后，调用publish方法发布行号
和当前行的文本。
```java
@Override public StringBuilder doInBackground() throws IOException, InterruptedException {
	int lineNumber = 0;
	Scanner in = new Scanner(new FileInputStream(file), "UTF-8");
	while (in.hasNextLine()) {
		String line = in.nextLine();
		lineNumber++;
		text.append(line).append("\n");
		ProgressData data = new ProgressData();
		data.number = lineNumber;
		data.line = line;
		publish(data);
		Thread.sleep(1); // to test cancellation; no need to do this in your programs
	}
	return text;
} 
```
在读取每一行之后休眠1毫秒，以便不使用重读就可以检测取消动作，但是，不要使用休眠来减慢程序的
执行速度。如果对这一行加注释，会发现The Count of Monte Cristo的加载相当快，只有几批用户接
口更新。

从工作器线程来更新文本区可以使这个程序的处理相当顺畅，但是，对大多数Swing组件来说不可能做到
这一点。这里，给出一个通用的方法，其中所有组件的更新都出现在事件分配线程中。
在这个process方法中，忽略除最后一行行号之外的所有行号，然后，我们把所有的行拼接在一起用于
文本区的一次更新。
```java
@Override public void process(List<ProgressData> data) {
	if (isCancelled()) return;
	StringBuilder b = new StringBuilder();
	statusLine.setText("" + data.get(data.size() -1).number);
	for (ProgressData d : data) b.append(d.line).append("\m");
	testArea.append(b.toString());
}
```
