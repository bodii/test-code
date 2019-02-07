#### tgmath.h库(C99)
C99标准提供的tgmath.h头文件中定义了泛类型宏。
如果在math.h中为一个函数定义了3种类型(float、double和long double)的版本，那么tgmath.h文件就创建一个
泛型类型宏，与原来double版本的函数名同名。例如，根据提供的参数类型，定义sqrt()宏展开为sqrtf()、sqrt
或sqrtl()函数。

如果编译器支持复数运算，就会支持complex.h头文件，其中声明了与复数运算相关的函数。例如，声明有csqrtf()
、csqrt()和csqrtl(),这些函数分别返回float complex、double complex和long double complex类型的复数平方根。
如果提供这些支持，那么tgmath.h中的sqrt()宏也能展开为相应的复数平方根函数。

如果包含了tgmath.h，要调用sqrt()函数而不是sqrt()宏，可以用圆括号把被调用的函数名括起来：
```c
#include <tgmath.h>
...
float x = 44.0;
double y;
y = sqrt(x);     // 调用宏，所以是sqrtf(x)
y = (sqrt)(x);   // 调用函数sqrt()
```
类函数宏的名称必须用圆括号括起来。
由于C语言奇怪而矛盾的函数指针规则，也可以使用(*sqrt)()的形式来调用sqrt()函数

不借助C标准以外的机制，C11新增的_Generic表达式是实现tgmath.h最简单的方式。
