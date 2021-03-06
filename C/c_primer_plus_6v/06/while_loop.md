#### C风格读取循环
如：
```c
status = scanf("%ld", &num);
while (status == 1)
{
	/* 循环行为 */
	status = scanf("%ld", &num);
}
```

可以用这些代码替换:
```c
while (scanf("%ld", &num) == 1)
{
	/* 循环行为 */
}
```

第二种形式同时使用scanf()的两种不同的特性。首先，如果函数调用成功，
scanf()会把一个值存入num。然后，利用scanf()的返回值(0或1,不是num的
值)控制while循环。因为每次迭代都会判断循环的条件，所以每次迭代都要
调用scanf()读取新的num值来判断。换句话说，C的语法特性让你可以用下面
的精简版本替换标准版本:
`
当获取值和判断值都成功
	处理该值
`


#### while语句
通用形式如下:
while ( expression )
	statement

statement部分可以是以分号结尾的简单语句，也可以是用花括号括起来的复合
语句。

expression是值之间的比较，可以使用任何表达式。


#### while:入口条件循环
while循环是使用入口条件的有条件循环。
所谓“有条件”指的是语句部分的执行取决于测试表达式描述的条件，如(index < 5).
该表达式是一个`入口条件(entry condition)`,因为必须满足条件才能进入循环体。


#### while无花括号时
语句从while开始执行，到第1个分号结束.


#### 空语句(null statement)
在C语言中，单独的分号表示空语句。
```c
while (scanf("%d", &num) == 1)
	; /* 跳过整数输入 */
```

#### 用关系运算符和表达式比较大小
while循环经常依赖测试表达式作为比较，这样的表达式被称为`关系表达式(relational expression)`,
出现在关系表达式中间的运算符叫做`关系运算符(relational operator)`。
关系运算符：
运算符: <、<=、==、>=、>、!= 


#### 新的_Bool类型
在C语言中，一直用int类型的变量表示真/假值。C99专门针对这种类型的变量新增了_Bool类型。
_Bool类型的变量只能储存1(真)或0(假).如果把其他非零数值赋给_Bool类型的变量，该变量会
被设置为1.
> C99提供了stdbool.h头文件，该头文件让bool成为_Bool的别名，而且还把true和false分别定义
为1和0的符号常量。包含该头文件后，写出的代码可以与C++兼容，因为C++把bool、true和false
定义为关键字。


#### 优先级和关系运算符
关系运算符的优先级比算术运算符(包括+和-)低，比赋值运算符高。


#### 不确定循环和计数循环
一些while循环是`不确定循环(indefinite loop)`。所谓不确定循环，指在测试表达式为假之前，
预先不知道要执行多少次循环。

`计数循环(counting loop)`,这类循环在执行循环之前就知道要重复多少次。


#### do while
形式：
`
do
	statement
while ( expression );
`
在test为假或0之前，重复执行statement部分。
示例：
```c
do
	scanf("%d", &number);
while (number != 20);
```
