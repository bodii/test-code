#### 复数和虚数类型

许多科学和工程计算都要用到复数和虚数。C99标准支持复数和虚数类型，
但是有所保留。一些独立实现，如嵌入式处理器的实现，就不需要使用复
数和虚数(VCR芯片就不需要复数)。一般而言，虚数类型都是可选项。
C11标准把整个复数软件包都作为可选项。
> 简而言之，C语言有3种复数类型：`float_Complex`、`double_Complex`
> 和`long double_Complex`。例如，float_Complex类型的变量包含两个
> float类型的值，分别表示复数的实部和虚部。

> 类似地，C语言的3种虚数类型是`float _Imaginary`、`double _Imaginary`
> 和`long double _Imaginary`

> 如果包含complex.h头文件，便可用complex代替_Complex,用imaginary代
> 代替_Imaginary,还可以用I代替-1的平方根。
