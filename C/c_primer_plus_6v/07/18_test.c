#include <stdio.h>
#include <stdbool.h>

bool isPrime(int n);

int main(void)
{

	unsigned int i;

	while (scanf("%d", &i) == 1)
	{
		for (int n = 1; n <= i; n++)
		{
			if( isPrime(n) )
			{
				printf("%d, ", n);
			}
		}
		printf("\n");

	}

	printf("Bye!\n");

	return 0;
}


/* 判断一个数是不是素数 */
bool isPrime(int n)
{
	if (n < 2)
		return false;

	for (int i = 1; i < n; i++) 
	{
		if (n % i == 0)
		{
			if (n == i)
				return true;
			else if (i != 1)
				return false;
		}
	}

	return true;
}



