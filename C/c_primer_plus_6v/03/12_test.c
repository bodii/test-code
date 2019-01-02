#include <stdio.h>

/* 练习题： 
 * 1个水分子的质量约为3.0X10^-23克。1夸脱水大约是950克。
 * 编写，提示用户输入水的夸脱数，并显示水分子的数量。
 */

int main(void)
{
	unsigned int u_quart_num;
	unsigned short int us_grams;
	double u_quality;

	u_quality = 3.0e-23;
	us_grams = 950;

	printf("Please enter the quart value of a water: ");
	scanf("%d", &u_quart_num);
	printf("\nWater molecular weight is : %e\n", u_quart_num * us_grams / u_quality);

	return 0;
}
