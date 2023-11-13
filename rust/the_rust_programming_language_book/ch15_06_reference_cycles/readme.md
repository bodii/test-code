### 引用循环与内存泄漏
Rust的内存安全性保证使其难以意外地制造永远也不会被清理的内存(被称为`内存泄漏(memory leak)`)，但并不是不可能。与在编译时拒绝数据竞争不同，Rust并不保证完全地避免内存泄漏，这意味着内存泄漏在Rust被认为是内存安全的。这一点可以通过Rc<T>和RefCell<T>看出：创建引用循环的可能性是存在的。这会造成内存泄漏，因为每一项的引用计数永远也到不了0, 其值也永远不会被丢弃。

### 避免引用循环：将Rc<T>变为Weak<T>
调用Rc::downgrade时会得到Weak<T>类型的智能指针。不同于将Rc<T>实例的strong_count加1，调用Rc::downgrade会将weak_count加1。
Rc<T>类型使用weak_count来记录其存在多少个Weak<T>引用，类似于strong_count。其区别在于weak_count无需计数0就能使Rc<T>实例被清理。

强引用代表如何共享Rc<T>实例的所有权，但弱引用并不属于所有权关系。他们不会造成引用循环，因为任何弱引用的循环会在其相关的强引用计数为0时被打断。

因为Weak<T>引用的值可能已经被丢弃了，为了使用Weak<T>所指向的值，我们必须确保其值仍然有效。为此可以调用Weak<T>实例的upgrade方法，这会返回Option<Rc<T>>。如果Rc<T>值还未被丢弃，则结果是Some;如果Rc<T>已被丢弃，则结果是None。因为upgrade返回一个Option<T>，我们确信Rust会处理Some和None的情况，所以它不会返回非法指针。

