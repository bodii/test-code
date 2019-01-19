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


#### 指向结构的指针
使用指向结构的指针有四个好处：
1. 就像指向数组的指针比数组本身更容易操控一样，指向结构的指针通常比结构本身更容易
操控。
2. 在一些早期的C实现中，结构不能作为参数传递给函数，但是可以传递指向结构的指针。
3. 即使能传递一个结构，传递指针通常更有效率。
4. 一些用于表示数据的结构中包含指向其他结构的指针。


#### 声明和初始化结构指针
声明结构指针：
```c
struct guy * him;
```
首先是关键字struct,其次是结构标记guy,然后是一个星号(*),其后跟着指针名。
该声明并未创建一个新的结构，但是指针him现在可以指向任意现有的guy类型的结构。
```c
him = &barney;
```
和数组不同的是，结构名并不是结构的地址，因此要在结构名前面加上&运算符。
```c
him = &fellow[0];
```
fellow是一个结构数组，这意味着fellow[0]是一个结构。所以，要让him指向fellow[0].
```c
him++;
```
him加1相当于him指向的地址加84.在十六进制中,874 - 820 = 54(十六进制) = 84(十进制),
因为每个guy结构都占有84字节的内存:names.first占用20字节，names.last占用20字节，favfood
占用20字节，job占用20字节，income占用4字节。

在一些系统中，一个结构的大小可能大于它各成员大小之和。这是因为系统对数据进行校准的
过程中产生了一些"缝隙".
有一些系统必须把每个成员都放在偶数地址上，或4倍数的地址上。在这种系统中，结构的内部就存
在未使用的"缝隙".


#### 用指针访问成员
第1种方法，也是最常用的方法： 使用->运算符。
该运算符由一个连接号(-)后跟一个大于号(>)组成。

如果him == &barney, 那么him->income 即是barney.income
如果him == &fellow[0], 那么him->income即是fellow[0].income

这里，him是一个指针，him->income是该指针所指向结构的一个成员。


第2种方法是，以这个样的顺序指定结构成员的值：
如果him == &fellow[0], 那么*him == fellow[0], 因为&和*是一对互逆运算符。因此:
fellow[0].income == (*him).income
必须要使用圆括号，因为.运算符比*运算符的优先级高。

下面的关系恒成立：
barney.income == (*him).income == him->income   // 假设him == &barney


#### 结构和结构指针的选择
把指针作为参数有两个优点：
无论是以前不是现在的C实现都有能使用这种方法，而且执行起来很快，只需要传递一个地址。
缺点是无法保护数据。被调函数中的某些操作可能会意外原来结构中的数据。
可以使用const限定符。


#### 结构中的字符数组和字符指针
```c
#define LEN 20

struct names {
	char first[LEN];
	char last[LEN];
};

// 可以这样写：
struct pnames {
	char * first;
	char * last;
};
// 但，panmes结构变量中的指针应该只用来在程序中管理那些已分配和在别处分配的字符串
```


#### 结构、指针和malloc()
如果使用malloc()分配内存并使用指针储存该地址，那么在结构中使用指针处理字符串就比较
合理。


#### 复合字面量和结构(C99)
C99的复合字面量特性可用于结构和数组。
如果只需要一个临时结构值，复合字面量很好用。
```c
(struct book) { "The Idiot", "Fyodor Dostoyevsky", 6.99 };
```
还可以把复合字面量作为函数的参数。
如果函数接受一个结构，可以把复合字面量作为实际参数传递：
```c
struct rect { double x; double y; };
dobule rect_area(struct rect r) { return r.x * r.y; }

double area;
area = rect_area((struct rect) { 10.5, 20.0});

double rect_areap(struct rect * rp) { return rp->x * rp->y; }
area = rect_areap( &(struct rect) { 10.5, 20.0} );


#### 伸缩型数组成员(C99)
`伸缩型数组成员(flexible array member)`
第1个特性是，该数组不会立即存在。
第2个特性是，使用这个伸缩型数组成员可以缩写合适的代码，就像它确定存在并具有所需数目的元素
一样。
首先，声明一个伸缩型数组成员有如下规则：
伸缩型数组成员必须是结构的最后一个成员;
结构中必须至少有一个成员;
伸缩数组的声明类似于普通数组，只是它的方括号中是空的。

```c
struct flex
{
	int count;
	double average;
	double scores[];   // 伸缩型数组成员
};
```
带伸缩型数组成员的结构确实有一些特殊的处理要求。
第一，不能用结构进行赋值或拷贝：
```c
struct flex * pf1, * pf2;   // *pf1 和 *pf2都是结构

*pf2 = *pf1;   // 不要这样做
```
这样做只能拷贝除伸缩型数组成员以外的其他成员。确实要进行拷贝，应使用memcpy()函数.

第二，不要以按值方式把这种结构传递给结构。原因相同，按值传递一个参数与赋值类似。要把
结构的地址传递给函数

第三，不要使用带伸缩型数组成员的结构作为数组成员或另一个结构的成员。


#### 匿名结构(C11)
```c
struct names
{
	char first[20];
	char last[20];
};

struct person
{
	int id;
	struct names name; // 嵌套结构成员
};

struct person ted = { 8483, { "Ted", "Grass" } };
```
在C11中，可以用嵌套的匿名成员结构定义person:
```c
struct person
{
	int id;
	struct { char first[20]; char last[20]; };   // 匿名结构
};

struct person ted = { 8483, {"Ted", "Grass"} };

puts(ted.first);
```


#### 使用结构数组的函数

#### 把结构内容保存到文件中

