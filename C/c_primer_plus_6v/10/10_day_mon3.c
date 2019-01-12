// 使用指针
#include <stdio.h>
#define MONTHS 12

int main(void)
{
	int days[MONTHS] = { 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31 };
	int index;

	for (index = 0; index < MONTHS; index++)
		printf("Month %2d has %d days.\n",
				index + 1,
				*(days + index)); // 与days[index]相同
	/*
	 这里，days是数组首元素的地址，days + index是元素days[index]的地址，而*(days 
	 + index)则是该元素的值，相当于days[index]。for循环依次引用数组中的每个元素，
	 并打印各元素的内容。
	 */
	
	return 0;
}
