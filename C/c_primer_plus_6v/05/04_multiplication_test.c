/* 计算1~20的平方 */

#include <stdio.h>

int main(void)
{
	int num = 1;

	while (num < 12)
	{
		printf("%4d %6d\n", num, num * num);
		num = num + 1;
	}

	return 0;
}
