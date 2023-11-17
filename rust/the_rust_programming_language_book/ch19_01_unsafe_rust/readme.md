### 不安全Rust
不会强制执行内存安全保证：被称为`不安全Rust(unsafe Rust)`

#### 不安全的超能力
可以通过`unsafe`关键字来切换到不安全Rust，接着可以开启一个新的存放不安全代码的块。这里有五类可以不安全Rust中进行而不能用于安全Rust的操作，它们称为"不安全的超能力":
* 解引用裸指针
* 调用不安全的函数或方法
* 访问或修改可变静态变量
* 实现不安全trait
* 访问union的字段

有一点很重要，unsafe并不会关闭借用检查器或禁用任何其他Rust安全检查: 如果在不安全代码中使用引用，它仍会被检查。unsafe关键字只是提供了那五个不会被编译器检查内存安全的功能。你仍然能在不安全块中获得某种程度的安全。
再者，unsafe不意味着块中的代码就一定是危险的或者必须导致内存安全问题：其意图在于作为开发者将会确保unsafe块中的代码以有效的方式访问内存

通过要求这五类操作必须位于标记为unsafe的块中，就能够知道任何与内存安全相关的错误必定位于unsafe块内。保持unsafe块尽可能小

#### 解引用祼指针
不安全Rust有两个被为`祼指针(raw pointers)`的类似于引用的新类型。和引用一亲，裸指针是不可变或可变的，分别写作`*const T `和`*mut T`。这里的星号不是解引用运算符；它是类型名称的一部分。在祼指针的上下文中，不可变意味着指针解引用之后不能直接赋值

祼指针与引用和智能指针的区别：
* 允许忽略借用规则，可以同时拥有不可变和可变的指针，或多个指向相同位置的可变指针
* 不保证指向有效的内存
* 允许为空
* 不能实现任何自动清理功能
通过去掉Rust强加的保证，你可以放弃安全保证以换取性能或使用另一个语言或硬件接口的能力，此时Rust的保证并不适用

从引用同时创建不可变和可变裸指针：
```Rust
let mut num = 5;

let r1 = &num as *const i32;
let r2 = &mut num as *mut i32;
```
可以在安全代码中创建裸指针，只是不能在不安全块之外解引用裸指针

这里使用`as`将不可变和可变引用强转为对应的裸指针类型。

创建一个指向任意内存地址的裸指针:
```Rust
let address = 0x012345usize;
let r = address as *const i32;
```

在unsafe块中解引用裸指针：
```Rust
let mut num = 5;

let r1 = &num as *const i32;
let r2 = &mut num as *mut i32;

unsafe {
    println!("r1 is: {}", *r1);
    println!("r2 is: {}", *r2);
}
```
创建一个指针不造成任何危险；只有当访问其指向的值时才有可能遇到无效的值。

同时创建num的不可变和可变引用，将无法通过编译，因为Rust的所有权规则不允许在拥有任何不可变引用的同时再创建一个可变引用。通过裸指针，就能够同时创建同一地址的不可变指针和可变指针，若通过可变指针修改数据，则可能潜在造成数据竞争

> 一个主要的应用场景便是调用C代码接口，另一个场景是构建借用检查无法理解的安全抽象

#### 调用不安全函或方法
不安全函数和方法和常规函数方法十分类似，除了其开头有一个额外的unsafe。
在此上下文中，关键字unsafe表示该函数具有调用时需要满足的要求，而Rust不会保证满足这些要求。
通过在unsafe块中调用不安全函数，表明已经阅读此函数的文档并对其是否满足函数自身的契约负责

没有任何操作的不安全函数
```Rust
unsafe fn dangerous() {}

unsafe {
    dangerous();
}
```

#### 创建不安全代码的安全抽象
仅仅因为函数包含不安全代码并不意味着整个函数都需要标记为不安全的
事实上，将不安全代码封装进安全函数是一个常见的抽象
```Rust
let mut v = vec![1, 2, 3, 4, 5, 6];

let r = &mut v[..];
let (a, b) = r.split_at_mut(3);

assert_eq!(a, &mut [1, 2, 3]);
assert_eq!(b, &mut [4, 5, 6]);
```
这个函数无法只通过安全Rust实现。
```Rust
fn split_at_mut(slice: &mut [i32], mid: usize) -> (&mut [i32], &mut [i32]) {
    let len = slice.len();

    assert!(mid <= len);

    // error: cannot borrow `*slice` as mutable more than once at a time
    // (&mut slice[..mid], &mut slice[mid..])

    let ptr = slice.as_mut_ptr();
    unsafe {
        (slice::from_raw_parts_mut(ptr, mid), 
         slice::from_raw_parts_mut(ptr.add(mid), len - mid))
    }

}
```
注意无需将split_at_mut函数的结果标记为unsafe，并可以估安全Rust中调用此函数。我们创建一个不安全代码的安全抽象，其代码以一种安全的方式使用了unsafe代码，因为其只从这个函数访问的数据中创建了有效的指针。

