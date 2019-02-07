#### 通用工具库
通用工具库包含各种函数，包括随机数生成器、查找和排序函数、转换函数和内存管理函数。
在ANSI C标准中，这些函数的原型都在stdlib.h头文件中。


#### exit()和atexit()函数
atexit()函数通过退出时注册被调用的函数提供这种功能，atexit()函数接受一个函数指针作
为参数。

1. atexit()函数的用法
这个函数使用函数指针。要使用atexit()函数，只需把退出时要调用的函数地址传递给atexit()
即可。函数名作为函数参数时相当于该函数的地址，所以该程序中把sign_off或too_bad作为参数。
然后，atexit()注册函数列表中的函数，当调用exit()时就会执行这些函数。在这个列表中至少
可以放32个函数。最后调用exit()函数时，exit()会执行这些函数(执行顺序与列表中的函数顺序
相反，即最后添加的函数最先执行)。

注意，输入失败时，会调用sign_off()和too_bad()函数;但是输入成功时只会调用sign_off().
因为只有输入失败时，才会进入if语句中注册too_bad().另外还要注意，最先调用的是最后一个
被注册的函数。

atexti()注册的函数(如sign_off()和too_bad())应该不带任何参数且返回类型为void.通常，这些
函数会执行一些清理任务，例如更新监视程序的文件或重置环境变量。

注意，即使没有显式调用exit()，还是会调用sign_off(),因为main()结束时会隐式调用exit().

2. exit()函数的用法
exit()执行完atexit()指定的函数后，会完成一些清理工作:`刷新所有输出流、关闭所有打开的流
和关闭由标准I/O函数tmpfile()创建的临时文件。然后exit()把控制权返回主机环境，如果可能的
话，向主机环境报告终止状态。`
通常，UNIX程序使用0表示成功终止，用非零值表示终止失败。UNIX返回的代码并不适用于所有的
系统，所以ANSI C为了可移植性的要求，定义了一个名为EXIT_FAILURE的宏表示终止失败。
类似地，ANSI C还定义了EXIT_SUCCESS表示成功终止。不过，exit()函数也接受0表示成功终止。
在ANSI C中，在非递归的main()中使用exit()函数等价于使用关键字return.尽管如此，在main()
以外的函数中使用exit()也会终止整个程序。


##### qsort()函数
对较大型的数组而言，"快速排序"方法是最有效的排序算法之一。
首先，把数组分成两部分，一部分的值都小于另一部分的值。这个过程一直持续到数组完成排序好
为止。

快速排序算法在C实现中的名称是qsort().qsort()函数排序数组的数组对象：
```c
void qsort(void * base, size_t nmemb, size_t size, int (* compar)(const void *, const void *));
```
ANSI C允许把指向任何数据类型的指针强制换成指向void的指针。

强制类型转换：
```c
const double * a1 = (const double *) p1;
```

下面再来看一个比较函数的例子。假设有下面的声明：
```c
struct names {
	char first[40];
	char last[40];
};

struct names staff[100];

// 调用qsort()
qsort(staff, 100, sizeof(struct names), comp);

// 这里comp是比较函数的函数名。
#include <string.h>
int comp(const void * p1, const void * p2)
{
	/* 得到正确类型的指针 */
	const struct names * ps1 = (const struct names *) p1;
	const struct names * ps2 = (const struct names *) p2;

	int res;
	res = strcmp(ps1->last, ps2->last);  
	if (res != 0 )
		return res;
	else
		return strcmp(ps1->first, ps2->first);
}
```
该函数使用strcmp()函数进行比较。strcmp()的返回值与比较函数的要求相匹配。
通过指针访问结构成员时必须使用->运算符。
