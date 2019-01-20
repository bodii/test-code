#### typedef
typedef工具是一个高级数据特性，利用typedef可以为某一类型自定义名称。
与#define类似，但是两者有3处不同：
1. 与#define不同，typedef创建的符号名只受限于类型，不能用于值。
2. typedef由编译器解释，不是预处理器。
3. 在其受限范围内，typedef比#define更灵活。

假设要用BYTE表示1字节的数组，只需:
```c
typedef unsigned char BYTE;

// 随后，便可使用BYTE来定义变量：
BYTE x, y[10], * z;
```
该定义的作用域取决于typedef定义所在的位置。如果定义在函数中，就具有局部
作用域，受限于定义所在的函数。如果定义在函数外面，就具有文件作用域。

通常，typedef定义中用大写字母表示被定义的名称，以提醒用户这个类型名实际
上是一个符号缩写。当然，也可以用小写：
```c
typedef unsigned char byte;
```
使用typedef还能提高程序的可移值性。
例如，我们之前提到的sizeof运算符的返回类型：size_t类型，time()函数的返回
类型：time_t类型。C标准规定sizeof和time()返回整数类型，但是让实现来决定具
体是什么整数类型。
```c
time_t time(time_t *);
```
time_t在一个系统中是unsigned long,在另一个系统中可以是unsigned long long.

typedef的一些特性与#define的功能重合：
```c
#define BYTE usigned char
```
这使预处理器用BYTE替换unsigned char.但是也有#define没有的功能：
```c
typedef char * STRING;
```
没有typedef关键字，编译器将把STRING识别为一个指向char的指针变量。有了typedef
关键字，编译器则把STRING解释成一个类型的标识符，该类型是指向char的指针：
```c
STRING name, sign;

// 相当于：
char * name, * sign;

```
但如果：
```c
#define STRING char *

STRING name, sign;
// 将被译成
char * name, sign;
```
这将导致只有name才是指针。

还可以把typedef用于结构:
```c
typedef struct complex {
	float real;
	float image;
} COMPLEX;
```
然后便可使用COMPLEX类型代替complex结构来表示复数。

使用typedef的第1个原因是：为经常出现的类型创建一个方便、易识别的类型名。
用typedef来命名一个结构类型时，可以省略该结构的标签：
```c
typedef struct { double x, double y } rect;

// 假设这样使用typedef定义的类型名：
rect r1 = { 3.0, 6.0 };
rect r2;

// 以上代码将被翻译成：
struct { double x; double y; } r1 = { 3.0, 6.0 };
struct { double x; double y; } r2;
r2 = r1;
```

使用typedef的第2个原因：typedef常用于给复杂的类型命名。
```c
typedef char (*FRPTC())[5];
```
把FRPTC声明为一个函数类型，该函数返回一个指针，该指针指向内含5个char类型元素的数组。

typedef并没有创建任何新类型，它只是为某个已存在的类型增加了一个方便使用的标签。



