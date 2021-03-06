### C语言的数据类型关键字

> 在C语言中，用int关键字来表示基本的整数类型。后3个关键字(long、short和
> unsigned)和C90新增的signed用于提供基本整数类型的变式，例如unsigned short
> int 和long long int。char关键字用于指定字母和其他字符(如，#、$、%和*）。
> 另外，char类型也可以表示较小的整数。float、double和long double表示带小数
> 点的数。_Bool类型表示布尔值(true或false), _complex和_Imaginary分别表示复数
> 和虚数

通过这些关键字创建的类型，按计算机的储存方式可分为两大基本类型： 整数类型和浮点
类型。



#### 位、字节和字

位、字节和字是描述计算机数据单元或存储单元的术语。这里主要指存储单元。
* 最小的存储单元是位（bit), 可以储存0或能1（或者说，位用于设置“开”或“关”），
虽然1位储存的信息有限，但是计算机中位的数量十分庞大。位是计算机内存的基本
构建块。

* 字节（byte）是常用的计算机存储单位。对于几乎所有的机器，1字节均为8位。这是
字节的标准定义，至少在衡量存储单位时是这样（但是，C语言对此有不同的定义）。
既然1位可以表示0或1,那么8位字节就有256(2的8次方）种可能的0、1的组合。通过二
进制编码（仅0和1便可表示数字），便可表示0~255的整数或一组字符。

* 字（word）是设计计算机时给定的自然存储单位。对于8位的微型计算机，1个字长只有
8位。从那以后，个人计算机字长增至16位、32位，直到目前的64位。计算机的字长越大，
其数据转移越快，允许的内存访问也更多。
