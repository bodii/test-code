#### 数字库
math.h头文件提供许多有用的数学函数。

ANSI C标准的一些数学函数:
| 原型 | 描述 |
|-|:-:|-:|
| double acos(double x) | 返回余弦值为x的角度(0~x弧度) |
| double asin(dobule x) | 返回正弦值为x的角度(-x/2~x/2弧度) |
| double atan(double x) | 返回正切值为x的角度(-pi/2~pi/2弧度) |
| double atan2(double y, double x) | 返回正弦值为y/x的角度(-pi～pi弧度) |
| double cos(double x) | 返回x的余弦值，x的单位为弧度 |
| double sin(double x) | 返回x的正弦值，x的单位为弧度 |
| double tan(double x) | 返回x的正切值，x的单位为弧度 |
| double exp(double x) | 返回x的指数函数的值(e^x) |
| double log(double x) | 返回x的自然对数值 |
| double log10(double x) | 返回x 的以10为底的对数值 |
| double pow(double x, double y) | 返回x的y次幂 |
| double sqrt(double x) | 返回x的平方值 |
| double cbrt(double x) | 返回x的立方值 |
| double ceil(double x) | 返回不小于x的最小整数值 |
| double fabs(double x) | 返回x的绝对值 |
| double floor(double x) | 返回不大于x的最大整数值 |

UNIX系统会要求使用-lm标记(flag)指示链接器搜索数学库：
cc rect_pol.c -lm
注意，-lm标记在命令行的末尾。因为链接器在编译器编译C文件后才开始处理。
在Linux中使用GCC编译器可能要这样写：
gcc rect_pol.c -lm


#### 类型变体
如果不需要双精度，那么用float类型的单精度值来计算会更快些。而且把long double
类型的值传递给double类型的形参会损失精度，形参获得的值可能不是原来的值。为了
解决这些潜在的问题，C标准专门为float类型和long double类型提供了标准函数，即在
原函数名前加上f或l前缀。因此，sqrtf()是sqrt()的float版本，sqrtl()是sqrt的long
double版本。


