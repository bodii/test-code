#### 命令行参数
Java程序在启动时会接收传入的命令行参数(command-line argument).
传入的命令行参数会在程序启动后（即main方法开始执行之前),作为main方法的参数进行传递。

```java
class PrintArgs {
	public static void main(String[] args) {
		for (int i = 0; i < args.length; i++)
			System.out.println("args[" + i + "] = " + args[i]);
	}
}
```

```shell
# 运行PrintArgs程序
java PrintArgs Turbo NA DOHC
```