#### 使用extern 函数调用外部代码
有时Rust代码可能需要与其他语言编写的代码交互。为些Rust有一个关键字，extern，有助于创建和使用`外部函数接口(Foreign Function interface, FFI)`。外部函数接口是一个编程语言用以定义函数的方式，其允许不同(外部)编程语言调用这些函数。
```Rust
extern "C" {
    fn abs(input: i32) -> i32;
}

fn main() {
    unsafe {
        println!("Absolute value of -3 according to C: {}", abs(-3));
    }
}
```
在extern "C"块中，列出了我们希望能够调用的另一个语言中的外部函数的签名和名称。"C"部分定义了外部函数所使用的`应用二进制接口(application binary interface, ABI)` -- ABI定义了如何在汇编语言层面调用此函数。"C"ABI是最常见的，并遵循C编程语言的ABI。

#### 从其它语言调用Rust函数
也可以使用extern来创建一个允许其他语言调用Rust函数的接口。
不同于extern块，就在fn关键字之前增加extern关键字并指定所用到的API。还需增加#[no_mangle]标注来告诉Rust编译器不要mangle此函数的名称。
```Rust
#[no_mangle]
pub extern "C" fn call_from_c() {
    println!("Just called a Rust function from C!");
}
```
extern的使用无需unsafe

#### 访问或修改可变静态变量
**全局变量(global variables)**，Rust确实支持，不过这对于Rust的所有权规则来说是有问题的。
如果有两个线程访问相同的可变全局变量，则可能会造成数据竞争。

全局变量在Rust中被称为`静态(static)`变量。
```Rust
static HELLO_WORLD: &str = "Hello, world!";

static mut COUNTER: u32 = 0;

fn add_to_count(inc: u32) {
    unsafe {
        COUNTER += inc;
    }
}

fn main() {
    println!("name is: {}", HELLO_WORLD);

    add_to_count(3);
    unsafe {
        println!("COUNTER: {}", COUNTER);
    }
}
```
静态变量只能储存拥有'static生命周期的引用，这意味着Rust编译器可以自己计算出其生命周期而无需显式标注。访问不可变静态变量是安全的。

常量与不可变静态变量可能看起来很类似，不过一个微妙的区别是静态变量中的值有一个固定的内存地址。使用这个值总是会访问相同的地址。另一方面，常量则允许在任何被用到的时候复制其数据。

常量与静态变量的另一个区别在于静态变量可以是可变的。访问和修改可变静态变量都是`不安全`的。
拥有可以全局访问的可变数据，难以保存不存在数据竞争，这就是为何Rust认为可变静态变量是不安全的。

#### 实现不安全trait
unsafe的另一个操作用例是实现不安全trait。当trait中至少有一个方法中包含编译器无法验证的不变式(invariant)时trait是不安全的。可以在trait之前增加unsafe关键字将trait声明为unsafe，同时trait的实现也必须标记为unsafe. 
```Rust
unsafe trait Foo {

}

unsafe impl Foo for i32 {

}
```
通过`unsafe impl`，我们承诺将保证编译器所不能验证的不变量。

Rust不能验证我们的类型保证可以安全的跨线程发送或在多线程间访问，所以需要我们自己进行检查并通过unsafe表明。

### 访问联合体中的字段
仅适用于`unsafe`的最后一个操作是访问`联合体`中的字段，union和struct类似，但是在一个实例中同时只能使用一个声明的字段。联合体主要用于和C代码中的联合体交互。访问联合体的字段是不安全的，因为Rust无法保证当前存储的联合体实例中数据的类型。

#### 何时使用不安全代码
使用unsafe来进行这五个操作(超能力)之一是没有问题的，甚至是不需要深思熟虑的，不过使得unsafe代码正确也实属不易，因为编译器不能帮助保证内存安全。当有理由使用unsafe代码时，是可以这么做的，通过使用显式的unsafe标注可以更容易地在错误发生时追踪问题的源头

