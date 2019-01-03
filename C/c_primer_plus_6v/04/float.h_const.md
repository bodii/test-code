#### float.h头文件中的一些明示常量

| 明示常量 | 含义 |
|-|:-:|-:|
| FLT_MANT_DIG | float类型的尾数位数 |
| FLT_DIG | float类型的最少有效数字位数（十进制) |
| FLT_MIN_10_EXP | 带全部有效数字的float类型的最小负指数(以10为底) |
| FLT_MAX_10_EXP | float类型的最大正指数(以10为底) |
| FLT_MIN | 保留全部精度的float类型的最小正数 |
| FLT_MAX | float类型的最大正数 |
| FLT_EPSILON | 1.00和比1.00大的最小float类型值之间的差值 |

表中所列都与float类型相关。把明示常量名中FLT分别替换成DBL和LDBL，即
可分别表示double和long double类型对应的明示常量。
