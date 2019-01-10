// 错误地使用函数
#include <stdio.h>

int imax();  /* 旧式函数声明 */

int main(void)
{
	printf("The maximum of %d and %d is %d.\n", 3, 5, imax(3));
	printf("The maximum of %d and %d is %d.\n", 3, 5, imax(3.0, 5.0));

	return 0;
}

int imax(n, m)
{
	return (n > m) ? n : m;
}
