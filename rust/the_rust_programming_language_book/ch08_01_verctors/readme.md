### vector 用来储存一系列的值
第一个类型是Vec<T>, 也被称为vector.
vector允许在一个单独的数据结构中储存多个值，所有值在内存中彼此相邻排列。
vector只能储存相同类型的值。
使用场景: 例如文件中的文本行或购物车中商品的价格

### 新建vector
```Rust
// 新建一个空的vector来储存i32类型的值
let v: Vec<i32> = Vec::new();
```
在实际的代码中，一旦插入值Rust就可以推断出想要存放的类型，所以很少会需要类型标注。
用宏来创建一个新的Vec:
```Rust
// 新建一个包含初始值的vector
let v = vec![1, 2, 3];
```

### 更新vector
```Rust
let mut v = Vec::new();

v.push(5);
v.push(6);
v.push(7);
v.push(8);
```

### 丢弃vector时也会丢弃其所有元素
```Rust
{
  let v = vec![1, 2, 3, 4, 5];
  // 处理变量v
} // 这里v离开作用域并被丢弃
```
当vector被丢弃时，所有其内容也会被丢弃，这意味着这里它包含的整数将被清理。

### 读取vector的元素
有两种方法引用vector中储存的值：索引语法或者get方法：
```Rust
let v = vec![1, 2, 3, 4, 5];

let third: &i32 = &v[2];
println!("The third element is {}", third);

match v.get(2) {
  Some(third) => println!("The third element is {}", third),
  None => println!("There is no third element."),
}
```
索引是从0开始的。
其次，这两个不同的获取第三个元素的方式分别为：使用&和[]返回一个引用；或者使用get方法以索引作为参数来返回一个Option<&T>。 

```Rust
let mut v = vec![1, 2, 3, 4, 5];

// error
let first = &v[0];

v.push(6);

println!("The first element is: {}", first);
```
不能这么做，原因是由于vector的工作方式:
在vector的结尾增加新元素时，在没有足够空间将所有元素依次相邻存放的情况下，可能会要求分配新内存并将老的元素拷贝到新的空间中。这时，第一个元素的引用就指向了被释放的内存。借用规则阻止程序陷入这种状况。

### 遍历vector中的元素
如果想要依次访问vector中的每一个元素，我们可以遍历其所有的元素而无需通过索引一次一个的访问. 
```Rust
let v = vec![1, 2, 3, 4, 5];
for i in &v {
  println("{}", i)
}
```
遍历vector中元素的可变引用:
```Rust
let mut v = vec![100, 32, 57];
for i in &mut v {
    *i += 50;
}
```
> 为了修改可变引用所指向的值，在使用+=运算符之前必须使用解引用运算符(*)获取i中的值。


### 使用枚举来储存多种类型
枚举的成员被定义为相同的枚举类型，所以当需要在vector中储存不同类型值时，我们可以定义并使用一个枚举

如果在编写程序时不能确切无遗地知道运行时会储存进vector的所有类型，枚举技术就行不通了。相反，你可以使用trait对象。

