#### 枚举类型
可以用枚举类型(enumerated type)声明符号名称来表示整型常量。使用
enum关键字,呆以创建一个新"类型"并指定它可具有的值（实际上，enum
常量是int类型，因此，只要能使用int类型的地方就可以使用枚举类型).
```c
enum spectrum { red, orange, yellow, green, blue, violet };
enum spectrum color;

int c;
color = blue;

if (color == yellow)
...

for(color = red; color <= violet; color++)
...
```
虽然枚举符是int类型，但是枚举变量可以是任意整数类型，前提是该整数类型
可以储存枚举常量。例如，spectrum的检举符范围是0~5,所以编译器可以用unsigned
char来表示color变量。
C枚举的一些特性并不适用于C++。
C允许枚举变量使用++运算符，但是C++标准不允许。


#### enum常量
如上例, blue和red是int类型的常量。
```c
printf("red = %d, orange = %d\n", red, orange);
```
// 其输出如下：
red = 0, orange = 1


#### 默认值
```c
enum kids { nippy, slats, kippy, nina, liz };
```
nina的值是3


#### 赋值
```c
enum levels { low = 100, medium = 500, high = 2000 };
```
```c
enum feline { cat, lynx = 10, puma, tiger };
```
cat的值是0(默认), lynx、puma和tiger的值分别是10, 11, 12


#### enum的用法
枚举类型的目的是为了提高程序的可读性和可维护性。


#### 共享名称空间
C语言使用名称空间(namespace)标识程序中的各部分，即通过名称来识别。
作用域是名称空间概念的一部分：两个不同作用域的同名量不冲突;
两个相同作用域的同名变量冲突。

名称空间是分类别的。在特定作用域中的结构标记、联合标记和枚举标记都共享相同的名称
空间，该名称空间与普通变量使用的空间不同。这意味着在相同作用域中变量和标记的名称
可以相同，不会引起冲突，但是不能在相同作用域中声明两个同名标签或同名变量.
```c
struct rect { double x; double y; };
int rect;  // 在C中不会产生冲突
```
尽管如此，以两种不同的方式使用相同的标识符会造成混乱。另外，C++不允许这样做，因为
它把标记名和变量名放在相同的名称空间中。

