#include <stdio.h>

#define I 25
#define SPACE ' '
#define PS() printf(SPACE)
#define BIG(X) (X+3)
#define SUMSQ(X, Y) ((X * X) + (Y * Y))

int main(void)
{
	printf("%d\n", SUMSQ(3, 5));

	return 0;
}
