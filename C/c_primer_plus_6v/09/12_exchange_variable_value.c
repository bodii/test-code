// 第1 个版本的交换函数 -- 错误的方法
#include <stdio.h>

void interchange(int u, int v); // 声明函数(函数原型)

int main(void)
{
	int x = 5, y = 10;

	printf("Originally x = %d and y = %d.\n", x, y);
	interchange(x, y);
	printf("Now x = %d and y = %d.\n", x, y);

	return 0;
}


void interchange(int u, int v) // 定义函数
{
	int temp;

	printf("In interchange(): Originally u = %d and v = %d.\n", u, v);
	temp = u;
	u = v;
	v = temp;
	printf("In interchange(): Originally u = %d and v = %d.\n", u, v);
	// 没有将交换的结果传给main()主调函数
	// 使用return语句只能把被调函数中的一个值传回主调函数，但是现在要传回两个值。
	// 要使用指针
}
