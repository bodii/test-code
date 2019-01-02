#include <stdio.h>

/* 演示float舍入错误 */

int main(void)
{
	float a, b;

	b = 2.0e20 + 1.0;
	a = b - 2.0e20;

	printf("%f \n", a);

	return 0;
}

/*
	得到奇怪答案的原因是，计算机缺少足够的小数位来完成正确的运算。
	2.0e20是2后面有20个0。如果把该数加1,那么发生变化的是第21位。要
	正确运算，程序至少要储存21位数字，而float类型的数字通常只能储存
	按指数比例缩小或放大的6或7位有效数字。在这种情况下，计算结果一
	定是错误的。另一方面，如果把2.0e20改成2.0e4，计算结果就没有问题。
	因为2.0e4加1只需改变第5位上的数字，float类型的精度足够进行这样的
	计算。
 */
