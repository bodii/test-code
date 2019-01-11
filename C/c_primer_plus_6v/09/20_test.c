#include <stdio.h>

void larger_of(double *, double *);

int main(void)
{
	double a = 3.54, b = 3.66;

	printf("old: %f, %f\n", a, b);
	larger_of(&a, &b);
	printf("new: %f, %f\n", a, b);
}


void larger_of(double * a, double * b)
{
	*a = (*a > *b) ? *a : *b;
	*b = *a;
}
