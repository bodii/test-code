#### 预定义宏
C标准规定了一些预定义宏：
| 宏 | 含义 |
|-|:-:|-:|
| __DATE__ | 预处理的日期("Mmm dd yyyy"形式的字符串字面量，如Nov 23 2013) |
| __FILE__ | 表示当前源代码文件名的字符串字面量 |
| __LINE__ | 表示当前源代码文件中行号的整型常量 |
| __STDC__ | 设置为1时，表明实现遵循C标准       |
| __STDC_HOSTED__ | 本机环境设置为1: 否则设置为0 |
| __STDC_VERSION__ | 支持C99标准，设置为199901L; 支持C11标准，设置为201112L |
| __TIME__ | 翻译代码的时间，格式为"hh:mm:ss" | 

C99 标准提供一个名为__func__的预定义标识符，它展开为一个代表函数名的字符串(该
函数包含该标识符).那么，__func__必须具有函数作用域，而从本质上看宏具有文件作用
域。因此，__func__是C语言的预定义标识符，而不是预定义宏。

如上表的一些预定义宏和预定义标识符。注意，其中一些是C99新增的，所以不支持C99的
编译器可能无法识别它们。如果使用GCC，必须设置`-std=c99`或`-std=c11`。
