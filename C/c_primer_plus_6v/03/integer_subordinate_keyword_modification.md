#### c语言整型关键字修饰

* short int 类型（或者简写为short）占用的存储空间可能比int类型少，常用
于较小数值的场合以节省空间。与int类似，short是有符号类型。

* long int 或long 占用的存储空间可能比int多，适用于较大数值的场合。与int
类似，long是有符号类型

* long long int 或 long long(C99标准加入) 占用的储存空间可能比long多，适用
于大数值的场合。该类型至少占用64位。与int类似，long long 是有符号类型。

* unsigned int 或 unsigned只用于非负值的场合。这种类型与有符号类型表示的范围
不同。例如，16位unsigned int 允许的取值范围是0~65535,而不是-32768~32767。用于
表示正负号的位现在用于表示另一个二进制位，所以无符号整型可以表示更大的数。

* 在C90标准中，添加了unsigned long int 或unsigned long 和unsigned int 或unsigned
 short类型。C99标准又添加了unsigned long long int 或unsigned long long。

 * 在任何有符号类型前面添加关键字unsigned,可强调使用有符号类型的意图。例如，
 short、short int、signed short、 signed short int 都表示同一种类型。


#### 整型的取值范围

> C标准对基本数据类型只规定了允许的最小大小。

> 对于16位机，short和int的最小取值范围是[-32767, 32767];

> 对于32位机，long的最小取值范围是[-2147483647, 2147483647];

> 对于unsigned short 和unsigned int, 最小取值范围是[0, 65535];

> 对于unsigned long,最小取值范围是[0, 4294967295]。

> long long类型是为了支持64位的需求，最小取值范围是[-9223372036854775807,
9223372036854775807];

> unsigned long long 的最小取值范围是[0, 18446744073709551615]。
如果要开支票，这个数是一千八百亿亿（兆）六千七百四十四万亿零七百
三十七亿零九百五十五万一千六百一十五。


#### long常量和long long 常量

要把一个较小的常量作为long类型对待,可以在值的末尾加上l(小写的L）或L
后缀。使用L后缀更好，因为l看上去和数字1很像。
因此， 在int为16位、long为32位的系统中，会把7作为16位储存，把7L作为32
位储存。l或L后缀也可用于八进制和十六进制整数，如020L和0x10L。
在支持long long类型的系统中，也可以使用ll或LL后缀来表示long long类型的值，
如3LL。另外，u或U后缀表示unsigned long long,如5ull、10LLU、6LLU或9Ull。


#### 整数溢出

当达到它能表示的最大值时，会重新从起始点开始。
主要区别是，在超过最大值时，unsigned int类型从0开始;而int类型则从-247483648
开始。
溢出行为是未定义的行为，C标准并未定义有符号类型的溢出规则。


#### 打印short、long、long long和unsigned类型

打印unsigned int类型的值，使用%u转换;
打印long类型的值，使用%ld转换;
如果系统中int和long的大小相同，使用%d就行。但是，这样的程序被植到其他系统
（int和long类型的大小不同）中会无法正常工作。在x和o前面可以使用l前缀，%lx
表示以十六进制格式打印long类型整数;
%lo表示以八进制格式打印long类型整数。注意，虽然C允许使用大写或小写的常量后
缀，但是在转换说明中只能用小写。
C语言有多种printf()格式。对于short类型，可以使用h前缀。%hd表示以十进制显示
short类型的整数，%ho表示以八进制显示short类型的整数。h和l前缀都可以和u一起
使用，用于表示无符号类型。例如，%lu表示打印unsigned long类型的值。
对于支持long long类型的系统，%lld和%llu分别表示有符号和无符类型。
