### 使用Drop Trait运行清理代码
对于智能指针模式来说第二个重要的trait是Drop，其允许我们在值要离开作用域时执行一些代码。
例如，Box<T>自定义了Drop用来释放box所指向的堆空间。

在Rust中，可以指定每当值离开作用域时被执行的代码，编译器会自动插入这些代码。于是就不需要在程序中到处编写在实例结束时清理这些变量的代码--而且还不会泄漏资源。

指定在值离开作用域时应该执行的代码的方式是实现`Drop trait`。
Drop trait 要求实现一个叫做`drop`的方法，它获取一个self的可变引用。

Drop trait 包含在`prelude`中，所以无需导入它。
Rust会调用放置于drop方法中的代码，注意`无需显式调用drop方法`

变量以被创建时相反的顺序被丢弃。

### 通过std::mem::drop提早丢弃值
一个例子是当使用智能指针管理锁时； 
当希望在作用域结束之前就强制释放变量的话，应该使用的是标准库std::mem::drop。

std::mem::drop 位于prelude。

Drop trait实现中指定的代码可以用于许多方面，来使得清理变得方便和安全：比如可以用其创建自己的内存分配器
