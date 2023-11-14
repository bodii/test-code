### 使用Sync 和 Send trait的可扩展并发
有两个并发概念是内嵌于语言中的: `std::marker`中的`Sync`和`Send trait`

### 通过Send允许在线程间转移所有权
Send标记trait表明类型的所有权可以在线程间传递。
几乎所有的Rust类型都是Send的，不过有一些例外，包括Rc<T>。

### Sync允许多线程访问
Sync标记trait表明一个实现了Sync的类型可以安全的在多个线程中拥有其值的引用。
换一种方式来说，对于任意类型T，如果&T是Send的话T就是Sync的，这意味着其引用就可以安全的发送到另一个线程。类似于Send的情况，基本类型是Sync的，完全由Sync的类型组成的类型也是Sync的。
智能指针Rc<T>也不是Sync的。RefCell<T>和Cell<T>系列类型不是Sync的。RefCell<T>在运行时所进行的借用检查也不是线程安全的。

### 手动实现Send和Sync是不安全的
通常并不需要手动实现Send和Sync trait，因为由Send和Sync的类型组成的类型，自动就是Send和Sync的。因为他们是标记trait，甚至都不需要实现任何方法。他们只是用来强并发相关不可变性的。
