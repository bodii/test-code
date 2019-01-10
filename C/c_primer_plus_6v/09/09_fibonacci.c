// 斐波那契数列

#include <stdio.h>

/* 斐波那契数列函数 - 函数原型 */
unsigned long Fibonacci(unsigned n);

int main(void)
{
	unsigned n;

	printf("Enter an integer (q to quit):\n");

	while (scanf("%d", &n) == 1)
	{
		printf("result: ");
		printf("%ld", Fibonacci(n));
		putchar('\n');

		printf("Enter an integer (q to quit):\n");
	}

	return 0;
}


/* 斐波那契数列函数 - 函数定义 */
unsigned long Fibonacci(unsigned n)
{
	if (n > 2)
		return Fibonacci(n-1) + Fibonacci(n-2);
	else
		return 1;

}
