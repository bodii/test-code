#include <stdio.h>

int main(void)
{
	int n, n2, n3;

	n = 5;
	n2 = n * n;
	n3 = n2 * n2; // 这里有语义错误 立方应为n2 * n 或 n * n * n
	printf("n = %d, n squared = %d, n cubed = %d \n", n, n2, n3);

	return 0;
}
