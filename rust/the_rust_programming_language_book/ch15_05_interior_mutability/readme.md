### RefCell<T>和内部可变性模式
`内部可变性(Interior mutability)` 是Rust中的一个设计模式，它允许你即使在有不可变引用时也可以改变数据，这通常是借用规则。
为了改变数据，该模式在数据结构中使用unsafe代码来模糊Rust通常的可变性和借用规则。

对于引用和Box<T>，借用规则的不可变性作用于编译时。对于RefCell<T>，这些不可变性作用于`运行时`。对于引用，如果违反这些规则，会得到一个编译错误。而对于RefCell<T>，如果违反这些规则会panic并退出。

RefCell<T>只能用于单线程场景。如果尝试在多线程上下文中使用RefCell<T>，会得到一个编译错误。

如下为选择Box<T>, Rc<T>或RefCell<T>的理由:
* Rc<T> 允许相同数据有多个所有者； Bxo<T>和RefCell有单一所有者。
* Box<T>允许在编译时执行不可变或可变借用检查；Rc<T>仅允许在编译时执行不可变借用检查；RefCell<T>允许在运行时执行不可变或可变用借用检查。
* 因为RefCell<T>允许在运行时执行可变借用检查，所以我们可以在即便RefCell<T>自身是不可变的情况下修改其内部的值。

在不可变值内部改变值变是`内部可变性`模式。

### 内部可变性：不可变值的可变借用
借用规则的一个推论是当有一个不可变值时，不能可变地借用它。
```Rust
fn main() {
    let x = 5;
    // error
    let y = &mut x;
}
```
特定情况下，令一个值在其方法内部能够修改自身，而在其他代码中仍视为不可变，是很有用的。

### RefCell<T> 在运行时记录使用
当创建不可变和可变引用时，我们分别使用&和&mut语法。对于RefCell<T>来说，则是borrow和borrow_mut方法，这属于RefCell<T>安全API的一部分。borrow方法返回Ref<T>类型的智能指针，borrow_mut方法返回RefMut类型的智能指针。这两个类型都实现了Deref，所以可以当作常规引用对待。

RefCell<T>记录当前有多少个活动的Ref<T>和RefMut<T>智能指针。每次调用borrow，RefCell<T>将活动的不可变借用计数加一。当Ref<T>值离开作用域时，不可变借用计数减一。就像编译时借用规则一样，RefCell<T>在任何时候只允许有多个不可变借用或一个不可变借用。

```Rust
impl Messenger for MockMessenger {
    fn send(&self, message: &str) {
        // panic: already borrowed: BorrowMutError
        let mut one_borrow = self.sent_messages.borrow_mut();
        let mut two_borrow = self.sent_messages.borrow_mut();

        one_borrow.push(String::from(message));
        two_borrow.push(String::from(message));
    }
}
```
在运行时捕获借用错误而不是编译时意味着将会在开发过程的后期才会发现错误，甚至有可能发布到生产环境才发现； 还会因为在运行时而不是编译时记录借用而导致少量的运行时性能惩罚。

### 结合Rc<T>和RefCell<T>来拥有多个可变数据所有者
RefCell<T>的一个常见用法是与Rc<T>结合。
Rc<T>允许对相同数据有多个所有者，不过只能提供数据的不可变访问。如果有一个储存了RefCell<T>的Rc<T>的话，就可以得到有多个所有者并且可以修改的值了

标准库中也有其他提供内部可变性的类型，比如Cell<T>，它类似RefCell<T>但有一点除外：它并非提供内部值的引用，而是把值拷贝进和拷贝出Cell<T>。还有Mutex<T>，其提供线程间安全的内部可变性
