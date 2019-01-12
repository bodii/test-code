// 指针地址

#include <stdio.h>
#define SIZE 4

int main(void)
{
	short dates[SIZE];
	short * pti;
	short index;
	double bills[SIZE];
	double * ptf;
	pti = dates;  // 把数组地址赋给指针
	ptf = bills;

	printf("%23s %15s\n", "short", "double");
	for (index = 0; index < SIZE; index++)
		printf("pointers + %d: %10p %10p\n", index, pti + index, ptf + index);
	/*
	 这里打印的是两个数组开始的地址，下一行打印的是指针加1后的地址，以此类推。
	 */

	return 0;
}

