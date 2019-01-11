#include <stdio.h>

double reciprocal(double, double);

int main(void)
{
	double a = 3.1323, b = 4.2432;

	printf("%f\n", reciprocal(a, b));

	return 0;
}

double reciprocal(double a, double b)
{
	a = 1 / a;
	b = 1 / b;

	return 1 / ( a + b );
}
	
