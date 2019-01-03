#### 转换说明(conversion specification)
它们指定了如何把数据转换成可显示的形式
printf()函数打数据的指令要与待打印数据的类型相匹配。

| 转换说明 | 输出 |
|-|:-:|-:|
| %a | 浮点数、十六进制数和p记数法(C99/C11) |
| %A | 浮点数、十六进制数和p记数法(C99/C11) |
| %c | 单个字符 |
| %d | 有符号十进制整数 |
| %e | 浮点数，e记数法 |
| %E | 浮点数，e记数法 |
| %f | 浮点数，十进制记数法 |
| %g | 根据值的不同，自动选择%f或%e。%e格式用于指数小于-4或者大于或等于精度时 |
| %G | 根据值的不同，自动选择%f或%E。%E格式用于指数小于-4或者大于或等于精度时 |
| %i | 有符号十进制整数（与%d相同) |
| %o | 无符号八进制整数 |
| %p | 指针 |
| %s | 字符串 |
| %u | 无符号十进制整数 |
| %x | 无符号十六进制整数，使用十六进制数0f |
| %X | 无符号十六进制整数，使用十六进制数0F |
| %% | 打印一个百分号 |

