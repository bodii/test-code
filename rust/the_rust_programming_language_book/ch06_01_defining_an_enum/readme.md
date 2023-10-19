### 枚举
可以将任意类型的数据放入枚举成员中：例如字符串、数字类型或者结构体。甚至可以包含另一个枚举

### Option 枚举和其相对于空值的优势
空值尝试表达的概念仍然是有意义的： 空值是一个因为某种原因目前无效或缺失的值.
> Rust 并没有空值，不过它可以编码存在或不存在概念的枚举。
```Rust
enum Option<T> {
    Some(T),
    None,
}
```
一些包含数字类型和字符串类型Option值的例子:
```Rust
let some_number = Some(5);
let some_string = Some("a string");

let absent_number: Option<i32> = None;
```
> 如果使用None而不是Some，需要告诉Rust Option<T> 是什么类型的，因为编译器只通过None值无法推断出Some成员保存的值的类型。

当有一个Somew值时，我们就知道存在一个值，而这个值保存在Some中。
当有一个None值时，在某种意义上，它跟空值具有相同的意义： 并没有一个有效的值。

在对Option<T>进行T运算之前必须将其转换成T。