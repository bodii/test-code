#### 存储类别和函数
函数也有存储类别，可以是外部函数(默认)或静态函数。C99新增了第3种类别--
内联函数。

外部函数可以被其他文件的函数访问，但是静态函数只能用于其定义所在的文件。
假设一个文件中包含了以下函数原型:
```c
double gamma(double);     // 该函数默认为外部函数
static double beta(int, int);
extern double delta(double, int);
```
在同一个程序中，其他文件中的函数可以调用gamma()和delta(),但是不能调用beta(),
因为以static存储类别说明符创建的函数属于特定模块私有。这样做避免了名称冲突的
问题，由于beta()受限于它所在的文件，所以在其他文件中可以使用与之同名的函数

通常的做法是：用extern关键字声明定义在其他文件中的函数。这样做是为了表明当前
文件中使用的函数被定义在别处。除非使用static关键字，否则一般函数声明都默认为
extern.
