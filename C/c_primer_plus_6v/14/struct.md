#### 建立结构声明
创建的结构，每个部分都称为成员(member)或字段(field).

结构声明(structure declaration) 描述了一个结构的组织布局。
```c
struct book {
	char title[MAXTITL];
	char author[MAXAUTL];
	float value;
}
```
该声明并未创建实际的数据对象，只描述了该对象由什么组成。
有时，我们把结构声明称为模板，因为它勾勒出结构是如何储存数据的。

```c
struct book library;
```
这把library声明为一个使用book结构布局的结构变量。

结构的标记名是可选的。但是以程序示例中的方式建立结构时（在一处定义结构布局
在另一处定义实际的结构变量),必须使用标记。


#### 定义结构变量
结构有两层含义。一层含义是“结构布局”，另一层是创建一个结构变量。
结构布局告诉编译器如何表示数据，但它并未让编译器为数据分配空间。
程序中创建结构变量的一行是：
```c
struct book library;
```
编译器执行这行代码便创建了一个结构变量library.
编译器使用book模板为该变量分配空间: 一个内含MAXTITL个元素的char数组、一个内
含MAXAUTL个元素的char数组和一个float类型的变量。这些存储空间都与一个名为library
结合在一起。

可以定义两个struct book类型的变量，或者甚至是指向struct book类型结构的指针：
```c
struct book doyle, panshin, *ptbook;
```
结构变量doyle和panshin中都包含title、author和value部分。
指针ptbook可以指向doyle、panshin或任何其他book类型的结构变量。

就计算机而言，下面的声明：
```c
struct book library;

// 简化为:
struct book {
	char title[MAXTITL];
	char author[AXAUTL];
	float value;
} library;   /* 声明的右花括号后跟变量名 */

// 换言之，声明结构的过程和定义结构变量的过程可以组合成一个步骤：
struct {     /* 无结构标记 */
	char title[MAXTITL];
	char author[MAXAUTL];
	float value;
} library;
```
然而，如果打算多次使用结构模板，就要使用带标记的形式；或者，使用typedef.


#### 初始化结构
初始化变量和数组：
```c
int count = 0;
int fibo[7] = { 0, 1, 2, 3, 4 };
```
初始化一个结构变量与初始化数组的语法类似：
```c
struct book library = {
	"The Pious Pirate and the Devious Damsel",
	"Renee Vivotte",
	1.95
};
```


#### 访问结构成员
使用结构成员运算符——点(.)访问结构中的成员。
例如，library.vlaue即访问library的value部分。

本质上，.title、.author和.value的作用相当于book结构的下标.
如果需要一个地址，可使用&library.float。.比&的优先级高,因此这个表达式和&(library.float)
一样。


#### 结构的初始化器
C99和C11为结构提供了`指定初始化器(designated initializer)`, 其语法与数组的指定初始化器类
似。但是，结构的指定初始化器使用点运算符和成员名(而不是方括号和下标)标识特定的元素。
例如，只初始化book结构的value成员，可以这样做：
```c
struct book surprise = {.value = 10.99 };

struct book gift = { 
	.value = 25.99,
	.author = "James Broadfool",
	.title = "Rue for the Toad"
};
```
对特定成员的最后一次赋值才是它实际获得的值:
```c
struct book gift = {
	.value = 18.90,
	.author = "Philionna Pestle",
	0.25
};
赋给value的最终值是0.25。


#### 结构数组 和声明结构数组
声明结构数组：
```c
struct book library[MAXBKS];
```
以上代码把library声明为一个内含MAXBKS个元素的数组的。
数组的每个元素都有是一个book类型的数组。


#### 标识结构数组的成员
```c
library[0].value; /* 第1个数组元素与value相关联 */
```


#### 嵌套结构
在一个结构中包含另一个结构, 即嵌套结构。
```c
struct names {  // 第1个结构
	char first[LEN];
	char last[LEN];
	1027
};

struct guy {   // 第2个结构
	struct names handle;    // 嵌套结构
	char favfood[LEN];
	char job[LEN];
	float income;
};


struct guy fellow = {     // 初始化一个结构变量
	{ "Ewen", "Villard" },
	"grilled salmon",
	"personality coach",
	68112.00
};

printf("Hello, %s!\n", fellow.handle.first);
```
从左往右解释fellow.handle.first:
(fellow.handle).firstl
也就是说，找到fellow,然后找到fellow的handle的成员，再找到handle的first成员。




