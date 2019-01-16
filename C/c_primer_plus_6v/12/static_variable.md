#### 块作用域的静态变量
静态变量(static variable)，是该变量在内存中原地不动，并不是说它的值不变。
具有文件作用域的变量自动具有(也必须是)静态存储期。

这些变量和自动变量一样，具有相同的作用域， 但是程序离开它们所在的函数后，这
些变量不会消失。也就是说，这种变量具有作用域、无链接，但是具有静态存储期。
计算机在多次函数调用之间会记录它们的值。在块中（提供块作用域和无链接）以存储
类别说明符static(提供静态存储期)声明的这种变量。


> 不能在函数的形参中使用static;
```c
int wontwork(static int flu);   // 不允许
```

"局部静态变量"是描述具有块作用域的静态变量的另一个术语。
在一些老的C文献时会发现，这种存储类别被称为`内部静态存储类别(internal static 
storage class)`.这里的内部指的是函数内部，而非内部链接。


####  外部链接的静态变量
外部链接的静态变量具有文件作用域、外部链接和静态存储期。
该类别有时称为`外部存储类别(external storage class)`，属于该类别的变量称为`外部
变量(external variable)`。把变量的定义性声明(defining declaration)放在所有函数的
外面便创建了外部变量。当然，为了指出该函数使用了外部变量，可以在函数中用关键字extern
再次声明。如果一个源代码文件使用的外部变量定义在另一个源代码中，则必须用extern在
该文件中声明该变量。

```c
int Errupt;     // 外部定义的变量
double Up[100];   // 外部定义的数组
extern char Coal;    // 如果Coalp被定义在另一个文件。则必须这样声明

void next(void);

int main(void)
{
	extern int Errupt;   // 可选的声明
	extern double Up[];   // 可选的声明
}


void next(void)
{
}
```

去掉下面声明中的extern:
```c
extern int Errupt;
```
便成为：
```c
int Errupt;
```
这使得编译器在main()中创建了一个名为Errupt的自动变量。它是一个独立的局部变量，与原来
的外部变量Errup不同。该局部变量仅main()中可见，但是外部变量Errupt对于该文件的其他函数
（如next())也可见。简而言之，在执行块中的语句时，块作用域中的变量将隐藏文件作用域串的
同名变量。如果不得已要使用与外部变量同名的局部变量同，可以在局部变量的声明中使用auto
存储类别说明符明确表达这种意图。

外部变量具有静态存储期。因此，无论程序执行到main()、next()还是其他函数，数组Up及其值都
一直存在。

示例1中有一个外部变量Hocus.该变量对main()和magic()均可见。
```c
/* 示例 1 */
int Hocus;
int magic();
int main(void)
{
	extern int Hocus;   // Hocus之前已声明为外部变量
}

int magic()
{
	extern int Hocus;   // 与上面的Hocus是同一个变量
}
```

示例2中有一个外部变量Hocus, 对两个函数均可见。
这次，在默认情况下对magic()可见:
```c
/* 示例 2 */
int Hocus;
int magic();
int main(void)
{
	extern int Hocus;  // Hocus 之前已声明为外部变量
}

int magic()
{
	// 并未在该函数中声明Hocus,但是仍可使用该变量
}
```

在示例3中，创建了4个独立的变量。main()中的Hocus变量默认是自动变量，属于main()私有。
magic()中Hocus变量被显示声明为自动，只有magic()，可用。外部变量Houcus对main()和magic
()均不可见，但是对该文件中未创建局部Hocus变量的其他函数可见。最后，Pocus是外部变量，
magic()可见，但是main()不可见，因为Pocus被声明在main()后面
```c
/* 示例 3 */
int Hocus;
int magic();
int main(void)
{
	int Hocus; // 声明 Houcs, 默认是自动变量
}

int Pocus;
int magic()
{
	auto int Hocus;  // 把局部变量Hocus显式声明为自动变量
}
```
