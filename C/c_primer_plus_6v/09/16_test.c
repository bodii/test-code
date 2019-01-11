#include <stdio.h>

int returnMax(int, int, int); 


int main(void)
{
	int a = 3, b = 5, c = 10;

	printf("%d, %d, %d, The largest number is : %d\n", 
			a, b, c, returnMax(a, b, c));

	return 0;
}


int returnMax(int a, int b, int c)
{
	if ( a > b && a > c)
		return a;
	else if ((a > b && a < c) || (a < b && b < c))
		return c;
	else if (a > c && a < b || (a < c && c < b))
		return b;
}
