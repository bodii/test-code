#### 基本数据类型
____


##### 关键字:
基本数据类型由于11个关键字组成：
> int、long、short、unsigned、char、float、double、
> signed、_Bool、_Complex和_Imaginary。


##### 有符号整型:
有符号整型可用于表示正整数和负整数。
* int  系统给定的基本整数类型。C语言规定int类型不小于16位。

* short或short int  最大的short类型整数小于或等于最大的int类型
整数。C语言规定short类型至少占16位。

* long 或long int  该类型可表示的整数大于或等于最大的int类型整数
。C语言规定long类型至少占32位。

* long long 或long long int  该类型可表示整数大于或等于最大的long
类型整数.long long类型至少占64位

一般而言，long类型占用的内存比short类型大，int类型的宽度要么和long
类型相同，要么和short类型相同。例如，旧DOS系统的PC提供16位的short
和int, 以及32位的long;windowns 95系统提供16位的short以及32位的int
和long。


##### 无符号整型:
无符号整型只能用于表示零和正整数，因此无符号整型可表示的正整数比有
符号整型的大。在整型类型前加上关键字unsigned表明该类型是无符号整型
：unsignedint、unsigned long、unsigned short。单独的unsigned相当于
unsignedint。


##### 字符类型:
可打印出来的符号(如A、&和+)都是字符。根据定义，char类型表示一个字符
要占用1字节内存。出于历史原因，1字节通常是8位，但是如果要表示基本字
符集，也可以是16位或更大。
* char   字符类型的关键字。有些编译器使用有符号的char,而有些则使用无
符号的char.在需要时，可在char前面加上关键字signed或unsigned来指明具
体哪一种类型。


##### 布尔类型:
布尔值表示ture和false。C语言用1表示true,0表示false.
* _Bool  布尔类型的关键字。布尔类型是无符号int类型，所占用的空间只要
能储存0或1即可。


##### 实浮点类型:
实浮点类型可表示正浮点和负浮点数。
* float  系统的基本浮点类型，可精确表示至少6位有效数字。
* double  储存浮点数的范围（可能）更大，能表示比float类型更多的有效数字
（至少10位，通常会更多）和更大的指数。
* long double   储存浮点数的范围（可能）比double更大，能表示比double更
多的有效数字和更大的指数。


##### 复数和虚数浮点数:
虚数类型是可选的类型。复数的实部和虚部类型都基于实浮点类型来构成：
* float _Complex
* double _Complex
* long double _Complex
* float _Imaginary
* double _Imaginary
* long long _Imaginary


##### 使用数据类型:
C编译器把浮点数转换成整数时，会直接丢弃（截断）小数部分，而不进行四舍五
入。
许多程序员和公司内部都有系统化的命名约定，在变量名中体现其类型。例如，用
i_前缀表示int类型，us_前缀表示unsigned short类型。这样，一眼就看出来i_smart
是int类型的变量，us_versmart是unsigned short类型的变量。


##### 参数和陷阱:
传递给函数的信息被称为参数。
C语言通过函数原型机制检查函数调用时参数的个数和类型是否正确。但是，该机制对
printf()和scanf()不起作用，因为这两个函数的参数个数可变。


##### 刷新输出
printf()何时把输出发送到屏幕上？
最初，printf()语句把输出发送到一个叫作缓冲区(buffer)的中间存储区域，然后缓冲
区中的内容再不断被发送到屏幕上。C标准明确规定了何时把缓冲区中的内容发送到屏
幕：当缓冲区满、遇到换行字符或需要输入的时候(从缓冲区把数据发送到屏幕或文件
被称为`刷新缓冲区`)。


##### 关键概念
C语言允许编写混合数据类型的表达式，但是会进行自动类型转换，以便在实际运算时
统一使用一种类型。


##### 数据类型
C有多种的数据类型。基本数据类型分为两大类：整数类型和浮点数类型。
通过为类型分配的储存量以及是有符号还是无符号，区分不同的整数类型。最小的整数
类型是char,因实现不同，可以是有符号char或无符号char，即unsigned char 或signed
char。
