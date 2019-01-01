#include <stdio.h>

/*
 编写一个程序，调用一个名为one_three()的函数。该函数
 在一行打印单调"one", 再调用第2个函数two(),然后在另一
 行打印单调“three”.
 two()函数在一行显示“two”.
 main()函数在调用one_three()函数前要打印短语“starting now:”，
 并在调用完毕后显示短语“done!”。
 */

// one_three函数的声明
void one_three(void);
// two 函数的声明
void two(void);

// two函数的定义
void two(void)
{
	printf("two\n");
}

// one_three函数的定义
void one_three(void)
{
	printf("one\n");
	two();
	printf("three\n");
}

int main(void)
{
	printf("starting now:\n");
	one_three();
	printf("done!\n");

	return 0;
}

