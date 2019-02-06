#### #pragma
#pragma把编译器指令放入源代码中。例如，在开发C99时，标准被称为C9X，可以
使用下面的编译指示(pragma)让编译器支持C9X:
```c
#pragma c9x on
```
一般而言，编译器都有自已的编译指示集。例如，编译指示可能用于控制分配给
自动变量的内存量，或者设置错误检查的严格程度，或者启用非标准语言特性等。
C99标准提供了3个标准编译指示。

C99还提供_Pragma预处理器运算符，该运算符把字符串转换成普通的编译指示:
```c
_Pragma("nonstandardtreatmenttypeB on")
// 等价于下面的指令：
#pragma nonstandardtreatmenttypeB on
```
由于该运算符不使用#符号，所以可以把它作为宏展开的一部分：
```c
#define PRAGMA(X) _Pragma(#X)
#define LIMRG(X) PRAGMA(STDC CX_LIMITED_RANGE X)
// 然后，可以使用类似下面的代码
LIMRG(ON)
```

下面的定义看上去没有问题，但实际上无法正常运行:
```c
#define LIMRG(X) _Pragma(STDC CX_LIMITED_RANGE #X)
```
问题在于这行代码依赖字符串的串联功能，而预处理过程完成之后才会串联字符串。
_Pragma运算符完成"解字符"(destringizing)的工作，即把字符串的转义序列转换成
它所代表的字符。因此，
```c
_Pragma("use_bool \"true \"false")
// 变成了:
#pragma use_bool "true "false
```

