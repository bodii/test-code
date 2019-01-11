#include <stdio.h>

int sum(int, int);
double doubleSum(double, double);
void alter(int *, int *);

int main(void)
{
	int a = 2, b = 5;
	double x = 1.223, y = 5.670;
	int a1 = 4, b1 = 3;

	printf("%d\n", sum(a, b));
	printf("%f\n", doubleSum(x, y));
	alter(&a1, &b1);
	printf("%d\t\t%d\n", a1, b1);

	return 0;
}


int sum(int a, int b)
{
	return a + b;
}


double doubleSum(double x, double y)
{
	return x + y;
}


void alter(int * a, int * b)
{
	int temp1, temp2;

	temp1 = *a + *b;
	temp2 = *a - *b;
	*a = temp1;
	*b = temp2;
}
