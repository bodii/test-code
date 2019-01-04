// 使用变宽输出字段
#include <stdio.h>

int main(void)
{
	unsigned width, precision;
	int number = 256;
	double weight = 242.5;

	printf("Enter a field width:\n");
	scanf("%d", &width);
	printf("The number is :%*d:\n", width, number);
	printf("Now enter a width and a precision:\n");
	scanf("%d %d", &width, &precision);
	printf("Weight = %*.*f\n", width, precision, weight);
	printf("Done!\n");

	return 0;
}

/* 这里，用户首先输入6,因此6是程序使用的字段宽度。类似地，接下来用户
 * 输入8和3,说明字段宽度是8，小数点后面显示3位数字。一般而言，程序应
 * 根据weight的值来决定这些变量的值。
 *
 * scanf()中*的用法与此不同.把*放在%和转换字符之间时，会使得scanf()跳
 * 过相应的输出项。
 */
