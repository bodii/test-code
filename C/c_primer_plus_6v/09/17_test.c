#include <stdio.h>

double min(double x, double y)
{
	return x > y ? y : x;
}


int main(void)
{
	double x = 2.3, y =1.666809;

	printf("%f,%f, The smallest number is: %f\n",
			x, y, min(x, y));

	return 0;
}
