#### 函数和指针
假设有一个指向int类型变量的指针，该指针储存着这个int类型变量储存在内存位置的地址。
同样，函数也有地址，因为函数的机器语言实现由载入内存的代码组成。指向函数的指针中
储存着函数代码的起始处的地址。
其次，声明一个数据指针时，必须声明指针所指向的数据类型。声明一个函数指针时，必须
声明指针指向的函数类型。为了指明函数类型，要指明函数签名，即函数的返回类型和形参
类型。
```c
void ToUpper(char *);
```
ToUpper()函数的类型是"带char *类型参数、返回类型是void的函数".
声明一个指针pf指向该函数类型：
```c
void (*pf)(char *);    // pf是一个指向函数的指针
```

要声明一个指向特定类型函数的指针，可以先声明一个该类型的函数，然后把函数名替换成
(*pf)形式的表达式。然后，pf就成为指向该类型函数的指针。

声明了函数指针后，可以把类型匹配的函数地址赋给它。
既然可以用数据指针访问数据，也可以用函数指针访问函数。
```c
void ToUpper(char *);

void (*pf)(char *);

char mis[] = "Nina Metier";

pf = ToUpper;

(*pf)(mis); // (语法1)

pf(mis); // (语法2)
```
第1种方法：由于pf指向ToUpper函数，那么*pf就相当于ToUpper函数，所以表达式(*pf)(mis)和
ToUpper(mis)相同.
第2种方法: 由于函数名是指针，那么指针和函数名可以互换使用，所以pf(mis)和ToUpper(mis)
相同。

```c
void show (void (* fp)(char *), char * str);
```
str是一个数据指针。fp形参是一个函数指针。
fp指向的函数接受char * 类型的参数，其返回类型为void; str指向一个char类型的值。
因此，可以这样调用函数：
```c
show(ToLower, mis);    /* show()使用ToLower()函数: fp = ToLower */
show(pf, mis);         /* show()使用pf指向的函数：fp = pf */

void show (void (* fp)(char *), char * str)
{
	(*fp)(str);  /* 把所选函数作用于str */
	puts(str);   /* 显示结果 */
}
```

把带返回值的函数作为参数传递给另一个函数有两种不同的方法：
```c
function1(sqrt);   /* 传递sqrt()函数的地址 */
function2(sqrt(4.0));  /* 传递sqrt()函数的返回值 */
```
