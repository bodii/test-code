#### #undef指令
#undef指令用于“取消”已定义的#define指令。
```c
#define LIMIT 400
// 然后，下面的指令
#undef LIMIT
```
将移除上面的定义。现在就可以把LIMIT重新定义为一个新值。即使原来没有定义LIMIT,
取消LIMIT的定义仍然有效。如果想使用一个名称，又不确定之前是否已经用过，为安全
起见，可以用#undef指令取消该名字的定义。

处理器在识别标识符时，遵循与C相同的规则：标识符可以由大写字母、小写字母、数字
和下划线字符组成，且首字符不能是数字。

已定义宏可以是对象宏，包括空宏或类函数宏：
```c
#define LIMIT 1000     // LIMIT 是已定义的
#define GOOD           // GOOD是已定义的
#define A(X) ((-(X))*(X))   // A是已定义的
int q;                 // q不是宏，因为是未定义的
#undef GOOD            // GOOD取消定义，是未定义的

#define宏的作用域从它在文件中的声明处开始，直到用#undef指令取消宏为止，或延伸至
文件尾(以二者中先满足的条件作为宏作用域的结束).


#### 条件编译(conditinal compilation)
可以使用这些指令告诉编译器根据编译时的条件执行或忽略信息(或代码)块。

1. #ifdef、#else和#endif指令

```c
#ifdef MAVIS
#include "horse.h"     // 如果已经用#define定义了MAVIS,则执行下面的指令
#define STABLES 5
#else
#include "cow.h"       // 如果没有用#define定义MAVIS,则执行下面的指令
#define STABLES 15
#endif
```
#ifdef指令说明，如果预处理器已定义了后面的标识符(MAVIS),则执行#else或#endif指令
之前的所有指令并编译所有C代码。

#ifdef #else很像C的if else.两者的主要区别是，预处理器不识别用于标记块的花括号，
这些指令结构可以嵌套。也可以用这些指令标记C语句块。

2. #ifndef指令
#ifndef指令与#ifdef指令的用法类似，也可以和#else、#endif一起使用，但是它们的逻辑
相反。#ifndef指令判断后面的标识符是否是未定义的，常用于定义之前未定义的常量。
```c
#ifndef SIZE
#define SIZE 100
#endif
```
通常，包含多个头文件时，其中的文件可能包含了相同宏定义。#ifndef指令可以防止相同的
宏被重复定义。在首次定义一个宏的头文件中用#ifndef指令激活定义。随后在其他头文件中
的定义都被忽略。

3. #if和#elif指令
#if指令很像C语言中的if。#if后面跟整型常量表达式，如果表达式为非零，则表达式为真。
```c
#if SYS == 1
#include "ibm.h"
#endif
```
可以按照if else的形式使用#elif
```c
#if SYS == 1
#include "ibmpc.h"
#elif SYS == 2
#include "vax.h"
#elif SYS == 3
#include "mac.h"
#else
#include "general.h"
#endif

较新的编译器提供另一种方法测试名称是否已定义，即用#if defined(VAX)代替#ifdef VAX.
这里，defined是一个预处理运算符，如果它的参数是用#define定义过，则返回1;否则返回0.
这种新方法的优点是，它可以和#elif一起使用。
```c
#if defined(IMPC)
#include "imbpc.h"
#elif defined(VAX)
#include "vax.h"
#elif defined(MAC)
#include "mac.h"
#else
#include "general.h"
#endif

