#### 其他赋值运算符：+=、-=、*=、/=、%=

scores += 20   与   scores = scores + 20   相同
dimes -= 2     与   dimes = dimes - 2      相同
bunnies *= 2   与   bunnies = bunnies * 2  相同
time /= 2.73   与   time = time / 2.73     相同
reduce %= 3    与   reduce = reduce % 3    相同


#### 逗号运算符
逗号运算符扩展了for 循环的灵活性，以便在循环头中包含更多的表达式。

逗号运算符有两个其他性质：
* 首先，它保证了被它分隔的表达式从左往右求值（换言之，逗号是一个序列
点，所似逗号左侧项的所有副作用都在程序执行逗号右侧项之前发生).


#### 当Zeno 遇到for循环


#### 使用带返回值的函数
> 编译器在程序中首次遇到power()时，需要知道power()的返回类型。此时，编译
器尚未执行到power()的定义，并不知道函数定义中的返回类型是double.因此，
必须通过`前置声明(forward declaration)`预先说明函数的返回类型。前置声明
告诉编译器，power()定义在别处，其返回类型为double.如果把power()函数的定
义置于main()的文件顶部，就可以省略前置声明，因为编译器在执行到main()之前
已经知道power()的所有信息。但是，这不是C的标准风格。因为main()通常只提供
整个程序的框架，最好把main()放在所有函数定义的前面。另外，通常把函数放在
其他文件中，所以前置声明必不可少。

> 函数原型是为了方便编译器查看程序中使用的函数是否正确，函数定义描述了如
何工作。现代的编程习惯是把程序要素分为接口部分和实现部分，例如函数原型和
函数定义。接口部分描述了如何使用一个特性，也就是函数原型所做的;实现部分
描述了具体的行为，这正是函数定义所做的。

