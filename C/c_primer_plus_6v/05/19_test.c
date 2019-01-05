/* 打印a~g */
#include <stdio.h>
#define SEVEN 7

int main(void)
{
	int n = 0;

	while (n++ < SEVEN)
		printf("%-5c", n + 96);
	printf("\n");

	return 0;
}
