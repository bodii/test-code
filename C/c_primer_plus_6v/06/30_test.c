#include <stdio.h>

int main(void)
{
	double f1, f2;

	while(scanf("%lf %lf", &f1, &f2) == 2)
	{
		printf("(%lf - %lf) / (%lf * %lf) = %.2lf.\n", 
				f1,f2,f1,f2, (f1 - f2) / (f1 * f2));
	}

	printf("bye!\n");

	return 0;
}
