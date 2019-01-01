#include <stdio.h>

/*
 * 编写一个把英寸单位转换为厘米单位(1英寸=2.54厘米)的程序
 */

int main(void)
{
	float inches;
	float centimeter;

	printf("Please enter an inch number:\n"); 
	scanf("%f", &inches);

	printf("= %.2f cm\n", inches * 2.54);

	return 0;
}
