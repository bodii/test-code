#include <stdio.h>

/*
 编写一个程序，创建一个整形变量toes，将toes设置为10.程序中
 还要计算toes的两倍和toes的平方。该程序应打印3个值，并分别
 描述以示区分。
 */

int main(void)
{
	int toes;

	toes = 10;

	printf("2 times the integer 'toes' is: %d\n", toes * 2);
	printf("Integer square of 'toes' is: %d\n", toes * toes);

	return 0;
}
