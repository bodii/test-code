#include <stdio.h>

void exchangeValue(double *, double *, double *);

int main(void)
{
	double a = 1.63, b = 1.44, c = 1.7355;

	printf("old: %f, %f, %f\n", a, b, c);
	exchangeValue(&a, &b, &c);
	printf("now: %f, %f, %f\n", a, b, c);
}

void exchangeValue(double * a, double * b, double * c)
{
	double min, middle, max;

	min = *a > *b ? *b : *a;
	min = min < *c ? min : *c;

	middle = *a < *b ? *b : *a;
	middle = middle < *c ? middle : *c; 

	max  = *a > *b ? *a : *b;
	max = *c < max ? max : *c;
	
	*a = min;
	*b = middle;
	*c = max;
}


